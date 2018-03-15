
// Code generated automatiko; DO NOT EDIT.
package kostka2

func (k *Kostka) WszystkieRuchy() []func() {
	return []func() {
		k.ObrYGoraA,
		k.ObrYGoraB,
		k.ObrYDolA,
		k.ObrYDolB,
		k.ObrXLewoA,
		k.ObrXLewoB,
		k.ObrXPrawoA,
		k.ObrXPrawoB,
		k.ObrZPrzodA,
		k.ObrZPrzodB,
		k.ObrZTylA,
		k.ObrZTylB,
	}
}
func (k *Kostka) ObrYGoraA() {
	buf = k.Kwadraciki[0]
	k.Kwadraciki[0] = k.Kwadraciki[5]
	k.Kwadraciki[5] = k.Kwadraciki[7]
	k.Kwadraciki[7] = k.Kwadraciki[2]
	k.Kwadraciki[2] = buf
	buf = k.Kwadraciki[3]
	k.Kwadraciki[3] = k.Kwadraciki[6]
	k.Kwadraciki[6] = k.Kwadraciki[4]
	k.Kwadraciki[4] = k.Kwadraciki[1]
	k.Kwadraciki[1] = buf
	buf = k.Kwadraciki[8]
	k.Kwadraciki[8] = k.Kwadraciki[16]
	k.Kwadraciki[16] = k.Kwadraciki[24]
	k.Kwadraciki[24] = k.Kwadraciki[47]
	k.Kwadraciki[47] = buf
	buf = k.Kwadraciki[9]
	k.Kwadraciki[9] = k.Kwadraciki[17]
	k.Kwadraciki[17] = k.Kwadraciki[25]
	k.Kwadraciki[25] = k.Kwadraciki[46]
	k.Kwadraciki[46] = buf
	buf = k.Kwadraciki[10]
	k.Kwadraciki[10] = k.Kwadraciki[18]
	k.Kwadraciki[18] = k.Kwadraciki[26]
	k.Kwadraciki[26] = k.Kwadraciki[45]
	k.Kwadraciki[45] = buf
}
func (k *Kostka) ObrYGoraB() {
	buf = k.Kwadraciki[2]
	k.Kwadraciki[2] = k.Kwadraciki[7]
	k.Kwadraciki[7] = k.Kwadraciki[5]
	k.Kwadraciki[5] = k.Kwadraciki[0]
	k.Kwadraciki[0] = buf
	buf = k.Kwadraciki[1]
	k.Kwadraciki[1] = k.Kwadraciki[4]
	k.Kwadraciki[4] = k.Kwadraciki[6]
	k.Kwadraciki[6] = k.Kwadraciki[3]
	k.Kwadraciki[3] = buf
	buf = k.Kwadraciki[47]
	k.Kwadraciki[47] = k.Kwadraciki[24]
	k.Kwadraciki[24] = k.Kwadraciki[16]
	k.Kwadraciki[16] = k.Kwadraciki[8]
	k.Kwadraciki[8] = buf
	buf = k.Kwadraciki[46]
	k.Kwadraciki[46] = k.Kwadraciki[25]
	k.Kwadraciki[25] = k.Kwadraciki[17]
	k.Kwadraciki[17] = k.Kwadraciki[9]
	k.Kwadraciki[9] = buf
	buf = k.Kwadraciki[45]
	k.Kwadraciki[45] = k.Kwadraciki[26]
	k.Kwadraciki[26] = k.Kwadraciki[18]
	k.Kwadraciki[18] = k.Kwadraciki[10]
	k.Kwadraciki[10] = buf
}
func (k *Kostka) ObrYDolA() {
	buf = k.Kwadraciki[32]
	k.Kwadraciki[32] = k.Kwadraciki[37]
	k.Kwadraciki[37] = k.Kwadraciki[39]
	k.Kwadraciki[39] = k.Kwadraciki[34]
	k.Kwadraciki[34] = buf
	buf = k.Kwadraciki[35]
	k.Kwadraciki[35] = k.Kwadraciki[38]
	k.Kwadraciki[38] = k.Kwadraciki[36]
	k.Kwadraciki[36] = k.Kwadraciki[33]
	k.Kwadraciki[33] = buf
	buf = k.Kwadraciki[13]
	k.Kwadraciki[13] = k.Kwadraciki[21]
	k.Kwadraciki[21] = k.Kwadraciki[29]
	k.Kwadraciki[29] = k.Kwadraciki[42]
	k.Kwadraciki[42] = buf
	buf = k.Kwadraciki[14]
	k.Kwadraciki[14] = k.Kwadraciki[22]
	k.Kwadraciki[22] = k.Kwadraciki[30]
	k.Kwadraciki[30] = k.Kwadraciki[41]
	k.Kwadraciki[41] = buf
	buf = k.Kwadraciki[15]
	k.Kwadraciki[15] = k.Kwadraciki[23]
	k.Kwadraciki[23] = k.Kwadraciki[31]
	k.Kwadraciki[31] = k.Kwadraciki[40]
	k.Kwadraciki[40] = buf
}
func (k *Kostka) ObrYDolB() {
	buf = k.Kwadraciki[34]
	k.Kwadraciki[34] = k.Kwadraciki[39]
	k.Kwadraciki[39] = k.Kwadraciki[37]
	k.Kwadraciki[37] = k.Kwadraciki[32]
	k.Kwadraciki[32] = buf
	buf = k.Kwadraciki[33]
	k.Kwadraciki[33] = k.Kwadraciki[36]
	k.Kwadraciki[36] = k.Kwadraciki[38]
	k.Kwadraciki[38] = k.Kwadraciki[35]
	k.Kwadraciki[35] = buf
	buf = k.Kwadraciki[42]
	k.Kwadraciki[42] = k.Kwadraciki[29]
	k.Kwadraciki[29] = k.Kwadraciki[21]
	k.Kwadraciki[21] = k.Kwadraciki[13]
	k.Kwadraciki[13] = buf
	buf = k.Kwadraciki[41]
	k.Kwadraciki[41] = k.Kwadraciki[30]
	k.Kwadraciki[30] = k.Kwadraciki[22]
	k.Kwadraciki[22] = k.Kwadraciki[14]
	k.Kwadraciki[14] = buf
	buf = k.Kwadraciki[40]
	k.Kwadraciki[40] = k.Kwadraciki[31]
	k.Kwadraciki[31] = k.Kwadraciki[23]
	k.Kwadraciki[23] = k.Kwadraciki[15]
	k.Kwadraciki[15] = buf
}
func (k *Kostka) ObrXLewoA() {
	buf = k.Kwadraciki[0]
	k.Kwadraciki[0] = k.Kwadraciki[16]
	k.Kwadraciki[16] = k.Kwadraciki[32]
	k.Kwadraciki[32] = k.Kwadraciki[40]
	k.Kwadraciki[40] = buf
	buf = k.Kwadraciki[3]
	k.Kwadraciki[3] = k.Kwadraciki[19]
	k.Kwadraciki[19] = k.Kwadraciki[35]
	k.Kwadraciki[35] = k.Kwadraciki[43]
	k.Kwadraciki[43] = buf
	buf = k.Kwadraciki[5]
	k.Kwadraciki[5] = k.Kwadraciki[21]
	k.Kwadraciki[21] = k.Kwadraciki[37]
	k.Kwadraciki[37] = k.Kwadraciki[45]
	k.Kwadraciki[45] = buf
	buf = k.Kwadraciki[8]
	k.Kwadraciki[8] = k.Kwadraciki[13]
	k.Kwadraciki[13] = k.Kwadraciki[15]
	k.Kwadraciki[15] = k.Kwadraciki[10]
	k.Kwadraciki[10] = buf
	buf = k.Kwadraciki[11]
	k.Kwadraciki[11] = k.Kwadraciki[14]
	k.Kwadraciki[14] = k.Kwadraciki[12]
	k.Kwadraciki[12] = k.Kwadraciki[9]
	k.Kwadraciki[9] = buf
}
func (k *Kostka) ObrXLewoB() {
	buf = k.Kwadraciki[40]
	k.Kwadraciki[40] = k.Kwadraciki[32]
	k.Kwadraciki[32] = k.Kwadraciki[16]
	k.Kwadraciki[16] = k.Kwadraciki[0]
	k.Kwadraciki[0] = buf
	buf = k.Kwadraciki[43]
	k.Kwadraciki[43] = k.Kwadraciki[35]
	k.Kwadraciki[35] = k.Kwadraciki[19]
	k.Kwadraciki[19] = k.Kwadraciki[3]
	k.Kwadraciki[3] = buf
	buf = k.Kwadraciki[45]
	k.Kwadraciki[45] = k.Kwadraciki[37]
	k.Kwadraciki[37] = k.Kwadraciki[21]
	k.Kwadraciki[21] = k.Kwadraciki[5]
	k.Kwadraciki[5] = buf
	buf = k.Kwadraciki[10]
	k.Kwadraciki[10] = k.Kwadraciki[15]
	k.Kwadraciki[15] = k.Kwadraciki[13]
	k.Kwadraciki[13] = k.Kwadraciki[8]
	k.Kwadraciki[8] = buf
	buf = k.Kwadraciki[9]
	k.Kwadraciki[9] = k.Kwadraciki[12]
	k.Kwadraciki[12] = k.Kwadraciki[14]
	k.Kwadraciki[14] = k.Kwadraciki[11]
	k.Kwadraciki[11] = buf
}
func (k *Kostka) ObrXPrawoA() {
	buf = k.Kwadraciki[2]
	k.Kwadraciki[2] = k.Kwadraciki[18]
	k.Kwadraciki[18] = k.Kwadraciki[34]
	k.Kwadraciki[34] = k.Kwadraciki[42]
	k.Kwadraciki[42] = buf
	buf = k.Kwadraciki[4]
	k.Kwadraciki[4] = k.Kwadraciki[20]
	k.Kwadraciki[20] = k.Kwadraciki[36]
	k.Kwadraciki[36] = k.Kwadraciki[44]
	k.Kwadraciki[44] = buf
	buf = k.Kwadraciki[7]
	k.Kwadraciki[7] = k.Kwadraciki[23]
	k.Kwadraciki[23] = k.Kwadraciki[39]
	k.Kwadraciki[39] = k.Kwadraciki[47]
	k.Kwadraciki[47] = buf
	buf = k.Kwadraciki[24]
	k.Kwadraciki[24] = k.Kwadraciki[29]
	k.Kwadraciki[29] = k.Kwadraciki[31]
	k.Kwadraciki[31] = k.Kwadraciki[26]
	k.Kwadraciki[26] = buf
	buf = k.Kwadraciki[27]
	k.Kwadraciki[27] = k.Kwadraciki[30]
	k.Kwadraciki[30] = k.Kwadraciki[28]
	k.Kwadraciki[28] = k.Kwadraciki[25]
	k.Kwadraciki[25] = buf
}
func (k *Kostka) ObrXPrawoB() {
	buf = k.Kwadraciki[42]
	k.Kwadraciki[42] = k.Kwadraciki[34]
	k.Kwadraciki[34] = k.Kwadraciki[18]
	k.Kwadraciki[18] = k.Kwadraciki[2]
	k.Kwadraciki[2] = buf
	buf = k.Kwadraciki[44]
	k.Kwadraciki[44] = k.Kwadraciki[36]
	k.Kwadraciki[36] = k.Kwadraciki[20]
	k.Kwadraciki[20] = k.Kwadraciki[4]
	k.Kwadraciki[4] = buf
	buf = k.Kwadraciki[47]
	k.Kwadraciki[47] = k.Kwadraciki[39]
	k.Kwadraciki[39] = k.Kwadraciki[23]
	k.Kwadraciki[23] = k.Kwadraciki[7]
	k.Kwadraciki[7] = buf
	buf = k.Kwadraciki[26]
	k.Kwadraciki[26] = k.Kwadraciki[31]
	k.Kwadraciki[31] = k.Kwadraciki[29]
	k.Kwadraciki[29] = k.Kwadraciki[24]
	k.Kwadraciki[24] = buf
	buf = k.Kwadraciki[25]
	k.Kwadraciki[25] = k.Kwadraciki[28]
	k.Kwadraciki[28] = k.Kwadraciki[30]
	k.Kwadraciki[30] = k.Kwadraciki[27]
	k.Kwadraciki[27] = buf
}
func (k *Kostka) ObrZPrzodA() {
	buf = k.Kwadraciki[16]
	k.Kwadraciki[16] = k.Kwadraciki[21]
	k.Kwadraciki[21] = k.Kwadraciki[23]
	k.Kwadraciki[23] = k.Kwadraciki[18]
	k.Kwadraciki[18] = buf
	buf = k.Kwadraciki[19]
	k.Kwadraciki[19] = k.Kwadraciki[22]
	k.Kwadraciki[22] = k.Kwadraciki[20]
	k.Kwadraciki[20] = k.Kwadraciki[17]
	k.Kwadraciki[17] = buf
	buf = k.Kwadraciki[5]
	k.Kwadraciki[5] = k.Kwadraciki[24]
	k.Kwadraciki[24] = k.Kwadraciki[34]
	k.Kwadraciki[34] = k.Kwadraciki[15]
	k.Kwadraciki[15] = buf
	buf = k.Kwadraciki[6]
	k.Kwadraciki[6] = k.Kwadraciki[27]
	k.Kwadraciki[27] = k.Kwadraciki[33]
	k.Kwadraciki[33] = k.Kwadraciki[12]
	k.Kwadraciki[12] = buf
	buf = k.Kwadraciki[7]
	k.Kwadraciki[7] = k.Kwadraciki[29]
	k.Kwadraciki[29] = k.Kwadraciki[32]
	k.Kwadraciki[32] = k.Kwadraciki[10]
	k.Kwadraciki[10] = buf
}
func (k *Kostka) ObrZPrzodB() {
	buf = k.Kwadraciki[18]
	k.Kwadraciki[18] = k.Kwadraciki[23]
	k.Kwadraciki[23] = k.Kwadraciki[21]
	k.Kwadraciki[21] = k.Kwadraciki[16]
	k.Kwadraciki[16] = buf
	buf = k.Kwadraciki[17]
	k.Kwadraciki[17] = k.Kwadraciki[20]
	k.Kwadraciki[20] = k.Kwadraciki[22]
	k.Kwadraciki[22] = k.Kwadraciki[19]
	k.Kwadraciki[19] = buf
	buf = k.Kwadraciki[15]
	k.Kwadraciki[15] = k.Kwadraciki[34]
	k.Kwadraciki[34] = k.Kwadraciki[24]
	k.Kwadraciki[24] = k.Kwadraciki[5]
	k.Kwadraciki[5] = buf
	buf = k.Kwadraciki[12]
	k.Kwadraciki[12] = k.Kwadraciki[33]
	k.Kwadraciki[33] = k.Kwadraciki[27]
	k.Kwadraciki[27] = k.Kwadraciki[6]
	k.Kwadraciki[6] = buf
	buf = k.Kwadraciki[10]
	k.Kwadraciki[10] = k.Kwadraciki[32]
	k.Kwadraciki[32] = k.Kwadraciki[29]
	k.Kwadraciki[29] = k.Kwadraciki[7]
	k.Kwadraciki[7] = buf
}
func (k *Kostka) ObrZTylA() {
	buf = k.Kwadraciki[40]
	k.Kwadraciki[40] = k.Kwadraciki[45]
	k.Kwadraciki[45] = k.Kwadraciki[47]
	k.Kwadraciki[47] = k.Kwadraciki[42]
	k.Kwadraciki[42] = buf
	buf = k.Kwadraciki[43]
	k.Kwadraciki[43] = k.Kwadraciki[46]
	k.Kwadraciki[46] = k.Kwadraciki[44]
	k.Kwadraciki[44] = k.Kwadraciki[41]
	k.Kwadraciki[41] = buf
	buf = k.Kwadraciki[37]
	k.Kwadraciki[37] = k.Kwadraciki[31]
	k.Kwadraciki[31] = k.Kwadraciki[2]
	k.Kwadraciki[2] = k.Kwadraciki[8]
	k.Kwadraciki[8] = buf
	buf = k.Kwadraciki[38]
	k.Kwadraciki[38] = k.Kwadraciki[28]
	k.Kwadraciki[28] = k.Kwadraciki[1]
	k.Kwadraciki[1] = k.Kwadraciki[11]
	k.Kwadraciki[11] = buf
	buf = k.Kwadraciki[39]
	k.Kwadraciki[39] = k.Kwadraciki[26]
	k.Kwadraciki[26] = k.Kwadraciki[0]
	k.Kwadraciki[0] = k.Kwadraciki[13]
	k.Kwadraciki[13] = buf
}
func (k *Kostka) ObrZTylB() {
	buf = k.Kwadraciki[42]
	k.Kwadraciki[42] = k.Kwadraciki[47]
	k.Kwadraciki[47] = k.Kwadraciki[45]
	k.Kwadraciki[45] = k.Kwadraciki[40]
	k.Kwadraciki[40] = buf
	buf = k.Kwadraciki[41]
	k.Kwadraciki[41] = k.Kwadraciki[44]
	k.Kwadraciki[44] = k.Kwadraciki[46]
	k.Kwadraciki[46] = k.Kwadraciki[43]
	k.Kwadraciki[43] = buf
	buf = k.Kwadraciki[8]
	k.Kwadraciki[8] = k.Kwadraciki[2]
	k.Kwadraciki[2] = k.Kwadraciki[31]
	k.Kwadraciki[31] = k.Kwadraciki[37]
	k.Kwadraciki[37] = buf
	buf = k.Kwadraciki[11]
	k.Kwadraciki[11] = k.Kwadraciki[1]
	k.Kwadraciki[1] = k.Kwadraciki[28]
	k.Kwadraciki[28] = k.Kwadraciki[38]
	k.Kwadraciki[38] = buf
	buf = k.Kwadraciki[13]
	k.Kwadraciki[13] = k.Kwadraciki[0]
	k.Kwadraciki[0] = k.Kwadraciki[26]
	k.Kwadraciki[26] = k.Kwadraciki[39]
	k.Kwadraciki[39] = buf
}
