package main

import "os"

func CreateFolder(folderPath string) {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		l(err)
	}
}
