package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
    // findfont "github.com/flopp/go-findfont"
    "github.com/golang/freetype/truetype"
    "os"
    "strings"
)

func init() { // 解决中文显示乱码问题
    //fontPath, err := findfont.Find("DouyinSansBold.ttf")
    //if err != nil {
    //    panic(err)
    //}
    fontPath := ".\\fonts\\DouyinSansBold.ttf"
    // load the font with the freetype library
    fontData, err := os.ReadFile(fontPath)
    if err != nil {
        panic(err)
    }
    _, err = truetype.Parse(fontData)
    if err != nil {
        panic(err)
    }
    os.Setenv("FYNE_FONT", fontPath)
}

func main() {
    a := app.NewWithID("whencome.github.com") // a id required
    win := a.NewWindow("Preferences")
    entry1 := widget.NewEntry()
    btn1 := widget.NewButton("save value", func() {
        txt := strings.TrimSpace(entry1.Text)
        if txt == "" {
            dialog.ShowInformation("TIPS", "please enter a string", win)
            return
        }
        a.Preferences().SetString("key1", txt)
    })
    btn2 := widget.NewButton("show value", func() {
        txt := a.Preferences().StringWithFallback("key1", "<empty data>")
        dialog.ShowInformation("Saved Data", txt, win)
    })
    vbox := container.NewVBox(entry1, btn1, btn2)
    win.SetContent(vbox)
    win.ShowAndRun()
}
