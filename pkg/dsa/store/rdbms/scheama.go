// scheama
package vrdbms

import (
	"log"
)

type Schema struct {
	Pageid           int
	name, user, pass string
	Pages            int
	HasTables        bool
	TableSize        byte
	TableMap         map[string]int
}

func newSchema(name string, user string, pass string) *Schema {
	if dbExists() {
		log.Println("Schema already existed")
		page := Read(0)
		log.Println("Page No:", page.Num)
		schema := &Schema{}
		return schema
	} else {
		return &Schema{
			Pageid: 0,
			name:   name,
			user:   user,
			pass:   pass,
		}
	}
}

func WriteSchema(schema *Schema) {

}
