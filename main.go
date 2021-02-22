package main

import (
	"io"
	"os"

	"log"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func getSource() (io.Reader, error) {
	if len(os.Args) > 1 && os.Args[1] != "-" {
		return os.Open(os.Args[1])
	}
	return os.Stdin, nil
}

func main() {
	if err := doWork(); err != nil {
		log.Fatalf("json2yaml failed: %s", err.Error())
		os.Exit(1)
	}
}

func doWork() error {
	r, err := getSource()
	if err != nil {
		return err
	}

	dec := jsoniter.NewDecoder(r)
	for dec.More() {
		data := map[string]interface{}{}
		if err := dec.Decode(&data); err != nil {
			return err
		}
		enc := yaml.NewEncoder(os.Stdout)
		os.Stdout.WriteString("---\n")
		if err := enc.Encode(data); err != nil {
			return err
		}
	}
	return nil
}
