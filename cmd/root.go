package cmd

import (
	"fmt"
	"os"

	"github.com/1995parham/gosimac/cmd/bing"
	"github.com/1995parham/gosimac/cmd/common"
	"github.com/1995parham/gosimac/cmd/unsplash"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := &cobra.Command{
		Use:     "GoSiMac",
		Short:   "Fetch the wallpaper from Bings, Wikimedia ...",
		Version: "4.0.0",
	}

	unsplash.Register(root)
	bing.Register(root)

	root.PersistentFlags().StringP(common.FlagPath, "p", common.DefaultPath(), "A path to store the photos")
	root.PersistentFlags().IntP(common.FlagCount, "n", common.DefaultCount, "The number of photos to return")

	if err := root.Execute(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(ExitFailure)
	}
}
