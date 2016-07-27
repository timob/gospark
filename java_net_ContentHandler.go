package gospark

import "github.com/timob/javabind"

type JavaNetContentHandlerInterface interface {
	JavaLangObjectInterface
}

type JavaNetContentHandler struct {
	JavaLangObject
}

// public java.net.ContentHandler()
func NewJavaNetContentHandler() (*JavaNetContentHandler) {

	obj, err := javabind.GetEnv().NewObject("java/net/ContentHandler")
	if err != nil {
		panic(err)
	}
	x := &JavaNetContentHandler{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract java.lang.Object getContent(java.net.URLConnection) throws java.io.IOException
func (jbobject *JavaNetContentHandler) GetContent(a JavaNetURLConnectionInterface) (*JavaLangObject, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContent", "java/lang/Object", conv_a.Value().Cast("java/net/URLConnection"))
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


