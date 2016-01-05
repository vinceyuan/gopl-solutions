package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Please input a zip/tar file")
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[1], ".zip") {
		r, err := zip.OpenReader(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		for _, f := range r.File {
			fmt.Println(f.Name)
		}
	} else if strings.HasSuffix(os.Args[1], ".tar") {
		r, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		tr := tar.NewReader(r)
		// Iterate through the files in the archive.
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				// end of tar archive
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(hdr.Name)
		}
	} else {
		fmt.Println("not supported file format")
	}

}
