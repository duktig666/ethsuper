//  Copyright © 2022-2023. duktig666 Limited.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// description: abi编码 @see https://github.com/0xsequence/ethkit/tree/master/ethcoder
// @author renshiwei
// Date: 2022/11/3 17:08

package cryptor

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

func buildArgumentsFromTypes(argTypes []string) (abi.Arguments, error) {
	args := abi.Arguments{}
	for _, argType := range argTypes {
		abiType, err := abi.NewType(argType, "", nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{Type: abiType})
	}
	return args, nil
}

func AbiCoder(argTypes []string, argValues []interface{}) ([]byte, error) {
	if len(argTypes) != len(argValues) {
		return nil, errors.New("invalid arguments - types and values do not match")
	}
	args, err := buildArgumentsFromTypes(argTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to build abi: %v", err)
	}
	return args.Pack(argValues...)
}

func AbiCoderHex(argTypes []string, argValues []interface{}) (string, error) {
	b, err := AbiCoder(argTypes, argValues)
	if err != nil {
		return "", err
	}
	h := hexutil.Encode(b)
	return h, nil
}

func AbiDecoder(argTypes []string, input []byte, argValues []interface{}) error {
	if len(argTypes) != len(argValues) {
		return errors.New("invalid arguments - types and values do not match")
	}
	args, err := buildArgumentsFromTypes(argTypes)
	if err != nil {
		return fmt.Errorf("failed to build abi: %v", err)
	}
	values, err := args.Unpack(input)
	if err != nil {
		return err
	}
	if len(args) > 1 {
		return args.Copy(&argValues, values)
	} else {
		return args.Copy(&argValues[0], values)
	}
}

func AbiDecoderWithReturnedValues(argTypes []string, input []byte) ([]interface{}, error) {
	args, err := buildArgumentsFromTypes(argTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to build abi: %v", err)
	}
	return args.UnpackValues(input)
}
