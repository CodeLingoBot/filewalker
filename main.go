package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"path/filepath"
	//_ "github.com/mattn/go-sqlite3"
)

var i = 1

var initialPath string = ""
var reportOnly bool = false

func listFiles(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		fullyQualifiedFilename := strings.Join([]string{path, filename}, "/")

		if file.IsDir() {
			listFiles(fullyQualifiedFilename)
		} else {

			f, err := os.Open(fullyQualifiedFilename)

			if err != nil {
				log.Fatal(err)
			}

			hash := sha256.New()
			if _, err := io.Copy(hash, f); err != nil {
				log.Fatal(err)
			}

			info, err := f.Stat()

			fmt.Printf("%v\t%x\t%v\t%v\n", i, hash.Sum(nil), info.Size(), fullyQualifiedFilename)

			i++
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var help bool = false
	flag.BoolVar(&help, "help", help, "Help")
	flag.StringVar(&initialPath, "initial-path", "./", "Set the top level directory to begin scan.")
	flag.BoolVar(&reportOnly, "report-only", false, "")

	flag.Parse()

	if help {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	initialPath := strings.TrimSpace(initialPath)

	fileInfo, err := os.Lstat(initialPath)

	if err != nil {
		log.Fatal(err)
	}

	if !fileInfo.IsDir() {
		log.Fatal(fmt.Sprintf("%v is not a directory.", initialPath))
	}

	if !filepath.IsAbs(initialPath) {
		absolutePath, err := filepath.Abs(initialPath)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(fmt.Sprintf("%v is not absolute, using %v instead.", initialPath, absolutePath))

		initialPath = absolutePath
	}

	//db, err := sql.Open("sqlite3", "./foo.db")
	//checkErr(err)

	//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	//checkErr(err)

	//baseDir := ""
	listFiles(initialPath)
	fmt.Println("Done. ", i, " files found")
}
