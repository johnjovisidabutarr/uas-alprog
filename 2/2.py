class Mahasiswa:
    def __init__(self, nim, nama, kuis, uts, uas):
        self.nim = nim
        self.nama = nama
        self.kuis = kuis
        self.uts = uts
        self.uas = uas
        self.nilai_akhir = 0.0

def input_data(arr_mhs, N):
    for i in range(N):
        nama, nim, kuis, uts, uas = input("Enter student's name, nim, kuis, uts, and uas scores separated by spaces: ").split()
        arr_mhs[i] = Mahasiswa(nim, nama, int(kuis), int(uts), int(uas))

def hitung_nilai_akhir(arr_mhs, N):
    for i in range(N):
        nilai_kuis = 0.2 * arr_mhs[i].kuis
        nilai_uts = 0.4 * arr_mhs[i].uts
        nilai_uas = 0.4 * arr_mhs[i].uas
        arr_mhs[i].nilai_akhir = nilai_kuis + nilai_uts + nilai_uas

def min_max(arr_mhs, N):
    min_mhs = max_mhs = 0
    for i in range(1, N):
        if arr_mhs[i].nilai_akhir < arr_mhs[min_mhs].nilai_akhir:
            min_mhs = i
        elif arr_mhs[i].nilai_akhir > arr_mhs[max_mhs].nilai_akhir:
            max_mhs = i
    return min_mhs, max_mhs

def median(arr_mhs, N):
    arr_mhs.sort(key=lambda x: x.nilai_akhir)
    median = 0.0
    if N == 1:
        median = arr_mhs[0].nilai_akhir
    elif N % 2 != 0:
        median = arr_mhs[N // 2].nilai_akhir
    else:
        median = (arr_mhs[N // 2].nilai_akhir + arr_mhs[N // 2 - 1].nilai_akhir) / 2
    return median

def mencetak(mhs):
    print(mhs.nama, mhs.nim)

if __name__ == "__main__":
    N = int(input("Enter number of students: "))
    arr_mhs = [Mahasiswa("", "", 0, 0, 0) for i in range(N)]
    input_data(arr_mhs, N)
    hitung_nilai_akhir(arr_mhs, N)
    min_mhs, max_mhs = min_max(arr_mhs, N)
    print("Mahasiswa yang mendapat nilai maksimum : ")
    mencetak(arr_mhs[max_mhs])
    print("Mahasiswa yang mendapat nilai minimum : ")
    mencetak(arr_mhs[min_mhs])
    nilai_median = median(arr_mhs, N)
    print("Median : ", nilai_median)