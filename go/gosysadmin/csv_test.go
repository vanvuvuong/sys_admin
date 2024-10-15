package gosysadmin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvDecode(t *testing.T) {
	t.Parallel()
	csvData := DecodeCsv("testdata/example.csv", ';')
	expectedFistRow := []string{"1", "programming language", "golang"}
	assert.Equal(t, expectedFistRow, csvData[0], "Matching first row")
	expectedLastRow := []string{"8", "container", "k0s"}
	assert.Equal(t, expectedLastRow, csvData[len(csvData)-1], "Matching first row")
}

// func TestCsvEncode(t *testing.T) {
// 	t.Parallel()
// 	csvData := DecodeCsv("testdata/example.csv", ';')
// 	expectedFistRow := []string{"1", "programming language", "golang"}
// 	assert.Equal(t, expectedFistRow, csvData[0], "Matching first row")
// 	expectedLastRow := []string{"8", "container", "k0s"}
// 	assert.Equal(t, expectedLastRow, csvData[len(csvData)-1], "Matching first row")
// }
