package gospark

import "github.com/timob/javabind"

type MapOutputStatisticsInterface interface {
	JavaLangObjectInterface

	// public int shuffleId()
	ShuffleId() int

	// public long[] bytesByPartitionId()
	BytesByPartitionId() []int64
}

type MapOutputStatistics struct {
	JavaLangObject
}

// public org.apache.spark.MapOutputStatistics(int, long[])
func NewMapOutputStatistics(a int, b []int64) (*MapOutputStatistics) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/MapOutputStatistics", a, b)
	if err != nil {
		panic(err)
	}
	x := &MapOutputStatistics{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int shuffleId()
func (jbobject *MapOutputStatistics) ShuffleId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long[] bytesByPartitionId()
func (jbobject *MapOutputStatistics) BytesByPartitionId() []int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "bytesByPartitionId", javabind.Long | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int64)
}


