package gospark

import "github.com/timob/javabind"

type SerializerSerializerInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.serializer.Serializer getSerializer(org.apache.spark.serializer.Serializer)
	GetSerializer(a SerializerSerializerInterface) *SerializerSerializer

	// public org.apache.spark.serializer.Serializer setDefaultClassLoader(java.lang.ClassLoader)
	SetDefaultClassLoader(a JavaLangClassLoaderInterface) *SerializerSerializer

	// public abstract org.apache.spark.serializer.SerializerInstance newInstance()
	NewInstance() *SerializerSerializerInstance

	// public boolean supportsRelocationOfSerializedObjects()
	SupportsRelocationOfSerializedObjects() bool
}

type SerializerSerializer struct {
	JavaLangObject
}

// public org.apache.spark.serializer.Serializer()
func NewSerializerSerializer() (*SerializerSerializer) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/serializer/Serializer")
	if err != nil {
		panic(err)
	}
	x := &SerializerSerializer{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.serializer.Serializer getSerializer(org.apache.spark.serializer.Serializer)
func (jbobject *SerializerSerializer) GetSerializer(a SerializerSerializerInterface) *SerializerSerializer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/serializer/Serializer", "getSerializer", "org/apache/spark/serializer/Serializer", conv_a.Value().Cast("org/apache/spark/serializer/Serializer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SerializerSerializer{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.serializer.Serializer setDefaultClassLoader(java.lang.ClassLoader)
func (jbobject *SerializerSerializer) SetDefaultClassLoader(a JavaLangClassLoaderInterface) *SerializerSerializer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setDefaultClassLoader", "org/apache/spark/serializer/Serializer", conv_a.Value().Cast("java/lang/ClassLoader"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SerializerSerializer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.serializer.SerializerInstance newInstance()
func (jbobject *SerializerSerializer) NewInstance() *SerializerSerializerInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "newInstance", "org/apache/spark/serializer/SerializerInstance")
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
	unique_x := &SerializerSerializerInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean supportsRelocationOfSerializedObjects()
func (jbobject *SerializerSerializer) SupportsRelocationOfSerializedObjects() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "supportsRelocationOfSerializedObjects", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


