package main

import (
	"github.com/cjdell/php-virus-scanner/definitions"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Flags struct {
	ScanDir string
}

type PhpFile struct {
	Path    string
	ModTime time.Time
}

type Definition interface {
	Init()
	Name() string
	Check(source string) bool
}

func Scan(flags *Flags) {
	defs := []Definition{
		&definitions.EvalBase64{},
		&definitions.EvalEscaped{},
		&definitions.ScriptInject{},
		&definitions.Assert{},
		&definitions.FileWriter{}}

	for _, v := range defs {
		v.Init()
		log.Printf("Registered \"%s\" definition", v.Name())
	}

	fileList := getFileList(flags.ScanDir)

	for _, file := range fileList {
		fullPath := filepath.Join(flags.ScanDir, file.Path)

		data, err := ioutil.ReadFile(fullPath)

		if err != nil {
			log.Printf("Error reading file: %s (%s)", fullPath, err.Error())
			continue
		}

		source := string(data)

		for _, def := range defs {
			found := def.Check(source)

			if found {
				log.Printf("FOUND VIRUS (Matcher=%s) (Path=%s) (ModTime=%s)", def.Name(), file.Path, file.ModTime)
			}
		}
	}
}

func getFileList(dir string) []*PhpFile {
	fileList := []*PhpFile{}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error resolving file: %s (%s)", path, err.Error())
		}

		if !f.IsDir() && filepath.Ext(f.Name()) == ".php" {
			relPath, _ := filepath.Rel(dir, path)

			fileList = append(fileList, &PhpFile{
				Path:    relPath,
				ModTime: f.ModTime()})
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking directory tree: ", err.Error())
	}

	return fileList
}
