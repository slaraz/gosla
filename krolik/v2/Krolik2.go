package v2



type Exchanger struct {
	Url         string
	Nazwa       string
	Typ         TypExchangera
	Durable     bool
	AutoDelete  bool
	Internal    bool
	NoWait      bool
	AlternateEx bool
	Confirm     bool
}

type TypExchangera string

const (
	ExFanout  TypExchangera = "fanout"
	ExDirect  TypExchangera = "direct"
	ExTopic   TypExchangera = "topic"
	ExHeaders TypExchangera = "headers"
)

// // MusiKrolik tworzy połączenie z RabbitMQ i exchanger.
// func MusiKrolik2(krolikURL, krolikEX string) *Krolik {

// 	
// 	log.Printf("[RabbitMQ.Dial] -> %q", krolikURL)
// 	conn, err := amqp.Dial(krolikURL)
// 	if err != nil {
// 		log.Fatalf("[RabbitMQ] błąd Dial(): %v", err)
// 	}
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalln("[RabbitMQ] błąd Channel():", err)
// 	}

// 	log.Printf("[RabbitMQ.Decl] (EX) %q", krolikEX)
// 	if err = ch.ExchangeDeclare(
// 		krolikEX, // nazwa exchangera
// 		"direct", // typ: direct, fanout, topic, headers
// 		true,     // durable - czy ma przerzyć reset serwera
// 		false,    // autodelete - czy skasować jeśli brak podłączonych kolejek
// 		false,    // internal
// 		false,    // noWait
// 		nil,      // arguments
// 	); err != nil {
// 		log.Fatalf("[RabbitMQ] błąd ExchangeDeclare(): %v", err)
// 	}

// 	dlx := krolikEX + ".rejected"
// 	log.Printf("[RabbitMQ.Decl] (DLX) %q", dlx)
// 	if err = ch.ExchangeDeclare(
// 		dlx,      // nazwa exchangera
// 		"direct", // typ: direct, fanout, topic, headers
// 		true,     // durable - czy ma przerzyć reset serwera
// 		false,    // autodelete - czy skasować jeśli nie ma bindingów
// 		false,    // internal
// 		false,    // noWait
// 		nil,      // arguments
// 	); err != nil {
// 		log.Fatalf("[RabbitMQ] błąd ExchangeDeclare(): %v", err)
// 	}

// 	return &Krolik{
// 		ex:  krolikEX,
// 		dlx: dlx,
// 		ch:  ch,
// 	}
// }

// // MusiPobierac binduje po kluczu kolejkę w rabicie
// // i podłącza handler do obsługi wiadomości z kolejki.
// func (krolik *Krolik) MusiPobierac(klucz string, handler func([]byte) error) {

// 	// Przygotowanie kolejki.
// 	nazwaKolejki := krolik.nazwaKolejki(klucz)
// 	args := amqp.Table{"x-dead-letter-exchange": krolik.dlx}
// 	_, err := krolik.ch.QueueDeclare(
// 		nazwaKolejki, // name of the queue
// 		true,         // durable
// 		false,        // delete kiedy nieużyta przez klienta
// 		false,        // exclusive
// 		false,        // noWait
// 		args,         // arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("[RabbitMQ] błąd QueueDeclare(): %v", err)
// 	}

// 	// Przyogotowanie kolejki dla odrzuconych.
// 	nazwaOdrzucone := krolik.nazwaOdrzucone(klucz)
// 	_, err = krolik.ch.QueueDeclare(
// 		nazwaOdrzucone, // name of the queue
// 		true,           // durable
// 		false,          // delete kiedy nieużyta przez klienta
// 		false,          // exclusive
// 		false,          // noWait
// 		nil,            // arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("[RabbitMQ] błąd QueueDeclare(): %v", err)
// 	}

// 	log.Printf("[RabbitMQ.Bind] (%s)----/%s/---->[%s]", krolik.dlx, klucz, nazwaOdrzucone)
// 	if err = krolik.ch.QueueBind(
// 		nazwaOdrzucone, // name of the queue
// 		klucz,          // bindingKey
// 		krolik.dlx,     // sourceExchange
// 		false,          // noWait
// 		nil,            // arguments
// 	); err != nil {
// 		log.Fatalf("[RabbitMQ] błąd QueueBind(): %v", err)
// 	}

