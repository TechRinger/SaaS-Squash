package sheets

import (
	"io"
	"os"

	"google.golang.org/api/drive/v3"
)

func downloadFile(clientDrive *drive.Service, fileId string, downloadPath string) error {

	resp, err2 := clientDrive.Files.Get(fileId).Download()
	if err2 != nil {
		return err2
	}

	defer resp.Body.Close()

	fileDownloaded, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		return err3
	}

	f, err4 := os.Create(downloadPath)
	if err4 != nil {
		return err4
	}
	_, err5 := f.Write(fileDownloaded)
	if err5 != nil {
		return err5
	}

	return nil
}
