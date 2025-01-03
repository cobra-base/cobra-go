abigen --abi .\erc20.json --pkg ethers --type ERC20 --out erc20.go
abigen --abi .\abi\PancakeswapV2Router.json --pkg pancakeswap_binding --type PancakeswapV2Router --out binding/pancakeswap/V2Router.go
abigen --abi .\abi\PancakeswapV3Router.json --pkg pancakeswap_binding --type PancakeswapV3Router --out binding/pancakeswap/V3Router.go
abigen --abi .\abi\PancakeswapV3Quoter.json --pkg pancakeswap_binding --type PancakeswapV3Quoter --out binding/pancakeswap/V3Quoter.go
abigen --abi .\abi\PancakeswapV3Pair.json --pkg pancakeswap_binding --type PancakeswapV3Pair --out binding/pancakeswap/V3Pair.go
abigen --abi .\abi\UniswapV2Router.json --pkg uniswap_binding --type UniswapV2Router --out binding/uniswap/V2Router.go
abigen --abi .\abi\UniswapV3Router.json --pkg uniswap_binding --type UniswapV3Router --out binding/uniswap/V3Router.go
abigen --abi .\abi\UniswapV3Quoter.json --pkg uniswap_binding --type UniswapV3Quoter --out binding/uniswap/V3Quoter.go
abigen --abi .\abi\UniswapV3Pair.json --pkg uniswap_binding --type UniswapV3Pair --out binding/uniswap/V3Pair.go