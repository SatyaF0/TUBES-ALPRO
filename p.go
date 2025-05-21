package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Mahasiswa struct {
	NIM    string
	Name   string
	Course string
	Origin string
}

type Buku struct {
	Author              string
	Judul               string
	Tahun               int
	TanggalPeminjaman   string
	TanggalPengembalian string
}

type Pegawai struct {
	NIP     string
	Nama    string
	Jabatan string
}

var studentData []Mahasiswa
var bookData []Buku
var employeeData []Pegawai

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Data")
		fmt.Println("2. View Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Delete Data")
		fmt.Println("5. Search Data")
		fmt.Println("6. Exit")
		fmt.Print("Choose menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addData(scanner)
		case "2":
			viewData()
		case "3":
			updateData(scanner)
		case "4":
			deleteData(scanner)
		case "5":
			searchData(scanner)
		case "6":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addData(scanner *bufio.Scanner) {
	for {
		fmt.Println("Add Data for:")
		fmt.Println("1. Mahasiswa")
		fmt.Println("2. Buku")
		fmt.Println("3. Pegawai")
		fmt.Println("4. Back")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			var m Mahasiswa
			for {
				fmt.Print("NIM: ")
				scanner.Scan()
				m.NIM = scanner.Text()
				if isNumeric(m.NIM) {
					break
				}
				fmt.Println("Error: NIM must be numeric.")
			}
			fmt.Print("Name: ")
			scanner.Scan()
			m.Name = scanner.Text()
			fmt.Print("Course: ")
			scanner.Scan()
			m.Course = scanner.Text()
			fmt.Print("Place of Origin (Country): ")
			scanner.Scan()
			m.Origin = scanner.Text()

			studentData = append(studentData, m)
			fmt.Println("Mahasiswa data successfully added.")
		case "2":
			var b Buku
			fmt.Print("Author: ")
			scanner.Scan()
			b.Author = scanner.Text()
			fmt.Print("Judul: ")
			scanner.Scan()
			b.Judul = scanner.Text()
			fmt.Print("Tahun: ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &b.Tahun)
			fmt.Print("Tanggal Peminjaman: ")
			scanner.Scan()
			b.TanggalPeminjaman = scanner.Text()
			fmt.Print("Tanggal Pengembalian: ")
			scanner.Scan()
			b.TanggalPengembalian = scanner.Text()

			bookData = append(bookData, b)
			fmt.Println("Buku data successfully added.")
		case "3":
			var p Pegawai
			fmt.Print("NIP: ")
			scanner.Scan()
			p.NIP = scanner.Text()
			fmt.Print("Nama: ")
			scanner.Scan()
			p.Nama = scanner.Text()
			fmt.Print("Jabatan: ")
			scanner.Scan()
			p.Jabatan = scanner.Text()

			employeeData = append(employeeData, p)
			fmt.Println("Pegawai data successfully added.")
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func isNumeric(input string) bool {
	match, _ := regexp.MatchString(`^\d+$`, input)
	return match
}

func viewData() {
	for {
		fmt.Println("View Data for:")
		fmt.Println("1. Mahasiswa")
		fmt.Println("2. Buku")
		fmt.Println("3. Pegawai")
		fmt.Println("4. Back")
		fmt.Print("Choose: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			if len(studentData) == 0 {
				fmt.Println("No Mahasiswa data available.")
				continue
			}
			for {
				fmt.Println("Sort by:")
				fmt.Println("1. Ascending (Selection Sort)")
				fmt.Println("2. Descending (Insertion Sort)")
				fmt.Println("3. Back")
				fmt.Print("Choose: ")
				scanner.Scan()
				sortChoice := scanner.Text()
				if sortChoice == "3" {
					break
				} else if sortChoice == "1" {
					selectionSortMahasiswa(studentData, true)
				} else if sortChoice == "2" {
					insertionSortMahasiswa(studentData, false)
				} else {
					fmt.Println("Invalid sort choice.")
					continue
				}
				for i, m := range studentData {
					fmt.Printf("%d. NIM: %s, Name: %s, Course: %s, Origin: %s\n",
						i+1, m.NIM, m.Name, m.Course, m.Origin)
				}
				break
			}
		case "2":
			if len(bookData) == 0 {
				fmt.Println("No Buku data available.")
				continue
			}
			for {
				fmt.Println("Sort by:")
				fmt.Println("1. Ascending (Selection Sort)")
				fmt.Println("2. Descending (Insertion Sort)")
				fmt.Println("3. Back")
				fmt.Print("Choose: ")
				scanner.Scan()
				sortChoice := scanner.Text()
				if sortChoice == "3" {
					break
				} else if sortChoice == "1" {
					selectionSortBuku(bookData, true)
				} else if sortChoice == "2" {
					insertionSortBuku(bookData, false)
				} else {
					fmt.Println("Invalid sort choice.")
					continue
				}
				for i, b := range bookData {
					fmt.Printf("%d. Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
						i+1, b.Author, b.Judul, b.Tahun, b.TanggalPeminjaman, b.TanggalPengembalian)
				}
				break
			}
		case "3":
			if len(employeeData) == 0 {
				fmt.Println("No Pegawai data available.")
				continue
			}
			for {
				fmt.Println("Sort by:")
				fmt.Println("1. Ascending (Selection Sort)")
				fmt.Println("2. Descending (Insertion Sort)")
				fmt.Println("3. Back")
				fmt.Print("Choose: ")
				scanner.Scan()
				sortChoice := scanner.Text()
				if sortChoice == "3" {
					break
				} else if sortChoice == "1" {
					selectionSortPegawai(employeeData, true)
				} else if sortChoice == "2" {
					insertionSortPegawai(employeeData, false)
				} else {
					fmt.Println("Invalid sort choice.")
					continue
				}
				for i, p := range employeeData {
					fmt.Printf("%d. NIP: %s, Nama: %s, Jabatan: %s\n",
						i+1, p.NIP, p.Nama, p.Jabatan)
				}
				break
			}
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func selectionSortMahasiswa(data []Mahasiswa, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		idx := i
		for j := i + 1; j < len(data); j++ {
			if ascending && data[j].NIM < data[idx].NIM {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortMahasiswa(data []Mahasiswa, descending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && descending && data[j].NIM < key.NIM {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func selectionSortBuku(data []Buku, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		idx := i
		for j := i + 1; j < len(data); j++ {
			if ascending && data[j].Judul < data[idx].Judul {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortBuku(data []Buku, descending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && descending && data[j].Judul < key.Judul {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func selectionSortPegawai(data []Pegawai, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		idx := i
		for j := i + 1; j < len(data); j++ {
			if ascending && data[j].NIP < data[idx].NIP {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortPegawai(data []Pegawai, descending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && descending && data[j].NIP < key.NIP {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func updateData(scanner *bufio.Scanner) {
	for {
		fmt.Println("Update Data for:")
		fmt.Println("1. Mahasiswa")
		fmt.Println("2. Buku")
		fmt.Println("3. Pegawai")
		fmt.Println("4. Back")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter NIM of the Mahasiswa to update: ")
			scanner.Scan()
			nim := scanner.Text()
			for i, m := range studentData {
				if m.NIM == nim {
					fmt.Println("Which one do you want to update?")
					fmt.Println("1. Name")
					fmt.Println("2. NIM")
					fmt.Println("3. Course")
					fmt.Println("4. Origin")
					fmt.Print("Choose: ")
					scanner.Scan()
					field := scanner.Text()
					switch field {
					case "1":
						fmt.Print("Enter new Name: ")
						scanner.Scan()
						studentData[i].Name = scanner.Text()
					case "2":
						fmt.Print("Enter new NIM: ")
						scanner.Scan()
						studentData[i].NIM = scanner.Text()
					case "3":
						fmt.Print("Enter new Course: ")
						scanner.Scan()
						studentData[i].Course = scanner.Text()
					case "4":
						fmt.Print("Enter new Origin: ")
						scanner.Scan()
						studentData[i].Origin = scanner.Text()
					default:
						fmt.Println("Invalid choice.")
					}
					fmt.Println("Mahasiswa data successfully updated.")
					return
				}
			}
			fmt.Println("Mahasiswa with the specified NIM not found.")
		case "2":
			fmt.Print("Enter Judul of the Buku to update: ")
			scanner.Scan()
			judul := scanner.Text()
			for i, b := range bookData {
				if strings.ToLower(b.Judul) == strings.ToLower(judul) {
					fmt.Println("Which one do you want to update?")
					fmt.Println("1. Author")
					fmt.Println("2. Judul")
					fmt.Println("3. Tahun")
					fmt.Println("4. Tanggal Peminjaman")
					fmt.Println("5. Tanggal Pengembalian")
					fmt.Print("Choose: ")
					scanner.Scan()
					field := scanner.Text()
					switch field {
					case "1":
						fmt.Print("Enter new Author: ")
						scanner.Scan()
						bookData[i].Author = scanner.Text()
					case "2":
						fmt.Print("Enter new Judul: ")
						scanner.Scan()
						bookData[i].Judul = scanner.Text()
					case "3":
						fmt.Print("Enter new Tahun: ")
						scanner.Scan()
						fmt.Sscanf(scanner.Text(), "%d", &bookData[i].Tahun)
					case "4":
						fmt.Print("Enter new Tanggal Peminjaman: ")
						scanner.Scan()
						bookData[i].TanggalPeminjaman = scanner.Text()
					case "5":
						fmt.Print("Enter new Tanggal Pengembalian: ")
						scanner.Scan()
						bookData[i].TanggalPengembalian = scanner.Text()
					default:
						fmt.Println("Invalid choice.")
					}
					fmt.Println("Buku data successfully updated.")
					return
				}
			}
			fmt.Println("Buku with the specified Judul not found.")
		case "3":
			fmt.Print("Enter NIP of the Pegawai to update: ")
			scanner.Scan()
			nip := scanner.Text()
			for i, p := range employeeData {
				if p.NIP == nip {
					fmt.Println("Which one do you want to update?")
					fmt.Println("1. Nama")
					fmt.Println("2. NIP")
					fmt.Println("3. Jabatan")
					fmt.Print("Choose: ")
					scanner.Scan()
					field := scanner.Text()
					switch field {
					case "1":
						fmt.Print("Enter new Nama: ")
						scanner.Scan()
						employeeData[i].Nama = scanner.Text()
					case "2":
						fmt.Print("Enter new NIP: ")
						scanner.Scan()
						employeeData[i].NIP = scanner.Text()
					case "3":
						fmt.Print("Enter new Jabatan: ")
						scanner.Scan()
						employeeData[i].Jabatan = scanner.Text()
					default:
						fmt.Println("Invalid choice.")
					}
					fmt.Println("Pegawai data successfully updated.")
					return
				}
			}
			fmt.Println("Pegawai with the specified NIP not found.")
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func deleteData(scanner *bufio.Scanner) {
	for {
		fmt.Println("Delete Data for:")
		fmt.Println("1. Mahasiswa")
		fmt.Println("2. Buku")
		fmt.Println("3. Pegawai")
		fmt.Println("4. Back")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter NIM of the Mahasiswa to delete: ")
			scanner.Scan()
			nim := scanner.Text()
			for i, m := range studentData {
				if m.NIM == nim {
					fmt.Printf("NIM: %s, Name: %s, Course: %s, Origin: %s\n", m.NIM, m.Name, m.Course, m.Origin)
					fmt.Print("Delete this data? (yes/no): ")
					scanner.Scan()
					confirm := strings.ToLower(scanner.Text())
					if confirm == "yes" {
						studentData = append(studentData[:i], studentData[i+1:]...)
						fmt.Println("Mahasiswa data successfully deleted.")
					} else {
						fmt.Println("Deletion canceled.")
					}
					return
				}
			}
			fmt.Println("Mahasiswa with the specified NIM not found.")
		case "2":
			fmt.Print("Enter Judul of the Buku to delete: ")
			scanner.Scan()
			judul := scanner.Text()
			for i, b := range bookData {
				if strings.ToLower(b.Judul) == strings.ToLower(judul) {
					fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
						b.Author, b.Judul, b.Tahun, b.TanggalPeminjaman, b.TanggalPengembalian)
					fmt.Print("Delete this data? (yes/no): ")
					scanner.Scan()
					confirm := strings.ToLower(scanner.Text())
					if confirm == "yes" {
						bookData = append(bookData[:i], bookData[i+1:]...)
						fmt.Println("Book data successfully deleted.")
					} else {
						fmt.Println("Deletion canceled.")
					}
					return
				}
			}
			fmt.Println("Buku with the specified Judul not found.")
		case "3":
			fmt.Print("Enter NIP of the Pegawai to delete: ")
			scanner.Scan()
			nip := scanner.Text()
			for i, p := range employeeData {
				if p.NIP == nip {
					fmt.Printf("NIP: %s, Nama: %s, Jabatan: %s\n", p.NIP, p.Nama, p.Jabatan)
					fmt.Print("Delete this data? (yes/no): ")
					scanner.Scan()
					confirm := strings.ToLower(scanner.Text())
					if confirm == "yes" {
						employeeData = append(employeeData[:i], employeeData[i+1:]...)
						fmt.Println("Pegawai data successfully deleted.")
					} else {
						fmt.Println("Deletion canceled.")
					}
					return
				}
			}
			fmt.Println("Pegawai with the specified NIP not found.")
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func searchData(scanner *bufio.Scanner) {
	for {
		fmt.Println("Search Buku:")
		fmt.Println("1. Buku yang tersedia")
		fmt.Println("2. Buku yang dipinjam")
		fmt.Println("3. Cari keanggotaan perpustakaan")
		fmt.Println("4. Back")
		fmt.Print("Choose: ")
		scanner.Scan()
		bukuChoice := scanner.Text()
		switch bukuChoice {
		case "1":
			availableBooks := []Buku{}
			for _, b := range bookData {
				if b.TanggalPeminjaman == "" {
					availableBooks = append(availableBooks, b)
				}
			}
			if len(availableBooks) == 0 {
				fmt.Println("No available books.")
				break
			}
			fmt.Println("Search method:")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			fmt.Print("Choose: ")
			scanner.Scan()
			method := scanner.Text()
			fmt.Print("Enter keyword for Buku: ")
			scanner.Scan()
			keyword := strings.ToLower(scanner.Text())
			if method == "1" {
				sequentialSearchBukuFiltered(availableBooks, keyword)
			} else if method == "2" {
				selectionSortBuku(availableBooks, true)
				binarySearchBukuFiltered(availableBooks, keyword)
			} else {
				fmt.Println("Invalid search method.")
			}
		case "2":
			borrowedBooks := []Buku{}
			for _, b := range bookData {
				if b.TanggalPeminjaman != "" {
					borrowedBooks = append(borrowedBooks, b)
				}
			}
			if len(borrowedBooks) == 0 {
				fmt.Println("No borrowed books.")
				break
			}
			fmt.Println("Search method:")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			fmt.Print("Choose: ")
			scanner.Scan()
			method := scanner.Text()
			fmt.Print("Enter keyword for Buku: ")
			scanner.Scan()
			keyword := strings.ToLower(scanner.Text())
			if method == "1" {
				sequentialSearchBukuFiltered(borrowedBooks, keyword)
			} else if method == "2" {
				selectionSortBuku(borrowedBooks, true)
				binarySearchBukuFiltered(borrowedBooks, keyword)
			} else {
				fmt.Println("Invalid search method.")
			}
		case "3":
			if len(studentData) == 0 {
				fmt.Println("No Mahasiswa data available.")
				break
			}
			fmt.Print("Enter NIM or Name for Mahasiswa: ")
			scanner.Scan()
			keyword := strings.ToLower(scanner.Text())
			found := false
			for _, m := range studentData {
				if strings.Contains(strings.ToLower(m.NIM), keyword) ||
					strings.Contains(strings.ToLower(m.Name), keyword) {
					printStudent(m)
					found = true
				}
			}
			if !found {
				fmt.Println("No Mahasiswa found as library member.")
			}
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func sequentialSearchMahasiswa(keyword string) {
	found := false
	for _, m := range studentData {
		if strings.Contains(strings.ToLower(m.NIM), keyword) ||
			strings.Contains(strings.ToLower(m.Name), keyword) ||
			strings.Contains(strings.ToLower(m.Course), keyword) ||
			strings.Contains(strings.ToLower(m.Origin), keyword) {
			printStudent(m)
			found = true
		}
	}
	if !found {
		fmt.Println("No Mahasiswa data found.")
	}
}

func binarySearchMahasiswa(keyword string) {
	low, high := 0, len(studentData)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(studentData[mid].NIM), keyword) {
			printStudent(studentData[mid])
			return
		} else if keyword < strings.ToLower(studentData[mid].NIM) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("No Mahasiswa data found.")
}

func sequentialSearchBuku(keyword string) {
	found := false
	for _, b := range bookData {
		if strings.Contains(strings.ToLower(b.Author), keyword) ||
			strings.Contains(strings.ToLower(b.Judul), keyword) {
			fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
				b.Author, b.Judul, b.Tahun, b.TanggalPeminjaman, b.TanggalPengembalian)
			found = true
		}
	}
	if !found {
		fmt.Println("No Buku data found.")
	}
}

func binarySearchBuku(keyword string) {
	low, high := 0, len(bookData)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(bookData[mid].Judul), keyword) {
			fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
				bookData[mid].Author, bookData[mid].Judul, bookData[mid].Tahun, bookData[mid].TanggalPeminjaman, bookData[mid].TanggalPengembalian)
			return
		} else if keyword < strings.ToLower(bookData[mid].Judul) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("No Buku data found.")
}

func sequentialSearchPegawai(keyword string) {
	found := false
	for _, p := range employeeData {
		if strings.Contains(strings.ToLower(p.NIP), keyword) ||
			strings.Contains(strings.ToLower(p.Nama), keyword) ||
			strings.Contains(strings.ToLower(p.Jabatan), keyword) {
			fmt.Printf("NIP: %s, Nama: %s, Jabatan: %s\n", p.NIP, p.Nama, p.Jabatan)
			found = true
		}
	}
	if !found {
		fmt.Println("No Pegawai data found.")
	}
}

func binarySearchPegawai(keyword string) {
	low, high := 0, len(employeeData)-1
	found := false
	
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(employeeData[mid].NIP), keyword) {
			fmt.Printf("NIP: %s, Nama: %s, Jabatan: %s\n", employeeData[mid].NIP, employeeData[mid].Nama, employeeData[mid].Jabatan)
			found = true
			break
		} else if keyword < strings.ToLower(employeeData[mid].NIP) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	
	if !found {
		fmt.Println("No Pegawai data found.")
	}
}

func printStudent(m Mahasiswa) {
	fmt.Printf("NIM: %s, Name: %s, Course: %s, Origin: %s\n",
		m.NIM, m.Name, m.Course, m.Origin)
}

// Tambahan fungsi pencarian filter Buku
func sequentialSearchBukuFiltered(data []Buku, keyword string) {
	found := false
	for _, b := range data {
		if strings.Contains(strings.ToLower(b.Author), keyword) ||
			strings.Contains(strings.ToLower(b.Judul), keyword) {
			fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
				b.Author, b.Judul, b.Tahun, b.TanggalPeminjaman, b.TanggalPengembalian)
			found = true
		}
	}
	if !found {
		fmt.Println("No Buku data found.")
	}
}

func binarySearchBukuFiltered(data []Buku, keyword string) {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(data[mid].Judul), keyword) {
			fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
				data[mid].Author, data[mid].Judul, data[mid].Tahun, data[mid].TanggalPeminjaman, data[mid].TanggalPengembalian)
			return
		} else if keyword < strings.ToLower(data[mid].Judul) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("No Buku data found.")
}
