package hazm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type WordInfo struct {
	Frequency int
	Tags      []string
}

const (
	embeddedWordsRel         = "data/words.dat"
	embeddedVerbsRel         = "data/verbs.dat"
	embeddedAbbreviationsRel = "data/abbreviations.dat"
)

// DefaultWordsPath returns the path used for embedded defaults (see openDataFile).
func DefaultWordsPath() string {
	return embeddedWordsRel
}

func DefaultVerbsPath() string {
	return embeddedVerbsRel
}

func DefaultAbbreviationsPath() string {
	return embeddedAbbreviationsRel
}

// openDataFile opens path for reading. If path is empty, it uses embedRel inside embeddedData.
// If path is non-empty: absolute paths and paths that exist on disk are read from the filesystem;
// otherwise the path is opened from the embedded FS.
func openDataFile(path, embedRel string) (io.ReadCloser, error) {
	if path == "" {
		return embeddedData.Open(embedRel)
	}
	if filepath.IsAbs(path) {
		return os.Open(path)
	}
	if _, err := os.Stat(path); err == nil {
		return os.Open(path)
	}
	f, err := embeddedData.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open data %q: %w", path, err)
	}
	return f, nil
}

func LoadWords(path string) (map[string]WordInfo, error) {
	rc, err := openDataFile(path, embeddedWordsRel)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	words := make(map[string]WordInfo)
	scanner := bufio.NewScanner(rc)
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
	rc, err := openDataFile(path, embeddedVerbsRel)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	verbs := make([]string, 0, 700)
	scanner := bufio.NewScanner(rc)
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
	rc, err := openDataFile(path, embeddedAbbreviationsRel)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	result := []string{}
	scanner := bufio.NewScanner(rc)
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
