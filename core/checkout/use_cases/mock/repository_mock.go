// Code generated by MockGen. DO NOT EDIT.
// Source: core/checkout/domain/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/guil95/grpcApi/core/checkout/domain"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetGiftProducts mocks base method.
func (m *MockRepository) GetGiftProducts() []domain.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGiftProducts")
	ret0, _ := ret[0].([]domain.Product)
	return ret0
}

// GetGiftProducts indicates an expected call of GetGiftProducts.
func (mr *MockRepositoryMockRecorder) GetGiftProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGiftProducts", reflect.TypeOf((*MockRepository)(nil).GetGiftProducts))
}

// GetProducts mocks base method.
func (m *MockRepository) GetProducts() []domain.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts")
	ret0, _ := ret[0].([]domain.Product)
	return ret0
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockRepositoryMockRecorder) GetProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockRepository)(nil).GetProducts))
}

// GetProductsByChart mocks base method.
func (m *MockRepository) GetProductsByChart(chart *domain.Chart) []domain.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByChart", chart)
	ret0, _ := ret[0].([]domain.Product)
	return ret0
}

// GetProductsByChart indicates an expected call of GetProductsByChart.
func (mr *MockRepositoryMockRecorder) GetProductsByChart(chart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByChart", reflect.TypeOf((*MockRepository)(nil).GetProductsByChart), chart)
}
