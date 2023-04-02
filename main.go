package main

import (
	"fmt"
	"log"
)

func main() {
	kanbanManager := &KanbanManager{}
	_, err := kanbanManager.importTodomdFile("TODO.md")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	for _, kanban := range kanbanManager.Kanbans {
		fmt.Printf("Title: '%v'\n", kanban.Title)
		for _, issue := range kanban.TodoList {
			fmt.Printf("Todo: '%v' %v %v %v\n", issue.Name, issue.Depth, issue.Check, issue.Tags)
		}
		for _, issue := range kanban.InProgressList {
			fmt.Printf("In Progress: '%v' %v %v %v\n", issue.Name, issue.Depth, issue.Check, issue.Tags)
		}
		fmt.Println()
	}
}
