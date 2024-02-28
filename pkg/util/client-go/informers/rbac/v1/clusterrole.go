package v1

import (
	"github.com/atompi/autom/pkg/util/apimachinery/labels"
	"github.com/atompi/autom/pkg/util/client-go/tools/cache"
)

// ClusterRoleLister helps list ClusterRoles.
// All objects returned here must be treated as read-only.
type ClusterRoleLister interface {
	// List lists all ClusterRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterRole, err error)
	// Get retrieves the ClusterRole from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterRole, error)
	ClusterRoleListerExpansion
}

// ClusterRoleInformer provides access to a shared informer and lister for
// ClusterRoles.
type ClusterRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterRoleLister
}
