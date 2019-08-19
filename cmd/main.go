package main

import (
    "github.com/katena-chain/key-generator/cmd/commands"
)

func main() {
    rootCmd := commands.GetRootCmd()
    rootCmd.AddCommand(
        commands.GetGenED25519KeysCmd(),
        commands.GetVersionCmd(),
        commands.GetGenX25519KeysCmd(),
    )

    if err := rootCmd.Execute(); err != nil {
        panic(err)
    }
}
