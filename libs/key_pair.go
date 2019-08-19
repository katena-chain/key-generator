package libs

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "os"

    "github.com/riobard/go-x25519"
    "golang.org/x/crypto/ed25519"
)

type KeyPair struct {
    PubKey  string
    PrivKey string
}

func (kp *KeyPair) Show() {
    fmt.Println("Public key :", kp.PubKey)
    fmt.Println("Private key :", kp.PrivKey)
}

func (kp *KeyPair) Save(filepath string) error {

    file, err := os.Create(filepath)

    if err != nil {
        return errors.New("couldn't write to the file" + err.Error())
    }

    defer func() {
        _ = file.Close()
    }()

    _, err = file.WriteString(fmt.Sprintf("Public key : %s \nPrivate key : %s \n", kp.PubKey, kp.PrivKey))
    if err != nil {
        return err
    }
    fmt.Println("Keys saved to :", filepath)

    return nil
}

func GenerateX25519() (KeyPair, error) {
    // Generates an X25519 keypair

    secretKey, err := x25519.GenerateKey(rand.Reader)

    if err != nil {
        return KeyPair{}, err
    }

    keys := KeyPair{
        PubKey:  base64.StdEncoding.EncodeToString(secretKey.Public()),
        PrivKey: base64.StdEncoding.EncodeToString(secretKey.Bytes()),
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
