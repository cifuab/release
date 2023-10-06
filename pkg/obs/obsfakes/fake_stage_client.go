/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package obsfakes

import (
	"sync"
)

type FakeStageClient struct {
	CheckPrerequisitesStub        func() error
	checkPrerequisitesMutex       sync.RWMutex
	checkPrerequisitesArgsForCall []struct {
	}
	checkPrerequisitesReturns struct {
		result1 error
	}
	checkPrerequisitesReturnsOnCall map[int]struct {
		result1 error
	}
	CheckReleaseBranchStateStub        func() error
	checkReleaseBranchStateMutex       sync.RWMutex
	checkReleaseBranchStateArgsForCall []struct {
	}
	checkReleaseBranchStateReturns struct {
		result1 error
	}
	checkReleaseBranchStateReturnsOnCall map[int]struct {
		result1 error
	}
	CheckoutOBSProjectStub        func() error
	checkoutOBSProjectMutex       sync.RWMutex
	checkoutOBSProjectArgsForCall []struct {
	}
	checkoutOBSProjectReturns struct {
		result1 error
	}
	checkoutOBSProjectReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateOBSProjectStub        func() error
	generateOBSProjectMutex       sync.RWMutex
	generateOBSProjectArgsForCall []struct {
	}
	generateOBSProjectReturns struct {
		result1 error
	}
	generateOBSProjectReturnsOnCall map[int]struct {
		result1 error
	}
	GeneratePackageArtifactsStub        func() error
	generatePackageArtifactsMutex       sync.RWMutex
	generatePackageArtifactsArgsForCall []struct {
	}
	generatePackageArtifactsReturns struct {
		result1 error
	}
	generatePackageArtifactsReturnsOnCall map[int]struct {
		result1 error
	}
	GeneratePackageVersionStub        func()
	generatePackageVersionMutex       sync.RWMutex
	generatePackageVersionArgsForCall []struct {
	}
	GenerateReleaseVersionStub        func() error
	generateReleaseVersionMutex       sync.RWMutex
	generateReleaseVersionArgsForCall []struct {
	}
	generateReleaseVersionReturns struct {
		result1 error
	}
	generateReleaseVersionReturnsOnCall map[int]struct {
		result1 error
	}
	InitOBSRootStub        func() error
	initOBSRootMutex       sync.RWMutex
	initOBSRootArgsForCall []struct {
	}
	initOBSRootReturns struct {
		result1 error
	}
	initOBSRootReturnsOnCall map[int]struct {
		result1 error
	}
	InitStateStub        func()
	initStateMutex       sync.RWMutex
	initStateArgsForCall []struct {
	}
	PushStub        func() error
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
	}
	pushReturns struct {
		result1 error
	}
	pushReturnsOnCall map[int]struct {
		result1 error
	}
	SubmitStub        func(bool) error
	submitMutex       sync.RWMutex
	submitArgsForCall []struct {
		arg1 bool
	}
	submitReturns struct {
		result1 error
	}
	submitReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateOptionsStub        func() error
	validateOptionsMutex       sync.RWMutex
	validateOptionsArgsForCall []struct {
	}
	validateOptionsReturns struct {
		result1 error
	}
	validateOptionsReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStageClient) CheckPrerequisites() error {
	fake.checkPrerequisitesMutex.Lock()
	ret, specificReturn := fake.checkPrerequisitesReturnsOnCall[len(fake.checkPrerequisitesArgsForCall)]
	fake.checkPrerequisitesArgsForCall = append(fake.checkPrerequisitesArgsForCall, struct {
	}{})
	stub := fake.CheckPrerequisitesStub
	fakeReturns := fake.checkPrerequisitesReturns
	fake.recordInvocation("CheckPrerequisites", []interface{}{})
	fake.checkPrerequisitesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) CheckPrerequisitesCallCount() int {
	fake.checkPrerequisitesMutex.RLock()
	defer fake.checkPrerequisitesMutex.RUnlock()
	return len(fake.checkPrerequisitesArgsForCall)
}

func (fake *FakeStageClient) CheckPrerequisitesCalls(stub func() error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = stub
}

