package blackjack

var blackjackMapping = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return blackjackMapping[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	myCalculatedCards := blackjackMapping[card1] + blackjackMapping[card2]

	dealerCardValue := blackjackMapping[dealerCard]
	switch {

	case myCalculatedCards == 22:
		{
			return "P"
		}
	case myCalculatedCards == 21:
		{
			if dealerCardValue < 10 {
				return "W"
			} else {
				return "S"
			}
		}
	case (myCalculatedCards >= 17) && (myCalculatedCards <= 20):
		{
			return "S"
		}

	case (myCalculatedCards <= 16) && (myCalculatedCards >= 12):
		{
			if dealerCardValue >= 7 {
				return "H"
			}
			return "S"
		}
	default:
		return "H"
	}

}
