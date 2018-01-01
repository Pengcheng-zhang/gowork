package model

import "time"

type Member struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	StartedAt time.Time `json:"started_at"`
	EndAt time.Time `json:"end_at"`
	CreatedAt time.Time `json:"created_at"`
} 
