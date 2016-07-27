package gospark

import "github.com/timob/javabind"

type JavaNetProxyInterface interface {
	JavaLangObjectInterface

	// public java.net.SocketAddress address()
	Address() *JavaNetSocketAddress
}

type JavaNetProxy struct {
	JavaLangObject
}

// public java.net.SocketAddress address()
func (jbobject *JavaNetProxy) Address() *JavaNetSocketAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "address", "java/net/SocketAddress")
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
	unique_x := &JavaNetSocketAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *JavaNetProxy) ToString() string {
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

// public final boolean equals(java.lang.Object)
func (jbobject *JavaNetProxy) Equals(a interface{}) bool {
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

// public final int hashCode()
func (jbobject *JavaNetProxy) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *JavaNetProxy) NO_PROXY() *JavaNetProxy {
	jret, err := javabind.GetEnv().GetStaticField("java/net/Proxy", "NO_PROXY", "java/net/Proxy")
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
	unique_x := &JavaNetProxy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaNetProxy) SetFieldNO_PROXY(val JavaNetProxyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/net/Proxy", "NO_PROXY", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


