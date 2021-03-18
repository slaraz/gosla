package krolik

import (
	"fmt"
	"testing"
	"time"
)

func TestMusiKrolik(t *testing.T) {
	// # serwer:
	// docker run -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management
	// # vhost:
	// curl -u guest:guest -X PUT http://localhost:15672/api/vhosts/krolik
	// # żeby wyczyścić:
	// curl -u guest:guest -X DELETE http://localhost:15672/api/vhosts/krolik

	krolikURL := "amqp://guest:guest@localhost:5672/test"
	krolikEX := "zwierzaki"
	zwierzaki := MusiKrolik(krolikURL, krolikEX)

	zwierzaki.MusiPobierac("xxx", workerXxx)
	zwierzaki.MusiPobierac("yyy", Loguj(workerYyy))
	zwierzaki.MusiWybierajOdrzucone("yyy", odrzuconeYyy)

	zwierzaki.Wyslij("xxx", []byte("Xenia kotka"))
	zwierzaki.Wyslij("yyy", []byte("yeż Yeży"))

	time.Sleep(time.Second)
}

func workerXxx(dane []byte) error {
	fmt.Println("pobrałem z xxx:", string(dane))
	return nil
}

func workerYyy(dane []byte) error {
	fmt.Println("odrzuciłem z yyy:", string(dane))
	return fmt.Errorf("błąd")
}

func odrzuconeYyy(dane []byte) {
	fmt.Println("wybrałem odrzucone yyy:", string(dane))
}
