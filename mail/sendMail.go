package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendMail(messages string) {

	auth := smtp.PlainAuth("", "wenbin171@163.com", "'sytwyx%100s.", "smtp.163.com")

	user := "wenbin171@163.com"
	host := "smtp.163.com:25"
	to := "nemo.emails@gmail.com;"

	subject := "mail subject"
	send_to := strings.Split(to, ";")
	msg := []byte("To: " + to + "\\r\\nFrom: " + user + "\\r\\nSubject: " + subject + "\\r\\n\\r\\n" + messages)
	e := smtp.SendMail(host, auth, user, send_to, msg)
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	sendMail("你好吗？")
}
