package app

/*Main struct to keep track of configuration settings required for the server
  one of the properties is data, which will container the interfaces for the
  data layer, this way we could init our app with a different source data.
*/
type App struct {
	Data Data
}

type Prime interface {
	Get(number int) int
}

type Data struct {
	Prime
}

type Parameters struct {
	Body []byte
}

type GetPrime interface {
	GetPrime(Parameters)
}
