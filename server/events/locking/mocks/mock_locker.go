// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/locking (interfaces: Locker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	locking "github.com/runatlantis/atlantis/server/events/locking"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockLocker() *MockLocker {
	return &MockLocker{fail: pegomock.GlobalFailHandler}
}

func (mock *MockLocker) TryLock(p models.Project, workspace string, pull models.PullRequest, user models.User) (locking.TryLockResponse, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{p, workspace, pull, user}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((*locking.TryLockResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 locking.TryLockResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(locking.TryLockResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) Unlock(key string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{key}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Unlock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) List() (map[string]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("List", params, []reflect.Type{reflect.TypeOf((*map[string]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) UnlockByPull(repoFullName string, pullNum int) ([]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{repoFullName, pullNum}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UnlockByPull", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) GetLock(key string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{key}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetLock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) VerifyWasCalledOnce() *VerifierLocker {
	return &VerifierLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockLocker) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierLocker {
	return &VerifierLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierLocker {
	return &VerifierLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierLocker {
	return &VerifierLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierLocker struct {
	mock                   *MockLocker
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierLocker) TryLock(p models.Project, workspace string, pull models.PullRequest, user models.User) *Locker_TryLock_OngoingVerification {
	params := []pegomock.Param{p, workspace, pull, user}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &Locker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Locker_TryLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *Locker_TryLock_OngoingVerification) GetCapturedArguments() (models.Project, string, models.PullRequest, models.User) {
	p, workspace, pull, user := c.GetAllCapturedArguments()
	return p[len(p)-1], workspace[len(workspace)-1], pull[len(pull)-1], user[len(user)-1]
}

func (c *Locker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Project, _param1 []string, _param2 []models.PullRequest, _param3 []models.User) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Project, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Project)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]models.PullRequest, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(models.PullRequest)
		}
		_param3 = make([]models.User, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(models.User)
		}
	}
	return
}

func (verifier *VerifierLocker) Unlock(key string) *Locker_Unlock_OngoingVerification {
	params := []pegomock.Param{key}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Unlock", params, verifier.timeout)
	return &Locker_Unlock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Locker_Unlock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *Locker_Unlock_OngoingVerification) GetCapturedArguments() string {
	key := c.GetAllCapturedArguments()
	return key[len(key)-1]
}

func (c *Locker_Unlock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierLocker) List() *Locker_List_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "List", params, verifier.timeout)
	return &Locker_List_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Locker_List_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *Locker_List_OngoingVerification) GetCapturedArguments() {
}

func (c *Locker_List_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierLocker) UnlockByPull(repoFullName string, pullNum int) *Locker_UnlockByPull_OngoingVerification {
	params := []pegomock.Param{repoFullName, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UnlockByPull", params, verifier.timeout)
	return &Locker_UnlockByPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Locker_UnlockByPull_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *Locker_UnlockByPull_OngoingVerification) GetCapturedArguments() (string, int) {
	repoFullName, pullNum := c.GetAllCapturedArguments()
	return repoFullName[len(repoFullName)-1], pullNum[len(pullNum)-1]
}

func (c *Locker_UnlockByPull_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierLocker) GetLock(key string) *Locker_GetLock_OngoingVerification {
	params := []pegomock.Param{key}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetLock", params, verifier.timeout)
	return &Locker_GetLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Locker_GetLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *Locker_GetLock_OngoingVerification) GetCapturedArguments() string {
	key := c.GetAllCapturedArguments()
	return key[len(key)-1]
}

func (c *Locker_GetLock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
