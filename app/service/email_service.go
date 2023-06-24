package service

import "github.com/rizalarfiyan/be-petang/app/model"

type EmailService interface {
	SendEmail(payload model.MailPayload) error
}
