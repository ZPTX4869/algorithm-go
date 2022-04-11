package util

import "log"

func CheckOrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func CheckOrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
