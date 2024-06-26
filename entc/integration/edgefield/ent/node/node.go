// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package node

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldPrevID holds the string denoting the prev_id field in the database.
	FieldPrevID = "prev_id"
	// EdgePrev holds the string denoting the prev edge name in mutations.
	EdgePrev = "prev"
	// EdgeNext holds the string denoting the next edge name in mutations.
	EdgeNext = "next"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// PrevTable is the table that holds the prev relation/edge.
	PrevTable = "nodes"
	// PrevColumn is the table column denoting the prev relation/edge.
	PrevColumn = "prev_id"
	// NextTable is the table that holds the next relation/edge.
	NextTable = "nodes"
	// NextColumn is the table column denoting the next relation/edge.
	NextColumn = "prev_id"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldValue,
	FieldPrevID,
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
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue int
)
