// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DhcpOptions DHCP options
//
// Spec for defining DHCP options.
//
// swagger:model dhcp_options
type DhcpOptions struct {

	// boot file name
	BootFileName string `json:"boot_file_name,omitempty"`

	// domain name
	DomainName string `json:"domain_name,omitempty"`

	// domain name server list
	// Max Items: 32
	DomainNameServerList []string `json:"domain_name_server_list"`

	// domain search list
	// Max Items: 32
	DomainSearchList []string `json:"domain_search_list"`

	// tftp server name
	TftpServerName string `json:"tftp_server_name,omitempty"`
}

// Validate validates this dhcp options
func (m *DhcpOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainNameServerList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainSearchList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpOptions) validateDomainNameServerList(formats strfmt.Registry) error {
	if swag.IsZero(m.DomainNameServerList) { // not required
		return nil
	}

	iDomainNameServerListSize := int64(len(m.DomainNameServerList))

	if err := validate.MaxItems("domain_name_server_list", "body", iDomainNameServerListSize, 32); err != nil {
		return err
	}

	for i := 0; i < len(m.DomainNameServerList); i++ {

		if err := validate.Pattern("domain_name_server_list"+"."+strconv.Itoa(i), "body", m.DomainNameServerList[i], `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *DhcpOptions) validateDomainSearchList(formats strfmt.Registry) error {
	if swag.IsZero(m.DomainSearchList) { // not required
		return nil
	}

	iDomainSearchListSize := int64(len(m.DomainSearchList))

	if err := validate.MaxItems("domain_search_list", "body", iDomainSearchListSize, 32); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dhcp options based on context it is used
func (m *DhcpOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DhcpOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DhcpOptions) UnmarshalBinary(b []byte) error {
	var res DhcpOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}