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



// Ellipse with uncertainty.
type UncertaintyEllipse struct {
	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor,omitempty" yaml:"semiMajor" bson:"semiMajor,omitempty"`
	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor,omitempty" yaml:"semiMinor" bson:"semiMinor,omitempty"`
	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor,omitempty" yaml:"orientationMajor" bson:"orientationMajor,omitempty"`
}

