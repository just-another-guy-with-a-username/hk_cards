package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/layout"
    "hk_cards/types"
    "hk_cards/cards"
    "fmt"
//    "time"
//    "math"
//    "slices"
)

func initDeck(i int, p int, h *types.Handler) *types.Group {
    deck := new(types.Group)
    deck.Length = 0
    deck.Type = "deck"
    if p == 1 {
        deck.Location = "Player1"
    }
    if p == 2 {
        deck.Location = "Player2"
    }
    for n := 0; n < 53; n++ {
        c := new(types.Card)
        c.Name = fmt.Sprintf("%d", n+1)
        deck.NewCard(*c)
    }
    deck.NewCard(cards.NewNailSlash(h))
    c := new(types.Card)
    c.Name = "shaman stone"
    c.Type = "charm"
    c.HandlerObj = h
    c.GroupType = "deck"
    c.GroupLocation = deck.Location
    c.Effect = cards.ShamanStone
    c.NotchCost = 3
    deck.NewCard(*c)
    c = new(types.Card)
    c.Name = "vengeful spirit"
    c.Type = "spell"
    c.HandlerObj = h
    c.FailChance = 0
    c.GroupType = "deck"
    c.GroupLocation = deck.Location
    c.Damage = 10
    c.Soul = 2
    deck.NewCard(*c)
    c = new(types.Card)
    c.Name = "vengeful spirit"
    c.Type = "spell"
    c.HandlerObj = h
    c.FailChance = 0
    c.GroupType = "deck"
    c.GroupLocation = deck.Location
    c.Damage = 10
    c.Soul = 2
    deck.NewCard(*c)
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
    game.Player1.Soul = 4
    game.Player1.Notches = 3
    game.Player1.ArtUses = 0
    game.Player1.WeakPointS = false
    game.Player1.WeakPointP = false
    game.Player1.Overcharmed = false
    game.Player2 = new(types.Player)
    game.Player2.Health = 100
    game.Player2.Soul = 4
    game.Player2.Notches = 3
    game.Player2.ArtUses = 0
    game.Player2.WeakPointS = false
    game.Player2.WeakPointP = false
    game.Player2.Overcharmed = false
    game.Player1.Hand = new(types.Group)
    game.Player1.Hand.Cards = []types.Card{}
    game.Player1.Hand.Length = 0
    game.Player1.Hand.Type = "hand"
    game.Player1.Hand.Location = "Player1"
    game.Player2.Hand = new(types.Group)
    game.Player2.Hand.Cards = []types.Card{}
    game.Player2.Hand.Length = 0
    game.Player2.Hand.Type = "hand"
    game.Player2.Hand.Location = "Player2"
    game.Player1.Discard = new(types.Group)
    game.Player1.Discard.Cards = []types.Card{}
    game.Player1.Discard.Length = 0
    game.Player1.Discard.Type = "discard"
    game.Player1.Discard.Location = "Player1"
    game.Player2.Discard = new(types.Group)
    game.Player2.Discard.Cards = []types.Card{}
    game.Player2.Discard.Length = 0
    game.Player2.Discard.Type = "discard"
    game.Player2.Discard.Location = "Player2"
    game.Player1.Charms = new(types.Group)
    game.Player1.Charms.Cards = []types.Card{}
    game.Player1.Charms.Length = 0
    game.Player1.Charms.Type = "charms"
    game.Player1.Charms.Location = "Player1"
    game.Player2.Charms = new(types.Group)
    game.Player2.Charms.Cards = []types.Card{}
    game.Player2.Charms.Length = 0
    game.Player2.Charms.Type = "charms"
    game.Player2.Charms.Location = "Player2"
    game.Player1.Deck = initDeck(i1, 1, game)
    game.Player2.Deck = initDeck(i2, 2, game)
    return *game
}

func displayCards(w fyne.Window, p int, h *types.Handler) {
    if p == 1 {
        cardsPresent := container.New(layout.NewHBoxLayout())
        centeredCards := container.NewCenter(cardsPresent)
        for i, card := range(h.Player1.Hand.Cards) {
            cardsPresent.Add(widget.NewButton(card.Name, func() {PlayCardHandling(p, i, h, centeredCards, w)}))
        }
        w.SetContent(centeredCards)
    } else if p == 2 {
        cardsPresent := container.New(layout.NewHBoxLayout())
        centeredCards := container.NewCenter(cardsPresent)
        for i, card := range(h.Player2.Hand.Cards) {
            cardsPresent.Add(widget.NewButton(card.Name, func() {PlayCardHandling(p, i, h, centeredCards, w)}))
        }
        w.SetContent(centeredCards)
    }
}

func PlayCardHandling(p int, i int, h *types.Handler, c *fyne.Container, w fyne.Window) {
    if p == 1 {
        h.Player1.Play(i, w)
        p = 2
    } else if p == 2 {
        h.Player2.Play(i, w)
        p = 1
    }
    c.RemoveAll()
    status := Turn(w, p, h)
    if status == 0 {
    }
}

func Turn(w fyne.Window, p int, h *types.Handler) int {
    if p == 1 {
        h.Player1.DrawHand()
    }
    if p == 2 {
        h.Player2.DrawHand()
    }
    displayCards(w, p, h)
    if (h.Player2.Deck.Length == 0 && h.Player2.Hand.Length == 0) || h.Player2.Health <= 0 {
        return 1
    }
    if (h.Player1.Deck.Length == 0 && h.Player1.Hand.Length == 0) || h.Player1.Health <= 0 {
        return 2
    }
    return 0
}

func main() {
    game := initGame(1, 1)
    a := app.New()
    w := a.NewWindow("Colosseum Of Fools")
    startButton := container.NewCenter(container.New(layout.NewHBoxLayout(), widget.NewButton("Start Game", func() {Turn(w, 1, &game)})))
    w.SetContent(startButton)
    w.ShowAndRun()
}
