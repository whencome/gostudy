package main

import (
    "fmt"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

func main() {
    fmt.Println("start ... ")
    a := app.New()
    fmt.Println("app created ... ")
    w1 := a.NewWindow("window 1")
    fmt.Println("window1 created ... ")
    btn := widget.NewButton("show window 2", func() {
        fmt.Println("button clicked ... ")
        w2 := a.NewWindow("window 2")
        fmt.Println("window2 created ... ")
        w2.SetContent(widget.NewLabel("this is window 2"))
        w2.Show()
        fmt.Println("show window2 ... ")
    })
    w1.SetContent(btn)
    fmt.Println("show window1 ... ")
    w1.Show()
    a.Run()
    fmt.Println("app exited ... ")
}
