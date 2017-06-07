package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

type Word struct {
	Word string `json:"word"`
	Definitions []struct {
		Definition string `json:"definition"`
		PartOfSpeech string `json:"partOfSpeech"`
	} `json:"definitions"`
}

var Words []Word = make([]Word, 0)

var RawWords string

func DictionarySetup() {
	file, e := ioutil.ReadFile("1st2000.json")
	if e != nil {
		fmt.Errorf("Error reading file: %v", e)
		os.Exit(1)
	}
	RawWords = string(file)
	e = json.Unmarshal(file, &Words)
	if e != nil {
		fmt.Errorf("Error parsing JSON: %v", e)
		os.Exit(1)
	}
}

func SingleWordAsString() string {
	n := rand.Intn(len(Words))
	jsonString, e := json.Marshal(Words[n])
	if e != nil {
		fmt.Errorf("Error stringifying JSON: %v", e)
		os.Exit(1)
	}
	return string(jsonString)
}

func DefineSingleWord(rWord string) string {
	for _, value := range Words {
		if value.Word == rWord {
			jsonString, e := json.Marshal(value)
			if e != nil {
				fmt.Errorf("Error stringifying JSON: %v", e)
				os.Exit(1)
			}
			return string(jsonString)
		}
	}
	return ""
}