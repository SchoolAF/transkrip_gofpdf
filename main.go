package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	// Membuat objek PDF baru
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Menambahkan judul
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "TRANSKRIP AKADEMIK", "0", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "i", 10)
	pdf.CellFormat(190, 0, "Academic Transcript", "0", 1, "C", false, 0, "")
	pdf.Ln(7.5)

	pdf.SetFont("Arial", "", 15)
	pdf.CellFormat(190, 0, "No: 0000/ULBI/xxxx/yyyy", "0", 1, "C", false, 0, "")
	pdf.Ln(10)

	// Menambahkan data mahasiswa
	mahasiswaData := []struct {
		Key   string
		Value string
	}{
		{"Nama", "Aliffathur Muhammad Revan"},
		{"NIM", "714220066"},
		{"Tempat, Tgl. Lahir", "3 November 2003"},
		{"Tahun Masuk", "2022"},
		{"Fakultas", "Vokasi"},
		{"Program Studi", "Tenik Informatika"},
	}

	// Menambahkan data mahasiswa dengan translasi bahasa Inggris di bawahnya
	for _, data := range mahasiswaData {
		key, value := data.Key, data.Value
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(70, 9, key+":", "0", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 9, value, "0", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 0, translateToEnglish(key), "0", 1, "L", false, 0, "")
		pdf.Ln(0.5)
	}

	pdf.Ln(10)

	// Menambahkan tabel nilai
	headers := []string{"No", "Mata Kuliah", "SKS", "Nilai"}
	data := [][]string{
		{"1", "Mathematics", "3", "A"},
		{"2", "Physics", "3", "B"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
		{"3", "Chemistry", "3", "A"},
	}

	// Menghitung lebar kolom
	colWidth := 115.0
	colWidthNo := 15.0    // Lebar kolom "No"
	colWidthSKS := 30.0   // Lebar kolom "SKS"
	colWidthNilai := 30.0 // Lebar kolom "Nilai"

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(colWidthNo, 10, headers[0], "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidth, 10, headers[1], "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidthSKS, 10, headers[2], "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidthNilai, 10, headers[3], "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	for _, row := range data {
		pdf.CellFormat(colWidthNo, 10, row[0], "1", 0, "C", false, 0, "")
		pdf.CellFormat(colWidth, 10, row[1], "1", 0, "C", false, 0, "")
		pdf.CellFormat(colWidthSKS, 10, row[2], "1", 0, "C", false, 0, "")
		pdf.CellFormat(colWidthNilai, 10, row[3], "1", 1, "C", false, 0, "")
	}
	pdf.Ln(10)

	// Menambahkan garis horizontal setelah tabel nilai
	pdf.SetLineWidth(0.5)                     // Mengatur lebar garis
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY()) // Garis horizontal dari (10, posisi Y) ke (200, posisi Y)
	pdf.Ln(5)

	// Menambahkan data mahasiswa
	footerData := []struct {
		Key   string
		Value string
	}{
		{"Jumlah SKS", "157"},
		{"Indeks Prestasi Kumulatif (IPK)", "3,99"},
		{"Lulus tanggal", "31 Februari 2051"},
		{"Predikat", "DENGAN PUJIAN"},
		{"Judul Skripsi/Tugas Akhir", "Implemetasi Komputasi Quantum"},
		{"Tittle of Thesis/Final Assignment", "Implmentation of Quantom Computing"},
	}

	// Menambahkan data footer dengan translasi bahasa Inggris di bawahnya
	for _, data := range footerData {
		key, value := data.Key, data.Value
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(70, 9, key+":", "0", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 9, value, "0", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 0, translateToEnglish(key), "0", 1, "L", false, 0, "")
		pdf.Ln(0.5)
	}

	pdf.Ln(5)

	// Menambahkan garis horizontal setelah tabel nilai
	pdf.SetLineWidth(0.5)                     // Mengatur lebar garis
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY()) // Garis horizontal dari (10, posisi Y) ke (200, posisi Y)
	pdf.Ln(5)

	// Menambahkan data mahasiswa
	publishData := []struct {
		Key   string
		Value string
	}{
		{"Diterbitkan di", "Bandung"},
		{"Tanggal", "32 Februari 2055"},
	}

	// Menambahkan data footer dengan translasi bahasa Inggris di bawahnya
	for _, data := range publishData {
		key, value := data.Key, data.Value
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(70, 9, key+":", "0", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 9, value, "0", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 0, translateToEnglish(key), "0", 1, "L", false, 0, "")
		pdf.Ln(0.5)
	}

	pdf.Ln(10)

	// Tentukan posisi x untuk membuat teks rata tengah di bagian kanan halaman
	x := 150.0

	fmt.Print(x)

	// Menambahkan tanda tangan dekan rata tengah di bagian kanan halaman
	pdf.Ln(10)
	pdf.SetX(x)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 5, "Dekan", "0", 1, "C", false, 0, "")
	pdf.SetX(x)
	pdf.CellFormat(0, 5, "Fakultas "+dekanData.Fakultas, "0", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetX(x)
	pdf.CellFormat(0, 5, " ", "0", 1, "C", false, 0, "")
	pdf.SetX(x)
	pdf.CellFormat(0, 5, dekanData.Nama, "0", 1, "C", false, 0, "")
	pdf.SetX(x)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, "NIK: "+dekanData.NIK, "0", 1, "C", false, 0, "")

	pdf.AddPage()
	// Tabel Predikat
	pred_headers := []string{"Huruf Mutu/Grade", "Bobot Nilai/Weight"}
	pred_data := [][]string{
		{"A", "4.00"},
		{"AB", "3.50"},
		{"B", "3.00"},
		{"BC", "2.50"},
		{"C", "2.00"},
		{"D", "1.00"},
	}

	// Menghitung lebar kolom
	colWidthPred := 94.0

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(colWidthPred, 10, pred_headers[0], "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidthPred, 10, pred_headers[1], "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	for _, row := range pred_data {
		pdf.CellFormat(colWidthPred, 10, row[0], "1", 0, "C", false, 0, "")
		pdf.CellFormat(colWidthPred, 10, row[1], "1", 1, "C", false, 0, "")
	}

	// Simpan file PDF
	err := pdf.OutputFileAndClose("transkrip.pdf")
	if err != nil {
		panic(err)
	}
}

// Struct untuk data dekan
type DekanData struct {
	Fakultas string
	Nama     string
	NIK      string
}

var dekanData = DekanData{
	Fakultas: "Vokasi",
	Nama:     "Aliffathur M. R. A. B. C, D.",
	NIK:      "896893",
}

// Fungsi untuk menerjemahkan teks ke dalam bahasa Inggris
func translateToEnglish(text string) string {
	translations := map[string]string{
		"Nama":                              "Name",
		"NIM":                               "Student Number",
		"Tempat, Tgl. Lahir":                "Place/date of Birth",
		"Tahun Masuk":                       "Entry Year",
		"Fakultas":                          "Faculty",
		"Program Studi":                     "Study Program",
		"Jumlah SKS":                        "Totla Credits",
		"Indeks Prestasi Kumulatif (IPK)":   "Cumulative Grade Point Average",
		"Lulus tanggal":                     "Graduated Date",
		"Predikat":                          "Predicate",
		"Judul Skripsi/Tugas Akhir":         " ",
		"Tittle of Thesis/Final Assignment": " ",
		"Diterbitkan di":                    "Issued in",
		"Tanggal":                           "Date",
	}

	englishText, found := translations[text]
	if found {
		return englishText
	}
	return text
}
