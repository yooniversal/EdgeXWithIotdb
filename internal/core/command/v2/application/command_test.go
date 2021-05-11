//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package application

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"

	"github.com/stretchr/testify/assert"
)

const (
	testDeviceName = "testDeviceName"
	testServiceUrl = "http://localhost:48082"
	resource1      = "str1_R"
	resource2      = "str2_W"
	resource3      = "str3_RW"
	resource4      = "str4_RW"
	resource5      = "str5_RW"
	resource6      = "str6_RW"
	command1       = "cmd1"
	command2       = "cmd2"
)

func TestBuildCoreCommands(t *testing.T) {
	profile := dtos.DeviceProfile{
		Name: "testProfile",
		DeviceResources: []dtos.DeviceResource{
			{Name: resource1, Properties: dtos.ResourceProperties{ValueType: v2.ValueTypeString, ReadWrite: v2.ReadWrite_R}},
			{Name: resource2, Properties: dtos.ResourceProperties{ValueType: v2.ValueTypeInt16, ReadWrite: v2.ReadWrite_W}},
			{Name: resource3, Properties: dtos.ResourceProperties{ValueType: v2.ValueTypeBool, ReadWrite: v2.ReadWrite_RW}},
			{Name: resource4, Properties: dtos.ResourceProperties{ValueType: v2.ValueTypeString, ReadWrite: v2.ReadWrite_RW}, IsHidden: true},
			{Name: resource5, Properties: dtos.ResourceProperties{ValueType: v2.ValueTypeInt16, ReadWrite: v2.ReadWrite_RW}},
			{Name: resource6, Properties: dtos.ResourceProperties{ValueType: v2.ValueTypeBool, ReadWrite: v2.ReadWrite_RW}},
		},
		DeviceCommands: []dtos.DeviceCommand{
			{
				Name: command1, ReadWrite: v2.ReadWrite_R,
				ResourceOperations: []dtos.ResourceOperation{
					{DeviceResource: resource1}, {DeviceResource: resource2}, {DeviceResource: resource3},
				},
			},
			{
				Name: command2, ReadWrite: v2.ReadWrite_W, IsHidden: true,
				ResourceOperations: []dtos.ResourceOperation{
					{DeviceResource: resource4}, {DeviceResource: resource5},
				},
			},
			{
				Name: resource6, ReadWrite: v2.ReadWrite_RW,
				ResourceOperations: []dtos.ResourceOperation{
					{DeviceResource: resource6},
				},
			},
		},
	}
	expectedCoreCommand := []dtos.CoreCommand{
		{
			Name: command1, Get: true, Path: commandPath(testDeviceName, command1), Url: testServiceUrl,
			Parameters: []dtos.CoreCommandParameter{
				{ResourceName: resource1, ValueType: v2.ValueTypeString},
				{ResourceName: resource2, ValueType: v2.ValueTypeInt16},
				{ResourceName: resource3, ValueType: v2.ValueTypeBool},
			},
		},
		{
			Name: resource6, Get: true, Set: true, Path: commandPath(testDeviceName, resource6), Url: testServiceUrl,
			Parameters: []dtos.CoreCommandParameter{{ResourceName: resource6, ValueType: v2.ValueTypeBool}},
		},
		{
			Name: resource1, Get: true, Path: commandPath(testDeviceName, resource1), Url: testServiceUrl,
			Parameters: []dtos.CoreCommandParameter{{ResourceName: resource1, ValueType: v2.ValueTypeString}},
		},
		{
			Name: resource2, Set: true, Path: commandPath(testDeviceName, resource2), Url: testServiceUrl,
			Parameters: []dtos.CoreCommandParameter{{ResourceName: resource2, ValueType: v2.ValueTypeInt16}},
		},
		{
			Name: resource3, Get: true, Set: true, Path: commandPath(testDeviceName, resource3), Url: testServiceUrl,
			Parameters: []dtos.CoreCommandParameter{{ResourceName: resource3, ValueType: v2.ValueTypeBool}},
		},
		{
			Name: resource5, Get: true, Set: true, Path: commandPath(testDeviceName, resource5), Url: testServiceUrl,
			Parameters: []dtos.CoreCommandParameter{{ResourceName: resource5, ValueType: v2.ValueTypeInt16}},
		},
	}

	result, err := buildCoreCommands(testDeviceName, testServiceUrl, profile)
	require.NoError(t, err)

	assert.ElementsMatch(t, expectedCoreCommand, result)
}
