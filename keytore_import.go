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

//     fmt.Println(account.Address.Hex()) 
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
    if err != nil && err != keystore.ErrAccountAlreadyExists {
        log.Fatal(err)
    }
    fmt.Println(account.Address.Hex()) 

    /*
    @dev code will not write to the same dir where the imported file sits, so no need to delete the old one. 
        Unless it's the case where you use a new password and want to remove the one encrypted with the old password.
    */
    // if err := os.Remove(file); err != nil {
    //     log.Fatal(err)
    // }
}

func main() {
    // createKs()
    importKs()
}