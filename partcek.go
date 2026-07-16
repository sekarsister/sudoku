for i := 1; i <= 9; i++ {
	for j := 1; j <= 9; j++ {
		sudocek[i][j].totalcandidate = 9
		b1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		if sudo[i][j] == 0 {
			for i2 := 1; i2 <= 9; i2++ {
				b1[sudo[i][i2-1]] = 0
			}
			for j2 := 1; j2 <= 9; j2++ {
				b1[sudo[j2][j-1]] = 0
			}
			for i2 := (i-1)/3*3 + 1; i2 <= (i-1)/3*3+3; i2++ {
				for j2 := (j-1)/3*3 + 1; j2 <= (j-1)/3*3+3; j2++ {
					b1[sudo[i2][j2-1]] = 0
				}
			}
			for k := 0; k < 9; k++ {
				if b1[k] == 0 {
					sudocek[i][j].totalcandidate--
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
for i := 1; i <= 9; i++ {
	for j := 1; j <= 9; j++ {
		if sudocek[i][j].totalcandidate == 9 {
			fmt.Printf("x=%d y=%d totalcandidate=%d\n", i, j, sudocek[i][j].totalcandidate)
		}
	}
}