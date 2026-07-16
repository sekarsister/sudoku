package main

import "fmt"

func main() {
	var janggal bool = false
	var count [10]int
	var sudo = [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var a string
	for true {
		fmt.Scan(&a)
		if a == "000" {
			break
		}
		sudo[(a[0] - 48)][(a[1] - 48)] = int(a[2] - 48)
		fmt.Println(a[0], a[1]-48, a[2]-48)
		fmt.Printf("x=%c y=%c value=%c\n", a[0], a[1], a[2])
	}
	for i := 1; i <= 9; i++ {
		fmt.Print("\n")

		for j := 1; j <= 9; j++ {
			fmt.Printf("%d ", sudo[i][j])
			if j%3 == 0 {
				fmt.Print(" ")
			}
		}
		if i%3 == 0 {
			fmt.Print("\n")
		}
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			count[sudo[i][j]]++
		}
		for k := 1; k <= 9; k++ {
			if count[k] > 1 {
				janggal = true
			}
		}
		count = [10]int{}
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			count[sudo[j][i]]++
		}
		for k := 1; k <= 9; k++ {
			if count[k] > 1 {
				janggal = true
			}
		}
		count = [10]int{}
	}
	if janggal {
		fmt.Println("sodoku tidak valid")
	} else {
		fmt.Println("sodoku valid")
	}
}

// for y => x to x++
// cari kasus 1 kotakan kondisi box kolom dan baris unique dari kasus ini kondisi masih kosong dan belum di dimainkan (mencari anomali)
// chalange box tidak bisa menggunakan func untuk if !=45
// untuk baris dan kolom kita bisa menggunakan count unique value
