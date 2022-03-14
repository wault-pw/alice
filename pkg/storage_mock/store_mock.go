// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package storage_mock is a generated GoMock package.
package storage_mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jwt "github.com/wault-pw/alice/lib/jwt"
	domain "github.com/wault-pw/alice/pkg/domain"
	storage "github.com/wault-pw/alice/pkg/storage"
)

// MockStore is a mock of IStore interface.
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

// ArchiveCard mocks base method.
func (m *MockStore) ArchiveCard(ctx context.Context, ID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArchiveCard", ctx, ID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveCard indicates an expected call of ArchiveCard.
func (mr *MockStoreMockRecorder) ArchiveCard(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveCard", reflect.TypeOf((*MockStore)(nil).ArchiveCard), ctx, ID)
}

// CandidateSession mocks base method.
func (m *MockStore) CandidateSession(ctx context.Context, jti, candidateID string, srp []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CandidateSession", ctx, jti, candidateID, srp)
	ret0, _ := ret[0].(error)
	return ret0
}

// CandidateSession indicates an expected call of CandidateSession.
func (mr *MockStoreMockRecorder) CandidateSession(ctx, jti, candidateID, srp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CandidateSession", reflect.TypeOf((*MockStore)(nil).CandidateSession), ctx, jti, candidateID, srp)
}

// CloneCard mocks base method.
func (m *MockStore) CloneCard(ctx context.Context, oldCardID string, titleEnc []byte) (domain.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneCard", ctx, oldCardID, titleEnc)
	ret0, _ := ret[0].(domain.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneCard indicates an expected call of CloneCard.
func (mr *MockStoreMockRecorder) CloneCard(ctx, oldCardID, titleEnc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneCard", reflect.TypeOf((*MockStore)(nil).CloneCard), ctx, oldCardID, titleEnc)
}

// CreateCardWithItems mocks base method.
func (m *MockStore) CreateCardWithItems(ctx context.Context, card *domain.Card, items []domain.CardItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCardWithItems", ctx, card, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCardWithItems indicates an expected call of CreateCardWithItems.
func (mr *MockStoreMockRecorder) CreateCardWithItems(ctx, card, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCardWithItems", reflect.TypeOf((*MockStore)(nil).CreateCardWithItems), ctx, card, items)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(ctx context.Context, user *domain.User, uw *domain.UserWorkspace, workspace *domain.Workspace, cards []domain.CardWithItems) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user, uw, workspace, cards)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(ctx, user, uw, workspace, cards interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), ctx, user, uw, workspace, cards)
}

// CreateWorkspace mocks base method.
func (m *MockStore) CreateWorkspace(ctx context.Context, uw *domain.UserWorkspace, workspace *domain.Workspace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkspace", ctx, uw, workspace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWorkspace indicates an expected call of CreateWorkspace.
func (mr *MockStoreMockRecorder) CreateWorkspace(ctx, uw, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkspace", reflect.TypeOf((*MockStore)(nil).CreateWorkspace), ctx, uw, workspace)
}

// DeleteCard mocks base method.
func (m *MockStore) DeleteCard(ctx context.Context, cardID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", ctx, cardID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockStoreMockRecorder) DeleteCard(ctx, cardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockStore)(nil).DeleteCard), ctx, cardID)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(ctx context.Context, jti string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, jti)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(ctx, jti interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), ctx, jti)
}

// DeleteWorkspace mocks base method.
func (m *MockStore) DeleteWorkspace(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkspace", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkspace indicates an expected call of DeleteWorkspace.
func (mr *MockStoreMockRecorder) DeleteWorkspace(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkspace", reflect.TypeOf((*MockStore)(nil).DeleteWorkspace), ctx, ID)
}

// FindUser mocks base method.
func (m *MockStore) FindUser(ctx context.Context, ID string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, ID)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockStoreMockRecorder) FindUser(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockStore)(nil).FindUser), ctx, ID)
}

// FindUserIdentity mocks base method.
func (m *MockStore) FindUserIdentity(ctx context.Context, identity string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserIdentity", ctx, identity)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserIdentity indicates an expected call of FindUserIdentity.
func (mr *MockStoreMockRecorder) FindUserIdentity(ctx, identity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserIdentity", reflect.TypeOf((*MockStore)(nil).FindUserIdentity), ctx, identity)
}

// FindUserWithWorkspace mocks base method.
func (m *MockStore) FindUserWithWorkspace(ctx context.Context, ID string) (domain.UserWithWorkspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserWithWorkspace", ctx, ID)
	ret0, _ := ret[0].(domain.UserWithWorkspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserWithWorkspace indicates an expected call of FindUserWithWorkspace.
func (mr *MockStoreMockRecorder) FindUserWithWorkspace(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserWithWorkspace", reflect.TypeOf((*MockStore)(nil).FindUserWithWorkspace), ctx, ID)
}

// IssueSession mocks base method.
func (m *MockStore) IssueSession(ctx context.Context, opts jwt.IOts) (domain.Session, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueSession", ctx, opts)
	ret0, _ := ret[0].(domain.Session)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IssueSession indicates an expected call of IssueSession.
func (mr *MockStoreMockRecorder) IssueSession(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueSession", reflect.TypeOf((*MockStore)(nil).IssueSession), ctx, opts)
}

// ListCardItems mocks base method.
func (m *MockStore) ListCardItems(ctx context.Context, cardID string) ([]domain.CardItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCardItems", ctx, cardID)
	ret0, _ := ret[0].([]domain.CardItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCardItems indicates an expected call of ListCardItems.
func (mr *MockStoreMockRecorder) ListCardItems(ctx, cardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCardItems", reflect.TypeOf((*MockStore)(nil).ListCardItems), ctx, cardID)
}

// ListCardsByWorkspace mocks base method.
func (m *MockStore) ListCardsByWorkspace(ctx context.Context, workspaceID string) ([]domain.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCardsByWorkspace", ctx, workspaceID)
	ret0, _ := ret[0].([]domain.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCardsByWorkspace indicates an expected call of ListCardsByWorkspace.
func (mr *MockStoreMockRecorder) ListCardsByWorkspace(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCardsByWorkspace", reflect.TypeOf((*MockStore)(nil).ListCardsByWorkspace), ctx, workspaceID)
}

// ListUserWithWorkspaces mocks base method.
func (m *MockStore) ListUserWithWorkspaces(ctx context.Context, userID string) ([]domain.UserWithWorkspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserWithWorkspaces", ctx, userID)
	ret0, _ := ret[0].([]domain.UserWithWorkspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserWithWorkspaces indicates an expected call of ListUserWithWorkspaces.
func (mr *MockStoreMockRecorder) ListUserWithWorkspaces(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserWithWorkspaces", reflect.TypeOf((*MockStore)(nil).ListUserWithWorkspaces), ctx, userID)
}

// NominateSession mocks base method.
func (m *MockStore) NominateSession(ctx context.Context, jti string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NominateSession", ctx, jti)
	ret0, _ := ret[0].(error)
	return ret0
}

// NominateSession indicates an expected call of NominateSession.
func (mr *MockStoreMockRecorder) NominateSession(ctx, jti interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NominateSession", reflect.TypeOf((*MockStore)(nil).NominateSession), ctx, jti)
}

// Ping mocks base method.
func (m *MockStore) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockStoreMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockStore)(nil).Ping), ctx)
}

// RetrieveSession mocks base method.
func (m *MockStore) RetrieveSession(ctx context.Context, opts jwt.IOts, token string) (domain.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveSession", ctx, opts, token)
	ret0, _ := ret[0].(domain.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveSession indicates an expected call of RetrieveSession.
func (mr *MockStoreMockRecorder) RetrieveSession(ctx, opts, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveSession", reflect.TypeOf((*MockStore)(nil).RetrieveSession), ctx, opts, token)
}

// TerminateUser mocks base method.
func (m *MockStore) TerminateUser(ctx context.Context, identity, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateUser", ctx, identity, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateUser indicates an expected call of TerminateUser.
func (mr *MockStoreMockRecorder) TerminateUser(ctx, identity, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateUser", reflect.TypeOf((*MockStore)(nil).TerminateUser), ctx, identity, userID)
}

// TruncateAll mocks base method.
func (m *MockStore) TruncateAll(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateAll", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateAll indicates an expected call of TruncateAll.
func (mr *MockStoreMockRecorder) TruncateAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateAll", reflect.TypeOf((*MockStore)(nil).TruncateAll), ctx)
}

// UpdateCardWithItems mocks base method.
func (m *MockStore) UpdateCardWithItems(ctx context.Context, card *domain.Card, items []domain.CardItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCardWithItems", ctx, card, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCardWithItems indicates an expected call of UpdateCardWithItems.
func (mr *MockStoreMockRecorder) UpdateCardWithItems(ctx, card, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCardWithItems", reflect.TypeOf((*MockStore)(nil).UpdateCardWithItems), ctx, card, items)
}

// MockIBuilder is a mock of IBuilder interface.
type MockIBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockIBuilderMockRecorder
}

// MockIBuilderMockRecorder is the mock recorder for MockIBuilder.
type MockIBuilderMockRecorder struct {
	mock *MockIBuilder
}

// NewMockIBuilder creates a new mock instance.
func NewMockIBuilder(ctrl *gomock.Controller) *MockIBuilder {
	mock := &MockIBuilder{ctrl: ctrl}
	mock.recorder = &MockIBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBuilder) EXPECT() *MockIBuilderMockRecorder {
	return m.recorder
}

// ToSql mocks base method.
func (m *MockIBuilder) ToSql() (string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSql")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ToSql indicates an expected call of ToSql.
func (mr *MockIBuilderMockRecorder) ToSql() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSql", reflect.TypeOf((*MockIBuilder)(nil).ToSql))
}

// MockIConn is a mock of IConn interface.
type MockIConn struct {
	ctrl     *gomock.Controller
	recorder *MockIConnMockRecorder
}

// MockIConnMockRecorder is the mock recorder for MockIConn.
type MockIConnMockRecorder struct {
	mock *MockIConn
}

// NewMockIConn creates a new mock instance.
func NewMockIConn(ctrl *gomock.Controller) *MockIConn {
	mock := &MockIConn{ctrl: ctrl}
	mock.recorder = &MockIConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConn) EXPECT() *MockIConnMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockIConn) ExecContext(ctx context.Context, query string, args ...interface{}) (storage.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(storage.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockIConnMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockIConn)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockIConn) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockIConnMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockIConn)(nil).GetContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockIConn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *storage.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*storage.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockIConnMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockIConn)(nil).QueryRowContext), varargs...)
}

// SelectContext mocks base method.
func (m *MockIConn) SelectContext(ctx context.Context, des interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, des, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockIConnMockRecorder) SelectContext(ctx, des, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, des, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockIConn)(nil).SelectContext), varargs...)
}

// MockIDb is a mock of IDb interface.
type MockIDb struct {
	ctrl     *gomock.Controller
	recorder *MockIDbMockRecorder
}

// MockIDbMockRecorder is the mock recorder for MockIDb.
type MockIDbMockRecorder struct {
	mock *MockIDb
}

// NewMockIDb creates a new mock instance.
func NewMockIDb(ctrl *gomock.Controller) *MockIDb {
	mock := &MockIDb{ctrl: ctrl}
	mock.recorder = &MockIDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDb) EXPECT() *MockIDbMockRecorder {
	return m.recorder
}

// BeginTxx mocks base method.
func (m *MockIDb) BeginTxx(ctx context.Context, opts *storage.TxOpts) (storage.ITransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxx", ctx, opts)
	ret0, _ := ret[0].(storage.ITransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTxx indicates an expected call of BeginTxx.
func (mr *MockIDbMockRecorder) BeginTxx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxx", reflect.TypeOf((*MockIDb)(nil).BeginTxx), ctx, opts)
}

// ExecContext mocks base method.
func (m *MockIDb) ExecContext(ctx context.Context, query string, args ...interface{}) (storage.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(storage.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockIDbMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockIDb)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockIDb) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockIDbMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockIDb)(nil).GetContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockIDb) QueryRowContext(ctx context.Context, query string, args ...interface{}) *storage.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*storage.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockIDbMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockIDb)(nil).QueryRowContext), varargs...)
}

