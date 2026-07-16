package main

import "fmt"

func main() {
	var janggal bool = false
	var count [10]int
	type candidate struct {
		a1             int
		a2             int
		a3             int
		a4             int
		a5             int
		a6             int
		a7             int
		a8             int
		a9             int
		totalcandidate int
	}
	var sudocek [10][10]candidate

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
		for i := 1; i <= 9; i++ {
			for j := 1; j <= 9; j++ {
				sudocek[i][j].totalcandidate = 9
				if sudo[i][j] == 0 {
					b1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
					for i2 := 1; i2 <= 9; i2++ {
						b1[sudo[i][i2]] = 0
					}
					for j2 := 1; j2 <= 9; j2++ {
						b1[sudo[j2][j]] = 0
					}
					for i2 := (i-1)/3*3 + 1; i2 <= (i-1)/3*3+3; i2++ {
						for j2 := (j-1)/3*3 + 1; j2 <= (j-1)/3*3+3; j2++ {
							b1[sudo[i2][j2]] = 0
						}
					}
					for k := 0; k < 10; k++ {
						if b1[k] == 0 {
							sudocek[i][j].totalcandidate--
						} else {
							fmt.Printf("x=%d y=%d candidate=%d\n", i, j, b1[k])
						}
					}
				} else {
					sudocek[i][j].totalcandidate = 10
					sudocek[i][j].a1 = 0
					sudocek[i][j].a2 = 0
					sudocek[i][j].a3 = 0
					sudocek[i][j].a4 = 0
					sudocek[i][j].a5 = 0
					sudocek[i][j].a6 = 0
					sudocek[i][j].a7 = 0
					sudocek[i][j].a8 = 0
					sudocek[i][j].a9 = 0
				}
			}
		}
	}
}

// for y => x to x++
// cari kasus 1 kotakan kondisi box kolom dan baris unique dari kasus ini kondisi masih kosong dan belum di dimainkan (mencari anomali)
// chalange box tidak bisa menggunakan func untuk if !=45
// untuk baris dan kolom kita bisa menggunakan count unique value
// 123456789 mod 10 for collect && div by 10 to elimitate
// no function
