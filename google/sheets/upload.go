package sheets

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
)

func uploadFile(clientDrive *drive.Service, localFilePath string, driveFolderId string) error {

	var parent []string
	parent = append(parent, driveFolderId)

	fileName := filepath.Base(localFilePath)

	f := &drive.File{
		Name:    fileName,
		DriveId: driveFolderId,
		Parents: parent,
	}

	file, err_os := os.Open(localFilePath)
	if err_os != nil {
		return fmt.Errorf("error opening %v", localFilePath)
	}

	_, err := clientDrive.Files.Create(f).Media(file).Do()

	return err

}
