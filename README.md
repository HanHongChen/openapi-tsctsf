# Go API client for TimeSynchronization

TSCTSF Time Synchronization Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0-alpha.6
- Package version: 1.0.0
- Build package: org.free5gc.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./TimeSynchronization"
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/ntsctsf-time-sync/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualTimeSynchronizationExposureConfigurationDocumentApi* | [**CreateIndividualTimeSynchronizationExposureConfiguration**](docs/IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#createindividualtimesynchronizationexposureconfiguration) | **Post** /subscriptions/{subscriptionId}/configurations | Craete a new Individual Time Synchronization Exposure Configuration
*IndividualTimeSynchronizationExposureConfigurationDocumentApi* | [**DeleteIndividualTimeSynchronizationExposureConfiguration**](docs/IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#deleteindividualtimesynchronizationexposureconfiguration) | **Delete** /subscriptions/{subscriptionId}/configurations/{configurationId} | Delete an Individual TimeSynchronization Exposure Configuration
*IndividualTimeSynchronizationExposureConfigurationDocumentApi* | [**GetIndividualTimeSynchronizationExposureConfiguration**](docs/IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#getindividualtimesynchronizationexposureconfiguration) | **Get** /subscriptions/{subscriptionId}/configurations/{configurationId} | Reads an existing Individual Time Synchronization Exposure Configuration
*IndividualTimeSynchronizationExposureConfigurationDocumentApi* | [**ReplaceIndividualTimeSynchronizationExposureConfiguration**](docs/IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#replaceindividualtimesynchronizationexposureconfiguration) | **Put** /subscriptions/{subscriptionId}/configurations/{configurationId} | Replace an individual Time Synchronization Exposure Configuration
*IndividualTimeSynchronizationExposureSubscriptionDocumentApi* | [**DeleteIndividualTimeSynchronizationExposureSubscription**](docs/IndividualTimeSynchronizationExposureSubscriptionDocumentApi.md#deleteindividualtimesynchronizationexposuresubscription) | **Delete** /subscriptions/{subscriptionId} | Delete an Individual TimeSynchronization Exposure Subscription
*IndividualTimeSynchronizationExposureSubscriptionDocumentApi* | [**GetIndividualTimeSynchronizationExposureSubscription**](docs/IndividualTimeSynchronizationExposureSubscriptionDocumentApi.md#getindividualtimesynchronizationexposuresubscription) | **Get** /subscriptions/{subscriptionId} | Reads an existing Individual Time Synchronization Exposure Subscription
*IndividualTimeSynchronizationExposureSubscriptionDocumentApi* | [**ReplaceIndividualTimeSynchronizationExposureSubscription**](docs/IndividualTimeSynchronizationExposureSubscriptionDocumentApi.md#replaceindividualtimesynchronizationexposuresubscription) | **Put** /subscriptions/{subscriptionId} | Replace an individual Time Synchronization Exposure Subscription
*TimeSynchronizationExposureSubscriptionsCollectionApi* | [**TimeSynchronizationExposureSubscriptions**](docs/TimeSynchronizationExposureSubscriptionsCollectionApi.md#timesynchronizationexposuresubscriptions) | **Post** /subscriptions | Creates a new subscription to notification of capability of time synchronization service resource


## Documentation For Models

 - [AcceptanceCriteriaResultIndication](docs/AcceptanceCriteriaResultIndication.md)
 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AsTimeResource](docs/AsTimeResource.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [ClockQuality](docs/ClockQuality.md)
 - [ClockQualityAcceptanceCriterion](docs/ClockQualityAcceptanceCriterion.md)
 - [ClockQualityDetailLevel](docs/ClockQualityDetailLevel.md)
 - [ConfigForPort](docs/ConfigForPort.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EventFilter](docs/EventFilter.md)
 - [GadShape](docs/GadShape.md)
 - [GeoServiceArea](docs/GeoServiceArea.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [GmCapable](docs/GmCapable.md)
 - [InstanceType](docs/InstanceType.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [NfType](docs/NfType.md)
 - [NoProfileMatchInfo](docs/NoProfileMatchInfo.md)
 - [NoProfileMatchReason](docs/NoProfileMatchReason.md)
 - [NotificationMethod](docs/NotificationMethod.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [Point](docs/Point.md)
 - [PointAltitude](docs/PointAltitude.md)
 - [PointAltitudeUncertainty](docs/PointAltitudeUncertainty.md)
 - [PointUncertaintyCircle](docs/PointUncertaintyCircle.md)
 - [PointUncertaintyEllipse](docs/PointUncertaintyEllipse.md)
 - [Polygon](docs/Polygon.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [Protocol](docs/Protocol.md)
 - [PtpCapabilitiesPerUe](docs/PtpCapabilitiesPerUe.md)
 - [PtpInstance](docs/PtpInstance.md)
 - [QueryParamCombination](docs/QueryParamCombination.md)
 - [QueryParameter](docs/QueryParameter.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [ServiceAreaCoverageInfo](docs/ServiceAreaCoverageInfo.md)
 - [Snssai](docs/Snssai.md)
 - [SpatialValidityCond](docs/SpatialValidityCond.md)
 - [StateOfConfiguration](docs/StateOfConfiguration.md)
 - [StateOfDstt](docs/StateOfDstt.md)
 - [SubsEventNotification](docs/SubsEventNotification.md)
 - [SubscribedEvent](docs/SubscribedEvent.md)
 - [SupportedGadShapes](docs/SupportedGadShapes.md)
 - [SynchronizationState](docs/SynchronizationState.md)
 - [Tai](docs/Tai.md)
 - [TemporalValidity](docs/TemporalValidity.md)
 - [TimeSource](docs/TimeSource.md)
 - [TimeSyncCapability](docs/TimeSyncCapability.md)
 - [TimeSyncExposureConfig](docs/TimeSyncExposureConfig.md)
 - [TimeSyncExposureConfigNotif](docs/TimeSyncExposureConfigNotif.md)
 - [TimeSyncExposureSubsNotif](docs/TimeSyncExposureSubsNotif.md)
 - [TimeSyncExposureSubsc](docs/TimeSyncExposureSubsc.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [WebsockNotifConfig](docs/WebsockNotifConfig.md)


## Documentation For Authorization



## oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **ntsctsf-timesynchronization**: Access to the Ntsctsf_TimeSynchronization API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Author



