package gospark

import "github.com/timob/javabind"

type SchedulerJobListenerInterface interface {

	// public abstract void taskSucceeded(int, java.lang.Object)
	TaskSucceeded(a int, b interface{}) 
}

type SchedulerJobListener struct {
	JavaLangObject
}

// public abstract void taskSucceeded(int, java.lang.Object)
func (jbobject *SchedulerJobListener) TaskSucceeded(a int, b interface{})  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "taskSucceeded", javabind.Void, a, conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}


