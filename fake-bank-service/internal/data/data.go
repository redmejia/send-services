package data

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Card Fake cards
type Card struct {
	FullName   string `json:"card_issuer"`
	CardNumber string `json:"card_number"`
	CvNumber   string `json:"cv_number"`
	StatusCode int    `json:"status_code"`
	Amount     int    `json:"amount"`
	Proceed    bool   `json:"proceed"`
}

const (
	// last four number on the card
	minNum = 1000
	maxNum = 3000
	// cv number card
	minCv = 100
	maxCv = 200
)

var CardNumberLenError = errors.New("card number length must be 12")

// GenFakeCards
func GenFakeCards(firstTwelveNum string) (string, error) {
	if len(firstTwelveNum) == 0 || len(firstTwelveNum) < 12 || len(firstTwelveNum) > 12 {
		return "", CardNumberLenError
	}
	rand.Seed(time.Now().UnixNano())

	genNum := rand.Intn(maxNum-minNum) + minNum
	lastFour := strconv.Itoa(genNum)

	card := fmt.Sprintf("%s%s", firstTwelveNum, lastFour)
	return card, nil
}

// GenFakeCv
func GenFakeCv() (cv string) {
	rand.Seed(time.Now().UnixNano())

	genCvNum := rand.Intn(maxCv-minCv) + minCv
	cvNum := strconv.Itoa(genCvNum)

	cv = fmt.Sprintf("%s", cvNum)

	return

}
