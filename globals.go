/*
 *
 * Copyright 2023 cornucopia authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"github.com/dave/jennifer/jen"
	"go.starlark.net/starlark"
)

func initCornucopiaGlobals() {
	addGlobals("NewFile", wrappedNewFile)
	addGlobals("Any", wrappedAny)
	addGlobals("Append", wrappedAppend)
	addGlobals("Bool", wrappedBool)
	addGlobals("Break", wrappedBreak)
	addGlobals("Byte", wrappedByte)
	addGlobals("Cap", wrappedCap)
	addGlobals("Case", wrappedCase)
	addGlobals("Chan", wrappedChan)
	addGlobals("Close", wrappedClose)
	addGlobals("Comment", wrappedComment)
	addGlobals("Comparable", wrappedComparable)
	addGlobals("Complex", wrappedComplex)
	addGlobals("Complex64", wrappedComplex64)
	addGlobals("Complex128", wrappedComplex128)
	addGlobals("Continue", wrappedContinue)
	addGlobals("Copy", wrappedCopy)
	addGlobals("Default", wrappedDefault)
	addGlobals("Defer", wrappedDefer)
	addGlobals("Delete", wrappedDelete)
	addGlobals("Empty", wrappedEmpty)
	addGlobals("Err", wrappedErr)
	addGlobals("Fallthrough", wrappedFallthrough)
	addGlobals("Float32", wrappedFloat32)
	addGlobals("Float64", wrappedFloat64)
	addGlobals("For", wrappedFor)
	addGlobals("Go", wrappedGo)
	addGlobals("Goto", wrappedGoto)
	addGlobals("Id", wrappedId)
	addGlobals("If", wrappedIf)
	addGlobals("Imag", wrappedImag)
	addGlobals("Index", wrappedIndex)
	addGlobals("Int", wrappedInt)
	addGlobals("Int8", wrappedInt8)
	addGlobals("Int16", wrappedInt16)
	addGlobals("Int32", wrappedInt32)
	addGlobals("Int64", wrappedInt64)
	addGlobals("Interface", wrappedInterface)
	addGlobals("Iota", wrappedIota)
	addGlobals("Len", wrappedLen)
	addGlobals("Line", wrappedLine)
	addGlobals("List", wrappedList)
	addGlobals("Lit", wrappedLit)
	addGlobals("LitByte", wrappedLitByte)
	addGlobals("LitRune", wrappedLitRune)
	addGlobals("Make", wrappedMake)
	addGlobals("Map", wrappedMap)
	addGlobals("Nil", wrappedNil)
	addGlobals("Null", wrappedNull)
	addGlobals("Op", wrappedOp)
	addGlobals("Parens", wrappedParens)
	addGlobals("Qual", wrappedQual)
	addGlobals("Real", wrappedReal)
	addGlobals("Recover", wrappedRecover)
	addGlobals("Return", wrappedReturn)
	addGlobals("Select", wrappedSelect)
	addGlobals("String", wrappedString)
	addGlobals("Switch", wrappedSwitch)
	addGlobals("Uint", wrappedUint)
	addGlobals("Uint8", wrappedUint8)
	addGlobals("Uint16", wrappedUint16)
	addGlobals("Uint32", wrappedUint32)
	addGlobals("Uint64", wrappedUint64)
	addGlobals("Uintptr", wrappedUintptr)
	addGlobals("Union", wrappedUnion)
	addGlobals("Var", wrappedVar)
	addGlobals("Values", wrappedValues)
}

func addGlobals(name string, wrapped func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)) {
	starlark.Universe[name] = starlark.NewBuiltin(name, wrapped)
}

func wrappedNewFile(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var packageName string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "packageName", &packageName); err != nil {
		return nil, err
	}

	file := jen.NewFile(packageName)
	return wrapper[*jen.File]{inner: file, wType: &jenFileWrappedType}, nil
}

func wrappedAny(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Any(), wType: &jenStatementWrappedType}, nil
}

func wrappedAppend(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Append(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedBool(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Bool(), wType: &jenStatementWrappedType}, nil
}

func wrappedBreak(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Break(), wType: &jenStatementWrappedType}, nil
}

func wrappedByte(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Byte(), wType: &jenStatementWrappedType}, nil
}

func wrappedCap(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	stmt := jen.Cap(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedCase(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	case_ := jen.Case(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: case_, wType: &jenStatementWrappedType}, nil
}

func wrappedChan(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Chan(), wType: &jenStatementWrappedType}, nil
}

func wrappedClose(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	stmt := jen.Close(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &value); err != nil {
		return nil, err
	}
	jen.Comment(value)
	return starlark.None, nil
}

func wrappedComparable(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Comparable(), wType: &jenStatementWrappedType}, nil
}

func wrappedComplex(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}
	stmt := jen.Complex(convertToCode(m), convertToCode(key))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedComplex64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Complex64(), wType: &jenStatementWrappedType}, nil
}

func wrappedComplex128(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Complex128(), wType: &jenStatementWrappedType}, nil
}

func wrappedContinue(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Continue(), wType: &jenStatementWrappedType}, nil
}

func wrappedCopy(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var dst starlark.Value
	var src starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "dst", &dst, "src", &src); err != nil {
		return nil, err
	}
	stmt := jen.Copy(convertToCode(dst), convertToCode(src))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedDefault(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Default(), wType: &jenStatementWrappedType}, nil
}

func wrappedDefer(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Defer(), wType: &jenStatementWrappedType}, nil
}

func wrappedDelete(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}
	stmt := jen.Delete(convertToCode(m), convertToCode(key))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedEmpty(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Empty(), wType: &jenStatementWrappedType}, nil
}

func wrappedErr(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Err(), wType: &jenStatementWrappedType}, nil
}

func wrappedFallthrough(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Fallthrough(), wType: &jenStatementWrappedType}, nil
}

func wrappedFloat32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Float32(), wType: &jenStatementWrappedType}, nil
}

func wrappedFloat64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Float64(), wType: &jenStatementWrappedType}, nil
}

func wrappedFor(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.For(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedGo(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Go(), wType: &jenStatementWrappedType}, nil
}

func wrappedGoto(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Goto(), wType: &jenStatementWrappedType}, nil
}

func wrappedId(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}
	id := jen.Id(name)
	return wrapper[*jen.Statement]{inner: id, wType: &jenStatementWrappedType}, nil
}

func wrappedIf(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.If(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedImag(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	stmt := jen.Imag(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedIndex(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Index(convertToCode(args))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedInt(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int8(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int16(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int32(), wType: &jenStatementWrappedType}, nil
}

func wrappedInt64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Int64(), wType: &jenStatementWrappedType}, nil
}

func wrappedInterface(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Interface(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedIota(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Iota(), wType: &jenStatementWrappedType}, nil
}

func wrappedLen(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	stmt := jen.Len(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedLine(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Line(), wType: &jenStatementWrappedType}, nil
}

func wrappedList(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	list := jen.List(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: list, wType: &jenStatementWrappedType}, nil
}

func wrappedLit(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.Lit(convertToGoBuiltin(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedLitByte(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.LitByte(convertToGoByte(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedLitRune(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}
	lit := jen.LitRune(convertToGoRune(value))
	return wrapper[*jen.Statement]{inner: lit, wType: &jenStatementWrappedType}, nil
}

func wrappedMake(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Make(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedMap(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var type_ starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "type", &type_); err != nil {
		return nil, err
	}
	stmt := jen.Map(convertToCode(type_))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedNil(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Nil(), wType: &jenStatementWrappedType}, nil
}

func wrappedNull(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Null(), wType: &jenStatementWrappedType}, nil
}

func wrappedOp(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "op", &value); err != nil {
		return nil, err
	}
	op := jen.Op(value)
	return wrapper[*jen.Statement]{inner: op, wType: &jenStatementWrappedType}, nil
}
func wrappedParens(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var item starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "item", &item); err != nil {
		return nil, err
	}
	parens := jen.Parens(convertToCode(item))
	return wrapper[*jen.Statement]{inner: parens, wType: &jenStatementWrappedType}, nil
}

func wrappedQual(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}
	qual := jen.Qual(path, name)
	return wrapper[*jen.Statement]{inner: qual, wType: &jenStatementWrappedType}, nil
}

func wrappedReal(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "c", &value); err != nil {
		return nil, err
	}
	stmt := jen.Real(convertToCode(value))
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedRecover(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Recover(), wType: &jenStatementWrappedType}, nil
}

func wrappedReturn(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return_ := jen.Return(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: return_, wType: &jenStatementWrappedType}, nil
}

func wrappedSelect(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Select(), wType: &jenStatementWrappedType}, nil
}

func wrappedString(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.String(), wType: &jenStatementWrappedType}, nil
}

func wrappedSwitch(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Switch(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedUint(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint8(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint8(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint16(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint16(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint32(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint32(), wType: &jenStatementWrappedType}, nil
}

func wrappedUint64(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uint64(), wType: &jenStatementWrappedType}, nil
}

func wrappedUintptr(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Uintptr(), wType: &jenStatementWrappedType}, nil
}

func wrappedUnion(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Union(convertToCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}

func wrappedVar(_ *starlark.Thread, _ *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	return wrapper[*jen.Statement]{inner: jen.Var(), wType: &jenStatementWrappedType}, nil
}

func wrappedValues(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	stmt := jen.Values(convertToDictOrCodeSlice(args)...)
	return wrapper[*jen.Statement]{inner: stmt, wType: &jenStatementWrappedType}, nil
}
