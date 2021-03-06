// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package processTransaction

import (
	"sync"
	"time"
	"transactions/model"
)

// Ensure, that EnvMock does implement Env.
// If this is not the case, regenerate this file with moq.
var _ Env = &EnvMock{}

// EnvMock is a mock implementation of Env.
//
// 	func TestSomethingThatUsesEnv(t *testing.T) {
//
// 		// make and configure a mocked Env
// 		mockedEnv := &EnvMock{
// 			GetUserByIDFunc: func(id int) (*model.User, error) {
// 				panic("mock out the GetUserByID method")
// 			},
// 			NowFunc: func() time.Time {
// 				panic("mock out the Now method")
// 			},
// 			UpdateTransactionFunc: func(data *model.Transaction) error {
// 				panic("mock out the UpdateTransaction method")
// 			},
// 			UpdateUserFunc: func(data *model.User) error {
// 				panic("mock out the UpdateUser method")
// 			},
// 		}
//
// 		// use mockedEnv in code that requires Env
// 		// and then make assertions.
//
// 	}
type EnvMock struct {
	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(id int) (*model.User, error)

	// NowFunc mocks the Now method.
	NowFunc func() time.Time

	// UpdateTransactionFunc mocks the UpdateTransaction method.
	UpdateTransactionFunc func(data *model.Transaction) error

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(data *model.User) error

	// calls tracks calls to the methods.
	calls struct {
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// ID is the id argument value.
			ID int
		}
		// Now holds details about calls to the Now method.
		Now []struct {
		}
		// UpdateTransaction holds details about calls to the UpdateTransaction method.
		UpdateTransaction []struct {
			// Data is the data argument value.
			Data *model.Transaction
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Data is the data argument value.
			Data *model.User
		}
	}
	lockGetUserByID       sync.RWMutex
	lockNow               sync.RWMutex
	lockUpdateTransaction sync.RWMutex
	lockUpdateUser        sync.RWMutex
}

// GetUserByID calls GetUserByIDFunc.
func (mock *EnvMock) GetUserByID(id int) (*model.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("EnvMock.GetUserByIDFunc: method is nil but Env.GetUserByID was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	mock.lockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(id)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//     len(mockedEnv.GetUserByIDCalls())
func (mock *EnvMock) GetUserByIDCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	mock.lockGetUserByID.RUnlock()
	return calls
}

// Now calls NowFunc.
func (mock *EnvMock) Now() time.Time {
	if mock.NowFunc == nil {
		panic("EnvMock.NowFunc: method is nil but Env.Now was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNow.Lock()
	mock.calls.Now = append(mock.calls.Now, callInfo)
	mock.lockNow.Unlock()
	return mock.NowFunc()
}

// NowCalls gets all the calls that were made to Now.
// Check the length with:
//     len(mockedEnv.NowCalls())
func (mock *EnvMock) NowCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNow.RLock()
	calls = mock.calls.Now
	mock.lockNow.RUnlock()
	return calls
}

// UpdateTransaction calls UpdateTransactionFunc.
func (mock *EnvMock) UpdateTransaction(data *model.Transaction) error {
	if mock.UpdateTransactionFunc == nil {
		panic("EnvMock.UpdateTransactionFunc: method is nil but Env.UpdateTransaction was just called")
	}
	callInfo := struct {
		Data *model.Transaction
	}{
		Data: data,
	}
	mock.lockUpdateTransaction.Lock()
	mock.calls.UpdateTransaction = append(mock.calls.UpdateTransaction, callInfo)
	mock.lockUpdateTransaction.Unlock()
	return mock.UpdateTransactionFunc(data)
}

// UpdateTransactionCalls gets all the calls that were made to UpdateTransaction.
// Check the length with:
//     len(mockedEnv.UpdateTransactionCalls())
func (mock *EnvMock) UpdateTransactionCalls() []struct {
	Data *model.Transaction
} {
	var calls []struct {
		Data *model.Transaction
	}
	mock.lockUpdateTransaction.RLock()
	calls = mock.calls.UpdateTransaction
	mock.lockUpdateTransaction.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *EnvMock) UpdateUser(data *model.User) error {
	if mock.UpdateUserFunc == nil {
		panic("EnvMock.UpdateUserFunc: method is nil but Env.UpdateUser was just called")
	}
	callInfo := struct {
		Data *model.User
	}{
		Data: data,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(data)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//     len(mockedEnv.UpdateUserCalls())
func (mock *EnvMock) UpdateUserCalls() []struct {
	Data *model.User
} {
	var calls []struct {
		Data *model.User
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
