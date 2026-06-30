package main

import (
	"context"
	"fmt"
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/database/mongodb"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/seed"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	mongodb.InitConnection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := seed.Run(ctx, mongodb.GetDatabase()); err != nil {
		panic(err)
	}

	fmt.Println("Seed executado com sucesso")
	fmt.Println("Admin:", seed.AdminEmail, "/", seed.AdminPassword)
	fmt.Println("Cliente:", seed.ClientEmail, "/", seed.ClientPassword)
	fmt.Println("Promocao:", seed.PromotionCode)
}
