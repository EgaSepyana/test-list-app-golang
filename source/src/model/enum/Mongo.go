package enum

type MongoCollection int

const (
	MongoCollection_User MongoCollection = iota
	MongoCollection_Checklist
	MongoCollection_ChecklistItem
)

func (index MongoCollection) String() string {
	return []string{
		"user",
		"checklist",
		"checklist_item",
	}[index]
}
