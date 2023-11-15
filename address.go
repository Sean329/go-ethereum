package main

import (
    "fmt"

    "github.com/ethereum/go-ethereum/common"
)

func main() {
    address := common.HexToAddress("0xa47E24BC59ff95685D50DFf57f958351559E5127")

    fmt.Println(address.Hex())        // 0xa47E24BC59ff95685D50DFf57f958351559E5127
    // fmt.Println(address.Hash().Hex()) // 0x000000000000000000000000a47E24BC59ff95685D50DFf57f958351559E5127 --- left padding with 0s, depreciated
    fmt.Println(address.Bytes())      // [164 126 36 188 89 255 149 104 93 80 223 245 127 149 131 81 85 158 81 39]
}