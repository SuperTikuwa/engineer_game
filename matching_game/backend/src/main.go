package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SuperTikuwa/matching_game/handler"
	"github.com/SuperTikuwa/matching_game/sheetclient"
	"github.com/joho/godotenv"
)

func init() {

}

func loadEnv() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	fmt.Println(os.Getenv("TEST"))
	ctx := context.Background()

	client, err := sheetclient.NewSheetClient(ctx, os.Getenv("SPREAD_SHEET_ID"))
	if err != nil {
		panic(err)
	}

	hash, err := client.GenerateHash(sheetclient.STUDENT)
	if err != nil {
		log.Fatal("Cannot generating hash")
	}

	http.HandleFunc("/", handler.WordsGetHandler)

	log.Println(hash)
}
