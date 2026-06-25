package read_packages

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadPackages(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Ошиюка при чтении пакетов которые нужно установить: %w", err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var packages []string
	for scanner.Scan() {
		packages = append(packages, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении пакетов: %w", err)
	}
	res := strings.Join(packages, " ")
	return res
}
