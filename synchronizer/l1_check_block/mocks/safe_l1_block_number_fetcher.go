// Code generated by mockery. DO NOT EDIT.

package mock_l1_check_block

import (
	context "context"

	l1_check_block "github.com/0xPolygonHermez/zkevm-node/synchronizer/l1_check_block"
	mock "github.com/stretchr/testify/mock"
)

// SafeL1BlockNumberFetcher is an autogenerated mock type for the SafeL1BlockNumberFetcher type
type SafeL1BlockNumberFetcher struct {
	mock.Mock
}

type SafeL1BlockNumberFetcher_Expecter struct {
	mock *mock.Mock
}

func (_m *SafeL1BlockNumberFetcher) EXPECT() *SafeL1BlockNumberFetcher_Expecter {
	return &SafeL1BlockNumberFetcher_Expecter{mock: &_m.Mock}
}

// Description provides a mock function with given fields:
func (_m *SafeL1BlockNumberFetcher) Description() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SafeL1BlockNumberFetcher_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type SafeL1BlockNumberFetcher_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *SafeL1BlockNumberFetcher_Expecter) Description() *SafeL1BlockNumberFetcher_Description_Call {
	return &SafeL1BlockNumberFetcher_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *SafeL1BlockNumberFetcher_Description_Call) Run(run func()) *SafeL1BlockNumberFetcher_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SafeL1BlockNumberFetcher_Description_Call) Return(_a0 string) *SafeL1BlockNumberFetcher_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SafeL1BlockNumberFetcher_Description_Call) RunAndReturn(run func() string) *SafeL1BlockNumberFetcher_Description_Call {
	_c.Call.Return(run)
	return _c
}

// GetSafeBlockNumber provides a mock function with given fields: ctx, l1Client
func (_m *SafeL1BlockNumberFetcher) GetSafeBlockNumber(ctx context.Context, l1Client l1_check_block.L1Requester) (uint64, error) {
	ret := _m.Called(ctx, l1Client)

	if len(ret) == 0 {
		panic("no return value specified for GetSafeBlockNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, l1_check_block.L1Requester) (uint64, error)); ok {
		return rf(ctx, l1Client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, l1_check_block.L1Requester) uint64); ok {
		r0 = rf(ctx, l1Client)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, l1_check_block.L1Requester) error); ok {
		r1 = rf(ctx, l1Client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSafeBlockNumber'
type SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call struct {
	*mock.Call
}

// GetSafeBlockNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - l1Client l1_check_block.L1Requester
func (_e *SafeL1BlockNumberFetcher_Expecter) GetSafeBlockNumber(ctx interface{}, l1Client interface{}) *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call {
	return &SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call{Call: _e.mock.On("GetSafeBlockNumber", ctx, l1Client)}
}

func (_c *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call) Run(run func(ctx context.Context, l1Client l1_check_block.L1Requester)) *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(l1_check_block.L1Requester))
	})
	return _c
}

func (_c *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call) Return(_a0 uint64, _a1 error) *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call) RunAndReturn(run func(context.Context, l1_check_block.L1Requester) (uint64, error)) *SafeL1BlockNumberFetcher_GetSafeBlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewSafeL1BlockNumberFetcher creates a new instance of SafeL1BlockNumberFetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSafeL1BlockNumberFetcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *SafeL1BlockNumberFetcher {
	mock := &SafeL1BlockNumberFetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
