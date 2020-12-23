package lib

import (
	"fmt"
	"os"
	"strings"

	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	
)

var ESPool *es.Client

func init() {
	urls := Conf["es"]["urls"].(string)
	esCfg := es.Config{
		Addresses: strings.Split(urls, ","),
		ConnectionPoolFunc: estransport.NewConnectionPool(nil)
	}
	client, err := es.NewClient(esCfg)
	if err != nil {
		fmt.Println("init elastic-search client fail", err)
		os.Exit(1)
	}
	ESClient = client
}
