// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import codec "github.com/eko/gocache/codec"

import mock "github.com/stretchr/testify/mock"

// MetricsInterface is an autogenerated mock type for the MetricsInterface type
type MetricsInterface struct {
	mock.Mock
}

// Record provides a mock function with given fields: store, metric, value
func (_m *MetricsInterface) Record(store string, metric string, value float64) {
	_m.Called(store, metric, value)
}

// RecordFromCodec provides a mock function with given fields: _a0
func (_m *MetricsInterface) RecordFromCodec(_a0 codec.CodecInterface) {
	_m.Called(_a0)
}
