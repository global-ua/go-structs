package structs

import (
	"database/sql/driver"
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

// RuleExpr struct for db store and execute scripted expressions
// based on github.com/antonmedv/expr engine
type RuleExpr struct {
	Source  string
	Program *vm.Program
}

// RunBool execute expression and return result as boolean
func (e *RuleExpr) RunBool(env interface{}) (bool, error) {
	res, err := expr.Run(e.Program, env)
	if err != nil {
		return false, fmt.Errorf("expression run error: %v", err.Error())
	}
	switch res.(type) {
	case bool:
		return res.(bool), nil
	default:
		return false, fmt.Errorf("expression return not boolean type : %v", e.Source)
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (e *RuleExpr) Scan(value interface{}) error {
	e.Source = string(value.([]byte))
	prg, err := expr.Compile(e.Source, expr.Env(nil))
	if err == nil {
		e.Program = prg
	}
	return err
}

// Value implements the driver.Valuer interface for database serialization.
func (e RuleExpr) Value() (driver.Value, error) {
	return e.Source, nil
}
