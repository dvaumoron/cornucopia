// This file is generated - do not edit.

package go_glu

import (
	jen "github.com/dave/jennifer/jen"
	glu "github.com/dvaumoron/cornucopia/glu"
	starlark "go.starlark.net/starlark"
)

func wrappedAny(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Any(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedBool(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Bool(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedBreak(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Break(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedByte(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Byte(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedChan(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Chan(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedComparable(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Comparable(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedComplex128(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Complex128(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedComplex64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Complex64(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedContinue(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Continue(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedDefault(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Default(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedDefer(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Defer(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedEmpty(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Empty(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedErr(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Err(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedError(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Error(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedFallthrough(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Fallthrough(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedFloat32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Float32(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedFloat64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Float64(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedFunc(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Func(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedGo(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Go(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedGoto(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Goto(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Int(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Int16(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Int32(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Int64(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedInt8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Int8(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedIota(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Iota(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedLine(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Line(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedNil(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Nil(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedNull(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Null(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedRecover(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Recover(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedRune(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Rune(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedSelect(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Select(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedString(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.String(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Uint(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Uint16(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Uint32(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Uint64(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUint8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Uint8(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUintptr(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Uintptr(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedVar(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Var(),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedAppend(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Append(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedCase(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Case(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedFor(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.For(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedIf(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.If(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedIndex(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Index(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedInterface(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Interface(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedList(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.List(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedMake(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Make(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedReturn(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Return(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedSwitch(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Switch(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedUnion(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return glu.Wrapper{
		Inner: jen.Union(convertToCodeSlice(args)...),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedCap(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Cap(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedClose(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Close(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedImag(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Imag(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedLen(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Len(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedMap(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Map(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedParens(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Parens(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func wrappedReal(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	return glu.Wrapper{
		Inner: jen.Real(convertToCode(value)),
		WType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Const(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Const()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Var(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Var()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Type(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Type()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenFile_Func(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Func()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Any(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Any()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Bool(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Bool()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Byte(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Byte()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Chan(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Chan()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Clone(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Clone()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Comparable(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Comparable()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Complex128(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Complex128()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Complex64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Complex64()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Else(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Else()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Err(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Err()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Error(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Error()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Float32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Float32()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Float64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Float64()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Func(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Func()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Int()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int16(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Int16()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Int32()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Int64()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Int8(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Int8()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Iota(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Iota()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Nil(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Nil()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Range(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Range()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Recover(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Recover()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Rune(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Rune()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_String(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.String()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Uint()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint16(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Uint16()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint32(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Uint32()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint64(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Uint64()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uint8(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Uint8()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Uintptr(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Uintptr()
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Add(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Add(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Append(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Append(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Block(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Block(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Call(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Call(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Defs(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Defs(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Index(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Index(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Interface(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Interface(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Make(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Make(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Params(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Params(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Struct(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Struct(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Types(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Types(convertToCodeSlice(args)...)
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Assert(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Assert(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Cap(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Cap(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Imag(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Imag(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Len(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Len(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Map(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Map(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Parens(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Parens(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}

func jenStatement_Real(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	casted, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	res := casted.Real(convertToCode(value))
	return glu.Wrapper{
		Inner: res,
		WType: &jenStatementWrappedType,
	}, nil
}
