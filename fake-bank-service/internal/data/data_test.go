package data

import (
	"errors"
	"testing"
)

func TestGenFakeCards(t *testing.T) {
	firstTwelve := "11112222333"
	card, err := GenFakeCards(firstTwelve)
	if err != nil {
		if errors.Is(err, CardNumberLenError) {
			t.Log("expected error need length of 12")
		}
		return
	}

	if len(card) == 16 {
		t.Log("expect len of 16 all pass")
	}
}

func TestGenFakeCv(t *testing.T) {
	expectedLen := 3
	cv := GenFakeCv()
	if len(cv) != 3 {
		t.Errorf("expected length of %d", expectedLen)
	}

}
