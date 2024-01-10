package lib

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFiles(srcPath, destDir string) error {
	return filepath.Walk(srcPath, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Determine the destination path
		destPath := filepath.Join(destDir, srcPath)

		// If it's a directory, create it in the destination directory
		if info.IsDir() {
			err := os.Mkdir(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			// If it's a file, copy it to the destination directory
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			// ensure the middle directory is created
			err = os.MkdirAll(filepath.Dir(destPath), 0755)
			if err != nil {
				return err
			}
			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
