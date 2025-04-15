package usecase

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dossantoscarlos/genpdf_cli/internal/model"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func ExtractPages(params model.Pdf) (string, error) {

	filesToMerge, err := SplitPDF(params)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	// Ordenar os arquivos, se necessário
	sort.Strings(filesToMerge)

	nameOutput := generatedName(params.OutputValueName, params.InputValueName)

	err = api.MergeCreateFile(filesToMerge, nameOutput, false, nil)
	if err != nil {
		log.Fatalf("Erro ao fazer merge dos PDFs: %v", err)
		return "", fmt.Errorf("%v", err)
	}

	// 3. Apagar arquivos temporários
	for _, f := range filesToMerge {
		os.Remove(f)
	}

	fmt.Println("PDF: Paginas extraida com sucesso!!!")

	return nameOutput, nil
}
