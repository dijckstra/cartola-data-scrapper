package model

// Player contains various information and performance metrics on a player.
type Player struct {
	ID          int     `json:"atleta_id" bson:"_id"`
	Name        string  `json:"nome" bson:"name"`
	Nickname    string  `json:"apelido" bson:"nickname"`
	TeamID      int     `json:"clube_id" bson:"team_id"`
	PositionID  int     `json:"posicao_id" bson:"position_id"`
	StatusID    int     `json:"status_id" bson:"status_id"`
	Points      float32 `json:"pontos_num" bson:"points"`
	Price       float32 `json:"preco_num" bson:"price"`
	Variation   float32 `json:"variacao_num" bson:"variation"`
	Average     float32 `json:"media_num" bson:"average"`
	GamesPlayed int     `json:"jogos_num" bson:"games_played"`
	Scout       `json:"scout" bson:"scout"`
}
