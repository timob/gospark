package gospark

import "github.com/timob/javabind"

type UnsafeMemoryMemoryLocationInterface interface {
	JavaLangObjectInterface

	// public void setObjAndOffset(java.lang.Object, long)
	SetObjAndOffset(a interface{}, b int64) 

	// public final java.lang.Object getBaseObject()
	GetBaseObject() *JavaLangObject

	// public final long getBaseOffset()
	GetBaseOffset() int64
}

type UnsafeMemoryMemoryLocation struct {
	JavaLangObject
}

// public org.apache.spark.unsafe.memory.MemoryLocation(java.lang.Object, long)
func NewUnsafeMemoryMemoryLocation2(a interface{}, b int64) (*UnsafeMemoryMemoryLocation) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/unsafe/memory/MemoryLocation", conv_a.Value().Cast("java/lang/Object"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &UnsafeMemoryMemoryLocation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.unsafe.memory.MemoryLocation()
func NewUnsafeMemoryMemoryLocation() (*UnsafeMemoryMemoryLocation) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/unsafe/memory/MemoryLocation")
	if err != nil {
		panic(err)
	}
	x := &UnsafeMemoryMemoryLocation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setObjAndOffset(java.lang.Object, long)
func (jbobject *UnsafeMemoryMemoryLocation) SetObjAndOffset(a interface{}, b int64)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setObjAndOffset", javabind.Void, conv_a.Value().Cast("java/lang/Object"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final java.lang.Object getBaseObject()
func (jbobject *UnsafeMemoryMemoryLocation) GetBaseObject() *JavaLangObject {
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

// public final long getBaseOffset()
func (jbobject *UnsafeMemoryMemoryLocation) GetBaseOffset() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBaseOffset", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


