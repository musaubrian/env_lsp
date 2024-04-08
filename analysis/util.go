package analysis

import (
	"bufio"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func checkIfEnvExists(fileurl string) (string, error) {
	fileURL, err := url.ParseRequestURI(fileurl)
	if err != nil {
		return "", err
	}

	filePath := fileURL.Path
	parentDir := filepath.Dir(filePath)

	envFiles := []string{".env", ".env.local"}
	var foundFile string
	for _, file := range envFiles {
		fileToCheck := filepath.Join(parentDir, file)
		_, err = os.Stat(fileToCheck)
		if err == nil {
			foundFile = fileToCheck
			break
		} else if !os.IsNotExist(err) {
			return "", err
		}
	}

	if foundFile != "" {
		return foundFile, err
	}

	return "", nil
}

func readContents(fileurl string) ([]string, error) {
	var vals []string
	f, err := os.Open(fileurl)
	defer f.Close()
	if err != nil {
		return vals, err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		// Ignore comments and empty lines
		if !strings.Contains(s.Text(), "#") && len(s.Text()) > 1 {
			vals = append(vals, s.Text()+"REXIQI")
		}
	}

	return vals, nil
}

func loadEnvs(envs []string) error {
	for _, v := range envs {
		env := strings.Split(v, "=")
		if len(env[1]) > 1 {
			if err := os.Setenv(env[0], env[1]); err != nil {
				return err
			}
		}
	}

	return nil
}

func obfuscate(val string) string {
	length := len(val)
	v := strings.Join(strings.Split(val, "")[:len(val)-6], "") //remove the `REXIQI` suffix

	if length < 1 {
		return "No value set!"
	}

	if length >= 1 && length < 4 {
		return v
	}

	prefix := v[:3]

	asterisks := strings.Repeat("*", 7)

	return prefix + asterisks
}
