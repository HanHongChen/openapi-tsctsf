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



// Indicates a Civic address.
type CivicAddress struct {
	Country string `json:"country,omitempty" yaml:"country" bson:"country,omitempty"`
	A1 string `json:"A1,omitempty" yaml:"A1" bson:"A1,omitempty"`
	A2 string `json:"A2,omitempty" yaml:"A2" bson:"A2,omitempty"`
	A3 string `json:"A3,omitempty" yaml:"A3" bson:"A3,omitempty"`
	A4 string `json:"A4,omitempty" yaml:"A4" bson:"A4,omitempty"`
	A5 string `json:"A5,omitempty" yaml:"A5" bson:"A5,omitempty"`
	A6 string `json:"A6,omitempty" yaml:"A6" bson:"A6,omitempty"`
	PRD string `json:"PRD,omitempty" yaml:"PRD" bson:"PRD,omitempty"`
	POD string `json:"POD,omitempty" yaml:"POD" bson:"POD,omitempty"`
	STS string `json:"STS,omitempty" yaml:"STS" bson:"STS,omitempty"`
	HNO string `json:"HNO,omitempty" yaml:"HNO" bson:"HNO,omitempty"`
	HNS string `json:"HNS,omitempty" yaml:"HNS" bson:"HNS,omitempty"`
	LMK string `json:"LMK,omitempty" yaml:"LMK" bson:"LMK,omitempty"`
	LOC string `json:"LOC,omitempty" yaml:"LOC" bson:"LOC,omitempty"`
	NAM string `json:"NAM,omitempty" yaml:"NAM" bson:"NAM,omitempty"`
	PC string `json:"PC,omitempty" yaml:"PC" bson:"PC,omitempty"`
	BLD string `json:"BLD,omitempty" yaml:"BLD" bson:"BLD,omitempty"`
	UNIT string `json:"UNIT,omitempty" yaml:"UNIT" bson:"UNIT,omitempty"`
	FLR string `json:"FLR,omitempty" yaml:"FLR" bson:"FLR,omitempty"`
	ROOM string `json:"ROOM,omitempty" yaml:"ROOM" bson:"ROOM,omitempty"`
	PLC string `json:"PLC,omitempty" yaml:"PLC" bson:"PLC,omitempty"`
	PCN string `json:"PCN,omitempty" yaml:"PCN" bson:"PCN,omitempty"`
	POBOX string `json:"POBOX,omitempty" yaml:"POBOX" bson:"POBOX,omitempty"`
	ADDCODE string `json:"ADDCODE,omitempty" yaml:"ADDCODE" bson:"ADDCODE,omitempty"`
	SEAT string `json:"SEAT,omitempty" yaml:"SEAT" bson:"SEAT,omitempty"`
	RD string `json:"RD,omitempty" yaml:"RD" bson:"RD,omitempty"`
	RDSEC string `json:"RDSEC,omitempty" yaml:"RDSEC" bson:"RDSEC,omitempty"`
	RDBR string `json:"RDBR,omitempty" yaml:"RDBR" bson:"RDBR,omitempty"`
	RDSUBBR string `json:"RDSUBBR,omitempty" yaml:"RDSUBBR" bson:"RDSUBBR,omitempty"`
	PRM string `json:"PRM,omitempty" yaml:"PRM" bson:"PRM,omitempty"`
	POM string `json:"POM,omitempty" yaml:"POM" bson:"POM,omitempty"`
	UsageRules string `json:"usageRules,omitempty" yaml:"usageRules" bson:"usageRules,omitempty"`
	Method string `json:"method,omitempty" yaml:"method" bson:"method,omitempty"`
	ProvidedBy string `json:"providedBy,omitempty" yaml:"providedBy" bson:"providedBy,omitempty"`
}


