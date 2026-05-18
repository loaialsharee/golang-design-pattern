package main

import "database/sql"

type Order struct{}

// Each of the following has ONE reason to change
type OrderValidator struct{}

func (v *OrderValidator) Validate(order Order) error { return nil }

//--------------------

type OrderRepository struct{ db *sql.DB }

func (r *OrderRepository) Save(order Order) error { return nil }

//--------------------

type EmailNotifier struct{ smtpAddr string }

func (e *EmailNotifier) Notify(order Order) error { return nil }

//--------------------

type AuditLogger struct{ filePath string }

func (a *AuditLogger) Log(order Order) error { return nil }

//--------------------

// OrderService now just orchestrates
type OrderService struct {
	validator *OrderValidator
	repo      *OrderRepository
	notifier  *EmailNotifier
	logger    *AuditLogger
}

func (o *OrderService) CreateOrder(order Order) error {
	if err := o.validator.Validate(order); err != nil {
		return err
	}
	if err := o.repo.Save(order); err != nil {
		return err
	}
	if err := o.notifier.Notify(order); err != nil {
		return err
	}
	return o.logger.Log(order)
}
