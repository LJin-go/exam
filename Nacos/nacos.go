package Nacos

//import (
//	"github.com/nacos-group/nacos-sdk-go/v2/clients"
//	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
//	"github.com/nacos-group/nacos-sdk-go/v2/vo"
//	"log"
//)
//
//func Nacos() {
//	//create clientConfig
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "/tmp/nacos/log",
//		CacheDir:            "/tmp/nacos/cache",
//		LogLevel:            "debug",
//	}
//
//	// At least one ServerConfig
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr:      "",
//			ContextPath: "/nacos",
//			Port:        0,
//			Scheme:      "http",
//		},
//	}
//
//	// Create config client for dynamic configuration
//	client, _ := clients.CreateConfigClient(map[string]interface{}{
//		"serverConfigs": serverConfigs,
//		"clientConfig":  clientConfig,
//	})
//
//	content, _ := client.GetConfig(vo.ConfigParam{
//		DataId: "",
//		Group:  ""})
//	log.Println(content)
//}
