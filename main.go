package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	kanbanManager := &KanbanManager{}
	inputFilename := os.Args[1]
	outputCsvFilename := os.Args[2]
	_, err := kanbanManager.importTodomdFile(inputFilename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	for _, kanban := range kanbanManager.Kanbans {
		fmt.Printf("Title: '%v'\n", kanban.Title)
		for _, issue := range kanban.Issues {
			fmt.Printf("Todo: '%v' %v %v %v\n", issue.Summary, issue.Depth, issue.Check, issue.Tags)
		}
		fmt.Println()
	}

	csvDate, err := kanbanManager.GenerateCsvData()
	if err != nil {
		log.Fatalf("%v", err)
	}
	ioutil.WriteFile(outputCsvFilename, []byte(csvDate), 0644)
	fmt.Println(string(csvDate))
}
