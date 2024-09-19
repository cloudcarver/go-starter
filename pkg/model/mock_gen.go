// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/model/model.go

// Package model is a generated GoMock package.
package model

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	querier "github.com/xich-dev/go-starter/pkg/model/querier"
)

// MockModelInterface is a mock of ModelInterface interface.
type MockModelInterface struct {
	ctrl     *gomock.Controller
	recorder *MockModelInterfaceMockRecorder
}

// MockModelInterfaceMockRecorder is the mock recorder for MockModelInterface.
type MockModelInterfaceMockRecorder struct {
	mock *MockModelInterface
}

// NewMockModelInterface creates a new mock instance.
func NewMockModelInterface(ctrl *gomock.Controller) *MockModelInterface {
	mock := &MockModelInterface{ctrl: ctrl}
	mock.recorder = &MockModelInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelInterface) EXPECT() *MockModelInterfaceMockRecorder {
	return m.recorder
}

// AddUserAccessRule mocks base method.
func (m *MockModelInterface) AddUserAccessRule(ctx context.Context, arg querier.AddUserAccessRuleParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserAccessRule", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserAccessRule indicates an expected call of AddUserAccessRule.
func (mr *MockModelInterfaceMockRecorder) AddUserAccessRule(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserAccessRule", reflect.TypeOf((*MockModelInterface)(nil).AddUserAccessRule), ctx, arg)
}

// CreateOrg mocks base method.
func (m *MockModelInterface) CreateOrg(ctx context.Context, name string) (*querier.Org, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrg", ctx, name)
	ret0, _ := ret[0].(*querier.Org)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrg indicates an expected call of CreateOrg.
func (mr *MockModelInterfaceMockRecorder) CreateOrg(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrg", reflect.TypeOf((*MockModelInterface)(nil).CreateOrg), ctx, name)
}

// CreateUser mocks base method.
func (m *MockModelInterface) CreateUser(ctx context.Context, arg querier.CreateUserParams) (*querier.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(*querier.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockModelInterfaceMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockModelInterface)(nil).CreateUser), ctx, arg)
}

