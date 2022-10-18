package tushare

import (
	"anguo/model"
	"fmt"
	"strings"
)

func GetBalanceSheetOfLastYearForGiveCode(code string) (*model.BalanceSheet, error) {
	if code == "" {
		return nil, fmt.Errorf("emtpy code when fetch balance sheet")
	}
	return GetBalanceSheetForGiveCodeAndDate(code, GetLastYearEndDate())
}

func GetBalanceSheetForGiveCodeAndDate(code, date string) (*model.BalanceSheet, error) {
	balanceSheet, err := getBalanceSheetFromTushare(code, date)
	if err != nil {
		return nil, err
	}
	return balanceSheet, nil
}

func getBalanceSheetFromTushare(code, date string) (*model.BalanceSheet, error) {
	fields := strings.Join([]string{
		fieldEndDate,
		fieldTsCode,
		fieldTotalShare,
		fieldCapRese,
		fieldUndistrPorfit,
		fieldSurplusRese,
		fieldSpecialRese,
		fieldMoneyCap,
		fieldTradAsset,
		fieldNotesReceiv,
		fieldAccountsReceiv,
		fieldOthReceiv,
		fieldPrePayment,
		fieldDivReceiv,
		fieldIntReceiv,
		fieldInventories,
		fieldAmorExp,
		fieldNcaWithin1y,
		fieldSettRsrv,
		fieldLoantoOthBankFi,
		fieldPremiumReceiv,
		fieldReinsurReceiv,
		fieldReinsurResReceiv,
		fieldPurResaleFa,
		fieldOthCurAssets,
		fieldTotalCurAssets,
		fieldFaAvailForSale,
		fieldHtmInvest,
		fieldLtEqtInvest,
		fieldInvestRealEstate,
		fieldTimeDeposits,
		fieldOthAssets,
		fieldLtRec,
		fieldFixAssets,
		fieldCip,
		fieldConstMaterials,
		fieldFixedAssetsDisp,
		fieldProducBioAssets,
		fieldOilAndGasAssets,
		fieldIntanAssets,
		fieldRAndD,
		fieldGoodwill,
		fieldLtAmorExp,
		fieldDeferTaxAssets,
		fieldDecrInDisbur,
		fieldOthNca,
		fieldTotalNca,
		fieldCashReserCb,
		fieldDeposInOthBfi,
		fieldPrecMetals,
		fieldDerivAssets,
		fieldRrReinsUnePrem,
		fieldRrReinsOutstdCla,
		fieldRrReinsLinsLiab,
		fieldRrReinsLthinsLiab,
		fieldRefundDepos,
		fieldPhPledgeLoans,
		fieldRefundCapDepos,
		fieldIndepAcctAssets,
		fieldClientDepos,
		fieldClientProv,
		fieldTransacSeatFee,
		fieldInvestAsReceiv,
		fieldTotalAssets,
		fieldLtBorr,
		fieldStBorr,
		fieldCbBorr,
		fieldDeposIbDeposits,
		fieldLoanOthBank,
		fieldTradingFl,
		fieldNotesPayable,
		fieldAcctPayable,
		fieldAdvReceipts,
		fieldSoldForRepurFa,
		fieldCommPayable,
		fieldPayrollPayable,
		fieldTaxesPayable,
		fieldIntPayable,
		fieldDivPayable,
		fieldOthPayable,
		fieldAccExp,
		fieldDeferredInc,
		fieldStBondsPayable,
		fieldPayableToReinsurer,
		fieldRsrvInsurCont,
		fieldActingTradingSec,
		fieldActingUwSec,
		fieldNonCurLiabDue1y,
		fieldOthCurLiab,
		fieldTotalCurLiab,
		fieldBondPayable,
		fieldLtPayable,
		fieldSpecificPayables,
		fieldEstimatedLiab,
		fieldDeferTaxLiab,
		fieldDeferIncNonCurLiab,
		fieldOthNcl,
		fieldTotalNcl,
		fieldDeposOthBfi,
		fieldDerivLiab,
		fieldDepos,
		fieldAgencyBusLiab,
		fieldOthLiab,
		fieldPremReceivAdva,
		fieldDeposReceived,
		fieldPhInvest,
		fieldReserUnePrem,
		fieldReserOutstdClaims,
		fieldReserLinsLiab,
		fieldReserLthinsLiab,
		fieldIndeptAccLiab,
		fieldPledgeBorr,
		fieldIndemPayable,
		fieldPolicyDivPayable,
		fieldTOtalLiab,
		fieldTreasuryShare,
		fieldOrdinRiskReser,
		fieldForexDiffer,
		fieldInvestLossUnconf,
		fieldMinorityInt,
		fieldTotalHldrEqyExcMinInt,
		fieldTotalLiabHldrEqy,
		fieldLtPayrollPayable,
		fieldOthCompIncome,
		fieldOthEqtTools,
		fieldOthEqtToolsPShr,
		fieldLendingFunds,
		fieldAccReceivable,
		fieldStFinPayable,
		fieldPayables,
		fieldHfsAssets,
		fieldHfsSales,
		fieldCostFinAssets,
		fieldFairValueFinAssets,
		fieldCipTotal,
		fieldOthPayTotal,
		fieldLongPayTotal,
		fieldDebtInvest,
		fieldOthDebtInvest,
		fieldOthEqInvest,
		fieldOthIlliqFinAssets,
		fieldOthEqPpbond,
		fieldReceivFinancing,
		fieldUseRightAssets,
		fieldLeaseLiab,
		fileContractAssets,
		fieldContractLiab,
		fieldAccountsReceivBill,
		fieldAccountsPay,
		fieldOthRcvTotal,
		fieldFixAssetsTotal,
		fieldUpdateFlag,
	}, ",")
	params := map[string]interface{}{
		fieldEndDate: date,
		fieldTsCode:  code,
	}
	resp, err := fetchTushareRawData("balancesheet", fields, params)
	if err != nil {
		return nil, err
	}
	return parseBalanceSheetRecord(resp)
}

