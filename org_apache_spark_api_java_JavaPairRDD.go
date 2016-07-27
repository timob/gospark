package gospark

import "github.com/timob/javabind"

type ApiJavaJavaPairRDDInterface interface {
	ApiJavaAbstractJavaRDDLikeInterface

	// public org.apache.spark.api.java.JavaPairRDD<K, V> cache()
	Cache() *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> persist(org.apache.spark.storage.StorageLevel)
	Persist(a StorageStorageLevelInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> unpersist()
	Unpersist() *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> unpersist(boolean)
	Unpersist2(a bool) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> distinct()
	Distinct() *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> distinct(int)
	Distinct2(a int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> coalesce(int)
	Coalesce(a int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> coalesce(int, boolean)
	Coalesce2(a int, b bool) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> repartition(int)
	Repartition(a int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sample(boolean, double)
	Sample(a bool, b float64) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sample(boolean, double, long)
	Sample2(a bool, b float64, c int64) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKey(boolean, java.util.Map<K, java.lang.Object>, long)
	SampleByKey2(a bool, b map[*JavaLangObject]*JavaLangObject, c int64) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKey(boolean, java.util.Map<K, java.lang.Object>)
	SampleByKey(a bool, b map[*JavaLangObject]*JavaLangObject) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKeyExact(boolean, java.util.Map<K, java.lang.Object>, long)
	SampleByKeyExact2(a bool, b map[*JavaLangObject]*JavaLangObject, c int64) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKeyExact(boolean, java.util.Map<K, java.lang.Object>)
	SampleByKeyExact(a bool, b map[*JavaLangObject]*JavaLangObject) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>)
	Union(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> intersection(org.apache.spark.api.java.JavaPairRDD<K, V>)
	Intersection(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD

	// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>, org.apache.spark.Partitioner, boolean, org.apache.spark.serializer.Serializer)
	CombineByKey4(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d PartitionerInterface, e bool, f SerializerSerializerInterface) *ApiJavaJavaPairRDD

	// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>, org.apache.spark.Partitioner)
	CombineByKey3(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d PartitionerInterface) *ApiJavaJavaPairRDD

	// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>, int)
	CombineByKey2(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> reduceByKey(org.apache.spark.Partitioner, org.apache.spark.api.java.function.Function2<V, V, V>)
	ReduceByKey3(a PartitionerInterface, b ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public java.util.Map<K, V> reduceByKeyLocally(org.apache.spark.api.java.function.Function2<V, V, V>)
	ReduceByKeyLocally(a ApiJavaFunctionFunction2Interface) map[*JavaLangObject]*JavaLangObject

	// public java.util.Map<K, java.lang.Object> countByKey()
	CountByKey() map[*JavaLangObject]*JavaLangObject

	// public org.apache.spark.partial.PartialResult<java.util.Map<K, org.apache.spark.partial.BoundedDouble>> countByKeyApprox(long)
	CountByKeyApprox(a int64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<java.util.Map<K, org.apache.spark.partial.BoundedDouble>> countByKeyApprox(long, double)
	CountByKeyApprox2(a int64, b float64) *PartialPartialResult

	// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> aggregateByKey(U, org.apache.spark.Partitioner, org.apache.spark.api.java.function.Function2<U, V, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
	AggregateByKey3(a interface{}, b PartitionerInterface, c ApiJavaFunctionFunction2Interface, d ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> aggregateByKey(U, int, org.apache.spark.api.java.function.Function2<U, V, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
	AggregateByKey2(a interface{}, b int, c ApiJavaFunctionFunction2Interface, d ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> aggregateByKey(U, org.apache.spark.api.java.function.Function2<U, V, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
	AggregateByKey(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> foldByKey(V, org.apache.spark.Partitioner, org.apache.spark.api.java.function.Function2<V, V, V>)
	FoldByKey3(a JavaLangObjectInterface, b PartitionerInterface, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> foldByKey(V, int, org.apache.spark.api.java.function.Function2<V, V, V>)
	FoldByKey2(a JavaLangObjectInterface, b int, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> foldByKey(V, org.apache.spark.api.java.function.Function2<V, V, V>)
	FoldByKey(a JavaLangObjectInterface, b ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> reduceByKey(org.apache.spark.api.java.function.Function2<V, V, V>, int)
	ReduceByKey2(a ApiJavaFunctionFunction2Interface, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Iterable<V>> groupByKey(org.apache.spark.Partitioner)
	GroupByKey3(a PartitionerInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Iterable<V>> groupByKey(int)
	GroupByKey2(a int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> subtract(org.apache.spark.api.java.JavaPairRDD<K, V>)
	Subtract(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> subtract(org.apache.spark.api.java.JavaPairRDD<K, V>, int)
	Subtract2(a ApiJavaJavaPairRDDInterface, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> subtract(org.apache.spark.api.java.JavaPairRDD<K, V>, org.apache.spark.Partitioner)
	Subtract3(a ApiJavaJavaPairRDDInterface, b PartitionerInterface) *ApiJavaJavaPairRDD

	// public <W> org.apache.spark.api.java.JavaPairRDD<K, V> subtractByKey(org.apache.spark.api.java.JavaPairRDD<K, W>)
	SubtractByKey(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD

	// public <W> org.apache.spark.api.java.JavaPairRDD<K, V> subtractByKey(org.apache.spark.api.java.JavaPairRDD<K, W>, int)
	SubtractByKey2(a ApiJavaJavaPairRDDInterface, b int) *ApiJavaJavaPairRDD

	// public <W> org.apache.spark.api.java.JavaPairRDD<K, V> subtractByKey(org.apache.spark.api.java.JavaPairRDD<K, W>, org.apache.spark.Partitioner)
	SubtractByKey3(a ApiJavaJavaPairRDDInterface, b PartitionerInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> partitionBy(org.apache.spark.Partitioner)
	PartitionBy(a PartitionerInterface) *ApiJavaJavaPairRDD

	// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>)
	CombineByKey(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> reduceByKey(org.apache.spark.api.java.function.Function2<V, V, V>)
	ReduceByKey(a ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Iterable<V>> groupByKey()
	GroupByKey() *ApiJavaJavaPairRDD

	// public java.util.Map<K, V> collectAsMap()
	CollectAsMap() map[*JavaLangObject]*JavaLangObject

	// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> mapValues(org.apache.spark.api.java.function.Function<V, U>)
	MapValues(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> flatMapValues(org.apache.spark.api.java.function.Function<V, java.lang.Iterable<U>>)
	FlatMapValues(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD

	// public java.util.List<V> lookup(K)
	Lookup(a JavaLangObjectInterface) []*JavaLangObject

	// public org.apache.spark.api.java.JavaPairRDD<K, V> repartitionAndSortWithinPartitions(org.apache.spark.Partitioner)
	RepartitionAndSortWithinPartitions(a PartitionerInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sortByKey()
	SortByKey() *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sortByKey(boolean)
	SortByKey2(a bool) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> sortByKey(boolean, int)
	SortByKey3(a bool, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaRDD<K> keys()
	Keys() *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<V> values()
	Values() *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Object> countApproxDistinctByKey(double, org.apache.spark.Partitioner)
	CountApproxDistinctByKey3(a float64, b PartitionerInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Object> countApproxDistinctByKey(double, int)
	CountApproxDistinctByKey2(a float64, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Object> countApproxDistinctByKey(double)
	CountApproxDistinctByKey(a float64) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<K, V> setName(java.lang.String)
	SetName(a string) *ApiJavaJavaPairRDD

	// public java.lang.Object first()
	First2() *JavaLangObject

	// public org.apache.spark.api.java.JavaRDDLike wrapRDD(org.apache.spark.rdd.RDD)
	WrapRDD(a RddRDDInterface) *ApiJavaJavaRDDLike
}

type ApiJavaJavaPairRDD struct {
	ApiJavaAbstractJavaRDDLike
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> cache()
func (jbobject *ApiJavaJavaPairRDD) Cache() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cache", "org/apache/spark/api/java/JavaPairRDD")
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> persist(org.apache.spark.storage.StorageLevel)
func (jbobject *ApiJavaJavaPairRDD) Persist(a StorageStorageLevelInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "persist", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> unpersist()
func (jbobject *ApiJavaJavaPairRDD) Unpersist() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/api/java/JavaPairRDD")
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> unpersist(boolean)
func (jbobject *ApiJavaJavaPairRDD) Unpersist2(a bool) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> distinct()
func (jbobject *ApiJavaJavaPairRDD) Distinct() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/api/java/JavaPairRDD")
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> distinct(int)
func (jbobject *ApiJavaJavaPairRDD) Distinct2(a int) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> coalesce(int)
func (jbobject *ApiJavaJavaPairRDD) Coalesce(a int) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "coalesce", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> coalesce(int, boolean)
func (jbobject *ApiJavaJavaPairRDD) Coalesce2(a int, b bool) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "coalesce", "org/apache/spark/api/java/JavaPairRDD", a, b)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> repartition(int)
func (jbobject *ApiJavaJavaPairRDD) Repartition(a int) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "repartition", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> sample(boolean, double)
func (jbobject *ApiJavaJavaPairRDD) Sample(a bool, b float64) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/api/java/JavaPairRDD", a, b)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> sample(boolean, double, long)
func (jbobject *ApiJavaJavaPairRDD) Sample2(a bool, b float64, c int64) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/api/java/JavaPairRDD", a, b, c)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKey(boolean, java.util.Map<K, java.lang.Object>, long)
func (jbobject *ApiJavaJavaPairRDD) SampleByKey2(a bool, b map[*JavaLangObject]*JavaLangObject, c int64) *ApiJavaJavaPairRDD {
	conv_b := javabind.NewGoToJavaMap(javabind.NewGoToJavaCallable(), javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleByKey", "org/apache/spark/api/java/JavaPairRDD", a, conv_b.Value().Cast("java/util/Map"), c)
	if err != nil {
		panic(err)
	}
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKey(boolean, java.util.Map<K, java.lang.Object>)
func (jbobject *ApiJavaJavaPairRDD) SampleByKey(a bool, b map[*JavaLangObject]*JavaLangObject) *ApiJavaJavaPairRDD {
	conv_b := javabind.NewGoToJavaMap(javabind.NewGoToJavaCallable(), javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleByKey", "org/apache/spark/api/java/JavaPairRDD", a, conv_b.Value().Cast("java/util/Map"))
	if err != nil {
		panic(err)
	}
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKeyExact(boolean, java.util.Map<K, java.lang.Object>, long)
func (jbobject *ApiJavaJavaPairRDD) SampleByKeyExact2(a bool, b map[*JavaLangObject]*JavaLangObject, c int64) *ApiJavaJavaPairRDD {
	conv_b := javabind.NewGoToJavaMap(javabind.NewGoToJavaCallable(), javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleByKeyExact", "org/apache/spark/api/java/JavaPairRDD", a, conv_b.Value().Cast("java/util/Map"), c)
	if err != nil {
		panic(err)
	}
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> sampleByKeyExact(boolean, java.util.Map<K, java.lang.Object>)
func (jbobject *ApiJavaJavaPairRDD) SampleByKeyExact(a bool, b map[*JavaLangObject]*JavaLangObject) *ApiJavaJavaPairRDD {
	conv_b := javabind.NewGoToJavaMap(javabind.NewGoToJavaCallable(), javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleByKeyExact", "org/apache/spark/api/java/JavaPairRDD", a, conv_b.Value().Cast("java/util/Map"))
	if err != nil {
		panic(err)
	}
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>)
func (jbobject *ApiJavaJavaPairRDD) Union(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> intersection(org.apache.spark.api.java.JavaPairRDD<K, V>)
func (jbobject *ApiJavaJavaPairRDD) Intersection(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersection", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"))
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

// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>, org.apache.spark.Partitioner, boolean, org.apache.spark.serializer.Serializer)
func (jbobject *ApiJavaJavaPairRDD) CombineByKey4(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d PartitionerInterface, e bool, f SerializerSerializerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "combineByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_d.Value().Cast("org/apache/spark/Partitioner"), e, conv_f.Value().Cast("org/apache/spark/serializer/Serializer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_f.CleanUp()
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

// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>, org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) CombineByKey3(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "combineByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_d.Value().Cast("org/apache/spark/Partitioner"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
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

// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>, int)
func (jbobject *ApiJavaJavaPairRDD) CombineByKey2(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "combineByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"), d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> reduceByKey(org.apache.spark.Partitioner, org.apache.spark.api.java.function.Function2<V, V, V>)
func (jbobject *ApiJavaJavaPairRDD) ReduceByKey3(a PartitionerInterface, b ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reduceByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/Partitioner"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"))
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

// public java.util.Map<K, V> reduceByKeyLocally(org.apache.spark.api.java.function.Function2<V, V, V>)
func (jbobject *ApiJavaJavaPairRDD) ReduceByKeyLocally(a ApiJavaFunctionFunction2Interface) map[*JavaLangObject]*JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reduceByKeyLocally", "java/util/Map", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoCallable(), javabind.NewJavaToGoCallable())
	dst := new(map[*JavaLangObject]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.Map<K, java.lang.Object> countByKey()
func (jbobject *ApiJavaJavaPairRDD) CountByKey() map[*JavaLangObject]*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countByKey", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoCallable(), javabind.NewJavaToGoCallable())
	dst := new(map[*JavaLangObject]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.partial.PartialResult<java.util.Map<K, org.apache.spark.partial.BoundedDouble>> countByKeyApprox(long)
func (jbobject *ApiJavaJavaPairRDD) CountByKeyApprox(a int64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countByKeyApprox", "org/apache/spark/partial/PartialResult", a)
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
	unique_x := &PartialPartialResult{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.partial.PartialResult<java.util.Map<K, org.apache.spark.partial.BoundedDouble>> countByKeyApprox(long, double)
func (jbobject *ApiJavaJavaPairRDD) CountByKeyApprox2(a int64, b float64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countByKeyApprox", "org/apache/spark/partial/PartialResult", a, b)
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
	unique_x := &PartialPartialResult{}
	unique_x.Callable = dst
	return unique_x
}

// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> aggregateByKey(U, org.apache.spark.Partitioner, org.apache.spark.api.java.function.Function2<U, V, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
func (jbobject *ApiJavaJavaPairRDD) AggregateByKey3(a interface{}, b PartitionerInterface, c ApiJavaFunctionFunction2Interface, d ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "aggregateByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/Partitioner"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_d.Value().Cast("org/apache/spark/api/java/function/Function2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
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

// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> aggregateByKey(U, int, org.apache.spark.api.java.function.Function2<U, V, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
func (jbobject *ApiJavaJavaPairRDD) AggregateByKey2(a interface{}, b int, c ApiJavaFunctionFunction2Interface, d ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "aggregateByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/Object"), b, conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_d.Value().Cast("org/apache/spark/api/java/function/Function2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
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

// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> aggregateByKey(U, org.apache.spark.api.java.function.Function2<U, V, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
func (jbobject *ApiJavaJavaPairRDD) AggregateByKey(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "aggregateByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> foldByKey(V, org.apache.spark.Partitioner, org.apache.spark.api.java.function.Function2<V, V, V>)
func (jbobject *ApiJavaJavaPairRDD) FoldByKey3(a JavaLangObjectInterface, b PartitionerInterface, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "foldByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/Partitioner"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> foldByKey(V, int, org.apache.spark.api.java.function.Function2<V, V, V>)
func (jbobject *ApiJavaJavaPairRDD) FoldByKey2(a JavaLangObjectInterface, b int, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "foldByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/Object"), b, conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"))
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> foldByKey(V, org.apache.spark.api.java.function.Function2<V, V, V>)
func (jbobject *ApiJavaJavaPairRDD) FoldByKey(a JavaLangObjectInterface, b ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "foldByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> reduceByKey(org.apache.spark.api.java.function.Function2<V, V, V>, int)
func (jbobject *ApiJavaJavaPairRDD) ReduceByKey2(a ApiJavaFunctionFunction2Interface, b int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reduceByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"), b)
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

// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Iterable<V>> groupByKey(org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) GroupByKey3(a PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "groupByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/Partitioner"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Iterable<V>> groupByKey(int)
func (jbobject *ApiJavaJavaPairRDD) GroupByKey2(a int) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "groupByKey", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> subtract(org.apache.spark.api.java.JavaPairRDD<K, V>)
func (jbobject *ApiJavaJavaPairRDD) Subtract(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> subtract(org.apache.spark.api.java.JavaPairRDD<K, V>, int)
func (jbobject *ApiJavaJavaPairRDD) Subtract2(a ApiJavaJavaPairRDDInterface, b int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"), b)
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> subtract(org.apache.spark.api.java.JavaPairRDD<K, V>, org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) Subtract3(a ApiJavaJavaPairRDDInterface, b PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"), conv_b.Value().Cast("org/apache/spark/Partitioner"))
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

// public <W> org.apache.spark.api.java.JavaPairRDD<K, V> subtractByKey(org.apache.spark.api.java.JavaPairRDD<K, W>)
func (jbobject *ApiJavaJavaPairRDD) SubtractByKey(a ApiJavaJavaPairRDDInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtractByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"))
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

// public <W> org.apache.spark.api.java.JavaPairRDD<K, V> subtractByKey(org.apache.spark.api.java.JavaPairRDD<K, W>, int)
func (jbobject *ApiJavaJavaPairRDD) SubtractByKey2(a ApiJavaJavaPairRDDInterface, b int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtractByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"), b)
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

// public <W> org.apache.spark.api.java.JavaPairRDD<K, V> subtractByKey(org.apache.spark.api.java.JavaPairRDD<K, W>, org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) SubtractByKey3(a ApiJavaJavaPairRDDInterface, b PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtractByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"), conv_b.Value().Cast("org/apache/spark/Partitioner"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> partitionBy(org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) PartitionBy(a PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitionBy", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/Partitioner"))
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

// public <C> org.apache.spark.api.java.JavaPairRDD<K, C> combineByKey(org.apache.spark.api.java.function.Function<V, C>, org.apache.spark.api.java.function.Function2<C, V, C>, org.apache.spark.api.java.function.Function2<C, C, C>)
func (jbobject *ApiJavaJavaPairRDD) CombineByKey(a ApiJavaFunctionFunctionInterface, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "combineByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> reduceByKey(org.apache.spark.api.java.function.Function2<V, V, V>)
func (jbobject *ApiJavaJavaPairRDD) ReduceByKey(a ApiJavaFunctionFunction2Interface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reduceByKey", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Iterable<V>> groupByKey()
func (jbobject *ApiJavaJavaPairRDD) GroupByKey() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "groupByKey", "org/apache/spark/api/java/JavaPairRDD")
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Map<K, V> collectAsMap()
func (jbobject *ApiJavaJavaPairRDD) CollectAsMap() map[*JavaLangObject]*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "collectAsMap", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoCallable(), javabind.NewJavaToGoCallable())
	dst := new(map[*JavaLangObject]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> mapValues(org.apache.spark.api.java.function.Function<V, U>)
func (jbobject *ApiJavaJavaPairRDD) MapValues(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapValues", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public <U> org.apache.spark.api.java.JavaPairRDD<K, U> flatMapValues(org.apache.spark.api.java.function.Function<V, java.lang.Iterable<U>>)
func (jbobject *ApiJavaJavaPairRDD) FlatMapValues(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "flatMapValues", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public java.util.List<V> lookup(K)
func (jbobject *ApiJavaJavaPairRDD) Lookup(a JavaLangObjectInterface) []*JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "lookup", "java/util/List", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> repartitionAndSortWithinPartitions(org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) RepartitionAndSortWithinPartitions(a PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "repartitionAndSortWithinPartitions", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/Partitioner"))
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

// public org.apache.spark.api.java.JavaPairRDD<K, V> sortByKey()
func (jbobject *ApiJavaJavaPairRDD) SortByKey() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sortByKey", "org/apache/spark/api/java/JavaPairRDD")
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> sortByKey(boolean)
func (jbobject *ApiJavaJavaPairRDD) SortByKey2(a bool) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sortByKey", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> sortByKey(boolean, int)
func (jbobject *ApiJavaJavaPairRDD) SortByKey3(a bool, b int) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sortByKey", "org/apache/spark/api/java/JavaPairRDD", a, b)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaRDD<K> keys()
func (jbobject *ApiJavaJavaPairRDD) Keys() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "keys", "org/apache/spark/api/java/JavaRDD")
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaRDD<V> values()
func (jbobject *ApiJavaJavaPairRDD) Values() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "values", "org/apache/spark/api/java/JavaRDD")
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Object> countApproxDistinctByKey(double, org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaPairRDD) CountApproxDistinctByKey3(a float64, b PartitionerInterface) *ApiJavaJavaPairRDD {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApproxDistinctByKey", "org/apache/spark/api/java/JavaPairRDD", a, conv_b.Value().Cast("org/apache/spark/Partitioner"))
	if err != nil {
		panic(err)
	}
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

// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Object> countApproxDistinctByKey(double, int)
func (jbobject *ApiJavaJavaPairRDD) CountApproxDistinctByKey2(a float64, b int) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApproxDistinctByKey", "org/apache/spark/api/java/JavaPairRDD", a, b)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, java.lang.Object> countApproxDistinctByKey(double)
func (jbobject *ApiJavaJavaPairRDD) CountApproxDistinctByKey(a float64) *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApproxDistinctByKey", "org/apache/spark/api/java/JavaPairRDD", a)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<K, V> setName(java.lang.String)
func (jbobject *ApiJavaJavaPairRDD) SetName(a string) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setName", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/String"))
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

// public java.lang.Object first()
func (jbobject *ApiJavaJavaPairRDD) First2() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "first", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaRDDLike wrapRDD(org.apache.spark.rdd.RDD)
func (jbobject *ApiJavaJavaPairRDD) WrapRDD(a RddRDDInterface) *ApiJavaJavaRDDLike {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "wrapRDD", "org/apache/spark/api/java/JavaRDDLike", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &ApiJavaJavaRDDLike{}
	unique_x.Callable = dst
	return unique_x
}


