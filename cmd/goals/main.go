package main

import (
	"fmt"
	"github.com/marco2704/goals/internal/config"
	"github.com/spf13/cobra"
	"os"
)

var goalsCmd *cobra.Command

func init() {
	goalsCmd = &cobra.Command{
		Use:           "goals GOAL",
		Short:         "Automation tool configured in YAML, customized in Go.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	goalsCmd.PersistentFlags().StringP("file", "f", "goals.yaml", "Name of the YAML goals config file")

	goalsCmd.RunE = func(cmd *cobra.Command, args []string) error {
		goalsFile, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		goalsConfig, err := config.From(goalsFile)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			fmt.Printf("\nAvailable goals:\n")

			for goalName, goal := range goalsConfig.Goals {
				fmt.Printf("%2s%-20s\t%s\n", "", goalName, goal.Description)
			}

			return nil
		}

		return goalsConfig.RunGoals(args...)
	}
}

func main() {
	if err := goalsCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
