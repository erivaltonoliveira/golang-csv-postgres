package util

import (
	"bufio"
	"os"
)

// lê o conteudo do arquivo e retorna um string com todas as linhas do arquivo
func ReadFile(filePath string) ([]string, error) {

	// Abre o arquivo
	file, err := os.Open(filePath)
	
	// retorna o erro encontrado ao tentar abrir o arquivo 
	if err != nil {
		return nil, err
	}
	
	// fecha o arquivo após uso
	defer file.Close()

	// lê o arquivo linha a linha
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Retorna as linhas lidas e um erro se ocorrer
	return lines, scanner.Err()
}
