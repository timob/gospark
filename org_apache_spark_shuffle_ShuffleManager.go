package gospark

import "github.com/timob/javabind"

type ShuffleShuffleManagerInterface interface {

	// public abstract <K, V, C> org.apache.spark.shuffle.ShuffleHandle registerShuffle(int, int, org.apache.spark.ShuffleDependency<K, V, C>)
	RegisterShuffle(a int, b int, c ShuffleDependencyInterface) *ShuffleShuffleHandle

	// public abstract <K, V> org.apache.spark.shuffle.ShuffleWriter<K, V> getWriter(org.apache.spark.shuffle.ShuffleHandle, int, org.apache.spark.TaskContext)
	GetWriter(a ShuffleShuffleHandleInterface, b int, c TaskContextInterface) *ShuffleShuffleWriter

	// public abstract <K, C> org.apache.spark.shuffle.ShuffleReader<K, C> getReader(org.apache.spark.shuffle.ShuffleHandle, int, int, org.apache.spark.TaskContext)
	GetReader(a ShuffleShuffleHandleInterface, b int, c int, d TaskContextInterface) *ShuffleShuffleReader

	// public abstract boolean unregisterShuffle(int)
	UnregisterShuffle(a int) bool

	// public abstract org.apache.spark.shuffle.ShuffleBlockResolver shuffleBlockResolver()
	ShuffleBlockResolver() *ShuffleShuffleBlockResolver

	// public abstract void stop()
	Stop() 
}

type ShuffleShuffleManager struct {
	JavaLangObject
}

// public abstract <K, V, C> org.apache.spark.shuffle.ShuffleHandle registerShuffle(int, int, org.apache.spark.ShuffleDependency<K, V, C>)
func (jbobject *ShuffleShuffleManager) RegisterShuffle(a int, b int, c ShuffleDependencyInterface) *ShuffleShuffleHandle {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "registerShuffle", "org/apache/spark/shuffle/ShuffleHandle", a, b, conv_c.Value().Cast("org/apache/spark/ShuffleDependency"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ShuffleShuffleHandle{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract <K, V> org.apache.spark.shuffle.ShuffleWriter<K, V> getWriter(org.apache.spark.shuffle.ShuffleHandle, int, org.apache.spark.TaskContext)
func (jbobject *ShuffleShuffleManager) GetWriter(a ShuffleShuffleHandleInterface, b int, c TaskContextInterface) *ShuffleShuffleWriter {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWriter", "org/apache/spark/shuffle/ShuffleWriter", conv_a.Value().Cast("org/apache/spark/shuffle/ShuffleHandle"), b, conv_c.Value().Cast("org/apache/spark/TaskContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ShuffleShuffleWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract <K, C> org.apache.spark.shuffle.ShuffleReader<K, C> getReader(org.apache.spark.shuffle.ShuffleHandle, int, int, org.apache.spark.TaskContext)
func (jbobject *ShuffleShuffleManager) GetReader(a ShuffleShuffleHandleInterface, b int, c int, d TaskContextInterface) *ShuffleShuffleReader {
	conv_a := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReader", "org/apache/spark/shuffle/ShuffleReader", conv_a.Value().Cast("org/apache/spark/shuffle/ShuffleHandle"), b, c, conv_d.Value().Cast("org/apache/spark/TaskContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_d.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ShuffleShuffleReader{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract boolean unregisterShuffle(int)
func (jbobject *ShuffleShuffleManager) UnregisterShuffle(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unregisterShuffle", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract org.apache.spark.shuffle.ShuffleBlockResolver shuffleBlockResolver()
func (jbobject *ShuffleShuffleManager) ShuffleBlockResolver() *ShuffleShuffleBlockResolver {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleBlockResolver", "org/apache/spark/shuffle/ShuffleBlockResolver")
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
	unique_x := &ShuffleShuffleBlockResolver{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void stop()
func (jbobject *ShuffleShuffleManager) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


