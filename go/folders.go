package main

import (
	"os"

	"github.com/go/basics"
)

func CreateFolder(folderPath string) {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		basics.L(err)
	}
}
