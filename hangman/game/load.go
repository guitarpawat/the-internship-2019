package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path"
	"strconv"
	"time"
)

type Words struct {
	Category string `json:"category"`
	Words    []Word `json:"words"`
}

type Word struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

var wordsDir = path.Join(".", "words")

var categories []Words

func load(signal chan<- struct{}) {
	defer func(signal chan<- struct{}) {
		signal <- struct{}{}
	}(signal)

	files, err := ioutil.ReadDir(wordsDir)
	if err != nil {
		log.Fatalln("error while opening words directory:", err)
	}

	for _, file := range files {
		filePath := path.Join(wordsDir, file.Name())
		raw, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("error while reading words at file %s: %s", file.Name(), err)
		}

		newCategory := Words{}
		err = json.Unmarshal(raw, &newCategory)
		if err != nil {
			log.Fatalf("error while unmarshal json from file `%s`: %s", file.Name(), err)
		}
		categories = append(categories, newCategory)
	}
}

func showLoading() {
	ticker := time.NewTicker(350 * time.Millisecond)
	signal := make(chan struct{})

	defer func(*time.Ticker) {
		ticker.Stop()
	}(ticker)

	go load(signal)

	fmt.Print("Loading ...")
	for {
		select {
		case <-signal:
			fmt.Println("")
			return

		case <-ticker.C:
			fmt.Print(".")
		}
	}
}

func selectCategory() int {

	fmt.Println("Select Category:")
	for i, v := range categories {
		fmt.Printf("%d\t%s\n", i+1, v.Category)
	}

	fmt.Print("Please select category number: ")
	var input string
	fmt.Scanln(&input)

	choice64, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return selectCategory()
	}

	choice := int(choice64)
	if choice <= 0 || choice > len(categories) {
		return selectCategory()
	}

	return choice - 1
}

func randWord(words Words) Word {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(words.Words))
	return words.Words[num]
}
