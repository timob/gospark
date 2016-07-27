package gospark

import "github.com/timob/javabind"

type RpcRpcEndpointRefInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public abstract org.apache.spark.rpc.RpcAddress address()
	Address() *RpcRpcAddress

	// public abstract java.lang.String name()
	Name() string

	// public abstract void send(java.lang.Object)
	Send(a interface{}) 
}

type RpcRpcEndpointRef struct {
	JavaLangObject
}

// public org.apache.spark.rpc.RpcEndpointRef(org.apache.spark.SparkConf)
func NewRpcRpcEndpointRef(a SparkConfInterface) (*RpcRpcEndpointRef) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/rpc/RpcEndpointRef", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &RpcRpcEndpointRef{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *RpcRpcEndpointRef) LogName() string {
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
func (jbobject *RpcRpcEndpointRef) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract org.apache.spark.rpc.RpcAddress address()
func (jbobject *RpcRpcEndpointRef) Address() *RpcRpcAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "address", "org/apache/spark/rpc/RpcAddress")
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
	unique_x := &RpcRpcAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.lang.String name()
func (jbobject *RpcRpcEndpointRef) Name() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "name", "java/lang/String")
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

// public abstract void send(java.lang.Object)
func (jbobject *RpcRpcEndpointRef) Send(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "send", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


