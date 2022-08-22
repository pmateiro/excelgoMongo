package main

import (
	"fmt"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("uploaded file:", handler.Filename)
	fmt.Println("file size:", handler.Size)
	fmt.Println("file header:", handler.Header)

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}
	fmt.Println("Successfully Uploaded")

}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
