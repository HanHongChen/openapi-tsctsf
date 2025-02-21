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

package TimeSynchronization

// APIClient manages communication with the Ntsctsf_TimeSynchronization Service API API v1.1.0-alpha.6
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg    *Configuration
    common service // Reuse a single struct instead of allocating one for each service on the heap.

    // API Services
    	IndividualTimeSynchronizationExposureConfigurationDocumentApi *IndividualTimeSynchronizationExposureConfigurationDocumentApiService
    	IndividualTimeSynchronizationExposureSubscriptionDocumentApi *IndividualTimeSynchronizationExposureSubscriptionDocumentApiService
    	TimeSynchronizationExposureSubscriptionsCollectionApi *TimeSynchronizationExposureSubscriptionsCollectionApiService
}

type service struct {
    client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
    c := &APIClient{}
    c.cfg = cfg
    c.common.client = c

    // API Services
    c.IndividualTimeSynchronizationExposureConfigurationDocumentApi = (*IndividualTimeSynchronizationExposureConfigurationDocumentApiService)(&c.common)
    c.IndividualTimeSynchronizationExposureSubscriptionDocumentApi = (*IndividualTimeSynchronizationExposureSubscriptionDocumentApiService)(&c.common)
    c.TimeSynchronizationExposureSubscriptionsCollectionApi = (*TimeSynchronizationExposureSubscriptionsCollectionApiService)(&c.common)

    return c
}
