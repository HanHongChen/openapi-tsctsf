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



// Contains the PTP port state of a DS-TT.
type StateOfDstt struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// When the PTP port state is Leader, Follower or Passive, it is included and set to true to indicate the state of configuration for DS-TT port is active; when PTP port state is in any other case, it is included and set to false to indicate the state of configuration for DS-TT port is inactive. Default value is false. 
	State bool `json:"state,omitempty" yaml:"state" bson:"state,omitempty"`
	ClkQltIndOfDstt AcceptanceCriteriaResultIndication `json:"clkQltIndOfDstt,omitempty" yaml:"clkQltIndOfDstt" bson:"clkQltIndOfDstt,omitempty"`
}


