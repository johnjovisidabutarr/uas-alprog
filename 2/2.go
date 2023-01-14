package main

import "fmt"

type mahasiswa struct {
	nim, nama      string
	kuis, uts, uas int
	nilai_akhir    float64
}

type ArrMhs [100]mahasiswa

func inputData(T *ArrMhs, N *int) {
	var nama, nim string
	var kuis, uts, uas int
	for i := 0; i < *N; i++ {
		fmt.Scanln(&nama, &nim, &kuis, &uts, &uas)
		(*T)[i].nama = nama
		(*T)[i].nim = nim
		(*T)[i].kuis = kuis
		(*T)[i].uts = uts
		(*T)[i].uas = uas
	}
}

func hitung_nilai_akhir(T *ArrMhs, N *int) {
	var nilaiKuis, nilaiUTS, nilaiUAS float64
	for i := 0; i < *N; i++ {
		nilaiKuis = 0.2 * float64((*T)[i].kuis)
		nilaiUTS = 0.4 * float64((*T)[i].uts)
		nilaiUAS = 0.4 * float64((*T)[i].uas)
		(*T)[i].nilai_akhir = nilaiKuis + nilaiUTS + nilaiUAS
	}
}

func min_max(T *ArrMhs, N int, min, max *int) {
	for i := 0; i < N; i++ {
		if (*T)[i].nilai_akhir < (*T)[*min].nilai_akhir {
			*min = i
		} else if (*T)[i].nilai_akhir > (*T)[*max].nilai_akhir {
			*max = i
		}
	}
}

func median(T *ArrMhs, N int) float64 {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if (*T)[j].nilai_akhir > (*T)[i].nilai_akhir {
				(*T)[j], (*T)[i] = (*T)[i], (*T)[j]
			}
		}
	}

	var median float64

	if N == 1 {
		median = (*T)[0].nilai_akhir
		return median
	}

	if N%2 != 0 {
		median = (*T)[(N/2)-1].nilai_akhir
	} else if N%2 == 0 {
		median = ((*T)[(N/2)-1].nilai_akhir + (*T)[(N/2)].nilai_akhir) * 0.5

	}
	return median
}

func mencetak(T *ArrMhs, N *int) {
	fmt.Println((*T)[*N].nama, (*T)[*N].nim)
}

func main() {
	var min, max int
	var T ArrMhs
	var N int

	fmt.Scanln(&N)

	inputData(&T, &N)

	hitung_nilai_akhir(&T, &N)

	min_max(&T, N, &min, &max)
	fmt.Printf("Mahasiswa yang mendapat nilai maksimum : ")
	mencetak(&T, &max)
	fmt.Printf("Mahasiswa yang mendapat nilai minimum : ")
	mencetak(&T, &min)

	nilai_median := median(&T, N)
	fmt.Println("Median : ", nilai_median)
}
