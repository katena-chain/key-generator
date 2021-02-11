package main

import (
    "github.com/katena-chain/key-generator/cli/commands"
    "github.com/katena-chain/key-generator/version"

    tcCliCommands "github.com/transchain/sdk-go/cli/commands"
)

func main() {
    rootCmd := commands.GetRootCmd()
    rootCmd.AddCommand(
        commands.GetGenED25519KeysCmd(),
        commands.GetGenX25519KeysCmd(),
        tcCliCommands.GetVersionCmd(version.Version, version.GitCommit),
    )

    if err := rootCmd.Execute(); err != nil {
        panic(err)
    }
}
