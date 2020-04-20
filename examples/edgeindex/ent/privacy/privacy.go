// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/edgeindex/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	for _, rule := range policy {
		switch err := rule.EvalQuery(ctx, q); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	for _, rule := range policy {
		switch err := rule.EvalMutation(ctx, m); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecisionRule{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecisionRule{Deny}
}

type fixedDecisionRule struct{ err error }

func (f fixedDecisionRule) EvalQuery(context.Context, ent.Query) error       { return f.err }
func (f fixedDecisionRule) EvalMutation(context.Context, ent.Mutation) error { return f.err }

// DenyMutationOperation returns a rule denying specifies mutation operation.
func DenyMutationOperation(op ent.Op) MutationRule {
	return MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return Denyf("ent/privacy: operation %s is not allowed", m.Op())
		}
		return Skip
	})
}

// The CityQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CityQueryRuleFunc func(context.Context, *ent.CityQuery) error

// EvalQuery return f(ctx, q).
func (f CityQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CityQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CityQuery", q)
}

// The CityMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CityMutationRuleFunc func(context.Context, *ent.CityMutation) error

// EvalMutation calls f(ctx, m).
func (f CityMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CityMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CityMutation", m)
}

// The StreetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StreetQueryRuleFunc func(context.Context, *ent.StreetQuery) error

// EvalQuery return f(ctx, q).
func (f StreetQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StreetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StreetQuery", q)
}

// The StreetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StreetMutationRuleFunc func(context.Context, *ent.StreetMutation) error

// EvalMutation calls f(ctx, m).
func (f StreetMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StreetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StreetMutation", m)
}