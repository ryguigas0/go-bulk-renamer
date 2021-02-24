package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	folder := os.Args[1]
	oldName := os.Args[2]
	newName := ""

	if folder == "info" {
		fmt.Println("./gorenamer [folder of the files] [pattern to find] [replace text]")
		fmt.Println("Rename all the files found on the folder, replacing the found pattern with the replace text provided")
	}

	if len(os.Args[1:]) < 2 {
		fmt.Printf("ERROR: Please fill all the arguments")
		fmt.Println("use help to get info")
		return
	}

	if len(os.Args[1:]) == 3 {
		newName = os.Args[3]
	}

	fileInfos, err := ioutil.ReadDir(folder)

	if err != nil {
		fmt.Printf("Error finding folder: %v\n", err)
	}

	for _, v := range fileInfos {
		if strings.Contains(v.Name(), oldName) {
			oldFilename := filepath.Join(folder, v.Name())
			newFilename := filepath.Join(folder, newPattern(v.Name(), oldName, newName))

			fmt.Printf("RENAMING FILE %v TO %v\n", oldFilename, newFilename)

			err := os.Rename(oldFilename, newFilename)

			if err != nil {
				fmt.Printf("Error renaming file: %v\n", err)
			} else {
				fmt.Println("SUCCESSFULY RENAMED FILE")
			}
		}
	}
}

func newPattern(toChange, toFind, replaceWith string) string {
	result := strings.Replace(toChange, toFind, replaceWith, -1)
	return result
}
