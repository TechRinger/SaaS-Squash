package sheets

import (
	"errors"
	"fmt"
	"io"
	"os"

	"google.golang.org/api/drive/v3"
)

func downloadFile(clientDrive *drive.Service, fileName string, downloadPath string) error {

	// Step 6: Search for the file by name in Google Drive
	files, err := clientDrive.Files.List().Q(fmt.Sprintf("name='%s'", fileName)).Do()
	if err != nil {
		return errors.New("unable to retrieve file list")
	}

	if len(files.Files) == 0 {
		return errors.New("no files found with the specified name: " + fileName + ".")
	}

	fileID := files.Files[0].Id
	resp, err2 := clientDrive.Files.Get(fileID).Download()

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
