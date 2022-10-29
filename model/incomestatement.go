package model

type IncomeStatement struct {
	Code                                                   string
	EndDate                                                string
	TotalRevenue                                           float64
	Revenue                                                float64
	InterestsIncome                                        float64
	PremiumEarned                                          float64
	HandlingChargesAndCommissionsIncome                    float64
	HandlingChargesAndCommissionsNetIncome                 float64
	OtherNetIncome                                         float64
	OtherBusinessNetIncome                                 float64
	PremiumIncome                                          float64
	OutPremium                                             float64
	AppropriationOfDepositForUndueDuty                     float64
	ReinsurancePremiumIncome                               float64
	NetIncomeFromSecuritiesTradingBrokerageBusiness        float64
	NetIncomeFromSecuritiesUnderwritingBusiness            float64
	NetIncomeFromEntrustedCustomerAssetsManagementBusiness float64
	OtherBusinessIncome                                    float64
	ProfitAndLossOnFluctuationOfFairMarketValue            float64
	InvestmentIncome                                       float64
	InvestmentIncomeFromAssociatesAndJointVentures         float64
	ExchangeNetIncome                                      float64
	TotalOperatingCost                                     float64
	OperatingCost                                          float64
	InterestExpense                                        float64
	HandlingChargesAndCommissionsExpense                   float64
	BusinessTaxesAndSurcharges                             float64
	SellExpense                                            float64
	AdministrationExpenses                                 float64
	FinancialExpenses                                      float64
	AssetImpairmentLoss                                    float64
	PremiumRefund                                          float64
	CompensationExpenditure                                float64
	AppropriationOfDepositForDuty                          float64
	DividendExpensesForTheInsured                          float64
	ReinsuranceExpense                                     float64
	OperatingExpense                                       float64
	AmortizedCompensationExpenses                          float64
	AmortizedDepositForDuty                                float64
	AmortizedReinsuranceExpense                            float64
	OtherBusinessCost                                      float64
	OperateProfit                                          float64
	NonOperateIncome                                       float64
	NonOperateExpense                                      float64
	NonCurrentAssetsDisposalNetLoss                        float64
	TotalProfit                                            float64
	IncomeTax                                              float64
	NetIncome                                              float64
	NetIncomeBelongsToOwner                                float64
	MinorityGain                                           float64
	OtherComprehensiveIncome                               float64
	TotalComprehensiveIncome                               float64
	TotalComprehensiveIncomeBelongToOwner                  float64
	TotalComprehensiveIncomeBelongToMinority               float64
	InsuranceExpense                                       float64
	UndistributedProfitOfBeginning                         float64
	DistributiveProfits                                    float64
	ResearchAndDevelopExpense                              float64
	InterestExpenseInFinancialExpense                      float64
	InterestIncomeInFinancialExpense                       float64
	TransferInFromSurplusReserve                           float64
	TransferInFromHousingRevolvingFund                     float64
	TransferInFromOther                                    float64
	BeforehandYearProfitAndLossAdjustment                  float64
	WithdrawLegalSurplus                                   float64
	WithdrawLegalPublicWelfareFund                         float64
	WithdrawEnterpriseDevelopmentFund                      float64
	WithdrawReserveFund                                    float64
	WithdrawOtherSurplus                                   float64
	WorkersWelfare                                         float64
	ProfitAvailableForDistributionByShareholders           float64
	PreferredStockDividendsPayable                         float64
	CommonStockDividendsPayable                            float64
	CommonStockDividendsConvertedIntoCapitalStock          float64
	NonRecurringGainsAndLosses                             float64
	CreditImpairmentLoss                                   float64
	NetExposureHedgingIncome                               float64
	OtherAssertImpairmentLoss                              float64
	IncomeFromStopOfFinancialAssetsMeasuredAtAmortizedCost float64
	AssertDisposalIncome                                   float64
	OtherIncome                                            float64
	ContinuedNetProfit                                     float64
	EndNetProfit                                           float64
}

func (is *IncomeStatement) RE(sheet *BalanceSheet, WACC float64) float64 {
	roce := is.ROCE(sheet)
	return (roce - WACC) * sheet.NetOperationAssert(is.TotalRevenue)
}

func (is *IncomeStatement) ROCE(sheet *BalanceSheet) float64 {
	return is.NetOperationProfit() / sheet.NetOperationAssert(is.TotalRevenue)
}

func (is *IncomeStatement) NetOperationProfit() float64 {
	var nop float64
	adjustRateAfterIncomeTax := 1 - is.AverageIncomeTaxRate()
	nop += is.TotalComprehensiveIncome
	nop += is.FinancialExpenses * adjustRateAfterIncomeTax
	nop -= is.OtherIncome * adjustRateAfterIncomeTax
	nop -= is.InvestmentIncome * adjustRateAfterIncomeTax
	nop += is.InvestmentIncomeFromAssociatesAndJointVentures * adjustRateAfterIncomeTax
	nop -= is.IncomeFromStopOfFinancialAssetsMeasuredAtAmortizedCost * adjustRateAfterIncomeTax
	nop -= is.ExchangeNetIncome * adjustRateAfterIncomeTax
	nop -= is.ProfitAndLossOnFluctuationOfFairMarketValue * adjustRateAfterIncomeTax
	nop -= is.NonOperateIncome * adjustRateAfterIncomeTax
	nop += is.NonOperateExpense * adjustRateAfterIncomeTax
	return nop
}

func (is *IncomeStatement) AverageIncomeTaxRate() float64 {
	return is.IncomeTax / is.TotalProfit
}
