// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	db "github.com/upper/db/v4"

	mock "github.com/stretchr/testify/mock"
)

// DataTable is an autogenerated mock type for the DataTable type
type DataTable struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *DataTable) Count() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: cond
func (_m *DataTable) Delete(cond db.Cond) error {
	ret := _m.Called(cond)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Cond) error); ok {
		r0 = rf(cond)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields:
func (_m *DataTable) Exists() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: _a0
func (_m *DataTable) Find(_a0 ...interface{}) db.Result {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	ret := _m.Called(_ca...)

	var r0 db.Result
	if rf, ok := ret.Get(0).(func(...interface{}) db.Result); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Result)
		}
	}

	return r0
}

// FindAll provides a mock function with given fields: dataAddress
func (_m *DataTable) FindAll(dataAddress interface{}) error {
	ret := _m.Called(dataAddress)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dataAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOne provides a mock function with given fields: cond, dataAddress
func (_m *DataTable) FindOne(cond db.Cond, dataAddress interface{}) error {
	ret := _m.Called(cond, dataAddress)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Cond, interface{}) error); ok {
		r0 = rf(cond, dataAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: _a0
func (_m *DataTable) Insert(_a0 interface{}) (db.InsertResult, error) {
	ret := _m.Called(_a0)

	var r0 db.InsertResult
	if rf, ok := ret.Get(0).(func(interface{}) db.InsertResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.InsertResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertReturning provides a mock function with given fields: _a0
func (_m *DataTable) InsertReturning(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *DataTable) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Session provides a mock function with given fields:
func (_m *DataTable) Session() db.Session {
	ret := _m.Called()

	var r0 db.Session
	if rf, ok := ret.Get(0).(func() db.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Session)
		}
	}

	return r0
}

// Truncate provides a mock function with given fields:
func (_m *DataTable) Truncate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReturning provides a mock function with given fields: _a0
func (_m *DataTable) UpdateReturning(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}