# Raizes do Nordeste API

API feita para o trabalho de Back-end do projeto multidisciplinar. A ideia do sistema e simular uma rede de lanchonetes, com cadastro de usuarios, unidades, produtos, estoque por unidade, pedidos, pagamento mock, fidelidade e promocoes.

O projeto foi feito em Go usando Gin e MongoDB. A API salva os dados no banco, usa JWT para autenticacao e tem controle simples por perfil.

## O que a API faz

- Cadastro e login de usuarios
- Perfis: `CLIENTE`, `ATENDENTE`, `COZINHA`, `GERENTE`, `ADMIN`
- Consentimento LGPD para fidelidade
- Cadastro de unidades
- Cadastro de produtos
- Cardapio por unidade, mostrando estoque e disponibilidade
- Entrada e saida de estoque
- Criacao de pedidos com `canalPedido`
- Filtro de pedidos por canal e status
- Pagamento mock aprovado ou recusado
- Pontos de fidelidade e resgate simples
- Promocoes com cupom de desconto
- Swagger/OpenAPI
- Collection do Postman

## Requisitos

Para rodar sem Docker:

- [Go](https://go.dev/dl/)
- [MongoDB](https://www.mongodb.com/try/download/community)

Para rodar com Docker:

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

## Variaveis de ambiente

Copie o arquivo de exemplo:

```bash
cp .env.example .env
```

No Windows PowerShell:

```powershell
Copy-Item .env.example .env
```

Exemplo usado no projeto:

```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=raizes_do_nordeste
JWT_SECRET=raizes-do-nordeste-dev-secret
JWT_EXPIRES_IN=3600
```

## Rodando localmente

Suba o MongoDB. Um jeito rapido com Docker:

```bash
docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest
```

Se esse container ja existir:

```bash
docker start mongodb
```

Depois rode a API:

```bash
go mod download
go run .
```

A API sobe em:

```text
http://localhost:8080
```

## Rodando tudo com Docker Compose

Esse e o jeito mais simples em outro PC, porque ja sobe API e MongoDB juntos:

```bash
docker compose up --build
```

Para parar:

```bash
docker compose down
```

Se quiser apagar tambem os dados do Mongo:

```bash
docker compose down -v
```

## Swagger

Depois que a API estiver rodando, abra:

```text
http://localhost:8080/swagger
```

O arquivo OpenAPI tambem fica disponivel em:

```text
http://localhost:8080/swagger/openapi.yaml
```

## Postman

A collection esta em:

```text
postman/raizes-do-nordeste-api.postman_collection.json
```

Ela usa a variavel:

```text
base_url = http://localhost:8080
```

Fluxo recomendado para testar:

1. Criar admin
2. Login admin
3. Criar cliente
4. Login cliente
5. Criar unidade
6. Criar produto
7. Dar entrada no estoque
8. Criar promocao
9. Consultar cardapio da unidade
10. Criar pedido
11. Fazer pagamento mock
12. Consultar fidelidade
13. Atualizar status do pedido

## Fluxo principal

O fluxo principal do projeto ficou assim:

```text
Cliente cria pedido
API valida unidade, produto, estoque e canalPedido
API baixa o estoque
Cliente faz pagamento mock
Se aprovado: pedido vira PAGO e gera pontos
Se recusado: pedido vira CANCELADO e estoque volta
Cozinha/atendimento atualizam status ate ENTREGUE
```

## Alguns endpoints

Auth:

- `POST /auth/login`

Usuarios:

- `POST /usuarios`
- `GET /usuarios/me`
- `PATCH /usuarios/me/consentimentos/fidelidade`
- `GET /usuarios/:userId`
- `GET /usuarios?email=...`
- `PATCH /usuarios/:userId`
- `DELETE /usuarios/:userId`

Unidades e cardapio:

- `POST /unidades`
- `GET /unidades?page=1&limit=10`
- `GET /unidades/:unitId`
- `GET /unidades/:unitId/cardapio`
- `PATCH /unidades/:unitId`

Produtos:

- `POST /produtos`
- `GET /produtos?category=comida&page=1&limit=10`
- `GET /produtos/:productId`
- `PATCH /produtos/:productId`

Estoque:

- `POST /estoque/movimentacoes`
- `GET /estoque?unidadeId=...&produtoId=...`

Pedidos:

- `POST /pedidos`
- `GET /pedidos?canalPedido=APP&status=AGUARDANDO_PAGAMENTO`
- `GET /pedidos/:orderId`
- `PATCH /pedidos/:orderId/status`
- `PATCH /pedidos/:orderId/cancelamento`

Pagamentos:

- `POST /pagamentos`
- `GET /pagamentos?pedidoId=...`
- `GET /pagamentos/:paymentId`

Fidelidade:

- `GET /fidelidade/saldo`
- `GET /fidelidade/historico`
- `POST /fidelidade/resgates`

Promocoes:

- `POST /promocoes`
- `GET /promocoes?active=true`
- `GET /promocoes/:promotionId`
- `PATCH /promocoes/:promotionId`

## Testes

O projeto ainda nao tem testes automatizados escritos, mas da para validar compilacao com:

```bash
go test ./...
```

Para teste manual, use o Swagger ou a collection do Postman.

## Observacoes

- Senha e salva com hash.
- Senha nao aparece nas respostas.
- O delete de usuario e logico.
- Estoque e controlado por unidade e produto.
- Pagamento e somente mock, sem integracao real.
- O campo `canalPedido` e obrigatorio nos pedidos.
- Listagens aceitam `page` e `limit`.
