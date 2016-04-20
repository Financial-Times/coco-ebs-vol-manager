package main

import (
	"github.com/Financial-Times/coco-ebs-vol-manager/mocks"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vektra/errors"
	"testing"
)

func TestAttachVolShouldExitOnFatal(t *testing.T) {
	assert := assert.New(t)
	*device = "/some/path"
	*instanceId = "insantce-id"
	*volId = "vol-id"

	m := new(mocks.MockEC2)
	expectedErr := errors.New("Boom")
	m.On("AttachVolume", mock.AnythingOfType("*ec2.AttachVolumeInput")).Return(nil, expectedErr)

	err := attachVol(&Ec2Client{m})

	m.AssertExpectations(t)

	assert.Equal(expectedErr, err)
}

func TestAttachVol(t *testing.T) {
	assert := assert.New(t)
	*device = "/some/path"
	*instanceId = "insantce-id"
	*volId = "vol-id"
	m := new(mocks.MockEC2)

	m.On("AttachVolume", mock.AnythingOfType("*ec2.AttachVolumeInput")).Return(nil, nil)

	attachVol(&Ec2Client{m})

	m.AssertExpectations(t)
	avi := m.Mock.Calls[0].Arguments.Get(0).(*ec2.AttachVolumeInput)
	assert.NotEmpty(avi)
	assert.Equal("/some/path", *avi.Device)
	assert.Equal("insantce-id", *avi.InstanceId)
	assert.Equal("vol-id", *avi.VolumeId)
	assert.Equal(false, *avi.DryRun)
}

func TestDetachVolShouldExitOnFatal(t *testing.T) {
	assert := assert.New(t)
	*device = "/some/path"
	*instanceId = "insantce-id"
	*volId = "vol-id"

	m := new(mocks.MockEC2)
	expectedErr := errors.New("Boom")
	m.On("DetachVolume", mock.AnythingOfType("*ec2.DetachVolumeInput")).Return(nil, expectedErr)

	err := detachVol(&Ec2Client{m})

	m.AssertExpectations(t)

	assert.Equal(expectedErr, err)
}

func TestDetachVol(t *testing.T) {
	assert := assert.New(t)
	*device = "/some/path"
	*instanceId = "insantce-id"
	*volId = "vol-id"
	m := new(mocks.MockEC2)

	m.On("DetachVolume", mock.AnythingOfType("*ec2.DetachVolumeInput")).Return(nil, nil)

	detachVol(&Ec2Client{m})

	m.AssertExpectations(t)
	dvi := m.Mock.Calls[0].Arguments.Get(0).(*ec2.DetachVolumeInput)
	assert.NotEmpty(dvi)
	assert.Equal("/some/path", *dvi.Device)
	assert.Equal("insantce-id", *dvi.InstanceId)
	assert.Equal("vol-id", *dvi.VolumeId)
	assert.Equal(false, *dvi.DryRun)
}
func TestCreateSnapshotShouldReturnErrorOnFail(t *testing.T) {
	assert := assert.New(t)
	*device = "/some/path"
	*instanceId = "insantce-id"
	*volId = "vol-id"
	*description = "snapshot description"

	m := new(mocks.MockEC2)
	expectedErr := errors.New("Boom")
	m.On("CreateSnapshot", mock.AnythingOfType("*ec2.CreateSnapshotInput")).Return(nil, expectedErr)

	err := createSnapshot(&Ec2Client{m})

	m.AssertExpectations(t)

	assert.Equal(expectedErr, err)
}

func TestCreateSnapshot(t *testing.T) {
	assert := assert.New(t)
	*device = "/some/path"
	*instanceId = "insantce-id"
	*volId = "vol-id"
	*description = "snapshot description"
	m := new(mocks.MockEC2)

	m.On("CreateSnapshot", mock.AnythingOfType("*ec2.CreateSnapshotInput")).Return(nil, nil)

	createSnapshot(&Ec2Client{m})

	m.AssertExpectations(t)
	csi := m.Mock.Calls[0].Arguments.Get(0).(*ec2.CreateSnapshotInput)
	assert.NotEmpty(csi)
	assert.Equal("snapshot description", *csi.Description)
	assert.Equal("vol-id", *csi.VolumeId)
	assert.Equal(false, *csi.DryRun)
}
