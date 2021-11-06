import (
	"bufio"
	"os"
	"sync"
	"containters"
)

type FileType int

const (
	// there are only 3 instancies we whant to write
	MAIN_FILE FileType = iota
	SOLUTION_FILE
	TEST_FILE
)

const (
	main_file_name     = "main.go"
	solution_file_name = "solution.go"
	test_file_name     = "test.json"
)

var fileTypeToName = map[FileType]string{
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
		Write(FileType, string) error
		WriteBuffered(FileType, bufio.Reader) error
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

func (w *writer) Write(t FileType, data string) error {
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

func (w *writer) WriteBuffered(t FileType, reader bufio.Reader) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	result := containers.NewQueue()
	buffer := make([]byte, 0, 1<<64)

	var err error
	for n := 1; n != 0; {
		n, err = reader.ReadString(buffer)
		if err != nil {
			return err
		}
		err = w.Write(t, string(buffer))
		if err != nil {
			return err
		}
	}

	return
}
