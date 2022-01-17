// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/google/uuid"
)

// Kyc is the model entity for the Kyc schema.
type Kyc struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// CardType holds the value of the "card_type" field.
	CardType string `json:"card_type,omitempty"`
	// CardID holds the value of the "card_id" field.
	CardID string `json:"card_id,omitempty"`
	// FrontCardImg holds the value of the "front_card_img" field.
	FrontCardImg string `json:"front_card_img,omitempty"`
	// BackCardImg holds the value of the "back_card_img" field.
	BackCardImg string `json:"back_card_img,omitempty"`
	// UserHandlingCardImg holds the value of the "user_handling_card_img" field.
	UserHandlingCardImg string `json:"user_handling_card_img,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Kyc) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kyc.FieldCreateAt, kyc.FieldUpdateAt:
			values[i] = new(sql.NullInt64)
		case kyc.FieldCardType, kyc.FieldCardID, kyc.FieldFrontCardImg, kyc.FieldBackCardImg, kyc.FieldUserHandlingCardImg:
			values[i] = new(sql.NullString)
		case kyc.FieldID, kyc.FieldUserID, kyc.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Kyc", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Kyc fields.
func (k *Kyc) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kyc.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				k.ID = *value
			}
		case kyc.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				k.UserID = *value
			}
		case kyc.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				k.AppID = *value
			}
		case kyc.FieldCardType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field card_type", values[i])
			} else if value.Valid {
				k.CardType = value.String
			}
		case kyc.FieldCardID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field card_id", values[i])
			} else if value.Valid {
				k.CardID = value.String
			}
		case kyc.FieldFrontCardImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field front_card_img", values[i])
			} else if value.Valid {
				k.FrontCardImg = value.String
			}
		case kyc.FieldBackCardImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field back_card_img", values[i])
			} else if value.Valid {
				k.BackCardImg = value.String
			}
		case kyc.FieldUserHandlingCardImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_handling_card_img", values[i])
			} else if value.Valid {
				k.UserHandlingCardImg = value.String
			}
		case kyc.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				k.CreateAt = uint32(value.Int64)
			}
		case kyc.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				k.UpdateAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Kyc.
// Note that you need to call Kyc.Unwrap() before calling this method if this Kyc
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Kyc) Update() *KycUpdateOne {
	return (&KycClient{config: k.config}).UpdateOne(k)
}

// Unwrap unwraps the Kyc entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Kyc) Unwrap() *Kyc {
	tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Kyc is not a transactional entity")
	}
	k.config.driver = tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Kyc) String() string {
	var builder strings.Builder
	builder.WriteString("Kyc(")
	builder.WriteString(fmt.Sprintf("id=%v", k.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", k.UserID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", k.AppID))
	builder.WriteString(", card_type=")
	builder.WriteString(k.CardType)
	builder.WriteString(", card_id=")
	builder.WriteString(k.CardID)
	builder.WriteString(", front_card_img=")
	builder.WriteString(k.FrontCardImg)
	builder.WriteString(", back_card_img=")
	builder.WriteString(k.BackCardImg)
	builder.WriteString(", user_handling_card_img=")
	builder.WriteString(k.UserHandlingCardImg)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", k.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", k.UpdateAt))
	builder.WriteByte(')')
	return builder.String()
}

// Kycs is a parsable slice of Kyc.
type Kycs []*Kyc

func (k Kycs) config(cfg config) {
	for _i := range k {
		k[_i].config = cfg
	}
}
