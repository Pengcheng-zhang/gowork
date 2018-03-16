package model

import (
	"time"
)

type EmailVerifyModel struct {
	Email string
	Code string
	MailType string
	Duration int
	CreatedAt time.Time
}

func (EmailVerifyModel) TableName() string {
	return "yz_email_verify"
}