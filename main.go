package main

import (
	"log"
	"os"

	"github.com/sergeykochiev/billgen/gen"
	"github.com/sergeykochiev/billgen/types"
)

func main() {
	f, err := os.Create("test.pdf")
	if err != nil {
		log.Fatal("failed to create pdf")
	}
	ci := types.CompanyInfo{
		Inn:     "SampleCompanyInn",
		Name:    "SampleCompanyName",
		Address: "SampleCompanyAddress",
		Number:  "SampleCompanyPhone",
		Details: types.CompanyDetails{
			Bank: "SampleCompanyBank",
			Rs:   "SampleCompanyRs",
			Ks:   "SampleCompanyKs",
			Bik:  "SampleCompanyBik",
		},
		PersonResp: "SampleCompanyPersonResp",
	}
	bil := types.BillItemList{
		Bia: []types.BillItem{
			{
				Name:          "SampleBillItemName",
				Cost:          0,
				Count:         0,
				One_is_called: "SampleOneIsCalled",
				Summ:          0,
			},
		},
		Len:  1,
		Summ: 0,
	}
	f.Write(gen.CreateBillPdf(ci, bil, "SampleClientCompanyName", "120938123", "2 oct 2025"))
}
