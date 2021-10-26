package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/liquipedia"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "feeds",
		Short: "Atom/RSS feed generator for websites that don't offer them",
	}

	liquipediaCmd := &cobra.Command{
		Use:   "liquipedia",
		Short: liquipedia.FeedTitle,
		Run:   liquipedia.Run,
	}

	rootCmd.AddCommand(liquipediaCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("rootCmd.Execute: ", err)
	}
}
