package v1

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterRoles returns a ClusterRoleInformer.
	ClusterRoles() ClusterRoleInformer
	// ClusterRoleBindings returns a ClusterRoleBindingInformer.
	ClusterRoleBindings() ClusterRoleBindingInformer
	// Roles returns a RoleInformer.
	Roles() RoleInformer
	// RoleBindings returns a RoleBindingInformer.
	RoleBindings() RoleBindingInformer
}
