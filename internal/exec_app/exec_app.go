package execapp

import (
	"log"
	"os"
	"strings"

	"github.com/dossantoscarlos/genpdf_cli/internal/model"
	"github.com/dossantoscarlos/genpdf_cli/internal/usecase"
	"github.com/dossantoscarlos/genpdf_cli/internal/util"
)

func MergePDF(pdf model.Pdf) {
	pdf.OutputValueName = "merge_pdf" + separatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	inputValueName := strings.Split(pdf.InputValueName, ",")
	nameOutput, err := usecase.MergePDFs(inputValueName, pdf.OutputValueName)
	if err != nil {
		log.Fatalf("Err: %v", err)
		return
	}
	usecase.OpenFileInBrowser(nameOutput)
}

func ExtractPDF(pdf model.Pdf) {
	pdf.OutputValueName = "pdf_extract" + separatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	nameOutput, err := usecase.ExtractPages(pdf)
	if err != nil {
		log.Fatalf("Erro ao extrair páginas: %v", err)
		return
	}
	usecase.OpenFileInBrowser(nameOutput)
}

func CompressPDF(pdf model.Pdf) {
	pdf.OutputValueName = "compress_pdf" + separatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	nameOutput, err := usecase.CompressPDF(pdf.InputValueName, pdf.OutputValueName)
	if err != nil {
		log.Fatalf("Erro ao comprimir o PDF: %v", err)
		return
	}
	usecase.OpenFileInBrowser(nameOutput)
}

func SplitPDF(pdf model.Pdf) {
	pdf.OutputValueName = "split_pdf" + separatorPath()
	util.VerificaDirectory(pdf.OutputValueName)
	_, err := usecase.SplitPDF(pdf)
	if err != nil {
		log.Fatalf("Erro ao dividir PDF: %v", err)
		return
	}
	usecase.OpenFolder(pdf.OutputValueName)
}

func separatorPath() string {
	return string(os.PathSeparator)
}
