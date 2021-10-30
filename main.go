package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/autossegredos"
	"github.com/myhro/feeds/copasa"
	"github.com/myhro/feeds/liquipedia"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "feeds",
		Short: "Atom/RSS feed generator for websites that don't offer them",
	}

	autosSegredosCmd := &cobra.Command{
		Use:   "autossegredos",
		Short: autossegredos.FeedTitle,
		Run:   autossegredos.Run,
	}

	copasaCmd := &cobra.Command{
		Use:   "copasa",
		Short: copasa.FeedTitle,
		Run:   copasa.Run,
	}

	liquipediaCmd := &cobra.Command{
		Use:   "liquipedia",
		Short: liquipedia.FeedTitle,
		Run:   liquipedia.Run,
	}

	rootCmd.AddCommand(autosSegredosCmd)
	rootCmd.AddCommand(copasaCmd)
	rootCmd.AddCommand(liquipediaCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("rootCmd.Execute: ", err)
	}
}
