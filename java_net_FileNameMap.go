package gospark

import "github.com/timob/javabind"

type JavaNetFileNameMapInterface interface {

	// public abstract java.lang.String getContentTypeFor(java.lang.String)
	GetContentTypeFor(a string) string
}

type JavaNetFileNameMap struct {
	JavaLangObject
}

// public abstract java.lang.String getContentTypeFor(java.lang.String)
func (jbobject *JavaNetFileNameMap) GetContentTypeFor(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContentTypeFor", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


