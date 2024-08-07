// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"insightstream/ent/cooccurrencenetworkpool"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CooccurrenceNetworkPool is the model entity for the CooccurrenceNetworkPool schema.
type CooccurrenceNetworkPool struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SiteURL holds the value of the "site_url" field.
	SiteURL string `json:"site_url,omitempty"`
	// Titles holds the value of the "titles" field.
	Titles []string `json:"titles,omitempty"`
	// Descriptions holds the value of the "descriptions" field.
	Descriptions []string `json:"descriptions,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CooccurrenceNetworkPool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cooccurrencenetworkpool.FieldTitles, cooccurrencenetworkpool.FieldDescriptions:
			values[i] = new([]byte)
		case cooccurrencenetworkpool.FieldSiteURL:
			values[i] = new(sql.NullString)
		case cooccurrencenetworkpool.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CooccurrenceNetworkPool fields.
func (cnp *CooccurrenceNetworkPool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cooccurrencenetworkpool.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cnp.ID = *value
			}
		case cooccurrencenetworkpool.FieldSiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field site_url", values[i])
			} else if value.Valid {
				cnp.SiteURL = value.String
			}
		case cooccurrencenetworkpool.FieldTitles:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field titles", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &cnp.Titles); err != nil {
					return fmt.Errorf("unmarshal field titles: %w", err)
				}
			}
		case cooccurrencenetworkpool.FieldDescriptions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field descriptions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &cnp.Descriptions); err != nil {
					return fmt.Errorf("unmarshal field descriptions: %w", err)
				}
			}
		default:
			cnp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CooccurrenceNetworkPool.
// This includes values selected through modifiers, order, etc.
func (cnp *CooccurrenceNetworkPool) Value(name string) (ent.Value, error) {
	return cnp.selectValues.Get(name)
}

// Update returns a builder for updating this CooccurrenceNetworkPool.
// Note that you need to call CooccurrenceNetworkPool.Unwrap() before calling this method if this CooccurrenceNetworkPool
// was returned from a transaction, and the transaction was committed or rolled back.
func (cnp *CooccurrenceNetworkPool) Update() *CooccurrenceNetworkPoolUpdateOne {
	return NewCooccurrenceNetworkPoolClient(cnp.config).UpdateOne(cnp)
}

// Unwrap unwraps the CooccurrenceNetworkPool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cnp *CooccurrenceNetworkPool) Unwrap() *CooccurrenceNetworkPool {
	_tx, ok := cnp.config.driver.(*txDriver)
	if !ok {
		panic("ent: CooccurrenceNetworkPool is not a transactional entity")
	}
	cnp.config.driver = _tx.drv
	return cnp
}

// String implements the fmt.Stringer.
func (cnp *CooccurrenceNetworkPool) String() string {
	var builder strings.Builder
	builder.WriteString("CooccurrenceNetworkPool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cnp.ID))
	builder.WriteString("site_url=")
	builder.WriteString(cnp.SiteURL)
	builder.WriteString(", ")
	builder.WriteString("titles=")
	builder.WriteString(fmt.Sprintf("%v", cnp.Titles))
	builder.WriteString(", ")
	builder.WriteString("descriptions=")
	builder.WriteString(fmt.Sprintf("%v", cnp.Descriptions))
	builder.WriteByte(')')
	return builder.String()
}

// CooccurrenceNetworkPools is a parsable slice of CooccurrenceNetworkPool.
type CooccurrenceNetworkPools []*CooccurrenceNetworkPool
