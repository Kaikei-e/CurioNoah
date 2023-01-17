package Init

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
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

	err = godotenv.Load(projectRootEnv)
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("failed to open .env: %v", err))
		if err != nil {
			panic(err)
		}

		return errors.New("failed to open .env")
	}

	var myEnv map[string]string
	myEnv, err = godotenv.Read()
	if err != nil {
		return errors.New("failed to open .env")
	}

	var envs []string
	for k := range myEnv {
		envs = append(envs, k)
	}

	err = envRequired(envs, getSlice())
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("%v", err))
		if err != nil {
			panic(err)
		}
		return errors.New("required env is not set")
	}

	return nil
}

func envRequired(envs []string, want []string) error {
	slices.Sort(envs)
	slices.Sort(want)

	for i, env := range envs {
		if strings.Compare(env, want[i]) == 0 {
			continue
		} else {
			return errors.New(fmt.Sprintf("failed to parse .env: %v is required", want[i]))
		}
	}
	return nil
}
