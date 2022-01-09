package cmd

import (
	"github.com/spf13/cobra"

	"github.com/myhro/feeds/generator"
	"github.com/myhro/feeds/oldnewthing"
)

func OldNewThing(cmd *cobra.Command, args []string) {
	generator.Print(oldnewthing.Command, oldnewthing.XML)
}
