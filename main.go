package main

import (
	"flag"
	"fmt"
	"project-app-inventaris-cli-zahra/database"
	"project-app-inventaris-cli-zahra/handler"
)

func main() {
	// Inisialisasi koneksi database
	if err := database.InitDB(); err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}

	// Definisikan flag command utama
	cmd := flag.String("cmd", "", "Perintah yang ingin dijalankan")
	flag.Parse()

	switch *cmd {
	// Fitur Kategori
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

	// Fitur Barang
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
	case "":
		showHelp()
	default:
		fmt.Println("Perintah tidak dikenali:", *cmd)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("\nUsage: go run main.go -cmd=<command>")
	fmt.Println("\nPerintah yang tersedia:")
	fmt.Println("  Kategori:")
	fmt.Println("    add-category, list-category, edit-category, delete-category, detail-category")
	fmt.Println("  Barang:")
	fmt.Println("    add-item, list-item, edit-item, delete-item, detail-item, check-replacement, search-item")
	fmt.Println("  Laporan:")
	fmt.Println("    report-investment, report-by-id")
}
