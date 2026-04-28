package bank

import (
	"errors"
	"fmt"
)

type Transaction interface {
	Process() error
	GetAmount() int
}

type BankTransfer struct {
	Jumlah int
}

func (b BankTransfer) GetAmount() int {
	return b.Jumlah
}

func (b BankTransfer) Process() error {
	if b.Jumlah <= 0 {
		return errors.New("error: amount harus > 0")
	}
	fmt.Printf("Success: Membayar %d menggunakan Bank Transfer\n", b.Jumlah)
	return nil
}
