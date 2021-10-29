package notifier

type INotifier interface {
	Send(in *NoticationInfo)
}

type IProvider interface {
	StockChanged(oldStock int, newStock int, product string)
}
