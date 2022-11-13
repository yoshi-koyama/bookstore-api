package response

import "net/http"

type BuyBooks struct {
	Message string `json:"message"`
}

func NewBuyBooks(message string) *BuyBooks {
	return &BuyBooks{
		Message: message,
	}
}

func (b *BuyBooks) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
