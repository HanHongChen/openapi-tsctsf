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


type NotificationMethod string

// List of NotificationMethod
const (
        NotificationMethod_PERIODIC NotificationMethod = "PERIODIC"
        NotificationMethod_ONE_TIME NotificationMethod = "ONE_TIME"
        NotificationMethod_ON_EVENT_DETECTION NotificationMethod = "ON_EVENT_DETECTION"
)




