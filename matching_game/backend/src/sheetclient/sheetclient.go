package sheetclient

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/SuperTikuwa/matching_game/models"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

const (
	STUDENT      = "就活生向け"
	NON_ENGINEER = "非エンジニア職種向け"
	ENGINEER     = "エンジニア目指す子"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}

func IdToString(id int) string {
	switch id {
	case 1:
		return STUDENT
	case 2:
		return NON_ENGINEER
	case 3:
		return ENGINEER
	default:
		return ""
	}
}

func StringToID(s string) int {
	switch s {
	case STUDENT:
		return 1
	case NON_ENGINEER:
		return 2
	case ENGINEER:
		return 3
	default:
		return 0
	}
}

func NewSheetClient(ctx context.Context, spreadsheetID string) (*SheetClient, error) {
	credential := option.WithCredentialsFile("secret.json")

	srv, err := sheets.NewService(ctx, credential)
	if err != nil {
		return nil, err
	}
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

func (s *SheetClient) Get(sheet string) ([]models.Word, error) {
	resp, err := s.srv.Spreadsheets.Values.Get(s.spreadsheetID, fmt.Sprintf("'%s'!A:B", sheet)).Do()
	if err != nil {
		return nil, err
	}

	var words []models.Word
	for i, row := range resp.Values {
		if i == 0 {
			continue
		}

		word := models.Word{}
		word.Word = row[0].(string)
		if len(row) == 2 {
			word.Meaning = row[1].(string)
		}
		word.DifficultyID = uint(StringToID(sheet))

		words = append(words, word)
	}

	return words, nil
}

func (s *SheetClient) GenerateHash(sheet string) (string, error) {
	values, err := s.srv.Spreadsheets.Values.Get(os.Getenv("SPREAD_SHEET_ID"), fmt.Sprintf("'%s'!A:B", sheet)).Do()
	if err != nil {
		return "", err
	}

	var hash string
	for _, row := range values.Values {
		hash += row[0].(string)
	}

	h := sha256.Sum256([]byte(hash))
	hash = hex.EncodeToString(h[:])
	return hash, nil
}
