package main

import (
        "crypto/rand"
        "crypto/rsa"
        "crypto/tls"
        "crypto/x509"
        "crypto/x509/pkix"
        "encoding/pem"
        "fmt"
        "math/big"
        "os"
        "time"
)

func NewTLSCert() (cert tls.Certificate, err error) {
        now := time.Now()

        t := &x509.Certificate{
                SerialNumber: big.NewInt(now.Unix()),
                NotBefore:    now,
                NotAfter:     now.AddDate(1, 0, 0), // one year\
                Subject: pkix.Name{
                        CommonName: "inmemory.com",
                },
                KeyUsage: x509.KeyUsageDigitalSignature |
                        x509.KeyUsageKeyEncipherment |
                        x509.KeyUsageCertSign,
                BasicConstraintsValid: true,
        }

        priv, err := rsa.GenerateKey(rand.Reader, 4096)
        if err != nil {
                return tls.Certificate{}, err
        }

        newCert, err := x509.CreateCertificate(rand.Reader, t, t, priv.Public(), priv)
        if err != nil {
                return tls.Certificate{}, err
        }

        cert.Certificate = append(cert.Certificate, newCert)
        cert.PrivateKey = priv
        return
}

func getPrivateKey(cert *tls.Certificate) {
        privBytes, _ := x509.MarshalPKCS8PrivateKey(cert.PrivateKey)
        _ = pem.Encode(os.Stdout, &pem.Block{Type:  "PRIVATE KEY", Bytes: privBytes})
}

func getPEMCert(cert *tls.Certificate) {
        _ = pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Certificate[0]})
}

func main() {
        cert, err := NewTLSCert()
        if err != nil {
                fmt.Printf("%s", err.Error())
                return
        }

        //getPrivateKey(&cert)
        getPEMCert(&cert)
}
