package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) error
	Name() string
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) error {
	fmt.Printf("Amount %f - paid using Credit Card", amount)
	return nil
}
func (c *CreditCard) Name() string { return "CreditCard" }

//-----------------------------

type BankTransfer struct {
	accountNo string
}

func (b *BankTransfer) Pay(amount float64) error {
	fmt.Printf("Amount %f - paid using Bank Transfer via account: %s", amount, b.accountNo)
	return nil
}
func (b *BankTransfer) Name() string { return "BankTransfer" }

//------------------------------

type PayPal struct{}

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Amount %f - paid using PayPal", amount)
	return nil
}

func (p *PayPal) Name() string { return "PayPal" }

//------------------------------

type PaymentProcessor struct{}

func (pp *PaymentProcessor) Process(method PaymentMethod, amount float64) error {
	if err := method.Pay(amount); err != nil {
		return err
	}
	fmt.Printf("Payment Processed with %s", method.Name())
	return nil
}
