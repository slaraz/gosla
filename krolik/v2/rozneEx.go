package v2

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

var rozneEx = map[string]func(*Ex, *amqp.Channel) error{
	"stdex":  stdEx,
	"szybki": szybkiEx,
	//"pewny": przygotujPewny,
}

// --- stdex ---

func stdEx(ex *Ex, chann *amqp.Channel) error {
	if err := chann.ExchangeDeclare(
		ex.nazwa, // nazwa exchangera
		ex.kind,  // sposób routingu: fanout, topic, headers
		true,     // durable - czy ma przeżyć restart serwera
		false,    // autodelete - czy skasować jeśli brak podłączonych kolejek
		false,    // internal - false oznacza moża normalnie publikować z zewnątrz
		false,    // noWait - serwer nie zwraca nic, ewentualne błędy są wysyłane asynchronicznie
		nil,      // arguments
	); err != nil {
		return fmt.Errorf("ExchangeDeclare(): %v", err)
	}
	ex.publikuj = ex.publikujStdEx
	return nil
}

func (ex *Ex) publikujStdEx(bajty []byte) error {
	if err := ex.sesja.chann.Publish(
		ex.nazwa,
		"",    // RoutingKey - dla exchangera typu topic
		false, // mandatory - czy upewnić się, że wiadomość gdzieś trafi (w wypadku braku kolejek lub zły routing - exception)
		false, // immediate - deprecated
		amqp.Publishing{
			// tu można poszaleć! patrz inne parametry amqp.Publishing
			Body:         bajty,
			DeliveryMode: amqp.Persistent,
		},
	); err != nil {
		return fmt.Errorf("ex.sesja.chann.Publish(): %v", err)
	}
	return nil
}

// --- szybki ---

func szybkiEx(ex *Ex, chann *amqp.Channel) error {
	if err := chann.ExchangeDeclare(
		ex.nazwa, // nazwa exchangera
		ex.kind,  // sposób routingu: fanout, topic, headers
		true,     // durable - czy ma przeżyć restart serwera
		false,    // autodelete - czy skasować jeśli brak podłączonych kolejek
		false,    // internal - false oznacza moża normalnie publikować z zewnątrz
		true,     // noWait - serwer nie zwraca nic, ewentualne błędy są wysyłane asynchronicznie
		nil,      // arguments
	); err != nil {
		return fmt.Errorf("ExchangeDeclare(): %v", err)
	}
	ex.publikuj = ex.publikujSzybkiEx
	return nil
}

func (ex *Ex) publikujSzybkiEx(bajty []byte) error {
	if err := ex.sesja.chann.Publish(
		ex.nazwa,
		"",    // RoutingKey - dla exchangera typu topic
		false, // mandatory - czy upewnić się, że wiadomość gdzieś trafi (w wypadku braku kolejek lub zły routing - exception)
		false, // immediate - deprecated
		amqp.Publishing{
			// tu można poszaleć! patrz inne parametry amqp.Publishing
			Body:         bajty,
			DeliveryMode: amqp.Transient,
			Timestamp:    time.Now(),
		},
	); err != nil {
		return fmt.Errorf("ex.sesja.chann.Publish(): %v", err)
	}
	return nil
}
