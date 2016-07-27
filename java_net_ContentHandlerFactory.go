package gospark

import "github.com/timob/javabind"

type JavaNetContentHandlerFactoryInterface interface {

	// public abstract java.net.ContentHandler createContentHandler(java.lang.String)
	CreateContentHandler(a string) *JavaNetContentHandler
}

type JavaNetContentHandlerFactory struct {
	JavaLangObject
}

// public abstract java.net.ContentHandler createContentHandler(java.lang.String)
func (jbobject *JavaNetContentHandlerFactory) CreateContentHandler(a string) *JavaNetContentHandler {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createContentHandler", "java/net/ContentHandler", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaNetContentHandler{}
	unique_x.Callable = dst
	return unique_x
}


