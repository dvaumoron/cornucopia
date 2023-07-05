// This file is generated - do not edit.

package main

import (
	jen "github.com/dave/jennifer/jen"
	starlark "go.starlark.net/starlark"
)

func wrappedAny(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Any(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedBool(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Bool(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedBreak(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Break(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedByte(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Byte(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedChan(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Chan(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedComparable(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Comparable(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedComplex64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Complex64(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedComplex128(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Complex128(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedContinue(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Continue(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedDefault(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Default(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedDefer(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Defer(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedEmpty(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Empty(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedErr(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Err(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedError(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Error(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedFallthrough(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Fallthrough(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedFloat32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Float32(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedFloat64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Float64(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedFunc(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Func(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedGo(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Go(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedGoto(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Goto(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Int(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Int8(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Int16(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Int32(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Int64(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedIota(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Iota(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedLine(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Line(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedNil(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Nil(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedNull(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Null(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedRecover(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Recover(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedRune(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Rune(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedSelect(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Select(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedString(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.String(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Uint(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Uint8(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Uint16(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Uint32(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Uint64(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedUintptr(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Uintptr(),
		wType: &jenStatementWrappedType,
	}, nil
}

func wrappedVar(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{
		inner: jen.Var(),
		wType: &jenStatementWrappedType,
	}, nil
}

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

func jenStatement_Func(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Func()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Iota(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Iota()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int8(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int8()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int16(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int16()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int32()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Int64()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Nil(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Nil()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Range(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Range()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Recover(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Recover()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Rune(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Rune()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_String(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.String()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint8(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint8()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint16(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint16()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint32()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uint64()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uintptr(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv := b.Receiver().(wrapper[*jen.Statement])
	stmt := recv.inner.Uintptr()
	return wrapper[*jen.Statement]{
		inner: stmt,
		wType: &jenStatementWrappedType,
	}, nil
}
