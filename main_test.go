package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/database/mongodb"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/seed"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type authTestResponse struct {
	AccessToken string `json:"accessToken"`
}

type orderTestResponse struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	TotalCents int64  `json:"totalCents"`
}

type paymentTestResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func TestHTTPFluxoPedidoPagamentoEProtecoes(t *testing.T) {
	router := setupHTTPTest(t)

	unauthorized := doJSON(t, router, http.MethodGet, "/usuarios/me", "", nil)
	if unauthorized.Code != http.StatusUnauthorized {
		t.Fatalf("esperava 401 sem token, recebeu %d: %s", unauthorized.Code, unauthorized.Body.String())
	}

	clientToken := login(t, router, seed.ClientEmail, seed.ClientPassword)

	forbidden := doJSON(t, router, http.MethodPost, "/unidades", clientToken, map[string]any{
		"name":    "Unidade Bloqueada",
		"address": "Rua Teste",
		"city":    "Recife",
		"state":   "PE",
	})
	if forbidden.Code != http.StatusForbidden {
		t.Fatalf("esperava 403 para cliente criar unidade, recebeu %d: %s", forbidden.Code, forbidden.Body.String())
	}

	orderRecorder := doJSON(t, router, http.MethodPost, "/pedidos", clientToken, map[string]any{
		"clienteId":      seed.ClientID,
		"unidadeId":      seed.UnitID,
		"canalPedido":    "APP",
		"formaPagamento": "MOCK",
		"codigoPromocao": seed.PromotionCode,
		"itens":          []map[string]any{{"produtoId": seed.ProductCuscuzID, "quantidade": 1}},
	})
	if orderRecorder.Code != http.StatusCreated {
		t.Fatalf("esperava 201 ao criar pedido, recebeu %d: %s", orderRecorder.Code, orderRecorder.Body.String())
	}

	var orderResponse orderTestResponse
	decodeJSON(t, orderRecorder, &orderResponse)
	if orderResponse.ID == "" || orderResponse.Status != "AGUARDANDO_PAGAMENTO" || orderResponse.TotalCents <= 0 {
		t.Fatalf("pedido criado com resposta inesperada: %+v", orderResponse)
	}

	paymentRecorder := doJSON(t, router, http.MethodPost, "/pagamentos", clientToken, map[string]any{
		"pedidoId":   orderResponse.ID,
		"valorCents": orderResponse.TotalCents,
		"aprovado":   true,
	})
	if paymentRecorder.Code != http.StatusCreated {
		t.Fatalf("esperava 201 ao pagar pedido, recebeu %d: %s", paymentRecorder.Code, paymentRecorder.Body.String())
	}

	var paymentResponse paymentTestResponse
	decodeJSON(t, paymentRecorder, &paymentResponse)
	if paymentResponse.ID == "" || paymentResponse.Status != "APROVADO" {
		t.Fatalf("pagamento criado com resposta inesperada: %+v", paymentResponse)
	}

	paidOrderRecorder := doJSON(t, router, http.MethodGet, "/pedidos/"+orderResponse.ID, clientToken, nil)
	if paidOrderRecorder.Code != http.StatusOK {
		t.Fatalf("esperava 200 ao buscar pedido pago, recebeu %d: %s", paidOrderRecorder.Code, paidOrderRecorder.Body.String())
	}

	var paidOrder orderTestResponse
	decodeJSON(t, paidOrderRecorder, &paidOrder)
	if paidOrder.Status != "PAGO" {
		t.Fatalf("esperava pedido PAGO, recebeu %+v", paidOrder)
	}

	conflict := doJSON(t, router, http.MethodPost, "/pedidos", clientToken, map[string]any{
		"clienteId":      seed.ClientID,
		"unidadeId":      seed.UnitID,
		"canalPedido":    "TOTEM",
		"formaPagamento": "MOCK",
		"itens":          []map[string]any{{"produtoId": seed.ProductTapiocaID, "quantidade": 999}},
	})
	if conflict.Code != http.StatusConflict {
		t.Fatalf("esperava 409 para estoque insuficiente, recebeu %d: %s", conflict.Code, conflict.Body.String())
	}
}

func setupHTTPTest(t *testing.T) *gin.Engine {
	t.Helper()
	_ = godotenv.Load()
	gin.SetMode(gin.TestMode)

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Skipf("MongoDB indisponivel para teste de integracao: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		t.Skipf("MongoDB indisponivel para teste de integracao: %v", err)
	}
	_ = client.Disconnect(ctx)

	os.Setenv("MONGODB_DATABASE", "raizes_do_nordeste_test")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_EXPIRES_IN", "3600")

	mongodb.InitConnection()
	database := mongodb.GetDatabase()
	if err := database.Drop(ctx); err != nil {
		t.Fatalf("erro ao limpar banco de teste: %v", err)
	}
	for range 2 {
		if err := seed.Run(ctx, database); err != nil {
			t.Fatalf("erro ao executar seed de teste: %v", err)
		}
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanupCancel()
		_ = database.Drop(cleanupCtx)
	})

	return buildRouter()
}

func login(t *testing.T, router *gin.Engine, email, password string) string {
	t.Helper()

	recorder := doJSON(t, router, http.MethodPost, "/auth/login", "", map[string]any{
		"email":    email,
		"password": password,
	})
	if recorder.Code != http.StatusOK {
		t.Fatalf("esperava 200 no login, recebeu %d: %s", recorder.Code, recorder.Body.String())
	}

	var response authTestResponse
	decodeJSON(t, recorder, &response)
	if response.AccessToken == "" {
		t.Fatal("login nao retornou accessToken")
	}

	return response.AccessToken
}

func doJSON(t *testing.T, router *gin.Engine, method, path, token string, body any) *httptest.ResponseRecorder {
	t.Helper()

	var payload bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&payload).Encode(body); err != nil {
			t.Fatalf("erro ao serializar body: %v", err)
		}
	}

	request := httptest.NewRequest(method, path, &payload)
	request.Header.Set("Content-Type", "application/json")
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	return recorder
}

func decodeJSON(t *testing.T, recorder *httptest.ResponseRecorder, target any) {
	t.Helper()
	if err := json.NewDecoder(recorder.Body).Decode(target); err != nil {
		t.Fatalf("erro ao decodificar resposta JSON: %v. body: %s", err, recorder.Body.String())
	}
}
