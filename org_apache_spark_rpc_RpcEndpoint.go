package gospark

import "github.com/timob/javabind"

type RpcRpcEndpointInterface interface {

	// public abstract org.apache.spark.rpc.RpcEnv rpcEnv()
	RpcEnv() *RpcRpcEnv

	// public abstract org.apache.spark.rpc.RpcEndpointRef self()
	Self() *RpcRpcEndpointRef

	// public abstract void onConnected(org.apache.spark.rpc.RpcAddress)
	OnConnected(a RpcRpcAddressInterface) 

	// public abstract void onDisconnected(org.apache.spark.rpc.RpcAddress)
	OnDisconnected(a RpcRpcAddressInterface) 

	// public abstract void onStart()
	OnStart() 

	// public abstract void onStop()
	OnStop() 

	// public abstract void stop()
	Stop() 
}

type RpcRpcEndpoint struct {
	JavaLangObject
}

// public abstract org.apache.spark.rpc.RpcEnv rpcEnv()
func (jbobject *RpcRpcEndpoint) RpcEnv() *RpcRpcEnv {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rpcEnv", "org/apache/spark/rpc/RpcEnv")
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
	unique_x := &RpcRpcEnv{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.rpc.RpcEndpointRef self()
func (jbobject *RpcRpcEndpoint) Self() *RpcRpcEndpointRef {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "self", "org/apache/spark/rpc/RpcEndpointRef")
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
	unique_x := &RpcRpcEndpointRef{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void onConnected(org.apache.spark.rpc.RpcAddress)
func (jbobject *RpcRpcEndpoint) OnConnected(a RpcRpcAddressInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onConnected", javabind.Void, conv_a.Value().Cast("org/apache/spark/rpc/RpcAddress"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void onDisconnected(org.apache.spark.rpc.RpcAddress)
func (jbobject *RpcRpcEndpoint) OnDisconnected(a RpcRpcAddressInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onDisconnected", javabind.Void, conv_a.Value().Cast("org/apache/spark/rpc/RpcAddress"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void onStart()
func (jbobject *RpcRpcEndpoint) OnStart()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStart", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void onStop()
func (jbobject *RpcRpcEndpoint) OnStop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void stop()
func (jbobject *RpcRpcEndpoint) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


