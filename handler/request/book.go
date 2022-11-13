package request

import (
	"errors"
	"net/http"
)

type BuyBooks struct {
	ID          *int `json:"id"`
	AmountToPay *int `json:"amount_to_pay"`
}

func (b *BuyBooks) Bind(r *http.Request) error {
	if b.ID == nil {
		return errors.New("missing required ID fields.")
	}

	if b.AmountToPay == nil {
		return errors.New("missing required AmountToPay fields.")
	}

	return nil
}
