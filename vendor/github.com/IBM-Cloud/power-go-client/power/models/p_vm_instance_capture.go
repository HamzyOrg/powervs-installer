// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceCapture p VM instance capture
// swagger:model PVMInstanceCapture
type PVMInstanceCapture struct {

	// Destination for the deployable image
	// Required: true
	// Enum: [image-catalog cloud-storage both]
	CaptureDestination *string `json:"captureDestination"`

	// Name of the deployable image created for the captured PVMInstance
	// Required: true
	CaptureName *string `json:"captureName"`

	// List of Data volume IDs to include in the captured PVMInstance
	CaptureVolumeIds []string `json:"captureVolumeIDs"`

	// Cloud Storage Access key
	CloudStorageAccessKey string `json:"cloudStorageAccessKey,omitempty"`

	// Cloud Storage Image Path (bucket-name [/folder/../..])
	CloudStorageImagePath string `json:"cloudStorageImagePath,omitempty"`

	// Cloud Storage Region
	CloudStorageRegion string `json:"cloudStorageRegion,omitempty"`

	// Cloud Storage Secret key
	CloudStorageSecretKey string `json:"cloudStorageSecretKey,omitempty"`
}

// Validate validates this p VM instance capture
func (m *PVMInstanceCapture) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaptureDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pVmInstanceCaptureTypeCaptureDestinationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["image-catalog","cloud-storage","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCaptureTypeCaptureDestinationPropEnum = append(pVmInstanceCaptureTypeCaptureDestinationPropEnum, v)
	}
}

const (

	// PVMInstanceCaptureCaptureDestinationImageCatalog captures enum value "image-catalog"
	PVMInstanceCaptureCaptureDestinationImageCatalog string = "image-catalog"

	// PVMInstanceCaptureCaptureDestinationCloudStorage captures enum value "cloud-storage"
	PVMInstanceCaptureCaptureDestinationCloudStorage string = "cloud-storage"

	// PVMInstanceCaptureCaptureDestinationBoth captures enum value "both"
	PVMInstanceCaptureCaptureDestinationBoth string = "both"
)

// prop value enum
func (m *PVMInstanceCapture) validateCaptureDestinationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pVmInstanceCaptureTypeCaptureDestinationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceCapture) validateCaptureDestination(formats strfmt.Registry) error {

	if err := validate.Required("captureDestination", "body", m.CaptureDestination); err != nil {
		return err
	}

	// value enum
	if err := m.validateCaptureDestinationEnum("captureDestination", "body", *m.CaptureDestination); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceCapture) validateCaptureName(formats strfmt.Registry) error {

	if err := validate.Required("captureName", "body", m.CaptureName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceCapture) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceCapture) UnmarshalBinary(b []byte) error {
	var res PVMInstanceCapture
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
