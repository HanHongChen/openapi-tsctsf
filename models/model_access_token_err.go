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



// Error returned in the access token response message
type AccessTokenErr struct {
	Error string `json:"error,omitempty" yaml:"error" bson:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty" yaml:"error_description" bson:"error_description,omitempty"`
	ErrorUri string `json:"error_uri,omitempty" yaml:"error_uri" bson:"error_uri,omitempty"`
}

