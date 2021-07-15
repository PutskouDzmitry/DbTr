package const_db

const (
	//Const for default for initialize connection
	ServerPort = "8080"
	Host       = "localhost"
	Port       = "5432"
	User       = "postgres"
	DbName     = "postgres"
	Password   = "1234"
	Sslmode    = "disable"
	//Tables const
	Publishers = "publishers"
	Books      = "books"
	//Errors const
	CantAddDataError      = "can't add data in database, error: %w"
	CantUpdateDataError   = "can't update information in db, error: %w"
	CantDeleteDataError   = "can't delete data in db, error: %w"
	TroubleWithConnection = "got an error when tried to make connection with database:%w"
	//Other const)
	AddInfoForConnection = "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
)
