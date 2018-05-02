package main

import (
	"flag"
	"fmt"
	"os"
	fp "path/filepath"
	"time"
)

func main() {
	var INDIR = flag.String("i", ".", "The input directory to be thumnail.")
	var OUTDIR = flag.String("o", "", "The output directory to save thumnail. Default [$INPUT/../thumbnail]")
	var MAX_WIDTH = flag.Uint("width", 1280, "The maximum width of output photo.")
	var MAX_HEIGHT = flag.Uint("height", 1280, "The maximum height of output photo.")
	var outdir = ""
	flag.Parse()

	fi, err := os.Stat(*INDIR)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !fi.IsDir() {
		fmt.Fprintln(os.Stderr, "The input path should be a directory!!")
		os.Exit(1)
	}

	if *OUTDIR != "" {
		outdir = *OUTDIR
	} else {
		outdir = fp.Join(fp.Dir(*INDIR), "thumbnail")
	}

	start := time.Now()
	fmt.Printf("To be process direcotry: [%v]\n", *INDIR)
	count, _ := thumb_directory(*INDIR, outdir, *MAX_WIDTH, *MAX_HEIGHT)
	fmt.Printf("Processed photos: %v; Time cost: %v\n", count, time.Since(start))
}
