package v2

import (
	"log"
	"testing"
	"time"
)

func Test_WYSLJI_GET_JEDEN(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, EX, "stdex", "fanout")
	defer mojex.Close()
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")
	defer mojqu.Close()

	wyslana := wiadomoscTestowa()
	if err := mojex.WyslijJSON(wyslana); err != nil {
		t.Fatal(err)
	}
	log.Printf("* wysłałem %q", wyslana.Id)

	time.Sleep(100 * time.Millisecond)

	odebrana := wiadomosc{}
	if err := mojqu.GetJSON(&odebrana); err != nil {
		t.Fatal(err)
	}
	log.Printf("* odebrałem %v", odebrana)

	if wyslana.Id != odebrana.Id {
		t.Errorf("błąd wysłałem: %v, odebrałem: %v", wyslana, odebrana)
	}
}

func Test_GET_100(t *testing.T) {
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")
	defer mojqu.Close()

	ile := 100
	odebrane := 0
	start := time.Now()
	for i := 1; i <= ile; i++ {
		odebrana := wiadomosc{}
		err := mojqu.GetJSON(&odebrana)
		if err != nil {
			t.Fatal(err)
		}
		if odebrana != wiadomoscPusta() {
			odebrane++
		}
		//log.Printf("* odebrałem %v", odebrana)
	}
	log.Printf("%d getów, odebrałem %d w %v", ile, odebrane, time.Since(start))

}
