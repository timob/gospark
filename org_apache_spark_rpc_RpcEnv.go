package gospark

import "github.com/timob/javabind"

type RpcRpcEnvInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.rpc.RpcEnv create(java.lang.String, java.lang.String, int, org.apache.spark.SparkConf, org.apache.spark.SecurityManager, boolean)
	Create(a string, b string, c int, d SparkConfInterface, e SecurityManagerInterface, f bool) *RpcRpcEnv

	// public org.apache.spark.rpc.RpcTimeout defaultLookupTimeout()
	DefaultLookupTimeout() *RpcRpcTimeout

	// public abstract org.apache.spark.rpc.RpcEndpointRef endpointRef(org.apache.spark.rpc.RpcEndpoint)
	EndpointRef(a RpcRpcEndpointInterface) *RpcRpcEndpointRef

	// public abstract org.apache.spark.rpc.RpcAddress address()
	Address() *RpcRpcAddress

	// public abstract org.apache.spark.rpc.RpcEndpointRef setupEndpoint(java.lang.String, org.apache.spark.rpc.RpcEndpoint)
	SetupEndpoint(a string, b RpcRpcEndpointInterface) *RpcRpcEndpointRef

	// public org.apache.spark.rpc.RpcEndpointRef setupEndpointRefByURI(java.lang.String)
	SetupEndpointRefByURI(a string) *RpcRpcEndpointRef

	// public org.apache.spark.rpc.RpcEndpointRef setupEndpointRef(java.lang.String, org.apache.spark.rpc.RpcAddress, java.lang.String)
	SetupEndpointRef(a string, b RpcRpcAddressInterface, c string) *RpcRpcEndpointRef

	// public abstract void stop(org.apache.spark.rpc.RpcEndpointRef)
	Stop(a RpcRpcEndpointRefInterface) 

	// public abstract void shutdown()
	Shutdown() 

	// public abstract void awaitTermination()
	AwaitTermination() 

	// public abstract java.lang.String uriOf(java.lang.String, org.apache.spark.rpc.RpcAddress, java.lang.String)
	UriOf(a string, b RpcRpcAddressInterface, c string) string

	// public abstract org.apache.spark.rpc.RpcEnvFileServer fileServer()
	FileServer() *RpcRpcEnvFileServer
}

type RpcRpcEnv struct {
	JavaLangObject
}

// public org.apache.spark.rpc.RpcEnv(org.apache.spark.SparkConf)
func NewRpcRpcEnv(a SparkConfInterface) (*RpcRpcEnv) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/rpc/RpcEnv", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &RpcRpcEnv{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.rpc.RpcEnv create(java.lang.String, java.lang.String, int, org.apache.spark.SparkConf, org.apache.spark.SecurityManager, boolean)
func (jbobject *RpcRpcEnv) Create(a string, b string, c int, d SparkConfInterface, e SecurityManagerInterface, f bool) *RpcRpcEnv {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/rpc/RpcEnv", "create", "org/apache/spark/rpc/RpcEnv", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), c, conv_d.Value().Cast("org/apache/spark/SparkConf"), conv_e.Value().Cast("org/apache/spark/SecurityManager"), f)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &RpcRpcEnv{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rpc.RpcTimeout defaultLookupTimeout()
func (jbobject *RpcRpcEnv) DefaultLookupTimeout() *RpcRpcTimeout {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultLookupTimeout", "org/apache/spark/rpc/RpcTimeout")
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
	unique_x := &RpcRpcTimeout{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.rpc.RpcEndpointRef endpointRef(org.apache.spark.rpc.RpcEndpoint)
func (jbobject *RpcRpcEnv) EndpointRef(a RpcRpcEndpointInterface) *RpcRpcEndpointRef {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "endpointRef", "org/apache/spark/rpc/RpcEndpointRef", conv_a.Value().Cast("org/apache/spark/rpc/RpcEndpoint"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &RpcRpcEndpointRef{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.rpc.RpcAddress address()
func (jbobject *RpcRpcEnv) Address() *RpcRpcAddress {
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

// public abstract org.apache.spark.rpc.RpcEndpointRef setupEndpoint(java.lang.String, org.apache.spark.rpc.RpcEndpoint)
func (jbobject *RpcRpcEnv) SetupEndpoint(a string, b RpcRpcEndpointInterface) *RpcRpcEndpointRef {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setupEndpoint", "org/apache/spark/rpc/RpcEndpointRef", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/rpc/RpcEndpoint"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &RpcRpcEndpointRef{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rpc.RpcEndpointRef setupEndpointRefByURI(java.lang.String)
func (jbobject *RpcRpcEnv) SetupEndpointRefByURI(a string) *RpcRpcEndpointRef {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setupEndpointRefByURI", "org/apache/spark/rpc/RpcEndpointRef", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &RpcRpcEndpointRef{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rpc.RpcEndpointRef setupEndpointRef(java.lang.String, org.apache.spark.rpc.RpcAddress, java.lang.String)
func (jbobject *RpcRpcEnv) SetupEndpointRef(a string, b RpcRpcAddressInterface, c string) *RpcRpcEndpointRef {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setupEndpointRef", "org/apache/spark/rpc/RpcEndpointRef", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/rpc/RpcAddress"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &RpcRpcEndpointRef{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void stop(org.apache.spark.rpc.RpcEndpointRef)
func (jbobject *RpcRpcEnv) Stop(a RpcRpcEndpointRefInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void, conv_a.Value().Cast("org/apache/spark/rpc/RpcEndpointRef"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void shutdown()
func (jbobject *RpcRpcEnv) Shutdown()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shutdown", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void awaitTermination()
func (jbobject *RpcRpcEnv) AwaitTermination()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "awaitTermination", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract java.lang.String uriOf(java.lang.String, org.apache.spark.rpc.RpcAddress, java.lang.String)
func (jbobject *RpcRpcEnv) UriOf(a string, b RpcRpcAddressInterface, c string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "uriOf", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/rpc/RpcAddress"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public abstract org.apache.spark.rpc.RpcEnvFileServer fileServer()
func (jbobject *RpcRpcEnv) FileServer() *RpcRpcEnvFileServer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "fileServer", "org/apache/spark/rpc/RpcEnvFileServer")
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
	unique_x := &RpcRpcEnvFileServer{}
	unique_x.Callable = dst
	return unique_x
}


