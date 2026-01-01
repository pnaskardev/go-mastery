package mylogger

import "log"

func Info(text string) {
	log.Println("[INFO]", text)
}
