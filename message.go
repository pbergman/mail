package mail

import (
	"crypto/rand"
	"encoding/hex"
	"bytes"
	"net/smtp"
)

const (
	CRLF string = "\r\n"
)

type MailMessage struct {
	Header 	*MailHeader
	Html	[]byte
	Text 	[]byte
	Charset string // charset ISO-8859-1, UTF-8 etc.
}

func (m *MailMessage) GetMessage() []byte {
	buffer := new(bytes.Buffer)
	switch {
	case len(m.Html) > 0 && len(m.Text) > 0: // multi message
		boundaries := []string{GetRandom(), GetRandom()}
		// Message Header
		m.Header.SetCustom("Content-type", "multipart/related;boundary=" + boundaries[0])
		m.Header.Write(buffer)
		// Message Body
		buffer.WriteString(CRLF + "--" + boundaries[0] + CRLF)
		buffer.WriteString("Content-Type: multipart/alternative;" + CRLF)
		buffer.WriteString("\tboundary=\"" + boundaries[1] + "\"")
		// Body Mode text
		buffer.WriteString(CRLF + "--" + boundaries[1] + CRLF)
		buffer.WriteString("Content-Type: text/plain; charset=" + m.Charset + CRLF)
		buffer.WriteString(CRLF)
		buffer.Write(m.Text)
		buffer.WriteString(CRLF)
		//Body Mode Html
		buffer.WriteString(CRLF + "--" + boundaries[1] + CRLF)
		buffer.WriteString("Content-Type: text/html; charset=" + m.Charset + CRLF)
		buffer.WriteString("Content-Transfer-Encoding: quoted-printable" + CRLF)
		buffer.WriteString(CRLF)
		buffer.Write(m.Html)
		buffer.WriteString(CRLF)
		// End Message Body
		buffer.WriteString(CRLF + "--" + boundaries[1] + "--" + CRLF)
		buffer.WriteString(CRLF + "--" + boundaries[0] + "--" + CRLF)
	case len(m.Html) > 0 && len(m.Text) <= 0: // html  message
		m.Header.SetCustom("Content-type", "text/html; charset=" + m.Charset)
		m.Header.Write(buffer)
		buffer.WriteString(CRLF)
		buffer.Write(m.Html)
	case len(m.Html) <= 0 && len(m.Text) > 0: // text  message
		m.Header.SetCustom("Content-type", "text/plain; charset=" + m.Charset)
		m.Header.Write(buffer)
		buffer.WriteString(CRLF)
		buffer.Write(m.Text)
	}

	return buffer.Bytes()
}

func (m *MailMessage) Send(addr string, auth smtp.Auth) error {
	return smtp.SendMail(addr, auth, m.Header.From.String(), m.Header.To.String(), m.GetMessage())
}

func GetRandom() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (m *MailMessage) AddFrom(c *Contact) {
	m.Header.From.Add(c)
}

func (m *MailMessage) AddTo(c *Contact) {
	m.Header.To.Add(c)
}

func (m *MailMessage) SetSubject(s string) {
	m.Header.Subject = s
}

func NewMailMessage() *MailMessage {
	return &MailMessage{
		Header:  NewMailHeader(),
		Html: 	 make([]byte, 0),
		Text: 	 make([]byte, 0),
		Charset: "UTF-8",
	}
}
