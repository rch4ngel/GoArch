//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

// Creating a suit type for our cards structure.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

// Making the ranks 1 bases will allow a simpler way of thinking when building out
// Scoring systems. Setting the first iteration to a throwaway value will achieve one
// based functionality.
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)
type Card struct {
	Suit
	Rank
}

func (c Card) String() string{
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}