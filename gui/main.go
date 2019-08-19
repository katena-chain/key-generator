//go:generate go-bindata -pkg assets -o ../assets/assets.go ../assets/
package main

import (
    "fmt"

    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/widget"
    "github.com/gen2brain/dlgs"

    "github.com/katena-chain/key-generator/assets"
    "github.com/katena-chain/key-generator/libs"
)

func main() {
    RunStartGui()
}

func RunStartGui() {

    appl := app.New()
    window := appl.NewWindow("Key generator")
    icon, err := assets.Asset("../assets/logo-katena.png")
    if err != nil {
        fmt.Println("icon can not be loaded : ", err.Error())
    }
    window.SetIcon(fyne.NewStaticResource("Icon transchain", icon))
    pubKeyEntry := widget.NewEntry()
    privKeyEntry := widget.NewEntry()
    window.SetContent(widget.NewVBox(
        widget.NewLabel("Public key : "),
        pubKeyEntry,
        widget.NewLabel("Private key : "),
        privKeyEntry,

        widget.NewButton("Generate X25519 keys", func() {
            // Generate X25519 keys
            xKeys, err := libs.GenerateX25519()
            if err != nil {
                _, err := dlgs.Error("Error", "Encountered an issue generating X25519 keys.")
                if err != nil {
                    fmt.Println("error opening generating keys dialog", err.Error())
                }
                return
            }

            pubKeyEntry.SetText(xKeys.PubKey)
            privKeyEntry.SetText(xKeys.PrivKey)
        }),

        widget.NewButton("Generate ED25519 keys", func() {
            // Generate ED25519 keys
            edKeys, err := libs.GenerateEd25519()
            if err != nil {
                _, err := dlgs.Error("Error", "Encountered an issue generating ED25519 keys.")
                if err != nil {
                    fmt.Println("error opening generating keys dialog", err.Error())
                }
                return
            }

            pubKeyEntry.SetText(edKeys.PubKey)
            privKeyEntry.SetText(edKeys.PrivKey)
        }),

        widget.NewButton("Save", func() {
            // Save the keypair to a file
            if pubKeyEntry.Text == "" || privKeyEntry.Text == "" {
                _, err := dlgs.Info("No keys generated", "Please generate a key pair before saving.")
                if err != nil {
                    fmt.Println("error opening no keys dialog", err.Error())
                }
                return
            }

            keys := libs.KeyPair{
                PubKey:  pubKeyEntry.Text,
                PrivKey: privKeyEntry.Text,
            }

            filePath, hasSelected, err := dlgs.File("Choose a file to save your keys", "", false)
            if err != nil {
                _, err := dlgs.Error("Invalid path", "Invalid path selected.")
                if err != nil {
                    fmt.Println("error opening invalid path dialog", err.Error())
                }
                return
            }

            if hasSelected {
                err := keys.Save(filePath)
                if err != nil {
                    _, err := dlgs.Error("Error saving", "Can't save the keys to the selected file.")
                    if err != nil {
                        fmt.Println("error saving to the file", err.Error())
                    }
                    return
                }
                return
            }
        }),

        widget.NewButton("Quit", func() {
            appl.Quit()
        }),
    ))
    window.Resize(fyne.Size{Width: 900, Height: 250})
    window.CenterOnScreen()
    window.ShowAndRun()
}
