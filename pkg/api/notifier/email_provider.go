package notifier

func NewEmailProvider() IProvider {
	return &EmailProvider{}
}

type EmailProvider struct {
}

func (ep *EmailProvider) StockChanged(oldStock int, newStock int, product string) {

}
