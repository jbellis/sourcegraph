// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package store

import (
	"context"
	"sync"

	types "github.com/sourcegraph/sourcegraph/enterprise/internal/github_apps/types"
	encryption "github.com/sourcegraph/sourcegraph/internal/encryption"
)

// MockGitHubAppsStore is a mock implementation of the GitHubAppsStore
// interface (from the package
// github.com/sourcegraph/sourcegraph/enterprise/internal/github_apps/store)
// used for unit testing.
type MockGitHubAppsStore struct {
	// CreateFunc is an instance of a mock function object controlling the
	// behavior of the method Create.
	CreateFunc *GitHubAppsStoreCreateFunc
	// DeleteFunc is an instance of a mock function object controlling the
	// behavior of the method Delete.
	DeleteFunc *GitHubAppsStoreDeleteFunc
	// GetByAppIDFunc is an instance of a mock function object controlling
	// the behavior of the method GetByAppID.
	GetByAppIDFunc *GitHubAppsStoreGetByAppIDFunc
	// GetByIDFunc is an instance of a mock function object controlling the
	// behavior of the method GetByID.
	GetByIDFunc *GitHubAppsStoreGetByIDFunc
	// GetBySlugFunc is an instance of a mock function object controlling
	// the behavior of the method GetBySlug.
	GetBySlugFunc *GitHubAppsStoreGetBySlugFunc
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *GitHubAppsStoreListFunc
	// UpdateFunc is an instance of a mock function object controlling the
	// behavior of the method Update.
	UpdateFunc *GitHubAppsStoreUpdateFunc
	// WithEncryptionKeyFunc is an instance of a mock function object
	// controlling the behavior of the method WithEncryptionKey.
	WithEncryptionKeyFunc *GitHubAppsStoreWithEncryptionKeyFunc
}

// NewMockGitHubAppsStore creates a new mock of the GitHubAppsStore
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockGitHubAppsStore() *MockGitHubAppsStore {
	return &MockGitHubAppsStore{
		CreateFunc: &GitHubAppsStoreCreateFunc{
			defaultHook: func(context.Context, *types.GitHubApp) (r0 int, r1 error) {
				return
			},
		},
		DeleteFunc: &GitHubAppsStoreDeleteFunc{
			defaultHook: func(context.Context, int) (r0 error) {
				return
			},
		},
		GetByAppIDFunc: &GitHubAppsStoreGetByAppIDFunc{
			defaultHook: func(context.Context, int, string) (r0 *types.GitHubApp, r1 error) {
				return
			},
		},
		GetByIDFunc: &GitHubAppsStoreGetByIDFunc{
			defaultHook: func(context.Context, int) (r0 *types.GitHubApp, r1 error) {
				return
			},
		},
		GetBySlugFunc: &GitHubAppsStoreGetBySlugFunc{
			defaultHook: func(context.Context, string, string) (r0 *types.GitHubApp, r1 error) {
				return
			},
		},
		ListFunc: &GitHubAppsStoreListFunc{
			defaultHook: func(context.Context) (r0 []*types.GitHubApp, r1 error) {
				return
			},
		},
		UpdateFunc: &GitHubAppsStoreUpdateFunc{
			defaultHook: func(context.Context, int, *types.GitHubApp) (r0 *types.GitHubApp, r1 error) {
				return
			},
		},
		WithEncryptionKeyFunc: &GitHubAppsStoreWithEncryptionKeyFunc{
			defaultHook: func(encryption.Key) (r0 GitHubAppsStore) {
				return
			},
		},
	}
}

