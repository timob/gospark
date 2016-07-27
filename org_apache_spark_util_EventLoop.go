package gospark

import "github.com/timob/javabind"

type UtilEventLoopInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public void start()
	Start() 

	// public void stop()
	Stop() 

	// public void post(E)
	Post(a JavaLangObjectInterface) 

	// public boolean isActive()
	IsActive() bool

	// public void onStart()
	OnStart() 

	// public void onStop()
	OnStop() 

	// public abstract void onReceive(E)
	OnReceive(a JavaLangObjectInterface) 
}

type UtilEventLoop struct {
	JavaLangObject
}

// public org.apache.spark.util.EventLoop(java.lang.String)
func NewUtilEventLoop(a string) (*UtilEventLoop) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/util/EventLoop", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &UtilEventLoop{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *UtilEventLoop) LogName() string {
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
func (jbobject *UtilEventLoop) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void start()
func (jbobject *UtilEventLoop) Start()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "start", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void stop()
func (jbobject *UtilEventLoop) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void post(E)
func (jbobject *UtilEventLoop) Post(a JavaLangObjectInterface)  {
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

// public boolean isActive()
func (jbobject *UtilEventLoop) IsActive() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isActive", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void onStart()
func (jbobject *UtilEventLoop) OnStart()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStart", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void onStop()
func (jbobject *UtilEventLoop) OnStop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void onReceive(E)
func (jbobject *UtilEventLoop) OnReceive(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onReceive", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


