package user

// Info describes a user that has been authenticated to the system.
type Info interface {
	// GetName returns the name that uniquely identifies this user among all
	// other active users.
	GetName() string
	// GetUID returns a unique value for a particular user that will change
	// if the user is removed from the system and another user is added with
	// the same name.
	GetUID() string
	// GetGroups returns the names of the groups the user is a member of
	GetGroups() []string

	// GetExtra can contain any additional information that the authenticator
	// thought was interesting.  One example would be scopes on a token.
	// Keys in this map should be namespaced to the authenticator or
	// authenticator/authorizer pair making use of them.
	// For instance: "example.org/foo" instead of "foo"
	// This is a map[string][]string because it needs to be serializeable into
	// a SubjectAccessReviewSpec.authorization.k8s.io for proper authorization
	// delegation flows
	// In order to faithfully round-trip through an impersonation flow, these keys
	// MUST be lowercase.
	GetExtra() map[string][]string
}

// well-known user and group names
const (
	SystemPrivilegedGroup = "system:masters"
	NodesGroup            = "system:nodes"
	MonitoringGroup       = "system:monitoring"
	AllUnauthenticated    = "system:unauthenticated"
	AllAuthenticated      = "system:authenticated"

	Anonymous     = "system:anonymous"
	APIServerUser = "system:apiserver"
)
