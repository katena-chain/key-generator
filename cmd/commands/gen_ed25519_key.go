package commands

import (
    "fmt"

    "github.com/spf13/viper"

    "github.com/katena-chain/key-generator/libs"

    "github.com/spf13/cobra"
)

func GetGenED25519KeysCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "gen-ed",
        Short: "Generate a new ED25519 key pair",
        RunE:  GetGenED25519KeysRunEFn(),
    }

    return cmd
}

func GetGenED25519KeysRunEFn() func(*cobra.Command, []string) error {
    return func(cmd *cobra.Command, args []string) error {
        fmt.Println("generate key ED25519")

        keys, err := libs.GenerateEd25519()
        if err != nil {
            return err
        }

        keys.Show()

        savePath := viper.GetString("save")
        // => Value is "" if no flag has been used
        if savePath != "" {
            err = keys.Save(savePath) // => flag parameter : path to the file
            if err != nil {
                fmt.Println("error while saving keys to a file :", err.Error())
            }
        }

        return nil
    }
}
