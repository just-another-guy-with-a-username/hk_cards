package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
//    "fyne.io/fyne/v2/layout"
    "hk_cards/types"
    "hk_cards/cards"
    "github.com/fstanis/screenresolution"
//    "fmt"
    "time"
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
    for n := 0; n < 50; n++ {
        deck.NewCard(cards.NewNailSlash(h))
    }
    deck.NewCard(cards.NewNailSlash(h))
    deck.NewCard(cards.NewDreamNail(h))
    deck.NewCard(cards.NewDoNotDream(h))
    deck.NewCard(cards.NewDreamNail(h))
    deck.NewCard(cards.NewNailSlash(h))
    deck.NewCard(cards.NewNailSlash(h))
    deck.NewCard(cards.NewNailSlash(h))
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

func displayCards(main *fyne.Container, p int, h *types.Handler, height int, width int) {
    if p == 1 {
        cardImgs := []fyne.CanvasObject{}
        cardBtns := []fyne.CanvasObject{}
        cardBtnsP := []*widget.Button{}
        for i, c := range(h.Player1.Hand.Cards) {
            cImage := canvas.NewImageFromFile(c.ImagePath)
            cImage.Move(fyne.NewPos(float32(width*(21-(3*i-h.Player1.Hand.Length))/40), float32(height-width*3/40)))
            cImage.Resize(fyne.NewSize(float32(width/20), float32(3*width/40)))
            cardImgs = append(cardImgs, cImage)
            cButton := widget.NewButton("", func() {CardEnlarge(p, i, h, cardImgs, cardBtns, cardBtnsP, main, height, width)})
            cButton.Move(fyne.NewPos(float32(width*(21-(3*i-h.Player1.Hand.Length))/40), float32(height-width*3/40)))
            cButton.Resize(fyne.NewSize(float32(width/20), float32(3*width/40)))
            cardBtns = append(cardBtns, cButton)
            cardBtnsP = append(cardBtnsP, cButton)
            main.Add(cButton)
            main.Add(cImage)
        }
        for i, _ := range(h.Player2.Hand.Cards) {
            cImage := canvas.NewImageFromFile("cards/images/card_back.png")
            cImage.Move(fyne.NewPos(float32(width*(20-(3*i-h.Player2.Hand.Length))/40), float32(0)))
            cImage.Resize(fyne.NewSize(float32(width/20), float32(3*width/40)))
            cardImgs = append(cardImgs, cImage)
            main.Add(cImage)
        }
    } else if p == 2 {
        cardImgs := []fyne.CanvasObject{}
        cardBtns := []fyne.CanvasObject{}
        cardBtnsP := []*widget.Button{}
        for i, c := range(h.Player2.Hand.Cards) {
            cImage := canvas.NewImageFromFile(c.ImagePath)
            cImage.Move(fyne.NewPos(float32(width*(21-(3*i-h.Player2.Hand.Length))/40), float32(height-width*3/40)))
            cImage.Resize(fyne.NewSize(float32(width/20), float32(3*width/40)))
            cardImgs = append(cardImgs, cImage)
            cButton := widget.NewButton("", func() {CardEnlarge(p, i, h, cardImgs, cardBtns, cardBtnsP, main, height, width)})
            cButton.Move(fyne.NewPos(float32(width*(21-(3*i-h.Player2.Hand.Length))/40), float32(height-width*3/40)))
            cButton.Resize(fyne.NewSize(float32(width/20), float32(3*width/40)))
            cardBtns = append(cardBtns, cButton)
            cardBtnsP = append(cardBtnsP, cButton)
            main.Add(cButton)
            main.Add(cImage)
        }
        for i, _ := range(h.Player1.Hand.Cards) {
            cImage := canvas.NewImageFromFile("cards/images/card_back.png")
            cImage.Move(fyne.NewPos(float32(width*(20-(3*i-h.Player1.Hand.Length))/40), float32(0)))
            cImage.Resize(fyne.NewSize(float32(width/20), float32(3*width/40)))
            cardImgs = append(cardImgs, cImage)
            main.Add(cImage)
        }
    }
}

func CardEnlarge(p int, i int, h *types.Handler, CardImgs []fyne.CanvasObject, CardBtns []fyne.CanvasObject, CardBtnsP []*widget.Button, main *fyne.Container, height int, width int) {
    for _, Btn := range(CardBtnsP) {
        Btn.Disable()
        Btn.Hide()
    }
    PosAnim := canvas.NewPositionAnimation(CardImgs[i].Position(), fyne.NewPos(float32(width/2-height/6), float32(height/4)), time.Millisecond*400, func(p fyne.Position) {
        CardImgs[i].Move(p)
    })
    SizeAnim := canvas.NewSizeAnimation(CardImgs[i].Size(), fyne.NewSize(float32(height/3), float32(height/2)), time.Millisecond*400, func(s fyne.Size) {
        CardImgs[i].Resize(s)
    })
    PosAnim.Start()
    SizeAnim.Start()
    CancelBtn := widget.NewButton("Cancel", func() {CancelPlay(p, i, h, CardImgs, CardBtns, CardBtnsP, main, height, width)})
    CancelBtn.Move(fyne.NewPos(float32(width*17/60), float32(height*19/40)))
    CancelBtn.Resize(fyne.NewSize(float32(width/10), float32(height/20)))
    PlayBtn := widget.NewButton("Play", func() {PlayCardHandling(p, i, h, CardImgs, CardBtns, main, height, width)})
    PlayBtn.Move(fyne.NewPos(float32(width*37/60), float32(height*19/40)))
    PlayBtn.Resize(fyne.NewSize(float32(width/10), float32(height/20)))
    CardBtns = append(CardBtns, PlayBtn, CancelBtn)
    main.Add(PlayBtn)
    main.Add(CancelBtn)
}

