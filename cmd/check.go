package cmd

import (
	"fmt"
	"sync"
	"tp2/internal/checker"
	"tp2/internal/config"

	"github.com/spf13/cobra"
)

var (
	inputFilePath string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "verify URLs access",
	Long:  "LOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOONNNNNNNNNNNNNNNNNNNNNNGGGGGGGGGGGGGGGGGGGGGGGGGG",
	Run: func(cmd *cobra.Command, args []string) {

		if inputFilePath == "" {
			fmt.Println("mathieu il sait !!")
			return
		}

		targets, err := config.LoadTargetsFromFile(inputFilePath)
		if err != nil {
			fmt.Println("error loading", err)
			return
		}

		if len(targets) == 0 {
			fmt.Println("no length", err)
			return
		}

		var wg sync.WaitGroup

		wg.Add(len(targets))
		for _, url := range targets {
			go func(u string) {
				result := checker.CheckUrl(u)
				defer wg.Done()
				if result.Err != nil {
					fmt.Printf("❌ KO - URL: %-50s | Erreur : %v\n", result.Target, result.Err)
				} else {
					fmt.Printf("✅ OK - URL: %-50s | Status : %s\n", result.Target, result.Status)
				}
			}(url)
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "chemin vers le fichier JSON d'entrées contenant les URLs")
	checkCmd.MarkFlagRequired("input")
}
