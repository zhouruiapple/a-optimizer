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

package optimizer

import "aoptimizer/rel"

type MExpr struct {
	op       rel.Operator
	rmask    []bool
	operands []GroupID
	group    GroupID
	next     *MExpr
}

func NewMExpr(op rel.Operator, operands []GroupID, group GroupID, next *MExpr) *MExpr {
	return &MExpr{
		op:       op,
		rmask:    make([]bool, op.Arity()),
		operands: operands,
		group:    group,
		next:     next,
	}
}

func (mexpr *MExpr) Operator() rel.Operator {
	return mexpr.op
}

func (mexpr *MExpr) Operands() []GroupID {
	return mexpr.operands
}

func (mexpr *MExpr) Group() GroupID {
	return mexpr.group
}

func (mexpr *MExpr) Next() *MExpr {
	return mexpr.next
}

func (mexpr *MExpr) Marked(at int) bool {
	return mexpr.rmask[at]
}
