/*****************************************************************************/
/* interpret.go                                                              */
/*****************************************************************************/
/*                           This file is part of:                           */
/*                                KAIJU ENGINE                               */
/*                          https://kaijuengine.org                          */
/*****************************************************************************/
/* MIT License                                                               */
/*                                                                           */
/* Copyright (c) 2022-present Kaiju Engine contributors (CONTRIBUTORS.md).   */
/* Copyright (c) 2015-2022 Brent Farris.                                     */
/*                                                                           */
/* May all those that this source may reach be blessed by the LORD and find  */
/* peace and joy in life.                                                    */
/* "Everyone who drinks of this water will be thirsty again; but whoever     */
/* drinks of the water that I will give him shall never thirst;" -Jesus      */
/*                                                                           */
/* Permission is hereby granted, free of charge, to any person obtaining a   */
/* copy of this software and associated documentation files (the "Software"),*/
/* to deal in the Software without restriction, including without limitation */
/* the rights to use, copy, modify, merge, publish, distribute, sublicense,  */
/* and/or sell copies of the Software, and to permit persons to whom the     */
/* Software is furnished to do so, subject to the following conditions:      */
/*                                                                           */
/* The above copyright, blessing, notice and this permission notice shall    */
/* be included in all copies or substantial portions of the Software.        */
/*                                                                           */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS   */
/* OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF                */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.    */
/* IN NO EVENT SHALL THE /* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT */
/* OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE     */
/* OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                             */
/*****************************************************************************/

package project

import (
	"kaiju/engine"
	"kaiju/filesystem"
	"kaiju/interpreter"
	"path/filepath"

	"github.com/KaijuEngine/yaegi/interp"
	"github.com/KaijuEngine/yaegi/stdlib"
)

func InterpretSource(projectPath string, host *engine.Host) error {
	//project.CreateNewProject(projectPath)
	itp := interp.New(interp.Options{ImportSub: interp.ImportSubstitution{
		Prefix:  "kaiju/source",
		Replace: projectPath + "/source",
	}})
	if err := itp.Use(stdlib.Symbols); err != nil {
		return err
	}
	if err := itp.Use(interpreter.Symbols); err != nil {
		return err
	}
	src, err := filesystem.ReadTextFile(filepath.Join(projectPath, "source/source.go"))
	if err != nil {
		return err
	}
	if _, err := itp.Eval(src); err != nil {
		return err
	}
	var entry func(*engine.Host)
	if sourceMain, err := itp.Eval("source.Main"); err != nil {
		return err
	} else {
		entry = sourceMain.Interface().(func(*engine.Host))
	}
	entry(host)
	return nil
}
