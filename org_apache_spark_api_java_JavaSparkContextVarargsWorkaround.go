package gospark

import "github.com/timob/javabind"

type ApiJavaJavaSparkContextVarargsWorkaroundInterface interface {
	JavaLangObjectInterface

	// public final <T> org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>...)
	Union3(a ...*ApiJavaJavaRDD) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD...)
	Union(a ...*ApiJavaJavaDoubleRDD) *ApiJavaJavaDoubleRDD

	// public final <K, V> org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>...)
	Union2(a ...*ApiJavaJavaPairRDD) *ApiJavaJavaPairRDD

	// public abstract <T> org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>, java.util.List<org.apache.spark.api.java.JavaRDD<T>>)
	Union6(a ApiJavaJavaRDDInterface, b []*ApiJavaJavaRDD) *ApiJavaJavaRDD

	// public abstract org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD, java.util.List<org.apache.spark.api.java.JavaDoubleRDD>)
	Union4(a ApiJavaJavaDoubleRDDInterface, b []*ApiJavaJavaDoubleRDD) *ApiJavaJavaDoubleRDD

	// public abstract <K, V> org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>, java.util.List<org.apache.spark.api.java.JavaPairRDD<K, V>>)
	Union5(a ApiJavaJavaPairRDDInterface, b []*ApiJavaJavaPairRDD) *ApiJavaJavaPairRDD
}

type ApiJavaJavaSparkContextVarargsWorkaround struct {
	JavaLangObject
}

// public final <T> org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>...)
func (jbobject *ApiJavaJavaSparkContextVarargsWorkaround) Union3(a ...*ApiJavaJavaRDD) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/apache/spark/api/java/JavaRDD<T>")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"))
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD...)
func (jbobject *ApiJavaJavaSparkContextVarargsWorkaround) Union(a ...*ApiJavaJavaDoubleRDD) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/apache/spark/api/java/JavaDoubleRDD")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"))
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public final <K, V> org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>...)
func (jbobject *ApiJavaJavaSparkContextVarargsWorkaround) Union2(a ...*ApiJavaJavaPairRDD) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/apache/spark/api/java/JavaPairRDD<K, V>")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"))
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract <T> org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>, java.util.List<org.apache.spark.api.java.JavaRDD<T>>)
func (jbobject *ApiJavaJavaSparkContextVarargsWorkaround) Union6(a ApiJavaJavaRDDInterface, b []*ApiJavaJavaRDD) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD, java.util.List<org.apache.spark.api.java.JavaDoubleRDD>)
func (jbobject *ApiJavaJavaSparkContextVarargsWorkaround) Union4(a ApiJavaJavaDoubleRDDInterface, b []*ApiJavaJavaDoubleRDD) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract <K, V> org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>, java.util.List<org.apache.spark.api.java.JavaPairRDD<K, V>>)
func (jbobject *ApiJavaJavaSparkContextVarargsWorkaround) Union5(a ApiJavaJavaPairRDDInterface, b []*ApiJavaJavaPairRDD) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}


