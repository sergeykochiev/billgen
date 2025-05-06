package types

import "github.com/a-h/templ"

type TemplateFunc func(ci CompanyInfo, bil BillItemList, client_company_name string, bill_number string, date string) templ.Component

type GenFunc func(ci CompanyInfo, bil BillItemList, client_company_name string, bill_number string, date string) error

type CompanyDetails struct {
	Bank string
	Rs   string
	Ks   string
	Bik  string
}

type CompanyInfo struct {
	Inn        string
	Name       string
	Address    string
	Number     string
	Details    CompanyDetails
	PersonResp string
}

type BillItemList struct {
	Bia        []BillItem
	Len        int
	Summ       float32
	SummString string
}

type BillItem struct {
	Name          string
	Cost          float32
	Count         int
	One_is_called string
	Summ          float32
}
