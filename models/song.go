package models

type Song struct {
	Name     string  `json:"name" bson:"name"`
	Artist   string  `json:"artist" bson:"artist"`
	Duration float32 `json:"duration" bson:"duration"`
}
