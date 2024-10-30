package eth_test

import (
	"crypto/ecdsa"
	"filestore/pkg/eth"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

const (
	testURL          = "http://localhost:8545"
	testContractAddr = "e7f1725e7734ce288f8367e1bb143e90bb3f0512"
	testFileHash     = "testFileHash"
	testIPFSHash     = "testIPFSHash"
	testFileName     = "testFileName"
	Contract_Address = "e7f1725e7734ce288f8367e1bb143e90bb3f0512"
)

func TestConnectToClient(t *testing.T) {
	client, err := eth.ConnectToClient(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestGetContractInstance(t *testing.T) {
	client, err := eth.ConnectToClient(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	instance, err := eth.GetContractInstance(client, Contract_Address)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestStoreFile(t *testing.T) {
	client, err := eth.ConnectToClient(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	instance, err := eth.GetContractInstance(client, Contract_Address)
	assert.NoError(t, err)
	assert.NotNil(t, instance)

	auth, err := eth.GetAuth(client)
	assert.NoError(t, err)
	assert.NotNil(t, auth)

	txHash, err := eth.StoreFile(instance, auth, testFileHash, testIPFSHash, testFileName)
	assert.NoError(t, err)
	assert.NotEmpty(t, txHash)
}

func TestGetUserFileCount(t *testing.T) {
	client, err := eth.ConnectToClient(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	instance, err := eth.GetContractInstance(client, testContractAddr)
	assert.NoError(t, err)
	assert.NotNil(t, instance)

	count, err := eth.GetUserFileCount(instance)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.True(t, count.Cmp(big.NewInt(0)) >= 0)
}

func TestGetAuth(t *testing.T) {
	client, err := eth.ConnectToClient(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	auth, err := eth.GetAuth(client)
	assert.NoError(t, err)
	assert.NotNil(t, auth)

	privateKey, err := crypto.HexToECDSA(eth.Private_Key)
	assert.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	assert.True(t, ok)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	assert.Equal(t, fromAddress, auth.From)
}
