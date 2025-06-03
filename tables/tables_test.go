package tables

import (
	"cube/core"
	"testing"
)

func TestInitphase1Table(t *testing.T) {
	Initphase1Table(false)
	//LoadPhase1Table("phase1.gob")

}

func TestInsertphase1TableItem(t *testing.T) {
	table, _ := Initphase1Table(false)
	c := core.NewCube()
	Insertphase1TableItem(&c, uint8(2), table)

	if table[0][0][0] != 2 {
		t.Error("table wasn't updated")
	}
	c.Move("b", 1)
	c.Move("d", 1)
	Insertphase1TableItem(&c, uint8(3), table)
	if table[1314][1048][303] != 3 {
		t.Error("table wasn't updated")
	}
}
