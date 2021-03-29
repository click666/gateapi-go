# Go API client for gateapi

Welcome to Gate.io API

APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.20.1
- Package version: 5.20.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.gate.io/page/contacts](https://www.gate.io/page/contacts)

## Versioning

Trying our best to follow the [semantic versioning](https://semver.org/), while enjoying recent features
provided by programming language and libraries, from 4.15.2, one major versioning difference will be
introduced:

If extra code rewrite is required when you upgrade the SDK, such as:

- some outdated programming language version support is dropped
- API method signature has breaking changes.

**the MAJOR version will be incremented, but the MINOR and PATCH version are still following REST API's
instead of resetting to 0**, so that you can recognize it has some breaking changes, but still getting
the idea of from which REST API version the change is introduced.

For example, the previous REST API and SDK version are both 4.14.0. But if we decide to introduce
some breaking changes in SDK along with REST API 4.15.2 upgrade, then the version of next SDK release
will be 5.15.2(the MAJOR version is incremented to denote breaking changes, but the MINOR and PATCH
version are identical to REST API's instead of resetting them to 0)

If MAJOR version is incremented, make sure you read the release note on
[Releases](https://github.com/gateio/gateapi-go/releases)
page

## Installation

```shell
go get github.com/gateio/gateapi-go/v5
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:


```golang
package main

import (
    "context"
    "fmt"

    "github.com/gateio/gateapi-go/v5"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.Background()
    settle := "usdt" // string - Settle currency
    
    result, _, err := client.DeliveryApi.ListDeliveryContracts(ctx, settle)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```



For a more complete API usage example, refer to the demo application in [example](_example) directory

## Documentation for API Endpoints

All URIs are relative to *https://api.gateio.ws/api/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeliveryApi* | [**ListDeliveryContracts**](docs/DeliveryApi.md#listdeliverycontracts) | **Get** /delivery/{settle}/contracts | List all futures contracts
*DeliveryApi* | [**GetDeliveryContract**](docs/DeliveryApi.md#getdeliverycontract) | **Get** /delivery/{settle}/contracts/{contract} | Get a single contract
*DeliveryApi* | [**ListDeliveryOrderBook**](docs/DeliveryApi.md#listdeliveryorderbook) | **Get** /delivery/{settle}/order_book | Futures order book
*DeliveryApi* | [**ListDeliveryTrades**](docs/DeliveryApi.md#listdeliverytrades) | **Get** /delivery/{settle}/trades | Futures trading history
*DeliveryApi* | [**ListDeliveryCandlesticks**](docs/DeliveryApi.md#listdeliverycandlesticks) | **Get** /delivery/{settle}/candlesticks | Get futures candlesticks
*DeliveryApi* | [**ListDeliveryTickers**](docs/DeliveryApi.md#listdeliverytickers) | **Get** /delivery/{settle}/tickers | List futures tickers
*DeliveryApi* | [**ListDeliveryInsuranceLedger**](docs/DeliveryApi.md#listdeliveryinsuranceledger) | **Get** /delivery/{settle}/insurance | Futures insurance balance history
*DeliveryApi* | [**ListDeliveryAccounts**](docs/DeliveryApi.md#listdeliveryaccounts) | **Get** /delivery/{settle}/accounts | Query futures account
*DeliveryApi* | [**ListDeliveryAccountBook**](docs/DeliveryApi.md#listdeliveryaccountbook) | **Get** /delivery/{settle}/account_book | Query account book
*DeliveryApi* | [**ListDeliveryPositions**](docs/DeliveryApi.md#listdeliverypositions) | **Get** /delivery/{settle}/positions | List all positions of a user
*DeliveryApi* | [**GetDeliveryPosition**](docs/DeliveryApi.md#getdeliveryposition) | **Get** /delivery/{settle}/positions/{contract} | Get single position
*DeliveryApi* | [**UpdateDeliveryPositionMargin**](docs/DeliveryApi.md#updatedeliverypositionmargin) | **Post** /delivery/{settle}/positions/{contract}/margin | Update position margin
*DeliveryApi* | [**UpdateDeliveryPositionLeverage**](docs/DeliveryApi.md#updatedeliverypositionleverage) | **Post** /delivery/{settle}/positions/{contract}/leverage | Update position leverage
*DeliveryApi* | [**UpdateDeliveryPositionRiskLimit**](docs/DeliveryApi.md#updatedeliverypositionrisklimit) | **Post** /delivery/{settle}/positions/{contract}/risk_limit | Update position risk limit
*DeliveryApi* | [**ListDeliveryOrders**](docs/DeliveryApi.md#listdeliveryorders) | **Get** /delivery/{settle}/orders | List futures orders
*DeliveryApi* | [**CreateDeliveryOrder**](docs/DeliveryApi.md#createdeliveryorder) | **Post** /delivery/{settle}/orders | Create a futures order
*DeliveryApi* | [**CancelDeliveryOrders**](docs/DeliveryApi.md#canceldeliveryorders) | **Delete** /delivery/{settle}/orders | Cancel all &#x60;open&#x60; orders matched
*DeliveryApi* | [**GetDeliveryOrder**](docs/DeliveryApi.md#getdeliveryorder) | **Get** /delivery/{settle}/orders/{order_id} | Get a single order
*DeliveryApi* | [**CancelDeliveryOrder**](docs/DeliveryApi.md#canceldeliveryorder) | **Delete** /delivery/{settle}/orders/{order_id} | Cancel a single order
*DeliveryApi* | [**GetMyDeliveryTrades**](docs/DeliveryApi.md#getmydeliverytrades) | **Get** /delivery/{settle}/my_trades | List personal trading history
*DeliveryApi* | [**ListDeliveryPositionClose**](docs/DeliveryApi.md#listdeliverypositionclose) | **Get** /delivery/{settle}/position_close | List position close history
*DeliveryApi* | [**ListDeliveryLiquidates**](docs/DeliveryApi.md#listdeliveryliquidates) | **Get** /delivery/{settle}/liquidates | List liquidation history
*DeliveryApi* | [**ListDeliverySettlements**](docs/DeliveryApi.md#listdeliverysettlements) | **Get** /delivery/{settle}/settlements | List settlement history
*DeliveryApi* | [**ListPriceTriggeredDeliveryOrders**](docs/DeliveryApi.md#listpricetriggereddeliveryorders) | **Get** /delivery/{settle}/price_orders | List all auto orders
*DeliveryApi* | [**CreatePriceTriggeredDeliveryOrder**](docs/DeliveryApi.md#createpricetriggereddeliveryorder) | **Post** /delivery/{settle}/price_orders | Create a price-triggered order
*DeliveryApi* | [**CancelPriceTriggeredDeliveryOrderList**](docs/DeliveryApi.md#cancelpricetriggereddeliveryorderlist) | **Delete** /delivery/{settle}/price_orders | Cancel all open orders
*DeliveryApi* | [**GetPriceTriggeredDeliveryOrder**](docs/DeliveryApi.md#getpricetriggereddeliveryorder) | **Get** /delivery/{settle}/price_orders/{order_id} | Get a single order
*DeliveryApi* | [**CancelPriceTriggeredDeliveryOrder**](docs/DeliveryApi.md#cancelpricetriggereddeliveryorder) | **Delete** /delivery/{settle}/price_orders/{order_id} | Cancel a single order
*FuturesApi* | [**ListFuturesContracts**](docs/FuturesApi.md#listfuturescontracts) | **Get** /futures/{settle}/contracts | List all futures contracts
*FuturesApi* | [**GetFuturesContract**](docs/FuturesApi.md#getfuturescontract) | **Get** /futures/{settle}/contracts/{contract} | Get a single contract
*FuturesApi* | [**ListFuturesOrderBook**](docs/FuturesApi.md#listfuturesorderbook) | **Get** /futures/{settle}/order_book | Futures order book
*FuturesApi* | [**ListFuturesTrades**](docs/FuturesApi.md#listfuturestrades) | **Get** /futures/{settle}/trades | Futures trading history
*FuturesApi* | [**ListFuturesCandlesticks**](docs/FuturesApi.md#listfuturescandlesticks) | **Get** /futures/{settle}/candlesticks | Get futures candlesticks
*FuturesApi* | [**ListFuturesTickers**](docs/FuturesApi.md#listfuturestickers) | **Get** /futures/{settle}/tickers | List futures tickers
*FuturesApi* | [**ListFuturesFundingRateHistory**](docs/FuturesApi.md#listfuturesfundingratehistory) | **Get** /futures/{settle}/funding_rate | Funding rate history
*FuturesApi* | [**ListFuturesInsuranceLedger**](docs/FuturesApi.md#listfuturesinsuranceledger) | **Get** /futures/{settle}/insurance | Futures insurance balance history
*FuturesApi* | [**ListContractStats**](docs/FuturesApi.md#listcontractstats) | **Get** /futures/{settle}/contract_stats | Futures stats
*FuturesApi* | [**ListLiquidatedOrders**](docs/FuturesApi.md#listliquidatedorders) | **Get** /futures/{settle}/liq_orders | Retrieve liquidation history
*FuturesApi* | [**ListFuturesAccounts**](docs/FuturesApi.md#listfuturesaccounts) | **Get** /futures/{settle}/accounts | Query futures account
*FuturesApi* | [**ListFuturesAccountBook**](docs/FuturesApi.md#listfuturesaccountbook) | **Get** /futures/{settle}/account_book | Query account book
*FuturesApi* | [**ListPositions**](docs/FuturesApi.md#listpositions) | **Get** /futures/{settle}/positions | List all positions of a user
*FuturesApi* | [**GetPosition**](docs/FuturesApi.md#getposition) | **Get** /futures/{settle}/positions/{contract} | Get single position
*FuturesApi* | [**UpdatePositionMargin**](docs/FuturesApi.md#updatepositionmargin) | **Post** /futures/{settle}/positions/{contract}/margin | Update position margin
*FuturesApi* | [**UpdatePositionLeverage**](docs/FuturesApi.md#updatepositionleverage) | **Post** /futures/{settle}/positions/{contract}/leverage | Update position leverage
*FuturesApi* | [**UpdatePositionRiskLimit**](docs/FuturesApi.md#updatepositionrisklimit) | **Post** /futures/{settle}/positions/{contract}/risk_limit | Update position risk limit
*FuturesApi* | [**SetDualMode**](docs/FuturesApi.md#setdualmode) | **Post** /futures/{settle}/dual_mode | Enable or disable dual mode
*FuturesApi* | [**GetDualModePosition**](docs/FuturesApi.md#getdualmodeposition) | **Get** /futures/{settle}/dual_comp/positions/{contract} | Retrieve position detail in dual mode
*FuturesApi* | [**UpdateDualModePositionMargin**](docs/FuturesApi.md#updatedualmodepositionmargin) | **Post** /futures/{settle}/dual_comp/positions/{contract}/margin | Update position margin in dual mode
*FuturesApi* | [**UpdateDualModePositionLeverage**](docs/FuturesApi.md#updatedualmodepositionleverage) | **Post** /futures/{settle}/dual_comp/positions/{contract}/leverage | Update position leverage in dual mode
*FuturesApi* | [**UpdateDualModePositionRiskLimit**](docs/FuturesApi.md#updatedualmodepositionrisklimit) | **Post** /futures/{settle}/dual_comp/positions/{contract}/risk_limit | Update position risk limit in dual mode
*FuturesApi* | [**ListFuturesOrders**](docs/FuturesApi.md#listfuturesorders) | **Get** /futures/{settle}/orders | List futures orders
*FuturesApi* | [**CreateFuturesOrder**](docs/FuturesApi.md#createfuturesorder) | **Post** /futures/{settle}/orders | Create a futures order
*FuturesApi* | [**CancelFuturesOrders**](docs/FuturesApi.md#cancelfuturesorders) | **Delete** /futures/{settle}/orders | Cancel all &#x60;open&#x60; orders matched
*FuturesApi* | [**GetFuturesOrder**](docs/FuturesApi.md#getfuturesorder) | **Get** /futures/{settle}/orders/{order_id} | Get a single order
*FuturesApi* | [**CancelFuturesOrder**](docs/FuturesApi.md#cancelfuturesorder) | **Delete** /futures/{settle}/orders/{order_id} | Cancel a single order
*FuturesApi* | [**GetMyTrades**](docs/FuturesApi.md#getmytrades) | **Get** /futures/{settle}/my_trades | List personal trading history
*FuturesApi* | [**ListPositionClose**](docs/FuturesApi.md#listpositionclose) | **Get** /futures/{settle}/position_close | List position close history
*FuturesApi* | [**ListLiquidates**](docs/FuturesApi.md#listliquidates) | **Get** /futures/{settle}/liquidates | List liquidation history
*FuturesApi* | [**ListPriceTriggeredOrders**](docs/FuturesApi.md#listpricetriggeredorders) | **Get** /futures/{settle}/price_orders | List all auto orders
*FuturesApi* | [**CreatePriceTriggeredOrder**](docs/FuturesApi.md#createpricetriggeredorder) | **Post** /futures/{settle}/price_orders | Create a price-triggered order
*FuturesApi* | [**CancelPriceTriggeredOrderList**](docs/FuturesApi.md#cancelpricetriggeredorderlist) | **Delete** /futures/{settle}/price_orders | Cancel all open orders
*FuturesApi* | [**GetPriceTriggeredOrder**](docs/FuturesApi.md#getpricetriggeredorder) | **Get** /futures/{settle}/price_orders/{order_id} | Get a single order
*FuturesApi* | [**CancelPriceTriggeredOrder**](docs/FuturesApi.md#cancelpricetriggeredorder) | **Delete** /futures/{settle}/price_orders/{order_id} | Cancel a single order
*MarginApi* | [**ListMarginCurrencyPairs**](docs/MarginApi.md#listmargincurrencypairs) | **Get** /margin/currency_pairs | List all supported currency pairs supported in margin trading
*MarginApi* | [**GetMarginCurrencyPair**](docs/MarginApi.md#getmargincurrencypair) | **Get** /margin/currency_pairs/{currency_pair} | Query one single margin currency pair
*MarginApi* | [**ListFundingBook**](docs/MarginApi.md#listfundingbook) | **Get** /margin/funding_book | Order book of lending loans
*MarginApi* | [**ListMarginAccounts**](docs/MarginApi.md#listmarginaccounts) | **Get** /margin/accounts | Margin account list
*MarginApi* | [**ListMarginAccountBook**](docs/MarginApi.md#listmarginaccountbook) | **Get** /margin/account_book | List margin account balance change history
*MarginApi* | [**ListFundingAccounts**](docs/MarginApi.md#listfundingaccounts) | **Get** /margin/funding_accounts | Funding account list
*MarginApi* | [**ListLoans**](docs/MarginApi.md#listloans) | **Get** /margin/loans | List all loans
*MarginApi* | [**CreateLoan**](docs/MarginApi.md#createloan) | **Post** /margin/loans | Lend or borrow
*MarginApi* | [**MergeLoans**](docs/MarginApi.md#mergeloans) | **Post** /margin/merged_loans | Merge multiple lending loans
*MarginApi* | [**GetLoan**](docs/MarginApi.md#getloan) | **Get** /margin/loans/{loan_id} | Retrieve one single loan detail
*MarginApi* | [**CancelLoan**](docs/MarginApi.md#cancelloan) | **Delete** /margin/loans/{loan_id} | Cancel lending loan
*MarginApi* | [**UpdateLoan**](docs/MarginApi.md#updateloan) | **Patch** /margin/loans/{loan_id} | Modify a loan
*MarginApi* | [**ListLoanRepayments**](docs/MarginApi.md#listloanrepayments) | **Get** /margin/loans/{loan_id}/repayment | List loan repayment records
*MarginApi* | [**RepayLoan**](docs/MarginApi.md#repayloan) | **Post** /margin/loans/{loan_id}/repayment | Repay a loan
*MarginApi* | [**ListLoanRecords**](docs/MarginApi.md#listloanrecords) | **Get** /margin/loan_records | List repayment records of specified loan
*MarginApi* | [**GetLoanRecord**](docs/MarginApi.md#getloanrecord) | **Get** /margin/loan_records/{loan_record_id} | Get one single loan record
*MarginApi* | [**UpdateLoanRecord**](docs/MarginApi.md#updateloanrecord) | **Patch** /margin/loan_records/{loan_record_id} | Modify a loan record
*MarginApi* | [**GetAutoRepayStatus**](docs/MarginApi.md#getautorepaystatus) | **Get** /margin/auto_repay | Retrieve user auto repayment setting
*MarginApi* | [**SetAutoRepay**](docs/MarginApi.md#setautorepay) | **Post** /margin/auto_repay | Update user&#39;s auto repayment setting
*SpotApi* | [**ListCurrencies**](docs/SpotApi.md#listcurrencies) | **Get** /spot/currencies | List all currencies&#39; detail
*SpotApi* | [**GetCurrency**](docs/SpotApi.md#getcurrency) | **Get** /spot/currencies/{currency} | Get detail of one particular currency
*SpotApi* | [**ListCurrencyPairs**](docs/SpotApi.md#listcurrencypairs) | **Get** /spot/currency_pairs | List all currency pairs supported
*SpotApi* | [**GetCurrencyPair**](docs/SpotApi.md#getcurrencypair) | **Get** /spot/currency_pairs/{currency_pair} | Get detail of one single order
*SpotApi* | [**ListTickers**](docs/SpotApi.md#listtickers) | **Get** /spot/tickers | Retrieve ticker information
*SpotApi* | [**ListOrderBook**](docs/SpotApi.md#listorderbook) | **Get** /spot/order_book | Retrieve order book
*SpotApi* | [**ListTrades**](docs/SpotApi.md#listtrades) | **Get** /spot/trades | Retrieve market trades
*SpotApi* | [**ListCandlesticks**](docs/SpotApi.md#listcandlesticks) | **Get** /spot/candlesticks | Market candlesticks
*SpotApi* | [**GetFee**](docs/SpotApi.md#getfee) | **Get** /spot/fee | Query user trading fee rates
*SpotApi* | [**ListSpotAccounts**](docs/SpotApi.md#listspotaccounts) | **Get** /spot/accounts | List spot accounts
*SpotApi* | [**CreateBatchOrders**](docs/SpotApi.md#createbatchorders) | **Post** /spot/batch_orders | Create a batch of orders
*SpotApi* | [**ListAllOpenOrders**](docs/SpotApi.md#listallopenorders) | **Get** /spot/open_orders | List all open orders
*SpotApi* | [**ListOrders**](docs/SpotApi.md#listorders) | **Get** /spot/orders | List orders
*SpotApi* | [**CreateOrder**](docs/SpotApi.md#createorder) | **Post** /spot/orders | Create an order
*SpotApi* | [**CancelOrders**](docs/SpotApi.md#cancelorders) | **Delete** /spot/orders | Cancel all &#x60;open&#x60; orders in specified currency pair
*SpotApi* | [**CancelBatchOrders**](docs/SpotApi.md#cancelbatchorders) | **Post** /spot/cancel_batch_orders | Cancel a batch of orders with an ID list
*SpotApi* | [**GetOrder**](docs/SpotApi.md#getorder) | **Get** /spot/orders/{order_id} | Get a single order
*SpotApi* | [**CancelOrder**](docs/SpotApi.md#cancelorder) | **Delete** /spot/orders/{order_id} | Cancel a single order
*SpotApi* | [**ListMyTrades**](docs/SpotApi.md#listmytrades) | **Get** /spot/my_trades | List personal trading history
*SpotApi* | [**ListSpotPriceTriggeredOrders**](docs/SpotApi.md#listspotpricetriggeredorders) | **Get** /spot/price_orders | Retrieve running auto order list
*SpotApi* | [**CreateSpotPriceTriggeredOrder**](docs/SpotApi.md#createspotpricetriggeredorder) | **Post** /spot/price_orders | Create a price-triggered order
*SpotApi* | [**CancelSpotPriceTriggeredOrderList**](docs/SpotApi.md#cancelspotpricetriggeredorderlist) | **Delete** /spot/price_orders | Cancel all open orders
*SpotApi* | [**GetSpotPriceTriggeredOrder**](docs/SpotApi.md#getspotpricetriggeredorder) | **Get** /spot/price_orders/{order_id} | Get a single order
*SpotApi* | [**CancelSpotPriceTriggeredOrder**](docs/SpotApi.md#cancelspotpricetriggeredorder) | **Delete** /spot/price_orders/{order_id} | Cancel a single order
*WalletApi* | [**GetDepositAddress**](docs/WalletApi.md#getdepositaddress) | **Get** /wallet/deposit_address | Generate currency deposit address
*WalletApi* | [**ListWithdrawals**](docs/WalletApi.md#listwithdrawals) | **Get** /wallet/withdrawals | Retrieve withdrawal records
*WalletApi* | [**ListDeposits**](docs/WalletApi.md#listdeposits) | **Get** /wallet/deposits | Retrieve deposit records
*WalletApi* | [**Transfer**](docs/WalletApi.md#transfer) | **Post** /wallet/transfers | Transfer between trading accounts
*WalletApi* | [**ListSubAccountTransfers**](docs/WalletApi.md#listsubaccounttransfers) | **Get** /wallet/sub_account_transfers | Transfer records between main and sub accounts
*WalletApi* | [**TransferWithSubAccount**](docs/WalletApi.md#transferwithsubaccount) | **Post** /wallet/sub_account_transfers | Transfer between main and sub accounts
*WalletApi* | [**ListWithdrawStatus**](docs/WalletApi.md#listwithdrawstatus) | **Get** /wallet/withdraw_status | Retrieve withdrawal status
*WalletApi* | [**ListSubAccountBalances**](docs/WalletApi.md#listsubaccountbalances) | **Get** /wallet/sub_account_balances | Retrieve sub account balances
*WalletApi* | [**GetTradeFee**](docs/WalletApi.md#gettradefee) | **Get** /wallet/fee | Retrieve personal trading fee
*WithdrawalApi* | [**Withdraw**](docs/WithdrawalApi.md#withdraw) | **Post** /withdrawals | Withdraw


## Documentation For Models

 - [AutoRepaySetting](docs/AutoRepaySetting.md)
 - [BatchOrder](docs/BatchOrder.md)
 - [CancelOrder](docs/CancelOrder.md)
 - [CancelOrderResult](docs/CancelOrderResult.md)
 - [Contract](docs/Contract.md)
 - [ContractStat](docs/ContractStat.md)
 - [Currency](docs/Currency.md)
 - [CurrencyPair](docs/CurrencyPair.md)
 - [DeliveryContract](docs/DeliveryContract.md)
 - [DeliverySettlement](docs/DeliverySettlement.md)
 - [DepositAddress](docs/DepositAddress.md)
 - [FundingAccount](docs/FundingAccount.md)
 - [FundingBookItem](docs/FundingBookItem.md)
 - [FundingRateRecord](docs/FundingRateRecord.md)
 - [FuturesAccount](docs/FuturesAccount.md)
 - [FuturesAccountBook](docs/FuturesAccountBook.md)
 - [FuturesCandlestick](docs/FuturesCandlestick.md)
 - [FuturesInitialOrder](docs/FuturesInitialOrder.md)
 - [FuturesLiquidate](docs/FuturesLiquidate.md)
 - [FuturesOrder](docs/FuturesOrder.md)
 - [FuturesOrderBook](docs/FuturesOrderBook.md)
 - [FuturesOrderBookItem](docs/FuturesOrderBookItem.md)
 - [FuturesPriceTrigger](docs/FuturesPriceTrigger.md)
 - [FuturesPriceTriggeredOrder](docs/FuturesPriceTriggeredOrder.md)
 - [FuturesTicker](docs/FuturesTicker.md)
 - [FuturesTrade](docs/FuturesTrade.md)
 - [InsuranceRecord](docs/InsuranceRecord.md)
 - [LedgerRecord](docs/LedgerRecord.md)
 - [Loan](docs/Loan.md)
 - [LoanPatch](docs/LoanPatch.md)
 - [LoanRecord](docs/LoanRecord.md)
 - [MarginAccount](docs/MarginAccount.md)
 - [MarginAccountBook](docs/MarginAccountBook.md)
 - [MarginAccountCurrency](docs/MarginAccountCurrency.md)
 - [MarginCurrencyPair](docs/MarginCurrencyPair.md)
 - [MultiChainAddressItem](docs/MultiChainAddressItem.md)
 - [MyFuturesTrade](docs/MyFuturesTrade.md)
 - [OpenOrders](docs/OpenOrders.md)
 - [Order](docs/Order.md)
 - [OrderBook](docs/OrderBook.md)
 - [Position](docs/Position.md)
 - [PositionClose](docs/PositionClose.md)
 - [PositionCloseOrder](docs/PositionCloseOrder.md)
 - [RepayRequest](docs/RepayRequest.md)
 - [Repayment](docs/Repayment.md)
 - [SpotAccount](docs/SpotAccount.md)
 - [SpotPricePutOrder](docs/SpotPricePutOrder.md)
 - [SpotPriceTrigger](docs/SpotPriceTrigger.md)
 - [SpotPriceTriggeredOrder](docs/SpotPriceTriggeredOrder.md)
 - [SubAccountBalance](docs/SubAccountBalance.md)
 - [SubAccountTransfer](docs/SubAccountTransfer.md)
 - [Ticker](docs/Ticker.md)
 - [Trade](docs/Trade.md)
 - [TradeFee](docs/TradeFee.md)
 - [Transfer](docs/Transfer.md)
 - [TriggerOrderResponse](docs/TriggerOrderResponse.md)
 - [WithdrawStatus](docs/WithdrawStatus.md)


## Documentation For Authorization



## apiv4

- **Type**: Gate APIv4

Example

```golang
ctx := context.WithValue(context.Background(), gateapi.ContextAPIKey, gateapi.GateAPIV4{
    Key:    "YOUR_API_KEY",
    Secret: "YOUR_API_SECRET",
})
r, err := client.Service.Operation(ctx, args)
```


## Author

support@mail.gate.io

