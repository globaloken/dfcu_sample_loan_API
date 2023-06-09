// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wizlif/dfcu_bank/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/wizlif/dfcu_bank/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateLoan mocks base method.
func (m *MockStore) CreateLoan(arg0 context.Context, arg1 db.CreateLoanParams) (db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoan", arg0, arg1)
	ret0, _ := ret[0].(db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoan indicates an expected call of CreateLoan.
func (mr *MockStoreMockRecorder) CreateLoan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoan", reflect.TypeOf((*MockStore)(nil).CreateLoan), arg0, arg1)
}

// CreateLog mocks base method.
func (m *MockStore) CreateLog(arg0 context.Context, arg1 db.CreateLogParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLog indicates an expected call of CreateLog.
func (mr *MockStoreMockRecorder) CreateLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLog", reflect.TypeOf((*MockStore)(nil).CreateLog), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// GetLoanByAccUsername mocks base method.
func (m *MockStore) GetLoanByAccUsername(arg0 context.Context, arg1 string) ([]db.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanByAccUsername", arg0, arg1)
	ret0, _ := ret[0].([]db.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanByAccUsername indicates an expected call of GetLoanByAccUsername.
func (mr *MockStoreMockRecorder) GetLoanByAccUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanByAccUsername", reflect.TypeOf((*MockStore)(nil).GetLoanByAccUsername), arg0, arg1)
}

// GetLogStats mocks base method.
func (m *MockStore) GetLogStats(arg0 context.Context) ([]db.GetLogStatsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogStats", arg0)
	ret0, _ := ret[0].([]db.GetLogStatsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogStats indicates an expected call of GetLogStats.
func (mr *MockStoreMockRecorder) GetLogStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogStats", reflect.TypeOf((*MockStore)(nil).GetLogStats), arg0)
}

// GetLogs mocks base method.
func (m *MockStore) GetLogs(arg0 context.Context) ([]db.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0)
	ret0, _ := ret[0].([]db.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockStoreMockRecorder) GetLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockStore)(nil).GetLogs), arg0)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByAccNo mocks base method.
func (m *MockStore) GetUserByAccNo(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAccNo", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAccNo indicates an expected call of GetUserByAccNo.
func (mr *MockStoreMockRecorder) GetUserByAccNo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAccNo", reflect.TypeOf((*MockStore)(nil).GetUserByAccNo), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
