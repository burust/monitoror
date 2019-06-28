// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/monitoror/monitoror/monitorable/jenkins/models"
import tiles "github.com/monitoror/monitoror/models/tiles"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Build provides a mock function with given fields: params
func (_m *Usecase) Build(params *models.JobParams) (*tiles.BuildTile, error) {
	ret := _m.Called(params)

	var r0 *tiles.BuildTile
	if rf, ok := ret.Get(0).(func(*models.JobParams) *tiles.BuildTile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tiles.BuildTile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.JobParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
