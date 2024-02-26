package rbac

import v1 "github.com/atompi/autom/pkg/util/client-go/informers/rbac/v1"

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() v1.Interface
}
