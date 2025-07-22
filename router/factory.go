package router

// RouterType тип роутера (можно расширить)
type RouterType string

const (
	Gin RouterType = "gin"
	// Chi  RouterType = "chi"
	// Fiber RouterType = "fib"
)

// NewRouter factory for create routers
func NewRouter(rType RouterType) Router {
	switch rType {
	case Gin:
		return NewGinRouter()
	// case Chi:
	//   return NewChiRouter()
	// case Fiber:
	//   return NewFiberRouter()
	default:
		panic("unknown router type")
	}
}
