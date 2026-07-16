# Sudoku Solver — SekarSendiriSaja Development Notes

Proyek ini adalah implementasi solver Sudoku berbasis bahasa Go yang dikembangkan secara bertahap, dari validasi papan hingga pencarian kandidat. Dokumentasi ini mencatat alur pengembangan, aturan domain, dan struktur program yang telah dibangun.

---

## Aturan Sudoku

Papan Sudoku terdiri dari **9 baris × 9 kolom**, membentuk **81 sel** yang dikelompokkan menjadi **9 kotak 3×3**.

**Baris** — setiap baris harus memuat angka 1–9 tepat satu kali, tidak ada pengulangan.

**Kolom** — setiap kolom harus memuat angka 1–9 tepat satu kali.

**Kotak 3×3** — setiap subgrid 3×3 harus memuat angka 1–9 tepat satu kali.

**Clue (angka petunjuk awal)** — sebagian sel sudah terisi sejak awal dan tidak boleh diubah.

**Solusi tunggal** — setiap teka-teki Sudoku yang valid memiliki tepat satu solusi. Penyelesaian dilakukan dengan deduksi logis murni, bukan tebak-tebakan.

Hanya angka 1–9 yang digunakan. Tidak ada angka 0, bilangan negatif, atau huruf.

---

## Struktur File

```
sudostep/
├── tipeawal.go   # Validasi awal papan: cek baris dan kolom
├── bagian2.go    # Pencarian kandidat per sel (part 2)
├── partcek.go    # Potongan logika iterasi kotak 3×3 dan cetak kandidat
├── input.txt     # Contoh data masukan puzzle (format: xyv)
└── Buku1.xlsx    # Catatan pengembangan dan pemetaan konsep
```

---

## Format Input

Program membaca input berupa string 3 karakter dari stdin:

```
xyv
```

- `x` = baris (1–9)
- `y` = kolom (1–9)
- `v` = nilai (1–9)

Input diakhiri dengan token `000` sebagai penanda berhenti.

Contoh:

```
115
123
134
...
000
```

Setiap baris mengisi satu sel pada papan. Sel yang tidak disebutkan dianggap kosong (nilai `0`).

---

## Alur Pengembangan

Checkpoint progres mengikuti jalur **Greedy → Dynamic Programming**:

### Part 1 — `tipeawal.go`

Tahap awal: membangun model papan dan melakukan validasi dasar.

- Papan direpresentasikan sebagai array `[10][10]int` (1-indexed, indeks 0 diabaikan).
- Input dibaca dalam loop hingga `000` ditemukan.
- Papan dicetak dengan pemisah tiap 3 kolom dan 3 baris.
- Validasi dilakukan untuk **baris** dan **kolom**: jika ada angka yang muncul lebih dari sekali, papan dinyatakan tidak valid.

Output: `sodoku valid` atau `sodoku tidak valid`.

### Part 2 — `bagian2.go`

Tahap lanjut: pencarian kandidat untuk setiap sel kosong.

Setelah validasi baris dan kolom berhasil, program melanjutkan ke fase analisis kandidat:

- Setiap sel kosong (`sudo[i][j] == 0`) dianalisis untuk menentukan angka mana saja yang mungkin diisikan.
- Angka dieliminasi berdasarkan tiga aturan sekaligus: baris, kolom, dan kotak 3×3.
- Formula iterasi kotak 3×3:
  ```go
  i2 := (i-1)/3*3 + 1; i2 <= (i-1)/3*3+3; i2++
  j2 := (j-1)/3*3 + 1; j2 <= (j-1)/3*3+3; j2++
  ```
- Sel yang sudah terisi (`sudo[i][j] != 0`) diberi `totalcandidate = 10` sebagai penanda "bukan kandidat aktif".
- Kandidat yang tersisa dicetak ke layar beserta koordinatnya.

Struct yang digunakan:

```go
type candidate struct {
    a1, a2, a3, a4, a5, a6, a7, a8, a9 int
    totalcandidate int
}
var sudocek [10][10]candidate
```

---

## Cara Menjalankan

### Dari terminal

```bash
go run bagian2.go < input.txt
```

Atau dengan input manual:

```bash
go run tipeawal.go
```

Ketik koordinat dan nilai satu per baris, akhiri dengan `000`.

---

## Catatan Teknis

- **Indeks papan** dimulai dari 1, bukan 0. Baris dan kolom ke-0 sengaja tidak dipakai agar indexing lebih natural.
- **Eliminasi kandidat** menggunakan array sementara `b1 [10]int` yang diisi angka 1–9, lalu posisi yang sudah terpakai di baris/kolom/kotak di-nol-kan. Angka yang tersisa adalah kandidat valid.
- **Validasi box** tidak menggunakan fungsi tersendiri karena keterbatasan kondisi yang sulit dirangkum ke dalam satu `if != 45`. Validasi baris dan kolom bisa menggunakan pendekatan `count unique value`.
- Komentar di akhir file mencatat rencana pengembangan berikutnya:
  ```
  // cari kasus 1 kotakan kondisi box kolom dan baris unique
  // dari kasus ini kondisi masih kosong dan belum dimainkan
  // 123456789 mod 10 for collect && div by 10 to eliminate
  // no function
  ```

---

## Rencana Pengembangan Selanjutnya

- Melengkapi pencatatan kandidat di struct `candidate` (field `a1`–`a9`).
- Implementasi strategi **naked single**: jika suatu sel hanya punya satu kandidat, langsung isi.
- Implementasi strategi **hidden single**: jika suatu angka hanya mungkin di satu posisi dalam baris/kolom/kotak.
- Transisi dari pendekatan greedy ke **dynamic programming** untuk kasus puzzle yang lebih kompleks.
- Validasi kotak 3×3 secara eksplisit (saat ini hanya baris dan kolom yang divalidasi di `tipeawal.go`).
