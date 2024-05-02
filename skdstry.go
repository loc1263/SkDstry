// seek destroy (sk_dstry)
// Busca archivos en un arbol de directorios, dando la opcion de eliminarlos.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func buscarArchivo(directorioRaiz string, nombreArchivo string) ([]string, int64) {
	var archivosEncontrados []string
	var pesoTotal int64

	fmt.Printf("Buscando..... \n")

	_, err := os.Stat(directorioRaiz)
	if err != nil {
		fmt.Printf("Error: El directorio raíz '%s' no existe\n", directorioRaiz)
		return archivosEncontrados, pesoTotal
	}

	err = filepath.Walk(directorioRaiz, func(path string, info os.FileInfo, err error) error {

		//if err != nil {
		//	fmt.Printf("Error al buscar en el directorio %s: %v\n", directorioRaiz, err)
		//	return err
		//}

		if matchesPattern(nombreArchivo, info.Name()) {
			archivosEncontrados = append(archivosEncontrados, path)
			pesoTotal += info.Size()
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error al buscar en el directorio %s: %v\n", directorioRaiz, err)
	}

	return archivosEncontrados, pesoTotal
}

func matchesPattern(pattern, name string) bool {
	return strings.Contains(name, pattern) || strings.Contains(pattern, "*") && matchWildcard(pattern, name)
}

func matchWildcard(pattern, name string) bool {

	regexPattern := strings.ReplaceAll(pattern, "*", ".*")
	match, err := regexp.MatchString(regexPattern, name)

	if err != nil {
		fmt.Printf("Error al comparar patrón con wildcard: %v\n", err)
		return false
	}
	return match
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Uso: programa directorio_raiz nombre_archivo")
		return
	}

	directorioRaiz := os.Args[1]
	nombreArchivo := os.Args[2]

	archivosEncontrados, pesoTotal := buscarArchivo(directorioRaiz, nombreArchivo)
	fmt.Println("Archivos encontrados:")
	for _, archivo := range archivosEncontrados {
		fmt.Println(archivo)
	}

	pesoKB := float64(pesoTotal) / 1024
	pesoMB := float64(pesoTotal) / (1024 * 1024)
	fmt.Printf("Peso total en KB: %.2f KB\n", pesoKB)
	fmt.Printf("Peso total en MB: %.2f MB\n", pesoMB)

	fmt.Print("¿Desea borrar los archivos encontrados? (Y/N): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	respuesta := strings.ToUpper(scanner.Text())
	if respuesta == "Y" {
		for _, archivo := range archivosEncontrados {
			err := os.Remove(archivo)
			if err != nil {
				fmt.Printf("Error al borrar el archivo %s: %v\n", archivo, err)
			} else {
				fmt.Printf("Archivo %s borrado exitosamente\n", archivo)
			}
		}
	} else {
		fmt.Println("No se han borrado los archivos.")
	}
}
