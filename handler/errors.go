package handler

import (
	"log"
)

func CheckErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
