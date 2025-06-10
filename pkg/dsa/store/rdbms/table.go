// dbstrcut
package vrdbms

type Table struct {
	Name      string
	Num       byte
	Primary   string
	Columns   *[]Column
	recSize   uint16
	IndexRoot int
	Rows      int
}
type Column struct {
	name  string
	pos   byte
	dtype byte
}

type Row struct {
}
