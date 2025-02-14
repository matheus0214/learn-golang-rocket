package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()
	measurements, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}

	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temperature, _ := strconv.ParseFloat(rawTemp, 64)

		measu, ok := data[location]
		if !ok {
			measu = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measu.Count++
			measu.Min = min(measu.Min, temperature)
			measu.Max = min(measu.Max, temperature)
			measu.Sum += temperature
		}

		data[location] = measu
	}

	locations := make([]string, 0, len(data))

	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measu := data[name]

		fmt.Printf("%s=%.1f/%.1f/%.1f, ", name, measu.Min, measu.Sum/float64(measu.Count), measu.Max)
	}
	fmt.Printf("}\n")

	fmt.Println(time.Since(start))
}
