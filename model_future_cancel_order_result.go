/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Order cancellation result
type FutureCancelOrderResult struct {
	// Order ID
	Id string `json:"id,omitempty"`
	// User ID
	UserId int64 `json:"user_id,omitempty"`
	// Whether cancellation succeeded
	Succeeded bool `json:"succeeded,omitempty"`
	// Error message when failed to cancel the order; empty if succeeded
	Message string `json:"message,omitempty"`
}