package store

import (
	"github.com/pokt-network/pocket-core/store/types"
)

// Import pocket-core/types/store.go for convenience.
// nolint
type (
	PruningOptions   = types.PruningOptions
	Store            = types.Store
	Committer        = types.Committer
	CommitStore      = types.CommitStore
	MultiStore       = types.MultiStore
	CacheMultiStore  = types.CacheMultiStore
	CommitMultiStore = types.CommitMultiStore
	KVStore          = types.KVStore
	KVPair           = types.KVPair
	Iterator         = types.Iterator
	CacheKVStore     = types.CacheKVStore
	CommitKVStore    = types.CommitKVStore
	CacheWrapper     = types.CacheWrapper
	CacheWrap        = types.CacheWrap
	CommitID         = types.CommitID
	StoreKey         = types.StoreKey
	StoreType        = types.StoreType
)

// nolint - reexport
var (
	PruneNothing    = types.PruneNothing
	PruneEverything = types.PruneEverything
	PruneSyncable   = types.PruneSyncable
)
