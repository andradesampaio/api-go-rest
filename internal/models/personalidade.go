package models

type Personalidade struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Profissao string `json:"profissao"`
	Historia  string `json:"historia"`
}
