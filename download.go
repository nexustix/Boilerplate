package boilerplate

//package nxdownload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	//bp "github.com/nexustix/boilerplate"
)

//Download represents a downloadable file
type Download struct {
	Filename string
	URL      string
}

//DownloadRemoteFile downloads a remote file via a remote adress
func DownloadRemoteFile(destination, remoteAdress string) {
	//remoteAdress := "http://www.bay12games.com/dwarves/df_40_24_win.zip"
	//filename := "filename"

	//FIXME could fail if remote filesize is unknown (needs testing)

	var sizeNow int64
	var sizeLast int64
	var remoteSize int64
	var percent int64

	//pool := x509.NewCertPool()
	tr := &http.Transport{
		//TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	remoteData, err := client.Get(remoteAdress)

	/*
		remoteData, err := http.Get(remoteAdress)
	*/
	if GotError(err) {
		panic("<!> Error getting remote data")
	}
	defer remoteData.Body.Close()
	remoteSize = remoteData.ContentLength
	//fmt.Printf("info;remoteSize: >%d<\n", remoteSize)

	//fileExists, err := bp.FileExists("filename")
	var exists bool

	tmpFile, err := os.Open(destination)
	if GotError(err) {
		//fmt.Printf("info;not opening\n")
		exists = false
	}
	tmpStat, err := tmpFile.Stat()
	if GotError(err) {
		//fmt.Printf("info;no stats\n")
		exists = false
	} else {
		//size := tmpStat.Size()
		//fmt.Printf("info;%ds\n", size)
		exists = true
	}

	var maybeEqual bool

	if (exists && tmpStat.Size() >= 0) && (tmpStat.Size() != remoteSize) {
		maybeEqual = false
		//fmt.Println("info;Size different")
	} else {
		//fmt.Println("info;Size equal")
		maybeEqual = true
	}
	tmpFile.Close()

	if remoteSize > 0 && (!maybeEqual || !exists) {
		outFile, err := os.Create(destination)
		if GotError(err) {
			panic("<!> Error creating local file")
		}
		go downloadFile(outFile, remoteData)

		for {
			fileStatus, err := outFile.Stat()
			if GotError(err) {
				panic("<!> Error getting local file info")
			}
			sizeNow = fileStatus.Size()
			if sizeLast != sizeNow {
				percent = int64((float64(sizeNow) / float64(remoteSize)) * 100)
				//fmt.Printf("\r<-> downloading >%s<[%d/%d] (%d%%)", filename, sizeNow, remoteSize, percent)
				//fmt.Printf("downloading;%s;%d;%d;%d\n", filename, sizeNow, remoteSize, percent)

				fmt.Printf("\r<-> downloading: %d%s", percent, "%")
			} else if sizeLast == remoteSize {

				break
			}
			sizeLast = sizeNow
		}
	}
	fmt.Printf("\n")
}

func downloadFile(outFile *os.File, remoteData *http.Response) {

	_, err := io.Copy(outFile, remoteData.Body)
	FailError(err)
	//fmt.Printf("<->Done >%d< Bytes Copied\n", n)
}
