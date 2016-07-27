package gospark

import "github.com/timob/javabind"

type SerializerSerializerInstanceInterface interface {
	JavaLangObjectInterface
}

type SerializerSerializerInstance struct {
	JavaLangObject
}

// public org.apache.spark.serializer.SerializerInstance()
func NewSerializerSerializerInstance() (*SerializerSerializerInstance) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/serializer/SerializerInstance")
	if err != nil {
		panic(err)
	}
	x := &SerializerSerializerInstance{}
	x.Callable = &javabind.Callable{obj}
	return x
}


