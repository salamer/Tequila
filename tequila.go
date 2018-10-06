package Tequila

import (
	"github.com/boltdb/bolt"
	"github.com/etcd-io/etcd/raft"
)

//Tequila : base Tequila client
type Tequila struct {
	boltConn *bolt.DB
	raftNode *raft.Node
}
