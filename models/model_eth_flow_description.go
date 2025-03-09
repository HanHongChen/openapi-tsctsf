/*
 * Ntsctsf_QoSandTSCAssistance Service API
 *
 * TSCTSF QoS and TSC Assistance Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.565 V18.3.0; 5G System; Time Sensitive Communication and Time Synchronization function  Services; Stage 3. 
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.565/
 *
 * API version: 1.1.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	
)



// Identifies an Ethernet flow.
type EthFlowDescription struct {
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddr string `json:"destMacAddr,omitempty" yaml:"destMacAddr" bson:"destMacAddr,omitempty"`
	EthType string `json:"ethType,omitempty" yaml:"ethType" bson:"ethType,omitempty"`
	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty" yaml:"fDesc" bson:"fDesc,omitempty"`
	FDir FlowDirection `json:"fDir,omitempty" yaml:"fDir" bson:"fDir,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SourceMacAddr string `json:"sourceMacAddr,omitempty" yaml:"sourceMacAddr" bson:"sourceMacAddr,omitempty"`
	VlanTags []string `json:"vlanTags,omitempty" yaml:"vlanTags" bson:"vlanTags,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SrcMacAddrEnd string `json:"srcMacAddrEnd,omitempty" yaml:"srcMacAddrEnd" bson:"srcMacAddrEnd,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddrEnd string `json:"destMacAddrEnd,omitempty" yaml:"destMacAddrEnd" bson:"destMacAddrEnd,omitempty"`
}

