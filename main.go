package main

import (
	"fmt"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	// TODO add spicific value degry rerate. e.g : 90 180 270 may be -90 -180 ...
	cmds := os.Args[:]
	if len(cmds) < 3 {
		fmt.Println("usage: \n\t$ pictor <command> photoName [newName]")
		os.Exit(1)
	}

	cmd1 := cmds[1]
	fmt.Println(cmd1)
	if cmd1 != "retate" {
		fmt.Println("\tthere is just \"retate\" command right now.\n\twe provid a mor usful commands soon")
		os.Exit(1)

	}

	photoName := cmds[2] // mybe flag is mor sinse

	src, err := imaging.Open(photoName)
	if err != nil {
		fmt.Printf("failed to open image: %v\n", err)
		os.Exit(1)
	}
	dst := imaging.Rotate90(src)

	err = imaging.Save(dst, photoName)
	if err != nil {
		fmt.Printf("failed to save image: %v\n", err)
	}
}
