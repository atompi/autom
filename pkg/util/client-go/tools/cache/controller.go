package cache

// ResourceEventHandler can handle notifications for events that
// happen to a resource. The events are informational only, so you
// can't return an error.  The handlers MUST NOT modify the objects
// received; this concerns not only the top level of structure but all
// the data structures reachable from it.
//   - OnAdd is called when an object is added.
//   - OnUpdate is called when an object is modified. Note that oldObj is the
//     last known state of the object-- it is possible that several changes
//     were combined together, so you can't use this to see every single
//     change. OnUpdate is also called when a re-list happens, and it will
//     get called even if nothing changed. This is useful for periodically
//     evaluating or syncing something.
//   - OnDelete will get the final state of the item if it is known, otherwise
//     it will get an object of type DeletedFinalStateUnknown. This can
//     happen if the watch is closed and misses the delete event and we don't
//     notice the deletion until the subsequent re-list.
type ResourceEventHandler interface {
	OnAdd(obj interface{}, isInInitialList bool)
	OnUpdate(oldObj, newObj interface{})
	OnDelete(obj interface{})
}

// Controller is a low-level controller that is parameterized by a
// Config and used in sharedIndexInformer.
type Controller interface {
	// Run does two things.  One is to construct and run a Reflector
	// to pump objects/notifications from the Config's ListerWatcher
	// to the Config's Queue and possibly invoke the occasional Resync
	// on that Queue.  The other is to repeatedly Pop from the Queue
	// and process with the Config's ProcessFunc.  Both of these
	// continue until `stopCh` is closed.
	Run(stopCh <-chan struct{})

	// HasSynced delegates to the Config's Queue
	HasSynced() bool

	// LastSyncResourceVersion delegates to the Reflector when there
	// is one, otherwise returns the empty string
	LastSyncResourceVersion() string
}
