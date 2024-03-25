package analysis

import (
	"bufio"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func CheckIfEnvExists(fileurl string) (string, error) {
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

func ReadContents(fileurl string) ([]string, error) {
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
			vals = append(vals, s.Text())
		}
	}

	return vals, nil
}

func Obfuscate(val string) string {
	length := len(val)

	// How many characters to not replace
	numToPreserve := 2
	if length > 5 {
		numToPreserve = 3
	} else if length <= 2 {
		numToPreserve = 1
	}

	// Preserve the first few characters
	prefix := val[:numToPreserve]

	// Replace the remaining characters with asterisks
	numToReplace := length - numToPreserve
	if numToReplace > 20 {
		numToReplace = numToReplace / 2
	}
	asterisks := strings.Repeat("*", numToReplace)

	return prefix + asterisks
}
