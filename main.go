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
	var INPUT = flag.String("i", ".", "The input directory or image to be thumnail.")
	var OUTDIR = flag.String("o", "", "The output directory to save thumnail. Default [dir($INPUT)/thumbnail]")
	var MAX_WIDTH = flag.Uint("width", 1280, "The maximum width of output photo.")
	var MAX_HEIGHT = flag.Uint("height", 1280, "The maximum height of output photo.")
	var outdir = ""
	flag.Parse()

	if *OUTDIR != "" {
		outdir = *OUTDIR
	} else {
		outdir = fp.Join(fp.Dir(*INPUT), "thumbnail")
	}

	start := time.Now()
	if libfs.IsDir(*INPUT) {
		fmt.Printf("To be process direcotry: [%v]\n", *INPUT)
		count, _ := thumb_directory(*INPUT, outdir, *MAX_WIDTH, *MAX_HEIGHT)
		fmt.Printf("Processed photos: %v; Time cost: %v\n", count, time.Since(start))
	} else if libfs.IsFile(*INPUT) {
		fmt.Printf("To be process image: %v\n", *INPUT)
		if !libfs.Exist(outdir) {
			err := os.MkdirAll(outdir, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		err := thumbnail(*INPUT, fp.Join(outdir, fp.Base(*INPUT)), *MAX_WIDTH, *MAX_HEIGHT)
		if err != nil {
			fmt.Printf("failed thumbnail: %v\n", err)
		} else {
			fmt.Printf("Time cost: %v\n", time.Since(start))
		}
	}
}
