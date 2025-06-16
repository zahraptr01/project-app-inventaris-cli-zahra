package handler

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-zahra/services"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func ListItemCLI() {
	items, err := services.ListItems()
	if err != nil {
		fmt.Println("Gagal mengambil data barang:", err)
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tNama\tKategori ID\tHarga\tTanggal Beli\tHari Pakai")
	for _, i := range items {
		days := int(time.Since(i.PurchaseDate).Hours() / 24)
		fmt.Fprintf(writer, "%d\t%s\t%d\t%.2f\t%s\t%d\n",
			i.ID, i.Name, i.CategoryID, i.Price, i.PurchaseDate.Format("2006-01-02"), days)
	}
	writer.Flush()
}

func AddItemCLI() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama barang: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Masukkan ID kategori: ")
	categoryStr, _ := reader.ReadString('\n')
	categoryID, _ := strconv.Atoi(strings.TrimSpace(categoryStr))

	fmt.Print("Masukkan harga barang: ")
	priceStr, _ := reader.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	fmt.Print("Masukkan tanggal beli (yyyy-mm-dd): ")
	dateStr, _ := reader.ReadString('\n')
	purchaseDate, err := time.Parse("2006-01-02", strings.TrimSpace(dateStr))
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}

	err = services.AddItem(strings.TrimSpace(name), categoryID, price, purchaseDate)
	if err != nil {
		fmt.Println("Gagal menambahkan barang:", err)
	} else {
		fmt.Println("Barang berhasil ditambahkan!")
	}
}

func DetailItemCLI() {
	var input string
	fmt.Print("Masukkan ID barang: ")
	fmt.Scanln(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID tidak valid")
		return
	}

	item, err := services.GetItemDetail(id)
	if err != nil {
		fmt.Println("Gagal:", err)
	} else {
		fmt.Println("ID:", item.ID)
		fmt.Println("Nama:", item.Name)
		fmt.Println("Kategori ID:", item.CategoryID)
		fmt.Println("Harga:", item.Price)
		fmt.Println("Tanggal Beli:", item.PurchaseDate.Format("2006-01-02"))
	}
}

func EditItemCLI() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID barang: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Nama baru: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("ID kategori baru: ")
	categoryStr, _ := reader.ReadString('\n')
	categoryID, _ := strconv.Atoi(strings.TrimSpace(categoryStr))

	fmt.Print("Harga baru: ")
	priceStr, _ := reader.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	fmt.Print("Tanggal beli baru (yyyy-mm-dd): ")
	dateStr, _ := reader.ReadString('\n')
	purchaseDate, err := time.Parse("2006-01-02", strings.TrimSpace(dateStr))
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}

	err = services.EditItem(id, strings.TrimSpace(name), categoryID, price, purchaseDate)
	if err != nil {
		fmt.Println("Gagal mengedit barang:", err)
	} else {
		fmt.Println("Barang berhasil diperbarui.")
	}
}

func DeleteItemCLI() {
	var input string
	fmt.Print("Masukkan ID barang: ")
	fmt.Scanln(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID tidak valid")
		return
	}
	err = services.RemoveItem(id)
	if err != nil {
		fmt.Println("Gagal menghapus barang:", err)
	} else {
		fmt.Println("Barang berhasil dihapus.")
	}
}
