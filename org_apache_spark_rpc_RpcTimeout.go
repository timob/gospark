package gospark

import "github.com/timob/javabind"

type RpcRpcTimeoutInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.rpc.RpcTimeout apply(org.apache.spark.SparkConf, java.lang.String, java.lang.String)
	Apply2(a SparkConfInterface, b string, c string) *RpcRpcTimeout

	// public static org.apache.spark.rpc.RpcTimeout apply(org.apache.spark.SparkConf, java.lang.String)
	Apply(a SparkConfInterface, b string) *RpcRpcTimeout

	// public java.lang.String timeoutProp()
	TimeoutProp() string
}

type RpcRpcTimeout struct {
	JavaLangObject
}

// public static org.apache.spark.rpc.RpcTimeout apply(org.apache.spark.SparkConf, java.lang.String, java.lang.String)
func (jbobject *RpcRpcTimeout) Apply2(a SparkConfInterface, b string, c string) *RpcRpcTimeout {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
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
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/rpc/RpcTimeout", "apply", "org/apache/spark/rpc/RpcTimeout", conv_a.Value().Cast("org/apache/spark/SparkConf"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
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
	unique_x := &RpcRpcTimeout{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.rpc.RpcTimeout apply(org.apache.spark.SparkConf, java.lang.String)
func (jbobject *RpcRpcTimeout) Apply(a SparkConfInterface, b string) *RpcRpcTimeout {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/rpc/RpcTimeout", "apply", "org/apache/spark/rpc/RpcTimeout", conv_a.Value().Cast("org/apache/spark/SparkConf"), conv_b.Value().Cast("java/lang/String"))
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
	unique_x := &RpcRpcTimeout{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String timeoutProp()
func (jbobject *RpcRpcTimeout) TimeoutProp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "timeoutProp", "java/lang/String")
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


