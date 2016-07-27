package gospark

import "github.com/timob/javabind"

type BroadcastBroadcastInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public long id()
	Id() int64

	// public T value()
	Value() *JavaLangObject

	// public void unpersist()
	Unpersist() 

	// public void unpersist(boolean)
	Unpersist2(a bool) 

	// public void destroy()
	Destroy() 

	// public void destroy(boolean)
	Destroy2(a bool) 

	// public boolean isValid()
	IsValid() bool

	// public abstract T getValue()
	GetValue() *JavaLangObject

	// public abstract void doUnpersist(boolean)
	DoUnpersist(a bool) 

	// public abstract void doDestroy(boolean)
	DoDestroy(a bool) 

	// public void assertValid()
	AssertValid() 
}

type BroadcastBroadcast struct {
	JavaLangObject
}

// public java.lang.String logName()
func (jbobject *BroadcastBroadcast) LogName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "logName", "java/lang/String")
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

// public boolean isTraceEnabled()
func (jbobject *BroadcastBroadcast) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long id()
func (jbobject *BroadcastBroadcast) Id() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public T value()
func (jbobject *BroadcastBroadcast) Value() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "value", "java/lang/Object")
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

// public void unpersist()
func (jbobject *BroadcastBroadcast) Unpersist()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void unpersist(boolean)
func (jbobject *BroadcastBroadcast) Unpersist2(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void destroy()
func (jbobject *BroadcastBroadcast) Destroy()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "destroy", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void destroy(boolean)
func (jbobject *BroadcastBroadcast) Destroy2(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "destroy", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean isValid()
func (jbobject *BroadcastBroadcast) IsValid() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isValid", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract T getValue()
func (jbobject *BroadcastBroadcast) GetValue() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "java/lang/Object")
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

// public abstract void doUnpersist(boolean)
func (jbobject *BroadcastBroadcast) DoUnpersist(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "doUnpersist", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public abstract void doDestroy(boolean)
func (jbobject *BroadcastBroadcast) DoDestroy(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "doDestroy", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void assertValid()
func (jbobject *BroadcastBroadcast) AssertValid()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "assertValid", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String toString()
func (jbobject *BroadcastBroadcast) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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


