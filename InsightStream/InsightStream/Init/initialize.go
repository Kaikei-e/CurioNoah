package Init

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
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
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	projectRoot := filepath.Dir(wd)

	projectRootEnv := filepath.Join(projectRoot, ".env")

	err = godotenv.Load(projectRootEnv)
	if err != nil {
		_, err := os.Stderr.WriteString(fmt.Sprintf("failed to open .env: %v", err))
		if err != nil {
			panic(err)
		}

		return fmt.Errorf("failed to open .env: %w", err)
	}

	var myEnv map[string]string
	myEnv, err = godotenv.Read()
	if err != nil {
		return fmt.Errorf("failed to read .env: %w", err)
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
		return fmt.Errorf("required env variables are not set: %w", err)
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
			return fmt.Errorf("required env variables are not set: %w", errors.New(strings.Join(want, ", ")))
		}
	}
	return nil
}
