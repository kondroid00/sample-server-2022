package errors

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

var (
	ErrTest = Global("ErrTest")
)

type ErrTestStruct struct {
	message string
}

func (e ErrTestStruct) Error() string { return e.message }

func TestCustomError_Unwrap_Is(t *testing.T) {
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Isで検出できる",
			args: args{
				err:    Stack(ErrTest),
				target: ErrTest,
			},
			want: true,
		},
		{
			name: "wrapしたエラーがIsで検出できる",
			args: args{
				err:    Stack(fmt.Errorf("test:%w", ErrTest)),
				target: ErrTest,
			},
			want: true,
		},
		{
			name: "ネストしてwrapしたエラーがIsで検出できる",
			args: args{
				err:    fmt.Errorf("test:%w", Stack(fmt.Errorf("test:%w", ErrTest))),
				target: ErrTest,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomError_Unwrap_As(t *testing.T) {
	errTest1 := &ErrTestStruct{message: "ErrTest"}
	errTest2 := Stack(New(""))

	type args struct {
		err error
	}
	tests := []struct {
		name           string
		args           args
		want           bool
		wantTarget     interface{}
		wantTargetType interface{}
	}{
		{
			name: "Asで取得できる",
			args: args{
				err: Stack(errTest1),
			},
			want:           true,
			wantTarget:     errTest1,
			wantTargetType: &ErrTestStruct{},
		},
		{
			name: "wrapしたエラーがAsで取得できる",
			args: args{
				err: Stack(fmt.Errorf("test:%w", errTest1)),
			},
			want:           true,
			wantTarget:     errTest1,
			wantTargetType: &ErrTestStruct{},
		},
		{
			name: "ネストしてwrapしたエラーがAsで取得できる",
			args: args{
				err: fmt.Errorf("test:%w", Stack(fmt.Errorf("test:%w", errTest1))),
			},
			want:           true,
			wantTarget:     errTest1,
			wantTargetType: &ErrTestStruct{},
		},
		{
			name: "StackでネストしていてもAsでエラーを取得できる",
			args: args{
				err: Stack(errTest2),
			},
			want:           true,
			wantTarget:     errTest2,
			wantTargetType: &CustomError{},
		},
	}
	opt := cmp.Options{
		cmp.AllowUnexported(ErrTestStruct{}, CustomError{}),
		cmpopts.IgnoreUnexported(CustomError{}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.wantTargetType.(type) {
			case *ErrTestStruct:
				if got := As(tt.args.err, &v); got != tt.want {
					t.Errorf("As() = %v, want %v", got, tt.want)
					return
				}
				if diff := cmp.Diff(v, tt.wantTarget.(*ErrTestStruct), opt); diff != "" {
					t.Error(diff)
				}
			case *CustomError:
				if got := As(tt.args.err, &v); got != tt.want {
					t.Errorf("As() = %v, want %v", got, tt.want)
					return
				}
				if diff := cmp.Diff(v, tt.wantTarget.(*CustomError), opt); diff != "" {
					t.Error(diff)
				}
			}
		})
	}
}
