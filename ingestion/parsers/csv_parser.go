package parsers

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func LoadCSV[T any](filename string) ([]*T, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read entire file and replace "N/A" with ""
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "N/A", "")
		line = strings.ReplaceAll(line, "NULL", "")
		_, _ = w.WriteString(line + "\n")
	}
	_ = w.Flush()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var records []*T
	if err := gocsv.Unmarshal(bytes.NewReader(buf.Bytes()), &records); err != nil {
		return nil, err
	}
	return records, nil
}
