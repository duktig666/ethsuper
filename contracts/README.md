```shell
abigen --abi=contracts/ssvnetwork/SSVNetwork.json --pkg=ssvnetwork --out=contracts/ssvnetwork/ssvnetwork.go

abigen --abi=contracts/erc721/ERC721.json --pkg=erc721 --out=contracts/erc721/erc721.go
abigen --abi=contracts/erc1155/ERC1155.json --pkg=erc721 --out=contracts/erc1155/erc1155.go
```