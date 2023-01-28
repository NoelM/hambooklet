package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	var filename string
	if len(args) > 0 {
		filename = args[0]
	} else {
		log.Fatalf("missing filename, usage: ./groupby [filename]")
	}

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to openfile %s: %s", filename, err.Error())
	}

	csvInputFile := csv.NewReader(inputFile)

	outputFile, err := os.Create(fmt.Sprintf("group_%s", filename))
	if err != nil {
		log.Fatalf("unable to openfile %s: %s", fmt.Sprintf("group_%s", filename), err.Error())
	}
	csvOutputFile := csv.NewWriter(outputFile)

	lineId := 0

	var prefixRange, firstPrefixRange, previousPrefixRange string
	var country, previousCountry string
	for {
		record, err := csvInputFile.Read()
		lineId++
		if err != nil && err != io.EOF {
			log.Fatalf("unable to read line %d", lineId)
		}
		if lineId == 1 {
			continue
		}

		previousPrefixRange = prefixRange
		previousCountry = country

		if len(record) > 0 {
			prefixRange = record[0]
			country = record[1]
		}

		if country != previousCountry || err == io.EOF {
			if firstPrefixRange != "" && previousPrefixRange != "" {
				cleanName := strings.TrimSpace(strings.Split(previousCountry, "(")[0])

				if firstPrefixRange == previousPrefixRange {
					csvOutputFile.Write([]string{firstPrefixRange, cleanName})
				} else {
					start := strings.TrimSpace(strings.Split(firstPrefixRange, "-")[0])
					stop := strings.TrimSpace(strings.Split(previousPrefixRange, "-")[1])
					completeRange := fmt.Sprintf("%s - %s", start, stop)

					csvOutputFile.Write([]string{completeRange, cleanName})
				}
			}

			firstPrefixRange = prefixRange
		}

		if err == io.EOF {
			break
		}
	}

	inputFile.Close()
	csvOutputFile.Flush()
	outputFile.Close()
}
