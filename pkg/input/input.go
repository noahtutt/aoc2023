package input

import (
	"fmt"
	"os"
)

func DeferFunc(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to close file: %w", err))
	}
}

func GetInputFile(exercise string, name string) (*os.File, error) {

	f, err := os.Open(fmt.Sprintf("./input/%s/%s", exercise, name))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return f, nil
}
