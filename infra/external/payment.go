package external

import (
	"dependency-injection-sample/domain/repository"
	"fmt"
)

type RakutenPay struct {
}

func NewRakutenPay() repository.Payment {
	return RakutenPay{}
}

func (r RakutenPay) MakePayment(amountToPay int) string {
	return fmt.Sprintf("I paid %d yen by RakutenPay", amountToPay)
}
