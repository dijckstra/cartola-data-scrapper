package model

// Player contains various information and performance metrics on a player.
type Player struct {
	ID          int     `json:"atleta_id"`
	Name        string  `json:"nome"`
	Nickname    string  `json:"apelido"`
	TeamID      int     `json:"clube_id"`
	PositionID  int     `json:"posicao_id"`
	StatusID    int     `json:"status_id"`
	Points      float32 `json:"pontos_num"`
	Price       float32 `json:"preco_num"`
	Variation   float32 `json:"variacao_num"`
	Average     float32 `json:"media_num"`
	GamesPlayed int     `json:"jogos_num"`
	Scout       `json:"scout"`
}
