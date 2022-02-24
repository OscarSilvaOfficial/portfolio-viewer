package ports

type NoSQLPort interface {
	createDocument(document map[string]string) map[string]string
}