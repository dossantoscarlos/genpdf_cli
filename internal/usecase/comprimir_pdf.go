package usecase

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func CompressPDF(inputPath string, outputFile string) (string, error) {

	nameOutput := generatedName(outputFile, inputPath)

	err := api.OptimizeFile(inputPath, nameOutput, nil)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	return nameOutput, nil
}
