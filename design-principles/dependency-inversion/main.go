package main

import "fmt"

// The abstraction — one contract for all channels
type Notifier interface {
	Send(contact, message string) error
}

// Low-level module 1
type SMSService struct{}

func (s *SMSService) Send(contact, message string) error {
	fmt.Printf("SMS to %s: %s\n", contact, message)
	return nil
}

// Low-level module 2
type EmailService struct{}

func (e *EmailService) Send(contact, message string) error {
	fmt.Printf("Email to %s: %s\n", contact, message)
	return nil
}

// Low-level module 3
type PushService struct{}

func (p *PushService) Send(contact, message string) error {
	fmt.Printf("Push to %s: %s\n", contact, message)
	return nil
}

//---------------------------------------

// High-level module — knows NOTHING about SMS, Email, or Push
type AlertSystem struct {
	notifier Notifier // depends on abstraction only
}

// Dependency injected at construction — not hardcoded inside
func NewAlertSystem(n Notifier) *AlertSystem {
	return &AlertSystem{notifier: n}
}

func (a *AlertSystem) SendAlert(contact, message string) error {
	return a.notifier.Send(contact, message)
}

func main() {
	message := "Battery critically low!"

	// Swap the notifier — AlertSystem untouched each time
	smsAlert := NewAlertSystem(&SMSService{})
	smsAlert.SendAlert("+60123456789", message)

	emailAlert := NewAlertSystem(&EmailService{})
	emailAlert.SendAlert("driver@tesla.com", message)

	pushAlert := NewAlertSystem(&PushService{})
	pushAlert.SendAlert("device-token-xyz", message)
}
