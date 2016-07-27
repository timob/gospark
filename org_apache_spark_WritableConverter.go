package gospark

import "github.com/timob/javabind"

type WritableConverterInterface interface {
	JavaLangObjectInterface

	// public static <T extends org.apache.hadoop.io.Writable> org.apache.spark.WritableConverter<T> writableWritableConverter()
	WritableWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.String> stringWritableConverter()
	StringWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<byte[]> bytesWritableConverter()
	BytesWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> booleanWritableConverter()
	BooleanWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> floatWritableConverter()
	FloatWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> doubleWritableConverter()
	DoubleWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> longWritableConverter()
	LongWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> intWritableConverter()
	IntWritableConverter() *WritableConverter
}

type WritableConverter struct {
	JavaLangObject
}

// public static <T extends org.apache.hadoop.io.Writable> org.apache.spark.WritableConverter<T> writableWritableConverter()
func (jbobject *WritableConverter) WritableWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "writableWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.String> stringWritableConverter()
func (jbobject *WritableConverter) StringWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "stringWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<byte[]> bytesWritableConverter()
func (jbobject *WritableConverter) BytesWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "bytesWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> booleanWritableConverter()
func (jbobject *WritableConverter) BooleanWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "booleanWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> floatWritableConverter()
func (jbobject *WritableConverter) FloatWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "floatWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> doubleWritableConverter()
func (jbobject *WritableConverter) DoubleWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "doubleWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> longWritableConverter()
func (jbobject *WritableConverter) LongWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "longWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> intWritableConverter()
func (jbobject *WritableConverter) IntWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/WritableConverter", "intWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}


