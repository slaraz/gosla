package main

func (k Kostka) obrotYGoraLewo() {
	return
}

func (k *Kostka) obrotYGoraPrawo() {
	return
}

type ruch func()

func (k *Kostka) wszystkieRuchy() []ruch {
	return []ruch{
		k.obrotYGoraLewo,
		k.obrotYGoraPrawo,
		// ObrotYGoraPrawo()
		// ObrotYSrodekLewo()
		// ObrotYSrodekPrawo()
		// ObrotYDolLewo()
		// ObrotYDolPrawo()

		// ObrotXLewoGora()
		// ObrotXLewoDol()
		// ObrotXSrodekGora()
		// ObrotXSrodekDol()
		// ObrotXPrawoGora()
		// ObrotXPrawoDol()

		// ObrotZPrzodPrawo()
		// ObrotZPrzodLewo()
		// ObrotZSrodekPrawo()
		// ObrotZSrodekLewo()
		// ObrotZTylPrawo()
		// ObrotZTylLewo()
	}
}
