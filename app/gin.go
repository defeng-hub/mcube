package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ginApps = map[string]GinApp{}
)

// HTTPService Http服务的实例
type GinApp interface {
	Registry(gin.IRouter)
	Config() error
	Name() string
}

// RegistryGinApp 服务实例注册
func RegistryGinApp(app GinApp) {
	// 已经注册的服务禁止再次注册
	_, ok := ginApps[app.Name()]
	if ok {
		panic(fmt.Sprintf("gin app %s has registed", app.Name()))
	}

	ginApps[app.Name()] = app
}

func GetGinApp(name string) GinApp {
	app, ok := ginApps[name]
	if !ok {
		panic(fmt.Sprintf("http app %s not registed", name))
	}

	return app
}

// LoadGinApp 装载所有的gin app
func LoadGinApp(pathPrefix string, root gin.IRouter) {
	for _, api := range ginApps {
		if pathPrefix != "" && !strings.HasPrefix(pathPrefix, "/") {
			pathPrefix = "/" + pathPrefix
		}
		api.Registry(root.Group(pathPrefix))
	}
}
