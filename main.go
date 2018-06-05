package main

import (
	"flag"
	"fmt"
	"os"
	fp "path/filepath"
	"time"

	libfs "github.com/weaming/golib/fs"
)

func main() {
	var INDIR = flag.String("i", ".", "The input directory to be thumnail.")
	var OUTDIR = flag.String("o", "", "The output directory to save thumnail. Default [$INPUT/../thumbnail]")
	var MAX_WIDTH = flag.Uint("width", 1280, "The maximum width of output photo.")
	var MAX_HEIGHT = flag.Uint("height", 1280, "The maximum height of output photo.")
	var outdir = ""
	flag.Parse()

	if *OUTDIR != "" {
		outdir = *OUTDIR
	} else {
		outdir = fp.Join(fp.Dir(*INDIR), "thumbnail")
	}

	start := time.Now()
	if libfs.IsDir(*INDIR) {
		fmt.Printf("To be process direcotry: [%v]\n", *INDIR)
		count, _ := thumb_directory(*INDIR, outdir, *MAX_WIDTH, *MAX_HEIGHT)
		fmt.Printf("Processed photos: %v; Time cost: %v\n", count, time.Since(start))
	} else if libfs.IsFile(*INDIR) {
		fmt.Printf("To be process image: %v\n", *INDIR)
		if !libfs.Exist(outdir) {
			err := os.MkdirAll(outdir, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		err := thumbnail(*INDIR, fp.Join(outdir, fp.Base(*INDIR)), *MAX_WIDTH, *MAX_HEIGHT)
		if err != nil {
			fmt.Printf("failed thumbnail: %v\n", err)
		} else {
			fmt.Printf("Time cost: %v\n", time.Since(start))
		}
	}
}
