package alphavantage

import (
	"encoding/json"
	"fmt"
)

type CashFlows struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}

type AnnualReport struct {
	FiscalDateEnding                      string  `json:"fiscalDateEnding"`
	ReportedCurrency                      string  `json:"reportedCurrency"`
	OperatingCashflow                     float64 `json:"operatingCashflow"`
	PaymentsForOperatingActivities        float64 `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities       float64 `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities          float64 `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets               float64 `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization  float64 `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                   float64 `json:"capitalExpenditures"`
	ChangeInReceivables                   float64 `json:"changeInReceivables"`
	ChangeInInventory                     float64 `json:"changeInInventory"`
	ProfitLoss                            float64 `json:"profitLoss"`
	CashflowFromInvestment                float64 `json:"cashflowFromInvestment"`
	CashflowFromFinancing                 float64 `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt float64 `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock    float64 `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity         float64 `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock float64 `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                        float64 `json:"dividendPayout"`
	DividendPayoutCommonStock             float64 `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock          float64 `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock     float64 `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebt    float64 `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock  float64 `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity        float64 `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock       float64 `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents        float64 `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                  float64 `json:"changeInExchangeRate"`
	NetIncome                             float64 `json:"netIncome"`
}

type QuarterlyReport struct {
	FiscalDateEnding                      string  `json:"fiscalDateEnding"`
	ReportedCurrency                      string  `json:"reportedCurrency"`
	OperatingCashflow                     float64 `json:"operatingCashflow"`
	PaymentsForOperatingActivities        float64 `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities       float64 `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities          float64 `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets               float64 `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization  float64 `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                   float64 `json:"capitalExpenditures"`
	ChangeInReceivables                   float64 `json:"changeInReceivables"`
	ChangeInInventory                     float64 `json:"changeInInventory"`
	ProfitLoss                            float64 `json:"profitLoss"`
	CashflowFromInvestment                float64 `json:"cashflowFromInvestment"`
	CashflowFromFinancing                 float64 `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt float64 `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock    float64 `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity         float64 `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock float64 `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                        float64 `json:"dividendPayout"`
	DividendPayoutCommonStock             float64 `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock          float64 `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock     float64 `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebt    float64 `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock  float64 `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity        float64 `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock       float64 `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents        float64 `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                  float64 `json:"changeInExchangeRate"`
	NetIncome                             float64 `json:"netIncome"`
}

func toCashFlows(buf []byte) (*CashFlows, error) {
	cashFlows := &CashFlows{}
	if err := json.Unmarshal(buf, cashFlows); err != nil {
		return nil, fmt.Errorf("unable to unmarshal cash flows: %v", err)
	}
	return cashFlows, nil
}

func (c *Client) CashFlows(symbol string) (*CashFlows, error) {
	const function = "CASH_FLOW"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	return toCashFlows(body)
}
