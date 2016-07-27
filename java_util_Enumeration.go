package gospark

import "github.com/timob/javabind"

type JavaUtilEnumerationInterface interface {

	// public abstract boolean hasMoreElements()
	HasMoreElements() bool

	// public abstract E nextElement()
	NextElement() *JavaLangObject
}

type JavaUtilEnumeration struct {
	JavaLangObject
}

// public abstract boolean hasMoreElements()
func (jbobject *JavaUtilEnumeration) HasMoreElements() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasMoreElements", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract E nextElement()
func (jbobject *JavaUtilEnumeration) NextElement() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "nextElement", "java/lang/Object")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}


