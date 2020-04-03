// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
	mock "github.com/stretchr/testify/mock"

	resource "github.com/kyma-project/kyma/components/console-backend-service/pkg/resource"

	v1alpha1 "github.com/knative/eventing/pkg/apis/eventing/v1alpha1"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CompareSubscribers provides a mock function with given fields: fir, sec
func (_m *Service) CompareSubscribers(fir *gqlschema.SubscriberInput, sec *v1alpha1.Trigger) bool {
	ret := _m.Called(fir, sec)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gqlschema.SubscriberInput, *v1alpha1.Trigger) bool); ok {
		r0 = rf(fir, sec)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Create provides a mock function with given fields: _a0
func (_m *Service) Create(_a0 *v1alpha1.Trigger) (*v1alpha1.Trigger, error) {
	ret := _m.Called(_a0)

	var r0 *v1alpha1.Trigger
	if rf, ok := ret.Get(0).(func(*v1alpha1.Trigger) *v1alpha1.Trigger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.Trigger) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMany provides a mock function with given fields: triggers
func (_m *Service) CreateMany(triggers []*v1alpha1.Trigger) ([]*v1alpha1.Trigger, error) {
	ret := _m.Called(triggers)

	var r0 []*v1alpha1.Trigger
	if rf, ok := ret.Get(0).(func([]*v1alpha1.Trigger) []*v1alpha1.Trigger); ok {
		r0 = rf(triggers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1alpha1.Trigger) error); ok {
		r1 = rf(triggers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *Service) Delete(_a0 gqlschema.TriggerMetadataInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(gqlschema.TriggerMetadataInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMany provides a mock function with given fields: triggers
func (_m *Service) DeleteMany(triggers []gqlschema.TriggerMetadataInput) error {
	ret := _m.Called(triggers)

	var r0 error
	if rf, ok := ret.Get(0).(func([]gqlschema.TriggerMetadataInput) error); ok {
		r0 = rf(triggers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: namespace, subscriber
func (_m *Service) List(namespace string, subscriber *gqlschema.SubscriberInput) ([]*v1alpha1.Trigger, error) {
	ret := _m.Called(namespace, subscriber)

	var r0 []*v1alpha1.Trigger
	if rf, ok := ret.Get(0).(func(string, *gqlschema.SubscriberInput) []*v1alpha1.Trigger); ok {
		r0 = rf(namespace, subscriber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *gqlschema.SubscriberInput) error); ok {
		r1 = rf(namespace, subscriber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: listener
func (_m *Service) Subscribe(listener resource.Listener) {
	_m.Called(listener)
}

// Unsubscribe provides a mock function with given fields: listener
func (_m *Service) Unsubscribe(listener resource.Listener) {
	_m.Called(listener)
}
