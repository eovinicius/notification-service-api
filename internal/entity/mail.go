package entity

import "time"

type Mail struct {
	From    string
	To      string
	Subject string
	Date    time.Time
	Body    string
}

func (m Mail) ToByte() []byte {
	return []byte("From: " + m.From + "\r\n" +
		"To: " + m.To + "\r\n" +
		"Subject: " + m.Subject + "\r\n" +
		"Date: " + m.Date.String() + "\r\n" +
		"\r\n" +
		m.Body)
}
