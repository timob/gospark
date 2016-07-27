package gospark

import "github.com/timob/javabind"

type UtilAsynchronousListenerBusInterface interface {
	JavaLangObjectInterface

	// public final void addListener(L)
	AddListener(a JavaLangObjectInterface) 

	// public final void postToAll(E)
	PostToAll(a JavaLangObjectInterface) 

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public void start(org.apache.spark.SparkContext)
	Start(a SparkContextInterface) 

	// public void post(E)
	Post(a JavaLangObjectInterface) 

	// public boolean listenerThreadIsAlive()
	ListenerThreadIsAlive() bool

	// public void stop()
	Stop() 

	// public abstract void onDropEvent(E)
	OnDropEvent(a JavaLangObjectInterface) 
}

type UtilAsynchronousListenerBus struct {
	JavaLangObject
}

// public org.apache.spark.util.AsynchronousListenerBus(java.lang.String)
func NewUtilAsynchronousListenerBus(a string) (*UtilAsynchronousListenerBus) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/util/AsynchronousListenerBus", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &UtilAsynchronousListenerBus{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final void addListener(L)
func (jbobject *UtilAsynchronousListenerBus) AddListener(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addListener", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final void postToAll(E)
func (jbobject *UtilAsynchronousListenerBus) PostToAll(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "postToAll", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String logName()
func (jbobject *UtilAsynchronousListenerBus) LogName() string {
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
func (jbobject *UtilAsynchronousListenerBus) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void start(org.apache.spark.SparkContext)
func (jbobject *UtilAsynchronousListenerBus) Start(a SparkContextInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "start", javabind.Void, conv_a.Value().Cast("org/apache/spark/SparkContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void post(E)
func (jbobject *UtilAsynchronousListenerBus) Post(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "post", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void waitUntilEmpty(long) throws java.util.concurrent.TimeoutException
func (jbobject *UtilAsynchronousListenerBus) WaitUntilEmpty(a int64) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "waitUntilEmpty", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public boolean listenerThreadIsAlive()
func (jbobject *UtilAsynchronousListenerBus) ListenerThreadIsAlive() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "listenerThreadIsAlive", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void stop()
func (jbobject *UtilAsynchronousListenerBus) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void onDropEvent(E)
func (jbobject *UtilAsynchronousListenerBus) OnDropEvent(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onDropEvent", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


