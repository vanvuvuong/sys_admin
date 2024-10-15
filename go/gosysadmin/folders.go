package gosysadmin

import (
	"os"

	"github.com/gosysadmin/basics"
)

func CreateFolder(folderPath string) {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		basics.L(err)
	}
}
