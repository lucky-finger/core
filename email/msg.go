package email

import "gopkg.in/gomail.v2"

type MsgOption func(m *gomail.Message)

// MsgWithHtmlContent 传入html格式的内容
func MsgWithHtmlContent(content string) MsgOption {
	return func(m *gomail.Message) {
		m.SetBody("text/html", content)
	}
}

// MsgWithTextContent 传入文本格式的内容
func MsgWithTextContent(content string) MsgOption {
	return func(m *gomail.Message) {
		m.SetBody("text/plain", content)
	}
}

// MsgWithSubject 传入主题
func MsgWithSubject(subject string) MsgOption {
	return func(m *gomail.Message) {
		m.SetHeader("Subject", subject)
	}
}
