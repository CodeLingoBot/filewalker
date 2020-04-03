package main

import (
	"crypto/sha512"
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



var (
	help = flag.Bool("help", false, "Help")
	initialPath = flag.String("initial-path", "./", "Set the top level directory to begin scan.")
	hashAlg = flag.String("hash", "sha256", "Wat Hash algorithm use. Options: md5, sha1, sha256, sha512")
	reportOnly = flag.Bool("report-only", false, "")
)

var i = 1

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

			hash := sha512.New()
			if _, err := io.Copy(hash, f); err != nil {
				log.Fatal(err)
			}

			info, err := f.Stat()

			fmt.Printf("%v\t%x\t%v\t%v\n", i, hash.Sum(nil), info.Size(), fullyQualifiedFilename)

			i++

			f.Close()
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	if *help {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	initialPath := strings.TrimSpace(*initialPath)

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
