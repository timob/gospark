package gospark

import "github.com/timob/javabind"

type JavaNetURLStreamHandlerInterface interface {
	JavaLangObjectInterface
}

type JavaNetURLStreamHandler struct {
	JavaLangObject
}

// public java.net.URLStreamHandler()
func NewJavaNetURLStreamHandler() (*JavaNetURLStreamHandler) {

	obj, err := javabind.GetEnv().NewObject("java/net/URLStreamHandler")
	if err != nil {
		panic(err)
	}
	x := &JavaNetURLStreamHandler{}
	x.Callable = &javabind.Callable{obj}
	return x
}


