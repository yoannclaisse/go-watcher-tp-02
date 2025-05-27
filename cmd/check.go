package cmd

import (
	"errors"
	"fmt"
	"sync"
	"tp2/internal/checker"
	"tp2/internal/config"
	"tp2/internal/reporter"

	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
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
		resultsChan := make(chan checker.CheckResult, len(targets))

		wg.Add(len(targets))
		for _, target := range targets {
			go func(t config.InputTarget) {
				result := checker.CheckUrl(t)
				resultsChan <- result
				defer wg.Done()
			}(target)
		}

		wg.Wait()
		close(resultsChan)

		var finalReport []checker.ReportEntry
		for res := range resultsChan {
			reportEntry := checker.ConvertToReportEntry(res)
			finalReport = append(finalReport, reportEntry)
			if res.Err != nil {
				var unreachable *checker.UnreachableURLError
				if errors.As(res.Err, &unreachable) {
					fmt.Printf("KO %s is unreachable: %v\n", res.InputTarget.URL, unreachable.Err)
				} else {
					fmt.Printf("KO %s: %v\n", res.InputTarget.URL, res.Err)
				}
			} else {
				fmt.Printf("OK %s: %s\n", res.InputTarget.URL, res.Status)
			}
		}
		if outputFilePath != "" {
			err := reporter.ExportResultsToJsonFile(outputFilePath, finalReport)
			if err != nil {
				fmt.Printf("Error exporting results to file %s: %v\n", outputFilePath, err)
			} else {
				fmt.Printf("Results successfully exported to %s\n", outputFilePath)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "chemin vers le fichier JSON d'entrées contenant les URLs")
	checkCmd.MarkFlagRequired("input")
}
