package types

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
    "errors"
    "math/rand"
    "slices"
    "time"
//    "fmt"
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
    DodgeChance bool
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
    ImagePath     string
    Type          string
    GroupType     string
    GroupLocation string
    Effect        func(*Card, *Card, *Handler, fyne.Window) error
    NotchCost     int
    TurnsLeft     int
    NailPlus      int
    Damage        int
    EndDamage     int
    Soul          int
    FailChance    float32
    AddsDodge     bool
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

func (p *Player) DrawHand() error {
    for i := p.Hand.Length; i < 7; i++ {
        err := p.Draw()
        if err != nil {
            return err
        }
    }
    return nil
}

func (p *Player) CharmEquip(i int) error {
    NotchesUsed := 0
    for _, charm := range(p.Charms.Cards) {
        NotchesUsed += charm.NotchCost
    }
    if i >= p.Hand.Length {
        return errors.New("number is too large")
    }
    c := p.Hand.Cards[i]
    if p.Notches - NotchesUsed >= c.NotchCost {
        p.Charms.NewCard(c)
    } else if p.Notches - NotchesUsed > 0 {
        p.Charms.NewCard(c)
        p.Overcharmed = true
    } else {
        return errors.New("all charm notches filled")
    }
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

func (p *Player) Play(toPlay int, w fyne.Window) error {
    if toPlay >= p.Hand.Length {
        return errors.New("number is too large")
    }
    if p.Hand.Cards[toPlay].Name == "Do Not Dream" {
        err, c := p.Hand.GetCard(toPlay)
        if err != nil {
            return err
        }
        err = c.Play(w)
        if err != nil {
            slices.Insert(p.Hand.Cards, toPlay, c)
            p.Hand.Length++
            return err
        }
        p.Discard.NewCard(c)
        return nil
    } else {
        err := p.Hand.Cards[toPlay].Play(w)
        if err != nil {
            return err
        }
        p.DiscardCard(toPlay)
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

func (g *Group) GetCard(i int) (error, Card) {
    if i >= g.Length {
        return errors.New("number is too large"), *new(Card)
    }
    ToGet := g.Cards[i]
    g.Cards = slices.Delete(g.Cards, i, i+1)
    g.Length--
    return nil, ToGet
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

func (c *Card) Play(w fyne.Window) error {
    if c.Type == "nailS" {
        if c.GroupLocation == "Player1" {
            if c.HandlerObj.Player1.NailBoost != nil {
                c.Damage += c.HandlerObj.Player1.NailBoost.Cards[c.HandlerObj.Player1.NailBoost.Length-1].NailPlus
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
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
            if c.HandlerObj.Player2.DodgeChance {
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
            if c.HandlerObj.Player2.NailBoost != nil {
                c.Damage += c.HandlerObj.Player2.NailBoost.Cards[c.HandlerObj.Player2.NailBoost.Length-1].NailPlus
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
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
            if c.HandlerObj.Player1.DodgeChance {
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
            if c.HandlerObj.Player1.NailBoost != nil {
                c.Damage += c.HandlerObj.Player1.NailBoost.Cards[c.HandlerObj.Player1.NailBoost.Length-1].NailPlus
            }
            if c.HandlerObj.Player1.ArtUses == 1 {
                c.Damage -= 3
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
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
            if c.HandlerObj.Player2.DodgeChance {
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
            if c.AddsDodge {
                c.HandlerObj.Player1.DodgeChance = true
            }
            c.HandlerObj.Player2.DodgeChance = false
            c.HandlerObj.Player1.ArtUses++
            c.HandlerObj.Player1.WeakPointS = false
            c.HandlerObj.Player2.WeakPointP = false
        }
        if c.GroupLocation == "Player2" {
            if c.HandlerObj.Player2.NailBoost != nil {
                c.Damage += c.HandlerObj.Player2.NailBoost.Cards[c.HandlerObj.Player2.NailBoost.Length-1].NailPlus
            }
            if c.HandlerObj.Player1.ArtUses == 1 {
                c.Damage -= 3
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
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
            if c.HandlerObj.Player1.DodgeChance {
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
            if c.AddsDodge {
                c.HandlerObj.Player2.DodgeChance = true
            }
            c.HandlerObj.Player1.DodgeChance = false
            c.HandlerObj.Player2.ArtUses++
            c.HandlerObj.Player2.WeakPointS = false
            c.HandlerObj.Player1.WeakPointP = false
        }
    } else if c.Type == "spell" {
        if c.GroupLocation == "Player1" {
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
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
            if c.HandlerObj.Player2.DodgeChance {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player2.Health -= c.Damage
            }
            if c.AddsDodge {
                c.HandlerObj.Player1.DodgeChance = true
            }
            c.HandlerObj.Player2.DodgeChance = false
            c.HandlerObj.Player1.Soul -= c.Soul
            c.HandlerObj.Player1.WeakPointS = false
            c.HandlerObj.Player2.WeakPointP = false
        }
        if c.GroupLocation == "Player2" {
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
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
            if c.HandlerObj.Player1.DodgeChance {
                r := rand.Intn(1)
                if r == 0 {
                    c.Damage = 0
                    c.Soul = 0
                }
            }
            if rand.Float32() >= c.FailChance {
                c.HandlerObj.Player1.Health -= c.Damage
            }
            if c.AddsDodge {
                c.HandlerObj.Player2.DodgeChance = true
            }
            c.HandlerObj.Player1.DodgeChance = false
            c.HandlerObj.Player2.Soul -= c.Soul
            c.HandlerObj.Player2.WeakPointS = false
            c.HandlerObj.Player1.WeakPointP = false
        }
    } else if c.Type == "precept" {
        c.Effect(c, new(Card), c.HandlerObj, w)
    } else if c.Type == "nailT" {
        if c.GroupLocation == "Player1" {
            c.HandlerObj.Player1.NailEquip(*c)
        }
        if c.GroupLocation == "Player2" {
            c.HandlerObj.Player2.NailEquip(*c)
        }
    } else if c.Type == "nailD" {
        if c.GroupLocation == "Player1" {
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            c.HandlerObj.Player1.Soul += c.Soul
            cardsSlice := []fyne.CanvasObject{}
            for _, card := range(c.HandlerObj.Player2.Hand.Cards) {
                cardsSlice = append(cardsSlice, widget.NewLabel(card.Name))
            }
            cardsPresent := container.NewCenter(container.New(layout.NewHBoxLayout(), cardsSlice...))
            w.SetContent(cardsPresent)
            time.Sleep(time.Second * 2)
            cardsPresent.RemoveAll()
            cardsPresent.Refresh()
        }
        if c.GroupLocation == "Player2" {
            for _, charm := range c.HandlerObj.Player2.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            for _, charm := range c.HandlerObj.Player1.Charms.Cards {
                charm.Effect(c, &charm, c.HandlerObj, w)
            }
            c.HandlerObj.Player2.Soul += c.Soul
            cardsSlice := []fyne.CanvasObject{}
            for _, card := range(c.HandlerObj.Player1.Hand.Cards) {
                cardsSlice = append(cardsSlice, widget.NewLabel(card.Name))
            }
            cardsPresent := container.NewCenter(container.New(layout.NewHBoxLayout(), cardsSlice...))
            w.SetContent(cardsPresent)
            time.Sleep(time.Second * 2)
            cardsPresent.RemoveAll()
            cardsPresent.Refresh()
        }
    } else {
        return errors.New("wrong card type")
    }
    return nil
}
