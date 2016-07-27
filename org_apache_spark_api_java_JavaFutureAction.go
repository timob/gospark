package gospark

import "github.com/timob/javabind"

type ApiJavaJavaFutureActionInterface interface {

	// public abstract java.util.List<java.lang.Integer> jobIds()
	JobIds() []int
}

type ApiJavaJavaFutureAction struct {
	JavaLangObject
}

// public abstract java.util.List<java.lang.Integer> jobIds()
func (jbobject *ApiJavaJavaFutureAction) JobIds() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jobIds", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoInteger())
	dst := new([]int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


