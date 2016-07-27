package gospark

import "github.com/timob/javabind"

type JavaLangIterableInterface interface {

	// public abstract java.util.Iterator<T> iterator()
	Iterator() []*JavaLangObject
}

type JavaLangIterable struct {
	JavaLangObject
}

// public abstract java.util.Iterator<T> iterator()
func (jbobject *JavaLangIterable) Iterator() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "iterator", "java/util/Iterator")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoIterator(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


