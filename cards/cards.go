package cards

import (
    "fyne.io/fyne/v2"
//    "fyne.io/fyne/v2/app"
//    "fyne.io/fyne/v2/widget"
//    "fyne.io/fyne/v2/container"
//    "fyne.io/fyne/v2/canvas"
//    "fyne.io/fyne/v2/layout"
    "hk_cards/types"
    "math"
    "slices"
//    "fmt"
)

func NewNailSlash(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailS"
    c.Name = "Nail Slash"
    c.ImagePath = "cards/images/nail_slash.png"
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
    c.ImagePath = "cards/images/great_slash.png"
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
    c.ImagePath = "cards/images/dash_slash.png"
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
    c.ImagePath = "cards/images/cyclone_slash.png"
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
    c.ImagePath = "cards/images/vengeful_spirit.png"
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
    c.ImagePath = "cards/images/shade_soul.png"
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
    c.ImagePath = "cards/images/desolate_dive.png"
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
    c.ImagePath = "cards/images/descending_dark.png"
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
    c.ImagePath = "cards/images/howling_wraiths.png"
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
    c.ImagePath = "cards/images/abyss_shriek.png"
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
    c.ImagePath = "cards/images/dream_nail.png"
    c.Soul = 2
    c.HandlerObj = h
    return *c
}

func NewOldNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Old Nail"
    c.ImagePath = "cards/images/old_nail.png"
    c.NailPlus = 2
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewSharpenedNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Sharpened Nail"
    c.ImagePath = "cards/images/sharpened_nail.png"
    c.NailPlus = 4
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewChanneledNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Channeled Nail"
    c.ImagePath = "cards/images/channeled_nail.png"
    c.NailPlus = 6
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewCoiledNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Coiled Nail"
    c.ImagePath = "cards/images/coiled_nail.png"
    c.NailPlus = 8
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewPureNail(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "nailT"
    c.Name = "Pure Nail"
    c.ImagePath = "cards/images/pure_nail.png"
    c.NailPlus = 10
    c.TurnsLeft = 2
    c.HandlerObj = h
    return *c
}

func NewRadiantOutburst(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "infection"
    c.Name = "Radiant Outburst"
    c.ImagePath = "cards/images/radiant_outburst.png"
    c.HandlerObj = h
    return *c
}

func NewVoidCovering(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "void"
    c.Name = "Void Covering"
    c.ImagePath = "cards/images/void_covering.png"
    c.TurnsLeft = 3
    c.HandlerObj = h
    return *c
}

func ShamanStone(c *types.Card, ch *types.Card, h *types.Handler, main *fyne.Container) error {
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
    c.ImagePath = "cards/images/shaman_stone.png"
    c.Effect = ShamanStone
    c.HandlerObj = h
    return *c
}

func DoNotDream(c *types.Card, ch *types.Card, h *types.Handler, main *fyne.Container) error {
    for i, card := range(slices.Backward(h.Player1.Hand.Cards)) {
        if card.Type == "nailD" {
            h.Player1.DiscardCard(i)
        }
    }
    for i, card := range(slices.Backward(h.Player2.Hand.Cards)) {
        if card.Type == "nailD" {
            h.Player2.DiscardCard(i)
        }
    }
    return nil
}

func NewDoNotDream(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "precept"
    c.Name = "Do Not Dream"
    c.ImagePath = "cards/images/do_not_dream.png"
    c.Effect = DoNotDream
    c.HandlerObj = h
    return *c
}

func StrikeTheFoesWeakPoint(c *types.Card, ch *types.Card, h *types.Handler, main *fyne.Container) error {
    if ch.GroupLocation == "Player1" {
        h.Player1.WeakPointS = true
    }
    if ch.GroupLocation == "Player2" {
        h.Player2.WeakPointS = true
    }
    return nil
}

func NewStrikeTheFoesWeakPoint(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "precept"
    c.Name = "Strike the Foe's Weak Point"
    c.ImagePath = "cards/images/strike_the_foes_weak_point.png"
    c.Effect = StrikeTheFoesWeakPoint
    c.HandlerObj = h
    return *c
}

func ProtectYourOwnWeakPoint(c *types.Card, ch *types.Card, h *types.Handler, main *fyne.Container) error {
    if ch.GroupLocation == "Player1" {
        h.Player1.WeakPointP = true
    }
    if ch.GroupLocation == "Player2" {
        h.Player2.WeakPointP = true
    }
    return nil
}

func NewProtectYourOwnWeakPoint(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "precept"
    c.Name = "Protect Your Own Weak Point"
    c.ImagePath = "cards/images/protect_your_own_weak_point.png"
    c.Effect = ProtectYourOwnWeakPoint
    c.HandlerObj = h
    return *c
}

func EatAsMuchAsYouCan(c *types.Card, ch *types.Card, h *types.Handler, main *fyne.Container) error {
    if ch.GroupLocation == "Player1" {
        h.Player1.Health += 10
        h.Player2.Health += 5
    }
    if ch.GroupLocation == "Player2" {
        h.Player2.Health += 10
        h.Player1.Health += 5
    }
    if h.Player1.Health > 100 {
        h.Player1.Health = 100
    }
    if h.Player2.Health > 100 {
        h.Player2.Health = 100
    }
    return nil
}

func NewEatAsMuchAsYouCan(h *types.Handler) types.Card {
    c := new(types.Card)
    c.Type = "precept"
    c.Name = "Eat as Much as You Can"
    c.ImagePath = "cards/images/eat_as_much_as_you_can.png"
    c.Effect = EatAsMuchAsYouCan
    c.HandlerObj = h
    return *c
}

