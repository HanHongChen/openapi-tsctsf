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


type RequestedQosMonitoringParameter string

// List of RequestedQosMonitoringParameter
const (
        RequestedQosMonitoringParameter_DOWNLINK RequestedQosMonitoringParameter = "DOWNLINK"
        RequestedQosMonitoringParameter_UPLINK RequestedQosMonitoringParameter = "UPLINK"
        RequestedQosMonitoringParameter_ROUND_TRIP RequestedQosMonitoringParameter = "ROUND_TRIP"
        RequestedQosMonitoringParameter_DOWNLINK_DATA_RATE RequestedQosMonitoringParameter = "DOWNLINK_DATA_RATE"
        RequestedQosMonitoringParameter_UPLINK_DATA_RATE RequestedQosMonitoringParameter = "UPLINK_DATA_RATE"
        RequestedQosMonitoringParameter_DOWNLINK_CONGESTION RequestedQosMonitoringParameter = "DOWNLINK_CONGESTION"
        RequestedQosMonitoringParameter_UPLINK_CONGESTION RequestedQosMonitoringParameter = "UPLINK_CONGESTION"
)



