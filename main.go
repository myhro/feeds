package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/errormap"
)

func main() {
	errormap.Init()

	rootCmd := &cobra.Command{
		Use:   "feeds",
		Short: "Atom/RSS feed generator for websites that don't offer them",
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("rootCmd.Execute: ", err)
	}
}
