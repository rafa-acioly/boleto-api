package models

import (
	"errors"
	"time"
)

// Title título de cobrança de entrada
type Title struct {
	CreateDate     time.Time
	ExpireDateTime time.Time
	ExpireDate     string
	AmountInCents  uint64
	OurNumber      uint
}

// NewTitle instancia um novo título
func NewTitle(expDate string, amountInCents uint64, ourNumber uint) (*Title, error) {
	eDate, err := parseDate(expDate)
	if err != nil {
		return nil, err
	}

	cDate, err := parseDate(time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	if amountInCents < 1 {
		return nil, errors.New("Valor não pode ser menor do que 1 centavo")
	}

	title := Title{}
	title.AmountInCents = amountInCents
	title.ExpireDateTime = eDate
	title.OurNumber = ourNumber
	title.CreateDate = cDate
	if title.CreateDate.After(title.ExpireDateTime) {
		return nil, errors.New("Data de expiração não pode ser menor que a data de hoje")
	}

	return &title, nil
}

// GetCreateDate Retorna a data de crição do título
func (t *Title) GetCreateDate() time.Time {
	return t.CreateDate
}

func parseDate(t string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Now(), err
	}
	return date, nil
}
