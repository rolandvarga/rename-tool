package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	dir = "/Volumes/stuff/tv/Batman_The_Animated_Series_1992"
)
type final struct {
	Files []outputFile `json:"files"`
}

type outputFile struct {
	Episode string `json:"episode"`
	Title   string `json:"title"`
	Season  int    `json:"season"`
}

func main() {
	sample, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer sample.Close()

	content, err := ioutil.ReadAll(sample)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	finalF := &final{}

	err = json.Unmarshal(content, finalF)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range finalF.Files {
		fmt.Println(f.Title)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count := 0
	match := 0
	for _, f := range files {
		count++
		for _, final := range finalF.Files {
			if strings.Contains(f.Name(), final.Title) {
				match++

				fmt.Println("***")
				oldName := fmt.Sprintf("%s/%s", dir, f.Name())
				newName := fmt.Sprintf("%s/%s - %s.mkv", dir, final.Episode, final.Title)
        fmt.Println(oldName)
				fmt.Println(newName)

				err := os.Rename(oldName, newName)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}
	fmt.Printf("Total count: %d\nTotal matches: %d\n", count, match)
}
