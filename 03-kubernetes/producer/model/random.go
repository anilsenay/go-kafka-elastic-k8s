package model

import (
	"math/rand"
)

var titles = []string{
	"Samsung Galaxy A51 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A71 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A31 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A21s 64 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A11 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A01 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A50 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A30s 64 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A20s 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A10s 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A10 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A70 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A40 64 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A20 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A30 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A7 2018 64 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A6 Plus 64 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A8 2018 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A6 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A5 2017 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A3 2017 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A8 Plus 2018 64 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A8 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A7 2017 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A8 2015 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A7 2016 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A5 2016 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A3 2016 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A5 2015 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A3 2015 16 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2018 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 Pro 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2015 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2018 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 Pro 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2015 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2018 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 Pro 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2015 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2018 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 Pro 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2015 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2018 128 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 Pro 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2016 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2015 32 GB (Samsung Türkiye Garantili)",
	"Samsung Galaxy A9 2018 128 GB (Samsung Türkiye Garantili)",
	"Apple iPhone 11 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 128 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 Pro 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 Pro 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 Pro 512 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 Pro Max 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 Pro Max 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone 11 Pro Max 512 GB (Apple Türkiye Garantili)",
	"Apple iPhone XS Max 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone XS Max 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone XS Max 512 GB (Apple Türkiye Garantili)",
	"Apple iPhone XS 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone XS 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone XS 512 GB (Apple Türkiye Garantili)",
	"Apple iPhone XR 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone XR 128 GB (Apple Türkiye Garantili)",
	"Apple iPhone XR 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone X 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone X 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone 8 Plus 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone 8 Plus 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone 8 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone 8 256 GB (Apple Türkiye Garantili)",
	"Apple iPhone 7 Plus 32 GB (Apple Türkiye Garantili)",
	"Apple iPhone 7 Plus 128 GB (Apple Türkiye Garantili)",
	"Apple iPhone 7 32 GB (Apple Türkiye Garantili)",
	"Apple iPhone 7 128 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6S Plus 32 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6S Plus 128 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6S 32 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6S 128 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6 Plus 16 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6 Plus 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6 16 GB (Apple Türkiye Garantili)",
	"Apple iPhone 6 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone SE 16 GB (Apple Türkiye Garantili)",
	"Apple iPhone SE 32 GB (Apple Türkiye Garantili)",
	"Apple iPhone SE 64 GB (Apple Türkiye Garantili)",
	"Apple iPhone SE 128 GB (Apple Türkiye Garantili)",
}

func GetRandomTitle() string {
	return titles[rand.Intn(len(titles))]
}

func RandomId() int64 {
	return rand.Int63n(1000000)
}

func RandomPrice() float64 {
	return rand.Float64() * 30000
}

func GenerateProduct() Product {
	return Product{
		Id:       RandomId(),
		Title:    GetRandomTitle(),
		Price:    RandomPrice(),
		Category: "Cep Telefonları",
	}
}
