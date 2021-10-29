package notifier

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type service struct {
	conf     *config
	provider IProvider
}

type NoticationInfo struct {
	SellerUUID  string
	SellerPhone string
	ProductName string
}

func NewNotifier() (INotifier, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, err
	}

	s := &service{
		conf: conf,
	}
	switch conf.Type {
	case "sms":
		s.provider = NewSmsProvider()
	case "email":
		s.provider = NewEmailProvider()
	}

	return s, nil
}

func (s *service) Send(in *NoticationInfo) {
	log.Info().Msg(fmt.Sprintf("SMS Warning sent to %s (Phone: %s): %s Product stock changed",
		in.SellerUUID, in.SellerPhone, in.ProductName))

	go s.provider.StockChanged(0, 0, "")
}