// SelectContext mocks base method.
func (m *MockIDb) SelectContext(ctx context.Context, des interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, des, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockIDbMockRecorder) SelectContext(ctx, des, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, des, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockIDb)(nil).SelectContext), varargs...)
}

// SqlDB mocks base method.
func (m *MockIDb) SqlDB() *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SqlDB")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// SqlDB indicates an expected call of SqlDB.
func (mr *MockIDbMockRecorder) SqlDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SqlDB", reflect.TypeOf((*MockIDb)(nil).SqlDB))
}

// MockITransaction is a mock of ITransaction interface.
type MockITransaction struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionMockRecorder
}

// MockITransactionMockRecorder is the mock recorder for MockITransaction.
type MockITransactionMockRecorder struct {
	mock *MockITransaction
}

// NewMockITransaction creates a new mock instance.
func NewMockITransaction(ctrl *gomock.Controller) *MockITransaction {
	mock := &MockITransaction{ctrl: ctrl}
	mock.recorder = &MockITransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransaction) EXPECT() *MockITransactionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockITransaction) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockITransactionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockITransaction)(nil).Commit))
}

// ExecContext mocks base method.
func (m *MockITransaction) ExecContext(ctx context.Context, query string, args ...interface{}) (storage.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(storage.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockITransactionMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockITransaction)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockITransaction) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockITransactionMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockITransaction)(nil).GetContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockITransaction) QueryRowContext(ctx context.Context, query string, args ...interface{}) *storage.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*storage.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockITransactionMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockITransaction)(nil).QueryRowContext), varargs...)
}

// Rollback mocks base method.
func (m *MockITransaction) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockITransactionMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockITransaction)(nil).Rollback))
}

// SelectContext mocks base method.
func (m *MockITransaction) SelectContext(ctx context.Context, des interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, des, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockITransactionMockRecorder) SelectContext(ctx, des, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, des, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockITransaction)(nil).SelectContext), varargs...)
}
