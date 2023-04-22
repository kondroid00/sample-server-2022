package errors

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	e "github.com/pkg/errors"
)

// Error StackTraceが追跡可能なErrorインターフェース
type Error interface {
	Error() string
	StackTrace() e.StackTrace
	Unwrap() error
}

//---------------------------------------------------------------------------------
//
//  CustomError
//  pkg/errorsのfundamentalをラップしてerrorインターフェースを満たす
//
//---------------------------------------------------------------------------------
type (
	CustomError struct {
		// github.com/pkg/errorsの*fundamentalを入れる
		err error
	}
)

// Error Errorインターフェースを満たす
func (ce *CustomError) Error() string {
	return ce.err.Error()
}

// StackTrace StackTrace()が実装されている場合にそれを呼び出す。sentryにstacktraceを送信する際に利用する
func (ce *CustomError) StackTrace() e.StackTrace {
	if stackTracer, ok := ce.err.(interface{ StackTrace() e.StackTrace }); ok {
		return stackTracer.StackTrace()
	}
	return nil
}

// Unwrap errors.Isとerrors.Asで利用する
func (ce *CustomError) Unwrap() error {
	return e.Cause(ce.err)
}

// New messageとstacktraceを格納したエラーを返す
func New(message string) Error {
	return &CustomError{
		err: e.New(message),
	}
}

// Stack 既存のエラーにstacktraceを格納したエラーを返す
func Stack(err error) Error {
	return &CustomError{
		err: e.WithStack(err),
	}
}

//---------------------------------------------------------------------------------
//
//  errorsパッケージを別でimportする必要をなくすためのラップ関数
//
//---------------------------------------------------------------------------------
// Global グローバル領域に格納するエラーを返す
// 送信時にはStack()と組み合わせて送信する必要がある（stacktraceが取れないので）
// ----------------------------------
// var Err = errors.Global("global")
//
// func foo() err {
//     return errors.Stack(Err)
// }
// ----------------------------------
func Global(message string) error {
	return errors.New(message)
}

// Is errors.Isと同じ
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As errors.Asと同じ
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

//---------------------------------------------------------------------------------
//
//  multiError用のラッパー
//
//---------------------------------------------------------------------------------
type (
	MultiError struct {
		err *multierror.Error
	}
)

// NewMultiError MultiErrorを生成
func NewMultiError() *MultiError {
	return &MultiError{}
}

// Append err != nilであればerrを追加する
func (me *MultiError) Append(err error) {
	if err != nil {
		me.err = multierror.Append(me.err, err)
	}
}

// Error Errorインターフェースを満たす
func (me *MultiError) Error() string {
	if me.err == nil {
		return ""
	}
	return me.err.Error()
}

// Unwrap errors.Isとerrors.Asで利用する
func (me *MultiError) Unwrap() error {
	if me.err == nil {
		return nil
	}
	return me.err.Unwrap()
}

// GetError multiErrorを取得する
func (me *MultiError) GetError() error {
	return me.err.ErrorOrNil()
}
