package main

import "testing"

//單元測試
//go test -v

//效能評估
//go test -v -bench=. .

func TestPrintInt2String01(t *testing.T) {
	if printInt2String01(100) != "100" {
		t.Fatal("error")
	}
}

func TestPrintInt2String02(t *testing.T) {
	if printInt2String02(int64(100)) != "100" {
		t.Fatal("error")
	}
}

func TestPrintInt2String03(t *testing.T) {
	if printInt2String03(100) != "100" {
		t.Fatal("error")
	}
}

func BenchmarkPrintInt2String01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String01(100)
	}
}

func BenchmarkPrintInt2String02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String02(int64(100))
	}
}

func BenchmarkPrintInt2String03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String03(100)
	}
}
