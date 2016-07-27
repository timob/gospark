package gospark

import "github.com/timob/javabind"

type MetricsSourceSourceInterface interface {

	// public abstract java.lang.String sourceName()
	SourceName() string
}

type MetricsSourceSource struct {
	JavaLangObject
}

// public abstract java.lang.String sourceName()
func (jbobject *MetricsSourceSource) SourceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sourceName", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


