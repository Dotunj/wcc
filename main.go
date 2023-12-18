package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dotunj/wc/wc"
)

var (
	byteCount bool
	lines     bool
	words     bool
	cc        bool
)

func getPath() string {
	var path string
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	return path
}

func main() {
	parseFlags()

	if byteCount {
		path := getPath()
		wc := &wc.Wc{Path: path}
		info, err := wc.ByteCount()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s\n", info, path)
	}

	if lines {
		path := getPath()
		wc := &wc.Wc{Path: path}

		count, err := wc.LineCount()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s\n", count, path)
	}

	if words {
		path := getPath()
		wc := &wc.Wc{Path: path}

		count, err := wc.WordCount()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s\n", count, path)
	}

	if cc {
		path := getPath()
		wc := &wc.Wc{Path: path}

		count, err := wc.CharacterCount()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s\n", count, path)
	}

	if !byteCount && !lines && !words {
		if len(os.Args) > 0 {
			path := getPath()
			ab := &wc.Wc{Path: path}

			wcc, err := ab.WordCount()
			if err != nil {
				log.Fatal(err)
			}

			lc, err := ab.LineCount()
			if err != nil {
				log.Fatal(err)
			}

			bc, err := ab.ByteCount()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%d %d %d %s\n", lc, wcc, bc, path)
		}

	}
}

func parseFlags() {
	flag.BoolVar(&byteCount, "c", false, "Outputs the number of bytes in a file")
	flag.BoolVar(&lines, "l", false, "Outputs the number of lines in a file")
	flag.BoolVar(&words, "w", false, "Outputs the number of words in a file")
	flag.BoolVar(&cc, "m", false, "Outputs the number of characters in a file")

	flag.Parse()
}