// NewStrictMockGitHubAppsStore creates a new mock of the GitHubAppsStore
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockGitHubAppsStore() *MockGitHubAppsStore {
	return &MockGitHubAppsStore{
		CreateFunc: &GitHubAppsStoreCreateFunc{
			defaultHook: func(context.Context, *types.GitHubApp) (int, error) {
				panic("unexpected invocation of MockGitHubAppsStore.Create")
			},
		},
		DeleteFunc: &GitHubAppsStoreDeleteFunc{
			defaultHook: func(context.Context, int) error {
				panic("unexpected invocation of MockGitHubAppsStore.Delete")
			},
		},
		GetByAppIDFunc: &GitHubAppsStoreGetByAppIDFunc{
			defaultHook: func(context.Context, int, string) (*types.GitHubApp, error) {
				panic("unexpected invocation of MockGitHubAppsStore.GetByAppID")
			},
		},
		GetByIDFunc: &GitHubAppsStoreGetByIDFunc{
			defaultHook: func(context.Context, int) (*types.GitHubApp, error) {
				panic("unexpected invocation of MockGitHubAppsStore.GetByID")
			},
		},
		GetBySlugFunc: &GitHubAppsStoreGetBySlugFunc{
			defaultHook: func(context.Context, string, string) (*types.GitHubApp, error) {
				panic("unexpected invocation of MockGitHubAppsStore.GetBySlug")
			},
		},
		ListFunc: &GitHubAppsStoreListFunc{
			defaultHook: func(context.Context) ([]*types.GitHubApp, error) {
				panic("unexpected invocation of MockGitHubAppsStore.List")
			},
		},
		UpdateFunc: &GitHubAppsStoreUpdateFunc{
			defaultHook: func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error) {
				panic("unexpected invocation of MockGitHubAppsStore.Update")
			},
		},
		WithEncryptionKeyFunc: &GitHubAppsStoreWithEncryptionKeyFunc{
			defaultHook: func(encryption.Key) GitHubAppsStore {
				panic("unexpected invocation of MockGitHubAppsStore.WithEncryptionKey")
			},
		},
	}
}

// NewMockGitHubAppsStoreFrom creates a new mock of the MockGitHubAppsStore
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockGitHubAppsStoreFrom(i GitHubAppsStore) *MockGitHubAppsStore {
	return &MockGitHubAppsStore{
		CreateFunc: &GitHubAppsStoreCreateFunc{
			defaultHook: i.Create,
		},
		DeleteFunc: &GitHubAppsStoreDeleteFunc{
			defaultHook: i.Delete,
		},
		GetByAppIDFunc: &GitHubAppsStoreGetByAppIDFunc{
			defaultHook: i.GetByAppID,
		},
		GetByIDFunc: &GitHubAppsStoreGetByIDFunc{
			defaultHook: i.GetByID,
		},
		GetBySlugFunc: &GitHubAppsStoreGetBySlugFunc{
			defaultHook: i.GetBySlug,
		},
		ListFunc: &GitHubAppsStoreListFunc{
			defaultHook: i.List,
		},
		UpdateFunc: &GitHubAppsStoreUpdateFunc{
			defaultHook: i.Update,
		},
		WithEncryptionKeyFunc: &GitHubAppsStoreWithEncryptionKeyFunc{
			defaultHook: i.WithEncryptionKey,
		},
	}
}

// GitHubAppsStoreCreateFunc describes the behavior when the Create method
// of the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreCreateFunc struct {
	defaultHook func(context.Context, *types.GitHubApp) (int, error)
	hooks       []func(context.Context, *types.GitHubApp) (int, error)
	history     []GitHubAppsStoreCreateFuncCall
	mutex       sync.Mutex
}

// Create delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitHubAppsStore) Create(v0 context.Context, v1 *types.GitHubApp) (int, error) {
	r0, r1 := m.CreateFunc.nextHook()(v0, v1)
	m.CreateFunc.appendCall(GitHubAppsStoreCreateFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Create method of the
// parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreCreateFunc) SetDefaultHook(hook func(context.Context, *types.GitHubApp) (int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Create method of the parent MockGitHubAppsStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitHubAppsStoreCreateFunc) PushHook(hook func(context.Context, *types.GitHubApp) (int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreCreateFunc) SetDefaultReturn(r0 int, r1 error) {
	f.SetDefaultHook(func(context.Context, *types.GitHubApp) (int, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreCreateFunc) PushReturn(r0 int, r1 error) {
	f.PushHook(func(context.Context, *types.GitHubApp) (int, error) {
		return r0, r1
	})
}

func (f *GitHubAppsStoreCreateFunc) nextHook() func(context.Context, *types.GitHubApp) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreCreateFunc) appendCall(r0 GitHubAppsStoreCreateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreCreateFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreCreateFunc) History() []GitHubAppsStoreCreateFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreCreateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreCreateFuncCall is an object that describes an invocation
// of method Create on an instance of MockGitHubAppsStore.
type GitHubAppsStoreCreateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *types.GitHubApp
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreCreateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreCreateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitHubAppsStoreDeleteFunc describes the behavior when the Delete method
// of the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreDeleteFunc struct {
	defaultHook func(context.Context, int) error
	hooks       []func(context.Context, int) error
	history     []GitHubAppsStoreDeleteFuncCall
	mutex       sync.Mutex
}

