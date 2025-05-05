package types

import (
    "errors"
    "math/rand"
)

type Handler struct {
    Player1   Player
    Player2   Player
    Turns     int
    HandSize  int
    VoidArea  Group
    Infection Group
}

type Player struct {
    Hand      Group
    Deck      Group
    Discard   Group
    Charms    Group
    Health    int
    Soul      int
    Notches   int
    WeakPoint bool
}

type Group struct {
    Cards    []Card
    Length   int
    Type     string
    Location string
}

type Card struct {
    Name          string
    Type          string
    GroupType     string
    GroupLocation string
    Effect        func(*Handler) error
    TurnsLeft     int
    HandlerObj    Handler
}

func (g *Group) shuffle() error {
    if g.Type != "deck" {
        return errors.New("wrong group type")
    }
    newOrder := []Card{}
    l := g.Length
    for i := 0; i < l; i++ {
        deckIndex := rand.Intn(g.Length)
        newOrder = append(newOrder, g.Cards[deckIndex])
        g.Cards[deckIndex] = g.Cards[g.Length-1]
        g.Cards = g.Cards[:g.Length-1]
        g.Length--
    }
    g.Length = l
    g.Cards = newOrder
    return nil
}

func (g *Group) draw() error, Card {
    if g.Length == 0 {
        return errors.New("group is empty"), nil
    }
    drawnCard := g.Cards[g.Length-1]
    g.Cards := g.Cards[:g.Length-1]
    g.Length--
    return nil, drawnCard
}

func (g *Group) newCard(c Card) {
    g.Cards = append(g.Cards, c)
    g.Length++
}

func (c *Card) play() error {
    err := c.effect(c.HandlerObj)
    if err != nil {
        return err
    }
    return nil
}
