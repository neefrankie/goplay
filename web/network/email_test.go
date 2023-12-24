package network

import (
	"goplay/web/config"
	"testing"
	"time"

	"github.com/go-mail/mail"
	sm "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

func TestSimpleMail(t *testing.T) {
	conn := config.MustLoad().GetEmailConn()

	t.Logf("%s\n", conn)

	server := sm.NewSMTPClient()
	server.Host = conn.Host
	server.Port = conn.Port
	server.Username = conn.User
	server.Password = conn.Pass
	// server.Encryption = sm.EncryptionSTARTTLS

	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	// server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	smtpClient, err := server.Connect()
	if err != nil {
		t.Fatal(err)
	}

	email := sm.NewMSG()
	email.SetFrom("From Example <neefrankie@163.com>").
		AddTo("neefrankie@outlook.com").
		SetSubject("New Go Email").
		SetBody(sm.TextHTML, htmlBody)

	if email.Error != nil {
		t.Fatal(err)
	}

	err = email.Send(smtpClient)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Email Sent")
	}
}

func TestGoMail(t *testing.T) {
	conn := config.MustLoad().GetEmailConn()

	t.Logf("%s\n", conn)

	d := mail.NewDialer(conn.Host, conn.Port, conn.User, conn.Pass)

	m := mail.NewMessage()
	m.SetAddressHeader("From", "neefrankie@163.com", "neefrankie")
	m.SetAddressHeader("To", "neefrankie@outlook.com", "Victor")
	m.SetHeader("Subject", "New Go Email")
	m.SetBody("text/html", htmlBody)

	err := d.DialAndSend(m)
	if err != nil {
		t.Fatal(err)
	}
}
