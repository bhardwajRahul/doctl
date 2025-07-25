// Code generated by MockGen. DO NOT EDIT.
// Source: genai.go
//
// Generated by this command:
//
//	mockgen -source genai.go -package=mocks GenAIService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "go.uber.org/mock/gomock"
)

// MockGenAIService is a mock of GenAIService interface.
type MockGenAIService struct {
	ctrl     *gomock.Controller
	recorder *MockGenAIServiceMockRecorder
	isgomock struct{}
}

// MockGenAIServiceMockRecorder is the mock recorder for MockGenAIService.
type MockGenAIServiceMockRecorder struct {
	mock *MockGenAIService
}

// NewMockGenAIService creates a new mock instance.
func NewMockGenAIService(ctrl *gomock.Controller) *MockGenAIService {
	mock := &MockGenAIService{ctrl: ctrl}
	mock.recorder = &MockGenAIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenAIService) EXPECT() *MockGenAIServiceMockRecorder {
	return m.recorder
}

// AddAgentRoute mocks base method.
func (m *MockGenAIService) AddAgentRoute(parentAgentID, childAgentID string) (*do.AgentRouteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAgentRoute", parentAgentID, childAgentID)
	ret0, _ := ret[0].(*do.AgentRouteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAgentRoute indicates an expected call of AddAgentRoute.
func (mr *MockGenAIServiceMockRecorder) AddAgentRoute(parentAgentID, childAgentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAgentRoute", reflect.TypeOf((*MockGenAIService)(nil).AddAgentRoute), parentAgentID, childAgentID)
}

// AddKnowledgeBaseDataSource mocks base method.
func (m *MockGenAIService) AddKnowledgeBaseDataSource(knowledgeBaseID string, req *godo.AddKnowledgeBaseDataSourceRequest) (*do.KnowledgeBaseDataSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKnowledgeBaseDataSource", knowledgeBaseID, req)
	ret0, _ := ret[0].(*do.KnowledgeBaseDataSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddKnowledgeBaseDataSource indicates an expected call of AddKnowledgeBaseDataSource.
func (mr *MockGenAIServiceMockRecorder) AddKnowledgeBaseDataSource(knowledgeBaseID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKnowledgeBaseDataSource", reflect.TypeOf((*MockGenAIService)(nil).AddKnowledgeBaseDataSource), knowledgeBaseID, req)
}

// AttachKnowledgeBaseToAgent mocks base method.
func (m *MockGenAIService) AttachKnowledgeBaseToAgent(agentId, knowledgeBaseID string) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachKnowledgeBaseToAgent", agentId, knowledgeBaseID)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachKnowledgeBaseToAgent indicates an expected call of AttachKnowledgeBaseToAgent.
func (mr *MockGenAIServiceMockRecorder) AttachKnowledgeBaseToAgent(agentId, knowledgeBaseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachKnowledgeBaseToAgent", reflect.TypeOf((*MockGenAIService)(nil).AttachKnowledgeBaseToAgent), agentId, knowledgeBaseID)
}

// CreateAgent mocks base method.
func (m *MockGenAIService) CreateAgent(req *godo.AgentCreateRequest) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgent", req)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAgent indicates an expected call of CreateAgent.
func (mr *MockGenAIServiceMockRecorder) CreateAgent(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgent", reflect.TypeOf((*MockGenAIService)(nil).CreateAgent), req)
}

// CreateFunctionRoute mocks base method.
func (m *MockGenAIService) CreateFunctionRoute(id string, req *godo.FunctionRouteCreateRequest) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFunctionRoute", id, req)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFunctionRoute indicates an expected call of CreateFunctionRoute.
func (mr *MockGenAIServiceMockRecorder) CreateFunctionRoute(id, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFunctionRoute", reflect.TypeOf((*MockGenAIService)(nil).CreateFunctionRoute), id, req)
}

// CreateKnowledgeBase mocks base method.
func (m *MockGenAIService) CreateKnowledgeBase(req *godo.KnowledgeBaseCreateRequest) (*do.KnowledgeBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKnowledgeBase", req)
	ret0, _ := ret[0].(*do.KnowledgeBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKnowledgeBase indicates an expected call of CreateKnowledgeBase.
func (mr *MockGenAIServiceMockRecorder) CreateKnowledgeBase(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKnowledgeBase", reflect.TypeOf((*MockGenAIService)(nil).CreateKnowledgeBase), req)
}

// CreateAgentAPIKey mocks base method.
func (m *MockGenAIService) CreateAgentAPIKey(agentID string, req *godo.AgentAPIKeyCreateRequest) (*do.ApiKeyInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgentAPIKey", agentID, req)
	ret0, _ := ret[0].(*do.ApiKeyInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAgentAPIKey indicates an expected call of CreateAgentAPIKey.
func (mr *MockGenAIServiceMockRecorder) CreateAgentAPIKey(agentID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgentAPIKey", reflect.TypeOf((*MockGenAIService)(nil).CreateAgentAPIKey), agentID, req)
}

// DeleteAgent mocks base method.
func (m *MockGenAIService) DeleteAgent(agentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgent", agentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgent indicates an expected call of DeleteAgent.
func (mr *MockGenAIServiceMockRecorder) DeleteAgent(agentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockGenAIService)(nil).DeleteAgent), agentID)
}

// DeleteAgentRoute mocks base method.
func (m *MockGenAIService) DeleteAgentRoute(parentAgentID, childAgentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgentRoute", parentAgentID, childAgentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgentRoute indicates an expected call of DeleteAgentRoute.
func (mr *MockGenAIServiceMockRecorder) DeleteAgentRoute(parentAgentID, childAgentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentRoute", reflect.TypeOf((*MockGenAIService)(nil).DeleteAgentRoute), parentAgentID, childAgentID)
}

// DeleteFunctionRoute mocks base method.
func (m *MockGenAIService) DeleteFunctionRoute(agent_id, function_id string) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFunctionRoute", agent_id, function_id)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFunctionRoute indicates an expected call of DeleteFunctionRoute.
func (mr *MockGenAIServiceMockRecorder) DeleteFunctionRoute(agent_id, function_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFunctionRoute", reflect.TypeOf((*MockGenAIService)(nil).DeleteFunctionRoute), agent_id, function_id)
}

// DeleteKnowledgeBase mocks base method.
func (m *MockGenAIService) DeleteKnowledgeBase(knowledgeBaseID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKnowledgeBase", knowledgeBaseID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKnowledgeBase indicates an expected call of DeleteKnowledgeBase.
func (mr *MockGenAIServiceMockRecorder) DeleteKnowledgeBase(knowledgeBaseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKnowledgeBase", reflect.TypeOf((*MockGenAIService)(nil).DeleteKnowledgeBase), knowledgeBaseID)
}

// DeleteKnowledgeBaseDataSource mocks base method.
func (m *MockGenAIService) DeleteKnowledgeBaseDataSource(knowledgeBaseID, dataSourceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKnowledgeBaseDataSource", knowledgeBaseID, dataSourceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKnowledgeBaseDataSource indicates an expected call of DeleteKnowledgeBaseDataSource.
func (mr *MockGenAIServiceMockRecorder) DeleteKnowledgeBaseDataSource(knowledgeBaseID, dataSourceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKnowledgeBaseDataSource", reflect.TypeOf((*MockGenAIService)(nil).DeleteKnowledgeBaseDataSource), knowledgeBaseID, dataSourceID)
}

// DetachKnowledgeBaseToAgent mocks base method.
func (m *MockGenAIService) DetachKnowledgeBaseToAgent(agentId, knowledgeBaseID string) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachKnowledgeBaseToAgent", agentId, knowledgeBaseID)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachKnowledgeBaseToAgent indicates an expected call of DetachKnowledgeBaseToAgent.
func (mr *MockGenAIServiceMockRecorder) DetachKnowledgeBaseToAgent(agentId, knowledgeBaseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachKnowledgeBaseToAgent", reflect.TypeOf((*MockGenAIService)(nil).DetachKnowledgeBaseToAgent), agentId, knowledgeBaseID)
}

// DeleteAgentAPIKey mocks base method.
func (m *MockGenAIService) DeleteAgentAPIKey(agentID, apikeyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgentAPIKey", agentID, apikeyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgentAPIKey indicates an expected call of DeleteAgentAPIKey.
func (mr *MockGenAIServiceMockRecorder) DeleteAgentAPIKey(agentID, apikeyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentAPIKey", reflect.TypeOf((*MockGenAIService)(nil).DeleteAgentAPIKey), agentID, apikeyID)
}

// GetAgent mocks base method.
func (m *MockGenAIService) GetAgent(agentID string) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgent", agentID)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent.
func (mr *MockGenAIServiceMockRecorder) GetAgent(agentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockGenAIService)(nil).GetAgent), agentID)
}

// GetKnowledgeBase mocks base method.
func (m *MockGenAIService) GetKnowledgeBase(knowledgeBaseID string) (*do.KnowledgeBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKnowledgeBase", knowledgeBaseID)
	ret0, _ := ret[0].(*do.KnowledgeBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKnowledgeBase indicates an expected call of GetKnowledgeBase.
func (mr *MockGenAIServiceMockRecorder) GetKnowledgeBase(knowledgeBaseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnowledgeBase", reflect.TypeOf((*MockGenAIService)(nil).GetKnowledgeBase), knowledgeBaseID)
}

// ListAgentAPIKeys mocks base method.
func (m *MockGenAIService) ListAgentAPIKeys(agentId string) (do.ApiKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAgentAPIKeys", agentId)
	ret0, _ := ret[0].(do.ApiKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgentAPIKeys indicates an expected call of ListAgentAPIKeys.
func (mr *MockGenAIServiceMockRecorder) ListAgentAPIKeys(agentId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgentAPIKeys", reflect.TypeOf((*MockGenAIService)(nil).ListAgentAPIKeys), agentId)
}

// ListAgentVersions mocks base method.
func (m *MockGenAIService) ListAgentVersions(agentID string) (do.AgentVersions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAgentVersions", agentID)
	ret0, _ := ret[0].(do.AgentVersions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgentVersions indicates an expected call of ListAgentVersions.
func (mr *MockGenAIServiceMockRecorder) ListAgentVersions(agentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgentVersions", reflect.TypeOf((*MockGenAIService)(nil).ListAgentVersions), agentID)
}

// ListAgents mocks base method.
func (m *MockGenAIService) ListAgents() (do.Agents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAgents")
	ret0, _ := ret[0].(do.Agents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgents indicates an expected call of ListAgents.
func (mr *MockGenAIServiceMockRecorder) ListAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgents", reflect.TypeOf((*MockGenAIService)(nil).ListAgents))
}

// ListKnowledgeBaseDataSources mocks base method.
func (m *MockGenAIService) ListKnowledgeBaseDataSources(knowledgeBaseID string) (do.KnowledgeBaseDataSources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKnowledgeBaseDataSources", knowledgeBaseID)
	ret0, _ := ret[0].(do.KnowledgeBaseDataSources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKnowledgeBaseDataSources indicates an expected call of ListKnowledgeBaseDataSources.
func (mr *MockGenAIServiceMockRecorder) ListKnowledgeBaseDataSources(knowledgeBaseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKnowledgeBaseDataSources", reflect.TypeOf((*MockGenAIService)(nil).ListKnowledgeBaseDataSources), knowledgeBaseID)
}

// ListKnowledgeBases mocks base method.
func (m *MockGenAIService) ListKnowledgeBases() (do.KnowledgeBases, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKnowledgeBases")
	ret0, _ := ret[0].(do.KnowledgeBases)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKnowledgeBases indicates an expected call of ListKnowledgeBases.
func (mr *MockGenAIServiceMockRecorder) ListKnowledgeBases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKnowledgeBases", reflect.TypeOf((*MockGenAIService)(nil).ListKnowledgeBases))
}

// RegenerateAgentAPIKey mocks base method.
func (m *MockGenAIService) RegenerateAgentAPIKey(agentID, apikeyID string) (*do.ApiKeyInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegenerateAgentAPIKey", agentID, apikeyID)
	ret0, _ := ret[0].(*do.ApiKeyInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegenerateAgentAPIKey indicates an expected call of RegenerateAgentAPIKey.
func (mr *MockGenAIServiceMockRecorder) RegenerateAgentAPIKey(agentID, apikeyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegenerateAgentAPIKey", reflect.TypeOf((*MockGenAIService)(nil).RegenerateAgentAPIKey), agentID, apikeyID)
}

// UpdateAgent mocks base method.
func (m *MockGenAIService) UpdateAgent(agentID string, req *godo.AgentUpdateRequest) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgent", agentID, req)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgent indicates an expected call of UpdateAgent.
func (mr *MockGenAIServiceMockRecorder) UpdateAgent(agentID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgent", reflect.TypeOf((*MockGenAIService)(nil).UpdateAgent), agentID, req)
}

// UpdateAgentRoute mocks base method.
func (m *MockGenAIService) UpdateAgentRoute(parentAgentID, childAgentID string, req *godo.AgentRouteUpdateRequest) (*do.AgentRouteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgentRoute", parentAgentID, childAgentID, req)
	ret0, _ := ret[0].(*do.AgentRouteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgentRoute indicates an expected call of UpdateAgentRoute.
func (mr *MockGenAIServiceMockRecorder) UpdateAgentRoute(parentAgentID, childAgentID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgentRoute", reflect.TypeOf((*MockGenAIService)(nil).UpdateAgentRoute), parentAgentID, childAgentID, req)
}

// UpdateAgentAPIKey mocks base method.
func (m *MockGenAIService) UpdateAgentAPIKey(agentID, apikeyID string, req *godo.AgentAPIKeyUpdateRequest) (*do.ApiKeyInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgentAPIKey", agentID, apikeyID, req)
	ret0, _ := ret[0].(*do.ApiKeyInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgentAPIKey indicates an expected call of UpdateAgentAPIKey.
func (mr *MockGenAIServiceMockRecorder) UpdateAgentAPIKey(agentID, apikeyID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgentAPIKey", reflect.TypeOf((*MockGenAIService)(nil).UpdateAgentAPIKey), agentID, apikeyID, req)
}

// UpdateAgentVisibility mocks base method.
func (m *MockGenAIService) UpdateAgentVisibility(agentID string, req *godo.AgentVisibilityUpdateRequest) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgentVisibility", agentID, req)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgentVisibility indicates an expected call of UpdateAgentVisibility.
func (mr *MockGenAIServiceMockRecorder) UpdateAgentVisibility(agentID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgentVisibility", reflect.TypeOf((*MockGenAIService)(nil).UpdateAgentVisibility), agentID, req)
}

// UpdateFunctionRoute mocks base method.
func (m *MockGenAIService) UpdateFunctionRoute(agent_id, function_id string, req *godo.FunctionRouteUpdateRequest) (*do.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFunctionRoute", agent_id, function_id, req)
	ret0, _ := ret[0].(*do.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFunctionRoute indicates an expected call of UpdateFunctionRoute.
func (mr *MockGenAIServiceMockRecorder) UpdateFunctionRoute(agent_id, function_id, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFunctionRoute", reflect.TypeOf((*MockGenAIService)(nil).UpdateFunctionRoute), agent_id, function_id, req)
}

// UpdateKnowledgeBase mocks base method.
func (m *MockGenAIService) UpdateKnowledgeBase(knowledgeBaseID string, req *godo.UpdateKnowledgeBaseRequest) (*do.KnowledgeBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKnowledgeBase", knowledgeBaseID, req)
	ret0, _ := ret[0].(*do.KnowledgeBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateKnowledgeBase indicates an expected call of UpdateKnowledgeBase.
func (mr *MockGenAIServiceMockRecorder) UpdateKnowledgeBase(knowledgeBaseID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKnowledgeBase", reflect.TypeOf((*MockGenAIService)(nil).UpdateKnowledgeBase), knowledgeBaseID, req)
}
