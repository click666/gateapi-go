/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Input for the portfolio margin calculator.
type UnifiedPortfolioInput struct {
	// Spot
	SpotBalances []MockSpotBalance `json:"spot_balances,omitempty"`
	// Spot orders
	SpotOrders []MockSpotOrder `json:"spot_orders,omitempty"`
	// Futures positions
	FuturesPositions []MockFuturesPosition `json:"futures_positions,omitempty"`
	// Futures order
	FuturesOrders []MockFuturesOrder `json:"futures_orders,omitempty"`
	// Options positions
	OptionsPositions []MockOptionsPosition `json:"options_positions,omitempty"`
	// Option orders
	OptionsOrders []MockOptionsOrder `json:"options_orders,omitempty"`
	// Whether to enable spot hedging.
	SpotHedge bool `json:"spot_hedge,omitempty"`
}