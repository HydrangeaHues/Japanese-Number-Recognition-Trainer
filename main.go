package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	htgotts "github.com/hegedustibor/htgo-tts"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

// TO DO:
// 1. Implement optional timer on answering questions
// 2. Implement stats report after the quiz is finished.
// 3. Clean up code.

func main() {

	var audioFolderPath = "audio"
	var quizSize = 5
	var correctAnswerCount = 0
	var userInput int

	for i := 0; i < quizSize; i++ {
		numberToRead := rand.Intn(9999) + 1
		PlayAudio(audioFolderPath, numberToRead)

		for {
			fmt.Print("Enter the number you heard: ")
			fmt.Scan(&userInput)
			if userInput == numberToRead {
				fmt.Println("Correct!")
				correctAnswerCount++
				break
			} else if userInput == -1 {
				fmt.Println("Retrying?")
				speech := htgotts.Speech{Folder: audioFolderPath, Language: voices.Japanese}
				res := speech.PlaySpeechFile("audio/current_question.mp3")
				if res != nil {
					fmt.Println("Error playing speech file: ", res)
				}
			} else {
				fmt.Println("Incorrect.")
				break
			}
		}
		os.RemoveAll(audioFolderPath)
	}
}

func PlayAudio(audioFolderPath string, numberToRead int) {
	speech := htgotts.Speech{Folder: audioFolderPath, Language: voices.Japanese}
	filename, error := speech.CreateSpeechFile(strconv.Itoa(numberToRead), "current_question")
	if error != nil {
		fmt.Println("Error creating speech file: ", error)
		return
	}
	speech.PlaySpeechFile(filename)
}
