package wc

import (
	"bufio"
	"io"
	"os"
)

type Wc struct {
	Path string
}

// Outputs the number of bytes in a file or Stdin
func (w *Wc) ByteCount() (int64, error) {
	if w.Path != "" {
		fs, err := os.Stat(w.Path)
		if err != nil {
			return 0, err
		}

		return fs.Size(), nil
	}

	fs, err := os.Stdin.Stat()
	if err != nil {
		return 0, err
	}

	return fs.Size(), nil
}

// Outputs the number of lines in a file or Stdin
func (w *Wc) LineCount() (int, error) {
	file, err := os.Open(w.Path)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lc := 0

	// Iterate over each line
	for scanner.Scan() {
		lc++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lc, nil
}

// Outputs the number of words in a file
func (w *Wc) WordCount() (int, error) {
	var file *os.File
	var input io.Reader
	var err error

	if w.Path != "" {
		file, err = os.Open(w.Path)
		if err != nil {
			return 0, err
		}

		defer file.Close()
	}

	input = os.Stdin
	if file != nil {
		input = file
	}

	// Create a scanner to read the file
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	wc := 0

	// Iterate over each word
	for scanner.Scan() {
		wc++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wc, nil
}

// Outputs the number of characters in a file
func (w *Wc) CharacterCount() (int, error) {
	file, err := os.Open(w.Path)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	cc := 0

	// Iterate over each character
	for scanner.Scan() {
		cc++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return cc, nil
}
