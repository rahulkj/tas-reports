package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetRaw(file string) []byte {

	raw, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}

func WriteReport(reportsDir string, fileName string, data []byte) {

	os.MkdirAll(reportsDir, os.ModePerm)

	filePath := filepath.Join(reportsDir, fileName)

	if err := os.WriteFile(filePath, []byte(data), 0644); err != nil {
		fmt.Printf("Failed to write file '%s': %s\n", fileName, err)
		return
	}
	fmt.Printf("File '%s' created successfully.\n", fileName)
}
