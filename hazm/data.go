package hazm

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type WordInfo struct {
	Frequency int
	Tags      []string
}

func defaultDataDir() string {
	candidates := []string{
		"data",
		"../hazm/data",
		"../../hazm/data",
	}

	if _, file, _, ok := runtime.Caller(0); ok {
		base := filepath.Dir(file)
		candidates = append(candidates,
			filepath.Join(base, "..", "..", "hazm", "data"),
			filepath.Join(base, "..", "data"),
		)
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	return "data"
}

func DefaultWordsPath() string {
	return filepath.Join(defaultDataDir(), "words.dat")
}

func DefaultVerbsPath() string {
	return filepath.Join(defaultDataDir(), "verbs.dat")
}

func DefaultAbbreviationsPath() string {
	return filepath.Join(defaultDataDir(), "abbreviations.dat")
}

func LoadWords(path string) (map[string]WordInfo, error) {
	if path == "" {
		path = DefaultWordsPath()
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make(map[string]WordInfo)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")
		if len(parts) != 3 {
			continue
		}

		freq, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		tags := strings.Split(parts[2], ",")
		words[parts[0]] = WordInfo{
			Frequency: freq,
			Tags:      tags,
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func LoadVerbs(path string) ([]string, error) {
	if path == "" {
		path = DefaultVerbsPath()
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	verbs := make([]string, 0, 700)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		verb := strings.TrimSpace(scanner.Text())
		if verb == "" {
			continue
		}
		verbs = append(verbs, verb)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return verbs, nil
}

func LoadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
