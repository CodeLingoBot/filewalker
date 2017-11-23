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
	//_ "github.com/mattn/go-sqlite3"
)

var initialPath = flag.String("initial-path", "", "Initial path")

var i = 0

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

			fmt.Printf("%v\t%x %v\n", i, hash.Sum(nil), fullyQualifiedFilename)
			//fmt.Println("%x, %s", hash.Sum(nil), )

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

	//args := os.Args

	fmt.Println(initialPath)

	//db, err := sql.Open("sqlite3", "./foo.db")
	//checkErr(err)

	//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	//checkErr(err)

	//baseDir := ""
	//listFiles(baseDir)
	fmt.Println("Done. ", i, " files found")
}
