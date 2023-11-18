package main

import (
    "fmt"
    // "io/ioutil"
    "log"
    // "os"

    "github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
    ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
    password := "YOUR PASSWORD HERE"
    account, err := ks.NewAccount(password)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(account.Address.Hex()) // 0x9580c87cfCE8497E215511A24513880Ca47ec694
}

// func importKs() {
//     file := "./wallets/UTC--2023-11-18T06-05-57.126947000Z--9580c87cfce8497e215511a24513880ca47ec694"
//     ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
//     jsonBytes, err := os.ReadFile(file)
//     if err != nil {
//         log.Fatal(err)
//     }

//     password := "secret"
//     account, err := ks.Import(jsonBytes, password, password)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

//     if err := os.Remove(file); err != nil {
//         log.Fatal(err)
//     }
// }

func main() {
    createKs()
    // importKs()
}