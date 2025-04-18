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



// Represents QoS monitoring information.
type QosMonitoringInformation struct {
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty" yaml:"reqQosMonParams" bson:"reqQosMonParams,omitempty"`
	RepFreqs []ReportingFrequency `json:"repFreqs,omitempty" yaml:"repFreqs" bson:"repFreqs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RepThreshDl int32 `json:"repThreshDl,omitempty" yaml:"repThreshDl" bson:"repThreshDl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RepThreshUl int32 `json:"repThreshUl,omitempty" yaml:"repThreshUl" bson:"repThreshUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RepThreshRp int32 `json:"repThreshRp,omitempty" yaml:"repThreshRp" bson:"repThreshRp,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ConThreshDl int32 `json:"conThreshDl,omitempty" yaml:"conThreshDl" bson:"conThreshDl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ConThreshUl int32 `json:"conThreshUl,omitempty" yaml:"conThreshUl" bson:"conThreshUl,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	WaitTime int32 `json:"waitTime,omitempty" yaml:"waitTime" bson:"waitTime,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	RepPeriod int32 `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	RepThreshDatRateDl string `json:"repThreshDatRateDl,omitempty" yaml:"repThreshDatRateDl" bson:"repThreshDatRateDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	RepThreshDatRateUl string `json:"repThreshDatRateUl,omitempty" yaml:"repThreshDatRateUl" bson:"repThreshDatRateUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ConsDataRateThrDl string `json:"consDataRateThrDl,omitempty" yaml:"consDataRateThrDl" bson:"consDataRateThrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ConsDataRateThrUl string `json:"consDataRateThrUl,omitempty" yaml:"consDataRateThrUl" bson:"consDataRateThrUl,omitempty"`
}

