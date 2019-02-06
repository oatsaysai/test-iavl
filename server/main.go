package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tendermint/iavl"
	dbm "github.com/tendermint/tendermint/libs/db"
)

var iavlTree *iavl.MutableTree

func main() {

	// IAVL
	var dbType = getEnv("DB_TYPE", "goleveldb")
	var dbDir = getEnv("DB_DIR_PATH", "./DB")
	name := "db"
	db := dbm.NewDB(name, dbm.DBBackendType(dbType), dbDir)
	tree := iavl.NewMutableTree(db, 0)
	tree.Load()
	iavlTree = tree

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.POST("/setKV/:key/:value", setKeyValue)

	// Prometheus
	go runProm()
	// Server
	var serverPort = getEnv("SERVER_PORT", "8080")
	e.Logger.Fatal(e.Start(":" + serverPort))
}

func runProm() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func setKeyValue(c echo.Context) error {
	go recordQueryMetrics()
	key := c.Param("key")
	value := c.Param("value")

	// Set KV
	startTime := time.Now()
	iavlTree.Set([]byte(key), []byte(value))
	go recordSetKVDurationMetrics(startTime)

	// Save version
	startTime = time.Now()
	iavlTree.SaveVersion()
	go recordSaveVersionDurationMetrics(startTime)

	return c.JSON(http.StatusCreated, nil)
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
