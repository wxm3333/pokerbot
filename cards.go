package cards


type Suit struct {
    Name string
}

const diamond := Suit("diamond")
const club := Suit("club")
const heart := Suit("heart")
const spade := Suit("spade")


type Card struct {
    """
    Value is an int between 1 and 13 inclusive.
    S is one of the 4 suits.
    """
    Value int
    S Suit
}

type Deck struct {
    """
    A deck is a list of cards.
    There are at most 52 cards in a deck.
    Cards get dealt from the lowest index.
    """
    Cards *[]Card
}

func (d *Deck) Shuffle() {
    """
    Randomize the order of cards in the deck
    """
}

func (d *Deck) Deal(numCards int) []Card {
    """
    Remove and return the top numCards cards from d.Cards
    """
    out := d.Cards[:numCards]
    d.Cards = d.Cards[numCards:]
    return out
}
