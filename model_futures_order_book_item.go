/*
 * Gate API v4
 *
 * APIv4 futures provides all sorts of futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 1.2.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesOrderBookItem struct {
	// Price
	P string `json:"p,omitempty"`
	// Size
	S int64 `json:"s,omitempty"`
}