package gospark

import "github.com/timob/javabind"

type UnsafeMemoryMemoryAllocatorInterface interface {

	// public abstract void free(org.apache.spark.unsafe.memory.MemoryBlock)
	Free(a UnsafeMemoryMemoryBlockInterface) 
}

type UnsafeMemoryMemoryAllocator struct {
	JavaLangObject
}

// public abstract org.apache.spark.unsafe.memory.MemoryBlock allocate(long) throws java.lang.OutOfMemoryError
func (jbobject *UnsafeMemoryMemoryAllocator) Allocate(a int64) (*UnsafeMemoryMemoryBlock, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "allocate", "org/apache/spark/unsafe/memory/MemoryBlock", a)
	if err != nil {
		var zero *UnsafeMemoryMemoryBlock
		return zero, err
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
	return unique_x, nil
}

// public abstract void free(org.apache.spark.unsafe.memory.MemoryBlock)
func (jbobject *UnsafeMemoryMemoryAllocator) Free(a UnsafeMemoryMemoryBlockInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "free", javabind.Void, conv_a.Value().Cast("org/apache/spark/unsafe/memory/MemoryBlock"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

func (jbobject *UnsafeMemoryMemoryAllocator) UNSAFE() *UnsafeMemoryMemoryAllocator {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/unsafe/memory/MemoryAllocator", "UNSAFE", "org/apache/spark/unsafe/memory/MemoryAllocator")
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
	unique_x := &UnsafeMemoryMemoryAllocator{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *UnsafeMemoryMemoryAllocator) SetFieldUNSAFE(val UnsafeMemoryMemoryAllocatorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/unsafe/memory/MemoryAllocator", "UNSAFE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *UnsafeMemoryMemoryAllocator) HEAP() *UnsafeMemoryMemoryAllocator {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/unsafe/memory/MemoryAllocator", "HEAP", "org/apache/spark/unsafe/memory/MemoryAllocator")
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
	unique_x := &UnsafeMemoryMemoryAllocator{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *UnsafeMemoryMemoryAllocator) SetFieldHEAP(val UnsafeMemoryMemoryAllocatorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/unsafe/memory/MemoryAllocator", "HEAP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


