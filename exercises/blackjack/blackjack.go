package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "joker":
		return 0
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	score := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)
	if score == 21 && dealerScore < 10 {
		return "W"
	}

	if score >= 17 && score <= 20 {
		return "S"
	}

	if score >= 12 && score <= 16 && dealerScore < 7 {
		return "S"
	}

	if score >= 12 && score <= 16 && dealerScore >= 7 {
		return "H"
	}

	if score <= 11 {
		return "H"
	}

	return "S"
}
