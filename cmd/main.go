package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	execapp "github.com/dossantoscarlos/genpdf_cli/internal/exec_app"
	"github.com/dossantoscarlos/genpdf_cli/internal/model"
)

func main() {
	var err error

	input := flag.String("input", "", "Arquivo PDF de entrada (obrigatório)")
	options := flag.String("option", "", "1:Extract PDF 2:Merge PDF 3:Compress PDF 4:Split PDF")
	pages := flag.String("pages", "", "Páginas a extrair (ex: 1,2,5)")

	flag.Parse()

	if *input == "" || *pages == "" {
		fmt.Println("Uso: go run main.go --input arquivo.pdf --option 1 --pages 1,2,3 or all para todas as páginas")
		flag.PrintDefaults()
		return
	}

	option, err := strconv.Atoi(*options)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	pageRanges := strings.Split(*pages, ",")

	pdf := model.Pdf{
		InputValueName: *input,
		PageRange:      pageRanges,
	}

	switch option {
	case 1:
		execapp.ExtractPDF(pdf)
	case 2:
		execapp.CompressPDF(pdf)
	case 3:
		execapp.MergePDF(pdf)
	case 4:
		execapp.SplitPDF(pdf)
	default:
		log.Panicf("Error ao selecionar diretorio")
	}
}
