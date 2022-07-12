// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/dinorain/checkoutaja/internal/models"
	utils "github.com/dinorain/checkoutaja/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockProductUseCase is a mock of ProductUseCase interface.
type MockProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUseCaseMockRecorder
}

// MockProductUseCaseMockRecorder is the mock recorder for MockProductUseCase.
type MockProductUseCaseMockRecorder struct {
	mock *MockProductUseCase
}

// NewMockProductUseCase creates a new mock instance.
func NewMockProductUseCase(ctrl *gomock.Controller) *MockProductUseCase {
	mock := &MockProductUseCase{ctrl: ctrl}
	mock.recorder = &MockProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUseCase) EXPECT() *MockProductUseCaseMockRecorder {
	return m.recorder
}

// CachedFindById mocks base method.
func (m *MockProductUseCase) CachedFindById(ctx context.Context, userID uuid.UUID) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CachedFindById", ctx, userID)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CachedFindById indicates an expected call of CachedFindById.
func (mr *MockProductUseCaseMockRecorder) CachedFindById(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CachedFindById", reflect.TypeOf((*MockProductUseCase)(nil).CachedFindById), ctx, userID)
}

// Create mocks base method.
func (m *MockProductUseCase) Create(ctx context.Context, user *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductUseCaseMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUseCase)(nil).Create), ctx, user)
}

// DeleteById mocks base method.
func (m *MockProductUseCase) DeleteById(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockProductUseCaseMockRecorder) DeleteById(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockProductUseCase)(nil).DeleteById), ctx, userID)
}

// FindAll mocks base method.
func (m *MockProductUseCase) FindAll(ctx context.Context, pagination *utils.Pagination) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, pagination)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockProductUseCaseMockRecorder) FindAll(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockProductUseCase)(nil).FindAll), ctx, pagination)
}

// FindById mocks base method.
func (m *MockProductUseCase) FindById(ctx context.Context, userID uuid.UUID) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, userID)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockProductUseCaseMockRecorder) FindById(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductUseCase)(nil).FindById), ctx, userID)
}

// UpdateById mocks base method.
func (m *MockProductUseCase) UpdateById(ctx context.Context, user *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", ctx, user)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockProductUseCaseMockRecorder) UpdateById(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockProductUseCase)(nil).UpdateById), ctx, user)
}
