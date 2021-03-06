// Code generated by MockGen. DO NOT EDIT.
// Source: ./trie/trie.go

// Package mock_trie is a generated GoMock package.
package mock_trie

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	db "github.com/iotexproject/iotex-core/db"
	hash "github.com/iotexproject/iotex-core/pkg/hash"
	reflect "reflect"
)

// MockTrie is a mock of Trie interface
type MockTrie struct {
	ctrl     *gomock.Controller
	recorder *MockTrieMockRecorder
}

// MockTrieMockRecorder is the mock recorder for MockTrie
type MockTrieMockRecorder struct {
	mock *MockTrie
}

// NewMockTrie creates a new mock instance
func NewMockTrie(ctrl *gomock.Controller) *MockTrie {
	mock := &MockTrie{ctrl: ctrl}
	mock.recorder = &MockTrieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrie) EXPECT() *MockTrieMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockTrie) Start(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockTrieMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTrie)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockTrie) Stop(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockTrieMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTrie)(nil).Stop), arg0)
}

// TrieDB mocks base method
func (m *MockTrie) TrieDB() db.KVStore {
	ret := m.ctrl.Call(m, "TrieDB")
	ret0, _ := ret[0].(db.KVStore)
	return ret0
}

// TrieDB indicates an expected call of TrieDB
func (mr *MockTrieMockRecorder) TrieDB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrieDB", reflect.TypeOf((*MockTrie)(nil).TrieDB))
}

// Upsert mocks base method
func (m *MockTrie) Upsert(arg0, arg1 []byte) error {
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockTrieMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTrie)(nil).Upsert), arg0, arg1)
}

// Get mocks base method
func (m *MockTrie) Get(arg0 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTrieMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTrie)(nil).Get), arg0)
}

// Delete mocks base method
func (m *MockTrie) Delete(arg0 []byte) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTrieMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrie)(nil).Delete), arg0)
}

// Commit mocks base method
func (m *MockTrie) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTrieMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTrie)(nil).Commit))
}

// RootHash mocks base method
func (m *MockTrie) RootHash() hash.Hash32B {
	ret := m.ctrl.Call(m, "RootHash")
	ret0, _ := ret[0].(hash.Hash32B)
	return ret0
}

// RootHash indicates an expected call of RootHash
func (mr *MockTrieMockRecorder) RootHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHash", reflect.TypeOf((*MockTrie)(nil).RootHash))
}

// SetRoot mocks base method
func (m *MockTrie) SetRoot(arg0 hash.Hash32B) error {
	ret := m.ctrl.Call(m, "SetRoot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRoot indicates an expected call of SetRoot
func (mr *MockTrieMockRecorder) SetRoot(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoot", reflect.TypeOf((*MockTrie)(nil).SetRoot), arg0)
}
