package gospark

import "github.com/timob/javabind"

type ApiJavaAbstractJavaRDDLikeInterface interface {
	JavaLangObjectInterface

	// public java.util.List<org.apache.spark.Partition> splits()
	Splits() []*Partition

	// public java.util.List<org.apache.spark.Partition> partitions()
	Partitions() []*Partition

	// public int getNumPartitions()
	GetNumPartitions() int

	// public org.apache.spark.SparkContext context()
	Context() *SparkContext

	// public int id()
	Id() int

	// public org.apache.spark.storage.StorageLevel getStorageLevel()
	GetStorageLevel() *StorageStorageLevel

	// public java.util.Iterator<T> iterator(org.apache.spark.Partition, org.apache.spark.TaskContext)
	Iterator(a PartitionInterface, b TaskContextInterface) []*JavaLangObject

	// public <R> org.apache.spark.api.java.JavaRDD<R> map(org.apache.spark.api.java.function.Function<T, R>)
	Map(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaRDD

	// public <R> org.apache.spark.api.java.JavaRDD<R> mapPartitionsWithIndex(org.apache.spark.api.java.function.Function2<java.lang.Integer, java.util.Iterator<T>, java.util.Iterator<R>>, boolean)
	MapPartitionsWithIndex(a ApiJavaFunctionFunction2Interface, b bool) *ApiJavaJavaRDD

	// public <R> org.apache.spark.api.java.JavaDoubleRDD mapToDouble(org.apache.spark.api.java.function.DoubleFunction<T>)
	MapToDouble(a ApiJavaFunctionDoubleFunctionInterface) *ApiJavaJavaDoubleRDD

	// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> mapToPair(org.apache.spark.api.java.function.PairFunction<T, K2, V2>)
	MapToPair(a ApiJavaFunctionPairFunctionInterface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaRDD<U> flatMap(org.apache.spark.api.java.function.FlatMapFunction<T, U>)
	FlatMap(a ApiJavaFunctionFlatMapFunctionInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaDoubleRDD flatMapToDouble(org.apache.spark.api.java.function.DoubleFlatMapFunction<T>)
	FlatMapToDouble(a ApiJavaFunctionDoubleFlatMapFunctionInterface) *ApiJavaJavaDoubleRDD

	// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> flatMapToPair(org.apache.spark.api.java.function.PairFlatMapFunction<T, K2, V2>)
	FlatMapToPair(a ApiJavaFunctionPairFlatMapFunctionInterface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaRDD<U> mapPartitions(org.apache.spark.api.java.function.FlatMapFunction<java.util.Iterator<T>, U>)
	MapPartitions(a ApiJavaFunctionFlatMapFunctionInterface) *ApiJavaJavaRDD

	// public <U> org.apache.spark.api.java.JavaRDD<U> mapPartitions(org.apache.spark.api.java.function.FlatMapFunction<java.util.Iterator<T>, U>, boolean)
	MapPartitions2(a ApiJavaFunctionFlatMapFunctionInterface, b bool) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaDoubleRDD mapPartitionsToDouble(org.apache.spark.api.java.function.DoubleFlatMapFunction<java.util.Iterator<T>>)
	MapPartitionsToDouble(a ApiJavaFunctionDoubleFlatMapFunctionInterface) *ApiJavaJavaDoubleRDD

	// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> mapPartitionsToPair(org.apache.spark.api.java.function.PairFlatMapFunction<java.util.Iterator<T>, K2, V2>)
	MapPartitionsToPair(a ApiJavaFunctionPairFlatMapFunctionInterface) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaDoubleRDD mapPartitionsToDouble(org.apache.spark.api.java.function.DoubleFlatMapFunction<java.util.Iterator<T>>, boolean)
	MapPartitionsToDouble2(a ApiJavaFunctionDoubleFlatMapFunctionInterface, b bool) *ApiJavaJavaDoubleRDD

	// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> mapPartitionsToPair(org.apache.spark.api.java.function.PairFlatMapFunction<java.util.Iterator<T>, K2, V2>, boolean)
	MapPartitionsToPair2(a ApiJavaFunctionPairFlatMapFunctionInterface, b bool) *ApiJavaJavaPairRDD

	// public void foreachPartition(org.apache.spark.api.java.function.VoidFunction<java.util.Iterator<T>>)
	ForeachPartition(a ApiJavaFunctionVoidFunctionInterface) 

	// public org.apache.spark.api.java.JavaRDD<java.util.List<T>> glom()
	Glom() *ApiJavaJavaRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<T, U> cartesian(org.apache.spark.api.java.JavaRDDLike<U, ?>)
	Cartesian(a ApiJavaJavaRDDLikeInterface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<U, java.lang.Iterable<T>> groupBy(org.apache.spark.api.java.function.Function<T, U>)
	GroupBy(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<U, java.lang.Iterable<T>> groupBy(org.apache.spark.api.java.function.Function<T, U>, int)
	GroupBy2(a ApiJavaFunctionFunctionInterface, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaRDD<java.lang.String> pipe(java.lang.String)
	Pipe(a string) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<java.lang.String> pipe(java.util.List<java.lang.String>)
	Pipe2(a []string) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<java.lang.String> pipe(java.util.List<java.lang.String>, java.util.Map<java.lang.String, java.lang.String>)
	Pipe3(a []string, b map[string]string) *ApiJavaJavaRDD

	// public <U> org.apache.spark.api.java.JavaPairRDD<T, U> zip(org.apache.spark.api.java.JavaRDDLike<U, ?>)
	Zip(a ApiJavaJavaRDDLikeInterface) *ApiJavaJavaPairRDD

	// public <U, V> org.apache.spark.api.java.JavaRDD<V> zipPartitions(org.apache.spark.api.java.JavaRDDLike<U, ?>, org.apache.spark.api.java.function.FlatMapFunction2<java.util.Iterator<T>, java.util.Iterator<U>, V>)
	ZipPartitions(a ApiJavaJavaRDDLikeInterface, b ApiJavaFunctionFlatMapFunction2Interface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaPairRDD<T, java.lang.Long> zipWithUniqueId()
	ZipWithUniqueId() *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<T, java.lang.Long> zipWithIndex()
	ZipWithIndex() *ApiJavaJavaPairRDD

	// public void foreach(org.apache.spark.api.java.function.VoidFunction<T>)
	Foreach(a ApiJavaFunctionVoidFunctionInterface) 

	// public java.util.List<T> collect()
	Collect() []*JavaLangObject

	// public java.util.Iterator<T> toLocalIterator()
	ToLocalIterator() []*JavaLangObject

	// public java.util.List<T> toArray()
	ToArray() []*JavaLangObject

	// public java.util.List<T>[] collectPartitions(int[])
	CollectPartitions(a []int) []*[]*JavaLangObject

	// public T reduce(org.apache.spark.api.java.function.Function2<T, T, T>)
	Reduce(a ApiJavaFunctionFunction2Interface) *JavaLangObject

	// public T treeReduce(org.apache.spark.api.java.function.Function2<T, T, T>, int)
	TreeReduce2(a ApiJavaFunctionFunction2Interface, b int) *JavaLangObject

	// public T treeReduce(org.apache.spark.api.java.function.Function2<T, T, T>)
	TreeReduce(a ApiJavaFunctionFunction2Interface) *JavaLangObject

	// public T fold(T, org.apache.spark.api.java.function.Function2<T, T, T>)
	Fold(a JavaLangObjectInterface, b ApiJavaFunctionFunction2Interface) *JavaLangObject

	// public <U> U aggregate(U, org.apache.spark.api.java.function.Function2<U, T, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
	Aggregate(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *JavaLangObject

	// public <U> U treeAggregate(U, org.apache.spark.api.java.function.Function2<U, T, U>, org.apache.spark.api.java.function.Function2<U, U, U>, int)
	TreeAggregate2(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d int) *JavaLangObject

	// public <U> U treeAggregate(U, org.apache.spark.api.java.function.Function2<U, T, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
	TreeAggregate(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *JavaLangObject

	// public long count()
	Count() int64

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> countApprox(long, double)
	CountApprox2(a int64, b float64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> countApprox(long)
	CountApprox(a int64) *PartialPartialResult

	// public java.util.Map<T, java.lang.Long> countByValue()
	CountByValue() map[*JavaLangObject]int64

	// public org.apache.spark.partial.PartialResult<java.util.Map<T, org.apache.spark.partial.BoundedDouble>> countByValueApprox(long, double)
	CountByValueApprox2(a int64, b float64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<java.util.Map<T, org.apache.spark.partial.BoundedDouble>> countByValueApprox(long)
	CountByValueApprox(a int64) *PartialPartialResult

	// public java.util.List<T> take(int)
	Take(a int) []*JavaLangObject

	// public java.util.List<T> takeSample(boolean, int)
	TakeSample(a bool, b int) []*JavaLangObject

	// public java.util.List<T> takeSample(boolean, int, long)
	TakeSample2(a bool, b int, c int64) []*JavaLangObject

	// public T first()
	First() *JavaLangObject

	// public boolean isEmpty()
	IsEmpty() bool

	// public void saveAsTextFile(java.lang.String)
	SaveAsTextFile(a string) 

	// public void saveAsObjectFile(java.lang.String)
	SaveAsObjectFile(a string) 

	// public <U> org.apache.spark.api.java.JavaPairRDD<U, T> keyBy(org.apache.spark.api.java.function.Function<T, U>)
	KeyBy(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD

	// public void checkpoint()
	Checkpoint() 

	// public boolean isCheckpointed()
	IsCheckpointed() bool

	// public java.lang.String toDebugString()
	ToDebugString() string

	// public java.util.List<T> top(int)
	Top(a int) []*JavaLangObject

	// public java.util.List<T> takeOrdered(int)
	TakeOrdered(a int) []*JavaLangObject

	// public long countApproxDistinct(double)
	CountApproxDistinct(a float64) int64

	// public java.lang.String name()
	Name() string

	// public org.apache.spark.api.java.JavaFutureAction<java.lang.Long> countAsync()
	CountAsync() *ApiJavaJavaFutureAction

	// public org.apache.spark.api.java.JavaFutureAction<java.util.List<T>> collectAsync()
	CollectAsync() *ApiJavaJavaFutureAction

	// public org.apache.spark.api.java.JavaFutureAction<java.util.List<T>> takeAsync(int)
	TakeAsync(a int) *ApiJavaJavaFutureAction

	// public org.apache.spark.api.java.JavaFutureAction<java.lang.Void> foreachAsync(org.apache.spark.api.java.function.VoidFunction<T>)
	ForeachAsync(a ApiJavaFunctionVoidFunctionInterface) *ApiJavaJavaFutureAction

	// public org.apache.spark.api.java.JavaFutureAction<java.lang.Void> foreachPartitionAsync(org.apache.spark.api.java.function.VoidFunction<java.util.Iterator<T>>)
	ForeachPartitionAsync(a ApiJavaFunctionVoidFunctionInterface) *ApiJavaJavaFutureAction
}

type ApiJavaAbstractJavaRDDLike struct {
	JavaLangObject
}

// public org.apache.spark.api.java.AbstractJavaRDDLike()
func NewApiJavaAbstractJavaRDDLike() (*ApiJavaAbstractJavaRDDLike) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/AbstractJavaRDDLike")
	if err != nil {
		panic(err)
	}
	x := &ApiJavaAbstractJavaRDDLike{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<org.apache.spark.Partition> splits()
func (jbobject *ApiJavaAbstractJavaRDDLike) Splits() []*Partition {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "splits", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*Partition)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<org.apache.spark.Partition> partitions()
func (jbobject *ApiJavaAbstractJavaRDDLike) Partitions() []*Partition {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*Partition)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getNumPartitions()
func (jbobject *ApiJavaAbstractJavaRDDLike) GetNumPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNumPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.SparkContext context()
func (jbobject *ApiJavaAbstractJavaRDDLike) Context() *SparkContext {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "context", "org/apache/spark/SparkContext")
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
	unique_x := &SparkContext{}
	unique_x.Callable = dst
	return unique_x
}

// public int id()
func (jbobject *ApiJavaAbstractJavaRDDLike) Id() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.storage.StorageLevel getStorageLevel()
func (jbobject *ApiJavaAbstractJavaRDDLike) GetStorageLevel() *StorageStorageLevel {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStorageLevel", "org/apache/spark/storage/StorageLevel")
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
	unique_x := &StorageStorageLevel{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Iterator<T> iterator(org.apache.spark.Partition, org.apache.spark.TaskContext)
func (jbobject *ApiJavaAbstractJavaRDDLike) Iterator(a PartitionInterface, b TaskContextInterface) []*JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "iterator", "java/util/Iterator", conv_a.Value().Cast("org/apache/spark/Partition"), conv_b.Value().Cast("org/apache/spark/TaskContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoIterator(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public <R> org.apache.spark.api.java.JavaRDD<R> map(org.apache.spark.api.java.function.Function<T, R>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Map(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "map", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public <R> org.apache.spark.api.java.JavaRDD<R> mapPartitionsWithIndex(org.apache.spark.api.java.function.Function2<java.lang.Integer, java.util.Iterator<T>, java.util.Iterator<R>>, boolean)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitionsWithIndex(a ApiJavaFunctionFunction2Interface, b bool) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitionsWithIndex", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"), b)
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

// public <R> org.apache.spark.api.java.JavaDoubleRDD mapToDouble(org.apache.spark.api.java.function.DoubleFunction<T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapToDouble(a ApiJavaFunctionDoubleFunctionInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapToDouble", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/DoubleFunction"))
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

// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> mapToPair(org.apache.spark.api.java.function.PairFunction<T, K2, V2>)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapToPair(a ApiJavaFunctionPairFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapToPair", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/PairFunction"))
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

// public <U> org.apache.spark.api.java.JavaRDD<U> flatMap(org.apache.spark.api.java.function.FlatMapFunction<T, U>)
func (jbobject *ApiJavaAbstractJavaRDDLike) FlatMap(a ApiJavaFunctionFlatMapFunctionInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "flatMap", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/FlatMapFunction"))
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

// public org.apache.spark.api.java.JavaDoubleRDD flatMapToDouble(org.apache.spark.api.java.function.DoubleFlatMapFunction<T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) FlatMapToDouble(a ApiJavaFunctionDoubleFlatMapFunctionInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "flatMapToDouble", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/DoubleFlatMapFunction"))
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

// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> flatMapToPair(org.apache.spark.api.java.function.PairFlatMapFunction<T, K2, V2>)
func (jbobject *ApiJavaAbstractJavaRDDLike) FlatMapToPair(a ApiJavaFunctionPairFlatMapFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "flatMapToPair", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/PairFlatMapFunction"))
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

// public <U> org.apache.spark.api.java.JavaRDD<U> mapPartitions(org.apache.spark.api.java.function.FlatMapFunction<java.util.Iterator<T>, U>)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitions(a ApiJavaFunctionFlatMapFunctionInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitions", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/FlatMapFunction"))
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

// public <U> org.apache.spark.api.java.JavaRDD<U> mapPartitions(org.apache.spark.api.java.function.FlatMapFunction<java.util.Iterator<T>, U>, boolean)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitions2(a ApiJavaFunctionFlatMapFunctionInterface, b bool) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitions", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/FlatMapFunction"), b)
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

// public org.apache.spark.api.java.JavaDoubleRDD mapPartitionsToDouble(org.apache.spark.api.java.function.DoubleFlatMapFunction<java.util.Iterator<T>>)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitionsToDouble(a ApiJavaFunctionDoubleFlatMapFunctionInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitionsToDouble", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/DoubleFlatMapFunction"))
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

// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> mapPartitionsToPair(org.apache.spark.api.java.function.PairFlatMapFunction<java.util.Iterator<T>, K2, V2>)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitionsToPair(a ApiJavaFunctionPairFlatMapFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitionsToPair", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/PairFlatMapFunction"))
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

// public org.apache.spark.api.java.JavaDoubleRDD mapPartitionsToDouble(org.apache.spark.api.java.function.DoubleFlatMapFunction<java.util.Iterator<T>>, boolean)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitionsToDouble2(a ApiJavaFunctionDoubleFlatMapFunctionInterface, b bool) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitionsToDouble", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/DoubleFlatMapFunction"), b)
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

// public <K2, V2> org.apache.spark.api.java.JavaPairRDD<K2, V2> mapPartitionsToPair(org.apache.spark.api.java.function.PairFlatMapFunction<java.util.Iterator<T>, K2, V2>, boolean)
func (jbobject *ApiJavaAbstractJavaRDDLike) MapPartitionsToPair2(a ApiJavaFunctionPairFlatMapFunctionInterface, b bool) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapPartitionsToPair", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/PairFlatMapFunction"), b)
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

// public void foreachPartition(org.apache.spark.api.java.function.VoidFunction<java.util.Iterator<T>>)
func (jbobject *ApiJavaAbstractJavaRDDLike) ForeachPartition(a ApiJavaFunctionVoidFunctionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "foreachPartition", javabind.Void, conv_a.Value().Cast("org/apache/spark/api/java/function/VoidFunction"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.apache.spark.api.java.JavaRDD<java.util.List<T>> glom()
func (jbobject *ApiJavaAbstractJavaRDDLike) Glom() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "glom", "org/apache/spark/api/java/JavaRDD")
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

// public <U> org.apache.spark.api.java.JavaPairRDD<T, U> cartesian(org.apache.spark.api.java.JavaRDDLike<U, ?>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Cartesian(a ApiJavaJavaRDDLikeInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cartesian", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDDLike"))
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

// public <U> org.apache.spark.api.java.JavaPairRDD<U, java.lang.Iterable<T>> groupBy(org.apache.spark.api.java.function.Function<T, U>)
func (jbobject *ApiJavaAbstractJavaRDDLike) GroupBy(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "groupBy", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public <U> org.apache.spark.api.java.JavaPairRDD<U, java.lang.Iterable<T>> groupBy(org.apache.spark.api.java.function.Function<T, U>, int)
func (jbobject *ApiJavaAbstractJavaRDDLike) GroupBy2(a ApiJavaFunctionFunctionInterface, b int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "groupBy", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"), b)
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

// public org.apache.spark.api.java.JavaRDD<java.lang.String> pipe(java.lang.String)
func (jbobject *ApiJavaAbstractJavaRDDLike) Pipe(a string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "pipe", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"))
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

// public org.apache.spark.api.java.JavaRDD<java.lang.String> pipe(java.util.List<java.lang.String>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Pipe2(a []string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "pipe", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/util/List"))
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

// public org.apache.spark.api.java.JavaRDD<java.lang.String> pipe(java.util.List<java.lang.String>, java.util.Map<java.lang.String, java.lang.String>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Pipe3(a []string, b map[string]string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaString())
	conv_b := javabind.NewGoToJavaMap(javabind.NewGoToJavaString(), javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "pipe", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/util/List"), conv_b.Value().Cast("java/util/Map"))
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

// public <U> org.apache.spark.api.java.JavaPairRDD<T, U> zip(org.apache.spark.api.java.JavaRDDLike<U, ?>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Zip(a ApiJavaJavaRDDLikeInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "zip", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDDLike"))
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

// public <U, V> org.apache.spark.api.java.JavaRDD<V> zipPartitions(org.apache.spark.api.java.JavaRDDLike<U, ?>, org.apache.spark.api.java.function.FlatMapFunction2<java.util.Iterator<T>, java.util.Iterator<U>, V>)
func (jbobject *ApiJavaAbstractJavaRDDLike) ZipPartitions(a ApiJavaJavaRDDLikeInterface, b ApiJavaFunctionFlatMapFunction2Interface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "zipPartitions", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDDLike"), conv_b.Value().Cast("org/apache/spark/api/java/function/FlatMapFunction2"))
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

// public org.apache.spark.api.java.JavaPairRDD<T, java.lang.Long> zipWithUniqueId()
func (jbobject *ApiJavaAbstractJavaRDDLike) ZipWithUniqueId() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "zipWithUniqueId", "org/apache/spark/api/java/JavaPairRDD")
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

// public org.apache.spark.api.java.JavaPairRDD<T, java.lang.Long> zipWithIndex()
func (jbobject *ApiJavaAbstractJavaRDDLike) ZipWithIndex() *ApiJavaJavaPairRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "zipWithIndex", "org/apache/spark/api/java/JavaPairRDD")
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

// public void foreach(org.apache.spark.api.java.function.VoidFunction<T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Foreach(a ApiJavaFunctionVoidFunctionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "foreach", javabind.Void, conv_a.Value().Cast("org/apache/spark/api/java/function/VoidFunction"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.List<T> collect()
func (jbobject *ApiJavaAbstractJavaRDDLike) Collect() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "collect", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.Iterator<T> toLocalIterator()
func (jbobject *ApiJavaAbstractJavaRDDLike) ToLocalIterator() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toLocalIterator", "java/util/Iterator")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoIterator(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<T> toArray()
func (jbobject *ApiJavaAbstractJavaRDDLike) ToArray() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toArray", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<T>[] collectPartitions(int[])
func (jbobject *ApiJavaAbstractJavaRDDLike) CollectPartitions(a []int) []*[]*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "collectPartitions", javabind.ObjectArrayType("java/util/List<T>"), a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoList(javabind.NewJavaToGoCallable()), "java/util/List<T>")
	dst := new([]*[]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public T reduce(org.apache.spark.api.java.function.Function2<T, T, T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Reduce(a ApiJavaFunctionFunction2Interface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reduce", "java/lang/Object", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public T treeReduce(org.apache.spark.api.java.function.Function2<T, T, T>, int)
func (jbobject *ApiJavaAbstractJavaRDDLike) TreeReduce2(a ApiJavaFunctionFunction2Interface, b int) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "treeReduce", "java/lang/Object", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"), b)
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public T treeReduce(org.apache.spark.api.java.function.Function2<T, T, T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) TreeReduce(a ApiJavaFunctionFunction2Interface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "treeReduce", "java/lang/Object", conv_a.Value().Cast("org/apache/spark/api/java/function/Function2"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public T fold(T, org.apache.spark.api.java.function.Function2<T, T, T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Fold(a JavaLangObjectInterface, b ApiJavaFunctionFunction2Interface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "fold", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public <U> U aggregate(U, org.apache.spark.api.java.function.Function2<U, T, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
func (jbobject *ApiJavaAbstractJavaRDDLike) Aggregate(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *JavaLangObject {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "aggregate", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public <U> U treeAggregate(U, org.apache.spark.api.java.function.Function2<U, T, U>, org.apache.spark.api.java.function.Function2<U, U, U>, int)
func (jbobject *ApiJavaAbstractJavaRDDLike) TreeAggregate2(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface, d int) *JavaLangObject {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "treeAggregate", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"), d)
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public <U> U treeAggregate(U, org.apache.spark.api.java.function.Function2<U, T, U>, org.apache.spark.api.java.function.Function2<U, U, U>)
func (jbobject *ApiJavaAbstractJavaRDDLike) TreeAggregate(a interface{}, b ApiJavaFunctionFunction2Interface, c ApiJavaFunctionFunction2Interface) *JavaLangObject {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "treeAggregate", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/api/java/function/Function2"), conv_c.Value().Cast("org/apache/spark/api/java/function/Function2"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public long count()
func (jbobject *ApiJavaAbstractJavaRDDLike) Count() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "count", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> countApprox(long, double)
func (jbobject *ApiJavaAbstractJavaRDDLike) CountApprox2(a int64, b float64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApprox", "org/apache/spark/partial/PartialResult", a, b)
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

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> countApprox(long)
func (jbobject *ApiJavaAbstractJavaRDDLike) CountApprox(a int64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApprox", "org/apache/spark/partial/PartialResult", a)
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

// public java.util.Map<T, java.lang.Long> countByValue()
func (jbobject *ApiJavaAbstractJavaRDDLike) CountByValue() map[*JavaLangObject]int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countByValue", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoCallable(), javabind.NewJavaToGoLong())
	dst := new(map[*JavaLangObject]int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.partial.PartialResult<java.util.Map<T, org.apache.spark.partial.BoundedDouble>> countByValueApprox(long, double)
func (jbobject *ApiJavaAbstractJavaRDDLike) CountByValueApprox2(a int64, b float64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countByValueApprox", "org/apache/spark/partial/PartialResult", a, b)
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

// public org.apache.spark.partial.PartialResult<java.util.Map<T, org.apache.spark.partial.BoundedDouble>> countByValueApprox(long)
func (jbobject *ApiJavaAbstractJavaRDDLike) CountByValueApprox(a int64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countByValueApprox", "org/apache/spark/partial/PartialResult", a)
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

// public java.util.List<T> take(int)
func (jbobject *ApiJavaAbstractJavaRDDLike) Take(a int) []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "take", "java/util/List", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<T> takeSample(boolean, int)
func (jbobject *ApiJavaAbstractJavaRDDLike) TakeSample(a bool, b int) []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "takeSample", "java/util/List", a, b)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<T> takeSample(boolean, int, long)
func (jbobject *ApiJavaAbstractJavaRDDLike) TakeSample2(a bool, b int, c int64) []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "takeSample", "java/util/List", a, b, c)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public T first()
func (jbobject *ApiJavaAbstractJavaRDDLike) First() *JavaLangObject {
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

// public boolean isEmpty()
func (jbobject *ApiJavaAbstractJavaRDDLike) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void saveAsTextFile(java.lang.String)
func (jbobject *ApiJavaAbstractJavaRDDLike) SaveAsTextFile(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "saveAsTextFile", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void saveAsObjectFile(java.lang.String)
func (jbobject *ApiJavaAbstractJavaRDDLike) SaveAsObjectFile(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "saveAsObjectFile", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public <U> org.apache.spark.api.java.JavaPairRDD<U, T> keyBy(org.apache.spark.api.java.function.Function<T, U>)
func (jbobject *ApiJavaAbstractJavaRDDLike) KeyBy(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "keyBy", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public void checkpoint()
func (jbobject *ApiJavaAbstractJavaRDDLike) Checkpoint()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "checkpoint", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean isCheckpointed()
func (jbobject *ApiJavaAbstractJavaRDDLike) IsCheckpointed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCheckpointed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toDebugString()
func (jbobject *ApiJavaAbstractJavaRDDLike) ToDebugString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toDebugString", "java/lang/String")
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

// public java.util.List<T> top(int)
func (jbobject *ApiJavaAbstractJavaRDDLike) Top(a int) []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "top", "java/util/List", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<T> takeOrdered(int)
func (jbobject *ApiJavaAbstractJavaRDDLike) TakeOrdered(a int) []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "takeOrdered", "java/util/List", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public long countApproxDistinct(double)
func (jbobject *ApiJavaAbstractJavaRDDLike) CountApproxDistinct(a float64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApproxDistinct", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.String name()
func (jbobject *ApiJavaAbstractJavaRDDLike) Name() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "name", "java/lang/String")
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

// public org.apache.spark.api.java.JavaFutureAction<java.lang.Long> countAsync()
func (jbobject *ApiJavaAbstractJavaRDDLike) CountAsync() *ApiJavaJavaFutureAction {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countAsync", "org/apache/spark/api/java/JavaFutureAction")
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
	unique_x := &ApiJavaJavaFutureAction{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaFutureAction<java.util.List<T>> collectAsync()
func (jbobject *ApiJavaAbstractJavaRDDLike) CollectAsync() *ApiJavaJavaFutureAction {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "collectAsync", "org/apache/spark/api/java/JavaFutureAction")
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
	unique_x := &ApiJavaJavaFutureAction{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaFutureAction<java.util.List<T>> takeAsync(int)
func (jbobject *ApiJavaAbstractJavaRDDLike) TakeAsync(a int) *ApiJavaJavaFutureAction {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "takeAsync", "org/apache/spark/api/java/JavaFutureAction", a)
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
	unique_x := &ApiJavaJavaFutureAction{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaFutureAction<java.lang.Void> foreachAsync(org.apache.spark.api.java.function.VoidFunction<T>)
func (jbobject *ApiJavaAbstractJavaRDDLike) ForeachAsync(a ApiJavaFunctionVoidFunctionInterface) *ApiJavaJavaFutureAction {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "foreachAsync", "org/apache/spark/api/java/JavaFutureAction", conv_a.Value().Cast("org/apache/spark/api/java/function/VoidFunction"))
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
	unique_x := &ApiJavaJavaFutureAction{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaFutureAction<java.lang.Void> foreachPartitionAsync(org.apache.spark.api.java.function.VoidFunction<java.util.Iterator<T>>)
func (jbobject *ApiJavaAbstractJavaRDDLike) ForeachPartitionAsync(a ApiJavaFunctionVoidFunctionInterface) *ApiJavaJavaFutureAction {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "foreachPartitionAsync", "org/apache/spark/api/java/JavaFutureAction", conv_a.Value().Cast("org/apache/spark/api/java/function/VoidFunction"))
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
	unique_x := &ApiJavaJavaFutureAction{}
	unique_x.Callable = dst
	return unique_x
}


