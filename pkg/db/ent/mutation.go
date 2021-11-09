// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"entgo.io/ent"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEmpty = "Empty"
)

// EmptyMutation represents an operation that mutates the Empty nodes in the graph.
type EmptyMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Empty, error)
	predicates    []predicate.Empty
}

var _ ent.Mutation = (*EmptyMutation)(nil)

// emptyOption allows management of the mutation configuration using functional options.
type emptyOption func(*EmptyMutation)

// newEmptyMutation creates new mutation for the Empty entity.
func newEmptyMutation(c config, op Op, opts ...emptyOption) *EmptyMutation {
	m := &EmptyMutation{
		config:        c,
		op:            op,
		typ:           TypeEmpty,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEmptyID sets the ID field of the mutation.
func withEmptyID(id int) emptyOption {
	return func(m *EmptyMutation) {
		var (
			err   error
			once  sync.Once
			value *Empty
		)
		m.oldValue = func(ctx context.Context) (*Empty, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Empty.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEmpty sets the old Empty of the mutation.
func withEmpty(node *Empty) emptyOption {
	return func(m *EmptyMutation) {
		m.oldValue = func(context.Context) (*Empty, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EmptyMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EmptyMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EmptyMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// Where appends a list predicates to the EmptyMutation builder.
func (m *EmptyMutation) Where(ps ...predicate.Empty) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *EmptyMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Empty).
func (m *EmptyMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EmptyMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EmptyMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EmptyMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Empty field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmptyMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Empty field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EmptyMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EmptyMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmptyMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Empty numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EmptyMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EmptyMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EmptyMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Empty nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EmptyMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Empty field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EmptyMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EmptyMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EmptyMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EmptyMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EmptyMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EmptyMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EmptyMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Empty unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EmptyMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Empty edge %s", name)
}
