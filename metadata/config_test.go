package metadata_test

import (
	"encoding/json"
	"flag"
	"io"
	"os"
	"testing"
)

var (
	caseTest   = flag.String("case", "", "specifies which test case to request")
	update     = flag.Bool("update", false, "update result file")
	htmlcreate = flag.Bool("create", false, "create new html files")
)

func readJSONFile(t *testing.T, fName string, data interface{}) {
	dataByte, err := os.ReadFile(fName)
	if err != nil {
		t.Fatalf("couldnt read file. error [%s]", err)
		return
	}
	err = json.Unmarshal(dataByte, data)
	if err != nil {
		t.Fatalf("couldnt unmarshal data to json. error [%s]", err)
	}
}

func createJSONFile(t *testing.T, fName string, data interface{}, indent ...bool) {
	var dataByte []byte
	var err error
	if len(indent) > 0 && indent[0] {
		dataByte, err = json.MarshalIndent(data, "", "\t")
	} else {
		dataByte, err = json.Marshal(data)
	}

	if err != nil {
		t.Fatalf("couldnt marshal data to json. error [%s]", err)
		return
	}
	err = os.WriteFile(fName, dataByte, 0644)
	if err != nil {
		t.Fatalf("couldnt create file [%s] error [%s]", fName, err)
	}
}

func createFileFromReader(t *testing.T, fName string, data io.Reader) {
	fhandler, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		t.Fatalf("couldnt create file [%s]. error [%s]", fName, err)
		return
	}
	defer func() {
		if err := fhandler.Close(); err != nil {
			t.Fatalf("couldnt close file [%s]. error [%s]", fName, err)
		}
	}()
	_, err = io.Copy(fhandler, data)
	if err != nil {
		t.Fatalf("couldnt copy data into file [%s]. error [%s]", fName, err)
	}
}

func readFile(t *testing.T, inFile string) (io.ReadCloser, func()) {
	html, err := os.Open(inFile)
	if err != nil {
		t.Fatalf("error on openning file. error [%s]", err)
		return nil, func() {}
	}
	return html, func() {
		err := html.Close()
		if err != nil {
			t.Errorf("couldnt close file. error [%s]", err)
		}
	}
}
