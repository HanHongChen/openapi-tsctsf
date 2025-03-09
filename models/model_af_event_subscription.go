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



// Describes the event information delivered in the subscription.
type AfEventSubscription struct {
	Event AfEvent `json:"event,omitempty" yaml:"event" bson:"event,omitempty"`
	NotifMethod AfNotifMethod `json:"notifMethod,omitempty" yaml:"notifMethod" bson:"notifMethod,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	RepPeriod int32 `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	WaitTime int32 `json:"waitTime,omitempty" yaml:"waitTime" bson:"waitTime,omitempty"`
	QosMonParamType QosMonitoringParamType `json:"qosMonParamType,omitempty" yaml:"qosMonParamType" bson:"qosMonParamType,omitempty"`
}

