package app

import (
	"github.com/gin-gonic/gin"
	"walletserver/api"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *api.Router
}

func (s *Server) StartWallet() {
	s.apiRouter.SetWalletRouter(s.engine)
	//if gin.Mode() == gin.ReleaseMode {
	//
	//} else {
	//	s.engine.Run(":8389")
	//}
	err := s.engine.Run(":8089")
	if err != nil {
		panic(err)
		return
	}
}
func NewServer(engine *gin.Engine, apiRouter *api.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}
