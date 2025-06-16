package cmd

import "project-app-inventaris-cli-zahra/handler"

// function to display report menu
func handleReport(command string) {
	switch command {
	case "report-investment":
		handler.ReportTotalInvestmentCLI()
	case "report-by-id":
		handler.ReportItemByIDCLI()
	}
}