// Delete delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitHubAppsStore) Delete(v0 context.Context, v1 int) error {
	r0 := m.DeleteFunc.nextHook()(v0, v1)
	m.DeleteFunc.appendCall(GitHubAppsStoreDeleteFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Delete method of the
// parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreDeleteFunc) SetDefaultHook(hook func(context.Context, int) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Delete method of the parent MockGitHubAppsStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitHubAppsStoreDeleteFunc) PushHook(hook func(context.Context, int) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreDeleteFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreDeleteFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int) error {
		return r0
	})
}

func (f *GitHubAppsStoreDeleteFunc) nextHook() func(context.Context, int) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreDeleteFunc) appendCall(r0 GitHubAppsStoreDeleteFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreDeleteFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreDeleteFunc) History() []GitHubAppsStoreDeleteFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreDeleteFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreDeleteFuncCall is an object that describes an invocation
// of method Delete on an instance of MockGitHubAppsStore.
type GitHubAppsStoreDeleteFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreDeleteFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreDeleteFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// GitHubAppsStoreGetByAppIDFunc describes the behavior when the GetByAppID
// method of the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreGetByAppIDFunc struct {
	defaultHook func(context.Context, int, string) (*types.GitHubApp, error)
	hooks       []func(context.Context, int, string) (*types.GitHubApp, error)
	history     []GitHubAppsStoreGetByAppIDFuncCall
	mutex       sync.Mutex
}

