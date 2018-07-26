package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/codahale/hdrhistogram"
)

const (
	intMultiplier = 10000
)

func splitIntoParams(line string) (x, y, e float64) {
	wordScanner := bufio.NewScanner(strings.NewReader(line))
	wordScanner.Split(bufio.ScanWords)
	var err error

	for wordScanner.Scan() {
		switch string(wordScanner.Text()[0]) {
		case "X":
			x, err = strconv.ParseFloat(wordScanner.Text()[1:], 64)
			if err != nil {
				panic(err)
			}
		case "Y":
			y, err = strconv.ParseFloat(wordScanner.Text()[1:], 64)
			if err != nil {
				panic(err)
			}
		case "E":
			e, err = strconv.ParseFloat(wordScanner.Text()[1:], 64)
			if err != nil {
				panic(err)
			}
		}
	}
	return x, y, e
}

func calcDistance(prevX, prevY, x, y float64) float64 {
	xDist := (x) - (prevX)
	yDist := (y) - (prevY)
	return math.Sqrt((xDist * xDist) + (yDist * yDist))
}

func main() {

	var verbose, summary bool
	var maxMove int64

	flag.BoolVar(&verbose, "verbose", false, "Verbose output, i.e. one line for each move")
	flag.BoolVar(&summary, "summary", false, "Show only summary (this overrules verbose mode)")
	flag.Int64Var(&maxMove, "maxMove", 300, "Maximum distance the longest axis can move in mm")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalln("At least one file must be provided")
		os.Exit(1)
	}

	multiFile := len(flag.Args()) > 1
	var hTotal *hdrhistogram.Histogram
	if multiFile || summary {
		hTotal = hdrhistogram.New(0, maxMove*intMultiplier, 5)
	}
	for _, file := range flag.Args() {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalln(err)
			os.Exit(2)
		}
		defer f.Close()
		lineScanner := bufio.NewScanner(f)
		var prevX, prevY float64

		// Find first valid line
		for lineScanner.Scan() {
			line := lineScanner.Text()
			if strings.HasPrefix(line, "G1 ") || strings.HasPrefix(line, "G0 ") {
				prevX, prevY, _ = splitIntoParams(line)
				break
			}
		}
		var h *hdrhistogram.Histogram
		if !summary {
			h = hdrhistogram.New(0, maxMove*intMultiplier, 5)
		}

		// Scan remaining lines
		for lineScanner.Scan() {
			line := lineScanner.Text()
			if !(strings.HasPrefix(line, "G1 ") || strings.HasPrefix(line, "G0 ")) {
				continue
			}

			x, y, e := splitIntoParams(line)

			// Non-head move, probably retract/unretract
			if x == 0 || y == 0 {
				continue
			}

			// Print-move has e > 0
			if e > 0 {
				distance := calcDistance(prevX, prevY, x, y)
				if distance == 0 {
					log.Println("Distance is", distance, " -> Skipping")
					continue
				}
				if !summary {
					h.RecordValue(int64(distance * intMultiplier))
				}
				if multiFile || summary {
					hTotal.RecordValue(int64(distance * intMultiplier))
				}

				if verbose && !summary {
					fmt.Println(distance, "->", e)
				}
			}

			prevX = x
			prevY = y
		}

		if !summary {
			printResult(file, h)
			if multiFile {
				fmt.Println("-------")
			}
		}
	}
	if multiFile || summary {
		printResult("Summary", hTotal)
	}
}

func printResult(heading string, h *hdrhistogram.Histogram) {
	fmt.Println(heading)
	fmt.Printf("Shortest Print Move: %vmm\n", (float64(h.Min()) / intMultiplier))
	fmt.Printf("Average Print Move: %.4fmm\n", (h.Mean() / intMultiplier))
	fmt.Printf("Longest Print Move: %vmm\n", (float64(h.Max()) / intMultiplier))
	fmt.Println("Percentiles:")

	// As we are more interested in very short moves be more fine-granular here
	for i := 1; i <= 9; i++ {
		fmt.Printf("%3d%% of print moves are <= %vmm\n", i, (float64(h.ValueAtQuantile(float64(i))))/intMultiplier)
	}

	// And a bit more coarse starting from 10%
	for i := 10; i <= 100; i += 5 {
		fmt.Printf("%3d%% of print moves are <= %vmm\n", i, (float64(h.ValueAtQuantile(float64(i))))/intMultiplier)
	}
}
