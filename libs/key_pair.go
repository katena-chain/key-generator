package libs

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "os"
    "syscall"

    tcOs "github.com/transchain/sdk-go/os"
    "golang.org/x/crypto/ed25519"
    "golang.org/x/crypto/nacl/box"
)

type KeyPair struct {
    PubKey  string
    PrivKey string
}

func (kp *KeyPair) Show() {
    fmt.Println(fmt.Sprintf("Public key: %s", kp.PubKey))
    fmt.Println(fmt.Sprintf("Private key: %s", kp.PrivKey))
}

func (kp *KeyPair) Save(filepath string) error {

    err := tcOs.EnsureFileDir(filepath, tcOs.DefaultDirPerm)
    if err != nil {
        return err
    }

    file, err := os.OpenFile(filepath, syscall.O_CREAT|syscall.O_WRONLY, tcOs.DefaultFilePerm)
    defer func() {
        _ = file.Close()
    }()
    if err != nil {
        return err
    }

    _, err = file.WriteString(fmt.Sprintf("Public key : %s \nPrivate key : %s \n", kp.PubKey, kp.PrivKey))
    if err != nil {
        return err
    }

    return nil
}

func GenerateX25519() (KeyPair, error) {
    // Generates an X25519 keypair

    publicKey, privateKey, err := box.GenerateKey(rand.Reader)
    if err != nil {
        return KeyPair{}, err
    }

    keys := KeyPair{
        PubKey:  base64.StdEncoding.EncodeToString(publicKey[:]),
        PrivKey: base64.StdEncoding.EncodeToString(privateKey[:]),
    }

    return keys, nil
}

func GenerateEd25519() (KeyPair, error) {
    // Generates an ED25519 keypair

    pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
    if err != nil {
        return KeyPair{}, err
    }

    keys := KeyPair{
        PubKey:  base64.StdEncoding.EncodeToString(pubKey),
        PrivKey: base64.StdEncoding.EncodeToString(privKey),
    }

    return keys, nil
}
