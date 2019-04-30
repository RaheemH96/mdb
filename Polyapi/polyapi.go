package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//convert json struct to g struct via: https://transform.now.sh/json-to-go/
type marketData []struct {
	Symbol                 string    `json:"symbol"`
	ReportDate             time.Time `json:"reportDate"`
	ReportDateStr          string    `json:"reportDateStr"`
	GrossProfit            int64     `json:"grossProfit"`
	CostOfRevenue          int64     `json:"costOfRevenue"`
	OperatingRevenue       int64     `json:"operatingRevenue"`
	TotalRevenue           int64     `json:"totalRevenue"`
	OperatingIncome        int64     `json:"operatingIncome"`
	NetIncome              int64     `json:"netIncome"`
	ResearchAndDevelopment int64     `json:"researchAndDevelopment"`
	OperatingExpense       int64     `json:"operatingExpense"`
	CurrentAssets          int64     `json:"currentAssets"`
	TotalAssets            int64     `json:"totalAssets"`
	TotalLiabilities       int64     `json:"totalLiabilities"`
	CurrentCash            int64     `json:"currentCash"`
	CurrentDebt            int64     `json:"currentDebt"`
	TotalCash              int64     `json:"totalCash"`
	TotalDebt              int64     `json:"totalDebt"`
	ShareholderEquity      int64     `json:"shareholderEquity"`
	CashChange             int64     `json:"cashChange"`
	CashFlow               int64     `json:"cashFlow"`
	OperatingGainsLosses   int       `json:"operatingGainsLosses"`
}

func main() {
	//variable created for the structure
	var details marketData
	//place polygon link
	response, _ := http.Get("https://api.polygon.io/v1/meta/symbols/aapl/financials?apiKey=tVKROaw3qfY_V4kFgEsjhSbs_5MmsUAF47whe3")
	bytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &details)
	// loop establishing how many quarters of data to run
	for i := 0; i < 5; i++ {
		fmt.Printf("Symbol: %v\n", details[i].Symbol)
		fmt.Printf("Report Date: %v\n", details[i].ReportDateStr)
		fmt.Printf("GrossProfit: %v\n", details[i].GrossProfit)
		fmt.Printf("Cost of Revenue: %v\n", details[i].CostOfRevenue)
		fmt.Printf("OperatingRevenue : %v\n", details[i].OperatingRevenue)
		fmt.Printf("TotalRevenue   : %v\n", details[i].TotalRevenue)
		fmt.Printf("OperatingIncome  : %v\n", details[i].OperatingIncome)
		fmt.Printf("CurrentAssets: %v\n", details[i].CurrentAssets)
		fmt.Printf("TotalAssets: %v\n", details[i].TotalAssets)
		fmt.Printf("TotalLiabilities: %v\n", details[0].TotalLiabilities)
		fmt.Printf("CurrentCash: %v\n", details[i].CurrentCash)
		fmt.Printf("CurrentDebt : %v\n", details[i].CurrentDebt)
		fmt.Printf("TotalCash: %v\n", details[i].TotalCash)
		fmt.Printf("TotalDebt : %v\n", details[i].TotalDebt)
		fmt.Printf("ShareholderEquity: %v\n", details[i].ShareholderEquity)
		fmt.Printf("CashChange: %v\n", details[i].CashChange)
		fmt.Printf("CashFlow: %v\n", details[i].CashFlow)
		fmt.Printf("OperatingGainsLosses: %v\n\n", details[i].OperatingGainsLosses)

	}

}
