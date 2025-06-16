package cmd

import "project-app-inventaris-cli-zahra/handler"

// function to display item menu
func handleItem(command string) {
	switch command {
	case "add-item":
		handler.AddItemCLI()
	case "list-item":
		handler.ListItemCLI()
	case "edit-item":
		handler.EditItemCLI()
	case "delete-item":
		handler.DeleteItemCLI()
	case "detail-item":
		handler.DetailItemCLI()
	case "search-item":
		handler.SearchItemCLI()
	case "check-replacement":
		handler.CheckReplacementCLI()
	}
}
