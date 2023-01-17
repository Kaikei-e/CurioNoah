package Init

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getSlice() []string {
	return []string{
		"MYSQL_USER",
		"MYSQL_PASSWORD",
		"MYSQL_DATABASE",
		"MYSQL_ADDR",
		"NET_TYPE",
	}
}

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
	envs := strings.Split(fileBody, "\n")
	err = envRequired(envs, getSlice())
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("failed to parse .env: %v", err))
		if err != nil {
			panic(err)
		}
		return errors.New("required env is not set")
	}

	fmt.Println(fileBody)

	return nil
}

func envRequired(envs []string, want []string) error {
	for i, env := range envs {
		if !cmp.Equal(env, want[i]) {
			return errors.New(fmt.Sprintf("failed to parse .env: %v is required", want))
		}
	}
	return nil
}
