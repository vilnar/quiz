package main

import (
	"quiz/internal/common"
	"quiz/internal/importdb"
)

func main() {
	c := common.AskForConfirmation("Do you really want import data?")
	if !c {
		fmt.Printf("Exit")
		return
	}

	importdb.RunImportDb()
}
