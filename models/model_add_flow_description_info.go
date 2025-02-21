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



// Contains additional flow description information.
type AddFlowDescriptionInfo struct {
	// 4-octet string representing the security parameter index of the IPSec packet in hexadecimal representation. 
	Spi string `json:"spi,omitempty" yaml:"spi" bson:"spi,omitempty"`
	// 3-octet string representing the IPv6 flow label header field in hexadecimal representation. 
	FlowLabel string `json:"flowLabel,omitempty" yaml:"flowLabel" bson:"flowLabel,omitempty"`
	FlowDir FlowDirection `json:"flowDir,omitempty" yaml:"flowDir" bson:"flowDir,omitempty"`
}


