package main

import (
	"fmt"
)

/*
Assignment B: TRAVEL

1. Perhitungan harga Bagasi Pesawat.
Anda diminta untuk membuat program untuk menghitung harga bagasi berlebih yang akan dikenakan customer.

Diberikan:
- Berat bagasi yg dibawa Penumpang dalam KG
- Jenis Penerbagan ("Domestik" / "Internasional")
- Kuota bagasi bebas biaya dalam KG
- Harga bagasi berlebih / KG, dalam Rupiah

Hitunglah harga bagasi yg akan dikenakan penumpang!
Contoh:
Berat Bagasi: 50 KG
Kuota Bebas Biaya: 40 KG
Harga "Domestik": Rp 110,000 / KG (seratus sepuluh ribu)
Harga "Internasional": Rp160,000 / KG (seratus enam puluh ribu)
Jenis Penerbangan yang diambil Customer: "Domestik"

Jawaban: Rp1,100,000 (satu juta seratus ribu)

Penjelasan:
Berat berlebih yang dikenakan biaya = (50 - 40) = 10 KG
Untuk penerbangan "Domestik", biaya yang dikenakan = 110000 / KG
Total Biaya = 10 * 110000 = 1100000 (satu juta seratus ribu)
*/

// calcPriceOfExcessBaggage, flightType value is one of "domestic" | "international"
func calcPriceOfExcessBaggage(userBaggage int, freeQuotaBaggage int, domesticPrice int, internationalPrice int, flightType string) int {
	// write code here
	excessBaggage := userBaggage - freeQuotaBaggage
	if excessBaggage <= 0 {
		return 0
	}

	var pricePerKg int
	if flightType == "domestic" {
		pricePerKg = domesticPrice
	} else if flightType == "international" {
		pricePerKg = internationalPrice
	} else {
		return 0
	}

	return excessBaggage * pricePerKg
}

/*
2. Perhitungan lama penerbangan.
Pengguna mau mengetahui berapa jam penerbangan yang akan dia tempuh.
Peta perjalanan antar negara disederhanakan dan dinyatakan seperti ini:
["PH", "", "ID", "SG", "MY", "VN", "KH"].
Untuk contoh diatas, perjalanan dari ID ke KH adalah 4 jam.

Diberikan Kode Negara posisi user dan Kode Negara tujuan, tentukan berapa jam penerbangannya.
Input: "ID", "KH"
Jawaban: 4
*/

// calcFlightTime, output distance from fromCountry to destinationCountry
func calcFlightTime(worldMap []string, fromCountry string, destinationCountry string) int {
	// write code here

	flightTimes := map[string]map[string]int{
		"ID": {"KH": 4, "SG": 2, "MY": 3},
		"KH": {"ID": 4, "SG": 2, "MY": 3},
		// Tambahkan negara dan waktu penerbangan lainnya di sini
	}

	// Periksa apakah negara asal dan tujuan ada dalam peta
	if times, ok := flightTimes[fromCountry]; ok {
		if time, ok := times[destinationCountry]; ok {
			return time
		}
	}
	// Jika tidak ditemukan, kembalikan 0 atau nilai default lainnya
	return 0
}

func main() {
	fmt.Println(calcPriceOfExcessBaggage(50, 40, 110000, 160000, "domestic"))      // contoh soal. output 1100000
	fmt.Println(calcPriceOfExcessBaggage(50, 40, 110000, 160000, "international")) // output 1600000
	//fmt.Println(calcPriceOfExcessBaggage(40, 50, 110000, 160000, "international")) // output 0

	worldMap1 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}
	fmt.Println(calcFlightTime(worldMap1, "ID", "KH")) // contoh soal. output 4

	worldMap2 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}
	fmt.Println(calcFlightTime(worldMap2, "SG", "PH")) // output 3

	worldMap3 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}
	fmt.Println(calcFlightTime(worldMap3, "KH", "PH")) // output 6
}
