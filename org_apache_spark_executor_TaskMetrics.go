package gospark

import "github.com/timob/javabind"

type ExecutorTaskMetricsInterface interface {
	JavaLangObjectInterface

	// public static java.lang.String getCachedHostName(java.lang.String)
	GetCachedHostName(a string) string

	// public static org.apache.spark.executor.TaskMetrics empty()
	Empty() *ExecutorTaskMetrics

	// public java.lang.String hostname()
	Hostname() string

	// public void setHostname(java.lang.String)
	SetHostname(a string) 

	// public long executorDeserializeTime()
	ExecutorDeserializeTime() int64

	// public void setExecutorDeserializeTime(long)
	SetExecutorDeserializeTime(a int64) 

	// public long executorRunTime()
	ExecutorRunTime() int64

	// public void setExecutorRunTime(long)
	SetExecutorRunTime(a int64) 

	// public long resultSize()
	ResultSize() int64

	// public void setResultSize(long)
	SetResultSize(a int64) 

	// public long jvmGCTime()
	JvmGCTime() int64

	// public void setJvmGCTime(long)
	SetJvmGCTime(a int64) 

	// public long resultSerializationTime()
	ResultSerializationTime() int64

	// public void setResultSerializationTime(long)
	SetResultSerializationTime(a int64) 

	// public long memoryBytesSpilled()
	MemoryBytesSpilled() int64

	// public void incMemoryBytesSpilled(long)
	IncMemoryBytesSpilled(a int64) 

	// public void decMemoryBytesSpilled(long)
	DecMemoryBytesSpilled(a int64) 

	// public long diskBytesSpilled()
	DiskBytesSpilled() int64

	// public void incDiskBytesSpilled(long)
	IncDiskBytesSpilled(a int64) 

	// public void decDiskBytesSpilled(long)
	DecDiskBytesSpilled(a int64) 

	// public synchronized org.apache.spark.executor.ShuffleReadMetrics createShuffleReadMetricsForDependency()
	CreateShuffleReadMetricsForDependency() *ExecutorShuffleReadMetrics

	// public synchronized void updateShuffleReadMetrics()
	UpdateShuffleReadMetrics() 

	// public synchronized void updateInputMetrics()
	UpdateInputMetrics() 

	// public synchronized void updateAccumulators()
	UpdateAccumulators() 
}

type ExecutorTaskMetrics struct {
	JavaLangObject
}

// public org.apache.spark.executor.TaskMetrics()
func NewExecutorTaskMetrics() (*ExecutorTaskMetrics) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/executor/TaskMetrics")
	if err != nil {
		panic(err)
	}
	x := &ExecutorTaskMetrics{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static java.lang.String getCachedHostName(java.lang.String)
func (jbobject *ExecutorTaskMetrics) GetCachedHostName(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/executor/TaskMetrics", "getCachedHostName", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static org.apache.spark.executor.TaskMetrics empty()
func (jbobject *ExecutorTaskMetrics) Empty() *ExecutorTaskMetrics {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/executor/TaskMetrics", "empty", "org/apache/spark/executor/TaskMetrics")
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
	unique_x := &ExecutorTaskMetrics{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String hostname()
func (jbobject *ExecutorTaskMetrics) Hostname() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hostname", "java/lang/String")
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

// public void setHostname(java.lang.String)
func (jbobject *ExecutorTaskMetrics) SetHostname(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostname", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public long executorDeserializeTime()
func (jbobject *ExecutorTaskMetrics) ExecutorDeserializeTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorDeserializeTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setExecutorDeserializeTime(long)
func (jbobject *ExecutorTaskMetrics) SetExecutorDeserializeTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExecutorDeserializeTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long executorRunTime()
func (jbobject *ExecutorTaskMetrics) ExecutorRunTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorRunTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setExecutorRunTime(long)
func (jbobject *ExecutorTaskMetrics) SetExecutorRunTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExecutorRunTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long resultSize()
func (jbobject *ExecutorTaskMetrics) ResultSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "resultSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setResultSize(long)
func (jbobject *ExecutorTaskMetrics) SetResultSize(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResultSize", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long jvmGCTime()
func (jbobject *ExecutorTaskMetrics) JvmGCTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jvmGCTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setJvmGCTime(long)
func (jbobject *ExecutorTaskMetrics) SetJvmGCTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJvmGCTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long resultSerializationTime()
func (jbobject *ExecutorTaskMetrics) ResultSerializationTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "resultSerializationTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void setResultSerializationTime(long)
func (jbobject *ExecutorTaskMetrics) SetResultSerializationTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResultSerializationTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long memoryBytesSpilled()
func (jbobject *ExecutorTaskMetrics) MemoryBytesSpilled() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryBytesSpilled", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void incMemoryBytesSpilled(long)
func (jbobject *ExecutorTaskMetrics) IncMemoryBytesSpilled(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incMemoryBytesSpilled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decMemoryBytesSpilled(long)
func (jbobject *ExecutorTaskMetrics) DecMemoryBytesSpilled(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decMemoryBytesSpilled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long diskBytesSpilled()
func (jbobject *ExecutorTaskMetrics) DiskBytesSpilled() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskBytesSpilled", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void incDiskBytesSpilled(long)
func (jbobject *ExecutorTaskMetrics) IncDiskBytesSpilled(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incDiskBytesSpilled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decDiskBytesSpilled(long)
func (jbobject *ExecutorTaskMetrics) DecDiskBytesSpilled(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decDiskBytesSpilled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public synchronized org.apache.spark.executor.ShuffleReadMetrics createShuffleReadMetricsForDependency()
func (jbobject *ExecutorTaskMetrics) CreateShuffleReadMetricsForDependency() *ExecutorShuffleReadMetrics {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createShuffleReadMetricsForDependency", "org/apache/spark/executor/ShuffleReadMetrics")
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
	unique_x := &ExecutorShuffleReadMetrics{}
	unique_x.Callable = dst
	return unique_x
}

// public synchronized void updateShuffleReadMetrics()
func (jbobject *ExecutorTaskMetrics) UpdateShuffleReadMetrics()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "updateShuffleReadMetrics", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public synchronized void updateInputMetrics()
func (jbobject *ExecutorTaskMetrics) UpdateInputMetrics()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "updateInputMetrics", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public synchronized void updateAccumulators()
func (jbobject *ExecutorTaskMetrics) UpdateAccumulators()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "updateAccumulators", javabind.Void)
	if err != nil {
		panic(err)
	}

}


