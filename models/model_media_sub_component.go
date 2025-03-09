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



// Identifies a media subcomponent.
type MediaSubComponent struct {
	AfSigProtocol AfSigProtocol `json:"afSigProtocol,omitempty" yaml:"afSigProtocol" bson:"afSigProtocol,omitempty"`
	EthfDescs []EthFlowDescription `json:"ethfDescs,omitempty" yaml:"ethfDescs" bson:"ethfDescs,omitempty"`
	FNum int32 `json:"fNum,omitempty" yaml:"fNum" bson:"fNum,omitempty"`
	FDescs []string `json:"fDescs,omitempty" yaml:"fDescs" bson:"fDescs,omitempty"`
	// Represents additional flow description information (flow label and IPsec SPI) per Uplink and/or Downlink IP flows. 
	AddInfoFlowDescs []AddFlowDescriptionInfo `json:"addInfoFlowDescs,omitempty" yaml:"addInfoFlowDescs" bson:"addInfoFlowDescs,omitempty"`
	FStatus FlowStatus `json:"fStatus,omitempty" yaml:"fStatus" bson:"fStatus,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwDl string `json:"marBwDl,omitempty" yaml:"marBwDl" bson:"marBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwUl string `json:"marBwUl,omitempty" yaml:"marBwUl" bson:"marBwUl,omitempty"`
	// 2-octet string, where each octet is encoded in hexadecimal representation. The first octet contains the IPv4 Type-of-Service or the IPv6 Traffic-Class field and the second octet contains the ToS/Traffic Class mask field. 
	TosTrCl string `json:"tosTrCl,omitempty" yaml:"tosTrCl" bson:"tosTrCl,omitempty"`
	FlowUsage FlowUsage `json:"flowUsage,omitempty" yaml:"flowUsage" bson:"flowUsage,omitempty"`
	EvSubsc *EventsSubscReqData `json:"evSubsc,omitempty" yaml:"evSubsc" bson:"evSubsc,omitempty"`
}