func parseBalanceSheetRecord(resp *Response) (*model.BalanceSheet, error) {
	err := resp.anyError()
	if err != nil {
		return nil, err
	}
	var sheets []model.BalanceSheet
	for _, item := range resp.Data.Items {
		var sheet = model.BalanceSheet{}
		for i, field := range resp.Data.Fields {
			switch field {
			case fieldEndDate:
				if str, ok := item[i].(string); ok {
					sheet.EndDate = str
				}
			case fieldTsCode:
				if str, ok := item[i].(string); ok {
					sheet.Code = str
				}
			case fieldTotalShare:
				if val, ok := item[i].(float64); ok {
					sheet.TotalShare = val
				}
			case fieldCapRese:
				if val, ok := item[i].(float64); ok {
					sheet.CapitalReserves = val
				}
			case fieldUndistrPorfit:
				if val, ok := item[i].(float64); ok {
					sheet.UndistributedProfits = val
				}
			case fieldSurplusRese:
				if val, ok := item[i].(float64); ok {
					sheet.SurplusReserves = val
				}
			case fieldSpecialRese:
				if val, ok := item[i].(float64); ok {
					sheet.SpecialReserve = val
				}
			case fieldMoneyCap:
				if val, ok := item[i].(float64); ok {
					sheet.CashAndCashEquivalents = val
				}
			case fieldTradAsset:
				if val, ok := item[i].(float64); ok {
					sheet.TradingFinancialAssets = val
				}
			case fieldNotesReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.NoteReceivable = val
				}
			case fieldAccountsReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.AccountsReceivable = val
				}
			case fieldOthReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.OtherReceivables = val
				}
			case fieldPrePayment:
				if val, ok := item[i].(float64); ok {
					sheet.Prepayments = val
				}
			case fieldDivReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.DividendReceivable = val
				}
			case fieldIntReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.InterestReceivable = val
				}
			case fieldInventories:
				if val, ok := item[i].(float64); ok {
					sheet.Inventories = val
				}
			case fieldAmorExp:
				if val, ok := item[i].(float64); ok {
					sheet.PrepaidExpenses = val
				}
			case fieldNcaWithin1y:
				if val, ok := item[i].(float64); ok {
					sheet.NonCurrentAssetsMaturingWithInOneYear = val
				}
			case fieldSettRsrv:
				if val, ok := item[i].(float64); ok {
					sheet.ProvisionOfSettlementFund = val
				}
			case fieldLoantoOthBankFi:
				if val, ok := item[i].(float64); ok {
					sheet.FundsLent = val
				}
			case fieldPremiumReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.InsurancePremiumsReceivable = val
				}
			case fieldReinsurReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.ReinsurancePremiumsReceivable = val
				}
			case fieldReinsurResReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.ProvisionOfCessionReceivable = val
				}
			case fieldPurResaleFa:
				if val, ok := item[i].(float64); ok {
					sheet.BuyingBackTheSaleOfFinancialAssets = val
				}
			case fieldOthCurAssets:
				if val, ok := item[i].(float64); ok {
					sheet.OtherCurrentAssert = val
				}
			case fieldTotalCurAssets:
				if val, ok := item[i].(float64); ok {
					sheet.TotalCurrentAsserts = val
				}
			case fieldFaAvailForSale:
				if val, ok := item[i].(float64); ok {
					sheet.AvailableForSaleFinancialAssets = val
				}
			case fieldHtmInvest:
				if val, ok := item[i].(float64); ok {
					sheet.HeldToMaturityInvestment = val
				}
			case fieldLtEqtInvest:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermEquityInvestment = val
				}
			case fieldInvestRealEstate:
				if val, ok := item[i].(float64); ok {
					sheet.InvestmentProperty = val
				}
			case fieldTimeDeposits:
				if val, ok := item[i].(float64); ok {
					sheet.TimeDeposit = val
				}
			case fieldOthAssets:
				if val, ok := item[i].(float64); ok {
					sheet.OtherAssets = val
				}
			case fieldLtRec:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermReceivables = val
				}
			case fieldFixAssets:
				if val, ok := item[i].(float64); ok {
					sheet.FixedAsset = val
				}
			case fieldCip:
				if val, ok := item[i].(float64); ok {
					sheet.ConstructionInProcess = val
				}
			case fieldConstMaterials:
				if val, ok := item[i].(float64); ok {
					sheet.ProjectGoodsAndMaterial = val
				}
			case fieldFixedAssetsDisp:
				if val, ok := item[i].(float64); ok {
					sheet.DisposalOfFixedAsset = val
				}
			case fieldProducBioAssets:
				if val, ok := item[i].(float64); ok {
					sheet.BiologicalAssets = val
				}
			case fieldOilAndGasAssets:
				if val, ok := item[i].(float64); ok {
					sheet.OilAndGasAssets = val
				}
			case fieldIntanAssets:
				if val, ok := item[i].(float64); ok {
					sheet.IntangibleAssets = val
				}
			case fieldRAndD:
				if val, ok := item[i].(float64); ok {
					sheet.ResearchAndDevelopmentExpenditure = val
				}
			case fieldGoodwill:
				if val, ok := item[i].(float64); ok {
					sheet.Goodwill = val
				}
			case fieldLtAmorExp:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermDeferredExpense = val
				}
			case fieldDeferTaxAssets:
				if val, ok := item[i].(float64); ok {
					sheet.DeferredTaxAsset = val
				}
			case fieldDecrInDisbur:
				if val, ok := item[i].(float64); ok {
					sheet.LoansAndPaymentsOnBehalf = val
				}
			case fieldOthNca:
				if val, ok := item[i].(float64); ok {
					sheet.OtherNonCurrentAssets = val
				}
			case fieldTotalNca:
				if val, ok := item[i].(float64); ok {
					sheet.TotalNonCurrentAssets = val
				}
			case fieldCashReserCb:
				if val, ok := item[i].(float64); ok {
					sheet.CashAndBalancesWithCentralBank = val
				}
			case fieldDeposInOthBfi:
				if val, ok := item[i].(float64); ok {
					sheet.DepositsInOtherBanks = val
				}
			case fieldPrecMetals:
				if val, ok := item[i].(float64); ok {
					sheet.PreciousMetals = val
				}
			case fieldDerivAssets:
				if val, ok := item[i].(float64); ok {
					sheet.DerivativeFinancialInstruments = val
				}
			case fieldRrReinsUnePrem:
				if val, ok := item[i].(float64); ok {
					sheet.ReceivableDepositForUndueDutyOfReinsurance = val
				}
			case fieldRrReinsOutstdCla:
				if val, ok := item[i].(float64); ok {
					sheet.ReinsurersShareOfClaimReserves = val
				}
			case fieldRrReinsLinsLiab:
				if val, ok := item[i].(float64); ok {
					sheet.ReinsurersShareOfLifeInsuranceReserves = val
				}
			case fieldRrReinsLthinsLiab:
				if val, ok := item[i].(float64); ok {
					sheet.ReinsurersShareOfLongTermHealthInsuranceReserves = val
				}
			case fieldRefundDepos:
				if val, ok := item[i].(float64); ok {
					sheet.RefundDeposits = val
				}
			case fieldPhPledgeLoans:
				if val, ok := item[i].(float64); ok {
					sheet.PolicyPledgeLoans = val
				}
			case fieldRefundCapDepos:
				if val, ok := item[i].(float64); ok {
					sheet.RefundCapitalDeposits = val
				}
			case fieldIndepAcctAssets:
				if val, ok := item[i].(float64); ok {
					sheet.AssetsForIndependenceAccount = val
				}
			case fieldClientDepos:
				continue
			case fieldClientProv:
				continue
			case fieldTransacSeatFee:
				continue
			case fieldInvestAsReceiv:
				if val, ok := item[i].(float64); ok {
					sheet.InvestmentAsReceivables = val
				}
			case fieldTotalAssets:
				if val, ok := item[i].(float64); ok {
					sheet.TotalAssets = val
				}
			case fieldLtBorr:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermLoans = val
				}
			case fieldStBorr:
				if val, ok := item[i].(float64); ok {
					sheet.ShortTermLoans = val
				}
			case fieldCbBorr:
				if val, ok := item[i].(float64); ok {
					sheet.BorrowingFromTheCentralBank = val
				}
			case fieldDeposIbDeposits:
				if val, ok := item[i].(float64); ok {
					sheet.DepositsFromCustomersAndInterBank = val
				}
			case fieldLoanOthBank:
				if val, ok := item[i].(float64); ok {
					sheet.LoansFromOtherBanks = val
				}
			case fieldTradingFl:
				if val, ok := item[i].(float64); ok {
					sheet.TradingFinancialLiabilities = val
				}
			case fieldNotesPayable:
				if val, ok := item[i].(float64); ok {
					sheet.NotesPayable = val
				}
			case fieldAcctPayable:
				if val, ok := item[i].(float64); ok {
					sheet.AccountsPayable = val
				}
			case fieldAdvReceipts:
				if val, ok := item[i].(float64); ok {
					sheet.AdvancesFromCustomers = val
				}
			case fieldSoldForRepurFa:
				if val, ok := item[i].(float64); ok {
					sheet.SellsBuysTheFinancialPropertyFunds = val
				}
			case fieldCommPayable:
				if val, ok := item[i].(float64); ok {
					sheet.HandlingChargesAndCommissionsPayable = val
				}
			case fieldPayrollPayable:
				if val, ok := item[i].(float64); ok {
					sheet.StaffSalaries = val
				}
			case fieldTaxesPayable:
				if val, ok := item[i].(float64); ok {
					sheet.TaxesPayable = val
				}
			case fieldIntPayable:
				if val, ok := item[i].(float64); ok {
					sheet.InterestPayable = val
				}
			case fieldDivPayable:
				if val, ok := item[i].(float64); ok {
					sheet.DividendsPayable = val
				}
			case fieldOthPayable:
				if val, ok := item[i].(float64); ok {
					sheet.OtherPayable = val
				}
			case fieldAccExp:
				if val, ok := item[i].(float64); ok {
					sheet.AccruedExpenses = val
				}
			case fieldDeferredInc:
				if val, ok := item[i].(float64); ok {
					sheet.DeferredRevenue = val
				}
			case fieldStBondsPayable:
				if val, ok := item[i].(float64); ok {
					sheet.ShortTermDebenturesPayable = val
				}
			case fieldPayableToReinsurer:
				if val, ok := item[i].(float64); ok {
					sheet.ThePayableReinsurance = val
				}
			case fieldRsrvInsurCont:
				if val, ok := item[i].(float64); ok {
					sheet.ProvisionForInsuranceContracts = val
				}
			case fieldActingTradingSec:
				if val, ok := item[i].(float64); ok {
					sheet.ActingTradingSecurities = val
				}
			case fieldActingUwSec:
				if val, ok := item[i].(float64); ok {
					sheet.SecuritiesUnderwritingBrokerageDeposits = val
				}
			case fieldNonCurLiabDue1y:
				if val, ok := item[i].(float64); ok {
					sheet.NonCurrentLiabilitiesMaturingWithinOneYear = val
				}
			case fieldOthCurLiab:
				if val, ok := item[i].(float64); ok {
					sheet.OtherCurrentLiability = val
				}
			case fieldTotalCurLiab:
				if val, ok := item[i].(float64); ok {
					sheet.TotalCurrentLiability = val
				}
			case fieldBondPayable:
				if val, ok := item[i].(float64); ok {
					sheet.BondPayable = val
				}
			case fieldLtPayable:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermPayable = val
				}
			case fieldSpecificPayables:
				if val, ok := item[i].(float64); ok {
					sheet.SpecificPayable = val
				}
			case fieldEstimatedLiab:
				if val, ok := item[i].(float64); ok {
					sheet.EstimatedLiabilities = val
				}
			case fieldDeferTaxLiab:
				if val, ok := item[i].(float64); ok {
					sheet.DeferTaxLiabilities = val
				}
			case fieldDeferIncNonCurLiab:
				if val, ok := item[i].(float64); ok {
					sheet.DeferIncomeNonCurLiabilities = val
				}
			case fieldOthNcl:
				if val, ok := item[i].(float64); ok {
					sheet.OtherNonCurrentLiability = val
				}
			case fieldTotalNcl:
				if val, ok := item[i].(float64); ok {
					sheet.TotalNonCurrentLiability = val
				}
			case fieldDeposOthBfi:
				if val, ok := item[i].(float64); ok {
					sheet.DueToBanksAndOtherFinancialInstitutions = val
				}
			case fieldDerivLiab:
				if val, ok := item[i].(float64); ok {
					sheet.DerivativeFinancialLiabilities = val
				}
			case fieldDepos:
				if val, ok := item[i].(float64); ok {
					sheet.Deposits = val
				}
			case fieldAgencyBusLiab:
				if val, ok := item[i].(float64); ok {
					sheet.AgencyBusinessLiabilities = val
				}
			case fieldOthLiab:
				if val, ok := item[i].(float64); ok {
					sheet.OtherLiabilities = val
				}
			case fieldPremReceivAdva:
				if val, ok := item[i].(float64); ok {
					sheet.PremiumsReceivedInAdvance = val
				}
			case fieldDeposReceived:
				if val, ok := item[i].(float64); ok {
					sheet.DepositsReceived = val
				}
			case fieldPhInvest:
				if val, ok := item[i].(float64); ok {
					sheet.DepositsFromPolicyHolders = val
				}
			case fieldReserUnePrem:
				if val, ok := item[i].(float64); ok {
					sheet.UnearnedPremiumReserve = val
				}
			case fieldReserOutstdClaims:
				if val, ok := item[i].(float64); ok {
					sheet.ReserveForOutstandingLosses = val
				}
			case fieldReserLinsLiab:
				if val, ok := item[i].(float64); ok {
					sheet.LifeInsuranceReserves = val
				}
			case fieldReserLthinsLiab:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermHealthInsuranceReserves = val
				}
			case fieldIndeptAccLiab:
				if val, ok := item[i].(float64); ok {
					sheet.LiabilitiesOfIndependentAccounts = val
				}
			case fieldPledgeBorr:
				continue
			case fieldIndemPayable:
				if val, ok := item[i].(float64); ok {
					sheet.ClaimsPayable = val
				}
			case fieldPolicyDivPayable:
				if val, ok := item[i].(float64); ok {
					sheet.PolicyDividendPayable = val
				}
			case fieldTOtalLiab:
				if val, ok := item[i].(float64); ok {
					sheet.TotalLiability = val
				}
			case fieldTreasuryShare:
				if val, ok := item[i].(float64); ok {
					sheet.TreasuryShare = val
				}
			case fieldOrdinRiskReser:
				if val, ok := item[i].(float64); ok {
					sheet.GeneralRiskPreparation = val
				}
			case fieldForexDiffer:
				if val, ok := item[i].(float64); ok {
					sheet.ConvertedDifferenceInForeignCurrencyStatements = val
				}
			case fieldInvestLossUnconf:
				if val, ok := item[i].(float64); ok {
					sheet.UnrealisedInvestmentLosses = val
				}
			case fieldMinorityInt:
				if val, ok := item[i].(float64); ok {
					sheet.MinorityInterests = val
				}
			case fieldTotalHldrEqyExcMinInt:
				if val, ok := item[i].(float64); ok {
					sheet.TotalEquityOfOwnersExcludeMinorityInterests = val
				}
			case fieldTotalHldrEqyIncMinInt:
				if val, ok := item[i].(float64); ok {
					sheet.TotalEquityOfOwnersIncludeMinorityInterests = val
				}
			case fieldTotalLiabHldrEqy:
				if val, ok := item[i].(float64); ok {
					sheet.TotalLiabilitiesAndOwnersEquity = val
				}
			case fieldLtPayrollPayable:
				if val, ok := item[i].(float64); ok {
					sheet.LongTermStaffSalaries = val
				}
			case fieldOthCompIncome:
				if val, ok := item[i].(float64); ok {
					sheet.OtherComprehensiveIncome = val
				}
			case fieldOthEqtTools:
				if val, ok := item[i].(float64); ok {
					sheet.OtherEquityTools = val
				}
			case fieldOthEqtToolsPShr:
				if val, ok := item[i].(float64); ok {
					sheet.OtherEquityToolsPreferredStock = val
				}
			case fieldLendingFunds:
				if val, ok := item[i].(float64); ok {
					sheet.LendingFunds = val
				}
			case fieldAccReceivable:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.AccountsReceivable) {
						sheet.AccountsReceivable = val
					}
				}
			case fieldStFinPayable:
				if val, ok := item[i].(float64); ok {
					sheet.ShortTermFinancingPayable = val
				}
			case fieldPayables:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.AccountsPayable) {
						sheet.AccountsPayable = val
					}
				}
			case fieldHfsAssets:
				if val, ok := item[i].(float64); ok {
					sheet.AssetsHeldForSale = val
				}
			case fieldHfsSales:
				if val, ok := item[i].(float64); ok {
					sheet.LiabilitiesHeldForSale = val
				}
			case fieldCostFinAssets:
				if val, ok := item[i].(float64); ok {
					sheet.FinancialAssetsMeasuredAtAmortizedCost = val
				}
			case fieldFairValueFinAssets:
				if val, ok := item[i].(float64); ok {
					sheet.FairValueThroughOtherComprehensiveIncome = val
				}
			case fieldCipTotal:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.ConstructionInProcess) {
						sheet.ConstructionInProcess = val
					}
				}
			case fieldOthPayTotal:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.OtherPayable) {
						sheet.OtherPayable = val
					}
				}
			case fieldLongPayTotal:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.LongTermPayable) {
						sheet.LongTermPayable = val
					}
				}
			case fieldDebtInvest:
				if val, ok := item[i].(float64); ok {
					sheet.DebtInvestment = val
				}
			case fieldOthDebtInvest:
				if val, ok := item[i].(float64); ok {
					sheet.OtherDebtInvestment = val
				}
			case fieldOthEqInvest:
				if val, ok := item[i].(float64); ok {
					sheet.OtherEquityInvestment = val
				}
			case fieldOthIlliqFinAssets:
				if val, ok := item[i].(float64); ok {
					sheet.OtherNonCurrentFinancialAssets = val
				}
			case fieldOthEqPpbond:
				if val, ok := item[i].(float64); ok {
					sheet.OtherEquityToolsPerpetualDebt = val
				}
			case fieldReceivFinancing:
				if val, ok := item[i].(float64); ok {
					sheet.ReceivableFinancing = val
				}
			case fieldUseRightAssets:
				if val, ok := item[i].(float64); ok {
					sheet.UseRightAssets = val
				}
			case fieldLeaseLiab:
				if val, ok := item[i].(float64); ok {
					sheet.LeaseLiabilities = val
				}
			case fileContractAssets:
				if val, ok := item[i].(float64); ok {
					sheet.ContractAssets = val
				}
			case fieldContractLiab:
				if val, ok := item[i].(float64); ok {
					sheet.ContractLiabilities = val
				}
			case fieldAccountsReceivBill:
				if val, ok := item[i].(float64); ok {
					sheet.AccountsReceivableAndBill = val
				}
			case fieldAccountsPay:
				if val, ok := item[i].(float64); ok {
					sheet.AccountsPayableAndBill = val
				}
			case fieldOthRcvTotal:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.OtherReceivables) {
						sheet.OtherReceivables = val
					}
				}
			case fieldFixAssetsTotal:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, sheet.FixedAsset) {
						sheet.FixedAsset = val
					}
				}
			default:
				continue
			}
		}
		sheets = append(sheets, sheet)
	}
	if len(sheets) == 0 {
		return nil, fmt.Errorf("can not fetch balance sheet from tushare")
	}
	return findLastBalanceSheet(sheets), nil
}

func findLastBalanceSheet(sheets []model.BalanceSheet) *model.BalanceSheet {
	return &sheets[len(sheets)-1]
}

func isNeedOverWrite(new, old float64) bool {
	return new > 1.0 || old < 1.0
}
