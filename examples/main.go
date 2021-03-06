package main

import (
	"flag"
	"log"

	"github.com/codingbeard/cbsession"
	"github.com/codingbeard/cbsession/memcache"
	"github.com/codingbeard/cbsession/memory"
	"github.com/codingbeard/cbsession/mysql"
	"github.com/codingbeard/cbsession/postgres"
	"github.com/codingbeard/cbsession/redis"
	"github.com/codingbeard/cbsession/sqlite3"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const defaultProvider = "memory"

var serverSession = cbsession.New(cbsession.NewDefaultConfig())

func init() {
	providerName := flag.String("provider", defaultProvider, "Name of provider")
	flag.Parse()

	var config cbsession.ProviderConfig
	switch *providerName {
	case "memory":
		config = &memory.Config{}
	case "memcache":
		config = &memcache.Config{
			ServerList: []string{
				"0.0.0.0:11211",
			},
			MaxIdleConns: 8,
			KeyPrefix:    "session",
		}
	case "mysql":
		config = mysql.NewConfigWith("127.0.0.1", 3306, "root", "session", "test", "session")
	case "postgres":
		config = postgres.NewConfigWith("127.0.0.1", 5432, "root", "session", "test", "session")
	case "redis":
		config = &redis.Config{
			Host:        "127.0.0.1",
			Port:        6379,
			PoolSize:    8,
			IdleTimeout: 300,
			KeyPrefix:   "session",
		}
	case "sqlite3":
		config = sqlite3.NewConfigWith("test.db", "session")
	default:
		panic("Invalid provider")
	}

	err := serverSession.SetProvider(*providerName, config)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Starting example with provider: " + *providerName)
}

func main() {
	r := router.New()
	r.GET("/", indexHandler)
	r.GET("/set", setHandler)
	r.GET("/get", getHandler)
	r.GET("/delete", deleteHandler)
	r.GET("/getAll", getAllHandler)
	r.GET("/flush", flushHandler)
	r.GET("/destroy", destroyHandler)
	r.GET("/sessionid", sessionIDHandler)
	r.GET("/regenerate", regenerateHandler)
	r.GET("/setexpiration", setExpirationHandler)
	r.GET("/getexpiration", getExpirationHandler)

	addr := "0.0.0.0:8086"
	log.Println("Session example server listen: http://" + addr)

	err := fasthttp.ListenAndServe(addr, r.Handler)
	if err != nil {
		log.Fatal(err)
	}
}
