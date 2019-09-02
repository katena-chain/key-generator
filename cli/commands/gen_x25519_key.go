package commands

import (
    "fmt"
    "errors"

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
            return errors.New(fmt.Sprintf("unable to generate the keys: %s", err.Error()))
        }

        keys.Show()

        savePath, _ := cmd.Flags().GetString(SaveFlag)
        // => Value is "" if no flag has been used
        if savePath != "" {
            err = keys.Save(savePath) // => flag parameter : path to the file
            if err != nil {
                return errors.New(fmt.Sprintf("unable to save the keys file: %s", err.Error()))
            }
            fmt.Println(fmt.Sprintf("Keys saved to: %s", savePath))
        }

        return nil
    }
}
