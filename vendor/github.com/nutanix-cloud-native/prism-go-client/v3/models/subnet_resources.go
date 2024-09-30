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

// SubnetResources Subnet creation/modification spec
//
// Subnet creation/modification spec.
//
// swagger:model subnet_resources
type SubnetResources struct {

	// Whether the subnet should be realized on OVN stack or not.
	AdvancedNetworking bool `json:"advanced_networking,omitempty"`

	// List of availability zones from which resources are derived (Only supported on Xi).
	//
	AvailabilityZoneReferenceList []*AvailabilityZoneReference `json:"availability_zone_reference_list"`

	// Whether NAT should be performed for VPCs attaching to the subnet. This field is supported only for external subnets. NAT is enabled by default on external subnets.
	//
	EnableNat bool `json:"enable_nat,omitempty"`

	// External connectivity state (Only supported on Xi)
	ExternalConnectivityState string `json:"external_connectivity_state,omitempty"`

	// ip config
	IPConfig *IPConfig `json:"ip_config,omitempty"`

	// Whether the subnet is external subnet or not.
	IsExternal bool `json:"is_external,omitempty"`

	// network function chain reference
	NetworkFunctionChainReference *NetworkFunctionChainReference `json:"network_function_chain_reference,omitempty"`

	// List of IPs that are not considered while allocating IP addresses to Atlas ports.
	//
	ReservedIPAddressList []string `json:"reserved_ip_address_list"`

	// subnet type
	// Required: true
	SubnetType *string `json:"subnet_type"`

	// The virtual network this subnet belongs to (Only supported on Xi). This reference is deprecated, use vpc_reference instead.
	//
	VirtualNetworkReference *VirtualNetworkReference `json:"virtual_network_reference,omitempty"`

	// Reference to virtual switch
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	VirtualSwitchUUID string `json:"virtual_switch_uuid,omitempty"`

	// vlan id
	// Minimum: 0
	VlanID *int32 `json:"vlan_id,omitempty"`

	// The VPC this subnet belongs to (Only supported on Xi).
	//
	VpcReference *VpcReference `json:"vpc_reference,omitempty"`

	// vswitch name
	// Max Length: 64
	VswitchName string `json:"vswitch_name,omitempty"`
}

// Validate validates this subnet resources
func (m *SubnetResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZoneReferenceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkFunctionChainReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservedIPAddressList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualSwitchUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVswitchName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubnetResources) validateAvailabilityZoneReferenceList(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailabilityZoneReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailabilityZoneReferenceList); i++ {
		if swag.IsZero(m.AvailabilityZoneReferenceList[i]) { // not required
			continue
		}

		if m.AvailabilityZoneReferenceList[i] != nil {
			if err := m.AvailabilityZoneReferenceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availability_zone_reference_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availability_zone_reference_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubnetResources) validateIPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.IPConfig) { // not required
		return nil
	}

	if m.IPConfig != nil {
		if err := m.IPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip_config")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) validateNetworkFunctionChainReference(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkFunctionChainReference) { // not required
		return nil
	}

	if m.NetworkFunctionChainReference != nil {
		if err := m.NetworkFunctionChainReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_function_chain_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_function_chain_reference")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) validateReservedIPAddressList(formats strfmt.Registry) error {
	if swag.IsZero(m.ReservedIPAddressList) { // not required
		return nil
	}

	for i := 0; i < len(m.ReservedIPAddressList); i++ {

		if err := validate.Pattern("reserved_ip_address_list"+"."+strconv.Itoa(i), "body", m.ReservedIPAddressList[i], `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *SubnetResources) validateSubnetType(formats strfmt.Registry) error {

	if err := validate.Required("subnet_type", "body", m.SubnetType); err != nil {
		return err
	}

	return nil
}

func (m *SubnetResources) validateVirtualNetworkReference(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualNetworkReference) { // not required
		return nil
	}

	if m.VirtualNetworkReference != nil {
		if err := m.VirtualNetworkReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_network_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual_network_reference")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) validateVirtualSwitchUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualSwitchUUID) { // not required
		return nil
	}

	if err := validate.Pattern("virtual_switch_uuid", "body", m.VirtualSwitchUUID, `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *SubnetResources) validateVlanID(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanID) { // not required
		return nil
	}

	if err := validate.MinimumInt("vlan_id", "body", int64(*m.VlanID), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SubnetResources) validateVpcReference(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcReference) { // not required
		return nil
	}

	if m.VpcReference != nil {
		if err := m.VpcReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_reference")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) validateVswitchName(formats strfmt.Registry) error {
	if swag.IsZero(m.VswitchName) { // not required
		return nil
	}

	if err := validate.MaxLength("vswitch_name", "body", m.VswitchName, 64); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this subnet resources based on the context it is used
func (m *SubnetResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailabilityZoneReferenceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkFunctionChainReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualNetworkReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubnetResources) contextValidateAvailabilityZoneReferenceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailabilityZoneReferenceList); i++ {

		if m.AvailabilityZoneReferenceList[i] != nil {

			if swag.IsZero(m.AvailabilityZoneReferenceList[i]) { // not required
				return nil
			}

			if err := m.AvailabilityZoneReferenceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availability_zone_reference_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availability_zone_reference_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubnetResources) contextValidateIPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.IPConfig != nil {

		if swag.IsZero(m.IPConfig) { // not required
			return nil
		}

		if err := m.IPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip_config")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) contextValidateNetworkFunctionChainReference(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkFunctionChainReference != nil {

		if swag.IsZero(m.NetworkFunctionChainReference) { // not required
			return nil
		}

		if err := m.NetworkFunctionChainReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_function_chain_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_function_chain_reference")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) contextValidateVirtualNetworkReference(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualNetworkReference != nil {

		if swag.IsZero(m.VirtualNetworkReference) { // not required
			return nil
		}

		if err := m.VirtualNetworkReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_network_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual_network_reference")
			}
			return err
		}
	}

	return nil
}

func (m *SubnetResources) contextValidateVpcReference(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcReference != nil {

		if swag.IsZero(m.VpcReference) { // not required
			return nil
		}

		if err := m.VpcReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_reference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubnetResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubnetResources) UnmarshalBinary(b []byte) error {
	var res SubnetResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
