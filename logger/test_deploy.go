package logger

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// TestDeployLogger Test logger contract gets deployed correctly
func TestDeployLogger(t *testing.T) {

	//Setup simulated block chain
	privateKey, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(privateKey)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	ethClient := backends.NewSimulatedBackend(alloc, 800000)

	//Deploy contract
	address, _, _, err := DeployLogger(auth, ethClient)

	// commit all pending transactions
	ethClient.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the Logger contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

}
