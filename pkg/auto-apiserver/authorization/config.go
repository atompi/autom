package authorization

import (
	"github.com/atompi/autom/pkg/auto-apiserver/authentication/user"
	"github.com/atompi/autom/pkg/auto-apiserver/authorization/authorizer"
	"github.com/atompi/autom/pkg/auto-apiserver/authorization/authorizerfactory"
	versionedinformers "github.com/atompi/autom/pkg/util/client-go/informers"
	"github.com/atompi/autom/pkg/util/client-go/informers/rbac"
)

type Config struct {
	VersionedInformerFactory versionedinformers.SharedInformerFactory
}

func (config Config) New() (authorizer.Authorizer, authorizer.RuleResolver, error) {
	var (
		authorizers   []authorizer.Authorizer
		ruleResolvers []authorizer.RuleResolver
	)

	// Add SystemPrivilegedGroup as an authorizing group
	superuserAuthorizer := authorizerfactory.NewPrivilegedGroups(user.SystemPrivilegedGroup)
	authorizers = append(authorizers, superuserAuthorizer)

	rbacAuthorizer := rbac.New(
		&rbac.RoleGetter{Lister: config.VersionedInformerFactory.Rbac().V1().Roles().Lister()},
		&rbac.RoleBindingLister{Lister: config.VersionedInformerFactory.Rbac().V1().RoleBindings().Lister()},
		&rbac.ClusterRoleGetter{Lister: config.VersionedInformerFactory.Rbac().V1().ClusterRoles().Lister()},
		&rbac.ClusterRoleBindingLister{Lister: config.VersionedInformerFactory.Rbac().V1().ClusterRoleBindings().Lister()},
	)
	authorizers = append(authorizers, rbacAuthorizer)
	ruleResolvers = append(ruleResolvers, rbacAuthorizer)
}