// GetAccessRule mocks base method.
func (m *MockModelInterface) GetAccessRule(ctx context.Context, name string) (*querier.AccessRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessRule", ctx, name)
	ret0, _ := ret[0].(*querier.AccessRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessRule indicates an expected call of GetAccessRule.
func (mr *MockModelInterfaceMockRecorder) GetAccessRule(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessRule", reflect.TypeOf((*MockModelInterface)(nil).GetAccessRule), ctx, name)
}

// GetOrgInfoByOrgId mocks base method.
func (m *MockModelInterface) GetOrgInfoByOrgId(ctx context.Context, id uuid.UUID) (*querier.Org, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgInfoByOrgId", ctx, id)
	ret0, _ := ret[0].(*querier.Org)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgInfoByOrgId indicates an expected call of GetOrgInfoByOrgId.
func (mr *MockModelInterfaceMockRecorder) GetOrgInfoByOrgId(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgInfoByOrgId", reflect.TypeOf((*MockModelInterface)(nil).GetOrgInfoByOrgId), ctx, id)
}

// GetPhoneCode mocks base method.
func (m *MockModelInterface) GetPhoneCode(ctx context.Context, arg querier.GetPhoneCodeParams) (*querier.PhoneCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhoneCode", ctx, arg)
	ret0, _ := ret[0].(*querier.PhoneCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhoneCode indicates an expected call of GetPhoneCode.
func (mr *MockModelInterfaceMockRecorder) GetPhoneCode(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoneCode", reflect.TypeOf((*MockModelInterface)(nil).GetPhoneCode), ctx, arg)
}

// GetUser mocks base method.
func (m *MockModelInterface) GetUser(ctx context.Context, phone string) (*querier.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, phone)
	ret0, _ := ret[0].(*querier.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockModelInterfaceMockRecorder) GetUser(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockModelInterface)(nil).GetUser), ctx, phone)
}

// GetUserAccessRuleNames mocks base method.
func (m *MockModelInterface) GetUserAccessRuleNames(ctx context.Context, userID uuid.UUID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAccessRuleNames", ctx, userID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccessRuleNames indicates an expected call of GetUserAccessRuleNames.
func (mr *MockModelInterfaceMockRecorder) GetUserAccessRuleNames(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccessRuleNames", reflect.TypeOf((*MockModelInterface)(nil).GetUserAccessRuleNames), ctx, userID)
}

// GetUserAccessRules mocks base method.
func (m *MockModelInterface) GetUserAccessRules(ctx context.Context, userID uuid.UUID) ([]*querier.GetUserAccessRulesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAccessRules", ctx, userID)
	ret0, _ := ret[0].([]*querier.GetUserAccessRulesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccessRules indicates an expected call of GetUserAccessRules.
func (mr *MockModelInterfaceMockRecorder) GetUserAccessRules(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccessRules", reflect.TypeOf((*MockModelInterface)(nil).GetUserAccessRules), ctx, userID)
}

// InTransaction mocks base method.
func (m *MockModelInterface) InTransaction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InTransaction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InTransaction indicates an expected call of InTransaction.
func (mr *MockModelInterfaceMockRecorder) InTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InTransaction", reflect.TypeOf((*MockModelInterface)(nil).InTransaction))
}

// IsPhoneExist mocks base method.
func (m *MockModelInterface) IsPhoneExist(ctx context.Context, phone string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPhoneExist", ctx, phone)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPhoneExist indicates an expected call of IsPhoneExist.
func (mr *MockModelInterfaceMockRecorder) IsPhoneExist(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPhoneExist", reflect.TypeOf((*MockModelInterface)(nil).IsPhoneExist), ctx, phone)
}

// IsUsernameExist mocks base method.
func (m *MockModelInterface) IsUsernameExist(ctx context.Context, name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUsernameExist", ctx, name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUsernameExist indicates an expected call of IsUsernameExist.
func (mr *MockModelInterfaceMockRecorder) IsUsernameExist(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUsernameExist", reflect.TypeOf((*MockModelInterface)(nil).IsUsernameExist), ctx, name)
}

// MarkPhoneCodeUsed mocks base method.
func (m *MockModelInterface) MarkPhoneCodeUsed(ctx context.Context, arg querier.MarkPhoneCodeUsedParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkPhoneCodeUsed", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkPhoneCodeUsed indicates an expected call of MarkPhoneCodeUsed.
func (mr *MockModelInterfaceMockRecorder) MarkPhoneCodeUsed(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkPhoneCodeUsed", reflect.TypeOf((*MockModelInterface)(nil).MarkPhoneCodeUsed), ctx, arg)
}

// RemoveUserAccessRule mocks base method.
func (m *MockModelInterface) RemoveUserAccessRule(ctx context.Context, arg querier.RemoveUserAccessRuleParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserAccessRule", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserAccessRule indicates an expected call of RemoveUserAccessRule.
func (mr *MockModelInterfaceMockRecorder) RemoveUserAccessRule(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserAccessRule", reflect.TypeOf((*MockModelInterface)(nil).RemoveUserAccessRule), ctx, arg)
}

// RunTransaction mocks base method.
func (m *MockModelInterface) RunTransaction(ctx context.Context, f func(ModelInterface) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockModelInterfaceMockRecorder) RunTransaction(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockModelInterface)(nil).RunTransaction), ctx, f)
}

// UpdateOrgOwnerID mocks base method.
func (m *MockModelInterface) UpdateOrgOwnerID(ctx context.Context, arg querier.UpdateOrgOwnerIDParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrgOwnerID", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrgOwnerID indicates an expected call of UpdateOrgOwnerID.
func (mr *MockModelInterfaceMockRecorder) UpdateOrgOwnerID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgOwnerID", reflect.TypeOf((*MockModelInterface)(nil).UpdateOrgOwnerID), ctx, arg)
}

// UpdateUserPasswordByPhone mocks base method.
func (m *MockModelInterface) UpdateUserPasswordByPhone(ctx context.Context, arg querier.UpdateUserPasswordByPhoneParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPasswordByPhone", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPasswordByPhone indicates an expected call of UpdateUserPasswordByPhone.
func (mr *MockModelInterfaceMockRecorder) UpdateUserPasswordByPhone(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPasswordByPhone", reflect.TypeOf((*MockModelInterface)(nil).UpdateUserPasswordByPhone), ctx, arg)
}

// UpsertPhoneCode mocks base method.
func (m *MockModelInterface) UpsertPhoneCode(ctx context.Context, arg querier.UpsertPhoneCodeParams) (*querier.PhoneCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPhoneCode", ctx, arg)
	ret0, _ := ret[0].(*querier.PhoneCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertPhoneCode indicates an expected call of UpsertPhoneCode.
func (mr *MockModelInterfaceMockRecorder) UpsertPhoneCode(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPhoneCode", reflect.TypeOf((*MockModelInterface)(nil).UpsertPhoneCode), ctx, arg)
}
