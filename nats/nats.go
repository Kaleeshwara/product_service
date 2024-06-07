package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"product/models"
	"product/service"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type ProductHandler struct {
	service service.APIService
}

var productHandler ProductHandler

func NatsInitalizeConnection(db *gorm.DB) {
	productHandler.service = service.NewProductService(db)
}

func ConnectAndSubscribe() {
	// Logic to connect to NATS and subscribe to topics
	fmt.Println("Connecting to NATS server and subscribing to topics...")
}

func Subscribe() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(nats.DefaultURL, err)
	}
	defer nc.Close()
	_, err = nc.Subscribe("product.created", func(msg *nats.Msg) {
		var product models.Product
		err := json.Unmarshal(msg.Data, &product)
		if err != nil {
			log.Fatal(err)
		}
		_, err = productHandler.service.CreateProduct(&product)
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
