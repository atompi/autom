package authorizerfactory

import (
	"context"
	"errors"

	"github.com/atompi/autom/pkg/auto-apiserver/authentication/user"
	"github.com/atompi/autom/pkg/auto-apiserver/authorization/authorizer"
)

// alwaysAllowAuthorizer is an implementation of authorizer.Attributes
// which always says yes to an authorization request.
// It is useful in tests and when using autom in an open manner.
type alwaysAllowAuthorizer struct{}

func (alwaysAllowAuthorizer) Authorize(ctx context.Context, a authorizer.Attributes) (authorized authorizer.Decision, reason string, err error) {
	return authorizer.DecisionAllow, "", nil
}

func (alwaysAllowAuthorizer) RulesFor(user user.Info, namespace string) ([]authorizer.ResourceRuleInfo, []authorizer.NonResourceRuleInfo, bool, error) {
	return []authorizer.ResourceRuleInfo{
			&authorizer.DefaultResourceRuleInfo{
				Verbs:     []string{"*"},
				APIGroups: []string{"*"},
				Resources: []string{"*"},
			},
		}, []authorizer.NonResourceRuleInfo{
			&authorizer.DefaultNonResourceRuleInfo{
				Verbs:           []string{"*"},
				NonResourceURLs: []string{"*"},
			},
		}, false, nil
}

func NewAlwaysAllowAuthorizer() *alwaysAllowAuthorizer {
	return new(alwaysAllowAuthorizer)
}

// alwaysDenyAuthorizer is an implementation of authorizer.Attributes
// which always says no to an authorization request.
// It is useful in unit tests to force an operation to be forbidden.
type alwaysDenyAuthorizer struct{}

func (alwaysDenyAuthorizer) Authorize(ctx context.Context, a authorizer.Attributes) (decision authorizer.Decision, reason string, err error) {
	return authorizer.DecisionNoOpinion, "Everything is forbidden.", nil
}

func (alwaysDenyAuthorizer) RulesFor(user user.Info, namespace string) ([]authorizer.ResourceRuleInfo, []authorizer.NonResourceRuleInfo, bool, error) {
	return []authorizer.ResourceRuleInfo{}, []authorizer.NonResourceRuleInfo{}, false, nil
}

func NewAlwaysDenyAuthorizer() *alwaysDenyAuthorizer {
	return new(alwaysDenyAuthorizer)
}

type privilegedGroupAuthorizer struct {
	groups []string
}

func (r *privilegedGroupAuthorizer) Authorize(ctx context.Context, attr authorizer.Attributes) (authorizer.Decision, string, error) {
	if attr.GetUser() == nil {
		return authorizer.DecisionNoOpinion, "Error", errors.New("no user on request")
	}
	for _, attr_group := range attr.GetUser().GetGroups() {
		for _, priv_group := range r.groups {
			if priv_group == attr_group {
				return authorizer.DecisionAllow, "", nil
			}
		}
	}
	return authorizer.DecisionNoOpinion, "", nil
}

// NewPrivilegedGroups is for use in loopback scenarios
func NewPrivilegedGroups(groups ...string) *privilegedGroupAuthorizer {
	return &privilegedGroupAuthorizer{
		groups: groups,
	}
}
