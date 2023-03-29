package core

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

/*
	LoadZkBNBInstance: load zkbnb instance if it is already deployed
*/
func LoadZkBNBInstance(cli *rpc.ProviderClient, addr string) (instance *ZkBNB, err error) {
	instance, err = NewZkBNB(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocks(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocksWithNonce(authCli *rpc.AuthClient, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonce(authCli, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	CommitBlocks: commit blocks with kms signature facility
*/
func CommitBlocksWithNonceAndKms(ctx context.Context, kmsSvc *kms.Client, keyId string, chainID *big.Int, address common.Address, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonceAndKms(ctx, kmsSvc, keyId, chainID, address, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocks(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocksWithNonce(authCli *rpc.AuthClient, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonce(authCli, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks with kms signature facility
*/
func VerifyAndExecuteBlocksWithNonceAndKms(ctx context.Context, kmsSvc *kms.Client, keyId string, chainID *big.Int, address common.Address, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonceAndKms(ctx, kmsSvc, keyId, chainID, address, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	RevertBlocks: revert blocks
*/
func RevertBlocks(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	revertBlocks []StorageStoredBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	RevertBlocks: revert blocks with kms signature facility
*/
func RevertBlocksWithKms(
	cli *rpc.ProviderClient, ctx context.Context, kmsSvc *kms.Client, keyId string, chainID *big.Int, address common.Address, instance *ZkBNB,
	revertBlocks []StorageStoredBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithKms(cli, ctx, kmsSvc, keyId, chainID, address, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func PerformDesert(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, storedBlockInfo StorageStoredBlockInfo, nftRoot *big.Int, assetExitData DesertVerifierAssetExitData, accountExitData DesertVerifierAccountExitData,
	assetMerkleProof [16]*big.Int, accountMerkleProof [32]*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.PerformDesert(transactOpts, storedBlockInfo, nftRoot, assetExitData, accountExitData, assetMerkleProof, accountMerkleProof)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func PerformDesertNft(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	storedBlockInfo StorageStoredBlockInfo, assetRoot *big.Int, accountExitData DesertVerifierAccountExitData, exitNfts []DesertVerifierNftExitData, accountMerkleProof [32]*big.Int, nftMerkleProofs [][40]*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.PerformDesertNft(transactOpts, storedBlockInfo, assetRoot, accountExitData, exitNfts, accountMerkleProof, nftMerkleProofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func WithdrawPendingBalance(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, owner common.Address, token common.Address, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.WithdrawPendingBalance(transactOpts, owner, token, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
func WithdrawPendingNFTBalance(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, nftIndex *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.WithdrawPendingNFTBalance(transactOpts, nftIndex)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func CancelOutstandingDepositsForExodusMode(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, priorityRequestId uint64, depositsPubData [][]byte, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CancelOutstandingDepositsForDesertMode(transactOpts, priorityRequestId, depositsPubData)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func ActivateDesertMode(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.ActivateDesertMode(transactOpts)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
