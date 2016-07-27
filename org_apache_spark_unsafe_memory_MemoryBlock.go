package gospark

import "github.com/timob/javabind"

type UnsafeMemoryMemoryBlockInterface interface {
	UnsafeMemoryMemoryLocationInterface

	// public long size()
	Size() int64

	// public static org.apache.spark.unsafe.memory.MemoryBlock fromLongArray(long[])
	FromLongArray(a []int64) *UnsafeMemoryMemoryBlock
}

type UnsafeMemoryMemoryBlock struct {
	UnsafeMemoryMemoryLocation
}

// public org.apache.spark.unsafe.memory.MemoryBlock(java.lang.Object, long, long)
func NewUnsafeMemoryMemoryBlock(a interface{}, b int64, c int64) (*UnsafeMemoryMemoryBlock) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/unsafe/memory/MemoryBlock", conv_a.Value().Cast("java/lang/Object"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &UnsafeMemoryMemoryBlock{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long size()
func (jbobject *UnsafeMemoryMemoryBlock) Size() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public static org.apache.spark.unsafe.memory.MemoryBlock fromLongArray(long[])
func (jbobject *UnsafeMemoryMemoryBlock) FromLongArray(a []int64) *UnsafeMemoryMemoryBlock {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/unsafe/memory/MemoryBlock", "fromLongArray", "org/apache/spark/unsafe/memory/MemoryBlock", a)
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

func (jbobject *UnsafeMemoryMemoryBlock) PageNumber() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "pageNumber", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *UnsafeMemoryMemoryBlock) SetFieldPageNumber(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "pageNumber", val)
	if err != nil {
		panic(err)
	}

}


