package tushare

import (
	"anguo/domain/common"
	"anguo/model"
	"fmt"
	"strings"
)

func GetIncomeStatementFromTushare(code string) (*model.IncomeStatement, error) {
	if code == "" {
		return nil, fmt.Errorf("emtpy code when fetch income statement")
	}
	return GetIncomeStatementFromTushareForGivenCodeAndDate(code, common.GetLastYearEndDate())
}

func GetIncomeStatementFromTushareForGivenCodeAndDate(code, date string) (*model.IncomeStatement, error) {
	fields := strings.Join([]string{
		fieldUpdateFlag,
		fieldBasicEps,
		fieldDilutedEps,
		filedTotalRevenue,
		fieldRevenue,
		fieldIntIncome,
		fieldPremEarned,
		fieldCommIncome,
		fieldNCommisIncome,
		fieldNOthIncome,
		filedNOthBIncome,
		fieldPremIncome,
		fieldOutPrem,
		fieldUnePremReser,
		fieldReinsIncome,
		fieldNSecTbIncome,
		fieldNSecUwIncome,
		fieldNAssetMgIncome,
		fieldOthBIncome,
		fieldFvValueChgGain,
		fieldInvestIncome,
		fieldAssInvestIncome,
		fieldForexAgain,
		fieldTotalCogs,
		fieldOperCost,
		fieldIntExp,
		fieldCommExp,
		fieldBizTaxSurChg,
		fieldSellExp,
		fieldAdminExp,
		fieldFinExp,
		fieldAssetsImpairLoss,
		fieldPremRefund,
		fieldCompensPayout,
		fieldReserInsurLiab,
		fieldDivPayt,
		fieldReinsExp,
		fieldOperExp,
		fieldCompensPayoutRefu,
		fieldInsurReserRefu,
		fieldReinsCostRefund,
		fieldOtherBusCost,
		fieldOperateProfit,
		fieldNonOperIncome,
		fieldNonOperExp,
		fieldNcaDisploss,
		fieldTotalProfit,
		fieldIncomeTax,
		fieldNIncome,
		fieldNIncomeAttrP,
		fieldMinorityGain,
		fieldOthComprIncome,
		fieldTComprIncome,
		fieldComprIncAttrP,
		fieldComprIncAttrMS,
		fieldEbit,
		fieldEbitda,
		fieldInsuranceExp,
		fieldUndistProfit,
		fieldDistableProfit,
		fieldRdExp,
		fieldFinExpIntExp,
		fieldFinExpIntInc,
		fieldTransferSurplusRese,
		fieldTransferHousingImprest,
		fieldTransferOth,
		fieldAdjLossGain,
		fieldWithdraLegalSurplus,
		fieldWithdraLegalPubFund,
		fieldWithdraBizDevFund,
		fieldWithdraReseFund,
		fieldWithdraOthErsu,
		fieldWorkersWelfare,
		fieldDistrProfitShrhder,
		fieldPrfSharePayableDvd,
		fieldComsharePayableDvd,
		fieldCapitComstockDiv,
		fieldNetAfterNrLpCorrect,
		fieldCreditImpaLoss,
		fieldNetExpoHedgingBenefits,
		fieldOthImpairLossAssets,
		fieldTotalOpCost,
		fieldAmodcostFinAssets,
		fieldOthIncome,
		fieldAssetDispIncome,
		fieldContinuedNetProfit,
		fieldEndNetProfit,
		fieldEndDate,
	}, ",")
	params := map[string]interface{}{
		fieldTsCode: code,
	}
	resp, err := fetchTushareRawData("income", fields, params)
	if err != nil {
		return nil, err
	}
	return parseIncomeStatementRecord(date, resp)
}

