package alphavantage

import (
	"encoding/json"
	"fmt"
)

type IncomeStatements struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}
type AnnualReport struct {
	FiscalDateEnding                  string  `json:"fiscalDateEnding"`
	ReportedCurrency                  string  `json:"reportedCurrency"`
	GrossProfit                       float64 `json:"grossProfit"`
	TotalRevenue                      float64 `json:"totalRevenue"`
	CostOfRevenue                     float64 `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        float64 `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   float64 `json:"operatingIncome"`
	SellingGeneralAndAdministrative   float64 `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            float64 `json:"researchAndDevelopment"`
	OperatingExpenses                 float64 `json:"operatingExpenses"`
	InvestmentIncomeNet               float64 `json:"investmentIncomeNet"` // Have to handle case where value is "None"
	NetInterestIncome                 float64 `json:"netInterestIncome"`
	InterestIncome                    float64 `json:"interestIncome"`
	InterestExpense                   float64 `json:"interestExpense"`
	NonInterestIncome                 float64 `json:"nonInterestIncome"`
	OtherNonOperatingIncome           float64 `json:"otherNonOperatingIncome"`
	Depreciation                      float64 `json:"depreciation"`
	DepreciationAndAmortization       float64 `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   float64 `json:"incomeBeforeTax"`
	IncomeTaxExpense                  float64 `json:"incomeTaxExpense"`
	InterestAndDebtExpense            float64 `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations float64 `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       float64 `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              float64 `json:"ebit"`
	Ebitda                            float64 `json:"ebitda"`
	NetIncome                         float64 `json:"netIncome"`
}

type QuarterlyReport struct {
	FiscalDateEnding                  string  `json:"fiscalDateEnding"`
	ReportedCurrency                  string  `json:"reportedCurrency"`
	GrossProfit                       float64 `json:"grossProfit"`
	TotalRevenue                      float64 `json:"totalRevenue"`
	CostOfRevenue                     float64 `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        float64 `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   float64 `json:"operatingIncome"`
	SellingGeneralAndAdministrative   float64 `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            float64 `json:"researchAndDevelopment"`
	OperatingExpenses                 float64 `json:"operatingExpenses"`
	InvestmentIncomeNet               float64 `json:"investmentIncomeNet"`
	NetInterestIncome                 float64 `json:"netInterestIncome"`
	InterestIncome                    float64 `json:"interestIncome"`
	InterestExpense                   float64 `json:"interestExpense"`
	NonInterestIncome                 float64 `json:"nonInterestIncome"`
	OtherNonOperatingIncome           float64 `json:"otherNonOperatingIncome"`
	Depreciation                      float64 `json:"depreciation"`
	DepreciationAndAmortization       float64 `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   float64 `json:"incomeBeforeTax"`
	IncomeTaxExpense                  float64 `json:"incomeTaxExpense"`
	InterestAndDebtExpense            float64 `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations float64 `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       float64 `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              float64 `json:"ebit"`
	Ebitda                            float64 `json:"ebitda"`
	NetIncome                         float64 `json:"netIncome"`
}

func toIncomeStatements(buf []byte) (*IncomeStatements, error) {
	incomeStatements := &IncomeStatements{}
	if err := json.Unmarshal(buf, incomeStatements); err != nil {
		return nil, err
	}
	return incomeStatements, nil
}

func (c *Client) IncomeStatements(symbol string) (*IncomeStatements, error) {
	const function = "INCOME_STATEMENT"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toIncomeStatements(body)
}
