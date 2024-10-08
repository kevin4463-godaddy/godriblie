// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	mock "github.com/stretchr/testify/mock"
)

// DynamoDbProvider is an autogenerated mock type for the DynamoDbProvider type
type DynamoDbProvider struct {
	mock.Mock
}

type DynamoDbProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *DynamoDbProvider) EXPECT() *DynamoDbProvider_Expecter {
	return &DynamoDbProvider_Expecter{mock: &_m.Mock}
}

// CreateTable provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTable")
	}

	var r0 *dynamodb.CreateTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) *dynamodb.CreateTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.CreateTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_CreateTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTable'
type DynamoDbProvider_CreateTable_Call struct {
	*mock.Call
}

// CreateTable is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.CreateTableInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) CreateTable(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_CreateTable_Call {
	return &DynamoDbProvider_CreateTable_Call{Call: _e.mock.On("CreateTable",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_CreateTable_Call) Run(run func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_CreateTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.CreateTableInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_CreateTable_Call) Return(_a0 *dynamodb.CreateTableOutput, _a1 error) *DynamoDbProvider_CreateTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_CreateTable_Call) RunAndReturn(run func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)) *DynamoDbProvider_CreateTable_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteItem")
	}

	var r0 *dynamodb.DeleteItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) *dynamodb.DeleteItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.DeleteItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_DeleteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItem'
type DynamoDbProvider_DeleteItem_Call struct {
	*mock.Call
}

// DeleteItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.DeleteItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) DeleteItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_DeleteItem_Call {
	return &DynamoDbProvider_DeleteItem_Call{Call: _e.mock.On("DeleteItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_DeleteItem_Call) Run(run func(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_DeleteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.DeleteItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_DeleteItem_Call) Return(_a0 *dynamodb.DeleteItemOutput, _a1 error) *DynamoDbProvider_DeleteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_DeleteItem_Call) RunAndReturn(run func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)) *DynamoDbProvider_DeleteItem_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteStatement provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteStatement")
	}

	var r0 *dynamodb.ExecuteStatementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.ExecuteStatementInput, ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.ExecuteStatementInput, ...func(*dynamodb.Options)) *dynamodb.ExecuteStatementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.ExecuteStatementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.ExecuteStatementInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_ExecuteStatement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteStatement'
type DynamoDbProvider_ExecuteStatement_Call struct {
	*mock.Call
}

// ExecuteStatement is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.ExecuteStatementInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) ExecuteStatement(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_ExecuteStatement_Call {
	return &DynamoDbProvider_ExecuteStatement_Call{Call: _e.mock.On("ExecuteStatement",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_ExecuteStatement_Call) Run(run func(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_ExecuteStatement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.ExecuteStatementInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_ExecuteStatement_Call) Return(_a0 *dynamodb.ExecuteStatementOutput, _a1 error) *DynamoDbProvider_ExecuteStatement_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_ExecuteStatement_Call) RunAndReturn(run func(context.Context, *dynamodb.ExecuteStatementInput, ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)) *DynamoDbProvider_ExecuteStatement_Call {
	_c.Call.Return(run)
	return _c
}

// GetItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 *dynamodb.GetItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) *dynamodb.GetItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.GetItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_GetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItem'
type DynamoDbProvider_GetItem_Call struct {
	*mock.Call
}

// GetItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.GetItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) GetItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_GetItem_Call {
	return &DynamoDbProvider_GetItem_Call{Call: _e.mock.On("GetItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_GetItem_Call) Run(run func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_GetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.GetItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_GetItem_Call) Return(_a0 *dynamodb.GetItemOutput, _a1 error) *DynamoDbProvider_GetItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_GetItem_Call) RunAndReturn(run func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)) *DynamoDbProvider_GetItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueryPaginator provides a mock function with given fields: client, params, optFns
func (_m *DynamoDbProvider) NewQueryPaginator(client dynamodb.Client, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) *dynamodb.QueryPaginator {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, client, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for NewQueryPaginator")
	}

	var r0 *dynamodb.QueryPaginator
	if rf, ok := ret.Get(0).(func(dynamodb.Client, *dynamodb.QueryInput, ...func(*dynamodb.Options)) *dynamodb.QueryPaginator); ok {
		r0 = rf(client, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.QueryPaginator)
		}
	}

	return r0
}

// DynamoDbProvider_NewQueryPaginator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewQueryPaginator'
type DynamoDbProvider_NewQueryPaginator_Call struct {
	*mock.Call
}

// NewQueryPaginator is a helper method to define mock.On call
//   - client dynamodb.Client
//   - params *dynamodb.QueryInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) NewQueryPaginator(client interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_NewQueryPaginator_Call {
	return &DynamoDbProvider_NewQueryPaginator_Call{Call: _e.mock.On("NewQueryPaginator",
		append([]interface{}{client, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_NewQueryPaginator_Call) Run(run func(client dynamodb.Client, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_NewQueryPaginator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(dynamodb.Client), args[1].(*dynamodb.QueryInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_NewQueryPaginator_Call) Return(_a0 *dynamodb.QueryPaginator) *DynamoDbProvider_NewQueryPaginator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DynamoDbProvider_NewQueryPaginator_Call) RunAndReturn(run func(dynamodb.Client, *dynamodb.QueryInput, ...func(*dynamodb.Options)) *dynamodb.QueryPaginator) *DynamoDbProvider_NewQueryPaginator_Call {
	_c.Call.Return(run)
	return _c
}

// PutItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutItem")
	}

	var r0 *dynamodb.PutItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) *dynamodb.PutItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_PutItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItem'
type DynamoDbProvider_PutItem_Call struct {
	*mock.Call
}

// PutItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.PutItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) PutItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_PutItem_Call {
	return &DynamoDbProvider_PutItem_Call{Call: _e.mock.On("PutItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_PutItem_Call) Run(run func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_PutItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.PutItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_PutItem_Call) Return(_a0 *dynamodb.PutItemOutput, _a1 error) *DynamoDbProvider_PutItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_PutItem_Call) RunAndReturn(run func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)) *DynamoDbProvider_PutItem_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *dynamodb.QueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) *dynamodb.QueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.QueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type DynamoDbProvider_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.QueryInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) Query(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_Query_Call {
	return &DynamoDbProvider_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_Query_Call) Run(run func(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.QueryInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_Query_Call) Return(_a0 *dynamodb.QueryOutput, _a1 error) *DynamoDbProvider_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_Query_Call) RunAndReturn(run func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)) *DynamoDbProvider_Query_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamoDbProvider) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateItem")
	}

	var r0 *dynamodb.UpdateItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) *dynamodb.UpdateItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.UpdateItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamoDbProvider_UpdateItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItem'
type DynamoDbProvider_UpdateItem_Call struct {
	*mock.Call
}

// UpdateItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.UpdateItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamoDbProvider_Expecter) UpdateItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamoDbProvider_UpdateItem_Call {
	return &DynamoDbProvider_UpdateItem_Call{Call: _e.mock.On("UpdateItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamoDbProvider_UpdateItem_Call) Run(run func(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options))) *DynamoDbProvider_UpdateItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.UpdateItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamoDbProvider_UpdateItem_Call) Return(_a0 *dynamodb.UpdateItemOutput, _a1 error) *DynamoDbProvider_UpdateItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamoDbProvider_UpdateItem_Call) RunAndReturn(run func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)) *DynamoDbProvider_UpdateItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewDynamoDbProvider creates a new instance of DynamoDbProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDynamoDbProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *DynamoDbProvider {
	mock := &DynamoDbProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
