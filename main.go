package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/a-h/templ"
	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/sergeykochiev/billgen/templates"
	"github.com/sergeykochiev/billgen/types"
)

func createPdfFromHtml(c templ.Component) []byte {
	if err := pdf.Init(); err != nil {
		log.Fatal(err)
	}
	defer pdf.Destroy()
	var sr bytes.Buffer
	c.Render(context.Background(), &sr)
	obj, err := pdf.NewObjectFromReader(&sr)
	if err != nil {
		log.Fatal(err)
	}
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()
	var sw bytes.Buffer
	converter.Add(obj)
	if err := converter.Run(&sw); err != nil {
		log.Fatal(err)
	}
	return sw.Bytes()
}

func CreateBillPdf(ci types.CompanyInfo, bil types.BillItemList, client_company_name string, bill_number string, date string) []byte {
	return createPdfFromHtml(templates.BillTemplate(ci, bil, client_company_name, bill_number, date))
}

func CreateInvoicePdf(ci types.CompanyInfo, bil types.BillItemList, client_company_name string, bill_number string, date string) []byte {
	return createPdfFromHtml(templates.InvoiceTemplate(ci, bil, client_company_name, bill_number, date))
}

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
	f.Write(CreateBillPdf(ci, bil, "SampleClientCompanyName", "120938123", "2 oct 2025"))
}
