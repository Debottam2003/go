package main

import "log"

func Error_Reaction(err error) {
	if err != nil {
		log.Fatal(err)
	}
}