func CancelPlay(p int, i int, h *types.Handler, CardImgs []fyne.CanvasObject, CardBtns []fyne.CanvasObject, CardBtnsP []*widget.Button, main *fyne.Container, height int, width int) {
    main.Remove(CardBtns[len(CardBtns)-1])
    main.Remove(CardBtns[len(CardBtns)-2])
    CardBtns = CardBtns[:len(CardBtns)-2]
    if p == 1 {
        PosAnim := canvas.NewPositionAnimation(CardImgs[i].Position(), fyne.NewPos(float32(width*(21-(3*i-h.Player1.Hand.Length))/40), float32(height-width*3/40)), time.Millisecond*400, func(p fyne.Position) {
            CardImgs[i].Move(p)
        })
        SizeAnim := canvas.NewSizeAnimation(CardImgs[i].Size(), fyne.NewSize(float32(width/20), float32(3*width/40)), time.Millisecond*400, func(s fyne.Size) {
            CardImgs[i].Resize(s)
        })
        PosAnim.Start()
        SizeAnim.Start()
    }
    if p == 2 {
        PosAnim := canvas.NewPositionAnimation(CardImgs[i].Position(), fyne.NewPos(float32(width*(21-(3*i-h.Player2.Hand.Length))/40), float32(height-width*3/40)), time.Millisecond*400, func(p fyne.Position) {
            CardImgs[i].Move(p)
        })
        SizeAnim := canvas.NewSizeAnimation(CardImgs[i].Size(), fyne.NewSize(float32(width/20), float32(3*width/40)), time.Millisecond*400, func(s fyne.Size) {
            CardImgs[i].Resize(s)
        })
        PosAnim.Start()
        SizeAnim.Start()
    }
    go func() {
        time.Sleep(time.Millisecond*400)
        for _, Btn := range(CardBtnsP) {
            Btn.Enable()
            Btn.Show()
        }
    }()
}

func PlayCardHandling(p int, i int, h *types.Handler, CardImgs []fyne.CanvasObject, CardBtns []fyne.CanvasObject, main *fyne.Container, height int, width int) {
    if p == 1 {
        h.Player1.Play(i, main, height, width)
        p = 2
    } else if p == 2 {
        h.Player2.Play(i, main, height, width)
        p = 1
    }
    for _, c := range(CardImgs) {
        main.Remove(c)
    }
    for _, c := range(CardBtns) {
        main.Remove(c)
    }
    status := Turn(main, p, h, height, width)
    if status != 0 {
    }
}

func StartGame(main *fyne.Container, p int, h *types.Handler, start *widget.Button, height int, width int) {
    main.Remove(start)
    main.Refresh()
    Turn(main, p, h, height, width)
}

func Turn(main *fyne.Container, p int, h *types.Handler, height int, width int) int {
    if p == 1 {
        h.Player1.DrawHand()
    }
    if p == 2 {
        h.Player2.DrawHand()
    }
    displayCards(main, p, h, height, width)
    if (h.Player2.Deck.Length == 0 && h.Player2.Hand.Length == 0) || h.Player2.Health <= 0 {
        return 1
    }
    if (h.Player1.Deck.Length == 0 && h.Player1.Hand.Length == 0) || h.Player1.Health <= 0 {
        return 2
    }
    return 0
}

func WinClose(w fyne.Window) {
    w.Close()
}

func main() {
    w_height := screenresolution.GetPrimary().Height
    w_width := screenresolution.GetPrimary().Width
    game := initGame(1, 1)
    a := app.New()
    w := a.NewWindow("Colosseum Of Fools")
    AllContent := container.NewWithoutLayout()
    w.SetFullScreen(true)
    w.SetFixedSize(true)
    StartButton := widget.NewButton("Start Game", func() {a = a})
    StartButton.OnTapped = func() {StartGame(AllContent, 1, &game, StartButton, w_height, w_width)}
    StartButton.Resize(fyne.NewSize(150, 50))
    StartButton.Move(fyne.NewPos(float32(w_width/2-75), float32(w_height/2-25)))
    CloseButton := widget.NewButton("Close", func() {WinClose(w)})
    CloseButton.Resize(fyne.NewSize(150, 50))
    CloseButton.Move(fyne.NewPos(0, 0))
    AllContent.Add(StartButton)
    AllContent.Add(CloseButton)
    w.SetContent(AllContent)
    w.ShowAndRun()
}
