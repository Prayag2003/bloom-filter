package storage

import (
	"bufio"
	"os"
)

type Storage interface {
	Load() ([]string, error)
	Save(item string) error
}

type FileStorage struct {
	FilePath string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{FilePath: path}
}

func (fs *FileStorage) Load() ([]string, error) {
	file, err := os.Open(fs.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (fs *FileStorage) Save(item string) error {
	f, err := os.OpenFile(fs.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(item + "\n")
	return err
}
