package main

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/app-2.git/handler"
)

func main() {
	http.HandleFunc("/employees", handler.GetEmployeList)
	http.HandleFunc("/companies", handler.GetCompanyList)

	if err := http.ListenAndServe(":2323", nil); err != nil {
		log.Fatalf("failed serve server on port 2323, err: %s", err)
	}
}
