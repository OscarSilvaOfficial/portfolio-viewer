package ports

type NoSQLPort interface {
	CreateDocument(document map[string]string)
}