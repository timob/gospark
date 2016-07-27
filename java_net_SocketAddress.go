package gospark

import "github.com/timob/javabind"

type JavaNetSocketAddressInterface interface {
	JavaLangObjectInterface
}

type JavaNetSocketAddress struct {
	JavaLangObject
}

// public java.net.SocketAddress()
func NewJavaNetSocketAddress() (*JavaNetSocketAddress) {

	obj, err := javabind.GetEnv().NewObject("java/net/SocketAddress")
	if err != nil {
		panic(err)
	}
	x := &JavaNetSocketAddress{}
	x.Callable = &javabind.Callable{obj}
	return x
}


