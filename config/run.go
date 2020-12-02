package config

import "github.com/gin-gonic/gin"

func Run(r *gin.Engine, host, port string) {
	InitMasterContainer(r)
	r.Run(host + ":" + port)
}
