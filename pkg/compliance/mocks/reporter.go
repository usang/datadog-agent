// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	event "github.com/DataDog/datadog-agent/pkg/compliance/event"
	mock "github.com/stretchr/testify/mock"
)

// Reporter is an autogenerated mock type for the Reporter type
type Reporter struct {
	mock.Mock
}

// Report provides a mock function with given fields: _a0
func (_m *Reporter) Report(_a0 *event.Event) {
	_m.Called(_a0)
}

// ReportRaw provides a mock function with given fields: content, tags
func (_m *Reporter) ReportRaw(content []byte, tags ...string) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, content)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}
