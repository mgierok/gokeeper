// +build ignore

package main

import (
	"os"
	"path/filepath"
	"regexp"

	bindata "github.com/tmthrgd/go-bindata"
)

func createBindata(output, input string, genOpts *bindata.GenerateOptions, ignore []*regexp.Regexp) error {
	if err := os.MkdirAll(filepath.Dir(output), 0744); err != nil {
		return err
	}

	files, err := bindata.FindFiles(input, &bindata.FindFilesOptions{
		Prefix:    input,
		Recursive: true,
		Ignore:    ignore,
	})
	if err != nil {
		return err
	}

	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	return files.Generate(f, genOpts)
}

func main() {
	if err := createBindata("assets/views/bindata.go", "assets/views", &bindata.GenerateOptions{
		Package:  "views",
		Metadata: true,
		Mode:     0444,
	}, nil); err != nil {
		panic(err)
	}
}
