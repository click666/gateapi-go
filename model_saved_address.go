/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SavedAddress struct {
	// Currency
	Currency string `json:"currency,omitempty"`
	// Chain name
	Chain string `json:"chain,omitempty"`
	// Address
	Address string `json:"address,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// Tag
	Tag string `json:"tag,omitempty"`
	// Whether to pass the verification 0-unverified, 1-verified
	Verified string `json:"verified,omitempty"`
}
