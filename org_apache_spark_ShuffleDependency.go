package gospark

import "github.com/timob/javabind"

type ShuffleDependencyInterface interface {
	DependencyInterface

	// public org.apache.spark.Partitioner partitioner()
	Partitioner() *Partitioner

	// public boolean mapSideCombine()
	MapSideCombine() bool

	// public java.lang.String keyClassName()
	KeyClassName() string

	// public java.lang.String valueClassName()
	ValueClassName() string

	// public int shuffleId()
	ShuffleId() int

	// public org.apache.spark.shuffle.ShuffleHandle shuffleHandle()
	ShuffleHandle() *ShuffleShuffleHandle
}

type ShuffleDependency struct {
	Dependency
}

// public org.apache.spark.Partitioner partitioner()
func (jbobject *ShuffleDependency) Partitioner() *Partitioner {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitioner", "org/apache/spark/Partitioner")
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
	unique_x := &Partitioner{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean mapSideCombine()
func (jbobject *ShuffleDependency) MapSideCombine() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapSideCombine", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String keyClassName()
func (jbobject *ShuffleDependency) KeyClassName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "keyClassName", "java/lang/String")
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

// public java.lang.String valueClassName()
func (jbobject *ShuffleDependency) ValueClassName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "valueClassName", "java/lang/String")
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

// public int shuffleId()
func (jbobject *ShuffleDependency) ShuffleId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.shuffle.ShuffleHandle shuffleHandle()
func (jbobject *ShuffleDependency) ShuffleHandle() *ShuffleShuffleHandle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleHandle", "org/apache/spark/shuffle/ShuffleHandle")
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
	unique_x := &ShuffleShuffleHandle{}
	unique_x.Callable = dst
	return unique_x
}


