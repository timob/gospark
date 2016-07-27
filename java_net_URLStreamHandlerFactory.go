package gospark

import "github.com/timob/javabind"

type JavaNetURLStreamHandlerFactoryInterface interface {

	// public abstract java.net.URLStreamHandler createURLStreamHandler(java.lang.String)
	CreateURLStreamHandler(a string) *JavaNetURLStreamHandler
}

type JavaNetURLStreamHandlerFactory struct {
	JavaLangObject
}

// public abstract java.net.URLStreamHandler createURLStreamHandler(java.lang.String)
func (jbobject *JavaNetURLStreamHandlerFactory) CreateURLStreamHandler(a string) *JavaNetURLStreamHandler {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createURLStreamHandler", "java/net/URLStreamHandler", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaNetURLStreamHandler{}
	unique_x.Callable = dst
	return unique_x
}


