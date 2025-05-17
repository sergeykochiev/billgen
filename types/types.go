package types

import (
	"fmt"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type THData struct {
	Value string
	Width float64
}

type TDData struct {
	Value string
	Align int
}

type TemplateFunc func(ci CompanyInfo, bil BillItemList, client_company_name string, bill_number string, date string) templ.Component

type GenFunc func(w io.Writer, ci CompanyInfo, bil BillItemList, client_company_name string, bill_number string, date string) error

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
	Summ       float64
	SummString string
}

type BillItem struct {
	Name          string
	Cost          float64
	Count         int
	One_is_called string
	Summ          float64
}

func (bi *BillItem) ToTRow(i int) []TDData {
	return []TDData{
		{Value: strconv.Itoa(i), Align: 1},
		{Value: bi.Name, Align: 0},
		{Value: fmt.Sprintf("%f", bi.Cost), Align: 2},
		{Value: strconv.Itoa(bi.Count), Align: 1},
		{Value: bi.One_is_called, Align: 1},
		{Value: fmt.Sprintf("%f", bi.Summ), Align: 2},
	}
}

func (bil *BillItemList) ToTData() ([]THData, [][]TDData, []TDData) {
	var tddata = make([][]TDData, bil.Len)
	for i, bi := range bil.Bia {
		tddata[i] = bi.ToTRow(i)
	}
	var ftddata = []TDData{
		{Value: "", Align: 1},
		{Value: "", Align: 1},
		{Value: "", Align: 1},
		{Value: "Итого:", Align: 2},
		{Value: "", Align: 1},
		{Value: fmt.Sprintf("%f", bil.Summ), Align: 2},
	}
	var thdata = []THData{
		{Value: "№", Width: 0},
		{Value: "Товар", Width: .45},
		{Value: "Цена", Width: .15},
		{Value: "Кол-во", Width: .1},
		{Value: "Ед.", Width: .1},
		{Value: "Сумма", Width: .15},
	}
	return thdata, tddata, ftddata
}
