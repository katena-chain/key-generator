package commands

import (
    "github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "key-generator",
        Short: "Katena Key generator command line interface",
    }

    // Flag attached to every command - if some commands get added in the future, consider removing it and adding it to the individual commands that need it
    cmd.PersistentFlags().StringP("save", "s", "", "Save keys into the specified file")

    return cmd
}
