package gospark

import "github.com/timob/javabind"

type MemoryExecutionMemoryPoolInterface interface {
	MemoryMemoryPoolInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public long getMemoryUsageForTask(long)
	GetMemoryUsageForTask(a int64) int64

	// public void releaseMemory(long, long)
	ReleaseMemory(a int64, b int64) 

	// public long releaseAllMemoryForTask(long)
	ReleaseAllMemoryForTask(a int64) int64
}

type MemoryExecutionMemoryPool struct {
	MemoryMemoryPool
}

// public org.apache.spark.memory.ExecutionMemoryPool(java.lang.Object, java.lang.String)
func NewMemoryExecutionMemoryPool(a interface{}, b string) (*MemoryExecutionMemoryPool) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/memory/ExecutionMemoryPool", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &MemoryExecutionMemoryPool{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *MemoryExecutionMemoryPool) LogName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "logName", "java/lang/String")
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

// public boolean isTraceEnabled()
func (jbobject *MemoryExecutionMemoryPool) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long memoryUsed()
func (jbobject *MemoryExecutionMemoryPool) MemoryUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getMemoryUsageForTask(long)
func (jbobject *MemoryExecutionMemoryPool) GetMemoryUsageForTask(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMemoryUsageForTask", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void releaseMemory(long, long)
func (jbobject *MemoryExecutionMemoryPool) ReleaseMemory(a int64, b int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseMemory", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public long releaseAllMemoryForTask(long)
func (jbobject *MemoryExecutionMemoryPool) ReleaseAllMemoryForTask(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "releaseAllMemoryForTask", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


