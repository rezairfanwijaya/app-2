package main

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/app-2.git/handler"
	"github.com/rezairfanwijaya/app-2.git/response"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := response.Success{Data: "pong app-2"}
		res, _ := resp.ToJSON()
		_, _ = w.Write(res)
	})
	http.HandleFunc("/employees", handler.GetEmployeList)
	http.HandleFunc("/companies", handler.GetCompanyList)

	if err := http.ListenAndServe(":2323", nil); err != nil {
		log.Fatalf("failed serve server on port 2323, err: %s", err)
	}
}
