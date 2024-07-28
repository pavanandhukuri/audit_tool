// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	entities "security_audit_tool/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// ReportGenerator is an autogenerated mock type for the ReportGenerator type
type ReportGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields: result
func (_m *ReportGenerator) Generate(result *entities.ValidationResult) error {
	ret := _m.Called(result)

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.ValidationResult) error); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewReportGenerator creates a new instance of ReportGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReportGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReportGenerator {
	mock := &ReportGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}