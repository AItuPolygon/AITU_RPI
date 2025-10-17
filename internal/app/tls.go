package app

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	CERTIFICATE = "CERTIFICATE"
)

func (app *App) InitTls() (*tls.Config, error) {

	root, err := newCertPool(app.Config.Cert.RootCAPath)
	if err != nil {
		return nil, err
	}

	cert, err := newCertificate(app.Config.Cert.ServerCert.PublicKey, app.Config.Cert.ServerCert.PrivateKey)
	if err != nil {
		return nil, err
	}

	tls := &tls.Config{
		RootCAs:                  root,
		ClientCAs:                root,
		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		Certificates:             []tls.Certificate{cert},
		//TODO: DEBUG ONLY
		ClientAuth:            tls.RequireAndVerifyClientCert,
		VerifyPeerCertificate: app.VerifyPeerCertificate,
	}

	tls.BuildNameToCertificate()
	return tls, nil
}

func (app *App) VerifyPeerCertificate(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	for _, chain := range verifiedChains {
		for _, cert := range chain {
			if !cert.IsCA {
				for _, client := range app.Tasks {
					if len(cert.Subject.CommonName) >= len(client.Name) && strings.HasSuffix(cert.Subject.CommonName, client.Name) {
						fmt.Printf("%+v\n", client)
						return nil
					}
				}
			}
		}
	}
	return errors.New("client sertificate verification failed")
}

func newCertificate(crt, key string) (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return tls.Certificate{}, err
	}

	return cert, nil
}

func newCertPool(caFilename string) (*x509.CertPool, error) {
	root := x509.NewCertPool()

	rootPEM, err := os.ReadFile(caFilename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(rootPEM)
	if block == nil || block.Type != CERTIFICATE || len(block.Headers) != 0 {
		return nil, errors.New("error parsing root certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	root.AddCert(cert)

	return root, nil
}
