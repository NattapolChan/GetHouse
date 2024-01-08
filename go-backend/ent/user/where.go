// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NattapolChan/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// EmailAddr applies equality check predicate on the "emailAddr" field. It's identical to EmailAddrEQ.
func EmailAddr(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailAddr, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// EmailAddrEQ applies the EQ predicate on the "emailAddr" field.
func EmailAddrEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailAddr, v))
}

// EmailAddrNEQ applies the NEQ predicate on the "emailAddr" field.
func EmailAddrNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmailAddr, v))
}

// EmailAddrIn applies the In predicate on the "emailAddr" field.
func EmailAddrIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmailAddr, vs...))
}

// EmailAddrNotIn applies the NotIn predicate on the "emailAddr" field.
func EmailAddrNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmailAddr, vs...))
}

// EmailAddrGT applies the GT predicate on the "emailAddr" field.
func EmailAddrGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmailAddr, v))
}

// EmailAddrGTE applies the GTE predicate on the "emailAddr" field.
func EmailAddrGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmailAddr, v))
}

// EmailAddrLT applies the LT predicate on the "emailAddr" field.
func EmailAddrLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmailAddr, v))
}

// EmailAddrLTE applies the LTE predicate on the "emailAddr" field.
func EmailAddrLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmailAddr, v))
}

// EmailAddrContains applies the Contains predicate on the "emailAddr" field.
func EmailAddrContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmailAddr, v))
}

// EmailAddrHasPrefix applies the HasPrefix predicate on the "emailAddr" field.
func EmailAddrHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmailAddr, v))
}

// EmailAddrHasSuffix applies the HasSuffix predicate on the "emailAddr" field.
func EmailAddrHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmailAddr, v))
}

// EmailAddrEqualFold applies the EqualFold predicate on the "emailAddr" field.
func EmailAddrEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmailAddr, v))
}

// EmailAddrContainsFold applies the ContainsFold predicate on the "emailAddr" field.
func EmailAddrContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmailAddr, v))
}

// HasPropertylistings applies the HasEdge predicate on the "propertylistings" edge.
func HasPropertylistings() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PropertylistingsTable, PropertylistingsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPropertylistingsWith applies the HasEdge predicate on the "propertylistings" edge with a given conditions (other predicates).
func HasPropertylistingsWith(preds ...predicate.PropertyListing) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPropertylistingsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
