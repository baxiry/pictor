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
	if cmd1 != "retate" {
		fmt.Println("\tthere is just \"retate\" command right now.\n\twe provid a mor usful commands soon")
		os.Exit(1)

	}

	photoName := cmds[2] // mybe flag is mor sinse
	retated := ""
	if len(cmds) == 4 {
		retated = cmds[3]
	}

	if retated == "" {
		for {
			fmt.Print("\tthe result photo will take some name of source photo\n\tthe source photo will removed. Continue ? (Y/n) ? ")
			var answr string
			fmt.Scanln(&answr)
			if answr == "n" {
				fmt.Println("Bye")
				os.Exit(1)
			} else if answr == "Y" || answr == "y" {
				retated = photoName
				break
			} else {
				fmt.Println("please read the message and type yes or no (Y/n)")
			}
		}
	}

	src, err := imaging.Open(photoName)
	if err != nil {
		fmt.Printf("failed to open image: %v\n", err)
		os.Exit(1)
	}
	dst := imaging.Rotate90(src)

	err = imaging.Save(dst, retated)
	if err != nil {
		fmt.Printf("failed to save image: %v\n", err)
	}
}
