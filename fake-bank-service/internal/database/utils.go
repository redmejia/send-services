package database

import (
	"bank/fake-cards/internal/data"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func GenFakesTestCardsToInsert(twelveNum string, amountInCent int, statusCode int, proceed bool) []data.Card {
	var fakeCardsList []data.Card
	for i := 0; i < 2; i++ {
		fakeCard, _ := data.GenFakeCards(twelveNum)
		fakeCvNum := data.GenFakeCv()

		card := data.Card{
			FullName:   "Elon Musk",
			CardNumber: fakeCard,
			CvNumber:   fakeCvNum,
			StatusCode: statusCode,
			Amount:     amountInCent,
			Proceed:    proceed,
		}

		fakeCardsList = append(fakeCardsList, card)
	}

	return fakeCardsList

}

func CardForTest(cards []data.Card) string {

	inserCards := `insert into cards (card_issuer, card_number, card_cv_number, status_code, amount, proceed) values `

	for _, card := range cards {

		inserCards += fmt.Sprintln()
		inserCards += fmt.Sprintf(`(%s, %s, %s, %d, %d, %t),`,
			card.FullName, card.CardNumber, card.CvNumber, card.StatusCode, card.Amount, card.Proceed,
		)
		inserCards += ``
	}
	inserCards = inserCards[:len(inserCards)-1]

	inserCards += `;`
	return inserCards

}

// GenerateFakeCards generate testing cards
func (p *Postgresql) GenerateFakeCards(twelveNum string, amountInCent int, statusCode int, proceed bool) (fakeCardPool []data.Card) {
	// cent 100000 = $1000
	// success 2 status code
	// decline 4 status code

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
 		insert into cards (card_issuer, card_number, card_cv_number, status_code, amount, proceed)
 		values ($1, $2, $3, $4, $5, $6)
 	`
	for i := 0; i < 2; i++ {

		fakeCard, _ := data.GenFakeCards(twelveNum)

		fakeCvNum := data.GenFakeCv()

		cards := data.Card{
			FullName:   "Elon Musk",
			CardNumber: fakeCard,
			CvNumber:   fakeCvNum,
			StatusCode: statusCode,
			Amount:     amountInCent,
			Proceed:    proceed,
		}

		_, err := p.Db.ExecContext(
			ctx,
			query,
			cards.FullName,
			cards.CardNumber,
			cards.CvNumber,
			statusCode,
			cards.Amount,
			proceed,
		)

		if err != nil {
			log.Println("err ", err)
			return
		}

		fakeCardPool = append(fakeCardPool, cards)

	}

	return
}

var NoRowInResultSet = errors.New("sql: no rows in result set")

// GeyInfo get ths card information
func (p *Postgresql) GetInfo(cardNum, cardCv string) (data.Card, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select 
			card_issuer, 
			card_number, 
			card_cv_number, 
			status_code, 
			amount, 
			proceed 
		from 
			cards
		where
			card_number = $1 and card_cv_number = $2
			`

	row := p.Db.QueryRowContext(ctx, query, cardNum, cardCv)

	var card data.Card

	err := row.Scan(
		&card.FullName,
		&card.CardNumber,
		&card.CvNumber,
		&card.StatusCode,
		&card.Amount,
		&card.Proceed,
	)
	if err != nil {
		return card, NoRowInResultSet
	}
	return card, nil
}
