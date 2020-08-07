package main 

import (
	"fmt"
	"errors"
	"math/rand"
	"bufio"
	"os"
	"time"
	"os/exec"
)

type Value uint8

const (
	_ Value = iota
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
	_
	Jack
	Queen
	King
)

func (v Value) GetValue() int {
	return int(v)
}

func (v Value) String() string {
	switch (v){
	case 1:
		return "Ace"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 12:
		return "Jack"
	case 13:
		return "Queen"
	default:
		return "King"
	}
}

type Suit uint8

const (
	Diamonds Suit = iota
	Hearts
	Spades
	Clubs
)

func (v Suit) String() string {
	switch (v){
	case 0:
		return "Diamonds"
	case 1:
		return "Hearts"
	case 2:
		return "Spades"
	default:
		return "Clubs"	
	}
}

type Card struct {
	Value
	Suit
}


func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value.String(), c.Suit.String())

}

type DeckOfCards struct {
	cards [52]Card
	size  int
}

func (d *DeckOfCards) Clear() {
	d.size = 0
}

func (d *DeckOfCards) AddCard(c Card) {
	d.cards[d.size] = c
	d.size ++
}

func (d *DeckOfCards) Exists(c Card) bool {
	for i := 0; i < d.size; i++ {
		if d.cards[i] == c {
			return true
		}
	}
	return false
}

func (d *DeckOfCards) GetRandomCard() (Card, error) {
	if d.size == 0 {
		return Card{Value: 1, Suit: 0}, errors.New("Deck is empty")
	}
	return d.cards[rand.Intn(d.size)], nil
}

func (d *DeckOfCards) RemoveCard(c Card) {
	pos := -1
	for i := 0; i < d.size; i++ {
		if d.cards[i] == c {
			pos = i
			break
		}
	}
	if pos == -1 {
		return
	}
	for i := pos + 1; i < d.size; i++ {
		d.cards[i - 1] = d.cards[i]
	}
	d.size --
}

func (d *DeckOfCards) ExtractRandomCard() (Card, error) {
	if d.size == 0 {
		return Card{Value: 1, Suit: 0}, errors.New("Deck is empty")
	}
	rand.Seed(time.Now().UnixNano())
	card := d.cards[rand.Intn(d.size)]
	d.RemoveCard(card)
	return card, nil
}

func (d *DeckOfCards) Init() {
	d.size = 52
	index := 0
	for value := 1; value < 15; value ++ {
		if value == 11 {
			continue
		}
		for suit := 0; suit < 4; suit ++ {
			d.cards[index] = Card {
				Value: Value(value),
				Suit: Suit(suit)}
			index++
		}
	}
}

type Player struct {
	name       string
	score      int
	cards 	   DeckOfCards
	sumOfCards int
	stopTaking bool

}

type BoardCardGame struct {
	startingDeckOfCards DeckOfCards
	playedCards         DeckOfCards
	players             [2]Player
	turn				int
	gameOver			bool
}

func (b *BoardCardGame) canMove() bool {
	return b.players[0].stopTaking == false ||
			b.players[1].stopTaking == false
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (b *BoardCardGame) AskPlayer() bool { 

	fmt.Printf("Player %d's turn.", b.turn)
	fmt.Println("Your cards are:")
	for i := 0; i < b.players[b.turn].cards.size; i++ {
		fmt.Println(b.players[b.turn].cards.cards[i].String())
	}
	fmt.Printf("Your score is: %d\n", b.players[b.turn].score)
	fmt.Println(" Do you want to extract a new card? y/n")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	ClearScreen()
	if text[0] == 'y' {
		return true
	}
	return false
}

func (b *BoardCardGame) addCardToPlayer(player int) {
	card, err := b.startingDeckOfCards.ExtractRandomCard()
	if err != nil {
		fmt.Println(err)
		return
	}
	b.players[player].cards.AddCard(card)
	b.players[player].score = b.players[player].score + card.GetValue()
}

func (b *BoardCardGame) move() {
	ans := b.AskPlayer()
	if ans {
		for ans && !b.gameOver {

			b.addCardToPlayer(b.turn)
			if b.players[b.turn].score == 21 {
				b.gameOver = true
				fmt.Printf("Player %d won. He scored %d points.", b.turn, b.players[b.turn].score)
			} else if b.players[b.turn].score > 21 {
				b.gameOver = true
				fmt.Printf("Player %d won, because player %d scored more than 21 points.",
							1 - b.turn, b.turn)
			}
			if !b.gameOver {
				ans = b.AskPlayer()
			}
		}
	} else {
		b.players[b.turn].stopTaking = true
	}
}

func (b *BoardCardGame) changeTurn() {
	b.turn = 1 - b.turn
}

func (b *BoardCardGame) Init() {
	for i := 0; i < 2; i++ {
		b.addCardToPlayer(0)
		b.addCardToPlayer(1)
	}
	fmt.Printf("Player 0 score: %d\n Player 1 score: %d\n", b.players[0].score, b.players[1].score)
	if b.players[0].score == 21 && b.players[1].score == 21 {
		fmt.Println("Draw")
		b.gameOver = true
	} else if b.players[0].score == 21 && b.players[1].score != 21 {
		fmt.Println("Player 0 won")
		b.gameOver = true
	} else 	if b.players[0].score != 21 && b.players[1].score == 21 {
		fmt.Println("Player 1 won")
		b.gameOver = true
	} else 	if b.players[0].score > 21 && b.players[1].score > 21 {
		fmt.Println("Draw")
		b.gameOver = true
	} else 	if b.players[0].score > 21 && b.players[1].score < 21 {
		fmt.Println("Player 1 won")
		b.gameOver = true
	} else 	if b.players[0].score < 21 && b.players[1].score > 21 {
		fmt.Println("Player 0 won")
		b.gameOver = true
	}
}

func (b *BoardCardGame) Start() {
	b.startingDeckOfCards.Init()
	b.Init()
	for !b.gameOver && b.canMove() {
		b.move()
		b.changeTurn()
	}

	if b.gameOver {
		return
	}

	// We are in !b.canMove() situation
	if b.players[0].score == b.players[1].score {
		fmt.Println("Draw")
	} else if b.players[0].score > b.players[1].score {
		fmt.Println("Player 0 won.")
	} else {
		fmt.Println("Player 1 won.")	
	}
}

func main() {
	b := BoardCardGame {}
	b.Start()
}