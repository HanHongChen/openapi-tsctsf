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



// Contains the Spatial Validity Condition.
type SpatialValidityCond struct {
	TrackingAreaList []Tai `json:"trackingAreaList,omitempty" yaml:"trackingAreaList" bson:"trackingAreaList,omitempty"`
	Countries []string `json:"countries,omitempty" yaml:"countries" bson:"countries,omitempty"`
	GeographicalServiceArea *GeoServiceArea `json:"geographicalServiceArea,omitempty" yaml:"geographicalServiceArea" bson:"geographicalServiceArea,omitempty"`
}

