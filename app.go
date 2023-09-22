package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/tjarratt/babble"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GeneratePassword(wordCount int, seperator string, includeNumbers bool, capitalizeLetters bool) string {
	babbler := babble.NewBabbler()
	babbler.Separator = seperator
	babbler.Count = wordCount

	passPhrase := babbler.Babble()
	finalPassPhrase := ""

	// Split the passPhrase at the seperator to get an array of words
	words := splitString(passPhrase, seperator)

	// If capitalizeLetters is true, capitalize the first letter of each word
	if capitalizeLetters {
		words = capitalizeRandomLetter(words)
	}

	// If includeNumbers is true, insert a random number at a random position in each word
	if includeNumbers {
		words = insertRandomNumber(words)
	}

	// Join the words in the array with the seperator
	finalPassPhrase = strings.Join(words, seperator)

	return finalPassPhrase
}

// generate a random number between 0 to 9
func generateRandomNumber() int {
	return rand.Intn(10)
}

// splitString splits a string at the seperator and returns an array of strings
func splitString(str string, seperator string) []string {
	return strings.Split(str, seperator)
}

// insertRandomNumber inserts a random number at a random position in the array
func insertRandomNumber(words []string) []string {
	// For each word in the array, insert a random number at a random position
	for i := 0; i < len(words); i++ {
		// Generate a random number between 0 to 9
		randomNumber := generateRandomNumber()

		// The word
		word := words[i]

		// Get a random position in the word
		randomPosition := rand.Intn(len(word) - 1)

		// Insert the random number at the random position
		word = fmt.Sprintf("%s%d%s", word[:randomPosition], randomNumber, word[randomPosition:])

		log.Println(word)

		// Update the word in the array
		words[i] = word
	}

	return words
}

// capitalizeRandomLetter capitalizes a random letter in the array
func capitalizeRandomLetter(words []string) []string {
	// For each word in the array, capitalize a random letter
	for i := 0; i < len(words); i++ {
		// The word
		word := words[i]

		// Lowercase the word
		word = strings.ToLower(word)

		// Get a random position in the word
		randomPosition := rand.Intn(len(word) - 1)

		// Capitalize the letter at the random position
		word = fmt.Sprintf("%s%s%s", word[:randomPosition], strings.ToUpper(string(word[randomPosition])), word[randomPosition+1:])

		log.Println(word)

		// Update the word in the array
		words[i] = word
	}

	return words
}
