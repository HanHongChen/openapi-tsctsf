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



// Relative Cartesian Location
type RelativeCartesianLocation struct {
	// string with format 'float' as defined in OpenAPI.
	X float32 `json:"x,omitempty" yaml:"x" bson:"x,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	Y float32 `json:"y,omitempty" yaml:"y" bson:"y,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	Z float32 `json:"z,omitempty" yaml:"z" bson:"z,omitempty"`
}

