package rotate

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	archiveExtention = ".tar.gz"
	timestampFormat  = "20060102150405"
)

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

	fileNameWithoutExt := strings.TrimSuffix(fileInfo.Name(), filepath.Ext(fileInfo.Name()))
	timestamp := time.Now().Format(timestampFormat)
	formatedFileName := fileNameWithoutExt + timestamp + archiveExtention

	archiveName := filepath.Join(archiveStorageDir, formatedFileName)
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
