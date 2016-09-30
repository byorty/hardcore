package plugin

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/meta/conf"
	"crypto/tls"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"github.com/byorty/hardcore/scope"
)

type TlsImpl struct {}

func NewTls() types.ApplicationPlugin {
	return new(TlsImpl)
}

func (t *TlsImpl) Run() {
	cipherSuites := []uint16{
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	}
	conf := &tls.Config{
		ClientAuth:             tls.RequireAndVerifyClientCert,
		CipherSuites:           cipherSuites,
		MinVersion:             tls.VersionTLS12,
		SessionTicketsDisabled: true,
	}
	pemBytes, err := ioutil.ReadFile(scope.App().GetCertFilename())
	if err == nil {
		pemBlock, _ := pem.Decode(pemBytes)
		cert, _ := x509.ParseCertificate(pemBlock.Bytes)
		pool := x509.NewCertPool()
		pool.AddCert(cert)
		conf.RootCAs = pool
	} else {
		scope.App().GetLogger().Error("tls can't read certificate %s, error - %v", scope.App().GetCertFilename(), err)
		scope.App().Exit()
	}
	cert, err := tls.LoadX509KeyPair(scope.App().GetCertFilename(), scope.App().GetPrivateKeyFilename())
	if err == nil {
		conf.Certificates = []tls.Certificate{cert}
	} else {
		scope.App().GetLogger().Error("tls can't load certificate %s or private key %s, error - %v", scope.App().GetCertFilename(), scope.App().GetPrivateKeyFilename(), err)
		scope.App().Exit()
	}
}