package main

// var card string = "Ace of Spades"
var decksize int

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"

	// card := newCard()
	// fmt.Println(card)

	// decksize = 52
	// fmt.Println(decksize)

	// printState()

	// cards := []string{"Ace of Spades", newCard()}

	// for i, card := range cards {
	// 	fmt.Println(i, card)

	//cards := deck{"Ace of Spades", newCard()}
	//cards = append(cards, "6 of Spades")
	//fmt.Println(cards)

	//cards := newDeck()
	// //cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	// cards.print()

	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "5 of Diamonds"
}
