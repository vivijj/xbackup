// Code generated by ethgo/abigen. DO NOT EDIT.
// Hash: b03b1d179ae319644cce116e6033536f35599ee2d98c99bceabf0c8b0b7cea10
// Version: 0.1.3
package main

import (
	"fmt"
	"math/big"

	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
)

var (
	_ = big.NewInt
	_ = jsonrpc.NewClient
)

// Xbackup is a solidity contract
type Xbackup struct {
	c *contract.Contract
}

// NewXbackup creates a new instance of the contract at a specific address
func NewXbackup(addr ethgo.Address, opts ...contract.ContractOption) *Xbackup {
	return &Xbackup{c: contract.NewContract(addr, abiXbackup, opts...)}
}

// calls

// AccessLinks calls the accessLinks method in the solidity contract
func (x *Xbackup) AccessLinks(val0 string, block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("accessLinks", ethgo.EncodeBlock(block...), val0)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// BalanceOf calls the balanceOf method in the solidity contract
func (x *Xbackup) BalanceOf(owner ethgo.Address, block ...ethgo.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("balanceOf", ethgo.EncodeBlock(block...), owner)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// FreeTokenId calls the freeTokenId method in the solidity contract
func (x *Xbackup) FreeTokenId(block ...ethgo.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("freeTokenId", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// GetApproved calls the getApproved method in the solidity contract
func (x *Xbackup) GetApproved(tokenId *big.Int, block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("getApproved", ethgo.EncodeBlock(block...), tokenId)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Infohash2tokenid calls the infohash2tokenid method in the solidity contract
func (x *Xbackup) Infohash2tokenid(val0 string, block ...ethgo.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("infohash2tokenid", ethgo.EncodeBlock(block...), val0)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// IsApprovedForAll calls the isApprovedForAll method in the solidity contract
func (x *Xbackup) IsApprovedForAll(owner ethgo.Address, operator ethgo.Address, block ...ethgo.BlockNumber) (retval0 bool, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("isApprovedForAll", ethgo.EncodeBlock(block...), owner, operator)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(bool)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Name calls the name method in the solidity contract
func (x *Xbackup) Name(block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("name", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Operator calls the operator method in the solidity contract
func (x *Xbackup) Operator(block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("operator", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// OwnerOf calls the ownerOf method in the solidity contract
func (x *Xbackup) OwnerOf(tokenId *big.Int, block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("ownerOf", ethgo.EncodeBlock(block...), tokenId)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// SupportsInterface calls the supportsInterface method in the solidity contract
func (x *Xbackup) SupportsInterface(interfaceId [4]byte, block ...ethgo.BlockNumber) (retval0 bool, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("supportsInterface", ethgo.EncodeBlock(block...), interfaceId)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(bool)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Symbol calls the symbol method in the solidity contract
func (x *Xbackup) Symbol(block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("symbol", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// TokenURI calls the tokenURI method in the solidity contract
func (x *Xbackup) TokenURI(tokenId *big.Int, block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = x.c.Call("tokenURI", ethgo.EncodeBlock(block...), tokenId)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// txns

// Approve sends a approve transaction in the solidity contract
func (x *Xbackup) Approve(to ethgo.Address, tokenId *big.Int) (contract.Txn, error) {
	return x.c.Txn("approve", to, tokenId)
}

// MintTokenForDataAuthor sends a mintTokenForDataAuthor transaction in the solidity contract
func (x *Xbackup) MintTokenForDataAuthor(infoHash string, tokenUri string, link string, author ethgo.Address) (contract.Txn, error) {
	return x.c.Txn("mintTokenForDataAuthor", infoHash, tokenUri, link, author)
}

// SafeTransferFrom sends a safeTransferFrom transaction in the solidity contract
func (x *Xbackup) SafeTransferFrom(from ethgo.Address, to ethgo.Address, tokenId *big.Int) (contract.Txn, error) {
	return x.c.Txn("safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom0 sends a safeTransferFrom0 transaction in the solidity contract
func (x *Xbackup) SafeTransferFrom0(from ethgo.Address, to ethgo.Address, tokenId *big.Int, data []byte) (contract.Txn, error) {
	return x.c.Txn("safeTransferFrom0", from, to, tokenId, data)
}

// SetApprovalForAll sends a setApprovalForAll transaction in the solidity contract
func (x *Xbackup) SetApprovalForAll(operator ethgo.Address, approved bool) (contract.Txn, error) {
	return x.c.Txn("setApprovalForAll", operator, approved)
}

// TransferFrom sends a transferFrom transaction in the solidity contract
func (x *Xbackup) TransferFrom(from ethgo.Address, to ethgo.Address, tokenId *big.Int) (contract.Txn, error) {
	return x.c.Txn("transferFrom", from, to, tokenId)
}

// UpdateDataMeta sends a updateDataMeta transaction in the solidity contract
func (x *Xbackup) UpdateDataMeta(infoHash string, newlinks string, newTokenUri string) (contract.Txn, error) {
	return x.c.Txn("updateDataMeta", infoHash, newlinks, newTokenUri)
}

// events

func (x *Xbackup) ApprovalEventSig() ethgo.Hash {
	return x.c.GetABI().Events["Approval"].ID()
}

func (x *Xbackup) ApprovalForAllEventSig() ethgo.Hash {
	return x.c.GetABI().Events["ApprovalForAll"].ID()
}

func (x *Xbackup) TransferEventSig() ethgo.Hash {
	return x.c.GetABI().Events["Transfer"].ID()
}