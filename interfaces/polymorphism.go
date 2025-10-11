// We can write functions that operate on different types as long as they share behavior.
package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Notify(msg string) { fmt.Println("Email:", msg) }

type SMSNotifier struct{}

func (SMSNotifier) Notify(msg string) { fmt.Println("SMS:", msg) }

func sendAlert(n Notifier) {
	n.Notify("System overload!")
}

func main() {
	sendAlert(EmailNotifier{})
	sendAlert(SMSNotifier{})
}
