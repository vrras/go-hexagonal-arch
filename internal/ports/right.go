package ports

// DBPort is the port for a db adapter
type DBPort interface {
	CloseDBConnection()
	AddToHistory(answer int32, operation string) error
}
