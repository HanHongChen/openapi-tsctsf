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



// Geographic area specified by different shape.
type GeographicArea struct {
	Shape SupportedGadShapes `json:"shape,omitempty" yaml:"shape" bson:"shape,omitempty"`
	Point *GeographicalCoordinates `json:"point,omitempty" yaml:"point" bson:"point,omitempty"`
	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty,omitempty" yaml:"uncertainty" bson:"uncertainty,omitempty"`
	UncertaintyEllipse *UncertaintyEllipse `json:"uncertaintyEllipse,omitempty" yaml:"uncertaintyEllipse" bson:"uncertaintyEllipse,omitempty"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
	// List of points.
	PointList []GeographicalCoordinates `json:"pointList,omitempty" yaml:"pointList" bson:"pointList,omitempty"`
	// Indicates value of altitude.
	Altitude float64 `json:"altitude,omitempty" yaml:"altitude" bson:"altitude,omitempty"`
	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude,omitempty" yaml:"uncertaintyAltitude" bson:"uncertaintyAltitude,omitempty"`
	// Indicates value of confidence.
	VConfidence int32 `json:"vConfidence,omitempty" yaml:"vConfidence" bson:"vConfidence,omitempty"`
	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius,omitempty" yaml:"innerRadius" bson:"innerRadius,omitempty"`
	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius,omitempty" yaml:"uncertaintyRadius" bson:"uncertaintyRadius,omitempty"`
	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle,omitempty" yaml:"offsetAngle" bson:"offsetAngle,omitempty"`
	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle,omitempty" yaml:"includedAngle" bson:"includedAngle,omitempty"`
}

