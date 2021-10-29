package notifier

func NewSmsProvider() IProvider {
	return &SmsProvider{}
}

type SmsProvider struct {
}

func (sp *SmsProvider) StockChanged(oldStock int, newStock int, product string) {

}
