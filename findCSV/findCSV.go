package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	const root = "task.zip"
	fileZip, err := zip.OpenReader(root)
	if err != nil {
		log.Fatal(err)
	}
	defer fileZip.Close()
}
