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

    fmt.Println(account.Address.Hex()) 
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
//     fmt.Println(account.Address.Hex()) 

//     if err := os.Remove(file); err != nil {
//         log.Fatal(err)
//     }
// }

func main() {
    createKs()
    // importKs()
}