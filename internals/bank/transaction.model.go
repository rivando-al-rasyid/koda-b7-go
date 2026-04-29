package bank

import (
	"errors"
	"fmt"
)

// PaymentProcessor adalah interface untuk semua metode pembayaran.
type PaymentProcessor interface {
	Pay(amount float64) error
}

// BankPayment mengimplementasikan PaymentProcessor untuk simulasi Bank.
type BankPayment struct{}

func (b BankPayment) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("error: biaya harus > 0")
	}
	fmt.Printf("Berhasil: Membayar %.2f menggunakan Bank\n", amount)
	return nil
}

// OnlinePayment mengimplementasikan PaymentProcessor untuk simulasi Online.
type OnlinePayment struct{}

func (o OnlinePayment) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("error: biaya harus > 0")
	}
	fmt.Printf("Berhasil: Membayar %.2f menggunakan Online Gateway\n", amount)
	return nil
}

// FictivePayment mengimplementasikan penyimpanan ke slice.
type FictivePayment struct {
	History []float64
}

func (f *FictivePayment) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("error: biaya pembayaran harus > 0")
	}
	f.History = append(f.History, amount)
	return nil
}

// CheckoutSystem hanya peduli pada interface PaymentProcessor.
type CheckoutSystem struct {
	Processor PaymentProcessor
}

func (c *CheckoutSystem) ExecuteCheckout(amounts []float64) {
	for _, amount := range amounts {
		err := c.Processor.Pay(amount)
		if err != nil {
			fmt.Println(err)
		}
	}
}
