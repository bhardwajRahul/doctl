package mocks

import "github.com/digitalocean/doctl/do"
import "github.com/stretchr/testify/mock"

import "github.com/digitalocean/godo"

// Generated: please do not edit by hand

type TagsService struct {
	mock.Mock
}

func (_m *TagsService) List() (do.Tags, error) {
	ret := _m.Called()

	var r0 do.Tags
	if rf, ok := ret.Get(0).(func() do.Tags); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(do.Tags)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *TagsService) Get(_a0 string) (*do.Tag, error) {
	ret := _m.Called(_a0)

	var r0 *do.Tag
	if rf, ok := ret.Get(0).(func(string) *do.Tag); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *TagsService) Create(_a0 *godo.TagCreateRequest) (*do.Tag, error) {
	ret := _m.Called(_a0)

	var r0 *do.Tag
	if rf, ok := ret.Get(0).(func(*godo.TagCreateRequest) *do.Tag); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.TagCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *TagsService) Update(_a0 string, _a1 *godo.TagUpdateRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.TagUpdateRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *TagsService) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *TagsService) TagResources(_a0 string, _a1 *godo.TagResourcesRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.TagResourcesRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *TagsService) UntagResources(_a0 string, _a1 *godo.UntagResourcesRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.UntagResourcesRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
