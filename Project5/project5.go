package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"flag"
)

type Chapter struct {
	Title    string   `json:"title"`
	Story    []string `json:"story"`
	Options  []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func GetChapterInput(optionsLength int) (int, error) {
	var chosen int
	_, err := fmt.Scan(&chosen)
	if err != nil {
		return chosen, err
	}

	for chosen >= optionsLength {
		fmt.Println("Not a valid index for a chapter")
		_, err := fmt.Scan(&chosen)
		if err != nil {
			return chosen, err
		}
	}
	return chosen, nil
}

func ReadChaptersFromJSON(jsonPath string) (map[string]Chapter, error) {

	var chapters map[string]Chapter
	jsonFile, err := os.Open(jsonPath)
	
	if err != nil {
		return chapters, err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return chapters, err
	}

	json.Unmarshal(byteValue, &chapters)
	if err != nil {
		fmt.Println(err)
		return chapters, err
	}
	return chapters, nil

}

func main() {

	
    jsonPath := flag.String("path", "story.json", "The path for the json file")
	flag.Parse()
	
	chapters, err := ReadChaptersFromJSON(*jsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	currentChapter := "intro"

	for {
		if value, ok := chapters[currentChapter]; ok {
			fmt.Println(value.Title)
			for _, story := range value.Story {
				fmt.Println(story)
			}

			if len(value.Options) == 0 {
				fmt.Println("End of the story :D")
				return
			}
			for index, option := range value.Options {
				fmt.Printf("%d. %s\n", index, option.Text)
			}

			fmt.Println("Please write the index of the answer you want to:")
			chosen, err := GetChapterInput(len(value.Options))
			if err != nil {
				fmt.Println(err)
				return
			}
			currentChapter = value.Options[chosen].Arc

		}
	}

}