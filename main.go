package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
)

func main() {

	ghInput := getGhInputVariables()
	ensureRequiredEnvsPresent(&ghInput)

	fmt.Println(ghInput.openAiApiKey)
	fmt.Println(ghInput.globs)
	fmt.Println(ghInput.targetLangs)

	r, err := git.PlainOpen("./")
	if err != nil {
		ErrLog("Not a git repository. Exiting.")
		os.Exit(1)
		return
	}

	h, err := r.Head()
	fmt.Println(h.Name())
}
