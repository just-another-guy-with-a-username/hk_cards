package types

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
    Cards  []Card
    Length int
}

type Card struct {
    Name   string
    Type   string
    Effect func(*Handler) error
    Turns  int
}

func (c Card) play (h *Handler) error {
    err := c.effect(h)
    if err != nil {
        return err
    }
    return nil
}
