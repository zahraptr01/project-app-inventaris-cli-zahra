package cmd

import (
	"flag"
	"fmt"
	"project-app-inventaris-cli-zahra/database"
)

func Execute() {
	// connect to DB
	if err := database.InitDB(); err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}

	// Main Flag
	cmd := flag.String("cmd", "", "Perintah yang ingin dijalankan")
	flag.Parse()

	// Routing to commands according to category
	switch *cmd {
	case "add-category", "list-category", "edit-category", "delete-category", "detail-category":
		handleCategory(*cmd)

	case "add-item", "list-item", "edit-item", "delete-item", "detail-item", "search-item", "check-replacement":
		handleItem(*cmd)

	case "report-investment", "report-by-id":
		handleReport(*cmd)

	case "":
		showHelp()

	default:
		fmt.Println("Perintah tidak dikenali:", *cmd)
		showHelp()
	}
}

// function displays the entire menu
func showHelp() {
	fmt.Println("\nUsage: go run main.go -cmd=<command>")
	fmt.Println("\nKategori:")
	fmt.Println("  add-category, list-category, edit-category, delete-category, detail-category")
	fmt.Println("Barang:")
	fmt.Println("  add-item, list-item, edit-item, delete-item, detail-item, search-item, check-replacement")
	fmt.Println("Laporan:")
	fmt.Println("  report-investment, report-by-id")
}
