package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== Started ===")

	var (
		inputFileName     = flag.String("i", "", "Use a file with the name file-name as an input.")
		sortingFile       = flag.Int("f", 0, "Sort input lines by value number N")
		outputFileName    = flag.String("o", "", "Use a file with the name file-name as an output.")
		ignoredHeaderLine = flag.Bool("h", false, "The first line is a header that must be ignored during sorting but included in the output.")
		sortingFileRevers = flag.Bool("r", false, "Sort input lines in reverse order.")
	)
	flag.Parse()
	if inputFileName != nil {
		fmt.Printf("Input file name: %s\n", *inputFileName)
	}
	if sortingFile != nil {
		fmt.Printf("sortingFile: %d\n", *sortingFile)
	}
	if outputFileName != nil {
		fmt.Printf("outputFileName: %s\n", *outputFileName)
	}
	if ignoredHeaderLine != nil {
		fmt.Printf("ignoredHeaderLine: %s\n", *inputFileName)
	}
	if sortingFileRevers != nil {
		fmt.Printf("sortingField: %v\n", *sortingFileRevers)
	}

	s := bufio.NewScanner(os.Stdin)

	n := 0

	content := [][]string{}

	for s.Scan() {
		line := s.Text()
		space := strings.TrimSpace(line)
		row := strings.Split(space, ",")
		sort.Slice(content, func(i, j int) bool { return content[i][0] < content[j][0] })
		if n == 0 {
			n = len(row)
		}
		if n != len(row) {
			log.Fatalf("ERROR: row has %d columns, but must %d\\n", len(row), n)
		}
		content = append(content, row)
		if line == "" {
			break
		}
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	fmt.Printf("result : %vdf\n", content)

	fmt.Println("=== Finished ===")
}
