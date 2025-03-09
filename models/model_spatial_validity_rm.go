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



// This data type is defined in the same way as the SpatialValidity data type, but with the OpenAPI nullable property set to true. 
type SpatialValidityRm struct {
	// Defines the presence information provisioned by the AF. The praId attribute within the  PresenceInfo data type is the key of the map. 
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList,omitempty" yaml:"presenceInfoList" bson:"presenceInfoList,omitempty"`
}

