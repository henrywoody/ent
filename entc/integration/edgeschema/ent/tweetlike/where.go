// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package tweetlike

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
)

// LikedAt applies equality check predicate on the "liked_at" field. It's identical to LikedAtEQ.
func LikedAt(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLikedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v int) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTweetID), v))
	})
}

// LikedAtEQ applies the EQ predicate on the "liked_at" field.
func LikedAtEQ(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLikedAt), v))
	})
}

// LikedAtNEQ applies the NEQ predicate on the "liked_at" field.
func LikedAtNEQ(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLikedAt), v))
	})
}

// LikedAtIn applies the In predicate on the "liked_at" field.
func LikedAtIn(vs ...time.Time) predicate.TweetLike {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLikedAt), v...))
	})
}

// LikedAtNotIn applies the NotIn predicate on the "liked_at" field.
func LikedAtNotIn(vs ...time.Time) predicate.TweetLike {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLikedAt), v...))
	})
}

// LikedAtGT applies the GT predicate on the "liked_at" field.
func LikedAtGT(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLikedAt), v))
	})
}

// LikedAtGTE applies the GTE predicate on the "liked_at" field.
func LikedAtGTE(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLikedAt), v))
	})
}

// LikedAtLT applies the LT predicate on the "liked_at" field.
func LikedAtLT(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLikedAt), v))
	})
}

// LikedAtLTE applies the LTE predicate on the "liked_at" field.
func LikedAtLTE(v time.Time) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLikedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.TweetLike {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.TweetLike {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v int) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTweetID), v))
	})
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v int) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTweetID), v))
	})
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...int) predicate.TweetLike {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTweetID), v...))
	})
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...int) predicate.TweetLike {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTweetID), v...))
	})
}

// HasTweet applies the HasEdge predicate on the "tweet" edge.
func HasTweet() predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TweetColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTweetWith applies the HasEdge predicate on the "tweet" edge with a given conditions (other predicates).
func HasTweetWith(preds ...predicate.Tweet) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TweetColumn),
			sqlgraph.To(TweetInverseTable, TweetFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.To(UserInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TweetLike) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TweetLike) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TweetLike) predicate.TweetLike {
	return predicate.TweetLike(func(s *sql.Selector) {
		p(s.Not())
	})
}
