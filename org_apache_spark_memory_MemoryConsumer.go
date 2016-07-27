package gospark

import "github.com/timob/javabind"

type MemoryMemoryConsumerInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.unsafe.array.LongArray allocateArray(long)
	AllocateArray(a int64) *UnsafeArrayLongArray

	// public void freeArray(org.apache.spark.unsafe.array.LongArray)
	FreeArray(a UnsafeArrayLongArrayInterface) 
}

type MemoryMemoryConsumer struct {
	JavaLangObject
}

// public void spill() throws java.io.IOException
func (jbobject *MemoryMemoryConsumer) Spill() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "spill", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public abstract long spill(long, org.apache.spark.memory.MemoryConsumer) throws java.io.IOException
func (jbobject *MemoryMemoryConsumer) Spill2(a int64, b MemoryMemoryConsumerInterface) (int64, error) {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "spill", javabind.Long, a, conv_b.Value().Cast("org/apache/spark/memory/MemoryConsumer"))
	if err != nil {
		var zero int64
		return zero, err
	}
	conv_b.CleanUp()
	return jret.(int64), nil
}

// public org.apache.spark.unsafe.array.LongArray allocateArray(long)
func (jbobject *MemoryMemoryConsumer) AllocateArray(a int64) *UnsafeArrayLongArray {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "allocateArray", "org/apache/spark/unsafe/array/LongArray", a)
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
	unique_x := &UnsafeArrayLongArray{}
	unique_x.Callable = dst
	return unique_x
}

// public void freeArray(org.apache.spark.unsafe.array.LongArray)
func (jbobject *MemoryMemoryConsumer) FreeArray(a UnsafeArrayLongArrayInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "freeArray", javabind.Void, conv_a.Value().Cast("org/apache/spark/unsafe/array/LongArray"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


