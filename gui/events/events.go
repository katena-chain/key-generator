package events

import (
    "fmt"

    "fyne.io/fyne"
    "fyne.io/fyne/widget"
    "github.com/gen2brain/dlgs"

    "github.com/katena-chain/key-generator/libs"
)

func OnClickGenX25519Keys(pubKeyEntry *widget.Entry, privKeyEntry *widget.Entry) func() {
    return func() {
        // Generate X25519 keys
        keys, err := libs.GenerateX25519()
        HandleGenError(err, keys, pubKeyEntry, privKeyEntry)
    }
}

func OnClickGenED25519Keys(pubKeyEntry *widget.Entry, privKeyEntry *widget.Entry) func() {
    return func() {
        // Generate X25519 keys
        keys, err := libs.GenerateEd25519()
        HandleGenError(err, keys, pubKeyEntry, privKeyEntry)
    }
}

func OnClickSaveKeys(pubKeyEntry *widget.Entry, privKeyEntry *widget.Entry) func() {
    return func() {
        // Save the keypair to a file
        if pubKeyEntry.Text == "" || privKeyEntry.Text == "" {
            _, uiErr := dlgs.Error("Error", "please generate a key pair before saving")
            if uiErr != nil {
                fmt.Println("please generate a key pair before saving")
            }
            return
        }

        keys := libs.KeyPair{
            PubKey:  pubKeyEntry.Text,
            PrivKey: privKeyEntry.Text,
        }

        filePath, ok, err := dlgs.Entry("Keys file path", "Enter the file path to save your keys", "")
        if err != nil {
            _, uiErr := dlgs.Error("Error", fmt.Sprintf("unable to retrieve the keys file path: %s", err.Error()))
            if uiErr != nil {
                fmt.Println("unable to retrieve the keys file path", err.Error())
            }
            return
        }

        if ok {
            err := keys.Save(filePath)
            if err != nil {
                _, uiErr := dlgs.Error("Error", fmt.Sprintf("unable to save the keys file: %s", err.Error()))
                if uiErr != nil {
                    fmt.Println(fmt.Sprintf("unable to save the keys file: %s", err.Error()))
                }
                return
            }
            _, uiErr := dlgs.Info("Done", fmt.Sprintf("Keys saved to: %s", filePath))
            if uiErr != nil {
                fmt.Println(fmt.Sprintf("Keys saved to: %s", filePath))
            }
            return
        }
    }
}

func OnClickQuit(application fyne.App) func() {
    return func() {
        application.Quit()
    }
}

func HandleGenError(err error, keys libs.KeyPair, pubKeyEntry *widget.Entry, privKeyEntry *widget.Entry) {
    if err != nil {
        _, uiErr := dlgs.Error("Error", fmt.Sprintf("unable to generate the keys: %s", err.Error()))
        if uiErr != nil {
            fmt.Println(fmt.Sprintf("unable to generate the keys: %s", err.Error()))
        }
        return
    }
    pubKeyEntry.SetText(keys.PubKey)
    privKeyEntry.SetText(keys.PrivKey)
}
