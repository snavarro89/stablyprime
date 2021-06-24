package app

/*Main struct to keep track of configuration settings required for the server
  one of the properties is data, which will container the interfaces for the
  data layer, this way we could init our app with a different source data.
*/
type App struct {
	Data Data
}

type Numbers interface {
	GetPrime(number int) (int, error)
}

type Data struct {
	Numbers
}
