package eth

import (
	"context"
	"crypto/ecdsa"
	contract "filestore/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const Private_Key = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

// ConnectToClient connects to the Ethereum client
func ConnectToClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetContractInstance returns a new contract instance
func GetContractInstance(client *ethclient.Client, address string) (*contract.Contract, error) {
	contractAddress := common.HexToAddress(address)
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// StoreFile stores a file in the contract
func StoreFile(instance *contract.Contract, auth *bind.TransactOpts, fileHash, ipfsHash, fileName string) (string, error) {
	tx, err := instance.StoreFile(auth, fileHash, ipfsHash, fileName)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// GetUserFileCount gets the user file count from the contract
func GetUserFileCount(instance *contract.Contract) (*big.Int, error) {
	count, err := instance.GetUserFileCount(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, err
	}
	return count, nil
}

// GetFileByHash retrieves file information by its hash
func GetFileByHash(instance *contract.Contract, fileHash string) (string, string, string, *big.Int, error) {
	fileHashBytes := [32]byte{}
	copy(fileHashBytes[:], fileHash)

	fileHash, ipfsHash, fileName, timestamp, err := instance.GetFileByHash(&bind.CallOpts{Context: context.Background()}, fileHashBytes)
	if err != nil {
		return "", "", "", nil, err
	}

	return fileHash, ipfsHash, fileName, timestamp, nil
}
func GetAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	// Replace with your Ethereum private key
	privateKey, err := crypto.HexToECDSA(Private_Key)
	if err != nil {
		return nil, err
	}

	// Get the public address of the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the nonce (the number of transactions sent from this address)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// Suggest gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// Specify the chain ID (e.g., 1 for Mainnet, 1337 for Ganache)
	chainID := big.NewInt(1337) // Change this to your network's chain ID

	// Create the transaction options (transactor)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce)) // The nonce
	auth.Value = big.NewInt(0)            // In wei (0 means no funds sent)
	auth.GasLimit = uint64(3000000)       // Gas limit
	auth.GasPrice = gasPrice              // Gas price in wei
	auth.Context = context.Background()   // Set the context for cancellation

	return auth, nil
}
