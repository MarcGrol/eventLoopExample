package main

import (
	"log"
)

func enter(name string) string {
	log.Printf("enter %s", name)
	return name
}

func leave(name string) {
	log.Printf("leave %s", name)
}
