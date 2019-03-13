package kubetoken

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/invia-de/kubetoken/internal/cert"

	ldap "gopkg.in/ldap.v2"
//	"log"
)

type LDAPCreds struct {
	Host     string
	Port     int
	BindDN   string
	Password string
	SkipVerify bool
}

func (l *LDAPCreds) Bind() (*ldap.Conn, error) {
	addr := fmt.Sprintf("%s:%d", l.Host, l.Port)
	config := tls.Config{
		ServerName: l.Host,
		InsecureSkipVerify: l.SkipVerify,
	}
	// TODO(dfc) should construct a net.Conn explicitly
	// to set the Dial and Read/Write Deadlines
	conn, err := ldap.DialTLS("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	err = conn.Bind(l.BindDN, l.Password)
//	log.Println("fehler beim binden:")
//	log.Println(l.BindDN)
//	log.Println(err)
//	log.Println(conn)

	return conn, err
}

type Signer struct {
	Cert    *x509.Certificate
	PrivKey *rsa.PrivateKey
}

func (s *Signer) Sign(csr *x509.CertificateRequest) ([]byte, error) {
	return cert.SignCSR(csr, s.Cert, s.PrivKey)
}