// GetByAppID delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockGitHubAppsStore) GetByAppID(v0 context.Context, v1 int, v2 string) (*types.GitHubApp, error) {
	r0, r1 := m.GetByAppIDFunc.nextHook()(v0, v1, v2)
	m.GetByAppIDFunc.appendCall(GitHubAppsStoreGetByAppIDFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByAppID method of
// the parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreGetByAppIDFunc) SetDefaultHook(hook func(context.Context, int, string) (*types.GitHubApp, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByAppID method of the parent MockGitHubAppsStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitHubAppsStoreGetByAppIDFunc) PushHook(hook func(context.Context, int, string) (*types.GitHubApp, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreGetByAppIDFunc) SetDefaultReturn(r0 *types.GitHubApp, r1 error) {
	f.SetDefaultHook(func(context.Context, int, string) (*types.GitHubApp, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreGetByAppIDFunc) PushReturn(r0 *types.GitHubApp, r1 error) {
	f.PushHook(func(context.Context, int, string) (*types.GitHubApp, error) {
		return r0, r1
	})
}

func (f *GitHubAppsStoreGetByAppIDFunc) nextHook() func(context.Context, int, string) (*types.GitHubApp, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreGetByAppIDFunc) appendCall(r0 GitHubAppsStoreGetByAppIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreGetByAppIDFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreGetByAppIDFunc) History() []GitHubAppsStoreGetByAppIDFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreGetByAppIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreGetByAppIDFuncCall is an object that describes an
// invocation of method GetByAppID on an instance of MockGitHubAppsStore.
type GitHubAppsStoreGetByAppIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.GitHubApp
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreGetByAppIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreGetByAppIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitHubAppsStoreGetByIDFunc describes the behavior when the GetByID method
// of the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreGetByIDFunc struct {
	defaultHook func(context.Context, int) (*types.GitHubApp, error)
	hooks       []func(context.Context, int) (*types.GitHubApp, error)
	history     []GitHubAppsStoreGetByIDFuncCall
	mutex       sync.Mutex
}

// GetByID delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitHubAppsStore) GetByID(v0 context.Context, v1 int) (*types.GitHubApp, error) {
	r0, r1 := m.GetByIDFunc.nextHook()(v0, v1)
	m.GetByIDFunc.appendCall(GitHubAppsStoreGetByIDFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByID method of
// the parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreGetByIDFunc) SetDefaultHook(hook func(context.Context, int) (*types.GitHubApp, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByID method of the parent MockGitHubAppsStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitHubAppsStoreGetByIDFunc) PushHook(hook func(context.Context, int) (*types.GitHubApp, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreGetByIDFunc) SetDefaultReturn(r0 *types.GitHubApp, r1 error) {
	f.SetDefaultHook(func(context.Context, int) (*types.GitHubApp, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreGetByIDFunc) PushReturn(r0 *types.GitHubApp, r1 error) {
	f.PushHook(func(context.Context, int) (*types.GitHubApp, error) {
		return r0, r1
	})
}

func (f *GitHubAppsStoreGetByIDFunc) nextHook() func(context.Context, int) (*types.GitHubApp, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreGetByIDFunc) appendCall(r0 GitHubAppsStoreGetByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreGetByIDFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreGetByIDFunc) History() []GitHubAppsStoreGetByIDFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreGetByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreGetByIDFuncCall is an object that describes an invocation
// of method GetByID on an instance of MockGitHubAppsStore.
type GitHubAppsStoreGetByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.GitHubApp
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreGetByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreGetByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitHubAppsStoreGetBySlugFunc describes the behavior when the GetBySlug
// method of the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreGetBySlugFunc struct {
	defaultHook func(context.Context, string, string) (*types.GitHubApp, error)
	hooks       []func(context.Context, string, string) (*types.GitHubApp, error)
	history     []GitHubAppsStoreGetBySlugFuncCall
	mutex       sync.Mutex
}

// GetBySlug delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitHubAppsStore) GetBySlug(v0 context.Context, v1 string, v2 string) (*types.GitHubApp, error) {
	r0, r1 := m.GetBySlugFunc.nextHook()(v0, v1, v2)
	m.GetBySlugFunc.appendCall(GitHubAppsStoreGetBySlugFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetBySlug method of
// the parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreGetBySlugFunc) SetDefaultHook(hook func(context.Context, string, string) (*types.GitHubApp, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetBySlug method of the parent MockGitHubAppsStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitHubAppsStoreGetBySlugFunc) PushHook(hook func(context.Context, string, string) (*types.GitHubApp, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreGetBySlugFunc) SetDefaultReturn(r0 *types.GitHubApp, r1 error) {
	f.SetDefaultHook(func(context.Context, string, string) (*types.GitHubApp, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreGetBySlugFunc) PushReturn(r0 *types.GitHubApp, r1 error) {
	f.PushHook(func(context.Context, string, string) (*types.GitHubApp, error) {
		return r0, r1
	})
}

func (f *GitHubAppsStoreGetBySlugFunc) nextHook() func(context.Context, string, string) (*types.GitHubApp, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreGetBySlugFunc) appendCall(r0 GitHubAppsStoreGetBySlugFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreGetBySlugFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreGetBySlugFunc) History() []GitHubAppsStoreGetBySlugFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreGetBySlugFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreGetBySlugFuncCall is an object that describes an
// invocation of method GetBySlug on an instance of MockGitHubAppsStore.
type GitHubAppsStoreGetBySlugFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.GitHubApp
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreGetBySlugFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreGetBySlugFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitHubAppsStoreListFunc describes the behavior when the List method of
// the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreListFunc struct {
	defaultHook func(context.Context) ([]*types.GitHubApp, error)
	hooks       []func(context.Context) ([]*types.GitHubApp, error)
	history     []GitHubAppsStoreListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitHubAppsStore) List(v0 context.Context) ([]*types.GitHubApp, error) {
	r0, r1 := m.ListFunc.nextHook()(v0)
	m.ListFunc.appendCall(GitHubAppsStoreListFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreListFunc) SetDefaultHook(hook func(context.Context) ([]*types.GitHubApp, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockGitHubAppsStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitHubAppsStoreListFunc) PushHook(hook func(context.Context) ([]*types.GitHubApp, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreListFunc) SetDefaultReturn(r0 []*types.GitHubApp, r1 error) {
	f.SetDefaultHook(func(context.Context) ([]*types.GitHubApp, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreListFunc) PushReturn(r0 []*types.GitHubApp, r1 error) {
	f.PushHook(func(context.Context) ([]*types.GitHubApp, error) {
		return r0, r1
	})
}

func (f *GitHubAppsStoreListFunc) nextHook() func(context.Context) ([]*types.GitHubApp, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreListFunc) appendCall(r0 GitHubAppsStoreListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreListFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreListFunc) History() []GitHubAppsStoreListFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreListFuncCall is an object that describes an invocation of
// method List on an instance of MockGitHubAppsStore.
type GitHubAppsStoreListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []*types.GitHubApp
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitHubAppsStoreUpdateFunc describes the behavior when the Update method
// of the parent MockGitHubAppsStore instance is invoked.
type GitHubAppsStoreUpdateFunc struct {
	defaultHook func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error)
	hooks       []func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error)
	history     []GitHubAppsStoreUpdateFuncCall
	mutex       sync.Mutex
}

// Update delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockGitHubAppsStore) Update(v0 context.Context, v1 int, v2 *types.GitHubApp) (*types.GitHubApp, error) {
	r0, r1 := m.UpdateFunc.nextHook()(v0, v1, v2)
	m.UpdateFunc.appendCall(GitHubAppsStoreUpdateFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Update method of the
// parent MockGitHubAppsStore instance is invoked and the hook queue is
// empty.
func (f *GitHubAppsStoreUpdateFunc) SetDefaultHook(hook func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Update method of the parent MockGitHubAppsStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *GitHubAppsStoreUpdateFunc) PushHook(hook func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreUpdateFunc) SetDefaultReturn(r0 *types.GitHubApp, r1 error) {
	f.SetDefaultHook(func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreUpdateFunc) PushReturn(r0 *types.GitHubApp, r1 error) {
	f.PushHook(func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error) {
		return r0, r1
	})
}

func (f *GitHubAppsStoreUpdateFunc) nextHook() func(context.Context, int, *types.GitHubApp) (*types.GitHubApp, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreUpdateFunc) appendCall(r0 GitHubAppsStoreUpdateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreUpdateFuncCall objects
// describing the invocations of this function.
func (f *GitHubAppsStoreUpdateFunc) History() []GitHubAppsStoreUpdateFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreUpdateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreUpdateFuncCall is an object that describes an invocation
// of method Update on an instance of MockGitHubAppsStore.
type GitHubAppsStoreUpdateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 *types.GitHubApp
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.GitHubApp
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreUpdateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreUpdateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitHubAppsStoreWithEncryptionKeyFunc describes the behavior when the
// WithEncryptionKey method of the parent MockGitHubAppsStore instance is
// invoked.
type GitHubAppsStoreWithEncryptionKeyFunc struct {
	defaultHook func(encryption.Key) GitHubAppsStore
	hooks       []func(encryption.Key) GitHubAppsStore
	history     []GitHubAppsStoreWithEncryptionKeyFuncCall
	mutex       sync.Mutex
}

// WithEncryptionKey delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockGitHubAppsStore) WithEncryptionKey(v0 encryption.Key) GitHubAppsStore {
	r0 := m.WithEncryptionKeyFunc.nextHook()(v0)
	m.WithEncryptionKeyFunc.appendCall(GitHubAppsStoreWithEncryptionKeyFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the WithEncryptionKey
// method of the parent MockGitHubAppsStore instance is invoked and the hook
// queue is empty.
func (f *GitHubAppsStoreWithEncryptionKeyFunc) SetDefaultHook(hook func(encryption.Key) GitHubAppsStore) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// WithEncryptionKey method of the parent MockGitHubAppsStore instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *GitHubAppsStoreWithEncryptionKeyFunc) PushHook(hook func(encryption.Key) GitHubAppsStore) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitHubAppsStoreWithEncryptionKeyFunc) SetDefaultReturn(r0 GitHubAppsStore) {
	f.SetDefaultHook(func(encryption.Key) GitHubAppsStore {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitHubAppsStoreWithEncryptionKeyFunc) PushReturn(r0 GitHubAppsStore) {
	f.PushHook(func(encryption.Key) GitHubAppsStore {
		return r0
	})
}

func (f *GitHubAppsStoreWithEncryptionKeyFunc) nextHook() func(encryption.Key) GitHubAppsStore {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitHubAppsStoreWithEncryptionKeyFunc) appendCall(r0 GitHubAppsStoreWithEncryptionKeyFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitHubAppsStoreWithEncryptionKeyFuncCall
// objects describing the invocations of this function.
func (f *GitHubAppsStoreWithEncryptionKeyFunc) History() []GitHubAppsStoreWithEncryptionKeyFuncCall {
	f.mutex.Lock()
	history := make([]GitHubAppsStoreWithEncryptionKeyFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitHubAppsStoreWithEncryptionKeyFuncCall is an object that describes an
// invocation of method WithEncryptionKey on an instance of
// MockGitHubAppsStore.
type GitHubAppsStoreWithEncryptionKeyFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 encryption.Key
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 GitHubAppsStore
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitHubAppsStoreWithEncryptionKeyFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitHubAppsStoreWithEncryptionKeyFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
