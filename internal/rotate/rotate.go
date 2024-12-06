package rotate

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type FileData struct {
	Header *tar.Header
	Reader io.Reader
}

func ProcessFile(fileName, storageDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("[Error] Cannot open file: %v\n", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("[Error] Mate I cannot take file stat in processing file: %v\n", err)
		return
	}

	var archiveStorageDir string

	if storageDir != "" {
		archiveStorageDir = storageDir

	} else {
		archiveStorageDir = filepath.Dir(fileName)
	}

	if err := os.MkdirAll(archiveStorageDir, os.ModePerm); err != nil {
		fmt.Printf("[Error] Cannot create common dir: %v\n", err)
		return
	}

	archiveName := filepath.Join(archiveStorageDir, fileInfo.Name()+".tar.gz")
	archiveFile, err := os.Create(archiveName)
	if err != nil {
		fmt.Printf("[Error] Cannot create a archive file: %v\n", err)
		return
	}

	defer archiveFile.Close()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	header := &tar.Header{
		Name: fileName,
		Size: fileInfo.Size(),
		Mode: int64(fileInfo.Mode()),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		fmt.Printf("[Error] Error to write header in archive %s: %v\n", archiveName, err)
		return
	}
	if _, err := io.Copy(tarWriter, file); err != nil {
		fmt.Printf("[Error] Error to recording file content in to archive %s: %v\n", archiveName, err)
		return
	}
}
