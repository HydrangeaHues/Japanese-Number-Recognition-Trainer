package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

// AudioData represents the data required to play audio.
type AudioData struct {
	AudioFilePath string // The file path of the audio file to be played.
	NumberToRead  int    // The number to be read in the audio.
}

// QuizData represents the data required for the quiz.
type QuizData struct {
	QuizSize           int // The size of the quiz.
	CorrectAnswerCount int // The count of correct answers.
	QuizWaitTime       int // The time to wait for the user to answer.
}

func main() {
	answerChannel := make(chan int)
	resultsChannel := make(chan QuizData)
	quizData := QuizData{QuizSize: 5, QuizWaitTime: 10}

	go RunQuiz(resultsChannel, answerChannel, quizData)
	go PollAnswers(answerChannel)

	quizResults := <-resultsChannel
	fmt.Printf("You got %d out of %d correct!\n", quizResults.CorrectAnswerCount, quizResults.QuizSize)
}

// RunQuiz loops through the quiz and plays audio for the user to listen to.
// At the end of the quiz, the results are sent to the resultsChannel and the resultsChannel is closed.
func RunQuiz(resultsChannel chan<- QuizData, answerChannel <-chan int, quizData QuizData) {
	for i := 0; i < quizData.QuizSize; i++ {
		numberToRead := rand.Intn(9999) + 1
		PlayAudio(AudioData{NumberToRead: numberToRead})
		fmt.Print("Enter the number you heard: ")

		// Wait for the user to answer or for the time to run out.
		// Either outcome results in the quiz moving to the next question.
		select {
		case answer := <-answerChannel:
			if answer == numberToRead {
				quizData.CorrectAnswerCount++
			}
		case <-time.After(time.Duration(quizData.QuizWaitTime) * time.Second):
			fmt.Println("Time's up! Moving to the next question.")
		}
		// Clear the audio file after each question to ensure
		// subsequent questions do not play the same audio.
		// This doesn't seem optimal, but I'm not sure if the htgotts package
		// allows for overwriting an audio file upon creating a new file of the same name.
		os.RemoveAll("audio")
	}
	// Sending on the resultsChannel signifies the end of the quiz.
	resultsChannel <- quizData
	close(resultsChannel)
}

// PollAnswers polls for user input and sends the answer to the answerChannel.
// If the user enters -1, the current question is replayed.
func PollAnswers(answerChannel chan<- int) {
	var userInput int
	for {
		fmt.Scan(&userInput)
		if userInput == -1 {
			PlayAudio(AudioData{AudioFilePath: "audio/current_question.mp3"})
		} else {
			answerChannel <- userInput
		}
	}
}

// PlayAudio uses the htgotts package to create mp3 files of the numbers being read and then plays those mp3 files.
func PlayAudio(audioData AudioData) {
	var audioFolderPath = "audio"
	speech := htgotts.Speech{Folder: audioFolderPath, Language: voices.Japanese}
	// The AudioFilePath attribute being blank signifies a new question,
	// where the attribute being populated represents a question being replayed.
	if audioData.AudioFilePath == "" {
		filename, error := speech.CreateSpeechFile(strconv.Itoa(audioData.NumberToRead), "current_question")
		if error != nil {
			fmt.Println("Error creating speech file: ", error)
			return
		}
		audioData.AudioFilePath = filename
	}
	speech.PlaySpeechFile(audioData.AudioFilePath)
}
