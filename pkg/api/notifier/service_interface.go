package notifier

//go:generate mockgen -source=$GOFILE -destination=./service_mock.go -package=$GOPACKAGE
type INotifier interface {
	Send(in *NoticationInfo)
}

type IProvider interface {
	StockChanged(oldStock int, newStock int, product string)
}
