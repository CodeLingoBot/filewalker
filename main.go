package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	pathSeparator = string(os.PathSeparator)
)

var (
	help        = flag.Bool("help", false, "Help")
	initialPath = flag.String("initial-path", "./", "Set the top level directory to begin scan.")
	hashAlg     = flag.String("hash", "sha256", "Wat Hash algorithm use. Options: md5, sha1, sha256, sha512")
	// reportOnly  = flag.Bool("report-only", false, "")
)

// var i = 0

func buildFQFN(path string, filename string) string {

	if strings.HasSuffix(path, pathSeparator) {
		path = path[:len(path)-1]
	}

	return strings.Join([]string{path, filename}, pathSeparator)
}

// func listFiles(path string) {

// 	files, err := ioutil.ReadDir(path)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, file := range files {
// 		filename := file.Name()
// 		fileMode := file.Mode()

// 		// get the fully quallified file name
// 		fqfn := buildFQFN(path, filename)

// 		if fileMode&os.ModeDir != 0 {
// 			listFiles(fqfn)
// 		} else if fileMode&os.ModeType == 0 {
// 			// i++

// 			f, err := os.Open(fqfn)

// 			if err == nil {
// 				// -----------------------------------------

// 				hash := sha512.New()

// 				// if _, err := io.Copy(hash, f); err != nil {
// 				// 	log.Println("err1")
// 				// 	log.Fatal(err)
// 				// }

// 				info, err := f.Stat()

// 				if err != nil {
// 					log.Println("err-stats")
// 				}

// 				// fmt.Printf("%v\t%x\t%v\t%v\n", i, hash.Sum(nil), info.Size(), fqfn)

// 				fmt.Printf("%x\t%v\t%v\n", hash.Sum(nil), info.Size(), fqfn)

// 				// -----------------------------------------

// 			} else {
// 				// if os.IsPermission(err) {
// 				// 	log.Println(err)
// 				// } else {
// 				// 	log.Fatal(err)
// 				// }
// 				log.Println(err)
// 			}

// 			f.Close()

// 		}
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func listFiles(ch <-chan string, path string) {

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
		log.Println("err2")
		log.Fatal(err)
	}

	if !fileInfo.IsDir() {
		log.Fatal(fmt.Sprintf("%v is not a directory.", initialPath))
	}

	if !filepath.IsAbs(initialPath) {
		absolutePath, err := filepath.Abs(initialPath)

		if err != nil {
			log.Println("err3")
			log.Fatal(err)
		}

		log.Println(fmt.Sprintf("%v is not absolute, using %v instead.", initialPath, absolutePath))

		initialPath = absolutePath
	}

	filesChan := make(chan string, 2)

	go func() { filesChan <- "ping" }()

	msg := <-filesChan
	fmt.Println(msg)

	listFiles(initialPath, filesChan)

}
