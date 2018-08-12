package cli

import (
	"errors"
	"fmt"

	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// handleNewDatabaseCommand - handle execution of handleNewDatabase method (wrapper)
func (term *Terminal) handleNewDatabaseCommand() {
	fmt.Println("attempting to initialize new NodeDatabase") // Log begin

	output, err := term.handleNewDatabase() // Attempt to init new db

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAddNodeCommand - handle execution of handleAddNode method (wrapper)
func (term *Terminal) handleAddNodeCommand(address string) {
	fmt.Println("attempting to add current node to database") // Log begin

	output, err := term.handleAddNode(address) // Attempt to append

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAttachDatabaseCommand - handle execution of handleAttachDatabase method (wrapper)
func (term *Terminal) handleAttachDatabaseCommand() {
	fmt.Println("attempting to attach to NodeDatabase") // Log begin

	output, err := term.handleAttachDatabase() // Attempt to attach to db

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleWriteDatabaseToMemoryCommand - handle execution of handleWritDatabaseToMemory method (wrapper)
func (term *Terminal) handleWriteDatabaseToMemoryCommand() {
	fmt.Println("attempting to write database to memory") // Log begin

	output, err := term.handleWriteDatabaseToMemory() // Attempt to write db

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleNewDatabase - attempt to initialize new NodeDatabase
func (term *Terminal) handleNewDatabase() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := database.NewDatabase(&foundNode, 5) // Attempt to create new database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(db, "NodeDatabase") // Add new database

	err = db.WriteToMemory(foundNode.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- created new nodedatabase with address " + foundNode.Address, nil // No error occurred, return success
}

// handleAddNode - attempt to append current node to NodeDatabase
func (term *Terminal) handleAddNode(address string) (string, error) {
	if address != "" {
		return term.handleAddSpecificNode(address)
	}

	return term.handleAddCurrentNode()
}

func (term *Terminal) handleAddSpecificNode(address string) (string, error) {
	foundDb := database.NodeDatabase{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "NodeDatabase" { // Verify element is NodeDatabase
			foundDb = term.Variables[x].(database.NodeDatabase) // Set to value

			break
		}
	}

	_, err := foundDb.QueryForAddress(address)

	if err == nil {
		return "", errors.New("node already added to database")
	}

	newNode, err := node.NewNode(address, false) // Attempt to init node with specified address

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundDb.AddNode(&newNode) // Attempt to add node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- added node with address " + address + " to attached node database", nil // Return success
}

// handleAddCurrentNode - attempt to add current node to attached node database
func (term *Terminal) handleAddCurrentNode() (string, error) {
	foundNode := node.Node{}           // Create placeholder
	foundDb := database.NodeDatabase{} // Create placeholder

	emptyNode := node.Node{}
	emptyDb := database.NodeDatabase{}

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value

			if foundDb != emptyDb { // Check for valid db
				break
			}
		} else if term.VariableTypes[x] == "NodeDatabase" { // Verify element is NodeDatabase
			foundDb = term.Variables[x].(database.NodeDatabase) // Set to value

			if foundNode != emptyNode { // Check for valid node
				break
			}
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	_, qErr := foundDb.QueryForAddress(foundNode.Address) // Check for already existing node

	if qErr != nil { // Check for already existing node
		err := foundDb.AddNode(&foundNode) // Attempt to add node

		if err != nil { // Check for errors
			return "", err // Return found error
		}
	} else { // Node already exists, return error
		return "", errors.New("node already exists in attached database") // Return found error
	}

	return "-- SUCCESS -- appended node with address " + foundNode.Address + " to NodeDatabase", nil // No error occurred, return success
}

// handleAttachDatabase - handle execution of database reading, write to term mem
func (term *Terminal) handleAttachDatabase() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := database.ReadDatabaseFromMemory(foundNode.Environment) // Attempt to read database from node environment memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = term.AddVariable(*db, "NodeDatabase") // Save for persistency

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- attached to nodedatabase with bootstrap address " + (*db.Nodes)[0].Address, nil
}

// handleWritDatabaseToMemory - handle execution of NodeDatabase writeToMemory() method
func (term *Terminal) handleWriteDatabaseToMemory() (string, error) {
	foundNode := node.Node{}           // Create placeholder
	foundDb := database.NodeDatabase{} // Create placeholder

	emptyNode := node.Node{}
	emptyDb := database.NodeDatabase{}

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value

			if foundDb != emptyDb { // Check for valid db
				break
			}
		} else if term.VariableTypes[x] == "NodeDatabase" { // Verify element is NodeDatabase
			foundDb = term.Variables[x].(database.NodeDatabase) // Set to value

			if foundNode != emptyNode { // Check for valid node
				break
			}
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	err := foundDb.WriteToMemory(foundNode.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- wrote nodedatabase with address " + foundNode.Address + " to memory", nil // No error occurred, return success
}