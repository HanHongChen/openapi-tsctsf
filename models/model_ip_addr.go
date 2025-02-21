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



// Contains an IP adresse.
type IpAddr struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr,omitempty"`
	Ipv6Addr string `json:"ipv6Addr,omitempty" yaml:"ipv6Addr" bson:"ipv6Addr,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty" yaml:"ipv6Prefix" bson:"ipv6Prefix,omitempty"`
}


