package splitstore

import (
	"fmt"
	"path/filepath"

	"github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
)

type TrackingStore interface {
	Put(cid.Cid, abi.ChainEpoch) error
	PutBatch([]cid.Cid, abi.ChainEpoch) error
	Get(cid.Cid) (abi.ChainEpoch, error)
	Delete(cid.Cid) error
	ForEach(func(cid.Cid, abi.ChainEpoch) error) error
	Close() error
}

func NewTrackingStore(path string, useLMDB bool) (TrackingStore, error) {
	if useLMDB {
		return NewLMDBTrackingStore(filepath.Join(path, "snoop.lmdb"))
	}

	return nil, fmt.Errorf("TODO: non-lmdb livesets")
}