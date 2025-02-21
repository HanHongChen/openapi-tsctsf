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


type MatchingOperator string

// List of MatchingOperator
const (
        MatchingOperator_FULL_MATCH MatchingOperator = "FULL_MATCH"
        MatchingOperator_MATCH_ALL MatchingOperator = "MATCH_ALL"
        MatchingOperator_STARTS_WITH MatchingOperator = "STARTS_WITH"
        MatchingOperator_NOT_START_WITH MatchingOperator = "NOT_START_WITH"
        MatchingOperator_ENDS_WITH MatchingOperator = "ENDS_WITH"
        MatchingOperator_NOT_END_WITH MatchingOperator = "NOT_END_WITH"
        MatchingOperator_CONTAINS MatchingOperator = "CONTAINS"
        MatchingOperator_NOT_CONTAIN MatchingOperator = "NOT_CONTAIN"
)




