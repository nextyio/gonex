// Copyright 2015 The Nexty Authors
// This file is part of the gonex library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ErrorLogInvalidOpCode            = "invalid opcode 0x%x"
	ErrorLogStackUnderflow           = "stack underflow (%d <=> %d)"
	ErrorLogStackLimitReached        = "stack limit reached %d (%d)"
	ErrorLogWriteProtection          = "write protection"
	ErrorLogOutOfGas                 = "out of gas"
	ErrorLogGasUintOverflow          = "gas overflow unsigned 64 bit integer"
	ErrorLogReturnDataOutOfBounds    = "return data out of bounds"
	ErrorLogMaxCodeSizeExceeded      = "max code size exceeded"
	ErrorLogInvalidJump              = "invalid jump destination"
	ErrorLogCodeStoreOutOfGas        = "contract creation code storage out of gas"
	ErrorLogDepth                    = "max call depth exceeded"
	ErrorLogInsufficientBalance      = "insufficient balance for transfer"
	ErrorLogContractAddressCollision = "contract address collision"
)

var (
	// TopicRevert is Keccak("REVERT")
	TopicRevert = common.HexToHash("e13872d662304a4be4efe6d4425b00781f90609ddf2ef6e5b5e5c8bc7f5ed47f")

	// TopicError is Keccak("ERROR")
	TopicError = common.HexToHash("6368faa35d5ea15ae80b929d8626383bb91c2157389a6ddb6239282e6aa9005d")
)

var (
	// SolidityErrorSignature is Keccak("Error(string)")
	SolidityErrorSignature = []byte{0x08, 0xc3, 0x79, 0xa0}
)

// GetSolidityRevertMessage handles Solidity revert and require message.
func GetSolidityRevertMessage(res []byte) string {
	if len(res) < 4+32+32 {
		return string(res)
	}
	if bytes.Compare(res[:4], SolidityErrorSignature) != 0 {
		return string(res)
	}
	res = res[4:]
	offset := int(new(big.Int).SetBytes(res[:32]).Uint64())
	res = res[32:]
	size := int(new(big.Int).SetBytes(res[:32]).Uint64())
	if len(res) < offset+size {
		return string(res)
	}
	msg := string(res[offset : offset+size])
	return msg
}