package usecase

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePDFs(inputFiles []string, outputFile string) (string, error) {
	outputFile = generatedName(outputFile, inputFiles[0])
	err := api.MergeCreateFile(inputFiles, outputFile, false, nil)
	if err != nil {
		return "", fmt.Errorf("Err: %v\n", err)
	}

	return outputFile, nil
}
