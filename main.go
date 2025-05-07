package main

import (
    "hk_cards/types"
    "fmt"
)

func initDeck(i int, p int) *types.Group {
    deck := new(types.Group)
    deck.Length = 0
    deck.Type = "deck"
    if p == 1 {
        deck.Location = "Player1"
    }
    if p == 2 {
        deck.Location = "Player2"
    }
    for n := 0; n < 57; n++ {
        c := new(types.Card)
        c.Name = fmt.Sprintf("%d", n+1)
        deck.NewCard(*c)
    }
    return deck
}

func initGame(i1 int, i2 int) types.Handler {
    game := new(types.Handler)
    game.VoidArea = new(types.Group)
    game.Infection = new(types.Group)
    game.Turns = 0
    game.HandSize = 6
    game.VoidArea.Cards = []types.Card{}
    game.VoidArea.Length = 0
    game.VoidArea.Type = "void"
    game.VoidArea.Location = "game"
    game.Infection.Cards = []types.Card{}
    game.Infection.Length = 0
    game.Infection.Type = "infection"
    game.Infection.Location = "game"
    game.Player1 = new(types.Player)
    game.Player1.Health = 100
    game.Player1.Soul = 0
    game.Player1.Notches = 3
    game.Player1.ArtUses = 0
    game.Player1.WeakPointS = false
    game.Player1.WeakPointP = false
    game.Player1.Overcharmed = false
    game.Player2 = new(types.Player)
    game.Player2.Health = 100
    game.Player2.Soul = 0
    game.Player2.Notches = 3
    game.Player2.ArtUses = 0
    game.Player2.WeakPointS = false
    game.Player2.WeakPointP = false
    game.Player2.Overcharmed = false
    game.Player1.Hand = new(types.Group)
    game.Player1.Hand.Cards = []types.Card{}
    game.Player1.Hand.Length = 0
    game.Player1.Hand.Type = "hand"
    game.Player1.Hand.Location = "player1"
    game.Player2.Hand = new(types.Group)
    game.Player2.Hand.Cards = []types.Card{}
    game.Player2.Hand.Length = 0
    game.Player2.Hand.Type = "hand"
    game.Player2.Hand.Location = "player2"
    game.Player1.Discard = new(types.Group)
    game.Player1.Discard.Cards = []types.Card{}
    game.Player1.Discard.Length = 0
    game.Player1.Discard.Type = "discard"
    game.Player1.Discard.Location = "player1"
    game.Player2.Discard = new(types.Group)
    game.Player2.Discard.Cards = []types.Card{}
    game.Player2.Discard.Length = 0
    game.Player2.Discard.Type = "discard"
    game.Player2.Discard.Location = "player2"
    game.Player1.Charms = new(types.Group)
    game.Player1.Charms.Cards = []types.Card{}
    game.Player1.Charms.Length = 0
    game.Player1.Charms.Type = "charms"
    game.Player1.Charms.Location = "player1"
    game.Player2.Charms = new(types.Group)
    game.Player2.Charms.Cards = []types.Card{}
    game.Player2.Charms.Length = 0
    game.Player2.Charms.Type = "charms"
    game.Player2.Charms.Location = "player2"
    game.Player1.Deck = initDeck(i1, 1)
    game.Player2.Deck = initDeck(i2, 2)
    return *game
}

func main() {
    game := initGame(1, 1)
    for _, c := range(game.Player2.Deck.Cards) {
        fmt.Println(c.Name, c.GroupType)
    }
    fmt.Println("shuffling...")
    game.Player2.Deck.Shuffle()
    for _, c := range(game.Player2.Deck.Cards) {
        fmt.Println(c.Name, c.GroupType)
    }
    fmt.Println("drawing 7 cards")
    game.Player2.Draw()
    game.Player2.Draw()
    game.Player2.Draw()
    game.Player2.Draw()
    game.Player2.Draw()
    game.Player2.Draw()
    game.Player2.Draw()
    for _, c := range(game.Player2.Hand.Cards) {
        fmt.Println(c.Name, c.GroupType)
    }
    fmt.Println("playing a card")
    game.Player2.Play(0)
    for _, c := range(game.Player2.Discard.Cards) {
        fmt.Println(c.Name, c.GroupType)
    }
}
