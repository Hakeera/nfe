// Package load 
package load 

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// LoadCSV carrega dados de um CSV para uma slice de structs.
// O struct de destino deve usar tags `csv:"nome_coluna"` para mapeamento.
func LoadCSV[T any](path string) ([]T, error) {
	var result []T

	// Abre o arquivo
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	// Cria um leitor CSV
	reader := csv.NewReader(file)

	// Lê o cabeçalho
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler cabeçalho: %w", err)
	}

	// Mapeia colunas para campos da struct via tags `csv`
	for {
		record, err := reader.Read()
		if err != nil {
			break // EOF ou erro
		}

		var item T
		itemVal := reflect.ValueOf(&item).Elem()
		itemType := itemVal.Type()

		for i, header := range headers {
			for j := 0; j < itemVal.NumField(); j++ {
				field := itemType.Field(j)
				tag := field.Tag.Get("csv")
				if strings.EqualFold(tag, header) {
					fieldVal := itemVal.Field(j)
					if fieldVal.CanSet() && fieldVal.Kind() == reflect.String {
						fieldVal.SetString(record[i])
					}
				}
			}
		}

		result = append(result, item)
	}

	return result, nil
}
