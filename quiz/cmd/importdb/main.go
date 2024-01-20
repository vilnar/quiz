package main

import (
	"log"
	"quiz/internal/common"
	"quiz/internal/importdb"
)

func main() {
	c := common.AskForConfirmation("Do you really want import data?")
	if !c {
		log.Printf("Exit")
		return
	}

	importdb.RunImportDb()
}
