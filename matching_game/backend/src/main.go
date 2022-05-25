package main

import (
	"context"
	"os"

	"github.com/SuperTikuwa/matching_game/sheetclient"
)

func main() {
	ctx := context.Background()
	client, err := sheetclient.NewSheetClient(ctx, os.Getenv("SPREAD_SHEET_ID"))
	if err != nil {
		panic(err)
	}
	client.PrintValue()
}
