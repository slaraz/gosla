package main

func (k *Kostka) obrotYGoraLewo() {
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
		// rotYSrodekLewo
		// rotYSrodekPrawo
		// rotYDolLewo
		// rotYDolPrawo

		// rotXLewoGora
		// rotXLewoDol
		// rotXSrodekGora
		// rotXSrodekDol
		// rotXPrawoGora
		// rotXPrawoDol

		// rotZPrzodPrawo
		// rotZPrzodLewo
		// rotZSrodekPrawo
		// rotZSrodekLewo
		// rotZTylPrawo
		// rotZTylLewo
	}
}
