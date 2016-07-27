package gospark

import "github.com/timob/javabind"

type UnsafeArrayLongArrayInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.unsafe.memory.MemoryBlock memoryBlock()
	MemoryBlock() *UnsafeMemoryMemoryBlock

	// public java.lang.Object getBaseObject()
	GetBaseObject() *JavaLangObject

	// public long getBaseOffset()
	GetBaseOffset() int64

	// public long size()
	Size() int64

	// public void zeroOut()
	ZeroOut() 

	// public void set(int, long)
	Set(a int, b int64) 

	// public long get(int)
	Get(a int) int64
}

type UnsafeArrayLongArray struct {
	JavaLangObject
}

// public org.apache.spark.unsafe.array.LongArray(org.apache.spark.unsafe.memory.MemoryBlock)
func NewUnsafeArrayLongArray(a UnsafeMemoryMemoryBlockInterface) (*UnsafeArrayLongArray) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/unsafe/array/LongArray", conv_a.Value().Cast("org/apache/spark/unsafe/memory/MemoryBlock"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &UnsafeArrayLongArray{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.unsafe.memory.MemoryBlock memoryBlock()
func (jbobject *UnsafeArrayLongArray) MemoryBlock() *UnsafeMemoryMemoryBlock {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryBlock", "org/apache/spark/unsafe/memory/MemoryBlock")
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
	unique_x := &UnsafeMemoryMemoryBlock{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object getBaseObject()
func (jbobject *UnsafeArrayLongArray) GetBaseObject() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBaseObject", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public long getBaseOffset()
func (jbobject *UnsafeArrayLongArray) GetBaseOffset() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBaseOffset", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long size()
func (jbobject *UnsafeArrayLongArray) Size() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void zeroOut()
func (jbobject *UnsafeArrayLongArray) ZeroOut()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "zeroOut", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void set(int, long)
func (jbobject *UnsafeArrayLongArray) Set(a int, b int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "set", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public long get(int)
func (jbobject *UnsafeArrayLongArray) Get(a int) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "get", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


