//			GNU GENERAL PUBLIC LICENSE
//			Version 3, 29 June 2007
//
//Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
//Everyone is permitted to copy and distribute verbatim copies
//of this license document, but changing it is not allowed.
//
//			Preamble
//
//The GNU General Public License is a free, copyleft license for
//software and other kinds of works.
//
//The licenses for most software and other practical works are designed
//to take away your freedom to share and change the works.  By contrast,
//the GNU General Public License is intended to guarantee your freedom to
//share and change all versions of a program--to make sure it remains free
//software for all its users.  We, the Free Software Foundation, use the
//GNU General Public License for most of our software; it applies also to
//any other work released this way by its authors.  You can apply it to
//your programs, too.

package rel

import (
	"fmt"
)

type LogicalScan struct {
	coll  string
	alias string
}

func NewLogicalScan(coll string, alias string) *LogicalScan {
	return &LogicalScan{
		coll:  coll,
		alias: alias,
	}
}

func (op *LogicalScan) Name() string {
	return "LogicalScan"
}
func (op *LogicalScan) Typ() OpID {
	return LOGICAL_SCAN
}
func (op *LogicalScan) Arity() int {
	return 0
}
func (op *LogicalScan) EqualTo(other Operator) bool {
	if o, ok := other.(*LogicalScan); ok {
		return op.coll == o.coll && op.alias == o.alias
	}
	return false
}
func (op *LogicalScan) ToString() string {
	return fmt.Sprintf("%s (coll=%s, alias=%s)", OperatorToString(op), op.coll, op.alias)
}
func (op *LogicalScan) Match(other LogicalOperator) bool {
	_, ok := other.(*LogicalScan)
	return ok
}
func (op *LogicalScan) FindLogicalProp() interface{} {
	return nil
}

type LogicalJoin struct {
	cond string
}

func NewLogicalJoin(cond string) *LogicalJoin {
	return &LogicalJoin{
		cond: cond,
	}
}

func (op *LogicalJoin) Name() string {
	return "LogicalJoin"
}
func (op *LogicalJoin) Typ() OpID {
	return LOGICAL_JOIN
}
func (op *LogicalJoin) Arity() int {
	return 2
}
func (op *LogicalJoin) EqualTo(other Operator) bool {
	if o, ok := other.(*LogicalJoin); ok {
		return op.cond == o.cond
	}
	return false
}
func (op *LogicalJoin) ToString() string {
	return fmt.Sprintf("%s (cond=%s)", OperatorToString(op), op.cond)
}
func (op *LogicalJoin) Match(other LogicalOperator) bool {
	_, ok := other.(*LogicalJoin)
	return ok
}
func (op *LogicalJoin) FindLogicalProp() interface{} {
	return nil
}

type LogicalProject struct {
	attrs []string
}

func NewLogicalProject(attrs []string) *LogicalProject {
	return &LogicalProject{
		attrs: attrs,
	}
}

func (op *LogicalProject) Name() string {
	return "LogicalProject"
}
func (op *LogicalProject) Typ() OpID {
	return LOGICAL_PROJECT
}
func (op *LogicalProject) Arity() int {
	return 1
}
func (op *LogicalProject) EqualTo(other Operator) bool {
	if o, ok := other.(*LogicalProject); ok {
		if len(op.attrs) != len(o.attrs) {
			return false
		}

		for i, attr := range op.attrs {
			if attr != o.attrs[i] {
				return false
			}
		}
		return true
	}
	return false
}
func (op *LogicalProject) ToString() string {
	return fmt.Sprintf("%s (attrs=%v)", OperatorToString(op), op.attrs)
}
func (op *LogicalProject) Match(other LogicalOperator) bool {
	_, ok := other.(*LogicalProject)
	return ok
}
func (op *LogicalProject) FindLogicalProp() interface{} {
	return nil
}

type LogicalSelect struct {
	pred string
}

func NewLogicalSelect(pred string) *LogicalSelect {
	return &LogicalSelect{
		pred: pred,
	}
}

func (op *LogicalSelect) Name() string {
	return "LogicalSelect"
}
func (op *LogicalSelect) Typ() OpID {
	return LOGICAL_SELECT
}
func (op *LogicalSelect) Arity() int {
	return 1
}
func (op *LogicalSelect) EqualTo(other Operator) bool {
	if o, ok := other.(*LogicalSelect); ok {
		return op.pred == o.pred
	}
	return false
}
func (op *LogicalSelect) ToString() string {
	return fmt.Sprintf("%s (pred=%s)", OperatorToString(op), op.pred)
}
func (op *LogicalSelect) Match(other LogicalOperator) bool {
	_, ok := other.(*LogicalSelect)
	return ok
}
func (op *LogicalSelect) FindLogicalProp() interface{} {
	return nil
}

type LogicalAggregate struct {
	aggs []string
}

func NewLogicalAggregate(aggs []string) *LogicalAggregate {
	return &LogicalAggregate{
		aggs: aggs,
	}
}

func (op *LogicalAggregate) Name() string {
	return "LogicalAggregate"
}
func (op *LogicalAggregate) Typ() OpID {
	return LOGICAL_SELECT
}
func (op *LogicalAggregate) Arity() int {
	return 1
}
func (op *LogicalAggregate) EqualTo(other Operator) bool {
	if o, ok := other.(*LogicalAggregate); ok {
		if len(op.aggs) != len(o.aggs) {
			return false
		}

		for i, agg := range op.aggs {
			if agg != o.aggs[i] {
				return false
			}
		}
		return true
	}
	return false
}
func (op *LogicalAggregate) ToString() string {
	return fmt.Sprintf("%s (aggs=%s)", OperatorToString(op), op.aggs)
}
func (op *LogicalAggregate) Match(other LogicalOperator) bool {
	_, ok := other.(*LogicalAggregate)
	return ok
}
func (op *LogicalAggregate) FindLogicalProp() interface{} {
	return nil
}
