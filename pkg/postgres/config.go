package postgres

// Config datas to instanciate a sql client type Postrgres sql
type Config struct {
	User     string `required:"true"`
	Password string `required:"true"`
	DbName   string `required:"true" split_words:"true"`
	Ssl      string `required:"true"`
	Host     string `required:"true"`
	Port     string `required:"true"`
}
