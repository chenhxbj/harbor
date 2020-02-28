// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/goharbor/harbor/src/pkg/p2p/preheat/models"
	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AppendHistory provides a mock function with given fields: record
func (_m *Manager) AppendHistory(record *models.HistoryRecord) error {
	ret := _m.Called(record)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.HistoryRecord) error); ok {
		r0 = rf(record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadHistories provides a mock function with given fields: params
func (_m *Manager) LoadHistories(params *models.QueryParam) (int64, []*models.HistoryRecord, error) {
	ret := _m.Called(params)

	var r0 []*models.HistoryRecord
	if rf, ok := ret.Get(1).(func(*models.QueryParam) []*models.HistoryRecord); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(1).([]*models.HistoryRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(*models.QueryParam) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(2)
	}

	return int64(ret.Int(0)), r0, r1
}

// UpdateStatus provides a mock function with given fields: taskID, status, startTime, endTime
func (_m *Manager) UpdateStatus(taskID string, status models.TrackStatus, startTime string, endTime string) error {
	ret := _m.Called(taskID, status, startTime, endTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.TrackStatus, string, string) error); ok {
		r0 = rf(taskID, status, startTime, endTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}