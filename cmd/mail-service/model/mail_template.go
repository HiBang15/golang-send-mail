package model

import (
	"context"
	db "github.com/HiBang15/golang-send-mail.git/db/mail/sqlc"
)

type Mail struct {
	ID      int32  `json:"id"`
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type CreateMailTemplateRequest struct {
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (tran *Transaction) CreateMailTemplate(ctx context.Context, request CreateMailTemplateRequest) (mail db.Mail, err error) {
	err := tran.execTx(ctx, func(queries *interface{}) error {
		var err error
		mail, err := q
	})
}
