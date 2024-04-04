package alphavantage

import (
	"encoding/json"
	"fmt"
)

type BalanceSheets struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}

type AnnualReport struct {
	FiscalDateEnding                       string  `json:"fiscalDateEnding"`
	ReportedCurrency                       string  `json:"reportedCurrency"`
	TotalAssets                            float64 `json:"totalAssets"`
	TotalCurrentAssets                     float64 `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  float64 `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            float64 `json:"cashAndShortTermInvestments"`
	Inventory                              float64 `json:"inventory"`
	CurrentNetReceivables                  float64 `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  float64 `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 float64 `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE float64 `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       float64 `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      float64 `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               float64 `json:"goodwill"`
	Investments                            float64 `json:"investments"`
	LongTermInvestments                    float64 `json:"longTermInvestments"`
	ShortTermInvestments                   float64 `json:"shortTermInvestments"`
	OtherCurrentAssets                     float64 `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  float64 `json:"otherNonCurrentAssets"`
	TotalLiabilities                       float64 `json:"totalLiabilities"`
	TotalCurrentLiabilities                float64 `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 float64 `json:"currentAccountsPayable"`
	DeferredRevenue                        float64 `json:"deferredRevenue"`
	CurrentDebt                            float64 `json:"currentDebt"`
	ShortTermDebt                          float64 `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             float64 `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                float64 `json:"capitalLeaseObligations"`
	LongTermDebt                           float64 `json:"longTermDebt"`
	CurrentLongTermDebt                    float64 `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 float64 `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 float64 `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                float64 `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             float64 `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 float64 `json:"totalShareholderEquity"`
	TreasuryStock                          float64 `json:"treasuryStock"`
	RetainedEarnings                       float64 `json:"retainedEarnings"`
	CommonStock                            float64 `json:"commonStock"`
	CommonStockSharesOutstanding           float64 `json:"commonStockSharesOutstanding"`
}

type QuarterlyReport struct {
	FiscalDateEnding                       string  `json:"fiscalDateEnding"`
	ReportedCurrency                       string  `json:"reportedCurrency"`
	TotalAssets                            float64 `json:"totalAssets"`
	TotalCurrentAssets                     float64 `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  float64 `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            float64 `json:"cashAndShortTermInvestments"`
	Inventory                              float64 `json:"inventory"`
	CurrentNetReceivables                  float64 `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  float64 `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 float64 `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE float64 `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       float64 `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      float64 `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               float64 `json:"goodwill"`
	Investments                            float64 `json:"investments"`
	LongTermInvestments                    float64 `json:"longTermInvestments"`
	ShortTermInvestments                   float64 `json:"shortTermInvestments"`
	OtherCurrentAssets                     float64 `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  float64 `json:"otherNonCurrentAssets"`
	TotalLiabilities                       float64 `json:"totalLiabilities"`
	TotalCurrentLiabilities                float64 `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 float64 `json:"currentAccountsPayable"`
	DeferredRevenue                        float64 `json:"deferredRevenue"`
	CurrentDebt                            float64 `json:"currentDebt"`
	ShortTermDebt                          float64 `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             float64 `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                float64 `json:"capitalLeaseObligations"`
	LongTermDebt                           float64 `json:"longTermDebt"`
	CurrentLongTermDebt                    float64 `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 float64 `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 float64 `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                float64 `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             float64 `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 float64 `json:"totalShareholderEquity"`
	TreasuryStock                          float64 `json:"treasuryStock"`
	RetainedEarnings                       float64 `json:"retainedEarnings"`
	CommonStock                            float64 `json:"commonStock"`
	CommonStockSharesOutstanding           float64 `json:"commonStockSharesOutstanding"`
}

func toBalanceSheets(buf []byte) (*BalanceSheets, error) {
	balanceSheets := &BalanceSheets{}
	if err := json.Unmarshal(buf, balanceSheets); err != nil {
		return nil, err
	}
	return balanceSheets, nil
}

func (c *Client) BalanceSheets(symbol string) (*BalanceSheets, error) {
	const function = "BALANCE_SHEET"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toBalanceSheets(body)
}
