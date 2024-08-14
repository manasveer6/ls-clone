package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

var suffixes [5]string

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {
	if size == 0 {
		return "0 B"
	}
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func main() {
	longFormat := flag.Bool("l", false, "display files in long format")
	reverseFormat := flag.Bool("r", false, "display files in reverse sorted order")
	allFilesFormat := flag.Bool("a", false, "display all files including dotfiles")
	flag.Parse()

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	if *reverseFormat {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() > files[j].Name()
		})
	} else {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() < files[j].Name()
		})
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	if *longFormat {
		fmt.Fprintf(w, "Name\tMode\tSize\n")
	}

	for _, file := range files {
		if *longFormat {
			info, err := file.Info()
			if err != nil {
				log.Printf("Failed to get info for file: %v", err)
				continue
			}

			name := info.Name()
			mode := info.Mode()
			size := HumanFileSize(float64(info.Size()))

			if !*allFilesFormat {
				if name[0] == '.' {
					continue
				}
			}

			if file.IsDir() {
				fmt.Fprintf(w, "%s/\t", name)
			} else {
				fmt.Fprintf(w, "%s\t", name)
			}
			fmt.Fprintf(w, "%v\t%s\n", mode, size)
		} else {
			fmt.Println(file.Name())
		}
	}
	w.Flush()
}
