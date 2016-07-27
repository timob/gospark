package gospark

import "github.com/timob/javabind"

type NetworkBufferManagedBufferInterface interface {
	JavaLangObjectInterface

	// public abstract long size()
	Size() int64

	// public abstract org.apache.spark.network.buffer.ManagedBuffer retain()
	Retain() *NetworkBufferManagedBuffer

	// public abstract org.apache.spark.network.buffer.ManagedBuffer release()
	Release() *NetworkBufferManagedBuffer
}

type NetworkBufferManagedBuffer struct {
	JavaLangObject
}

// public org.apache.spark.network.buffer.ManagedBuffer()
func NewNetworkBufferManagedBuffer() (*NetworkBufferManagedBuffer) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/network/buffer/ManagedBuffer")
	if err != nil {
		panic(err)
	}
	x := &NetworkBufferManagedBuffer{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract long size()
func (jbobject *NetworkBufferManagedBuffer) Size() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public abstract org.apache.spark.network.buffer.ManagedBuffer retain()
func (jbobject *NetworkBufferManagedBuffer) Retain() *NetworkBufferManagedBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "retain", "org/apache/spark/network/buffer/ManagedBuffer")
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
	unique_x := &NetworkBufferManagedBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.network.buffer.ManagedBuffer release()
func (jbobject *NetworkBufferManagedBuffer) Release() *NetworkBufferManagedBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "release", "org/apache/spark/network/buffer/ManagedBuffer")
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
	unique_x := &NetworkBufferManagedBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.lang.Object convertToNetty() throws java.io.IOException
func (jbobject *NetworkBufferManagedBuffer) ConvertToNetty() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "convertToNetty", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
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
	return unique_x, nil
}


