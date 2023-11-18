package main

import (
    "fmt"
    // "io/ioutil"
    "log"
    "os"

    "github.com/ethereum/go-ethereum/accounts/keystore"
)

// func createKs() {
//     ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
//     password := "secret"
//     account, err := ks.NewAccount(password)
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println(account.Address.Hex()) // 0x9580c87cfCE8497E215511A24513880Ca47ec694
// }

func importKs() {
    file := "./wallets/UTC--2023-11-18T07-12-31.958364000Z--21e389561ca61cca960055312af1fced7d320604"
    ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
    jsonBytes, err := os.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }

    password := "YOUR PASSWORD HERE"
    newPassword := password // Or you can use a new one
    account, err := ks.Import(jsonBytes, password, newPassword)
    if err != keystore.ErrAccountAlreadyExists {
        log.Fatal(err)
    }
    fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

    // if err := os.Remove(file); err != nil {
    //     log.Fatal(err)
    // }
}

func main() {
    // createKs()
    importKs()
}