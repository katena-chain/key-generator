package commands

import (
    "fmt"

    "github.com/spf13/viper"

    "github.com/katena-chain/key-generator/libs"

    "github.com/spf13/cobra"
)

func GetGenX25519KeysCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "gen-x",
        Short: "Generate a new X25519 key pair",
        RunE:  GetGenX25519KeysRunEFn(),
    }

    return cmd
}

func GetGenX25519KeysRunEFn() func(*cobra.Command, []string) error {
    return func(cmd *cobra.Command, args []string) error {
        fmt.Println("generate key X25519")

        keys, err := libs.GenerateX25519()
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
