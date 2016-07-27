package gospark

import "github.com/timob/javabind"

type ShuffleShuffleHandleInterface interface {
	JavaLangObjectInterface

	// public int shuffleId()
	ShuffleId() int
}

type ShuffleShuffleHandle struct {
	JavaLangObject
}

// public org.apache.spark.shuffle.ShuffleHandle(int)
func NewShuffleShuffleHandle(a int) (*ShuffleShuffleHandle) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/shuffle/ShuffleHandle", a)
	if err != nil {
		panic(err)
	}
	x := &ShuffleShuffleHandle{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int shuffleId()
func (jbobject *ShuffleShuffleHandle) ShuffleId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


