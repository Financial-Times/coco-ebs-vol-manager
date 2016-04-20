package mocks

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/mock"
)

type MockEC2 struct {
	mock.Mock
}

func (_m *MockEC2) AcceptVpcPeeringConnectionRequest(_a0 *ec2.AcceptVpcPeeringConnectionInput) (*request.Request, *ec2.AcceptVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AcceptVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AcceptVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.AcceptVpcPeeringConnectionInput) *ec2.AcceptVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AcceptVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AcceptVpcPeeringConnection(_a0 *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AcceptVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.AcceptVpcPeeringConnectionInput) *ec2.AcceptVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AcceptVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AllocateAddressRequest(_a0 *ec2.AllocateAddressInput) (*request.Request, *ec2.AllocateAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AllocateAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AllocateAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AllocateAddressInput) *ec2.AllocateAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AllocateAddressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AllocateAddress(_a0 *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AllocateAddressOutput
	if rf, ok := ret.Get(0).(func(*ec2.AllocateAddressInput) *ec2.AllocateAddressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AllocateAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AllocateAddressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AllocateHostsRequest(_a0 *ec2.AllocateHostsInput) (*request.Request, *ec2.AllocateHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AllocateHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AllocateHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.AllocateHostsInput) *ec2.AllocateHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AllocateHostsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AllocateHosts(_a0 *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AllocateHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.AllocateHostsInput) *ec2.AllocateHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AllocateHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AllocateHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AssignPrivateIpAddressesRequest(_a0 *ec2.AssignPrivateIpAddressesInput) (*request.Request, *ec2.AssignPrivateIpAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssignPrivateIpAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssignPrivateIpAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssignPrivateIpAddressesInput) *ec2.AssignPrivateIpAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssignPrivateIpAddressesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AssignPrivateIpAddresses(_a0 *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssignPrivateIpAddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssignPrivateIpAddressesInput) *ec2.AssignPrivateIpAddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssignPrivateIpAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssignPrivateIpAddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AssociateAddressRequest(_a0 *ec2.AssociateAddressInput) (*request.Request, *ec2.AssociateAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateAddressInput) *ec2.AssociateAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateAddressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AssociateAddress(_a0 *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateAddressOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateAddressInput) *ec2.AssociateAddressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateAddressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AssociateDhcpOptionsRequest(_a0 *ec2.AssociateDhcpOptionsInput) (*request.Request, *ec2.AssociateDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateDhcpOptionsInput) *ec2.AssociateDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateDhcpOptionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AssociateDhcpOptions(_a0 *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateDhcpOptionsInput) *ec2.AssociateDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AssociateRouteTableRequest(_a0 *ec2.AssociateRouteTableInput) (*request.Request, *ec2.AssociateRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateRouteTableInput) *ec2.AssociateRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateRouteTableOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AssociateRouteTable(_a0 *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateRouteTableInput) *ec2.AssociateRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AttachClassicLinkVpcRequest(_a0 *ec2.AttachClassicLinkVpcInput) (*request.Request, *ec2.AttachClassicLinkVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachClassicLinkVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachClassicLinkVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachClassicLinkVpcInput) *ec2.AttachClassicLinkVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachClassicLinkVpcOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AttachClassicLinkVpc(_a0 *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachClassicLinkVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachClassicLinkVpcInput) *ec2.AttachClassicLinkVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachClassicLinkVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachClassicLinkVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AttachInternetGatewayRequest(_a0 *ec2.AttachInternetGatewayInput) (*request.Request, *ec2.AttachInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachInternetGatewayInput) *ec2.AttachInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachInternetGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AttachInternetGateway(_a0 *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachInternetGatewayInput) *ec2.AttachInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AttachNetworkInterfaceRequest(_a0 *ec2.AttachNetworkInterfaceInput) (*request.Request, *ec2.AttachNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachNetworkInterfaceInput) *ec2.AttachNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachNetworkInterfaceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AttachNetworkInterface(_a0 *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachNetworkInterfaceInput) *ec2.AttachNetworkInterfaceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachNetworkInterfaceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AttachVolumeRequest(_a0 *ec2.AttachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.VolumeAttachment
	if rf, ok := ret.Get(1).(func(*ec2.AttachVolumeInput) *ec2.VolumeAttachment); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.VolumeAttachment)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AttachVolume(_a0 *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.VolumeAttachment
	if rf, ok := ret.Get(0).(func(*ec2.AttachVolumeInput) *ec2.VolumeAttachment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AttachVpnGatewayRequest(_a0 *ec2.AttachVpnGatewayInput) (*request.Request, *ec2.AttachVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachVpnGatewayInput) *ec2.AttachVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachVpnGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AttachVpnGateway(_a0 *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachVpnGatewayInput) *ec2.AttachVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AuthorizeSecurityGroupEgressRequest(_a0 *ec2.AuthorizeSecurityGroupEgressInput) (*request.Request, *ec2.AuthorizeSecurityGroupEgressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupEgressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AuthorizeSecurityGroupEgressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupEgressInput) *ec2.AuthorizeSecurityGroupEgressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AuthorizeSecurityGroupEgressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AuthorizeSecurityGroupEgress(_a0 *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AuthorizeSecurityGroupEgressOutput
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupEgressInput) *ec2.AuthorizeSecurityGroupEgressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AuthorizeSecurityGroupEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupEgressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) AuthorizeSecurityGroupIngressRequest(_a0 *ec2.AuthorizeSecurityGroupIngressInput) (*request.Request, *ec2.AuthorizeSecurityGroupIngressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupIngressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AuthorizeSecurityGroupIngressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupIngressInput) *ec2.AuthorizeSecurityGroupIngressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AuthorizeSecurityGroupIngressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) AuthorizeSecurityGroupIngress(_a0 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AuthorizeSecurityGroupIngressOutput
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupIngressInput) *ec2.AuthorizeSecurityGroupIngressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AuthorizeSecurityGroupIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupIngressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) BundleInstanceRequest(_a0 *ec2.BundleInstanceInput) (*request.Request, *ec2.BundleInstanceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.BundleInstanceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.BundleInstanceOutput
	if rf, ok := ret.Get(1).(func(*ec2.BundleInstanceInput) *ec2.BundleInstanceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.BundleInstanceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) BundleInstance(_a0 *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.BundleInstanceOutput
	if rf, ok := ret.Get(0).(func(*ec2.BundleInstanceInput) *ec2.BundleInstanceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.BundleInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.BundleInstanceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelBundleTaskRequest(_a0 *ec2.CancelBundleTaskInput) (*request.Request, *ec2.CancelBundleTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelBundleTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelBundleTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelBundleTaskInput) *ec2.CancelBundleTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelBundleTaskOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelBundleTask(_a0 *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelBundleTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelBundleTaskInput) *ec2.CancelBundleTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelBundleTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelBundleTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelConversionTaskRequest(_a0 *ec2.CancelConversionTaskInput) (*request.Request, *ec2.CancelConversionTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelConversionTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelConversionTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelConversionTaskInput) *ec2.CancelConversionTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelConversionTaskOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelConversionTask(_a0 *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelConversionTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelConversionTaskInput) *ec2.CancelConversionTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelConversionTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelConversionTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelExportTaskRequest(_a0 *ec2.CancelExportTaskInput) (*request.Request, *ec2.CancelExportTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelExportTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelExportTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelExportTaskInput) *ec2.CancelExportTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelExportTaskOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelExportTask(_a0 *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelExportTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelExportTaskInput) *ec2.CancelExportTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelExportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelExportTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelImportTaskRequest(_a0 *ec2.CancelImportTaskInput) (*request.Request, *ec2.CancelImportTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelImportTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelImportTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelImportTaskInput) *ec2.CancelImportTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelImportTaskOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelImportTask(_a0 *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelImportTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelImportTaskInput) *ec2.CancelImportTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelImportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelImportTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelReservedInstancesListingRequest(_a0 *ec2.CancelReservedInstancesListingInput) (*request.Request, *ec2.CancelReservedInstancesListingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelReservedInstancesListingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelReservedInstancesListingOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelReservedInstancesListingInput) *ec2.CancelReservedInstancesListingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelReservedInstancesListingOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelReservedInstancesListing(_a0 *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelReservedInstancesListingOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelReservedInstancesListingInput) *ec2.CancelReservedInstancesListingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelReservedInstancesListingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelReservedInstancesListingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelSpotFleetRequestsRequest(_a0 *ec2.CancelSpotFleetRequestsInput) (*request.Request, *ec2.CancelSpotFleetRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotFleetRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelSpotFleetRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotFleetRequestsInput) *ec2.CancelSpotFleetRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelSpotFleetRequestsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelSpotFleetRequests(_a0 *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelSpotFleetRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotFleetRequestsInput) *ec2.CancelSpotFleetRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelSpotFleetRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotFleetRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CancelSpotInstanceRequestsRequest(_a0 *ec2.CancelSpotInstanceRequestsInput) (*request.Request, *ec2.CancelSpotInstanceRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotInstanceRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelSpotInstanceRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotInstanceRequestsInput) *ec2.CancelSpotInstanceRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelSpotInstanceRequestsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CancelSpotInstanceRequests(_a0 *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelSpotInstanceRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotInstanceRequestsInput) *ec2.CancelSpotInstanceRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelSpotInstanceRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotInstanceRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ConfirmProductInstanceRequest(_a0 *ec2.ConfirmProductInstanceInput) (*request.Request, *ec2.ConfirmProductInstanceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ConfirmProductInstanceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ConfirmProductInstanceOutput
	if rf, ok := ret.Get(1).(func(*ec2.ConfirmProductInstanceInput) *ec2.ConfirmProductInstanceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ConfirmProductInstanceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ConfirmProductInstance(_a0 *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ConfirmProductInstanceOutput
	if rf, ok := ret.Get(0).(func(*ec2.ConfirmProductInstanceInput) *ec2.ConfirmProductInstanceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ConfirmProductInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ConfirmProductInstanceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CopyImageRequest(_a0 *ec2.CopyImageInput) (*request.Request, *ec2.CopyImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CopyImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CopyImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.CopyImageInput) *ec2.CopyImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CopyImageOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CopyImage(_a0 *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CopyImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.CopyImageInput) *ec2.CopyImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopyImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CopyImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CopySnapshotRequest(_a0 *ec2.CopySnapshotInput) (*request.Request, *ec2.CopySnapshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CopySnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CopySnapshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.CopySnapshotInput) *ec2.CopySnapshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CopySnapshotOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CopySnapshot(_a0 *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CopySnapshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.CopySnapshotInput) *ec2.CopySnapshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopySnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CopySnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateCustomerGatewayRequest(_a0 *ec2.CreateCustomerGatewayInput) (*request.Request, *ec2.CreateCustomerGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateCustomerGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateCustomerGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateCustomerGatewayInput) *ec2.CreateCustomerGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateCustomerGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateCustomerGateway(_a0 *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateCustomerGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateCustomerGatewayInput) *ec2.CreateCustomerGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateCustomerGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateCustomerGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateDhcpOptionsRequest(_a0 *ec2.CreateDhcpOptionsInput) (*request.Request, *ec2.CreateDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateDhcpOptionsInput) *ec2.CreateDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateDhcpOptionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateDhcpOptions(_a0 *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateDhcpOptionsInput) *ec2.CreateDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateFlowLogsRequest(_a0 *ec2.CreateFlowLogsInput) (*request.Request, *ec2.CreateFlowLogsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateFlowLogsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateFlowLogsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateFlowLogsInput) *ec2.CreateFlowLogsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateFlowLogsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateFlowLogs(_a0 *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateFlowLogsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateFlowLogsInput) *ec2.CreateFlowLogsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateFlowLogsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateImageRequest(_a0 *ec2.CreateImageInput) (*request.Request, *ec2.CreateImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateImageInput) *ec2.CreateImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateImageOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateImage(_a0 *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateImageInput) *ec2.CreateImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateInstanceExportTaskRequest(_a0 *ec2.CreateInstanceExportTaskInput) (*request.Request, *ec2.CreateInstanceExportTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateInstanceExportTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateInstanceExportTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateInstanceExportTaskInput) *ec2.CreateInstanceExportTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateInstanceExportTaskOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateInstanceExportTask(_a0 *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateInstanceExportTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateInstanceExportTaskInput) *ec2.CreateInstanceExportTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateInstanceExportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateInstanceExportTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateInternetGatewayRequest(_a0 *ec2.CreateInternetGatewayInput) (*request.Request, *ec2.CreateInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateInternetGatewayInput) *ec2.CreateInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateInternetGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateInternetGateway(_a0 *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateInternetGatewayInput) *ec2.CreateInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateKeyPairRequest(_a0 *ec2.CreateKeyPairInput) (*request.Request, *ec2.CreateKeyPairOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateKeyPairInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateKeyPairOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateKeyPairInput) *ec2.CreateKeyPairOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateKeyPairOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateKeyPair(_a0 *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateKeyPairOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateKeyPairInput) *ec2.CreateKeyPairOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateKeyPairInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateNatGatewayRequest(_a0 *ec2.CreateNatGatewayInput) (*request.Request, *ec2.CreateNatGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNatGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNatGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNatGatewayInput) *ec2.CreateNatGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNatGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateNatGateway(_a0 *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNatGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNatGatewayInput) *ec2.CreateNatGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNatGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNatGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateNetworkAclRequest(_a0 *ec2.CreateNetworkAclInput) (*request.Request, *ec2.CreateNetworkAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkAclOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclInput) *ec2.CreateNetworkAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkAclOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateNetworkAcl(_a0 *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNetworkAclOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclInput) *ec2.CreateNetworkAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateNetworkAclEntryRequest(_a0 *ec2.CreateNetworkAclEntryInput) (*request.Request, *ec2.CreateNetworkAclEntryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclEntryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkAclEntryOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclEntryInput) *ec2.CreateNetworkAclEntryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkAclEntryOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateNetworkAclEntry(_a0 *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclEntryInput) *ec2.CreateNetworkAclEntryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclEntryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateNetworkInterfaceRequest(_a0 *ec2.CreateNetworkInterfaceInput) (*request.Request, *ec2.CreateNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkInterfaceInput) *ec2.CreateNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkInterfaceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateNetworkInterface(_a0 *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkInterfaceInput) *ec2.CreateNetworkInterfaceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkInterfaceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreatePlacementGroupRequest(_a0 *ec2.CreatePlacementGroupInput) (*request.Request, *ec2.CreatePlacementGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreatePlacementGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreatePlacementGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreatePlacementGroupInput) *ec2.CreatePlacementGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreatePlacementGroupOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreatePlacementGroup(_a0 *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreatePlacementGroupOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreatePlacementGroupInput) *ec2.CreatePlacementGroupOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreatePlacementGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreatePlacementGroupInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateReservedInstancesListingRequest(_a0 *ec2.CreateReservedInstancesListingInput) (*request.Request, *ec2.CreateReservedInstancesListingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateReservedInstancesListingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateReservedInstancesListingOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateReservedInstancesListingInput) *ec2.CreateReservedInstancesListingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateReservedInstancesListingOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateReservedInstancesListing(_a0 *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateReservedInstancesListingOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateReservedInstancesListingInput) *ec2.CreateReservedInstancesListingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateReservedInstancesListingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateReservedInstancesListingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateRouteRequest(_a0 *ec2.CreateRouteInput) (*request.Request, *ec2.CreateRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteInput) *ec2.CreateRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateRouteOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateRoute(_a0 *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteInput) *ec2.CreateRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateRouteTableRequest(_a0 *ec2.CreateRouteTableInput) (*request.Request, *ec2.CreateRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteTableInput) *ec2.CreateRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateRouteTableOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateRouteTable(_a0 *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteTableInput) *ec2.CreateRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateSecurityGroupRequest(_a0 *ec2.CreateSecurityGroupInput) (*request.Request, *ec2.CreateSecurityGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSecurityGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateSecurityGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateSecurityGroupInput) *ec2.CreateSecurityGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateSecurityGroupOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateSecurityGroup(_a0 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateSecurityGroupOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateSecurityGroupInput) *ec2.CreateSecurityGroupOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSecurityGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateSecurityGroupInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateSnapshotRequest(_a0 *ec2.CreateSnapshotInput) (*request.Request, *ec2.Snapshot) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.Snapshot
	if rf, ok := ret.Get(1).(func(*ec2.CreateSnapshotInput) *ec2.Snapshot); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.Snapshot)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateSnapshot(_a0 *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Snapshot
	if rf, ok := ret.Get(0).(func(*ec2.CreateSnapshotInput) *ec2.Snapshot); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateSnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateSpotDatafeedSubscriptionRequest(_a0 *ec2.CreateSpotDatafeedSubscriptionInput) (*request.Request, *ec2.CreateSpotDatafeedSubscriptionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSpotDatafeedSubscriptionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateSpotDatafeedSubscriptionInput) *ec2.CreateSpotDatafeedSubscriptionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateSpotDatafeedSubscriptionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateSpotDatafeedSubscription(_a0 *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateSpotDatafeedSubscriptionInput) *ec2.CreateSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateSpotDatafeedSubscriptionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateSubnetRequest(_a0 *ec2.CreateSubnetInput) (*request.Request, *ec2.CreateSubnetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSubnetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateSubnetOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateSubnetInput) *ec2.CreateSubnetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateSubnetOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateSubnet(_a0 *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateSubnetOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateSubnetInput) *ec2.CreateSubnetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateSubnetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateTagsRequest(_a0 *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateTagsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateTagsInput) *ec2.CreateTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateTagsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateTags(_a0 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateTagsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateTagsInput) *ec2.CreateTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVolumeRequest(_a0 *ec2.CreateVolumeInput) (*request.Request, *ec2.Volume) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.Volume
	if rf, ok := ret.Get(1).(func(*ec2.CreateVolumeInput) *ec2.Volume); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.Volume)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVolume(_a0 *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Volume
	if rf, ok := ret.Get(0).(func(*ec2.CreateVolumeInput) *ec2.Volume); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpcRequest(_a0 *ec2.CreateVpcInput) (*request.Request, *ec2.CreateVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcInput) *ec2.CreateVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpc(_a0 *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcInput) *ec2.CreateVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpcEndpointRequest(_a0 *ec2.CreateVpcEndpointInput) (*request.Request, *ec2.CreateVpcEndpointOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcEndpointOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointInput) *ec2.CreateVpcEndpointOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcEndpointOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpcEndpoint(_a0 *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcEndpointOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointInput) *ec2.CreateVpcEndpointOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpcPeeringConnectionRequest(_a0 *ec2.CreateVpcPeeringConnectionInput) (*request.Request, *ec2.CreateVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcPeeringConnectionInput) *ec2.CreateVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpcPeeringConnection(_a0 *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcPeeringConnectionInput) *ec2.CreateVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpnConnectionRequest(_a0 *ec2.CreateVpnConnectionInput) (*request.Request, *ec2.CreateVpnConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpnConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionInput) *ec2.CreateVpnConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpnConnectionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpnConnection(_a0 *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpnConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionInput) *ec2.CreateVpnConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpnConnectionRouteRequest(_a0 *ec2.CreateVpnConnectionRouteInput) (*request.Request, *ec2.CreateVpnConnectionRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpnConnectionRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionRouteInput) *ec2.CreateVpnConnectionRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpnConnectionRouteOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpnConnectionRoute(_a0 *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpnConnectionRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionRouteInput) *ec2.CreateVpnConnectionRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnConnectionRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpnGatewayRequest(_a0 *ec2.CreateVpnGatewayInput) (*request.Request, *ec2.CreateVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnGatewayInput) *ec2.CreateVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpnGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) CreateVpnGateway(_a0 *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnGatewayInput) *ec2.CreateVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteCustomerGatewayRequest(_a0 *ec2.DeleteCustomerGatewayInput) (*request.Request, *ec2.DeleteCustomerGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteCustomerGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteCustomerGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteCustomerGatewayInput) *ec2.DeleteCustomerGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteCustomerGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteCustomerGateway(_a0 *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteCustomerGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteCustomerGatewayInput) *ec2.DeleteCustomerGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteCustomerGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteCustomerGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteDhcpOptionsRequest(_a0 *ec2.DeleteDhcpOptionsInput) (*request.Request, *ec2.DeleteDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteDhcpOptionsInput) *ec2.DeleteDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteDhcpOptionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteDhcpOptions(_a0 *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteDhcpOptionsInput) *ec2.DeleteDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteFlowLogsRequest(_a0 *ec2.DeleteFlowLogsInput) (*request.Request, *ec2.DeleteFlowLogsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFlowLogsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteFlowLogsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFlowLogsInput) *ec2.DeleteFlowLogsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteFlowLogsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteFlowLogs(_a0 *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteFlowLogsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFlowLogsInput) *ec2.DeleteFlowLogsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFlowLogsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteInternetGatewayRequest(_a0 *ec2.DeleteInternetGatewayInput) (*request.Request, *ec2.DeleteInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteInternetGatewayInput) *ec2.DeleteInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteInternetGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteInternetGateway(_a0 *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteInternetGatewayInput) *ec2.DeleteInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteKeyPairRequest(_a0 *ec2.DeleteKeyPairInput) (*request.Request, *ec2.DeleteKeyPairOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteKeyPairInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteKeyPairOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteKeyPairInput) *ec2.DeleteKeyPairOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteKeyPairOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteKeyPair(_a0 *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteKeyPairOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteKeyPairInput) *ec2.DeleteKeyPairOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteKeyPairInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNatGatewayRequest(_a0 *ec2.DeleteNatGatewayInput) (*request.Request, *ec2.DeleteNatGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNatGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNatGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNatGatewayInput) *ec2.DeleteNatGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNatGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNatGateway(_a0 *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNatGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNatGatewayInput) *ec2.DeleteNatGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNatGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNatGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNetworkAclRequest(_a0 *ec2.DeleteNetworkAclInput) (*request.Request, *ec2.DeleteNetworkAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkAclOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclInput) *ec2.DeleteNetworkAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkAclOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNetworkAcl(_a0 *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNetworkAclOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclInput) *ec2.DeleteNetworkAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNetworkAclEntryRequest(_a0 *ec2.DeleteNetworkAclEntryInput) (*request.Request, *ec2.DeleteNetworkAclEntryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclEntryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkAclEntryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclEntryInput) *ec2.DeleteNetworkAclEntryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkAclEntryOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNetworkAclEntry(_a0 *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclEntryInput) *ec2.DeleteNetworkAclEntryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclEntryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNetworkInterfaceRequest(_a0 *ec2.DeleteNetworkInterfaceInput) (*request.Request, *ec2.DeleteNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkInterfaceInput) *ec2.DeleteNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkInterfaceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteNetworkInterface(_a0 *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkInterfaceInput) *ec2.DeleteNetworkInterfaceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkInterfaceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeletePlacementGroupRequest(_a0 *ec2.DeletePlacementGroupInput) (*request.Request, *ec2.DeletePlacementGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeletePlacementGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeletePlacementGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeletePlacementGroupInput) *ec2.DeletePlacementGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeletePlacementGroupOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeletePlacementGroup(_a0 *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeletePlacementGroupOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeletePlacementGroupInput) *ec2.DeletePlacementGroupOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeletePlacementGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeletePlacementGroupInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteRouteRequest(_a0 *ec2.DeleteRouteInput) (*request.Request, *ec2.DeleteRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteInput) *ec2.DeleteRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteRouteOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteRoute(_a0 *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteInput) *ec2.DeleteRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteRouteTableRequest(_a0 *ec2.DeleteRouteTableInput) (*request.Request, *ec2.DeleteRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteTableInput) *ec2.DeleteRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteRouteTableOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteRouteTable(_a0 *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteTableInput) *ec2.DeleteRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSecurityGroupRequest(_a0 *ec2.DeleteSecurityGroupInput) (*request.Request, *ec2.DeleteSecurityGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSecurityGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSecurityGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSecurityGroupInput) *ec2.DeleteSecurityGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSecurityGroupOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSecurityGroup(_a0 *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSecurityGroupOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSecurityGroupInput) *ec2.DeleteSecurityGroupOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSecurityGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSecurityGroupInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSnapshotRequest(_a0 *ec2.DeleteSnapshotInput) (*request.Request, *ec2.DeleteSnapshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSnapshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSnapshotInput) *ec2.DeleteSnapshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSnapshotOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSnapshot(_a0 *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSnapshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSnapshotInput) *ec2.DeleteSnapshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSpotDatafeedSubscriptionRequest(_a0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) *ec2.DeleteSpotDatafeedSubscriptionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSpotDatafeedSubscriptionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSpotDatafeedSubscription(_a0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) *ec2.DeleteSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSubnetRequest(_a0 *ec2.DeleteSubnetInput) (*request.Request, *ec2.DeleteSubnetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSubnetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSubnetOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSubnetInput) *ec2.DeleteSubnetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSubnetOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteSubnet(_a0 *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSubnetOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSubnetInput) *ec2.DeleteSubnetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSubnetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteTagsRequest(_a0 *ec2.DeleteTagsInput) (*request.Request, *ec2.DeleteTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteTagsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteTagsInput) *ec2.DeleteTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteTagsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteTags(_a0 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteTagsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteTagsInput) *ec2.DeleteTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVolumeRequest(_a0 *ec2.DeleteVolumeInput) (*request.Request, *ec2.DeleteVolumeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVolumeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVolumeInput) *ec2.DeleteVolumeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVolumeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVolume(_a0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVolumeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVolumeInput) *ec2.DeleteVolumeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpcRequest(_a0 *ec2.DeleteVpcInput) (*request.Request, *ec2.DeleteVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcInput) *ec2.DeleteVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpc(_a0 *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcInput) *ec2.DeleteVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpcEndpointsRequest(_a0 *ec2.DeleteVpcEndpointsInput) (*request.Request, *ec2.DeleteVpcEndpointsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcEndpointsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointsInput) *ec2.DeleteVpcEndpointsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcEndpointsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpcEndpoints(_a0 *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcEndpointsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointsInput) *ec2.DeleteVpcEndpointsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpcPeeringConnectionRequest(_a0 *ec2.DeleteVpcPeeringConnectionInput) (*request.Request, *ec2.DeleteVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcPeeringConnectionInput) *ec2.DeleteVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpcPeeringConnection(_a0 *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcPeeringConnectionInput) *ec2.DeleteVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpnConnectionRequest(_a0 *ec2.DeleteVpnConnectionInput) (*request.Request, *ec2.DeleteVpnConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpnConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionInput) *ec2.DeleteVpnConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpnConnectionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpnConnection(_a0 *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpnConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionInput) *ec2.DeleteVpnConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpnConnectionRouteRequest(_a0 *ec2.DeleteVpnConnectionRouteInput) (*request.Request, *ec2.DeleteVpnConnectionRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpnConnectionRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionRouteInput) *ec2.DeleteVpnConnectionRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpnConnectionRouteOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpnConnectionRoute(_a0 *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpnConnectionRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionRouteInput) *ec2.DeleteVpnConnectionRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnConnectionRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpnGatewayRequest(_a0 *ec2.DeleteVpnGatewayInput) (*request.Request, *ec2.DeleteVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnGatewayInput) *ec2.DeleteVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpnGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeleteVpnGateway(_a0 *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnGatewayInput) *ec2.DeleteVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DeregisterImageRequest(_a0 *ec2.DeregisterImageInput) (*request.Request, *ec2.DeregisterImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeregisterImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeregisterImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeregisterImageInput) *ec2.DeregisterImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeregisterImageOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DeregisterImage(_a0 *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeregisterImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeregisterImageInput) *ec2.DeregisterImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeregisterImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeregisterImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeAccountAttributesRequest(_a0 *ec2.DescribeAccountAttributesInput) (*request.Request, *ec2.DescribeAccountAttributesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAccountAttributesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAccountAttributesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAccountAttributesInput) *ec2.DescribeAccountAttributesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAccountAttributesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeAccountAttributes(_a0 *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeAccountAttributesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAccountAttributesInput) *ec2.DescribeAccountAttributesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAccountAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAccountAttributesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeAddressesRequest(_a0 *ec2.DescribeAddressesInput) (*request.Request, *ec2.DescribeAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAddressesInput) *ec2.DescribeAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAddressesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeAddresses(_a0 *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeAddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAddressesInput) *ec2.DescribeAddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeAvailabilityZonesRequest(_a0 *ec2.DescribeAvailabilityZonesInput) (*request.Request, *ec2.DescribeAvailabilityZonesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAvailabilityZonesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAvailabilityZonesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAvailabilityZonesInput) *ec2.DescribeAvailabilityZonesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAvailabilityZonesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeAvailabilityZones(_a0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeAvailabilityZonesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAvailabilityZonesInput) *ec2.DescribeAvailabilityZonesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAvailabilityZonesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAvailabilityZonesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeBundleTasksRequest(_a0 *ec2.DescribeBundleTasksInput) (*request.Request, *ec2.DescribeBundleTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeBundleTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeBundleTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeBundleTasksInput) *ec2.DescribeBundleTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeBundleTasksOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeBundleTasks(_a0 *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeBundleTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeBundleTasksInput) *ec2.DescribeBundleTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeBundleTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeBundleTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeClassicLinkInstancesRequest(_a0 *ec2.DescribeClassicLinkInstancesInput) (*request.Request, *ec2.DescribeClassicLinkInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeClassicLinkInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeClassicLinkInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeClassicLinkInstancesInput) *ec2.DescribeClassicLinkInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeClassicLinkInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeClassicLinkInstances(_a0 *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeClassicLinkInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeClassicLinkInstancesInput) *ec2.DescribeClassicLinkInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeClassicLinkInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeClassicLinkInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeConversionTasksRequest(_a0 *ec2.DescribeConversionTasksInput) (*request.Request, *ec2.DescribeConversionTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeConversionTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeConversionTasksInput) *ec2.DescribeConversionTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeConversionTasksOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeConversionTasks(_a0 *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeConversionTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) *ec2.DescribeConversionTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeConversionTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeConversionTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeCustomerGatewaysRequest(_a0 *ec2.DescribeCustomerGatewaysInput) (*request.Request, *ec2.DescribeCustomerGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeCustomerGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeCustomerGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeCustomerGatewaysInput) *ec2.DescribeCustomerGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeCustomerGatewaysOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeCustomerGateways(_a0 *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeCustomerGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeCustomerGatewaysInput) *ec2.DescribeCustomerGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeCustomerGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeCustomerGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeDhcpOptionsRequest(_a0 *ec2.DescribeDhcpOptionsInput) (*request.Request, *ec2.DescribeDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeDhcpOptionsInput) *ec2.DescribeDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeDhcpOptionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeDhcpOptions(_a0 *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeDhcpOptionsInput) *ec2.DescribeDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeExportTasksRequest(_a0 *ec2.DescribeExportTasksInput) (*request.Request, *ec2.DescribeExportTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeExportTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeExportTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeExportTasksInput) *ec2.DescribeExportTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeExportTasksOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeExportTasks(_a0 *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeExportTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeExportTasksInput) *ec2.DescribeExportTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeExportTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeExportTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeFlowLogsRequest(_a0 *ec2.DescribeFlowLogsInput) (*request.Request, *ec2.DescribeFlowLogsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFlowLogsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFlowLogsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFlowLogsInput) *ec2.DescribeFlowLogsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFlowLogsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeFlowLogs(_a0 *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFlowLogsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFlowLogsInput) *ec2.DescribeFlowLogsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFlowLogsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeHostsRequest(_a0 *ec2.DescribeHostsInput) (*request.Request, *ec2.DescribeHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostsInput) *ec2.DescribeHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeHostsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeHosts(_a0 *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostsInput) *ec2.DescribeHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeIdFormatRequest(_a0 *ec2.DescribeIdFormatInput) (*request.Request, *ec2.DescribeIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIdFormatInput) *ec2.DescribeIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeIdFormatOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeIdFormat(_a0 *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIdFormatInput) *ec2.DescribeIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImageAttributeRequest(_a0 *ec2.DescribeImageAttributeInput) (*request.Request, *ec2.DescribeImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImageAttributeInput) *ec2.DescribeImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImageAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImageAttribute(_a0 *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImageAttributeInput) *ec2.DescribeImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImagesRequest(_a0 *ec2.DescribeImagesInput) (*request.Request, *ec2.DescribeImagesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImagesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImagesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImagesInput) *ec2.DescribeImagesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImagesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImages(_a0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImagesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImagesInput) *ec2.DescribeImagesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImagesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImagesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImportImageTasksRequest(_a0 *ec2.DescribeImportImageTasksInput) (*request.Request, *ec2.DescribeImportImageTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportImageTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImportImageTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportImageTasksInput) *ec2.DescribeImportImageTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImportImageTasksOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImportImageTasks(_a0 *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImportImageTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportImageTasksInput) *ec2.DescribeImportImageTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImportImageTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportImageTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImportSnapshotTasksRequest(_a0 *ec2.DescribeImportSnapshotTasksInput) (*request.Request, *ec2.DescribeImportSnapshotTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportSnapshotTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImportSnapshotTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportSnapshotTasksInput) *ec2.DescribeImportSnapshotTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImportSnapshotTasksOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeImportSnapshotTasks(_a0 *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImportSnapshotTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportSnapshotTasksInput) *ec2.DescribeImportSnapshotTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImportSnapshotTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportSnapshotTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstanceAttributeRequest(_a0 *ec2.DescribeInstanceAttributeInput) (*request.Request, *ec2.DescribeInstanceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstanceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceAttributeInput) *ec2.DescribeInstanceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstanceAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstanceAttribute(_a0 *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceAttributeInput) *ec2.DescribeInstanceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstanceStatusRequest(_a0 *ec2.DescribeInstanceStatusInput) (*request.Request, *ec2.DescribeInstanceStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstanceStatusOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceStatusInput) *ec2.DescribeInstanceStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstanceStatusOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstanceStatus(_a0 *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInstanceStatusOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput) *ec2.DescribeInstanceStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstanceStatusPages(_a0 *ec2.DescribeInstanceStatusInput, _a1 func(*ec2.DescribeInstanceStatusOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput, func(*ec2.DescribeInstanceStatusOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeInstancesRequest(_a0 *ec2.DescribeInstancesInput) (*request.Request, *ec2.DescribeInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstancesInput) *ec2.DescribeInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstances(_a0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) *ec2.DescribeInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInstancesPages(_a0 *ec2.DescribeInstancesInput, _a1 func(*ec2.DescribeInstancesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput, func(*ec2.DescribeInstancesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeInternetGatewaysRequest(_a0 *ec2.DescribeInternetGatewaysInput) (*request.Request, *ec2.DescribeInternetGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInternetGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInternetGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInternetGatewaysInput) *ec2.DescribeInternetGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInternetGatewaysOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeInternetGateways(_a0 *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInternetGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInternetGatewaysInput) *ec2.DescribeInternetGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInternetGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInternetGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeKeyPairsRequest(_a0 *ec2.DescribeKeyPairsInput) (*request.Request, *ec2.DescribeKeyPairsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeKeyPairsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeKeyPairsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeKeyPairsInput) *ec2.DescribeKeyPairsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeKeyPairsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeKeyPairs(_a0 *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeKeyPairsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeKeyPairsInput) *ec2.DescribeKeyPairsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeKeyPairsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeKeyPairsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeMovingAddressesRequest(_a0 *ec2.DescribeMovingAddressesInput) (*request.Request, *ec2.DescribeMovingAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeMovingAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeMovingAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeMovingAddressesInput) *ec2.DescribeMovingAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeMovingAddressesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeMovingAddresses(_a0 *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeMovingAddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeMovingAddressesInput) *ec2.DescribeMovingAddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeMovingAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeMovingAddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNatGatewaysRequest(_a0 *ec2.DescribeNatGatewaysInput) (*request.Request, *ec2.DescribeNatGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNatGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNatGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNatGatewaysInput) *ec2.DescribeNatGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNatGatewaysOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNatGateways(_a0 *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNatGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNatGatewaysInput) *ec2.DescribeNatGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNatGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNatGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNetworkAclsRequest(_a0 *ec2.DescribeNetworkAclsInput) (*request.Request, *ec2.DescribeNetworkAclsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkAclsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkAclsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkAclsInput) *ec2.DescribeNetworkAclsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkAclsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNetworkAcls(_a0 *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNetworkAclsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkAclsInput) *ec2.DescribeNetworkAclsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkAclsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkAclsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNetworkInterfaceAttributeRequest(_a0 *ec2.DescribeNetworkInterfaceAttributeInput) (*request.Request, *ec2.DescribeNetworkInterfaceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfaceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfaceAttributeInput) *ec2.DescribeNetworkInterfaceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkInterfaceAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNetworkInterfaceAttribute(_a0 *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfaceAttributeInput) *ec2.DescribeNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNetworkInterfacesRequest(_a0 *ec2.DescribeNetworkInterfacesInput) (*request.Request, *ec2.DescribeNetworkInterfacesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfacesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkInterfacesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfacesInput) *ec2.DescribeNetworkInterfacesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkInterfacesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeNetworkInterfaces(_a0 *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNetworkInterfacesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfacesInput) *ec2.DescribeNetworkInterfacesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfacesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfacesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribePlacementGroupsRequest(_a0 *ec2.DescribePlacementGroupsInput) (*request.Request, *ec2.DescribePlacementGroupsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribePlacementGroupsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribePlacementGroupsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribePlacementGroupsInput) *ec2.DescribePlacementGroupsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribePlacementGroupsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribePlacementGroups(_a0 *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribePlacementGroupsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribePlacementGroupsInput) *ec2.DescribePlacementGroupsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePlacementGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribePlacementGroupsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribePrefixListsRequest(_a0 *ec2.DescribePrefixListsInput) (*request.Request, *ec2.DescribePrefixListsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribePrefixListsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribePrefixListsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribePrefixListsInput) *ec2.DescribePrefixListsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribePrefixListsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribePrefixLists(_a0 *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribePrefixListsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribePrefixListsInput) *ec2.DescribePrefixListsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePrefixListsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribePrefixListsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeRegionsRequest(_a0 *ec2.DescribeRegionsInput) (*request.Request, *ec2.DescribeRegionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRegionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeRegionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRegionsInput) *ec2.DescribeRegionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeRegionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeRegions(_a0 *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeRegionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRegionsInput) *ec2.DescribeRegionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeRegionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRegionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesRequest(_a0 *ec2.DescribeReservedInstancesInput) (*request.Request, *ec2.DescribeReservedInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesInput) *ec2.DescribeReservedInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstances(_a0 *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesInput) *ec2.DescribeReservedInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesListingsRequest(_a0 *ec2.DescribeReservedInstancesListingsInput) (*request.Request, *ec2.DescribeReservedInstancesListingsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesListingsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesListingsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesListingsInput) *ec2.DescribeReservedInstancesListingsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesListingsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesListings(_a0 *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesListingsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesListingsInput) *ec2.DescribeReservedInstancesListingsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesListingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesListingsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesModificationsRequest(_a0 *ec2.DescribeReservedInstancesModificationsInput) (*request.Request, *ec2.DescribeReservedInstancesModificationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesModificationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesModificationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesModificationsInput) *ec2.DescribeReservedInstancesModificationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesModificationsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesModifications(_a0 *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesModificationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesModificationsInput) *ec2.DescribeReservedInstancesModificationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesModificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesModificationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesModificationsPages(_a0 *ec2.DescribeReservedInstancesModificationsInput, _a1 func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesModificationsInput, func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeReservedInstancesOfferingsRequest(_a0 *ec2.DescribeReservedInstancesOfferingsInput) (*request.Request, *ec2.DescribeReservedInstancesOfferingsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesOfferingsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesOfferingsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesOfferingsInput) *ec2.DescribeReservedInstancesOfferingsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesOfferingsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesOfferings(_a0 *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesOfferingsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesOfferingsInput) *ec2.DescribeReservedInstancesOfferingsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesOfferingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesOfferingsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeReservedInstancesOfferingsPages(_a0 *ec2.DescribeReservedInstancesOfferingsInput, _a1 func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesOfferingsInput, func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeRouteTablesRequest(_a0 *ec2.DescribeRouteTablesInput) (*request.Request, *ec2.DescribeRouteTablesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRouteTablesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeRouteTablesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRouteTablesInput) *ec2.DescribeRouteTablesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeRouteTablesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeRouteTables(_a0 *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeRouteTablesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRouteTablesInput) *ec2.DescribeRouteTablesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeRouteTablesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRouteTablesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeScheduledInstanceAvailabilityRequest(_a0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*request.Request, *ec2.DescribeScheduledInstanceAvailabilityOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeScheduledInstanceAvailabilityOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) *ec2.DescribeScheduledInstanceAvailabilityOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeScheduledInstanceAvailabilityOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeScheduledInstanceAvailability(_a0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeScheduledInstanceAvailabilityOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) *ec2.DescribeScheduledInstanceAvailabilityOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeScheduledInstanceAvailabilityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeScheduledInstancesRequest(_a0 *ec2.DescribeScheduledInstancesInput) (*request.Request, *ec2.DescribeScheduledInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeScheduledInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstancesInput) *ec2.DescribeScheduledInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeScheduledInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeScheduledInstances(_a0 *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstancesInput) *ec2.DescribeScheduledInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSecurityGroupsRequest(_a0 *ec2.DescribeSecurityGroupsInput) (*request.Request, *ec2.DescribeSecurityGroupsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSecurityGroupsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSecurityGroupsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSecurityGroupsInput) *ec2.DescribeSecurityGroupsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSecurityGroupsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSecurityGroups(_a0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSecurityGroupsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSecurityGroupsInput) *ec2.DescribeSecurityGroupsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSecurityGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSecurityGroupsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSnapshotAttributeRequest(_a0 *ec2.DescribeSnapshotAttributeInput) (*request.Request, *ec2.DescribeSnapshotAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSnapshotAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotAttributeInput) *ec2.DescribeSnapshotAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSnapshotAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSnapshotAttribute(_a0 *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotAttributeInput) *ec2.DescribeSnapshotAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSnapshotsRequest(_a0 *ec2.DescribeSnapshotsInput) (*request.Request, *ec2.DescribeSnapshotsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSnapshotsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotsInput) *ec2.DescribeSnapshotsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSnapshotsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSnapshots(_a0 *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSnapshotsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput) *ec2.DescribeSnapshotsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSnapshotsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSnapshotsPages(_a0 *ec2.DescribeSnapshotsInput, _a1 func(*ec2.DescribeSnapshotsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput, func(*ec2.DescribeSnapshotsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeSpotDatafeedSubscriptionRequest(_a0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) *ec2.DescribeSpotDatafeedSubscriptionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotDatafeedSubscriptionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotDatafeedSubscription(_a0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) *ec2.DescribeSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotFleetInstancesRequest(_a0 *ec2.DescribeSpotFleetInstancesInput) (*request.Request, *ec2.DescribeSpotFleetInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotFleetInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetInstancesInput) *ec2.DescribeSpotFleetInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotFleetInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotFleetInstances(_a0 *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotFleetInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetInstancesInput) *ec2.DescribeSpotFleetInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotFleetRequestHistoryRequest(_a0 *ec2.DescribeSpotFleetRequestHistoryInput) (*request.Request, *ec2.DescribeSpotFleetRequestHistoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestHistoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotFleetRequestHistoryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestHistoryInput) *ec2.DescribeSpotFleetRequestHistoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotFleetRequestHistoryOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotFleetRequestHistory(_a0 *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotFleetRequestHistoryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestHistoryInput) *ec2.DescribeSpotFleetRequestHistoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetRequestHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestHistoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotFleetRequestsRequest(_a0 *ec2.DescribeSpotFleetRequestsInput) (*request.Request, *ec2.DescribeSpotFleetRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotFleetRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestsInput) *ec2.DescribeSpotFleetRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotFleetRequestsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotFleetRequests(_a0 *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotFleetRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestsInput) *ec2.DescribeSpotFleetRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotInstanceRequestsRequest(_a0 *ec2.DescribeSpotInstanceRequestsInput) (*request.Request, *ec2.DescribeSpotInstanceRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotInstanceRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotInstanceRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotInstanceRequestsInput) *ec2.DescribeSpotInstanceRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotInstanceRequestsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotInstanceRequests(_a0 *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotInstanceRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotInstanceRequestsInput) *ec2.DescribeSpotInstanceRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotInstanceRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotInstanceRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotPriceHistoryRequest(_a0 *ec2.DescribeSpotPriceHistoryInput) (*request.Request, *ec2.DescribeSpotPriceHistoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotPriceHistoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotPriceHistoryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotPriceHistoryInput) *ec2.DescribeSpotPriceHistoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotPriceHistoryOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotPriceHistory(_a0 *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotPriceHistoryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotPriceHistoryInput) *ec2.DescribeSpotPriceHistoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotPriceHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotPriceHistoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSpotPriceHistoryPages(_a0 *ec2.DescribeSpotPriceHistoryInput, _a1 func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotPriceHistoryInput, func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeSubnetsRequest(_a0 *ec2.DescribeSubnetsInput) (*request.Request, *ec2.DescribeSubnetsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSubnetsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSubnetsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSubnetsInput) *ec2.DescribeSubnetsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSubnetsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeSubnets(_a0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSubnetsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSubnetsInput) *ec2.DescribeSubnetsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSubnetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSubnetsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeTagsRequest(_a0 *ec2.DescribeTagsInput) (*request.Request, *ec2.DescribeTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeTagsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeTagsInput) *ec2.DescribeTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeTagsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeTags(_a0 *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeTagsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeTagsInput) *ec2.DescribeTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeTagsPages(_a0 *ec2.DescribeTagsInput, _a1 func(*ec2.DescribeTagsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeTagsInput, func(*ec2.DescribeTagsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeVolumeAttributeRequest(_a0 *ec2.DescribeVolumeAttributeInput) (*request.Request, *ec2.DescribeVolumeAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumeAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeAttributeInput) *ec2.DescribeVolumeAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumeAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVolumeAttribute(_a0 *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumeAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeAttributeInput) *ec2.DescribeVolumeAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumeAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVolumeStatusRequest(_a0 *ec2.DescribeVolumeStatusInput) (*request.Request, *ec2.DescribeVolumeStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumeStatusOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeStatusInput) *ec2.DescribeVolumeStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumeStatusOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVolumeStatus(_a0 *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumeStatusOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeStatusInput) *ec2.DescribeVolumeStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumeStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVolumeStatusPages(_a0 *ec2.DescribeVolumeStatusInput, _a1 func(*ec2.DescribeVolumeStatusOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeStatusInput, func(*ec2.DescribeVolumeStatusOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeVolumesRequest(_a0 *ec2.DescribeVolumesInput) (*request.Request, *ec2.DescribeVolumesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumesInput) *ec2.DescribeVolumesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVolumes(_a0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) *ec2.DescribeVolumesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVolumesPages(_a0 *ec2.DescribeVolumesInput, _a1 func(*ec2.DescribeVolumesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput, func(*ec2.DescribeVolumesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockEC2) DescribeVpcAttributeRequest(_a0 *ec2.DescribeVpcAttributeInput) (*request.Request, *ec2.DescribeVpcAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcAttributeInput) *ec2.DescribeVpcAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcAttribute(_a0 *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcAttributeInput) *ec2.DescribeVpcAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcClassicLinkRequest(_a0 *ec2.DescribeVpcClassicLinkInput) (*request.Request, *ec2.DescribeVpcClassicLinkOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcClassicLinkOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkInput) *ec2.DescribeVpcClassicLinkOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcClassicLinkOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcClassicLink(_a0 *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkInput) *ec2.DescribeVpcClassicLinkOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcClassicLinkDnsSupportRequest(_a0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DescribeVpcClassicLinkDnsSupportOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) *ec2.DescribeVpcClassicLinkDnsSupportOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcClassicLinkDnsSupportOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcClassicLinkDnsSupport(_a0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) *ec2.DescribeVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcEndpointServicesRequest(_a0 *ec2.DescribeVpcEndpointServicesInput) (*request.Request, *ec2.DescribeVpcEndpointServicesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServicesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointServicesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServicesInput) *ec2.DescribeVpcEndpointServicesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointServicesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcEndpointServices(_a0 *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointServicesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServicesInput) *ec2.DescribeVpcEndpointServicesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServicesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServicesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcEndpointsRequest(_a0 *ec2.DescribeVpcEndpointsInput) (*request.Request, *ec2.DescribeVpcEndpointsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointsInput) *ec2.DescribeVpcEndpointsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcEndpoints(_a0 *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointsInput) *ec2.DescribeVpcEndpointsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcPeeringConnectionsRequest(_a0 *ec2.DescribeVpcPeeringConnectionsInput) (*request.Request, *ec2.DescribeVpcPeeringConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcPeeringConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcPeeringConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcPeeringConnectionsInput) *ec2.DescribeVpcPeeringConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcPeeringConnectionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcPeeringConnections(_a0 *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcPeeringConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcPeeringConnectionsInput) *ec2.DescribeVpcPeeringConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcPeeringConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcPeeringConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcsRequest(_a0 *ec2.DescribeVpcsInput) (*request.Request, *ec2.DescribeVpcsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcsInput) *ec2.DescribeVpcsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpcs(_a0 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcsInput) *ec2.DescribeVpcsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpnConnectionsRequest(_a0 *ec2.DescribeVpnConnectionsInput) (*request.Request, *ec2.DescribeVpnConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpnConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnConnectionsInput) *ec2.DescribeVpnConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpnConnectionsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpnConnections(_a0 *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpnConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnConnectionsInput) *ec2.DescribeVpnConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpnConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpnGatewaysRequest(_a0 *ec2.DescribeVpnGatewaysInput) (*request.Request, *ec2.DescribeVpnGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpnGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnGatewaysInput) *ec2.DescribeVpnGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpnGatewaysOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DescribeVpnGateways(_a0 *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpnGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnGatewaysInput) *ec2.DescribeVpnGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpnGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DetachClassicLinkVpcRequest(_a0 *ec2.DetachClassicLinkVpcInput) (*request.Request, *ec2.DetachClassicLinkVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachClassicLinkVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachClassicLinkVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachClassicLinkVpcInput) *ec2.DetachClassicLinkVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachClassicLinkVpcOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DetachClassicLinkVpc(_a0 *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachClassicLinkVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachClassicLinkVpcInput) *ec2.DetachClassicLinkVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachClassicLinkVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachClassicLinkVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DetachInternetGatewayRequest(_a0 *ec2.DetachInternetGatewayInput) (*request.Request, *ec2.DetachInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachInternetGatewayInput) *ec2.DetachInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachInternetGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DetachInternetGateway(_a0 *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachInternetGatewayInput) *ec2.DetachInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DetachNetworkInterfaceRequest(_a0 *ec2.DetachNetworkInterfaceInput) (*request.Request, *ec2.DetachNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachNetworkInterfaceInput) *ec2.DetachNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachNetworkInterfaceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DetachNetworkInterface(_a0 *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachNetworkInterfaceInput) *ec2.DetachNetworkInterfaceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachNetworkInterfaceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DetachVolumeRequest(_a0 *ec2.DetachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.VolumeAttachment
	if rf, ok := ret.Get(1).(func(*ec2.DetachVolumeInput) *ec2.VolumeAttachment); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.VolumeAttachment)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DetachVolume(_a0 *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.VolumeAttachment
	if rf, ok := ret.Get(0).(func(*ec2.DetachVolumeInput) *ec2.VolumeAttachment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DetachVpnGatewayRequest(_a0 *ec2.DetachVpnGatewayInput) (*request.Request, *ec2.DetachVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachVpnGatewayInput) *ec2.DetachVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachVpnGatewayOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DetachVpnGateway(_a0 *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachVpnGatewayInput) *ec2.DetachVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DisableVgwRoutePropagationRequest(_a0 *ec2.DisableVgwRoutePropagationInput) (*request.Request, *ec2.DisableVgwRoutePropagationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisableVgwRoutePropagationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisableVgwRoutePropagationOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisableVgwRoutePropagationInput) *ec2.DisableVgwRoutePropagationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisableVgwRoutePropagationOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DisableVgwRoutePropagation(_a0 *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisableVgwRoutePropagationOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisableVgwRoutePropagationInput) *ec2.DisableVgwRoutePropagationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVgwRoutePropagationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisableVgwRoutePropagationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DisableVpcClassicLinkRequest(_a0 *ec2.DisableVpcClassicLinkInput) (*request.Request, *ec2.DisableVpcClassicLinkOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisableVpcClassicLinkOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkInput) *ec2.DisableVpcClassicLinkOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisableVpcClassicLinkOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DisableVpcClassicLink(_a0 *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisableVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkInput) *ec2.DisableVpcClassicLinkOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DisableVpcClassicLinkDnsSupportRequest(_a0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DisableVpcClassicLinkDnsSupportOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) *ec2.DisableVpcClassicLinkDnsSupportOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisableVpcClassicLinkDnsSupportOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DisableVpcClassicLinkDnsSupport(_a0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) *ec2.DisableVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DisassociateAddressRequest(_a0 *ec2.DisassociateAddressInput) (*request.Request, *ec2.DisassociateAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateAddressInput) *ec2.DisassociateAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateAddressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DisassociateAddress(_a0 *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisassociateAddressOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateAddressInput) *ec2.DisassociateAddressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateAddressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) DisassociateRouteTableRequest(_a0 *ec2.DisassociateRouteTableInput) (*request.Request, *ec2.DisassociateRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateRouteTableInput) *ec2.DisassociateRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateRouteTableOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) DisassociateRouteTable(_a0 *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisassociateRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateRouteTableInput) *ec2.DisassociateRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) EnableVgwRoutePropagationRequest(_a0 *ec2.EnableVgwRoutePropagationInput) (*request.Request, *ec2.EnableVgwRoutePropagationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVgwRoutePropagationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVgwRoutePropagationOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVgwRoutePropagationInput) *ec2.EnableVgwRoutePropagationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVgwRoutePropagationOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) EnableVgwRoutePropagation(_a0 *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVgwRoutePropagationOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVgwRoutePropagationInput) *ec2.EnableVgwRoutePropagationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVgwRoutePropagationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVgwRoutePropagationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) EnableVolumeIORequest(_a0 *ec2.EnableVolumeIOInput) (*request.Request, *ec2.EnableVolumeIOOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVolumeIOInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVolumeIOOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVolumeIOInput) *ec2.EnableVolumeIOOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVolumeIOOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) EnableVolumeIO(_a0 *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVolumeIOOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVolumeIOInput) *ec2.EnableVolumeIOOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVolumeIOOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVolumeIOInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) EnableVpcClassicLinkRequest(_a0 *ec2.EnableVpcClassicLinkInput) (*request.Request, *ec2.EnableVpcClassicLinkOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVpcClassicLinkOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkInput) *ec2.EnableVpcClassicLinkOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVpcClassicLinkOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) EnableVpcClassicLink(_a0 *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkInput) *ec2.EnableVpcClassicLinkOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) EnableVpcClassicLinkDnsSupportRequest(_a0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.EnableVpcClassicLinkDnsSupportOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) *ec2.EnableVpcClassicLinkDnsSupportOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVpcClassicLinkDnsSupportOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) EnableVpcClassicLinkDnsSupport(_a0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) *ec2.EnableVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) GetConsoleOutputRequest(_a0 *ec2.GetConsoleOutputInput) (*request.Request, *ec2.GetConsoleOutputOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetConsoleOutputInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetConsoleOutputOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetConsoleOutputInput) *ec2.GetConsoleOutputOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetConsoleOutputOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) GetConsoleOutput(_a0 *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetConsoleOutputOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetConsoleOutputInput) *ec2.GetConsoleOutputOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetConsoleOutputOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetConsoleOutputInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) GetPasswordDataRequest(_a0 *ec2.GetPasswordDataInput) (*request.Request, *ec2.GetPasswordDataOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetPasswordDataInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetPasswordDataOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetPasswordDataInput) *ec2.GetPasswordDataOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetPasswordDataOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) GetPasswordData(_a0 *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetPasswordDataOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetPasswordDataInput) *ec2.GetPasswordDataOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetPasswordDataOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetPasswordDataInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ImportImageRequest(_a0 *ec2.ImportImageInput) (*request.Request, *ec2.ImportImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportImageInput) *ec2.ImportImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportImageOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ImportImage(_a0 *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportImageInput) *ec2.ImportImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ImportInstanceRequest(_a0 *ec2.ImportInstanceInput) (*request.Request, *ec2.ImportInstanceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportInstanceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportInstanceOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportInstanceInput) *ec2.ImportInstanceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportInstanceOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ImportInstance(_a0 *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportInstanceOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportInstanceInput) *ec2.ImportInstanceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportInstanceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ImportKeyPairRequest(_a0 *ec2.ImportKeyPairInput) (*request.Request, *ec2.ImportKeyPairOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportKeyPairInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportKeyPairOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportKeyPairInput) *ec2.ImportKeyPairOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportKeyPairOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ImportKeyPair(_a0 *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportKeyPairOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportKeyPairInput) *ec2.ImportKeyPairOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportKeyPairInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ImportSnapshotRequest(_a0 *ec2.ImportSnapshotInput) (*request.Request, *ec2.ImportSnapshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportSnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportSnapshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportSnapshotInput) *ec2.ImportSnapshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportSnapshotOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ImportSnapshot(_a0 *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportSnapshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportSnapshotInput) *ec2.ImportSnapshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportSnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportSnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ImportVolumeRequest(_a0 *ec2.ImportVolumeInput) (*request.Request, *ec2.ImportVolumeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportVolumeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportVolumeInput) *ec2.ImportVolumeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportVolumeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ImportVolume(_a0 *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportVolumeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportVolumeInput) *ec2.ImportVolumeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyHostsRequest(_a0 *ec2.ModifyHostsInput) (*request.Request, *ec2.ModifyHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyHostsInput) *ec2.ModifyHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyHostsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyHosts(_a0 *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyHostsInput) *ec2.ModifyHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyIdFormatRequest(_a0 *ec2.ModifyIdFormatInput) (*request.Request, *ec2.ModifyIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyIdFormatInput) *ec2.ModifyIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyIdFormatOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyIdFormat(_a0 *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyIdFormatInput) *ec2.ModifyIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyImageAttributeRequest(_a0 *ec2.ModifyImageAttributeInput) (*request.Request, *ec2.ModifyImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyImageAttributeInput) *ec2.ModifyImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyImageAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyImageAttribute(_a0 *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyImageAttributeInput) *ec2.ModifyImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyInstanceAttributeRequest(_a0 *ec2.ModifyInstanceAttributeInput) (*request.Request, *ec2.ModifyInstanceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstanceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyInstanceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstanceAttributeInput) *ec2.ModifyInstanceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyInstanceAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyInstanceAttribute(_a0 *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstanceAttributeInput) *ec2.ModifyInstanceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstanceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyInstancePlacementRequest(_a0 *ec2.ModifyInstancePlacementInput) (*request.Request, *ec2.ModifyInstancePlacementOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstancePlacementInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyInstancePlacementOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstancePlacementInput) *ec2.ModifyInstancePlacementOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyInstancePlacementOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyInstancePlacement(_a0 *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyInstancePlacementOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstancePlacementInput) *ec2.ModifyInstancePlacementOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstancePlacementOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstancePlacementInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyNetworkInterfaceAttributeRequest(_a0 *ec2.ModifyNetworkInterfaceAttributeInput) (*request.Request, *ec2.ModifyNetworkInterfaceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyNetworkInterfaceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyNetworkInterfaceAttributeInput) *ec2.ModifyNetworkInterfaceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyNetworkInterfaceAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyNetworkInterfaceAttribute(_a0 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyNetworkInterfaceAttributeInput) *ec2.ModifyNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyReservedInstancesRequest(_a0 *ec2.ModifyReservedInstancesInput) (*request.Request, *ec2.ModifyReservedInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyReservedInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyReservedInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyReservedInstancesInput) *ec2.ModifyReservedInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyReservedInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyReservedInstances(_a0 *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyReservedInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyReservedInstancesInput) *ec2.ModifyReservedInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyReservedInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyReservedInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifySnapshotAttributeRequest(_a0 *ec2.ModifySnapshotAttributeInput) (*request.Request, *ec2.ModifySnapshotAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifySnapshotAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifySnapshotAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifySnapshotAttributeInput) *ec2.ModifySnapshotAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifySnapshotAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifySnapshotAttribute(_a0 *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifySnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifySnapshotAttributeInput) *ec2.ModifySnapshotAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifySnapshotAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifySpotFleetRequestRequest(_a0 *ec2.ModifySpotFleetRequestInput) (*request.Request, *ec2.ModifySpotFleetRequestOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifySpotFleetRequestInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifySpotFleetRequestOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifySpotFleetRequestInput) *ec2.ModifySpotFleetRequestOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifySpotFleetRequestOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifySpotFleetRequest(_a0 *ec2.ModifySpotFleetRequestInput) (*ec2.ModifySpotFleetRequestOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifySpotFleetRequestOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifySpotFleetRequestInput) *ec2.ModifySpotFleetRequestOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySpotFleetRequestOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifySpotFleetRequestInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifySubnetAttributeRequest(_a0 *ec2.ModifySubnetAttributeInput) (*request.Request, *ec2.ModifySubnetAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifySubnetAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifySubnetAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifySubnetAttributeInput) *ec2.ModifySubnetAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifySubnetAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifySubnetAttribute(_a0 *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifySubnetAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifySubnetAttributeInput) *ec2.ModifySubnetAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySubnetAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifySubnetAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyVolumeAttributeRequest(_a0 *ec2.ModifyVolumeAttributeInput) (*request.Request, *ec2.ModifyVolumeAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVolumeAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVolumeAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVolumeAttributeInput) *ec2.ModifyVolumeAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVolumeAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyVolumeAttribute(_a0 *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVolumeAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVolumeAttributeInput) *ec2.ModifyVolumeAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVolumeAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVolumeAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyVpcAttributeRequest(_a0 *ec2.ModifyVpcAttributeInput) (*request.Request, *ec2.ModifyVpcAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcAttributeInput) *ec2.ModifyVpcAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyVpcAttribute(_a0 *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcAttributeInput) *ec2.ModifyVpcAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ModifyVpcEndpointRequest(_a0 *ec2.ModifyVpcEndpointInput) (*request.Request, *ec2.ModifyVpcEndpointOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcEndpointOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointInput) *ec2.ModifyVpcEndpointOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcEndpointOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ModifyVpcEndpoint(_a0 *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcEndpointOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointInput) *ec2.ModifyVpcEndpointOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) MonitorInstancesRequest(_a0 *ec2.MonitorInstancesInput) (*request.Request, *ec2.MonitorInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.MonitorInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.MonitorInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.MonitorInstancesInput) *ec2.MonitorInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.MonitorInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) MonitorInstances(_a0 *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.MonitorInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.MonitorInstancesInput) *ec2.MonitorInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.MonitorInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.MonitorInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) MoveAddressToVpcRequest(_a0 *ec2.MoveAddressToVpcInput) (*request.Request, *ec2.MoveAddressToVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.MoveAddressToVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.MoveAddressToVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.MoveAddressToVpcInput) *ec2.MoveAddressToVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.MoveAddressToVpcOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) MoveAddressToVpc(_a0 *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.MoveAddressToVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.MoveAddressToVpcInput) *ec2.MoveAddressToVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.MoveAddressToVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.MoveAddressToVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) PurchaseReservedInstancesOfferingRequest(_a0 *ec2.PurchaseReservedInstancesOfferingInput) (*request.Request, *ec2.PurchaseReservedInstancesOfferingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseReservedInstancesOfferingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.PurchaseReservedInstancesOfferingOutput
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseReservedInstancesOfferingInput) *ec2.PurchaseReservedInstancesOfferingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.PurchaseReservedInstancesOfferingOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) PurchaseReservedInstancesOffering(_a0 *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.PurchaseReservedInstancesOfferingOutput
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseReservedInstancesOfferingInput) *ec2.PurchaseReservedInstancesOfferingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseReservedInstancesOfferingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseReservedInstancesOfferingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) PurchaseScheduledInstancesRequest(_a0 *ec2.PurchaseScheduledInstancesInput) (*request.Request, *ec2.PurchaseScheduledInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseScheduledInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.PurchaseScheduledInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseScheduledInstancesInput) *ec2.PurchaseScheduledInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.PurchaseScheduledInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) PurchaseScheduledInstances(_a0 *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.PurchaseScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseScheduledInstancesInput) *ec2.PurchaseScheduledInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseScheduledInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RebootInstancesRequest(_a0 *ec2.RebootInstancesInput) (*request.Request, *ec2.RebootInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RebootInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RebootInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.RebootInstancesInput) *ec2.RebootInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RebootInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RebootInstances(_a0 *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RebootInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.RebootInstancesInput) *ec2.RebootInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RebootInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RebootInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RegisterImageRequest(_a0 *ec2.RegisterImageInput) (*request.Request, *ec2.RegisterImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RegisterImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RegisterImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.RegisterImageInput) *ec2.RegisterImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RegisterImageOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RegisterImage(_a0 *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RegisterImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.RegisterImageInput) *ec2.RegisterImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RegisterImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RegisterImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RejectVpcPeeringConnectionRequest(_a0 *ec2.RejectVpcPeeringConnectionInput) (*request.Request, *ec2.RejectVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RejectVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RejectVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.RejectVpcPeeringConnectionInput) *ec2.RejectVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RejectVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RejectVpcPeeringConnection(_a0 *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RejectVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.RejectVpcPeeringConnectionInput) *ec2.RejectVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RejectVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RejectVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReleaseAddressRequest(_a0 *ec2.ReleaseAddressInput) (*request.Request, *ec2.ReleaseAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReleaseAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReleaseAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReleaseAddressInput) *ec2.ReleaseAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReleaseAddressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReleaseAddress(_a0 *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReleaseAddressOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReleaseAddressInput) *ec2.ReleaseAddressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReleaseAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReleaseAddressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReleaseHostsRequest(_a0 *ec2.ReleaseHostsInput) (*request.Request, *ec2.ReleaseHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReleaseHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReleaseHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReleaseHostsInput) *ec2.ReleaseHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReleaseHostsOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReleaseHosts(_a0 *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReleaseHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReleaseHostsInput) *ec2.ReleaseHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReleaseHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReleaseHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceNetworkAclAssociationRequest(_a0 *ec2.ReplaceNetworkAclAssociationInput) (*request.Request, *ec2.ReplaceNetworkAclAssociationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclAssociationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceNetworkAclAssociationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclAssociationInput) *ec2.ReplaceNetworkAclAssociationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceNetworkAclAssociationOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceNetworkAclAssociation(_a0 *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceNetworkAclAssociationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclAssociationInput) *ec2.ReplaceNetworkAclAssociationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceNetworkAclAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclAssociationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceNetworkAclEntryRequest(_a0 *ec2.ReplaceNetworkAclEntryInput) (*request.Request, *ec2.ReplaceNetworkAclEntryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclEntryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceNetworkAclEntryOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclEntryInput) *ec2.ReplaceNetworkAclEntryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceNetworkAclEntryOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceNetworkAclEntry(_a0 *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclEntryInput) *ec2.ReplaceNetworkAclEntryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclEntryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceRouteRequest(_a0 *ec2.ReplaceRouteInput) (*request.Request, *ec2.ReplaceRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteInput) *ec2.ReplaceRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceRouteOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceRoute(_a0 *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteInput) *ec2.ReplaceRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceRouteTableAssociationRequest(_a0 *ec2.ReplaceRouteTableAssociationInput) (*request.Request, *ec2.ReplaceRouteTableAssociationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteTableAssociationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceRouteTableAssociationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteTableAssociationInput) *ec2.ReplaceRouteTableAssociationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceRouteTableAssociationOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReplaceRouteTableAssociation(_a0 *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceRouteTableAssociationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteTableAssociationInput) *ec2.ReplaceRouteTableAssociationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceRouteTableAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteTableAssociationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ReportInstanceStatusRequest(_a0 *ec2.ReportInstanceStatusInput) (*request.Request, *ec2.ReportInstanceStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReportInstanceStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReportInstanceStatusOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReportInstanceStatusInput) *ec2.ReportInstanceStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReportInstanceStatusOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ReportInstanceStatus(_a0 *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReportInstanceStatusOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReportInstanceStatusInput) *ec2.ReportInstanceStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReportInstanceStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReportInstanceStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RequestSpotFleetRequest(_a0 *ec2.RequestSpotFleetInput) (*request.Request, *ec2.RequestSpotFleetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotFleetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RequestSpotFleetOutput
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotFleetInput) *ec2.RequestSpotFleetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RequestSpotFleetOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RequestSpotFleet(_a0 *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RequestSpotFleetOutput
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotFleetInput) *ec2.RequestSpotFleetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RequestSpotFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotFleetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RequestSpotInstancesRequest(_a0 *ec2.RequestSpotInstancesInput) (*request.Request, *ec2.RequestSpotInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RequestSpotInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotInstancesInput) *ec2.RequestSpotInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RequestSpotInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RequestSpotInstances(_a0 *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RequestSpotInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotInstancesInput) *ec2.RequestSpotInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RequestSpotInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ResetImageAttributeRequest(_a0 *ec2.ResetImageAttributeInput) (*request.Request, *ec2.ResetImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetImageAttributeInput) *ec2.ResetImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetImageAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ResetImageAttribute(_a0 *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetImageAttributeInput) *ec2.ResetImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ResetInstanceAttributeRequest(_a0 *ec2.ResetInstanceAttributeInput) (*request.Request, *ec2.ResetInstanceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetInstanceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetInstanceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetInstanceAttributeInput) *ec2.ResetInstanceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetInstanceAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ResetInstanceAttribute(_a0 *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetInstanceAttributeInput) *ec2.ResetInstanceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetInstanceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ResetNetworkInterfaceAttributeRequest(_a0 *ec2.ResetNetworkInterfaceAttributeInput) (*request.Request, *ec2.ResetNetworkInterfaceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetNetworkInterfaceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetNetworkInterfaceAttributeInput) *ec2.ResetNetworkInterfaceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetNetworkInterfaceAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ResetNetworkInterfaceAttribute(_a0 *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetNetworkInterfaceAttributeInput) *ec2.ResetNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) ResetSnapshotAttributeRequest(_a0 *ec2.ResetSnapshotAttributeInput) (*request.Request, *ec2.ResetSnapshotAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetSnapshotAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetSnapshotAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetSnapshotAttributeInput) *ec2.ResetSnapshotAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetSnapshotAttributeOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) ResetSnapshotAttribute(_a0 *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetSnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetSnapshotAttributeInput) *ec2.ResetSnapshotAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetSnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetSnapshotAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RestoreAddressToClassicRequest(_a0 *ec2.RestoreAddressToClassicInput) (*request.Request, *ec2.RestoreAddressToClassicOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RestoreAddressToClassicInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RestoreAddressToClassicOutput
	if rf, ok := ret.Get(1).(func(*ec2.RestoreAddressToClassicInput) *ec2.RestoreAddressToClassicOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RestoreAddressToClassicOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RestoreAddressToClassic(_a0 *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RestoreAddressToClassicOutput
	if rf, ok := ret.Get(0).(func(*ec2.RestoreAddressToClassicInput) *ec2.RestoreAddressToClassicOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RestoreAddressToClassicOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RestoreAddressToClassicInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RevokeSecurityGroupEgressRequest(_a0 *ec2.RevokeSecurityGroupEgressInput) (*request.Request, *ec2.RevokeSecurityGroupEgressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupEgressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RevokeSecurityGroupEgressOutput
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupEgressInput) *ec2.RevokeSecurityGroupEgressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RevokeSecurityGroupEgressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RevokeSecurityGroupEgress(_a0 *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RevokeSecurityGroupEgressOutput
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupEgressInput) *ec2.RevokeSecurityGroupEgressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RevokeSecurityGroupEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupEgressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RevokeSecurityGroupIngressRequest(_a0 *ec2.RevokeSecurityGroupIngressInput) (*request.Request, *ec2.RevokeSecurityGroupIngressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupIngressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RevokeSecurityGroupIngressOutput
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupIngressInput) *ec2.RevokeSecurityGroupIngressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RevokeSecurityGroupIngressOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RevokeSecurityGroupIngress(_a0 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RevokeSecurityGroupIngressOutput
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupIngressInput) *ec2.RevokeSecurityGroupIngressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RevokeSecurityGroupIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupIngressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RunInstancesRequest(_a0 *ec2.RunInstancesInput) (*request.Request, *ec2.Reservation) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RunInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.Reservation
	if rf, ok := ret.Get(1).(func(*ec2.RunInstancesInput) *ec2.Reservation); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.Reservation)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RunInstances(_a0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Reservation
	if rf, ok := ret.Get(0).(func(*ec2.RunInstancesInput) *ec2.Reservation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RunInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) RunScheduledInstancesRequest(_a0 *ec2.RunScheduledInstancesInput) (*request.Request, *ec2.RunScheduledInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RunScheduledInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RunScheduledInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.RunScheduledInstancesInput) *ec2.RunScheduledInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RunScheduledInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) RunScheduledInstances(_a0 *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RunScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.RunScheduledInstancesInput) *ec2.RunScheduledInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RunScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RunScheduledInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) StartInstancesRequest(_a0 *ec2.StartInstancesInput) (*request.Request, *ec2.StartInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.StartInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.StartInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.StartInstancesInput) *ec2.StartInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.StartInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) StartInstances(_a0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.StartInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.StartInstancesInput) *ec2.StartInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.StartInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.StartInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) StopInstancesRequest(_a0 *ec2.StopInstancesInput) (*request.Request, *ec2.StopInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.StopInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.StopInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.StopInstancesInput) *ec2.StopInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.StopInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) StopInstances(_a0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.StopInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.StopInstancesInput) *ec2.StopInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.StopInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.StopInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) TerminateInstancesRequest(_a0 *ec2.TerminateInstancesInput) (*request.Request, *ec2.TerminateInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.TerminateInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.TerminateInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.TerminateInstancesInput) *ec2.TerminateInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.TerminateInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) TerminateInstances(_a0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.TerminateInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.TerminateInstancesInput) *ec2.TerminateInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.TerminateInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.TerminateInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) UnassignPrivateIpAddressesRequest(_a0 *ec2.UnassignPrivateIpAddressesInput) (*request.Request, *ec2.UnassignPrivateIpAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UnassignPrivateIpAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UnassignPrivateIpAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.UnassignPrivateIpAddressesInput) *ec2.UnassignPrivateIpAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UnassignPrivateIpAddressesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) UnassignPrivateIpAddresses(_a0 *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.UnassignPrivateIpAddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.UnassignPrivateIpAddressesInput) *ec2.UnassignPrivateIpAddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnassignPrivateIpAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.UnassignPrivateIpAddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockEC2) UnmonitorInstancesRequest(_a0 *ec2.UnmonitorInstancesInput) (*request.Request, *ec2.UnmonitorInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UnmonitorInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UnmonitorInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.UnmonitorInstancesInput) *ec2.UnmonitorInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UnmonitorInstancesOutput)
		}
	}

	return r0, r1
}
func (_m *MockEC2) UnmonitorInstances(_a0 *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.UnmonitorInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.UnmonitorInstancesInput) *ec2.UnmonitorInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnmonitorInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.UnmonitorInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
