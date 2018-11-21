package proto

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/mitsukomegumi/GoP2P/common"
)

/* BEGIN EXPORTED METHODS */

// ReadGuideFromMemory - read protobufGuide from memory
func ReadGuideFromMemory(path string) (*ProtobufGuide, error) {
	protoGuide := &ProtobufGuide{} // Init guide buffer

	err := common.ReadGob(fmt.Sprintf("%s.goP2PGuide", path), protoGuide) // Read protoGuide

	if err != nil { // Check for errors
		return nil, err // Return error
	}

	return protoGuide, nil // No error occurred, return nil
}

// WriteToMemory - write protobufGuide to memory
func (protoGuide *ProtobufGuide) WriteToMemory(path string) error {
	file, err := os.Create(path) // Attempt to create file at path

	if err != nil { // Check for errors
		return err // Return found error
	}

	encoder := gob.NewEncoder(file) // Init file writer

	err = encoder.Encode(*protoGuide) // Encode object

	if err != nil { // Check for errors
		return err // Return found error
	}

	return file.Close() // Return output of file.Close()
}

/* END EXPORTED METHODS */
