package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Printf("Usage : cat ugly.xml | xmlpretty > unfugly.xml\n")
}

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		usage()
		os.Exit(1)
	}

	bytes, _ := ioutil.ReadAll(os.Stdin)
	formatted, err := formatXML(bytes)

	if err != nil {
		fmt.Printf("Unable to format data. %v", err)
		os.Exit(1)
	}

	fmt.Println(string(formatted))
}
