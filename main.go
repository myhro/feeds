package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/autossegredos"
	"github.com/myhro/feeds/cmd"
	"github.com/myhro/feeds/copasa"
	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/liquipedia"
)

func main() {
	errormap.Init(
		autossegredos.Command,
		copasa.Command,
		liquipedia.Command,
	)

	rootCmd := &cobra.Command{
		Use:   "feeds",
		Short: "Atom/RSS feed generator for websites that don't offer them",
	}

	autosSegredosCmd := &cobra.Command{
		Use:   autossegredos.Command,
		Short: autossegredos.FeedTitle,
		Run:   cmd.AutosSegredos,
	}

	copasaCmd := &cobra.Command{
		Use:   copasa.Command,
		Short: copasa.FeedTitle,
		Run:   cmd.Copasa,
	}

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate multiple feeds at once",
		Run:   cmd.Generate,
	}
	generateCmd.Flags().StringSliceP("feed", "f", []string{}, "feed to generate, can be specified multiple times")

	err := generateCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("generateCmd.MarkFlagRequired: ", err)
	}

	liquipediaCmd := &cobra.Command{
		Use:   liquipedia.Command,
		Short: liquipedia.FeedTitle,
		Run:   cmd.Liquipedia,
	}

	rootCmd.AddCommand(autosSegredosCmd)
	rootCmd.AddCommand(copasaCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(liquipediaCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err = rootCmd.Execute()
	if err != nil {
		log.Fatal("rootCmd.Execute: ", err)
	}
}
