package types

import (
    "errors"
    "math/rand"
    "slices"
)

type Handler struct {
    Player1   *Player
    Player2   *Player
    Turns     int
    HandSize  int
    VoidArea  *Group
    Infection *Group
}

type Player struct {
    Hand        *Group
    Deck        *Group
    Discard     *Group
    Charms      *Group
    NailBoost   *Group
    Health      int
    Soul        int
    Notches     int
    ArtUses     int
    WeakPointS  bool
    WeakPointP  bool
    Overcharmed bool
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
    Effect        func(*Card, *Handler) error
    NotchCost     int
    TurnsLeft     int
    NailPlus      int
    Damage        int
    EndDamage     int
    Soul          int
    FailChance    float32
    HandlerObj    *Handler
}

func (p *Player) Draw() error {
    err, card := p.Deck.Draw()
    if err != nil {
        return err
    }
    p.Hand.NewCard(card)
    return nil
}

func (p *Player) NailEquip(c Card) {
    p.NailBoost.NewCard(c)
}

func (p *Player) DiscardCard(toDiscard int) error {
    if toDiscard >= p.Hand.Length {
        return errors.New("number is too large")
    }
    p.Discard.NewCard(p.Hand.Cards[toDiscard])
    err := p.Hand.RmCard(toDiscard)
    if err != nil {
        return err
    }
    return nil
}

func (p *Player) Play(toPlay int) error {
    if toPlay >= p.Hand.Length {
        return errors.New("number is too large")
    }
    p.Hand.Cards[toPlay].Play()
    err := p.DiscardCard(toPlay)
    if err != nil {
        return err
    }
    return nil
}

