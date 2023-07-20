# How to capture Ethereum's NFT data

## ERC165 supportsInterface

ERC165 

```solidity
interface ERC165 {
    /// @notice Query if a contract implements an interface
    /// @param interfaceID The interface identifier, as specified in ERC-165
    /// @dev Interface identification is specified in ERC-165. This function
    ///  uses less than 30,000 gas.
    /// @return `true` if the contract implements `interfaceID` and
    ///  `interfaceID` is not 0xffffffff, `false` otherwise
    function supportsInterface(bytes4 interfaceID) external view returns (bool);
}
```

`supportsInterface` 是一个用于判断合约是否支持特定接口的函数。

参数 `interfaceID` 是一个 4 字节的标识符，用来表示特定接口。通常，接口的 ID 是通过对接口中所有函数的函数签名进行哈希运算，并取前四个字节得到的。



[ERC721](https://github.com/chiru-labs/ERC721A-Upgradeable/blob/6ed67b38c7d4ebd33531b12e72cb455a4ac747f8/contracts/ERC721AUpgradeable.sol#L225)

```solidity
/**
 * @dev Returns true if this contract implements the interface defined by
 * `interfaceId`. See the corresponding
 * [EIP section](https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified)
 * to learn more about how these ids are created.
 *
 * This function call must use less than 30000 gas.
 */
function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
    // The interface IDs are constants representing the first 4 bytes
    // of the XOR of all function selectors in the interface.
    // See: [ERC165](https://eips.ethereum.org/EIPS/eip-165)
    // (e.g. `bytes4(i.functionA.selector ^ i.functionB.selector ^ ...)`)
    return
        interfaceId == 0x01ffc9a7 || // ERC165 interface ID for ERC165.
        interfaceId == 0x80ac58cd || // ERC165 interface ID for ERC721.
        interfaceId == 0x5b5e139f; // ERC165 interface ID for ERC721Metadata.
}
```



[ERC1155](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/4d9d9073b84f56fe3eea360e5067c6ffd864c43d/contracts/token/ERC1155/ERC1155Upgradeable.sol#L49)

```solidity
/**
 * @dev See {IERC165-supportsInterface}.
 */
function supportsInterface(bytes4 interfaceId) public view virtual override(ERC165Upgradeable, IERC165Upgradeable) returns (bool) {
    return
        interfaceId == type(IERC1155Upgradeable).interfaceId ||
        interfaceId == type(IERC1155MetadataURIUpgradeable).interfaceId ||
        super.supportsInterface(interfaceId);
}
```

`type(IERC1155Upgradeable).interfaceId` ？

> `type(IERC1155Upgradeable).interfaceId` 是 Solidity 中的特殊语法，用于获取给定接口的接口标识符（interfaceId）。这个语法通常用于实现 ERC-165 的合约，用于进行合约接口的检测。
>
> 下面是对 `type(IERC1155Upgradeable).interfaceId` 表达式的解释：
>
> 1. `IERC1155Upgradeable`：这是我们要获取接口标识符的接口名称。它是在 ERC-1155 标准中定义的一个接口，包含了合约必须实现的函数，以使其符合 ERC-1155 标准。
> 2. `type(IERC1155Upgradeable)`：`type(...)` 表达式用于获取 `IERC1155Upgradeable` 接口的类型。这是 Solidity 中的一个特殊运算符，允许我们在不创建合约实例的情况下访问接口类型。
> 3. `.interfaceId`：这是接口的一个属性，代表它的唯一标识符。它是通过对接口中所有函数的函数签名进行哈希运算，并取前四个字节得到的。这个 interfaceId 用于在 ERC-165 的支持下进行接口检测，以确定合约是否实现了特定的接口。

```solidity
// forge test -vvvv --match-test testErc1155InterfaceId
function testErc1155InterfaceId() public {
    bytes4 interfaceId = type(IERC1155Upgradeable).interfaceId;
    console2.logBytes4(interfaceId);
}
```

输出结果：`0xd9b67a26`



## 创建ERC721或者ERC1155合约的交易

**其中 ERC721和ERC1155都实现了ERC165**。

1. 如果交易的 `tx.To() == nil`，说明这笔交易是创建合约。

2. `ContractAddress != address(0)`

3. 可以通过 `supportsInterface` 来 **判断一笔交易是否是 创建ERC721或者ERC1155合约**。

```go
const (
	Erc721 = iota
	Erc1155
	Unknown
)

func JudgmentErcType(contract common.Address, backend bind.ContractBackend) (int, error) {
	instance, err := erc721.NewErc721(contract, backend)
	if err != nil {
		return 0, err
	}

	if ok, err := instance.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xac, 0x58, 0xcd}); err == nil && ok {
		return Erc721, nil
	}

	if ok, err := instance.SupportsInterface(&bind.CallOpts{}, [4]byte{0xd9, 0xb6, 0x7a, 0x26}); err == nil && ok {
		return Erc1155, nil
	}

	return Unknown, nil
}
```





## NFT Transfer

### 判断是否属于 ERC721 Transfer

```go
func checkErc721Transfer(log *types.Log) bool {
	if len(log.Topics) == 4 && log.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
		return true
	} else {
		return false
	}
}
```

`log.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"` ?

参看这段solidity代码：

[https://github.com/chiru-labs/ERC721A-Upgradeable/blob/6ed67b38c7d4ebd33531b12e72cb455a4ac747f8/contracts/ERC721AUpgradeable.sol#L556-L566](https://github.com/chiru-labs/ERC721A-Upgradeable/blob/6ed67b38c7d4ebd33531b12e72cb455a4ac747f8/contracts/ERC721AUpgradeable.sol#L556-L566)

```solidity
// The `Transfer` event signature is given by:
// `keccak256(bytes("Transfer(address,address,uint256)"))`.
bytes32 private constant _TRANSFER_EVENT_SIGNATURE =
        0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef;

assembly {
            // Emit the `Transfer` event.
            log4(
                0, // Start of data (0, since no data).
                0, // End of data (0, since no data).
                _TRANSFER_EVENT_SIGNATURE, // Signature.
                from, // `from`.
                toMasked, // `to`.
                tokenId // `tokenId`.
            )
        }
```

也可以对比如下的log：[https://etherscan.io/tx/0x061355142a9c60b8fc34ec49b8721830bef2ab55f97b73654bf4e62643a015f4#eventlog](https://etherscan.io/tx/0x061355142a9c60b8fc34ec49b8721830bef2ab55f97b73654bf4e62643a015f4#eventlog)



Transfer event:

https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/cce46e03434b72e060a726057365f249f3e13a61/contracts/token/ERC721/IERC721Upgradeable.sol#L15

```solidity
/**
  * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
  */
event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
```



### 判断ERC721 Transfer Type

```go
func checkType(from string, to string) string {
	if from == "0x0000000000000000000000000000000000000000" {
		return "mint"
	} else if to == "0x0000000000000000000000000000000000000000" {
		return "burn"
	} else {
		return "transfer"
	}
}
```



## NFT Matadata

```solidity
/**
  * @dev Returns the Uniform Resource Identifier (URI) for `tokenId` token.
  */
function tokenURI(uint256 tokenId) external view returns (string memory);
```

`tokenURI()`方法返回指定 NFT 的 Metadata URI。Metadata URI 是一个指向包含有关 NFT 信息的 JSON 文件的 URL。

元数据扩展对于ERC-721智能合约是可选的（请参阅下面的“注意事项”）。这允许您的智能合约被询问其名称以及有关您的 NFT 所代表的资产的详细信息:

```solidity
/// @title ERC-721 Non-Fungible Token Standard, optional metadata extension
/// @dev See https://eips.ethereum.org/EIPS/eip-721
///  Note: the ERC-165 identifier for this interface is 0x5b5e139f.
interface ERC721Metadata /* is ERC721 */ {
    /// @notice A descriptive name for a collection of NFTs in this contract
    function name() external view returns (string _name);

    /// @notice An abbreviated name for NFTs in this contract
    function symbol() external view returns (string _symbol);

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    /// @dev Throws if `_tokenId` is not a valid NFT. URIs are defined in RFC
    ///  3986. The URI may point to a JSON file that conforms to the "ERC721
    ///  Metadata JSON Schema".
    function tokenURI(uint256 _tokenId) external view returns (string);
}
```

一般Metadata JSON的值：

```json
{
    "title": "Asset Metadata",
    "type": "object",
    "properties": {
        "name": {
            "type": "string",
            "description": "Identifies the asset to which this NFT represents"
        },
        "description": {
            "type": "string",
            "description": "Describes the asset to which this NFT represents"
        },
        "image": {
            "type": "string",
            "description": "A URI pointing to a resource with mime type image/* representing the asset to which this NFT represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive."
        }
    }
}
```

eg:

- Contract: https://etherscan.io/address/0x49c36afa15c7fdbd57ce3d61d80f39b6615a76ef#readContract

-  [https://invisiblefriends.io/api/metadata/3d/1](https://invisiblefriends.io/api/metadata/3d/1)





