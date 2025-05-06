package gen

import (
	"bytes"
	"context"
	"log"

	"github.com/a-h/templ"
	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/sergeykochiev/billgen/templates"
	"github.com/sergeykochiev/billgen/types"
)

func createPdfFromHtml(c templ.Component) []byte {
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
