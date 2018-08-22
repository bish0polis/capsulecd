// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/config/interface.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	gomock "github.com/golang/mock/gomock"
	viper "github.com/spf13/viper"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockInterface) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockInterfaceMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockInterface)(nil).Init))
}

// ReadConfig mocks base method
func (m *MockInterface) ReadConfig(configFilePath string) error {
	ret := m.ctrl.Call(m, "ReadConfig", configFilePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadConfig indicates an expected call of ReadConfig
func (mr *MockInterfaceMockRecorder) ReadConfig(configFilePath interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConfig", reflect.TypeOf((*MockInterface)(nil).ReadConfig), configFilePath)
}

// Set mocks base method
func (m *MockInterface) Set(key string, value interface{}) {
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set
func (mr *MockInterfaceMockRecorder) Set(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInterface)(nil).Set), key, value)
}

// SetDefault mocks base method
func (m *MockInterface) SetDefault(key string, value interface{}) {
	m.ctrl.Call(m, "SetDefault", key, value)
}

// SetDefault indicates an expected call of SetDefault
func (mr *MockInterfaceMockRecorder) SetDefault(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefault", reflect.TypeOf((*MockInterface)(nil).SetDefault), key, value)
}

// AllSettings mocks base method
func (m *MockInterface) AllSettings() map[string]interface{} {
	ret := m.ctrl.Call(m, "AllSettings")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// AllSettings indicates an expected call of AllSettings
func (mr *MockInterfaceMockRecorder) AllSettings() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSettings", reflect.TypeOf((*MockInterface)(nil).AllSettings))
}

// IsSet mocks base method
func (m *MockInterface) IsSet(key string) bool {
	ret := m.ctrl.Call(m, "IsSet", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSet indicates an expected call of IsSet
func (mr *MockInterfaceMockRecorder) IsSet(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSet", reflect.TypeOf((*MockInterface)(nil).IsSet), key)
}

// Get mocks base method
func (m *MockInterface) Get(key string) interface{} {
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), key)
}

// GetBool mocks base method
func (m *MockInterface) GetBool(key string) bool {
	ret := m.ctrl.Call(m, "GetBool", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool
func (mr *MockInterfaceMockRecorder) GetBool(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockInterface)(nil).GetBool), key)
}

// GetInt mocks base method
func (m *MockInterface) GetInt(key string) int {
	ret := m.ctrl.Call(m, "GetInt", key)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt
func (mr *MockInterfaceMockRecorder) GetInt(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockInterface)(nil).GetInt), key)
}

// GetString mocks base method
func (m *MockInterface) GetString(key string) string {
	ret := m.ctrl.Call(m, "GetString", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString
func (mr *MockInterfaceMockRecorder) GetString(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockInterface)(nil).GetString), key)
}

// GetStringSlice mocks base method
func (m *MockInterface) GetStringSlice(key string) []string {
	ret := m.ctrl.Call(m, "GetStringSlice", key)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStringSlice indicates an expected call of GetStringSlice
func (mr *MockInterfaceMockRecorder) GetStringSlice(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringSlice", reflect.TypeOf((*MockInterface)(nil).GetStringSlice), key)
}

// GetBase64Decoded mocks base method
func (m *MockInterface) GetBase64Decoded(key string) (string, error) {
	ret := m.ctrl.Call(m, "GetBase64Decoded", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBase64Decoded indicates an expected call of GetBase64Decoded
func (mr *MockInterfaceMockRecorder) GetBase64Decoded(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBase64Decoded", reflect.TypeOf((*MockInterface)(nil).GetBase64Decoded), key)
}

// UnmarshalKey mocks base method
func (m *MockInterface) UnmarshalKey(key string, rawVal interface{}, decoder ...viper.DecoderConfigOption) error {
	varargs := []interface{}{key, rawVal}
	for _, a := range decoder {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnmarshalKey", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalKey indicates an expected call of UnmarshalKey
func (mr *MockInterfaceMockRecorder) UnmarshalKey(key, rawVal interface{}, decoder ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key, rawVal}, decoder...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalKey", reflect.TypeOf((*MockInterface)(nil).UnmarshalKey), varargs...)
}
