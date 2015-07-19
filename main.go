package main

import (
	"flag"
	"log"
	"path/filepath"
)

func getFlags() *Flags {
	scanDir := flag.String("scandir", ".", "Directory to scan")

	flag.Parse()

	absPath, _ := filepath.Abs(*scanDir)

	flags := &Flags{
		ScanDir: absPath}

	log.Printf("Scan directory: %s", flags.ScanDir)

	return flags
}

func main() {
	log.Println("Starting PHP Virus Scanner")

	flags := getFlags()

	Scan(flags)
}
