package main

func main() {
	cards := newDeckFromFile("mycards")
	cards.print()
	//cards.saveToFile("my_cards")
}
