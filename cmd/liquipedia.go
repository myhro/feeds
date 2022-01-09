package cmd

import (
	"github.com/spf13/cobra"

	"github.com/myhro/feeds/generator"
	"github.com/myhro/feeds/liquipedia"
)

func Liquipedia(cmd *cobra.Command, args []string) {
	generator.Print(liquipedia.Command, liquipedia.XML)
}
