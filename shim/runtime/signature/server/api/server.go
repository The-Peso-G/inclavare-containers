package api

import (
	"crypto/rsa"
	"crypto/x509"

	"github.com/alibaba/inclavare-containers/shim/runtime/signature/server/conf"
	"github.com/alibaba/inclavare-containers/shim/runtime/signature/server/util"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	router      *gin.Engine
	listenAddr  string
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
	certificate *x509.Certificate
}

func NewApiServer(listenAddr string, conf *conf.Config) (*ApiServer, error) {
	privateKey, err := util.ParseRsaPrivateKey(conf.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	publicKey, err := util.ParseRsaPublicKey(conf.PublicKeyPath)
	if err != nil {
		return nil, err
	}
	s := &ApiServer{
		router:     gin.Default(),
		listenAddr: listenAddr,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	s.installRoutes()
	return s, nil
}

func (s *ApiServer) RunForeground() error {
	return s.router.Run(s.listenAddr)
}
