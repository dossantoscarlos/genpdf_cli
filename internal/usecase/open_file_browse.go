package usecase

import "os/exec"

func OpenFileInBrowser(filePath string) error {
	// Comando para abrir o arquivo no navegador padrão
	cmd := exec.Command("cmd", "/c", "start", filePath) // Para Windows
	return cmd.Start()
}
