// Package parser implements source code parsing.
// It uses parser (and lexer) generated by ANTLR4 from neva.g4 grammar file.
package parser

import (
	"fmt"
	"runtime/debug"

	"github.com/antlr4-go/antlr/v4"

	"github.com/nevalang/neva/internal/compiler"
	generated "github.com/nevalang/neva/internal/compiler/parser/generated"
	src "github.com/nevalang/neva/internal/compiler/sourcecode"
)

type treeShapeListener struct {
	*generated.BasenevaListener
	file src.File
}

type Parser struct {
	isDebug bool
}

func (p Parser) ParseModules(
	rawMods map[src.ModuleRef]compiler.RawModule,
) (map[src.ModuleRef]src.Module, *compiler.Error) {
	parsedMods := make(map[src.ModuleRef]src.Module, len(rawMods))

	for modRef, rawMod := range rawMods {
		parsedPkgs, err := p.ParsePackages(modRef, rawMod.Packages)
		if err != nil {
			return nil, err
		}

		parsedMods[modRef] = src.Module{
			Manifest: rawMod.Manifest,
			Packages: parsedPkgs,
		}
	}

	return parsedMods, nil
}

func (p Parser) ParsePackages(
	modRef src.ModuleRef,
	rawPkgs map[string]compiler.RawPackage,
) (
	map[string]src.Package,
	*compiler.Error,
) {
	packages := make(map[string]src.Package, len(rawPkgs))

	for pkgName, pkgFiles := range rawPkgs {
		parsedFiles, err := p.ParseFiles(modRef, pkgName, pkgFiles)
		if err != nil {
			return nil, err
		}
		packages[pkgName] = parsedFiles
	}

	return packages, nil
}

func (p Parser) ParseFiles(
	modRef src.ModuleRef,
	pkgName string,
	files map[string][]byte,
) (map[string]src.File, *compiler.Error) {
	result := make(map[string]src.File, len(files))

	for fileName, fileBytes := range files {
		parsedFile, err := p.parseFile(fileBytes)
		if err != nil {
			err.Location = compiler.Pointer(src.Location{
				ModRef:   modRef,
				PkgName:  pkgName,
				FileName: fileName,
			})
			return nil, err
		}
		result[fileName] = parsedFile
	}

	return result, nil
}

func (p Parser) parseFile(bb []byte) (src.File, *compiler.Error) {
	input := antlr.NewInputStream(string(bb))
	lexer := generated.NewnevaLexer(input)
	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	parserErrors := &CustomErrorListener{}
	prsr := generated.NewnevaParser(tokenStream)
	prsr.RemoveErrorListeners()
	prsr.AddErrorListener(parserErrors)
	if p.isDebug {
		prsr.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	}
	prsr.BuildParseTrees = true

	listener := &treeShapeListener{}

	if err := walkTree(listener, prsr.Prog()); err != nil {
		return src.File{}, err
	}

	if len(lexerErrors.Errors) > 0 {
		return src.File{}, lexerErrors.Errors[0]
	}

	if len(parserErrors.Errors) > 0 {
		return src.File{}, parserErrors.Errors[0]
	}

	return listener.file, nil
}

func walkTree(listener antlr.ParseTreeListener, tree antlr.ParseTree) (err *compiler.Error) {
	defer func() {
		if e := recover(); e != nil {
			if _, ok := e.(*compiler.Error); !ok {
				err = &compiler.Error{
					Err: fmt.Errorf(
						"%v: %v",
						e,
						string(debug.Stack()),
					),
				}
			}
		}
	}()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return nil
}

func New(isDebug bool) Parser {
	return Parser{isDebug: isDebug}
}
