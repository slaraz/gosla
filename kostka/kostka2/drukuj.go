package kostka2

import "fmt"

func (k *Kostka) Drukuj() {
	fmt.Printf("%20v %4v %4v\n", k.Kwadraciki[0], k.Kwadraciki[1], k.Kwadraciki[2])
	fmt.Printf("%20v CZAR %4v\n", k.Kwadraciki[3], k.Kwadraciki[4])
	fmt.Printf("%20v %4v %4v\n", k.Kwadraciki[5], k.Kwadraciki[6], k.Kwadraciki[7])
	fmt.Println()
	fmt.Printf("%4v %4v %4v  %4v %4v %4v  %4v %4v %4v\n",
		k.Kwadraciki[8], k.Kwadraciki[9], k.Kwadraciki[10],
		k.Kwadraciki[16], k.Kwadraciki[17], k.Kwadraciki[18],
		k.Kwadraciki[24], k.Kwadraciki[25], k.Kwadraciki[26])
	fmt.Printf("%4v ZIEL %4v  %4v CZER %4v  %4v NIEB %4v\n",
		k.Kwadraciki[11], k.Kwadraciki[12],
		k.Kwadraciki[19], k.Kwadraciki[20],
		k.Kwadraciki[27], k.Kwadraciki[28])
	fmt.Printf("%4v %4v %4v  %4v %4v %4v  %4v %4v %4v\n",
		k.Kwadraciki[13], k.Kwadraciki[14], k.Kwadraciki[15],
		k.Kwadraciki[21], k.Kwadraciki[22], k.Kwadraciki[23],
		k.Kwadraciki[29], k.Kwadraciki[30], k.Kwadraciki[31])
	fmt.Println()
	fmt.Printf("%20v %4v %4v\n", k.Kwadraciki[32], k.Kwadraciki[33], k.Kwadraciki[34])
	fmt.Printf("%20v ŻÓŁT %4v\n", k.Kwadraciki[35], k.Kwadraciki[36])
	fmt.Printf("%20v %4v %4v\n", k.Kwadraciki[37], k.Kwadraciki[38], k.Kwadraciki[39])
	fmt.Println()
	fmt.Printf("%20v %4v %4v\n", k.Kwadraciki[40], k.Kwadraciki[41], k.Kwadraciki[42])
	fmt.Printf("%20v POMA %4v\n", k.Kwadraciki[43], k.Kwadraciki[44])
	fmt.Printf("%20v %4v %4v\n", k.Kwadraciki[45], k.Kwadraciki[46], k.Kwadraciki[47])
	fmt.Println()
}
