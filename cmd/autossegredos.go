package cmd

import (
	"github.com/spf13/cobra"

	"github.com/myhro/feeds/autossegredos"
	"github.com/myhro/feeds/generator"
)

func AutosSegredos(cmd *cobra.Command, args []string) {
	generator.Print(autossegredos.Command, autossegredos.XML)
}
