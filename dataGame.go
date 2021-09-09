package main

import (
	"fmt"
	"sort"
	"strconv"
)

var jenis string
var rating string
var kualitas string
var name string
var arr1 []string
var arr [][]string
var idgame int

func addData() {
	fmt.Print("\nMasukkan Judul Game\t\t: ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan Ratting (0.0 - 5.0)\t: ")
	fmt.Scan(&rating)

	arr1 = []string{jenis, rating}
	arr = append(arr, arr1)

	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

}

func deleteData(ID int) {
	var before, after [][]string
	before = arr[:ID-1]
	after = arr[ID:]
	arr = append(after, before...)

}

func main() {
	var pilih int
	var x = true

	for x {
		fmt.Println("\n\t====Daftar Pilihan==== ")
		fmt.Println("\n1. Input Data game baru \n2. Hapus data game berdasarkan id_game \n3. Tampilkan seluruh data game beserta jumlah data yang tersimpan dalam list")
		fmt.Println("4. Cari data game berdasarkan nama \n5. Menampilkan top 3 game terfavorit \n6. Menampilkan seluruh data dengan rating diatas 4.0 \n0. Keluar")
		fmt.Print("Masukkan Pilihan\t\t: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			addData()

		case 2:
			fmt.Print("\nMasukkan ID game yang akan dihapus: ")
			fmt.Scan(&idgame)

			fmt.Printf("Nama Game: %s \t Rating: %s \nBERHASIL TERHAPUS", arr[idgame-1][0], arr[idgame-1][1])

			deleteData(idgame)
			sort.Slice(arr[:], func(i, j int) bool {
				return arr[i][1] > arr[j][1]
			})
		case 3:
			fmt.Println("ID_GAME \tNama Game \t\tRating \t\tKualitas")
			for i := 0; i < len(arr); i++ {
				var fltrating, err = strconv.ParseFloat(arr[i][1], 64)
				if err == nil {
					if fltrating >= 4.0 {
						kualitas = "good"
					} else if fltrating >= 2.0 {
						kualitas = "average"
					} else if fltrating < 2.0 {
						kualitas = "poor"
					}
				}
				fmt.Printf("%d \t\t%s \t\t\t%s \t\t%s\n", i+1, arr[i][0], arr[i][1], kualitas)
			}
			fmt.Printf("\njumlah \t= %d", len(arr))
		case 4:
			fmt.Print("Masukkan nama game: ")
			fmt.Scan(&name)

			data := 0
			for i := 0; i < len(arr); i++ {
				if name == arr[i][0] {
					fmt.Printf("ID Game = %d \tNama game = %s \trating = %s\n", i+1, arr[i][0], arr[i][1])
				} else {
					data = data + 1
				}
			}
			if data == len(arr) {
				fmt.Println("Data tidak ditemukan!")
			}
		case 5:
			fmt.Println("TOP 3 GAMES")
			fmt.Println("Rank \tNama Game \t\tRating")
			for i := 0; i <= 2; i++ {
				fmt.Printf("%d \t%s \t\t%s\n", i+1, arr[i][0], arr[i][1])
			}

		case 6:
			fmt.Println("ID_GAME \tNama Game \t\tRating \t\tKualitas")
			data := 0
			for i := 0; i < len(arr); i++ {
				var fltrating, err = strconv.ParseFloat(arr[i][1], 64)
				if err == nil {
					if fltrating >= 4.0 {
						if fltrating >= 4.0 {
							kualitas = "good"
						} else if fltrating >= 2.0 {
							kualitas = "average"
						} else if fltrating < 2.0 {
							kualitas = "poor"
						}
						fmt.Printf("%d \t\t%s \t\t\t%s \t\t%s\n", i+1, arr[i][0], arr[i][1], kualitas)
					} else {
						data = data + 1
					}
				}
			}
			if data == len(arr) {
				fmt.Println("Data game dengan rating diatas 4.0 tidak ditemukan!")
			}
		case 0:
			x = false
		}
	}

}
