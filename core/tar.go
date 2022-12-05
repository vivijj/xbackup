// Provides a Tar tool with methods to create tar archives
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateTarFile creates a tar file from the given directory
// It returns the name of the created file
func CreateTarFile(base, dir string) (string, error) {
	fileBaseName := filepath.Base(dir) + ".tar"
	tarFile, err := os.Create(filepath.Join(base, fileBaseName))
	defer func() {
		err = tarFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()
	dir = filepath.Clean(dir)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 || info.IsDir() {
			if dir == path {
				return nil
			}
			err = writeTarHeader(tarWriter, dir, path, info, true)
			return err
		}
		// regular files are copied into tar, if accessible
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Ignoring file %s: %v", path, err)
			return nil
		}
		defer file.Close()
		if err = writeTarHeader(tarWriter, dir, path, info, true); err != nil {
			fmt.Printf("Error writing header for %q: %v", info.Name(), err)
			return err
		}
		if _, err = io.Copy(tarWriter, file); err != nil {
			fmt.Printf("Error copying file %q to tar: %v", path, err)
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return tarFile.Name(), nil
}

// WriteTarHeader writes tar header for given file, returns error if operation fails
func writeTarHeader(tarWriter *tar.Writer, dir string, path string, info os.FileInfo, includeDirInPath bool) error {
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	prefix := dir
	if includeDirInPath {
		prefix = filepath.Dir(prefix)
	}
	fileName := path
	if prefix != "." {
		fileName = path[1+len(prefix):]
	}
	header.Name = filepath.ToSlash(fileName)
	return tarWriter.WriteHeader(header)
}
