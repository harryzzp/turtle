package main

import "github.com/jung-kurt/gofpdf"

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world 中文")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err.Error())
	}
}