func (g *Group) Shuffle() error {
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

func (g *Group) Draw() (error, Card) {
    if g.Length == 0 {
        return errors.New("group is empty"), *new(Card)
    }
    drawnCard := g.Cards[g.Length-1]
    g.Cards = g.Cards[:g.Length-1]
    g.Length--
    return nil, drawnCard
}

func (g *Group) RmCard(i int) error {
    if i >= g.Length {
        return errors.New("number is too large")
    }
    g.Cards = slices.Delete(g.Cards, i, i+1)
    g.Length--
    return nil
}

func (g *Group) NewCard(c Card) {
    c.GroupType = g.Type
    c.GroupLocation = g.Location
    g.Cards = append(g.Cards, c)
    g.Length++
}

func (c *Card) Play() error {
    if c.Type == "nailS" {
        if c.GroupLocation == "Player1" {
            c.Damage += c.HandlerObj.Player1.NailBoost.Cards[c.HandlerObj.Player1.NailBoost.Length-1].NailPlus
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            if c.HandlerObj.Player1.WeakPointS {
                c.Damage = c.Damage * 2
            }
            if c.HandlerObj.Player1.WeakPointS && c.HandlerObj.Player2.WeakPointP {
                r := rand.Intn(3)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                } else if r == 1 || r == 2 {
                    c.Damage = c.Damage/2
                }
            }
            if c.HandlerObj.Player2.WeakPointP && !c.HandlerObj.Player1.WeakPointS {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player2.Health -= c.Damage
                c.HandlerObj.Player1.Soul += c.Soul
            }
            c.HandlerObj.Player1.WeakPointS = false
            c.HandlerObj.Player2.WeakPointP = false
        }
        if c.GroupLocation == "Player2" {
            c.Damage += c.HandlerObj.Player2.NailBoost.Cards[c.HandlerObj.Player2.NailBoost.Length-1].NailPlus
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            if c.HandlerObj.Player2.WeakPointS {
                c.Damage = c.Damage * 2
            }
            if c.HandlerObj.Player2.WeakPointS && c.HandlerObj.Player1.WeakPointP {
                r := rand.Intn(3)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                } else if r == 1 || r == 2 {
                    c.Damage = c.Damage/2
                }
            }
            if c.HandlerObj.Player1.WeakPointP && !c.HandlerObj.Player2.WeakPointS {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player1.Health -= c.Damage
                c.HandlerObj.Player2.Soul += c.Soul
            }
            c.HandlerObj.Player2.WeakPointS = false
            c.HandlerObj.Player1.WeakPointP = false
        }
    } else if c.Type == "nailA" {
        if c.GroupLocation == "Player1" {
            c.Damage += c.HandlerObj.Player1.NailBoost.Cards[c.HandlerObj.Player1.NailBoost.Length-1].NailPlus
            if c.HandlerObj.Player1.ArtUses == 1 {
                c.Damage -= 3
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            if c.HandlerObj.Player1.WeakPointS {
                c.Damage = c.Damage * 2
            }
            if c.HandlerObj.Player1.WeakPointS && c.HandlerObj.Player2.WeakPointP {
                r := rand.Intn(3)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                } else if r == 1 || r == 2 {
                    c.Damage = c.Damage/2
                }
            }
            if c.HandlerObj.Player2.WeakPointP && !c.HandlerObj.Player1.WeakPointS {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player2.Health -= c.Damage
                c.HandlerObj.Player1.Soul += c.Soul
            }
            c.HandlerObj.Player1.ArtUses++
            c.HandlerObj.Player1.WeakPointS = false
            c.HandlerObj.Player2.WeakPointP = false
        }
        if c.GroupLocation == "Player2" {
            c.Damage += c.HandlerObj.Player2.NailBoost.Cards[c.HandlerObj.Player2.NailBoost.Length-1].NailPlus
            if c.HandlerObj.Player1.ArtUses == 1 {
                c.Damage -= 3
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            if c.HandlerObj.Player2.WeakPointS {
                c.Damage = c.Damage * 2
            }
            if c.HandlerObj.Player2.WeakPointS && c.HandlerObj.Player1.WeakPointP {
                r := rand.Intn(3)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                } else if r == 1 || r == 2 {
                    c.Damage = c.Damage/2
                }
            }
            if c.HandlerObj.Player1.WeakPointP && !c.HandlerObj.Player2.WeakPointS {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player1.Health -= c.Damage
                c.HandlerObj.Player2.Soul += c.Soul
            }
            c.HandlerObj.Player2.ArtUses++
            c.HandlerObj.Player2.WeakPointS = false
            c.HandlerObj.Player1.WeakPointP = false
        }
    } else if c.Type == "spell" {
        if c.GroupLocation == "Player1" {
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            if c.HandlerObj.Player1.WeakPointS {
                c.Damage = c.Damage * 2
            }
            if c.HandlerObj.Player1.WeakPointS && c.HandlerObj.Player2.WeakPointP {
                r := rand.Intn(3)
                if r == 0 {
                    c.Damage = 0
                } else if r == 1 || r == 2 {
                    c.Damage = c.Damage/2
                }
            }
            if c.HandlerObj.Player2.WeakPointP && !c.HandlerObj.Player1.WeakPointS {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player2.Health -= c.Damage
            }
            c.HandlerObj.Player1.Soul -= c.Soul
            c.HandlerObj.Player1.WeakPointS = false
            c.HandlerObj.Player2.WeakPointP = false
        }
        if c.GroupLocation == "Player2" {
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, c.HandlerObj)
            }
            if c.HandlerObj.Player2.WeakPointS {
                c.Damage = c.Damage * 2
            }
            if c.HandlerObj.Player2.WeakPointS && c.HandlerObj.Player1.WeakPointP {
                r := rand.Intn(3)
                if r == 0 {
                    c.Damage = 0
                } else if r == 1 || r == 2 {
                    c.Damage = c.Damage/2
                }
            }
            if c.HandlerObj.Player1.WeakPointP && !c.HandlerObj.Player2.WeakPointS {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player1.Health -= c.Damage
            }
            c.HandlerObj.Player2.Soul -= c.Soul
            c.HandlerObj.Player2.WeakPointS = false
            c.HandlerObj.Player1.WeakPointP = false
        }
    } else if c.Type == "precept" {
        c.Effect(c, c.HandlerObj)
    } else if c.Type == "nailT" {
        if c.GroupLocation == "Player1" {
            c.HandlerObj.Player1.NailEquip(*c)
        }
        if c.GroupLocation == "Player2" {
            c.HandlerObj.Player2.NailEquip(*c)
        }
    } else {
        return errors.New("wrong card type")
    }
    return nil
}
