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


type NoProfileMatchReason string

// List of NoProfileMatchReason
const (
        NoProfileMatchReason_REQUESTER_PLMN_NOT_ALLOWED NoProfileMatchReason = "REQUESTER_PLMN_NOT_ALLOWED"
        NoProfileMatchReason_TARGET_NF_SUSPENDED NoProfileMatchReason = "TARGET_NF_SUSPENDED"
        NoProfileMatchReason_TARGET_NF_UNDISCOVERABLE NoProfileMatchReason = "TARGET_NF_UNDISCOVERABLE"
        NoProfileMatchReason_QUERY_PARAMS_COMBINATION_NO_MATCH NoProfileMatchReason = "QUERY_PARAMS_COMBINATION_NO_MATCH"
        NoProfileMatchReason_TARGET_NF_TYPE_NOT_SUPPORTED NoProfileMatchReason = "TARGET_NF_TYPE_NOT_SUPPORTED"
        NoProfileMatchReason_UNSPECIFIED NoProfileMatchReason = "UNSPECIFIED"
)



