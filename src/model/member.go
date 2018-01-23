package model

import "time"

type MemberModel struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	StartedAt time.Time `json:"started_at"`
	EndAt time.Time `json:"end_at"`
	CreatedAt time.Time `json:"created_at"`
} 
func (MemberModel) TableName() string {
	return "yz_member"
}