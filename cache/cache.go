package cache

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/ArtisanCloud/go-libs/cache"
	"github.com/ArtisanCloud/go-libs/str"
)

var (
	CacheConnection *cache.GRedis

)

func SetupCache() (err error) {

	c := config.CacheConn
	//fmt.Dump(c)

	options := cache.RedisOptions{
		Host:       c.Host,
		Password:   c.Password,
		DB:         c.DB,
		SSLEnabled: c.SSLEnabled,
	}

	CacheConnection = cache.NewGRedis(&options)
	//fmt2.Printf("CacheConnection:%+v \r\n", CacheConnection.Pool.String())

	//CacheMapLockers = make(map[string]*sync.Mutex)

	return nil

}

func GetKeyPrefix() string {
	strAppName := str.Snake(config.AppConfigure.Name, "_")
	return strAppName + "_database_" + strAppName + "_cache:"
}
