package handler

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-zahra/services"
	"strconv"
	"text/tabwriter"
	"time"
)

// Displays items function
func CheckReplacementCLI() {
	items, err := services.GetItemsOver100Days()
	if err != nil {
		fmt.Println("Gagal mendapatkan data barang:", err)
		return
	}

	if len(items) == 0 {
		fmt.Println("Tidak ada barang yang perlu diganti.")
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tNama\tHari Penggunaan\tTanggal Beli")
	for _, item := range items {
		days := int(time.Since(item.PurchaseDate).Hours() / 24)
		fmt.Fprintf(writer, "%d\t%s\t%d\t%s\n", item.ID, item.Name, days, item.PurchaseDate.Format("2006-01-02"))
	}
	writer.Flush()
}

// Investment and Depreciation Report function
func ReportTotalInvestmentCLI() {
	total, err := services.GetTotalInvestmentValue()
	if err != nil {
		fmt.Println("Gagal menghitung nilai investasi:", err)
		return
	}
	fmt.Printf("Total nilai investasi setelah depresiasi: Rp%.2f\n", total)
}

// Displays investment and depreciation values for specific items based on ID.
func ReportItemByIDCLI() {
	var input string
	fmt.Print("Masukkan ID barang: ")
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID tidak valid")
		return
	}

	original, depreciated, err := services.GetItemInvestmentValueByID(id)
	if err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Printf("Nilai awal: Rp%.2f\n", original)
	fmt.Printf("Nilai setelah depresiasi: Rp%.2f\n", depreciated)
}
