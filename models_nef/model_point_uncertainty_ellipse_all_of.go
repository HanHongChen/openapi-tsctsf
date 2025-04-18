/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/HanHongChen/openapi-tsctsf/models"
)

type PointUncertaintyEllipseAllOf struct {
	Point models.GeographicalCoordinates `json:"point"`

	UncertaintyEllipse models.UncertaintyEllipse `json:"uncertaintyEllipse"`

	Confidence int32 `json:"confidence"`
}
