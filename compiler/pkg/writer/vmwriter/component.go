package vmwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/compilation_engine"
	"fmt"
)

// getSegmentFromSymbolAttribute SymbolAttributeから対応するSegmentを習得する
func getSegmentFromSymbolAttribute(attribute string) string {

	switch attribute {
	case string(compilation_engine.StaticVariableKind):
		return segmentStatic
	case string(compilation_engine.FieldVariableKind):
		return segmentThis
	case string(compilation_engine.LocalVariableKind):
		return segmentLocal
	case string(compilation_engine.ArgumentVariableKind):
		return segmentArgument
	default:
		chk.SE(fmt.Errorf("未定義のSymbolAttributeが渡されました:%s", attribute))
		return ""
	}
}
