package utils

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const UTF8BOM string = "\xEF\xBB\xBF"

func CreateCSVFile(filePath string) (f *os.File, err error) {
	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return
	}
	return os.Create(filePath)
}

func WriteCSVFile(f *os.File, data *[][]string) error {
	f.WriteString(UTF8BOM)
	w := csv.NewWriter(f)
	for _, v := range *data {
		err := w.Write(v)
		if err != nil {
			return err
		}
	}
	w.Flush()

	defer f.Close()

	return nil
}

func DataToMap(data *[][]string) (records *[]map[string]string, err error) {
	header := []string{}
	records = &[]map[string]string{}
	for i, record := range *data {
		if i == 0 {
			for j := 0; j < len(record); j++ {
				header = append(header, strings.TrimSpace(record[j]))
			}
		} else {
			l := map[string]string{}
			for j := 0; j < len(record); j++ {
				l[header[j]] = strings.TrimSpace(record[j])
			}
			*records = append(*records, l)
		}
	}
	return
}

func BOMAwareCSVReader(reader io.Reader) *csv.Reader {
	var transformer = unicode.BOMOverride(encoding.Nop.NewDecoder())
	return csv.NewReader(transform.NewReader(reader, transformer))
}
