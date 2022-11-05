package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/cmd"
	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/liquipedia"
)

func main() {
	errormap.Init(
		liquipedia.Command,
	)

	rootCmd := &cobra.Command{
		Use:   "feeds",
		Short: "Atom/RSS feed generator for websites that don't offer them",
	}

	liquipediaCmd := &cobra.Command{
		Use:   liquipedia.Command,
		Short: liquipedia.FeedTitle,
		Run:   cmd.Liquipedia,
	}

	rootCmd.AddCommand(liquipediaCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("rootCmd.Execute: ", err)
	}
}
