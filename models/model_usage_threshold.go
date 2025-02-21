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



// Represents a usage threshold.
type UsageThreshold struct {
	// Unsigned integer identifying a period of time in units of seconds.
	Duration int32 `json:"duration,omitempty" yaml:"duration" bson:"duration,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	TotalVolume int64 `json:"totalVolume,omitempty" yaml:"totalVolume" bson:"totalVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume int64 `json:"downlinkVolume,omitempty" yaml:"downlinkVolume" bson:"downlinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume int64 `json:"uplinkVolume,omitempty" yaml:"uplinkVolume" bson:"uplinkVolume,omitempty"`
}


