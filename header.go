package mail

import (
	"bytes"
	"time"
)

type MailHeader struct {
	To 		 *Contacts
	From 	 *Contacts
	Cc		 *Contacts
	Bcc      *Contacts
	ReplyTo  *Contacts
	Subject  string
	Custom   map[string]string
}

func (c *MailHeader) RemoveCustom(name string) {
	if _, ok := c.Custom[name]; ok {
		delete(c.Custom, name)
	}
}

func (c *MailHeader) SetCustom(name, value string) {
	c.Custom[name] = value
}

func (c *MailHeader) Write(buffer *bytes.Buffer) {
	if _, ok := c.Custom["Date"]; !ok {
		buffer.WriteString("Date: " + time.Now().Format("Tue Sep 16 21:58:58 +0000 2014") + CRLF)
	}
	if c.To.Len() > 0 {
		buffer.WriteString("To: " + c.To.String() + CRLF)
	}
	if c.From.Len() > 0 {
		buffer.WriteString("From: " + c.From.String() + CRLF)
	}
	if c.Cc.Len() > 0 {
		buffer.WriteString("Cc: " + c.Cc.String() + CRLF)
	}
	if c.Bcc.Len() > 0 {
		buffer.WriteString("Bcc: " + c.Bcc.String() + CRLF)
	}
	if c.ReplyTo.Len() > 0 {
		buffer.WriteString("Reply-To: " + c.ReplyTo.String() + CRLF)
	}
	if c.Subject != "" {
		buffer.WriteString("Subject: " + c.Subject + CRLF)
	}
	if len(c.Custom) > 0 {
		for name, value := range c.Custom {
			buffer.WriteString(name + ": " + value + CRLF)
		}
	}
}

func (c *MailHeader) String() string {
	buffer := new(bytes.Buffer)
	c.Write(buffer)
	return string(buffer.Bytes())
}

func NewMailHeader() *MailHeader {
	return &MailHeader{
		To: 		NewContacts(),
		From: 		NewContacts(),
		Cc: 		NewContacts(),
		Bcc: 		NewContacts(),
		ReplyTo: 	NewContacts(),
		Custom:		map[string]string{"MIME-Version": "1.0"},
	}
}