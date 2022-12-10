package csv

import (
	"os"

	"github.com/gocarina/gocsv"
)

type CsvParser struct {
	separator string
}

const defaultSeparator = ","

func NewCsvParser() CsvParser {
	return CsvParser{
		separator: defaultSeparator,
	}
}

func (c *CsvParser) SetSeparator(separator string) {
	c.separator = separator
}

func (c *CsvParser) ParseFile(path string, results interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, results); err != nil {
		return err
	}
	return nil
}
