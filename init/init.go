package init

import pdf "github.com/adrg/go-wkhtmltopdf"

func Init() error {
	return pdf.Init()
}

func Destroy() {
	pdf.Destroy()
}
