package symboltable

// createSymbol .
func createSymbol(varName string, t string, attribute string, num int32) *Symbol {
	return &Symbol{
		VarName:   varName,
		Type:      t,
		Attribute: attribute,
		Num:       num,
	}
}
