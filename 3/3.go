package main

import "fmt"

const NMAX = 10

type siswa struct {
	nama             string
	mtk, bahasa, ipa int
}

type kelas [NMAX]siswa

func main() {
	var data kelas
	var N, i int
	var nama string

	input_data(&data, &N)
	sort_data(&data, &N)
	fmt.Scanln(&nama)
	i = search_data(&data, &N, nama)
	if i != -1 {
		fmt.Println("Nilai rata-rata dari", data[i].nama, "adalah", rata_rata(&data[i]))

	}
	print_data(&data, &N)
	fmt.Println("Ok")
}

func input_data(T *kelas, N *int) {
	var nama string
	var mtk, bahasa, ipa int
	for i := 0; i < NMAX; i++ {
		fmt.Scanln(&nama, &mtk, &bahasa, &ipa)
		if nama == "stop" {
			break
		}
		(*T)[i].nama = nama
		(*T)[i].mtk = mtk
		(*T)[i].bahasa = bahasa
		(*T)[i].ipa = ipa
		*N = i + 1
	}
}

func rata_rata(M *siswa) float64 {
	var avg float64
	for i := 0; i < NMAX; i++ {
		avg = float64((*M).mtk+(*M).bahasa+(*M).ipa) / 3
	}
	return avg
}

func search_data(T *kelas, N *int, nama string) int {
	var i int
	for i = 0; i < *N; i++ {
		if (*T)[i].nama == nama {
			return i
		}
	}
	return -1
}

func sort_data(T *kelas, N *int) {
	for i := 0; i < *N; i++ {
		for j := 0; j < *N; j++ {
			if float64((*T)[i].mtk+(*T)[i].bahasa+(*T)[i].ipa) < float64((*T)[j].mtk+(*T)[j].bahasa+(*T)[j].ipa) {
				(*T)[j], (*T)[i] = (*T)[i], (*T)[j]
			}
		}
	}
}

func print_data(T *kelas, N *int) {
	var j int = 1
	for i := *N - 1; i > *N-4; i-- {
		fmt.Println((*T)[i].nama, "rangking", j, "dengan rata-rata", float64((*T)[i].mtk+(*T)[i].bahasa+(*T)[i].ipa)/3)
		j++
	}
}
