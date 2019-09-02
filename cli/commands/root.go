package commands

import (
    "github.com/spf13/cobra"
)

const (
    SaveFlag      = "save"
    SaveShorthand = "s"
)

func GetRootCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "key-generator",
        Short: "Katena Key Generator command line interface",
    }

    cmd.PersistentFlags().StringP(SaveFlag, SaveShorthand, "", "Save keys into the specified file")

    return cmd
}
