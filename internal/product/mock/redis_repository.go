// Code generated by MockGen. DO NOT EDIT.
// Source: redis_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/dinorain/checkoutaja/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockProductRedisRepository is a mock of ProductRedisRepository interface.
type MockProductRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRedisRepositoryMockRecorder
}

// MockProductRedisRepositoryMockRecorder is the mock recorder for MockProductRedisRepository.
type MockProductRedisRepositoryMockRecorder struct {
	mock *MockProductRedisRepository
}

// NewMockProductRedisRepository creates a new mock instance.
func NewMockProductRedisRepository(ctrl *gomock.Controller) *MockProductRedisRepository {
	mock := &MockProductRedisRepository{ctrl: ctrl}
	mock.recorder = &MockProductRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRedisRepository) EXPECT() *MockProductRedisRepositoryMockRecorder {
	return m.recorder
}

// DeleteProductCtx mocks base method.
func (m *MockProductRedisRepository) DeleteProductCtx(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductCtx", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductCtx indicates an expected call of DeleteProductCtx.
func (mr *MockProductRedisRepositoryMockRecorder) DeleteProductCtx(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductCtx", reflect.TypeOf((*MockProductRedisRepository)(nil).DeleteProductCtx), ctx, key)
}

// GetByIDCtx mocks base method.
func (m *MockProductRedisRepository) GetByIDCtx(ctx context.Context, key string) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDCtx", ctx, key)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDCtx indicates an expected call of GetByIDCtx.
func (mr *MockProductRedisRepositoryMockRecorder) GetByIDCtx(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDCtx", reflect.TypeOf((*MockProductRedisRepository)(nil).GetByIDCtx), ctx, key)
}

// SetProductCtx mocks base method.
func (m *MockProductRedisRepository) SetProductCtx(ctx context.Context, key string, seconds int, user *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProductCtx", ctx, key, seconds, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProductCtx indicates an expected call of SetProductCtx.
func (mr *MockProductRedisRepositoryMockRecorder) SetProductCtx(ctx, key, seconds, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProductCtx", reflect.TypeOf((*MockProductRedisRepository)(nil).SetProductCtx), ctx, key, seconds, user)
}
