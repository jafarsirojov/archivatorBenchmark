package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var (
	con = flag.String("con", "default", "concurrency archive")
	seq = flag.String("seq", "default", "sequence archive")
)

func main() {
	flag.Parse()
	files := os.Args[2:]
	if *seq != "default" {
		seqArchive(files)
		return
	}
	if *con != "default" {
		conArchive(files)
		return
	}
	fmt.Println(`Команды:
по одному:
		archivator -seq <имя файлов ....>
по ктнкурентности
		archivator -con <имя файлов ....>`)
	return
}

func conArchive(files []string) {
	wg:=sync.WaitGroup{}
	for _, file := range files {
		filePath:=fileInputPath+file
		wg.Add(1)
		go func(wg *sync.WaitGroup, file string, filePath string) {
			defer wg.Done()
			archiveZip(file,filePath,fileOutputConPath)

		}(&wg, file, filePath)
	}
	wg.Wait()
}

func seqArchive(files []string) {
	for _, file := range files {
		filePath:=fileInputPath+file
		archiveZip(file,filePath,fileOutputSeqPath)
	}
}

func archiveZip(fileName string,inputFile string, fileOutZipPath string) {
	outputFile := fileOutZipPath + fileName + fileTypeZip
	zipFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("can't creat new zip file: %v\n", err)
	}
	defer func() {
		err := zipFile.Close()
		if err != nil {
			log.Fatalf("can't close zipFile: %v\n", err)
		}
	}()

	writer := zip.NewWriter(zipFile)
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Fatalf("can't close writer: %v\n", err)
		}
	}()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("can't open input file: %v\n", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("can't file close: %v\n", err)
		}
	}()

	info, err := file.Stat()
	if err != nil {
		log.Fatalf("can't stat file: %v\n", err)
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Fatalf("can't FileInfoHeader: %v\n", err)
	}
	header.Name = fileName
	header.Method = zip.Deflate
	writerHeader, err := writer.CreateHeader(header)
	if err != nil {
		log.Fatalf("can't CreatHeader: %v\n", err)
	}
	_, err = io.Copy(writerHeader, file)
	if err != nil {
		log.Fatalf("can't copy file: %v\n", err)
	}
	return
}
