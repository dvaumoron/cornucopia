// This file is generated - do not edit.

package main

import (
	jen "github.com/dave/jennifer/jen"
	starlark "go.starlark.net/starlark"
)

func jenFile_Const(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Const()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Var(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Var()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Type(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Type()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Func(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.File])
	stmt := recv.inner.Func()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Any(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Any()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Bool(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Bool()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Byte(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Byte()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Chan(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Chan()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Clone(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Clone()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Comparable(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Comparable()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Complex64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Complex64()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Complex128(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Complex128()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Else(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Else()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Err(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Err()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Error(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Error()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Float32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Float32()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Float64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Float64()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Line(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Line()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}
