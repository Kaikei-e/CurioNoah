// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"insightstream/ent/feedaudittrail"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// FeedAuditTrail is the model entity for the FeedAuditTrail schema.
type FeedAuditTrail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeedAuditTrail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feedaudittrail.FieldID:
			values[i] = new(sql.NullInt64)
		case feedaudittrail.FieldAction:
			values[i] = new(sql.NullString)
		case feedaudittrail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FeedAuditTrail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeedAuditTrail fields.
func (fat *FeedAuditTrail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feedaudittrail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fat.ID = int(value.Int64)
		case feedaudittrail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fat.UpdatedAt = value.Time
			}
		case feedaudittrail.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				fat.Action = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FeedAuditTrail.
// Note that you need to call FeedAuditTrail.Unwrap() before calling this method if this FeedAuditTrail
// was returned from a transaction, and the transaction was committed or rolled back.
func (fat *FeedAuditTrail) Update() *FeedAuditTrailUpdateOne {
	return (&FeedAuditTrailClient{config: fat.config}).UpdateOne(fat)
}

// Unwrap unwraps the FeedAuditTrail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fat *FeedAuditTrail) Unwrap() *FeedAuditTrail {
	_tx, ok := fat.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeedAuditTrail is not a transactional entity")
	}
	fat.config.driver = _tx.drv
	return fat
}

// String implements the fmt.Stringer.
func (fat *FeedAuditTrail) String() string {
	var builder strings.Builder
	builder.WriteString("FeedAuditTrail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fat.ID))
	builder.WriteString("updated_at=")
	builder.WriteString(fat.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("action=")
	builder.WriteString(fat.Action)
	builder.WriteByte(')')
	return builder.String()
}

// FeedAuditTrails is a parsable slice of FeedAuditTrail.
type FeedAuditTrails []*FeedAuditTrail

func (fat FeedAuditTrails) config(cfg config) {
	for _i := range fat {
		fat[_i].config = cfg
	}
}
