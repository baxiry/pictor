package main

import (
	"flag"
	"fmt"

	"github.com/disintegration/imaging"
	_ "github.com/disintegration/imaging"
)

func main() {
	photoName := flag.String("old", "", "name of photo that need retate, <old photo>")
	resultName := flag.String("new", "", "name of resutl, if not set new name then result tak old name. \n NOTE that old phot will delete in this case")
	flag.Parse()
	if *resultName == "" {
		//resultName = photoName
		fmt.Println("please set new photo name")
	}

	 src, err := imaging.Open(*photoName)
	 if err != nil {
		fmt.Printf("failed to open image: %v", err)
	 }
	 dst := imaging.Rotate90(src)

	 err = imaging.Save(dst, *resultName)
	 if err != nil {
		fmt.Printf("failed to save image: %v", err)
	}
}

