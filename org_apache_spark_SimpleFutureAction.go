package gospark

import "github.com/timob/javabind"

type SimpleFutureActionInterface interface {
	JavaLangObjectInterface

	// public void cancel()
	Cancel() 

	// public boolean isCompleted()
	IsCompleted() bool

	// public boolean isCancelled()
	IsCancelled() bool
}

type SimpleFutureAction struct {
	JavaLangObject
}

// public T get() throws java.lang.Exception
func (jbobject *SimpleFutureAction) Get() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "get", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
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
	return unique_x, nil
}

// public void cancel()
func (jbobject *SimpleFutureAction) Cancel()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancel", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean isCompleted()
func (jbobject *SimpleFutureAction) IsCompleted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCompleted", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isCancelled()
func (jbobject *SimpleFutureAction) IsCancelled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCancelled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


