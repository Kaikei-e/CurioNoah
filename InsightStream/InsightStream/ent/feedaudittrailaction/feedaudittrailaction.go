// Code generated by ent, DO NOT EDIT.

package feedaudittrailaction

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the feedaudittrailaction type in the database.
	Label = "feed_audit_trail_action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// Table holds the table name of the feedaudittrailaction in the database.
	Table = "feed_audit_trail_actions"
)

// Columns holds all SQL columns for feedaudittrailaction fields.
var Columns = []string{
	FieldID,
	FieldAction,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ActionValidator is a validator for the "action" field. It is called by the builders before save.
	ActionValidator func(string) error
)

// OrderOption defines the ordering options for the FeedAuditTrailAction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}
