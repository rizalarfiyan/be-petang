package adapter

import (
	"crypto/tls"
	"time"

	"github.com/rizalarfiyan/be-petang/config"
	gomail "github.com/xhit/go-simple-mail/v2"
)

func EmailConnection() *gomail.SMTPServer {
	conf := config.Get()
	server := gomail.NewSMTPClient()
	server.Host = conf.Email.Host
	server.Port = conf.Email.Port
	server.Username = conf.Email.User
	server.Password = conf.Email.Password
	server.Encryption = gomail.EncryptionTLS
	server.KeepAlive = true
	server.ConnectTimeout = 30 * time.Second
	server.SendTimeout = 30 * time.Second
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return server
}
