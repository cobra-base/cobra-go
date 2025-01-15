# cobra-go

# ABI
  - abigen --abi .\erc20.json --pkg ethers --type ERC20 --out erc20.go
  - abigen --abi .\abi\PancakeswapV2Router.json --pkg binding --type PancakeswapV2Router --out binding/PancakeswapV2Router.go
  - abigen --abi .\abi\PancakeswapV3Router.json --pkg binding --type PancakeswapV3Router --out binding/PancakeswapV3Router.go
  - abigen --abi .\abi\PancakeswapV3Quoter.json --pkg binding --type PancakeswapV3Quoter --out binding/PancakeswapV3Quoter.go
  - abigen --abi .\abi\UniswapV2Router.json --pkg binding --type UniswapV2Router --out binding/UniswapV2Router.go
  - abigen --abi .\abi\UniswapV3Router.json --pkg binding --type UniswapV3Router --out binding/UniswapV3Router.go
  - abigen --abi .\abi\UniswapV3Quoter.json --pkg binding --type UniswapV3Quoter --out binding/UniswapV3Quoter.go

# publish
  - git tag v0.0.46
  - git push --tags

# notice
  - public,切勿上传敏感信息