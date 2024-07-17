// Code generated by ent, DO NOT EDIT.

package feedaudittrailaction

import (
	"insightstream/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldLTE(FieldID, id))
}

// Action applies equality check predicate on the "action" field. It's identical to ActionEQ.
func Action(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldEQ(FieldAction, v))
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldEQ(FieldAction, v))
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldNEQ(FieldAction, v))
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldIn(FieldAction, vs...))
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldNotIn(FieldAction, vs...))
}

// ActionGT applies the GT predicate on the "action" field.
func ActionGT(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldGT(FieldAction, v))
}

// ActionGTE applies the GTE predicate on the "action" field.
func ActionGTE(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldGTE(FieldAction, v))
}

// ActionLT applies the LT predicate on the "action" field.
func ActionLT(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldLT(FieldAction, v))
}

// ActionLTE applies the LTE predicate on the "action" field.
func ActionLTE(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldLTE(FieldAction, v))
}

// ActionContains applies the Contains predicate on the "action" field.
func ActionContains(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldContains(FieldAction, v))
}

// ActionHasPrefix applies the HasPrefix predicate on the "action" field.
func ActionHasPrefix(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldHasPrefix(FieldAction, v))
}

// ActionHasSuffix applies the HasSuffix predicate on the "action" field.
func ActionHasSuffix(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldHasSuffix(FieldAction, v))
}

// ActionEqualFold applies the EqualFold predicate on the "action" field.
func ActionEqualFold(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldEqualFold(FieldAction, v))
}

// ActionContainsFold applies the ContainsFold predicate on the "action" field.
func ActionContainsFold(v string) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.FieldContainsFold(FieldAction, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeedAuditTrailAction) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeedAuditTrailAction) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FeedAuditTrailAction) predicate.FeedAuditTrailAction {
	return predicate.FeedAuditTrailAction(sql.NotPredicates(p))
}
