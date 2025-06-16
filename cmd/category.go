package cmd

import "project-app-inventaris-cli-zahra/handler"

// function to display category menu
func handleCategory(command string) {
	switch command {
	case "add-category":
		handler.AddCategoryCLI()
	case "list-category":
		handler.ListCategoryCLI()
	case "edit-category":
		handler.EditCategoryCLI()
	case "delete-category":
		handler.DeleteCategoryCLI()
	case "detail-category":
		handler.DetailCategoryCLI()
	}
}
