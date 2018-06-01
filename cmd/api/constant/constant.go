package constant

// Constant define the parameters shard with all packages
type Constant struct {
	LifeAPIToken int64 `json:"lifeApiToken" valid:"required"`
}

var Public Constant

// Init save the constant's instance on global variable to expose it on all packages
func Init(c Constant) (err error) {
	Public = c
	return nil
}
