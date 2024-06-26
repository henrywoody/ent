// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package comment

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldPostID holds the string denoting the post_id field in the database.
	FieldPostID = "post_id"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "comments"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_id"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldPostID,
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
