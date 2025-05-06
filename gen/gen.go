package gen

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/sergeykochiev/billgen/templates"
	"github.com/sergeykochiev/billgen/types"
)

func createPdfFromHtml(c templ.Component, w io.Writer) error {
	var sr bytes.Buffer
	c.Render(context.Background(), &sr)
	obj, err := pdf.NewObjectFromReader(&sr)
	if err != nil {
		return err
	}
	converter, err := pdf.NewConverter()
	if err != nil {
		return err
	}
	defer converter.Destroy()
	converter.Add(obj)
	if err := converter.Run(w); err != nil {
		return err
	}
	return nil
}

func CreateBillPdf(w io.Writer, ci types.CompanyInfo, bil types.BillItemList, client_company_name string, bill_number string, date string) error {
	return createPdfFromHtml(templates.BillTemplate(ci, bil, client_company_name, bill_number, date), w)
}

func CreateInvoicePdf(w io.Writer, ci types.CompanyInfo, bil types.BillItemList, client_company_name string, bill_number string, date string) error {
	return createPdfFromHtml(templates.InvoiceTemplate(ci, bil, client_company_name, bill_number, date), w)
}
