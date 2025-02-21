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



// Contains the filter conditions to match for notifying the event(s) of time synchronization capabilities. 
type EventFilter struct {
	// Indicates the PTP instance type(s). 
	InstanceTypes []InstanceType `json:"instanceTypes,omitempty" yaml:"instanceTypes" bson:"instanceTypes,omitempty"`
	// Indicates the transport protocol type(s). 
	TransProtocols []Protocol `json:"transProtocols,omitempty" yaml:"transProtocols" bson:"transProtocols,omitempty"`
	// Identifies the supported PTP profiles. 
	PtpProfiles []string `json:"ptpProfiles,omitempty" yaml:"ptpProfiles" bson:"ptpProfiles,omitempty"`
}


