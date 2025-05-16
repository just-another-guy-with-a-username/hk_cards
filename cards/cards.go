package cards

import (
    "hk_cards/types"
    "math"
)

func NewNailSlash(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailS"
    c.Name = "Nail Slash"
    c.Damage = 1
    c.Soul = 1
    c.FailChance = float32(0)
    c.HandlerObj = h
    return *c
}

func NewGreatSlash(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailA"
    c.Name = "Great Slash"
    c.Damage = 5
    c.FailChance = float32(0)
    c.AddsDodge = false
    c.HandlerObj = h
    return *c
}

func NewDashSlash(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailA"
    c.Name = "Dash Slash"
    c.Damage = 3
    c.FailChance = float32(0)
    c.AddsDodge = true
    c.HandlerObj = h
    return *c
}

func NewCycloneSlash(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailA"
    c.Name = "Cyclone Slash"
    c.Damage = 8
    c.FailChance = float32(0.5)
    c.AddsDodge = false
    c.HandlerObj = h
    return *c
}

func NewVengefulSpirit(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "spell"
    c.Name = "Vengeful Spirit"
    c.Damage = 8
    c.Soul = 2
    c.FailChance = float32(0)
    c.AddsDodge = false
    c.HandlerObj = h
    return *c
}

func NewShadeSoul(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "spell"
    c.Name = "Shade Soul"
    c.Damage = 16
    c.Soul = 4
    c.FailChance = float32(0)
    c.AddsDodge = false
    c.HandlerObj = h
    return *c
}

func NewDesolateDive(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "spell"
    c.Name = "Desolate Dive"
    c.Damage = 6
    c.Soul = 2
    c.FailChance = float32(0)
    c.AddsDodge = true
    c.HandlerObj = h
    return *c
}

func NewDescendingDark(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "spell"
    c.Name = "Descending Dark"
    c.Damage = 12
    c.Soul = 4
    c.FailChance = float32(0)
    c.AddsDodge = true
    c.HandlerObj = h
    return *c
}

func NewHowlingWraiths(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "spell"
    c.Name = "Howling Wraiths"
    c.Damage = 10
    c.Soul = 2
    c.FailChance = float32(0.5)
    c.AddsDodge = false
    c.HandlerObj = h
    return *c
}

func NewAbyssShriek(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "spell"
    c.Name = "Abyss Shriek"
    c.Damage = 20
    c.Soul = 4
    c.FailChance = float32(0.5)
    c.AddsDodge = false
    c.HandlerObj = h
    return *c
}

func NewDreamNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailD"
    c.Name = "Dream Nail"
    c.HandlerObj = h
    return *c
}

func NewOldNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Old Nail"
    c.NailPlus = 2
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewSharpenedNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Sharpened Nail"
    c.NailPlus = 4
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewChanneledNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Channeled Nail"
    c.NailPlus = 6
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewCoiledNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Coiled Nail"
    c.NailPlus = 8
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewPureNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Pure Nail"
    c.NailPlus = 10
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewRadiantOutburst(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "infection"
    c.Name = "Radiant Outburst"
    c.HandlerObj = h
    return *c
}

func NewVoidCovering(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "void"
    c.Name = "Void Covering"
    c.TurnsLeft = 3
    c.HandlerObj = h
    return *c
}

func ShamanStone(c *types.Card, ch *types.Card, h *types.Handler) error {
    if c.Type == "spell" {
        if c.GroupLocation == ch.GroupLocation {
            c.Damage = int(math.Floor(float64(c.Damage) * float64(1.5)))
        }
    }
    return nil
}

func NewShamanStone(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "charm"
    c.Name = "Shaman Stone"
    c.Effect = ShamanStone
    c.HandlerObj = h
    return *c
}

func StrikeTheFoesWeakPoint(c *types.Card, ch *types.Card, h *types.Handler) error {
    if ch.GroupLocation == "Player1" {
        h.Player1.WeakPointS = true
    }
    if ch.GroupLocation == "Player2" {
        h.Player2.WeakPointS = true
    }
    return nil
}

func ProtectYourOwnWeakPoint(c *types.Card, ch *types.Card, h *types.Handler) error {
    if ch.GroupLocation == "Player1" {
        h.Player1.WeakPointP = true
    }
    if ch.GroupLocation == "Player2" {
        h.Player2.WeakPointP = true
    }
    return nil
}
