package repository

type Payment interface {
	MakePayment(amountToPay int) string
}
