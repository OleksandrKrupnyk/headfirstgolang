package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from scanDirrectory(\"%s\") call\n", path)
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirrectory(\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println("File:", file.Name())
		}
	}
	return nil
}

func main() {
	err := scanDirectory("c:\\users")
	if err != nil {
		log.Fatal(err)
	}
}
