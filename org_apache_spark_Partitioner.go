package gospark

import "github.com/timob/javabind"

type PartitionerInterface interface {
	JavaLangObjectInterface

	// public abstract int numPartitions()
	NumPartitions() int

	// public abstract int getPartition(java.lang.Object)
	GetPartition(a interface{}) int
}

type Partitioner struct {
	JavaLangObject
}

// public org.apache.spark.Partitioner()
func NewPartitioner() (*Partitioner) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/Partitioner")
	if err != nil {
		panic(err)
	}
	x := &Partitioner{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract int numPartitions()
func (jbobject *Partitioner) NumPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int getPartition(java.lang.Object)
func (jbobject *Partitioner) GetPartition(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPartition", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}


