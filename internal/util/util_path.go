package util

import (
	"fmt"
	"log"
	"os"
)

func VerificaDirectory(dir string) {
	_, err := directory(dir)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}

func directory(pathName string) (string, error) {
	if directoryExists(pathName) {
		os.RemoveAll(pathName)
	}
	err := os.Mkdir(pathName, 0700)
	if err != nil {
		return "", fmt.Errorf("%c", err)
	}
	return pathName, nil
}

func directoryExists(path string) bool {
	dir, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false // O diretório não existe
	}
	return dir.IsDir()
}
