/*
 * Gate API v4
 *
 * APIv4 futures provides all sorts of futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 1.2.0
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesTicker struct {
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Last trading price
	Last string `json:"last,omitempty"`
	// Change percentage.
	ChangePercentage string `json:"change_percentage,omitempty"`
	// Contract total size
	TotalSize string `json:"total_size,omitempty"`
	// Trade size in recent 24h
	Volume24h string `json:"volume_24h,omitempty"`
	// Recent mark price
	MarkPrice string `json:"mark_price,omitempty"`
	// Funding rate
	FundingRate string `json:"funding_rate,omitempty"`
	// Index price
	IndexPrice string `json:"index_price,omitempty"`
	// Exchange rate of base currency and settlement currency in Quanto contract. Not existed in contract of other types
	QuantoBaseRate string `json:"quanto_base_rate,omitempty"`
}
