package model

import "github.com/dijckstra/cartola-data-scrapper/util"

// Round describes basic information in a game.
type Round struct {
	RoundID   int             `json:"rodada_id" bson:"_id"`
	StartTime util.CustomTime `json:"inicio" bson:"inicio"`
	EndTime   util.CustomTime `json:"fim" bson:"fim"`
}