func parseIncomeStatementRecord(endDate string, resp *Response) (*model.IncomeStatement, error) {
	err := resp.anyError()
	if err != nil {
		return nil, err
	}
	var statements []model.IncomeStatement
	for _, item := range resp.Data.Items {
		var statement = model.IncomeStatement{}
		for i, field := range resp.Data.Fields {
			switch field {
			case fieldEndDate:
				if str, ok := item[i].(string); ok {
					statement.EndDate = str
				}
			case fieldTsCode:
				if str, ok := item[i].(string); ok {
					statement.Code = str
				}
			case fieldBasicEps:
				continue
			case fieldDilutedEps:
				continue
			case filedTotalRevenue:
				if val, ok := item[i].(float64); ok {
					statement.TotalRevenue = val
				}
			case fieldRevenue:
				if val, ok := item[i].(float64); ok {
					statement.Revenue = val
				}
			case fieldIntIncome:
				if val, ok := item[i].(float64); ok {
					statement.InterestsIncome = val
				}
			case fieldPremEarned:
				if val, ok := item[i].(float64); ok {
					statement.PremiumEarned = val
				}
			case fieldCommIncome:
				if val, ok := item[i].(float64); ok {
					statement.HandlingChargesAndCommissionsIncome = val
				}
			case fieldNCommisIncome:
				if val, ok := item[i].(float64); ok {
					statement.HandlingChargesAndCommissionsNetIncome = val
				}
			case fieldNOthIncome:
				if val, ok := item[i].(float64); ok {
					statement.OtherNetIncome = val
				}
			case filedNOthBIncome:
				if val, ok := item[i].(float64); ok {
					statement.OtherBusinessNetIncome = val
				}
			case fieldPremIncome:
				if val, ok := item[i].(float64); ok {
					statement.PremiumIncome = val
				}
			case fieldOutPrem:
				if val, ok := item[i].(float64); ok {
					statement.OutPremium = val
				}
			case fieldUnePremReser:
				if val, ok := item[i].(float64); ok {
					statement.AppropriationOfDepositForUndueDuty = val
				}
			case fieldReinsIncome:
				if val, ok := item[i].(float64); ok {
					statement.ReinsurancePremiumIncome = val
				}
			case fieldNSecTbIncome:
				if val, ok := item[i].(float64); ok {
					statement.NetIncomeFromSecuritiesTradingBrokerageBusiness = val
				}
			case fieldNSecUwIncome:
				if val, ok := item[i].(float64); ok {
					statement.NetIncomeFromSecuritiesUnderwritingBusiness = val
				}
			case fieldNAssetMgIncome:
				if val, ok := item[i].(float64); ok {
					statement.NetIncomeFromEntrustedCustomerAssetsManagementBusiness = val
				}
			case fieldOthBIncome:
				if val, ok := item[i].(float64); ok {
					statement.OtherBusinessIncome = val
				}
			case fieldFvValueChgGain:
				if val, ok := item[i].(float64); ok {
					statement.ProfitAndLossOnFluctuationOfFairMarketValue = val
				}
			case fieldInvestIncome:
				if val, ok := item[i].(float64); ok {
					statement.InvestmentIncome = val
				}
			case fieldAssInvestIncome:
				if val, ok := item[i].(float64); ok {
					statement.InvestmentIncomeFromAssociatesAndJointVentures = val
				}
			case fieldForexAgain:
				if val, ok := item[i].(float64); ok {
					statement.ExchangeNetIncome = val
				}
			case fieldTotalCogs:
				if val, ok := item[i].(float64); ok {
					statement.TotalOperatingCost = val
				}
			case fieldOperCost:
				if val, ok := item[i].(float64); ok {
					statement.OperatingCost = val
				}
			case fieldIntExp:
				if val, ok := item[i].(float64); ok {
					statement.InterestExpense = val
				}
			case fieldCommExp:
				if val, ok := item[i].(float64); ok {
					statement.HandlingChargesAndCommissionsExpense = val
				}
			case fieldBizTaxSurChg:
				if val, ok := item[i].(float64); ok {
					statement.BusinessTaxesAndSurcharges = val
				}
			case fieldSellExp:
				if val, ok := item[i].(float64); ok {
					statement.SellExpense = val
				}
			case fieldAdminExp:
				if val, ok := item[i].(float64); ok {
					statement.AdministrationExpenses = val
				}
			case fieldFinExp:
				if val, ok := item[i].(float64); ok {
					statement.FinancialExpenses = val
				}
			case fieldAssetsImpairLoss:
				if val, ok := item[i].(float64); ok {
					statement.AssetImpairmentLoss = val
				}
			case fieldPremRefund:
				if val, ok := item[i].(float64); ok {
					statement.PremiumRefund = val
				}
			case fieldCompensPayout:
				if val, ok := item[i].(float64); ok {
					statement.CompensationExpenditure = val
				}
			case fieldReserInsurLiab:
				if val, ok := item[i].(float64); ok {
					statement.AppropriationOfDepositForDuty = val
				}
			case fieldDivPayt:
				if val, ok := item[i].(float64); ok {
					statement.DividendExpensesForTheInsured = val
				}
			case fieldReinsExp:
				if val, ok := item[i].(float64); ok {
					statement.ReinsuranceExpense = val
				}
			case fieldOperExp:
				if val, ok := item[i].(float64); ok {
					statement.OperatingExpense = val
				}
			case fieldCompensPayoutRefu:
				if val, ok := item[i].(float64); ok {
					statement.AmortizedCompensationExpenses = val
				}
			case fieldInsurReserRefu:
				if val, ok := item[i].(float64); ok {
					statement.AmortizedDepositForDuty = val
				}
			case fieldReinsCostRefund:
				if val, ok := item[i].(float64); ok {
					statement.AmortizedReinsuranceExpense = val
				}
			case fieldOtherBusCost:
				if val, ok := item[i].(float64); ok {
					statement.OtherBusinessCost = val
				}
			case fieldOperateProfit:
				if val, ok := item[i].(float64); ok {
					statement.OperateProfit = val
				}
			case fieldNonOperIncome:
				if val, ok := item[i].(float64); ok {
					statement.NonOperateIncome = val
				}
			case fieldNonOperExp:
				if val, ok := item[i].(float64); ok {
					statement.NonOperateExpense = val
				}
			case fieldNcaDisploss:
				if val, ok := item[i].(float64); ok {
					statement.NonCurrentAssetsDisposalNetLoss = val
				}
			case fieldTotalProfit:
				if val, ok := item[i].(float64); ok {
					statement.TotalProfit = val
				}
			case fieldIncomeTax:
				if val, ok := item[i].(float64); ok {
					statement.IncomeTax = val
				}
			case fieldNIncome:
				if val, ok := item[i].(float64); ok {
					statement.NetIncome = val
				}
			case fieldNIncomeAttrP:
				if val, ok := item[i].(float64); ok {
					statement.NetIncomeBelongsToOwner = val
				}
			case fieldMinorityGain:
				if val, ok := item[i].(float64); ok {
					statement.MinorityGain = val
				}
			case fieldOthComprIncome:
				if val, ok := item[i].(float64); ok {
					statement.OtherComprehensiveIncome = val
				}
			case fieldTComprIncome:
				if val, ok := item[i].(float64); ok {
					statement.TotalComprehensiveIncome = val
				}
			case fieldComprIncAttrP:
				if val, ok := item[i].(float64); ok {
					statement.TotalComprehensiveIncomeBelongToOwner = val
				}
			case fieldComprIncAttrMS:
				if val, ok := item[i].(float64); ok {
					statement.TotalComprehensiveIncomeBelongToMinority = val
				}
			case fieldEbit:
				continue
			case fieldEbitda:
				continue
			case fieldInsuranceExp:
				if val, ok := item[i].(float64); ok {
					statement.InsuranceExpense = val
				}
			case fieldUndistProfit:
				if val, ok := item[i].(float64); ok {
					statement.UndistributedProfitOfBeginning = val
				}
			case fieldDistableProfit:
				if val, ok := item[i].(float64); ok {
					statement.DistributiveProfits = val
				}
			case fieldRdExp:
				if val, ok := item[i].(float64); ok {
					statement.ResearchAndDevelopExpense = val
				}
			case fieldFinExpIntExp:
				if val, ok := item[i].(float64); ok {
					statement.InterestExpenseInFinancialExpense = val
				}
			case fieldFinExpIntInc:
				if val, ok := item[i].(float64); ok {
					statement.InterestIncomeInFinancialExpense = val
				}
			case fieldTransferSurplusRese:
				if val, ok := item[i].(float64); ok {
					statement.TransferInFromSurplusReserve = val
				}
			case fieldTransferHousingImprest:
				if val, ok := item[i].(float64); ok {
					statement.TransferInFromHousingRevolvingFund = val
				}
			case fieldTransferOth:
				if val, ok := item[i].(float64); ok {
					statement.TransferInFromOther = val
				}
			case fieldAdjLossGain:
				if val, ok := item[i].(float64); ok {
					statement.BeforehandYearProfitAndLossAdjustment = val
				}
			case fieldWithdraLegalSurplus:
				if val, ok := item[i].(float64); ok {
					statement.WithdrawLegalSurplus = val
				}
			case fieldWithdraLegalPubFund:
				if val, ok := item[i].(float64); ok {
					statement.WithdrawLegalPublicWelfareFund = val
				}
			case fieldWithdraBizDevFund:
				if val, ok := item[i].(float64); ok {
					statement.WithdrawEnterpriseDevelopmentFund = val
				}
			case fieldWithdraReseFund:
				if val, ok := item[i].(float64); ok {
					statement.WithdrawReserveFund = val
				}
			case fieldWithdraOthErsu:
				if val, ok := item[i].(float64); ok {
					statement.WithdrawOtherSurplus = val
				}
			case fieldWorkersWelfare:
				if val, ok := item[i].(float64); ok {
					statement.WorkersWelfare = val
				}
			case fieldDistrProfitShrhder:
				if val, ok := item[i].(float64); ok {
					statement.ProfitAvailableForDistributionByShareholders = val
				}
			case fieldPrfSharePayableDvd:
				if val, ok := item[i].(float64); ok {
					statement.PreferredStockDividendsPayable = val
				}
			case fieldComsharePayableDvd:
				if val, ok := item[i].(float64); ok {
					statement.CommonStockDividendsPayable = val
				}
			case fieldCapitComstockDiv:
				if val, ok := item[i].(float64); ok {
					statement.CommonStockDividendsConvertedIntoCapitalStock = val
				}
			case fieldNetAfterNrLpCorrect:
				if val, ok := item[i].(float64); ok {
					statement.NonRecurringGainsAndLosses = val
				}
			case fieldCreditImpaLoss:
				if val, ok := item[i].(float64); ok {
					statement.CreditImpairmentLoss = val
				}
			case fieldNetExpoHedgingBenefits:
				if val, ok := item[i].(float64); ok {
					statement.NetExposureHedgingIncome = val
				}
			case fieldOthImpairLossAssets:
				if val, ok := item[i].(float64); ok {
					statement.OtherAssertImpairmentLoss = val
				}
			case fieldTotalOpCost:
				if val, ok := item[i].(float64); ok {
					if isNeedOverWrite(val, statement.TotalOperatingCost) {
						statement.TotalOperatingCost = val
					}
				}
			case fieldAmodcostFinAssets:
				if val, ok := item[i].(float64); ok {
					statement.IncomeFromStopOfFinancialAssetsMeasuredAtAmortizedCost = val
				}
			case fieldOthIncome:
				if val, ok := item[i].(float64); ok {
					statement.OtherIncome = val
				}
			case fieldAssetDispIncome:
				if val, ok := item[i].(float64); ok {
					statement.AssertDisposalIncome = val
				}
			case fieldContinuedNetProfit:
				if val, ok := item[i].(float64); ok {
					statement.ContinuedNetProfit = val
				}
			case fieldEndNetProfit:
				if val, ok := item[i].(float64); ok {
					statement.EndNetProfit = val
				}
			default:
				continue
			}
		}
		statements = append(statements, statement)
	}
	if len(statements) == 0 {
		return nil, fmt.Errorf("can not fetch balance statement from tushare")
	}
	return findLastIncomeStatement(endDate, statements), nil
}

func findLastIncomeStatement(endData string, statements []model.IncomeStatement) *model.IncomeStatement {
	buf := make([]model.IncomeStatement, 0)
	for _, states := range statements {
		// end_date in the params CAN NOT filter the statements, we must filter it manually
		if states.EndDate == endData {
			buf = append(buf, states)
		}
	}
	return &buf[0]
}
