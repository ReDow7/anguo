package model

type BalanceSheet struct {
	Code                     string
	EndDate                  string
	TotalCurrentAsserts      float64
	TotalNonCurrentAssets    float64
	totalAssets              float64
	TotalCurrentLiability    float64
	TotalNonCurrentLiability float64
	TotalLiability           float64
	_                        OwnersEquity
	_                        CurrentAssert
	_                        NonCurrentAssets
	_                        CurrentLiability
	_                        NonCurrentLiability
}

type OwnersEquity struct {
	TotalShare                                     float64
	CapitalReserves                                float64
	UndistributedProfits                           float64
	SurplusReserves                                float64
	SpecialReserve                                 float64
	TreasuryShare                                  float64
	GeneralRiskPreparation                         float64
	ConvertedDifferenceInForeignCurrencyStatements float64
	UnrealisedInvestmentLosses                     float64
	MinorityInterests                              float64
	TotalEquityOfOwnersExcludeMinorityInterests    float64
	TotalEquityOfOwnersIncludeMinorityInterests    float64
	TotalLiabilitiesAndOwnersEquity                float64
	OtherComprehensiveIncome                       float64
	OtherEquityTools                               float64
	OtherEquityToolsPreferredStock                 float64
	OtherEquityToolsPerpetualDebt                  float64
}

type CurrentLiability struct {
	ShortTermLoans                             float64
	BorrowingFromTheCentralBank                float64
	DepositsFromCustomersAndInterBank          float64
	LoansFromOtherBanks                        float64
	TransactionFinancialLiabilities            float64
	NotesPayable                               float64
	AccountsPayable                            float64
	AdvancesFromCustomers                      float64
	SellsBuysTheFinancialPropertyFunds         float64
	HandlingChargesAndCommissionsPayable       float64
	StaffSalaries                              float64
	TaxesPayable                               float64
	InterestPayable                            float64
	DividendsPayable                           float64
	OtherPayable                               float64
	AccruedExpenses                            float64
	DeferredRevenue                            float64
	ShortTermDebenturesPayable                 float64
	ThePayableReinsurance                      float64
	ProvisionForInsuranceContracts             float64
	ActingTradingSecurities                    float64
	SecuritiesUnderwritingBrokerageDeposits    float64
	NonCurrentLiabilitiesMaturingWithinOneYear float64
	OtherCurrentLiability                      float64
	DueToBanksAndOtherFinancialInstitutions    float64
	DerivativeFinancialLiabilities             float64
	Deposits                                   float64
	AgencyBusinessLiabilities                  float64
	OtherLiabilities                           float64
	PremiumsReceivedInAdvance                  float64
	DepositsReceived                           float64
	DepositsFromPolicyHolders                  float64
	UnearnedPremiumReserve                     float64
	ReserveForOutstandingLosses                float64
	LifeInsuranceReserves                      float64
	LiabilitiesOfIndependentAccounts           float64
	ClaimsPayable                              float64
	PolicyDividendPayable                      float64
	ShortTermFinancingPayable                  float64
	LiabilitiesHeldForSale                     float64
	LeaseLiabilities                           float64
	ContractLiabilities                        float64
	AccountsPayableAndBill                     float64
}

type NonCurrentLiability struct {
	LongTermLoans                   float64
	BondPayable                     float64
	LongTermPayable                 float64
	SpecificPayable                 float64
	EstimatedLiabilities            float64
	DeferTaxLiabilities             float64
	DeferIncomeNonCurLiabilities    float64
	OtherNonCurrentLiability        float64
	LongTermHealthInsuranceReserves float64
	LongTermStaffSalaries           float64
}

type NonCurrentAssets struct {
	AvailableForSaleFinancialAssets   float64
	HeldToMaturityInvestment          float64
	LongTermEquityInvestment          float64
	InvestmentProperty                float64
	TimeDeposit                       float64
	OtherAssets                       float64
	LongTermReceivables               float64
	FixedAsset                        float64
	ConstructionInProcess             float64
	ProjectGoodsAndMaterial           float64
	DisposalOfFixedAsset              float64
	BiologicalAssets                  float64
	OilAndGasAssets                   float64
	IntangibleAssets                  float64
	ResearchAndDevelopmentExpenditure float64
	Goodwill                          float64
	LongTermDeferredExpense           float64
	DeferredTaxAsset                  float64
	LoansAndPaymentsOnBehalf          float64
	OtherNonCurrentAssets             float64
	UseRightAssets                    float64
}

type CurrentAssert struct {
	CashAndCashEquivalents                           float64
	TradingFinancialAssets                           float64
	NoteReceivable                                   float64
	AccountsReceivable                               float64
	OtherReceivables                                 float64
	Prepayments                                      float64
	DividendReceivable                               float64
	InterestReceivable                               float64
	Inventories                                      float64
	PrepaidExpenses                                  float64
	NonCurrentAssetsMaturingWithinOneYear            float64
	ProvisionOfSettlementFund                        float64
	FundsLent                                        float64
	InsurancePremiumsReceivable                      float64
	AccountsReceivableReinsurance                    float64
	ProvisionOfCessionReceivable                     float64
	BuyingBackTheSaleOfFinancialAssets               float64
	OtherCurrentAssert                               float64
	CashAndBalancesWithCentralBank                   float64
	DepositsInOtherBanks                             float64
	PreciousMetals                                   float64
	DerivativeFinancialInstruments                   float64
	ReceivableDepositForUndueDutyOfReinsurance       float64
	ReinsurersShareOfClaimReserves                   float64
	ReinsurersShareOfLifeInsuranceReserves           float64
	ReinsurersShareOfLongTermHealthInsuranceReserves float64
	RefundDeposits                                   float64
	PolicyPledgeLoans                                float64
	RefundCapitalDeposits                            float64
	AssetsForIndependenceAccount                     float64
	InvestmentAsReceivables                          float64
	LendingFunds                                     float64
	AssetsHeldForSale                                float64
	FinancialAssetsMeasuredAtAmortizedCost           float64
	FairValueThroughOtherComprehensiveIncome         float64
	DebtInvestment                                   float64
	OtherDebtInvestment                              float64
	OtherEquityInvestment                            float64
	OtherNonCurrentFinancialAssets                   float64
	ReceivableFinancing                              float64
	ContractAssets                                   float64
	AccountsReceivableAndBill                        float64
}
