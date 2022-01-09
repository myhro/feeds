package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/autossegredos"
	"github.com/myhro/feeds/copasa"
	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/liquipedia"
	"github.com/myhro/feeds/oldnewthing"
)

func main() {
	errormap.Init(
		autossegredos.Command,
		copasa.Command,
		liquipedia.Command,
		oldnewthing.Command,
	)

	rootCmd := &cobra.Command{
		Use:   "feeds",
		Short: "Atom/RSS feed generator for websites that don't offer them",
	}

	autosSegredosCmd := &cobra.Command{
		Use:   autossegredos.Command,
		Short: autossegredos.FeedTitle,
		Run:   autossegredos.Run,
	}

	copasaCmd := &cobra.Command{
		Use:   copasa.Command,
		Short: copasa.FeedTitle,
		Run:   copasa.Run,
	}

	liquipediaCmd := &cobra.Command{
		Use:   liquipedia.Command,
		Short: liquipedia.FeedTitle,
		Run:   liquipedia.Run,
	}

	oldnewthingCmd := &cobra.Command{
		Use:   oldnewthing.Command,
		Short: oldnewthing.FeedTitle,
		Run:   oldnewthing.Run,
	}

	rootCmd.AddCommand(autosSegredosCmd)
	rootCmd.AddCommand(copasaCmd)
	rootCmd.AddCommand(liquipediaCmd)
	rootCmd.AddCommand(oldnewthingCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("rootCmd.Execute: ", err)
	}
}
