package codewriter

import (
	"os"
	"sync"
)

type FILE_TYPE int

const (
	// there are only 3 instancies we whant to write
	MAIN_FILE FILE_TYPE = iota
	SOLUTION_FILE
	TEST_FILE
)

const (
	main_file_name     = "main.go"
	solution_file_name = "solution.go"
	test_file_name     = "test.json"
)

var fileTypeToName = map[FILE_TYPE]string{
	MAIN_FILE:     main_file_name,
	SOLUTION_FILE: solution_file_name,
	TEST_FILE:     test_file_name,
}
var once sync.Once

type (
	writer struct {
		mutex *sync.Mutex
		path  string
	}

	Writer interface {
		Write(FILE_TYPE, string) error
	}
)

var theOnlyWriter *writer

// Returns the only instance of a writer
func Get(path string) Writer {
	once.Do(func() {
		theOnlyWriter = &writer{&sync.Mutex{}, path}
	})
	return theOnlyWriter
}

func (w *writer) Write(t FILE_TYPE, data string) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	file, err := os.Open(w.path + fileTypeToName[t])
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}
