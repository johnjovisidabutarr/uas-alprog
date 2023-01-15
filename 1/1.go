package main

import "fmt"

const NMAX = 10

type datanilai struct {
	nim, nama string
	nilai     float64
}

type arrnilai struct {
	arr [NMAX]datanilai
	jml int
}

func main() {
	var arrayA, arrayB, arr2MK arrnilai
	var carinama string
	var hasilCari int

	inputarr(&arrayA)
	inputarr(&arrayB)

	irisanDuaDataNilai(&arr2MK, &arrayA, &arrayB)
	urutData(&arr2MK)

	fmt.Scanln(&carinama)

	hasilCari = binarySearch(&arr2MK, carinama)
	if hasilCari == -1 {
		fmt.Println("Belum mengikuti kedua MK")
	} else {
		fmt.Println(arr2MK.arr[hasilCari].nilai)
	}

}

func inputarr(A *arrnilai) {
	var N int
	var nim, nama string
	var nilai float64
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&nim, &nama, &nilai)
		(*A).arr[i].nim = nim
		(*A).arr[i].nama = nama
		(*A).arr[i].nilai = nilai
		(*A).jml = N
	}
}

func irisanDuaDataNilai(C, A, B *arrnilai) {
	var k int
	var avg float64
	for i := 0; i < (*A).jml; i++ {
		for j := 0; j < (*B).jml; j++ {
			if (*A).arr[i].nim == (*B).arr[j].nim {
				avg = ((*A).arr[i].nilai + (*B).arr[j].nilai) / 2
				(*C).arr[k] = (*A).arr[i]
				(*C).arr[k].nilai = avg
				k++
			}
		}
	}
	(*C).jml = k
}

func urutData(A *arrnilai) {
	for i := 1; i < (*A).jml; i++ {
		for j := i; j > 0; j-- {
			if (*A).arr[j-1].nama < (*A).arr[j].nama {
				(*A).arr[j-1], (*A).arr[j] = (*A).arr[j], (*A).arr[j-1]
			}
		}
	}
}

func binarySearch(A *arrnilai, x string) int {
	left := 0
	right := (*A).jml - 1

	for left <= right {
		mid := (left + right) / 2

		if (*A).arr[mid].nama == x {
			return mid
		} else if (*A).arr[mid].nama < x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
