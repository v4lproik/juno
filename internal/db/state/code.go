package state

import (
	"fmt"

	"github.com/NethermindEth/juno/internal/db"
	"google.golang.org/protobuf/proto"
)

// GetCode returns the ContractCode associated with the given contract address.
// If the contract code is not found, then error is returned.
func (x *Manager) GetCode(contractAddress []byte) (*Code, error) {
	rawData, err := x.codeDatabase.Get(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("GetCode: failed get operation: %w", err)
	}
	if rawData == nil {
		// notest
		return nil, fmt.Errorf("GetCode: %w", db.ErrNotFound)
	}
	code := new(Code)
	if err := proto.Unmarshal(rawData, code); err != nil {
		return nil, fmt.Errorf("GetCode: %w: %s", db.ErrUnmarshal, err)
	}
	return code, nil
}

// PutCode stores a new contract code into the database, associated with the
// given contract address. If the contract address already have a contract code
// in the database, then the value is updated.
func (x *Manager) PutCode(contractAddress []byte, code *Code) error {
	rawData, err := proto.Marshal(code)
	if err != nil {
		return fmt.Errorf("PutCode: %s: %w", db.ErrMarshal, err)
	}
	if err := x.codeDatabase.Put(contractAddress, rawData); err != nil {
		return fmt.Errorf("PutCode: failed put operation: %w", err)
	}
	return nil
}
