package v2

import (
	"fmt"

	"github.com/streadway/amqp"
)

var rozneEx = map[string]func(*Ex, *amqp.Channel) error{
	"std": przygotujStd,
	//"szybki": przygotujSzybki,
	//"pewny": przygotujPewny,
}

// std

func przygotujStd(ex *Ex, chann *amqp.Channel) error {
	err := chann.ExchangeDeclare(
		ex.nazwa, // nazwa exchangera
		ex.kind,  // sposób routingu: fanout, topic, headers
		true,     // durable - czy ma przeżyć restart serwera
		false,    // autodelete - czy skasować jeśli brak podłączonych kolejek
		false,    // internal - false oznacza moża normalnie publikować z zewnątrz
		false,    // noWait - serwer nie zwraca nic, ewentualne błędy są wysyłane asynchronicznie
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("ExchangeDeclare(): %v", err)
	}
	ex.publikuj = ex.publikujStd
	return nil
}

func (ex *Ex) publikujStd(bajty []byte) error {
	if !ex.sesja.czyOK {
		return fmt.Errorf("brak połączenia")
	}
	err := ex.sesja.chann.Publish(
		ex.nazwa,
		"",    // RoutingKey - dla exchangera typu topic
		false, // mandatory - czy upewnić się, że wiadomość gdzieś trafi (w wypadku braku kolejek lub zły routing - exception)
		false, // immediate - deprecated
		amqp.Publishing{
			// tu można poszaleć! patrz inne parametry amqp.Publishing
			Body: bajty,
		})
	if err != nil {
		return fmt.Errorf("ch.Publish(): %v", err)
	}
	return nil
}

// pewny
