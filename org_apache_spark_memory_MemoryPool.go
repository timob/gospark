package gospark

import "github.com/timob/javabind"

type MemoryMemoryPoolInterface interface {
	JavaLangObjectInterface

	// public final long poolSize()
	PoolSize() int64

	// public final long memoryFree()
	MemoryFree() int64

	// public final void incrementPoolSize(long)
	IncrementPoolSize(a int64) 

	// public final void decrementPoolSize(long)
	DecrementPoolSize(a int64) 

	// public abstract long memoryUsed()
	MemoryUsed() int64
}

type MemoryMemoryPool struct {
	JavaLangObject
}

// public org.apache.spark.memory.MemoryPool(java.lang.Object)
func NewMemoryMemoryPool(a interface{}) (*MemoryMemoryPool) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/memory/MemoryPool", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &MemoryMemoryPool{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final long poolSize()
func (jbobject *MemoryMemoryPool) PoolSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "poolSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final long memoryFree()
func (jbobject *MemoryMemoryPool) MemoryFree() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryFree", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final void incrementPoolSize(long)
func (jbobject *MemoryMemoryPool) IncrementPoolSize(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incrementPoolSize", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final void decrementPoolSize(long)
func (jbobject *MemoryMemoryPool) DecrementPoolSize(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decrementPoolSize", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public abstract long memoryUsed()
func (jbobject *MemoryMemoryPool) MemoryUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


