package cliCodeGenerator

import (
	"github.com/bamtols/dbgo/cli/cliCodeGenerator/cliGoPrinter"
	"github.com/bamtols/dbgo/cli/cliTypes"
	"github.com/bamtols/dbgo/func/errFn"
)

type (
	V1 struct {
		configure *cliTypes.Configure
		template  *cliGoPrinter.Template
	}
)

func NewV1(t *cliTypes.Configure) (*V1, error) {
	template, err := cliGoPrinter.NewTemplatePrinter(
		t.Version.String(),
		t.GoPackageNm(),
	)

	if err != nil {
		return nil, err
	}

	res := &V1{
		configure: t,
		template:  template,
	}

	res.GenerateCode()

	return res, nil
}

func (x *V1) GenerateCode() {
	x.createImports()
	x.createModels()
}

func (x *V1) Print() string {
	return x.template.PrintLn()
}

func (x *V1) createImports() {
	for _, definition := range x.configure.ModelConfigure().GoTypeMap {
		if definition.ImportGoCode == nil {
			continue
		}
		x.template.AddImport(*definition.ImportGoCode)
	}
}

func (x *V1) createModels() {
	// 기본 값 입력
	modelConfigure := x.configure.ModelConfigure()
	for modelNm, config := range modelConfigure.ModelMap {
		printer := x.genStructPrinter(modelNm, config)
		x.template.AddBody(printer)
	}
}

func (x *V1) genStructPrinter(modelNm cliTypes.ModelNm, config cliTypes.TmpModel) cliGoPrinter.PrinterFn {
	modelConfigure := x.configure.ModelConfigure()
	goTypeMap := modelConfigure.GoTypeMap

	structPrinter := cliGoPrinter.NewStructPrinter(modelNm.String())

	// 템플릿의 데이터 컬럼 입력
	for fieldNm, fieldType := range config.Columns {
		if fieldType.HasTagOne(
			cliTypes.ColumnTagPrimaryKey,
			cliTypes.ColumnTagNullable,
		) {
			structPrinter.AddField(fieldNm.String(), goTypeMap.PrintPtrGoType(fieldType.GoType))
		} else {
			structPrinter.AddField(fieldNm.String(), goTypeMap.PrintGoType(fieldType.GoType))
		}
	}

	// Hooks 추가
	if config.Hooks != nil {
		// Hooks.TimeAt 컬럼 추가
		if len(config.Hooks.TimeAt) != 0 {
			for _, timeAt := range config.Hooks.TimeAt {
				switch timeAt {
				case cliTypes.TimeAtCreatedAt, cliTypes.TimeAtUpdatedAt:
					structPrinter.AddField(timeAt.String(), goTypeMap.PrintGoType(cliTypes.GoTypeTime))
				case cliTypes.TimeAtDeletedAt:
					structPrinter.AddField(timeAt.String(), goTypeMap.PrintPtrGoType(cliTypes.GoTypeTime))
				default:
					panic(errFn.BindF(ErrInvalid, "HookTimeAt : TimeAt=%s", timeAt))
				}
			}
		}
	}

	// 바인딩된 객체 추가
	if config.Binds != nil {
		for fieldNm, bindConfig := range config.Binds {
			if !modelConfigure.HasModel(bindConfig.Model) {
				panic(errFn.BindF(ErrNotFound, "Model : ModelNm=%s", bindConfig.Model))
			}

			switch bindConfig.BindType {
			case cliTypes.BindTypeBelongsTo, cliTypes.BindTypeHasOne:
				structPrinter.AddField(fieldNm.GoPublicModelNm(), fieldNm.PrintPtrGoType())
			case cliTypes.BindTypeHasMany:
				structPrinter.AddField(fieldNm.GoPublicModelNm(), fieldNm.PrintSliceGoType())
			default:
				panic(errFn.BindF(ErrInvalid, "BindType : BindType%s", bindConfig.BindType))
			}
		}
	}

	return structPrinter
}
