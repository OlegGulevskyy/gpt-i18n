package main

import "fmt"

func InfoLog(msg string) {
	fmt.Printf("[gpt-i18n] INFO: %s.", msg)
}

func ErrLog(msg string) {
	fmt.Printf("[gpt-i18n] ERROR: %s.", msg)
}
