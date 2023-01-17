package Init

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Initialize() error {
	wd, err := os.Getwd()
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("failed to get working directory: %v", err))
		if err != nil {
			panic(err)
		}
		return errors.New("failed to get working directory")
	}

	projectRoot := filepath.Dir(wd)

	projectRootEnv := filepath.Join(projectRoot, ".env")

	f, err := os.OpenFile(projectRootEnv, os.O_RDWR, 0644)
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("open .env: %v", err))
		if err != nil {
			panic(err)
		}
		return errors.New("failed to open .env")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	by, err := io.ReadAll(f)
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("failed to read .env: %v", err))
		if err != nil {
			panic(err)
		}
		return errors.New("failed to read .env")
	}

	fileBody := string(by)
	fmt.Println(fileBody)

	return nil
}
