package cache

// TransformFunc allows for transforming an object before it will be processed.
// TransformFunc (similarly to ResourceEventHandler functions) should be able
// to correctly handle the tombstone of type cache.DeletedFinalStateUnknown.
//
// New in v1.27: In such cases, the contained object will already have gone
// through the transform object separately (when it was added / updated prior
// to the delete), so the TransformFunc can likely safely ignore such objects
// (i.e., just return the input object).
//
// The most common usage pattern is to clean-up some parts of the object to
// reduce component memory usage if a given component doesn't care about them.
//
// New in v1.27: unless the object is a DeletedFinalStateUnknown, TransformFunc
// sees the object before any other actor, and it is now safe to mutate the
// object in place instead of making a copy.
//
// Note that TransformFunc is called while inserting objects into the
// notification queue and is therefore extremely performance sensitive; please
// do not do anything that will take a long time.
type TransformFunc func(interface{}) (interface{}, error)
