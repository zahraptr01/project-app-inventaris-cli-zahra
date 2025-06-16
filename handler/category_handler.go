package handler

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-zahra/services"
	"strconv"
	"strings"
	"text/tabwriter"
)

func ListCategoryCLI() {
	categories, err := services.ListCategories()
	if err != nil {
		fmt.Println("Gagal mengambil data kategori:", err)
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tNama\tDeskripsi")
	for _, c := range categories {
		fmt.Fprintf(writer, "%d\t%s\t%s\n", c.ID, c.Name, c.Description)
	}
	writer.Flush()
}

func AddCategoryCLI() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama kategori: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Masukkan deskripsi kategori: ")
	description, _ := reader.ReadString('\n')

	err := services.AddCategory(strings.TrimSpace(name), strings.TrimSpace(description))
	if err != nil {
		fmt.Println("Gagal menambahkan kategori:", err)
	} else {
		fmt.Println("Kategori berhasil ditambahkan!")
	}
}

func DetailCategoryCLI() {
	var input string
	fmt.Print("Masukkan ID kategori: ")
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID tidak valid")
		return
	}

	category, err := services.GetCategoryDetail(id)
	if err != nil {
		fmt.Println("Gagal:", err)
	} else {
		fmt.Println("ID:", category.ID)
		fmt.Println("Nama:", category.Name)
		fmt.Println("Deskripsi:", category.Description)
	}
}

func EditCategoryCLI() {
	var input string
	fmt.Print("Masukkan ID kategori: ")
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID tidak valid")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama baru: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Masukkan deskripsi baru: ")
	description, _ := reader.ReadString('\n')

	err = services.EditCategory(id, strings.TrimSpace(name), strings.TrimSpace(description))
	if err != nil {
		fmt.Println("Gagal mengedit kategori:", err)
	} else {
		fmt.Println("Kategori berhasil diubah.")
	}
}

func DeleteCategoryCLI() {
	var input string
	fmt.Print("Masukkan ID kategori yang ingin dihapus: ")
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID tidak valid")
		return
	}

	err = services.RemoveCategory(id)
	if err != nil {
		fmt.Println("Gagal menghapus kategori:", err)
	} else {
		fmt.Println("Kategori berhasil dihapus.")
	}
}
