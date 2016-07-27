package gospark

import "github.com/timob/javabind"

type JavaLangObjectInterface interface {

	// public native int hashCode()
	HashCode() int

	// public boolean equals(java.lang.Object)
	Equals(a interface{}) bool

	// public java.lang.String toString()
	ToString() string

	// public final native void notify()
	Notify() 

	// public final native void notifyAll()
	NotifyAll() 
}

type JavaLangObject struct {
	*javabind.Callable
}

// public java.lang.Object()
func NewJavaLangObject() (*JavaLangObject) {

	obj, err := javabind.GetEnv().NewObject("java/lang/Object")
	if err != nil {
		panic(err)
	}
	x := &JavaLangObject{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public native int hashCode()
func (jbobject *JavaLangObject) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *JavaLangObject) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *JavaLangObject) ToString() string {
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

// public final native void notify()
func (jbobject *JavaLangObject) Notify()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "notify", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public final native void notifyAll()
func (jbobject *JavaLangObject) NotifyAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "notifyAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public final native void wait(long) throws java.lang.InterruptedException
func (jbobject *JavaLangObject) Wait2(a int64) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "wait", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public final void wait(long, int) throws java.lang.InterruptedException
func (jbobject *JavaLangObject) Wait3(a int64, b int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "wait", javabind.Void, a, b)
	if err != nil {
		return err
	}
	return nil
}

// public final void wait() throws java.lang.InterruptedException
func (jbobject *JavaLangObject) Wait() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "wait", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}


