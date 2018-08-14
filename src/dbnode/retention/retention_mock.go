// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/retention/types.go

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package retention is a generated GoMock package.
package retention

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockOptions is a mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsMockRecorder
}

// MockOptionsMockRecorder is the mock recorder for MockOptions
type MockOptionsMockRecorder struct {
	mock *MockOptions
}

// NewMockOptions creates a new mock instance
func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &MockOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockOptions) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// Equal mocks base method
func (m *MockOptions) Equal(value Options) bool {
	ret := m.ctrl.Call(m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockOptionsMockRecorder) Equal(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockOptions)(nil).Equal), value)
}

// SetRetentionPeriod mocks base method
func (m *MockOptions) SetRetentionPeriod(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetRetentionPeriod", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRetentionPeriod indicates an expected call of SetRetentionPeriod
func (mr *MockOptionsMockRecorder) SetRetentionPeriod(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRetentionPeriod", reflect.TypeOf((*MockOptions)(nil).SetRetentionPeriod), value)
}

// RetentionPeriod mocks base method
func (m *MockOptions) RetentionPeriod() time.Duration {
	ret := m.ctrl.Call(m, "RetentionPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetentionPeriod indicates an expected call of RetentionPeriod
func (mr *MockOptionsMockRecorder) RetentionPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetentionPeriod", reflect.TypeOf((*MockOptions)(nil).RetentionPeriod))
}

// SetBlockSize mocks base method
func (m *MockOptions) SetBlockSize(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBlockSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockSize indicates an expected call of SetBlockSize
func (mr *MockOptionsMockRecorder) SetBlockSize(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockSize", reflect.TypeOf((*MockOptions)(nil).SetBlockSize), value)
}

// BlockSize mocks base method
func (m *MockOptions) BlockSize() time.Duration {
	ret := m.ctrl.Call(m, "BlockSize")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockSize indicates an expected call of BlockSize
func (mr *MockOptionsMockRecorder) BlockSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSize", reflect.TypeOf((*MockOptions)(nil).BlockSize))
}

// SetBufferFuture mocks base method
func (m *MockOptions) SetBufferFuture(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBufferFuture", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBufferFuture indicates an expected call of SetBufferFuture
func (mr *MockOptionsMockRecorder) SetBufferFuture(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBufferFuture", reflect.TypeOf((*MockOptions)(nil).SetBufferFuture), value)
}

// BufferFuture mocks base method
func (m *MockOptions) BufferFuture() time.Duration {
	ret := m.ctrl.Call(m, "BufferFuture")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BufferFuture indicates an expected call of BufferFuture
func (mr *MockOptionsMockRecorder) BufferFuture() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferFuture", reflect.TypeOf((*MockOptions)(nil).BufferFuture))
}

// SetBufferPast mocks base method
func (m *MockOptions) SetBufferPast(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBufferPast", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBufferPast indicates an expected call of SetBufferPast
func (mr *MockOptionsMockRecorder) SetBufferPast(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBufferPast", reflect.TypeOf((*MockOptions)(nil).SetBufferPast), value)
}

// BufferPast mocks base method
func (m *MockOptions) BufferPast() time.Duration {
	ret := m.ctrl.Call(m, "BufferPast")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BufferPast indicates an expected call of BufferPast
func (mr *MockOptionsMockRecorder) BufferPast() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferPast", reflect.TypeOf((*MockOptions)(nil).BufferPast))
}

// SetBlockDataExpiry mocks base method
func (m *MockOptions) SetBlockDataExpiry(on bool) Options {
	ret := m.ctrl.Call(m, "SetBlockDataExpiry", on)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockDataExpiry indicates an expected call of SetBlockDataExpiry
func (mr *MockOptionsMockRecorder) SetBlockDataExpiry(on interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockDataExpiry", reflect.TypeOf((*MockOptions)(nil).SetBlockDataExpiry), on)
}

// BlockDataExpiry mocks base method
func (m *MockOptions) BlockDataExpiry() bool {
	ret := m.ctrl.Call(m, "BlockDataExpiry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockDataExpiry indicates an expected call of BlockDataExpiry
func (mr *MockOptionsMockRecorder) BlockDataExpiry() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockDataExpiry", reflect.TypeOf((*MockOptions)(nil).BlockDataExpiry))
}

// SetBlockDataExpiryAfterNotAccessedPeriod mocks base method
func (m *MockOptions) SetBlockDataExpiryAfterNotAccessedPeriod(period time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBlockDataExpiryAfterNotAccessedPeriod", period)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockDataExpiryAfterNotAccessedPeriod indicates an expected call of SetBlockDataExpiryAfterNotAccessedPeriod
func (mr *MockOptionsMockRecorder) SetBlockDataExpiryAfterNotAccessedPeriod(period interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockDataExpiryAfterNotAccessedPeriod", reflect.TypeOf((*MockOptions)(nil).SetBlockDataExpiryAfterNotAccessedPeriod), period)
}

// BlockDataExpiryAfterNotAccessedPeriod mocks base method
func (m *MockOptions) BlockDataExpiryAfterNotAccessedPeriod() time.Duration {
	ret := m.ctrl.Call(m, "BlockDataExpiryAfterNotAccessedPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockDataExpiryAfterNotAccessedPeriod indicates an expected call of BlockDataExpiryAfterNotAccessedPeriod
func (mr *MockOptionsMockRecorder) BlockDataExpiryAfterNotAccessedPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockDataExpiryAfterNotAccessedPeriod", reflect.TypeOf((*MockOptions)(nil).BlockDataExpiryAfterNotAccessedPeriod))
}

// SetAnyWriteTimeEnabled mocks base method
func (m *MockOptions) SetAnyWriteTimeEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetAnyWriteTimeEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetAnyWriteTimeEnabled indicates an expected call of SetAnyWriteTimeEnabled
func (mr *MockOptionsMockRecorder) SetAnyWriteTimeEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnyWriteTimeEnabled", reflect.TypeOf((*MockOptions)(nil).SetAnyWriteTimeEnabled), value)
}

// AnyWriteTimeEnabled mocks base method
func (m *MockOptions) AnyWriteTimeEnabled() bool {
	ret := m.ctrl.Call(m, "AnyWriteTimeEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AnyWriteTimeEnabled indicates an expected call of AnyWriteTimeEnabled
func (mr *MockOptionsMockRecorder) AnyWriteTimeEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnyWriteTimeEnabled", reflect.TypeOf((*MockOptions)(nil).AnyWriteTimeEnabled))
}

// SetFlushAfterNoMetricPeriod mocks base method
func (m *MockOptions) SetFlushAfterNoMetricPeriod(period time.Duration) Options {
	ret := m.ctrl.Call(m, "SetFlushAfterNoMetricPeriod", period)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetFlushAfterNoMetricPeriod indicates an expected call of SetFlushAfterNoMetricPeriod
func (mr *MockOptionsMockRecorder) SetFlushAfterNoMetricPeriod(period interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFlushAfterNoMetricPeriod", reflect.TypeOf((*MockOptions)(nil).SetFlushAfterNoMetricPeriod), period)
}

// FlushAfterNoMetricPeriod mocks base method
func (m *MockOptions) FlushAfterNoMetricPeriod() time.Duration {
	ret := m.ctrl.Call(m, "FlushAfterNoMetricPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// FlushAfterNoMetricPeriod indicates an expected call of FlushAfterNoMetricPeriod
func (mr *MockOptionsMockRecorder) FlushAfterNoMetricPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAfterNoMetricPeriod", reflect.TypeOf((*MockOptions)(nil).FlushAfterNoMetricPeriod))
}

// SetMaxWritesBeforeFlush mocks base method
func (m *MockOptions) SetMaxWritesBeforeFlush(value uint64) Options {
	ret := m.ctrl.Call(m, "SetMaxWritesBeforeFlush", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetMaxWritesBeforeFlush indicates an expected call of SetMaxWritesBeforeFlush
func (mr *MockOptionsMockRecorder) SetMaxWritesBeforeFlush(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxWritesBeforeFlush", reflect.TypeOf((*MockOptions)(nil).SetMaxWritesBeforeFlush), value)
}

// MaxWritesBeforeFlush mocks base method
func (m *MockOptions) MaxWritesBeforeFlush() uint64 {
	ret := m.ctrl.Call(m, "MaxWritesBeforeFlush")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// MaxWritesBeforeFlush indicates an expected call of MaxWritesBeforeFlush
func (mr *MockOptionsMockRecorder) MaxWritesBeforeFlush() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxWritesBeforeFlush", reflect.TypeOf((*MockOptions)(nil).MaxWritesBeforeFlush))
}
