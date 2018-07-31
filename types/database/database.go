package database

import (
	"errors"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// NodeDatabase - database containing list of node addresses, as well as bootstrap addresses
type NodeDatabase struct {
	Nodes *[]node.Node // Nodes - primary list of nodes

	AcceptableTimeout uint // AcceptableTimeout - database-wide definition for operation timeout
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewDatabase - attempts creates new instance of the NodeDatabase struct
func NewDatabase(bootstrapNode *node.Node, acceptableTimeout uint) (NodeDatabase, error) {

	db := NodeDatabase{AcceptableTimeout: acceptableTimeout} // Create empty database with specified timeout

	err := db.AddNode(bootstrapNode) // Attempt to add bootstrapnode

	if err != nil { // Check for errors
		return NodeDatabase{}, err // Return empty node database, error
	}

	return db, nil // No error occurred, return database
}

// AddNode - adds node to specified nodedatabase, after checking address of node
func (db *NodeDatabase) AddNode(node *node.Node) error {
	if !common.CheckAddress(node.Address) { // Check for invalid address
		return errors.New("invalid address") // Return new error
	}

	*db.Nodes = append(*db.Nodes, *node) // Append node to node list

	return nil // No error occurred, return nil
}

// RemoveNode - removes node with specified address from database
func (db *NodeDatabase) RemoveNode(address string) error {
	nodeIndex, err := db.QueryForAddress(address) // Finds index of node with address

	if err != nil { // Checks for error
		return err // Returns error
	}

	db.remove(int(nodeIndex)) // Removes value at index
	return nil                // Returns nil, no error
}

// QueryForAddress - attempts to search specified node database for specified address, returns index of node
func (db *NodeDatabase) QueryForAddress(address string) (uint, error) {
	x := 0

	for x != len(*db.Nodes) { // Wait until entire db has been queried
		if address == (*db.Nodes)[x].Address { // Check for match
			return uint(x), nil // If provided value matches value of node in list, return index
		}
		x++ // Increment index
	}

	return 0, errors.New("no value found") // Could not find index of address, return new error
}

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS:
*/

func (db *NodeDatabase) remove(s int) { // Removes address at index
	*db.Nodes = append((*db.Nodes)[:s], (*db.Nodes)[s+1:]...) // Remove index
}

/*
	END INTERNAL METHODS
*/