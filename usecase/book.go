package usecase

import (
	"context"
	"dependency-injection-sample/domain/repository"
	"errors"
)

type Book interface {
	BuyBooks(ctx context.Context, id int, amountToBuy int) (*string, error)
}

type bookUseCase struct {
	bookRepo    repository.Book
	paymentRepo repository.Payment
}

func NewBook(bookRepo repository.Book, paymentRepo repository.Payment) Book {
	return bookUseCase{
		bookRepo:    bookRepo,
		paymentRepo: paymentRepo,
	}
}

func (b bookUseCase) BuyBooks(ctx context.Context, id int, amountToBuy int) (*string, error) {
	book := b.bookRepo.FindByID(id)
	if book == nil {
		return nil, errors.New("cannot find book")
	}

	result := b.paymentRepo.MakePayment(book.Price() * amountToBuy)
	return &result, nil
}
