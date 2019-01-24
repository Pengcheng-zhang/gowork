package model

import (
	"time"
)

type EmailVerifyModel struct {
	Email string
	Code string
	MailType string
	Duration int
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (EmailVerifyModel) TableName() string {
	return "yz_email_verify"
}