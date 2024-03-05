package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.NewWithID("whencome.github.com")
    win := a.NewWindow("Menu")
    win.SetMaster()
    win.Resize(fyne.Size{
        Width:  600,
        Height: 400,
    })
    // 增加关闭前确认提示
    win.SetCloseIntercept(func() {
        // 创建一个退出确认对话框
        confirmation := dialog.NewConfirm("Confirmation", "Are you sure you want to exit?", func(confirm bool) {
            if confirm {
                // 如果用户确认，关闭窗口
                win.Close()
            }
            // 如果用户取消，不执行任何操作，窗口保持打开状态
        }, win)
        confirmation.Show()
    })
    // 设置菜单
    win.SetMainMenu(fyne.NewMainMenu(
        fyne.NewMenu(
            "File",
            fyne.NewMenuItem("Open", func() {
                fmt.Println("Open a file")
            }),
            fyne.NewMenuItem("Save", func() {
                fmt.Println("Save a file")
            }),
            // 将分隔符作为菜单项添加
            fyne.NewMenuItemSeparator(),
            fyne.NewMenuItem("Print", func() {
                fmt.Println("Print a file")
            }),
        ),
        fyne.NewMenu(
            "Help",
            fyne.NewMenuItem("About", func() {
                fmt.Println("help -> about")
            }),
            fyne.NewMenuItem("Update", func() {
                fmt.Println("help -> update")
            },
            ),
        )))
    // 设置内容
    win.SetContent(widget.NewLabel("menu demo"))
    win.ShowAndRun()
}
