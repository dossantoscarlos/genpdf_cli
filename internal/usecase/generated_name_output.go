package usecase

import "time"

func generatedName(outputFile string, inputPath string) string {
	dataHora := time.Now()
	nameOutput := outputFile + dataHora.Format("2006_01_02-15_04_05-") + inputPath
	return nameOutput
}
