package cards

import (
    "hk_cards/types"
    "math"
)

func ShamanStone (c *types.Card, ch *types.Card, h *types.Handler) error {
    if c.Type == "spell" {
        if c.GroupLocation == ch.GroupLocation {
            c.Damage = int(math.Floor(float64(c.Damage) * float64(1.5)))
        }
    }
    return nil
}
