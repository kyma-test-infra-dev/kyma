// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"

// appConverter is an autogenerated mock type for the appConverter type
type appConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *appConverter) ToGQL(in *v1alpha1.Application) gqlschema.Application {
	ret := _m.Called(in)

	var r0 gqlschema.Application
	if rf, ok := ret.Get(0).(func(*v1alpha1.Application) gqlschema.Application); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(gqlschema.Application)
	}

	return r0
}
