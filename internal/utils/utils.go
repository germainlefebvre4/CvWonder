package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func ReadDirectory(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	CheckError(err)

	result := make([]string, 0)
	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func CopyDirectory(scrDir string, dest string) error {
	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return err
	}
	// Cleanup the .git/ dir
	if ok, _ := regexp.MatchString(".git", scrDir); ok {
		return nil
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0750); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			// Exclude the templated index.html theme file and path with .git/ dir from copying
			if destPath == filepath.Join(dest, "index.html") {
				continue
			}
			if ok, _ := regexp.MatchString(".git", destPath); ok {
				continue
			}
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		fInfo, err := entry.Info()
		if err != nil {
			return err
		}

		isSymlink := fInfo.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, fInfo.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
