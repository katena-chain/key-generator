//go:generate go-bindata -pkg assets -o ../assets/assets.go ../assets/
package main

import (
    "fmt"

    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/theme"
    "fyne.io/fyne/widget"

    "github.com/katena-chain/key-generator/assets"
    "github.com/katena-chain/key-generator/gui/events"
)

func main() {

    application := app.New()
    application.Settings().SetTheme(theme.LightTheme())

    mainWindow := application.NewWindow("Katena Key Generator")

    // Try to load the Katena icon
    icon, err := assets.Asset("../assets/logo-katena.png")
    if err != nil {
        fmt.Println(fmt.Sprintf("unable to load icon: %s", err.Error()))
    } else {
        mainWindow.SetIcon(fyne.NewStaticResource("TransChain icon", icon))
    }

    pubKeyEntry := widget.NewEntry()
    privKeyEntry := widget.NewEntry()

    mainWindow.SetContent(widget.NewVBox(
        widget.NewLabel("Public key : "),
        pubKeyEntry,
        widget.NewLabel("Private key : "),
        privKeyEntry,
        widget.NewButton("Generate X25519 keys", events.OnClickGenX25519Keys(pubKeyEntry, privKeyEntry)),
        widget.NewButton("Generate ED25519 keys", events.OnClickGenED25519Keys(pubKeyEntry, privKeyEntry)),
        widget.NewButton("Save keys", events.OnClickSaveKeys(pubKeyEntry, privKeyEntry)),
        widget.NewButton("Quit", events.OnClickQuit(application)),
    ))

    mainWindow.Resize(fyne.Size{
        Width:  800,
        Height: 250,
    })
    mainWindow.CenterOnScreen()
    mainWindow.ShowAndRun()

}
