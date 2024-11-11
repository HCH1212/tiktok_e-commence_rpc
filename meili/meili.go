package meili

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/viper"
)

var Client meilisearch.ServiceManager

func InitMeili() {
	client := meilisearch.New(viper.GetString("meilisearch.addr"), meilisearch.WithAPIKey(viper.GetString("meilisearch.key")))
	Client = client
}
