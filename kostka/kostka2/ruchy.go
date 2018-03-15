
// Code generated automatiko; DO NOT EDIT.
package kostka2
func (k *Kostka) ObrótXGóraDół() {
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
