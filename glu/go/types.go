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

package go_glu

import (
	"path"

	"github.com/dave/jennifer/jen"
	"github.com/dvaumoron/cornucopia/common"
	glu "github.com/dvaumoron/cornucopia/glu"
	"go.starlark.net/starlark"
)

var (
	jenFileWrappedType      glu.WrappedType
	jenStatementWrappedType glu.WrappedType
)

func jenFile_HeaderComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.HeaderComment(comment)
	return starlark.None, nil
}

func jenFile_PackageComment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.PackageComment(comment)
	return starlark.None, nil
}

func jenFile_Comment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.Comment(comment)
	return starlark.None, nil
}

func jenFile_Anon(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.Anon(convertToStringSlice(args)...)
	return starlark.None, nil
}

func jenFile_ImportAlias(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var alias string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &alias); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.ImportAlias(path, alias)
	return starlark.None, nil
}

func jenFile_ImportName(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.ImportName(path, name)
	return starlark.None, nil
}

func jenFile_ImportNames(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var names starlark.Dict
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "names", &names); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.ImportNames(convertToMapString(names.Items()))
	return starlark.None, nil
}

func jenFile_Line(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	file.Line()
	return starlark.None, nil
}

func jenFile_Save(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var fileName string
	err := starlark.UnpackArgs(b.Name(), args, kwargs, "fileName", &fileName)
	if err != nil {
		return nil, err
	}

	if path.IsAbs(fileName) {
		return nil, glu.ErrForbidAbsolute
	}

	if err = common.EnsureWrite(fileName); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	file, ok := recv.Inner.(*jen.File)
	if !ok {
		return nil, glu.ErrCast
	}

	if err = file.Save(fileName); err != nil {
		return nil, err
	}

	glu.GeneratedFileNames = append(glu.GeneratedFileNames, fileName)
	return starlark.None, nil
}

func jenStatement_Comment(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var comment string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "comment", &comment); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	stmt.Comment(comment)
	return starlark.None, nil
}

func jenStatement_Complex(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var m starlark.Value
	var key starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "m", &m, "key", &key); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Complex(convertToCode(m), convertToCode(key))
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Dot(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Dot(name)
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Id(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Id(name)
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Lit(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	lit := stmt.Lit(glu.ConvertToGoBaseType(value))
	return glu.Wrapper{Inner: lit, WType: &jenStatementWrappedType}, nil
}

func jenStatement_LitByte(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	lit := stmt.LitByte(convertToGoByte(value))
	return glu.Wrapper{Inner: lit, WType: &jenStatementWrappedType}, nil
}

func jenStatement_LitRune(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var value starlark.Value
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "value", &value); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	lit := stmt.LitRune(convertToGoRune(value))
	return glu.Wrapper{Inner: lit, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Op(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var op string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "op", &op); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Op(op)
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Qual(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var path string
	var name string
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "path", &path, "name", &name); err != nil {
		return nil, err
	}

	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Qual(path, name)
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Tag(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Tag(convertToMapString(kwargs))
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}

func jenStatement_Values(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	recv, _ := b.Receiver().(glu.Wrapper)
	stmt, ok := recv.Inner.(*jen.Statement)
	if !ok {
		return nil, glu.ErrCast
	}

	newStmt := stmt.Values(convertToDictOrCodeSlice(args)...)
	return glu.Wrapper{Inner: newStmt, WType: &jenStatementWrappedType}, nil
}
