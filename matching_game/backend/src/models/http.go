package models

type WordGetResponse struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
}

type GameStatus struct {
	Turn          int  `json:"turn"`
	Card          int  `json:"card"`
	Count         int  `json:"count"`
	PlayerChanged bool `json:"playerChanged"`
}
