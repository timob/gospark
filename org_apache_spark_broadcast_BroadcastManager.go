package gospark

import "github.com/timob/javabind"

type BroadcastBroadcastManagerInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public boolean isDriver()
	IsDriver() bool

	// public void stop()
	Stop() 

	// public void unbroadcast(long, boolean, boolean)
	Unbroadcast(a int64, b bool, c bool) 
}

type BroadcastBroadcastManager struct {
	JavaLangObject
}

// public org.apache.spark.broadcast.BroadcastManager(boolean, org.apache.spark.SparkConf, org.apache.spark.SecurityManager)
func NewBroadcastBroadcastManager(a bool, b SparkConfInterface, c SecurityManagerInterface) (*BroadcastBroadcastManager) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/broadcast/BroadcastManager", a, conv_b.Value().Cast("org/apache/spark/SparkConf"), conv_c.Value().Cast("org/apache/spark/SecurityManager"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &BroadcastBroadcastManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *BroadcastBroadcastManager) LogName() string {
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
func (jbobject *BroadcastBroadcastManager) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isDriver()
func (jbobject *BroadcastBroadcastManager) IsDriver() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDriver", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void stop()
func (jbobject *BroadcastBroadcastManager) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void unbroadcast(long, boolean, boolean)
func (jbobject *BroadcastBroadcastManager) Unbroadcast(a int64, b bool, c bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unbroadcast", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}