// 	log.Printf("[RabbitMQ.Bind] (%s)-----/%s/---->[%s]", krolik.ex, klucz, nazwaKolejki)
// 	if err = krolik.ch.QueueBind(
// 		nazwaKolejki, // name of the queue
// 		klucz,        // bindingKey
// 		krolik.ex,    // sourceExchange
// 		false,        // noWait
// 		nil,          // arguments
// 	); err != nil {
// 		log.Fatalf("[RabbitMQ] błąd QueueBind(): %v", err)
// 	}

// 	deliveries, err := krolik.ch.Consume(
// 		nazwaKolejki, // name
// 		"",           // consumerTag, dostanę od serwera
// 		false,        // noAck
// 		false,        // exclusive
// 		false,        // noLocal
// 		false,        // noWait
// 		nil,          // arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("[RabbitMQ] błąd Consume(): %v", err)
// 	}

// 	go func(deliveries <-chan amqp.Delivery) {
// 		for d := range deliveries {
// 			err := handler(d.Body)
// 			if err != nil {
// 				d.Nack( /*multiple*/ false /*requeue*/, false)
// 				//d.Reject( /*requeue*/ false)
// 			} else {
// 				d.Ack( /*multiple*/ false)
// 			}
// 		}
// 		log.Printf("[RabbitMQ] pobieranie zakończone: kanał zamknięty")
// 	}(deliveries)
// 	log.Printf("[RabbitMQ.Cons] <~~ [%s]", nazwaKolejki)
// }

// // Wyslij do exchangera w Rabbicie
// func (krolik *Krolik) Wyslij(klucz string, dane []byte) error {
// 	return krolik.ch.Publish(
// 		krolik.ex, // publikuj do exchangera
// 		klucz,     // routuj do odpowiednich kolejek
// 		false,     // nie interesuje nas czy jest jakaś kolejka
// 		false,     // nie interesuje nas czy jest jakiś konsumer
// 		amqp.Publishing{
// 			Body:         dane,
// 			DeliveryMode: amqp.Persistent,
// 		})
// }

// // MusiWybierajOdrzucone podłącza handler do kolejki odrzuconych
// // dla danego klucza
// func (krolik *Krolik) MusiWybierajOdrzucone(klucz string, handler func([]byte)) {

// 	nazwOdrzuc := krolik.nazwaOdrzucone(klucz)
// 	deliveries, err := krolik.ch.Consume(
// 		nazwOdrzuc, // nazwa kolejki
// 		"",         // consumerTag, dostanę od serwera
// 		false,      // noAck
// 		false,      // exclusive
// 		false,      // noLocal
// 		false,      // noWait
// 		nil,        // arguments
// 	)
// 	if err != nil {
// 		log.Fatalf("[RabbitMQ] błąd Consume(): %v", err)
// 	}

// 	go func(deliveries <-chan amqp.Delivery) {
// 		for d := range deliveries {
// 			handler(d.Body)
// 			// odrzucone zabrane na dobre
// 			d.Ack(false)
// 		}
// 		log.Printf("[RabbitMQ] pobieranie zakończone: kanał zamknięty")
// 		// TODO: Jakoś trzebaby gdzieś zareagować na powyższe?
// 	}(deliveries)
// 	log.Printf("[RabbitMQ.Cons] <~~ [%s]", nazwOdrzuc)

// }

// func (krolik *Krolik) nazwaKolejki(klucz string) string {
// 	return fmt.Sprintf("%s %s", krolik.ex, klucz)
// }

// func (krolik *Krolik) nazwaOdrzucone(klucz string) string {
// 	return fmt.Sprintf("%s %s", krolik.dlx, klucz)
// }

// // Loguj middleware
// func Loguj(hand func([]byte) error) func([]byte) error {
// 	return func(dane []byte) error {
// 		st := time.Now()
// 		err := hand(dane)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		log.Printf("królik middleware %v", time.Since(st))
// 		return err
// 	}
// }
