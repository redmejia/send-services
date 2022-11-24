package database

import (
	"bank/fake-cards/internal/data"
	"errors"
	"log"
)

type NoneDB struct {
	Db                []data.Card
	InfoLog, ErrorLog *log.Logger
}

func NewNoneDb() *NoneDB {
	return new(NoneDB)
}

func (n *NoneDB) GenerateFakeCards(twelveNum string, amountInCent int, statusCode int, proceed bool) (fakeCardPool []data.Card) {

	for i := 0; i < 2; i++ {
		fakeCard, _ := data.GenFakeCards(twelveNum)
		fakeCvNum := data.GenFakeCv()

		card := data.Card{
			FullName:   "Elon Munsk",
			CardNumber: fakeCard,
			CvNumber:   fakeCvNum,
			StatusCode: statusCode,
			Amount:     amountInCent,
			Proceed:    proceed,
		}

		n.Db = append(n.Db, card)
	}

	return n.Db
}

var ErrorNonedDBRowInResultSet = errors.New("nonodb: no rows in result set")

func (n *NoneDB) GetInfo(cardNum string, cardCv string) (data.Card, error) {

	for _, card := range n.Db {
		if card.CardNumber == cardNum && card.CvNumber == cardCv {
			return card, nil
		}
	}
	return data.Card{}, ErrorNonedDBRowInResultSet
}