func (fake *FakeStageClient) CheckPrerequisitesReturns(result1 error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = nil
	fake.checkPrerequisitesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckPrerequisitesReturnsOnCall(i int, result1 error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = nil
	if fake.checkPrerequisitesReturnsOnCall == nil {
		fake.checkPrerequisitesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkPrerequisitesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckReleaseBranchState() error {
	fake.checkReleaseBranchStateMutex.Lock()
	ret, specificReturn := fake.checkReleaseBranchStateReturnsOnCall[len(fake.checkReleaseBranchStateArgsForCall)]
	fake.checkReleaseBranchStateArgsForCall = append(fake.checkReleaseBranchStateArgsForCall, struct {
	}{})
	stub := fake.CheckReleaseBranchStateStub
	fakeReturns := fake.checkReleaseBranchStateReturns
	fake.recordInvocation("CheckReleaseBranchState", []interface{}{})
	fake.checkReleaseBranchStateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) CheckReleaseBranchStateCallCount() int {
	fake.checkReleaseBranchStateMutex.RLock()
	defer fake.checkReleaseBranchStateMutex.RUnlock()
	return len(fake.checkReleaseBranchStateArgsForCall)
}

func (fake *FakeStageClient) CheckReleaseBranchStateCalls(stub func() error) {
	fake.checkReleaseBranchStateMutex.Lock()
	defer fake.checkReleaseBranchStateMutex.Unlock()
	fake.CheckReleaseBranchStateStub = stub
}

func (fake *FakeStageClient) CheckReleaseBranchStateReturns(result1 error) {
	fake.checkReleaseBranchStateMutex.Lock()
	defer fake.checkReleaseBranchStateMutex.Unlock()
	fake.CheckReleaseBranchStateStub = nil
	fake.checkReleaseBranchStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckReleaseBranchStateReturnsOnCall(i int, result1 error) {
	fake.checkReleaseBranchStateMutex.Lock()
	defer fake.checkReleaseBranchStateMutex.Unlock()
	fake.CheckReleaseBranchStateStub = nil
	if fake.checkReleaseBranchStateReturnsOnCall == nil {
		fake.checkReleaseBranchStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkReleaseBranchStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckoutOBSProject() error {
	fake.checkoutOBSProjectMutex.Lock()
	ret, specificReturn := fake.checkoutOBSProjectReturnsOnCall[len(fake.checkoutOBSProjectArgsForCall)]
	fake.checkoutOBSProjectArgsForCall = append(fake.checkoutOBSProjectArgsForCall, struct {
	}{})
	stub := fake.CheckoutOBSProjectStub
	fakeReturns := fake.checkoutOBSProjectReturns
	fake.recordInvocation("CheckoutOBSProject", []interface{}{})
	fake.checkoutOBSProjectMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) CheckoutOBSProjectCallCount() int {
	fake.checkoutOBSProjectMutex.RLock()
	defer fake.checkoutOBSProjectMutex.RUnlock()
	return len(fake.checkoutOBSProjectArgsForCall)
}

func (fake *FakeStageClient) CheckoutOBSProjectCalls(stub func() error) {
	fake.checkoutOBSProjectMutex.Lock()
	defer fake.checkoutOBSProjectMutex.Unlock()
	fake.CheckoutOBSProjectStub = stub
}

func (fake *FakeStageClient) CheckoutOBSProjectReturns(result1 error) {
	fake.checkoutOBSProjectMutex.Lock()
	defer fake.checkoutOBSProjectMutex.Unlock()
	fake.CheckoutOBSProjectStub = nil
	fake.checkoutOBSProjectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckoutOBSProjectReturnsOnCall(i int, result1 error) {
	fake.checkoutOBSProjectMutex.Lock()
	defer fake.checkoutOBSProjectMutex.Unlock()
	fake.CheckoutOBSProjectStub = nil
	if fake.checkoutOBSProjectReturnsOnCall == nil {
		fake.checkoutOBSProjectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkoutOBSProjectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GenerateOBSProject() error {
	fake.generateOBSProjectMutex.Lock()
	ret, specificReturn := fake.generateOBSProjectReturnsOnCall[len(fake.generateOBSProjectArgsForCall)]
	fake.generateOBSProjectArgsForCall = append(fake.generateOBSProjectArgsForCall, struct {
	}{})
	stub := fake.GenerateOBSProjectStub
	fakeReturns := fake.generateOBSProjectReturns
	fake.recordInvocation("GenerateOBSProject", []interface{}{})
	fake.generateOBSProjectMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) GenerateOBSProjectCallCount() int {
	fake.generateOBSProjectMutex.RLock()
	defer fake.generateOBSProjectMutex.RUnlock()
	return len(fake.generateOBSProjectArgsForCall)
}

func (fake *FakeStageClient) GenerateOBSProjectCalls(stub func() error) {
	fake.generateOBSProjectMutex.Lock()
	defer fake.generateOBSProjectMutex.Unlock()
	fake.GenerateOBSProjectStub = stub
}

func (fake *FakeStageClient) GenerateOBSProjectReturns(result1 error) {
	fake.generateOBSProjectMutex.Lock()
	defer fake.generateOBSProjectMutex.Unlock()
	fake.GenerateOBSProjectStub = nil
	fake.generateOBSProjectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GenerateOBSProjectReturnsOnCall(i int, result1 error) {
	fake.generateOBSProjectMutex.Lock()
	defer fake.generateOBSProjectMutex.Unlock()
	fake.GenerateOBSProjectStub = nil
	if fake.generateOBSProjectReturnsOnCall == nil {
		fake.generateOBSProjectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateOBSProjectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GeneratePackageArtifacts() error {
	fake.generatePackageArtifactsMutex.Lock()
	ret, specificReturn := fake.generatePackageArtifactsReturnsOnCall[len(fake.generatePackageArtifactsArgsForCall)]
	fake.generatePackageArtifactsArgsForCall = append(fake.generatePackageArtifactsArgsForCall, struct {
	}{})
	stub := fake.GeneratePackageArtifactsStub
	fakeReturns := fake.generatePackageArtifactsReturns
	fake.recordInvocation("GeneratePackageArtifacts", []interface{}{})
	fake.generatePackageArtifactsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) GeneratePackageArtifactsCallCount() int {
	fake.generatePackageArtifactsMutex.RLock()
	defer fake.generatePackageArtifactsMutex.RUnlock()
	return len(fake.generatePackageArtifactsArgsForCall)
}

func (fake *FakeStageClient) GeneratePackageArtifactsCalls(stub func() error) {
	fake.generatePackageArtifactsMutex.Lock()
	defer fake.generatePackageArtifactsMutex.Unlock()
	fake.GeneratePackageArtifactsStub = stub
}

func (fake *FakeStageClient) GeneratePackageArtifactsReturns(result1 error) {
	fake.generatePackageArtifactsMutex.Lock()
	defer fake.generatePackageArtifactsMutex.Unlock()
	fake.GeneratePackageArtifactsStub = nil
	fake.generatePackageArtifactsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GeneratePackageArtifactsReturnsOnCall(i int, result1 error) {
	fake.generatePackageArtifactsMutex.Lock()
	defer fake.generatePackageArtifactsMutex.Unlock()
	fake.GeneratePackageArtifactsStub = nil
	if fake.generatePackageArtifactsReturnsOnCall == nil {
		fake.generatePackageArtifactsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generatePackageArtifactsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GeneratePackageVersion() {
	fake.generatePackageVersionMutex.Lock()
	fake.generatePackageVersionArgsForCall = append(fake.generatePackageVersionArgsForCall, struct {
	}{})
	stub := fake.GeneratePackageVersionStub
	fake.recordInvocation("GeneratePackageVersion", []interface{}{})
	fake.generatePackageVersionMutex.Unlock()
	if stub != nil {
		fake.GeneratePackageVersionStub()
	}
}

func (fake *FakeStageClient) GeneratePackageVersionCallCount() int {
	fake.generatePackageVersionMutex.RLock()
	defer fake.generatePackageVersionMutex.RUnlock()
	return len(fake.generatePackageVersionArgsForCall)
}

func (fake *FakeStageClient) GeneratePackageVersionCalls(stub func()) {
	fake.generatePackageVersionMutex.Lock()
	defer fake.generatePackageVersionMutex.Unlock()
	fake.GeneratePackageVersionStub = stub
}

func (fake *FakeStageClient) GenerateReleaseVersion() error {
	fake.generateReleaseVersionMutex.Lock()
	ret, specificReturn := fake.generateReleaseVersionReturnsOnCall[len(fake.generateReleaseVersionArgsForCall)]
	fake.generateReleaseVersionArgsForCall = append(fake.generateReleaseVersionArgsForCall, struct {
	}{})
	stub := fake.GenerateReleaseVersionStub
	fakeReturns := fake.generateReleaseVersionReturns
	fake.recordInvocation("GenerateReleaseVersion", []interface{}{})
	fake.generateReleaseVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) GenerateReleaseVersionCallCount() int {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	return len(fake.generateReleaseVersionArgsForCall)
}

func (fake *FakeStageClient) GenerateReleaseVersionCalls(stub func() error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = stub
}

func (fake *FakeStageClient) GenerateReleaseVersionReturns(result1 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	fake.generateReleaseVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GenerateReleaseVersionReturnsOnCall(i int, result1 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	if fake.generateReleaseVersionReturnsOnCall == nil {
		fake.generateReleaseVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateReleaseVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) InitOBSRoot() error {
	fake.initOBSRootMutex.Lock()
	ret, specificReturn := fake.initOBSRootReturnsOnCall[len(fake.initOBSRootArgsForCall)]
	fake.initOBSRootArgsForCall = append(fake.initOBSRootArgsForCall, struct {
	}{})
	stub := fake.InitOBSRootStub
	fakeReturns := fake.initOBSRootReturns
	fake.recordInvocation("InitOBSRoot", []interface{}{})
	fake.initOBSRootMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) InitOBSRootCallCount() int {
	fake.initOBSRootMutex.RLock()
	defer fake.initOBSRootMutex.RUnlock()
	return len(fake.initOBSRootArgsForCall)
}

func (fake *FakeStageClient) InitOBSRootCalls(stub func() error) {
	fake.initOBSRootMutex.Lock()
	defer fake.initOBSRootMutex.Unlock()
	fake.InitOBSRootStub = stub
}

func (fake *FakeStageClient) InitOBSRootReturns(result1 error) {
	fake.initOBSRootMutex.Lock()
	defer fake.initOBSRootMutex.Unlock()
	fake.InitOBSRootStub = nil
	fake.initOBSRootReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) InitOBSRootReturnsOnCall(i int, result1 error) {
	fake.initOBSRootMutex.Lock()
	defer fake.initOBSRootMutex.Unlock()
	fake.InitOBSRootStub = nil
	if fake.initOBSRootReturnsOnCall == nil {
		fake.initOBSRootReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initOBSRootReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) InitState() {
	fake.initStateMutex.Lock()
	fake.initStateArgsForCall = append(fake.initStateArgsForCall, struct {
	}{})
	stub := fake.InitStateStub
	fake.recordInvocation("InitState", []interface{}{})
	fake.initStateMutex.Unlock()
	if stub != nil {
		fake.InitStateStub()
	}
}

func (fake *FakeStageClient) InitStateCallCount() int {
	fake.initStateMutex.RLock()
	defer fake.initStateMutex.RUnlock()
	return len(fake.initStateArgsForCall)
}

func (fake *FakeStageClient) InitStateCalls(stub func()) {
	fake.initStateMutex.Lock()
	defer fake.initStateMutex.Unlock()
	fake.InitStateStub = stub
}

func (fake *FakeStageClient) Push() error {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
	}{})
	stub := fake.PushStub
	fakeReturns := fake.pushReturns
	fake.recordInvocation("Push", []interface{}{})
	fake.pushMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeStageClient) PushCalls(stub func() error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = stub
}

func (fake *FakeStageClient) PushReturns(result1 error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) PushReturnsOnCall(i int, result1 error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) Submit(arg1 bool) error {
	fake.submitMutex.Lock()
	ret, specificReturn := fake.submitReturnsOnCall[len(fake.submitArgsForCall)]
	fake.submitArgsForCall = append(fake.submitArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SubmitStub
	fakeReturns := fake.submitReturns
	fake.recordInvocation("Submit", []interface{}{arg1})
	fake.submitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) SubmitCallCount() int {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	return len(fake.submitArgsForCall)
}

func (fake *FakeStageClient) SubmitCalls(stub func(bool) error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = stub
}

func (fake *FakeStageClient) SubmitArgsForCall(i int) bool {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	argsForCall := fake.submitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStageClient) SubmitReturns(result1 error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = nil
	fake.submitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) SubmitReturnsOnCall(i int, result1 error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = nil
	if fake.submitReturnsOnCall == nil {
		fake.submitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.submitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) ValidateOptions() error {
	fake.validateOptionsMutex.Lock()
	ret, specificReturn := fake.validateOptionsReturnsOnCall[len(fake.validateOptionsArgsForCall)]
	fake.validateOptionsArgsForCall = append(fake.validateOptionsArgsForCall, struct {
	}{})
	stub := fake.ValidateOptionsStub
	fakeReturns := fake.validateOptionsReturns
	fake.recordInvocation("ValidateOptions", []interface{}{})
	fake.validateOptionsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) ValidateOptionsCallCount() int {
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	return len(fake.validateOptionsArgsForCall)
}

func (fake *FakeStageClient) ValidateOptionsCalls(stub func() error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = stub
}

func (fake *FakeStageClient) ValidateOptionsReturns(result1 error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = nil
	fake.validateOptionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) ValidateOptionsReturnsOnCall(i int, result1 error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = nil
	if fake.validateOptionsReturnsOnCall == nil {
		fake.validateOptionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateOptionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) Wait() error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	stub := fake.WaitStub
	fakeReturns := fake.waitReturns
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeStageClient) WaitCalls(stub func() error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeStageClient) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkPrerequisitesMutex.RLock()
	defer fake.checkPrerequisitesMutex.RUnlock()
	fake.checkReleaseBranchStateMutex.RLock()
	defer fake.checkReleaseBranchStateMutex.RUnlock()
	fake.checkoutOBSProjectMutex.RLock()
	defer fake.checkoutOBSProjectMutex.RUnlock()
	fake.generateOBSProjectMutex.RLock()
	defer fake.generateOBSProjectMutex.RUnlock()
	fake.generatePackageArtifactsMutex.RLock()
	defer fake.generatePackageArtifactsMutex.RUnlock()
	fake.generatePackageVersionMutex.RLock()
	defer fake.generatePackageVersionMutex.RUnlock()
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	fake.initOBSRootMutex.RLock()
	defer fake.initOBSRootMutex.RUnlock()
	fake.initStateMutex.RLock()
	defer fake.initStateMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStageClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
