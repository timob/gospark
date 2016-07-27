package gospark

import "github.com/timob/javabind"

type MemoryTaskMemoryManagerInterface interface {
	JavaLangObjectInterface

	// public long acquireExecutionMemory(long, org.apache.spark.memory.MemoryMode, org.apache.spark.memory.MemoryConsumer)
	AcquireExecutionMemory(a int64, b MemoryMemoryModeInterface, c MemoryMemoryConsumerInterface) int64

	// public void releaseExecutionMemory(long, org.apache.spark.memory.MemoryMode, org.apache.spark.memory.MemoryConsumer)
	ReleaseExecutionMemory(a int64, b MemoryMemoryModeInterface, c MemoryMemoryConsumerInterface) 

	// public void showMemoryUsage()
	ShowMemoryUsage() 

	// public long pageSizeBytes()
	PageSizeBytes() int64

	// public org.apache.spark.unsafe.memory.MemoryBlock allocatePage(long, org.apache.spark.memory.MemoryConsumer)
	AllocatePage(a int64, b MemoryMemoryConsumerInterface) *UnsafeMemoryMemoryBlock

	// public void freePage(org.apache.spark.unsafe.memory.MemoryBlock, org.apache.spark.memory.MemoryConsumer)
	FreePage(a UnsafeMemoryMemoryBlockInterface, b MemoryMemoryConsumerInterface) 

	// public long encodePageNumberAndOffset(org.apache.spark.unsafe.memory.MemoryBlock, long)
	EncodePageNumberAndOffset2(a UnsafeMemoryMemoryBlockInterface, b int64) int64

	// public static long encodePageNumberAndOffset(int, long)
	EncodePageNumberAndOffset(a int, b int64) int64

	// public static int decodePageNumber(long)
	DecodePageNumber(a int64) int

	// public java.lang.Object getPage(long)
	GetPage(a int64) *JavaLangObject

	// public long getOffsetInPage(long)
	GetOffsetInPage(a int64) int64

	// public long cleanUpAllAllocatedMemory()
	CleanUpAllAllocatedMemory() int64

	// public long getMemoryConsumptionForThisTask()
	GetMemoryConsumptionForThisTask() int64
}

type MemoryTaskMemoryManager struct {
	JavaLangObject
}

// public org.apache.spark.memory.TaskMemoryManager(org.apache.spark.memory.MemoryManager, long)
func NewMemoryTaskMemoryManager(a MemoryMemoryManagerInterface, b int64) (*MemoryTaskMemoryManager) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/memory/TaskMemoryManager", conv_a.Value().Cast("org/apache/spark/memory/MemoryManager"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &MemoryTaskMemoryManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long acquireExecutionMemory(long, org.apache.spark.memory.MemoryMode, org.apache.spark.memory.MemoryConsumer)
func (jbobject *MemoryTaskMemoryManager) AcquireExecutionMemory(a int64, b MemoryMemoryModeInterface, c MemoryMemoryConsumerInterface) int64 {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "acquireExecutionMemory", javabind.Long, a, conv_b.Value().Cast("org/apache/spark/memory/MemoryMode"), conv_c.Value().Cast("org/apache/spark/memory/MemoryConsumer"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	return jret.(int64)
}

// public void releaseExecutionMemory(long, org.apache.spark.memory.MemoryMode, org.apache.spark.memory.MemoryConsumer)
func (jbobject *MemoryTaskMemoryManager) ReleaseExecutionMemory(a int64, b MemoryMemoryModeInterface, c MemoryMemoryConsumerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseExecutionMemory", javabind.Void, a, conv_b.Value().Cast("org/apache/spark/memory/MemoryMode"), conv_c.Value().Cast("org/apache/spark/memory/MemoryConsumer"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()

}

// public void showMemoryUsage()
func (jbobject *MemoryTaskMemoryManager) ShowMemoryUsage()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showMemoryUsage", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public long pageSizeBytes()
func (jbobject *MemoryTaskMemoryManager) PageSizeBytes() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "pageSizeBytes", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.unsafe.memory.MemoryBlock allocatePage(long, org.apache.spark.memory.MemoryConsumer)
func (jbobject *MemoryTaskMemoryManager) AllocatePage(a int64, b MemoryMemoryConsumerInterface) *UnsafeMemoryMemoryBlock {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "allocatePage", "org/apache/spark/unsafe/memory/MemoryBlock", a, conv_b.Value().Cast("org/apache/spark/memory/MemoryConsumer"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
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

// public void freePage(org.apache.spark.unsafe.memory.MemoryBlock, org.apache.spark.memory.MemoryConsumer)
func (jbobject *MemoryTaskMemoryManager) FreePage(a UnsafeMemoryMemoryBlockInterface, b MemoryMemoryConsumerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "freePage", javabind.Void, conv_a.Value().Cast("org/apache/spark/unsafe/memory/MemoryBlock"), conv_b.Value().Cast("org/apache/spark/memory/MemoryConsumer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public long encodePageNumberAndOffset(org.apache.spark.unsafe.memory.MemoryBlock, long)
func (jbobject *MemoryTaskMemoryManager) EncodePageNumberAndOffset2(a UnsafeMemoryMemoryBlockInterface, b int64) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "encodePageNumberAndOffset", javabind.Long, conv_a.Value().Cast("org/apache/spark/unsafe/memory/MemoryBlock"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public static long encodePageNumberAndOffset(int, long)
func (jbobject *MemoryTaskMemoryManager) EncodePageNumberAndOffset(a int, b int64) int64 {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/memory/TaskMemoryManager", "encodePageNumberAndOffset", javabind.Long, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public static int decodePageNumber(long)
func (jbobject *MemoryTaskMemoryManager) DecodePageNumber(a int64) int {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/memory/TaskMemoryManager", "decodePageNumber", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object getPage(long)
func (jbobject *MemoryTaskMemoryManager) GetPage(a int64) *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPage", "java/lang/Object", a)
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

// public long getOffsetInPage(long)
func (jbobject *MemoryTaskMemoryManager) GetOffsetInPage(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffsetInPage", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long cleanUpAllAllocatedMemory()
func (jbobject *MemoryTaskMemoryManager) CleanUpAllAllocatedMemory() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cleanUpAllAllocatedMemory", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getMemoryConsumptionForThisTask()
func (jbobject *MemoryTaskMemoryManager) GetMemoryConsumptionForThisTask() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMemoryConsumptionForThisTask", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *MemoryTaskMemoryManager) MAXIMUM_PAGE_SIZE_BYTES() int64 {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/memory/TaskMemoryManager", "MAXIMUM_PAGE_SIZE_BYTES", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *MemoryTaskMemoryManager) SetFieldMAXIMUM_PAGE_SIZE_BYTES(val int64) {
	err := javabind.GetEnv().SetStaticField("org/apache/spark/memory/TaskMemoryManager", "MAXIMUM_PAGE_SIZE_BYTES", val)
	if err != nil {
		panic(err)
	}

}


