package security

import "regexp"

// store the last four number of a card
func LastFour(src *string, strRegX string, replaceChar string) {

	re := regexp.MustCompile(strRegX)
	replacedBytes := re.ReplaceAll([]byte(*src), []byte(replaceChar))

	*src = string(replacedBytes)
}
