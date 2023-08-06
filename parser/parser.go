package parser

import (
	"go/ast"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

func ExtractEnumeration(packagePath string, targetEnumName string) (string, map[string]string, error) {
	packageConf := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedName | packages.NeedTypesInfo | packages.NeedTypes,
	}
	data, err := packages.Load(packageConf, packagePath)
	if err != nil {
		return "", nil, errors.Wrap(err, "could not load packages")
	}

	result := make(map[string]string)
	for node := range data[0].TypesInfo.Scopes {
		if fileNode, ok := node.(*ast.File); ok {
			for objectName, objectData := range fileNode.Scope.Objects {
				if valueData, ok := objectData.Decl.(*ast.ValueSpec); ok && len(valueData.Values) == 1 && objectData.Kind == ast.Con {
					if identType, ok := valueData.Type.(*ast.Ident); ok && identType.Name == targetEnumName {
						if litData, ok := valueData.Values[0].(*ast.BasicLit); ok {
							result[objectName] = litData.Value
						}
					}
				}
			}
		}
	}

	return data[0].Name, result, nil
}
