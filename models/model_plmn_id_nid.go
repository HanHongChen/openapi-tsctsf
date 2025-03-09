/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.565 V18.5.0; 5G System; Time Sensitive Communication and Time Synchronization Function  Services; Stage 3. 
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.565/
 *
 * API version: 1.1.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	
)



// Contains the serving core network operator PLMN ID and, for an SNPN, the NID that together with the PLMN ID identifies the SNPN. 
type PlmnIdNid struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.  
	Mcc string `json:"mcc,omitempty" yaml:"mcc" bson:"mcc,omitempty"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in  clause 9.3.3.5 of 3GPP TS 38.413.  
	Mnc string `json:"mnc,omitempty" yaml:"mnc" bson:"mnc,omitempty"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty" yaml:"nid" bson:"nid,omitempty"`
}

