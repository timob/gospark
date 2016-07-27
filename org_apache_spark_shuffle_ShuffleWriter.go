package gospark

import "github.com/timob/javabind"

type ShuffleShuffleWriterInterface interface {
	JavaLangObjectInterface
}

type ShuffleShuffleWriter struct {
	JavaLangObject
}

// public org.apache.spark.shuffle.ShuffleWriter()
func NewShuffleShuffleWriter() (*ShuffleShuffleWriter) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/shuffle/ShuffleWriter")
	if err != nil {
		panic(err)
	}
	x := &ShuffleShuffleWriter{}
	x.Callable = &javabind.Callable{obj}
	return x
}


