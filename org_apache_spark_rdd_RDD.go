package gospark

import "github.com/timob/javabind"

type RddRDDInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.rdd.DoubleRDDFunctions doubleRDDToDoubleRDDFunctions(org.apache.spark.rdd.RDD<java.lang.Object>)
	DoubleRDDToDoubleRDDFunctions(a RddRDDInterface) *RddDoubleRDDFunctions

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.SparkConf conf()
	Conf() *SparkConf

	// public abstract org.apache.spark.Partition[] getPartitions()
	GetPartitions() []*Partition

	// public org.apache.spark.SparkContext sparkContext()
	SparkContext() *SparkContext

	// public int id()
	Id() int

	// public java.lang.String name()
	Name() string

	// public org.apache.spark.rdd.RDD<T> setName(java.lang.String)
	SetName(a string) *RddRDD

	// public org.apache.spark.rdd.RDD<T> persist(org.apache.spark.storage.StorageLevel)
	Persist2(a StorageStorageLevelInterface) *RddRDD

	// public org.apache.spark.rdd.RDD<T> persist()
	Persist() *RddRDD

	// public org.apache.spark.rdd.RDD<T> cache()
	Cache() *RddRDD

	// public org.apache.spark.rdd.RDD<T> unpersist(boolean)
	Unpersist(a bool) *RddRDD

	// public org.apache.spark.storage.StorageLevel getStorageLevel()
	GetStorageLevel() *StorageStorageLevel

	// public final org.apache.spark.Partition[] partitions()
	Partitions() []*Partition

	// public final int getNumPartitions()
	GetNumPartitions() int

	// public org.apache.spark.rdd.RDD<T> distinct()
	Distinct() *RddRDD

	// public org.apache.spark.rdd.RDD<T> sample(boolean, double, long)
	Sample(a bool, b float64, c int64) *RddRDD

	// public org.apache.spark.rdd.RDD<T>[] randomSplit(double[], long)
	RandomSplit(a []float64, b int64) []*RddRDD

	// public org.apache.spark.rdd.RDD<T> randomSampleWithRange(double, double, long)
	RandomSampleWithRange(a float64, b float64, c int64) *RddRDD

	// public java.lang.Object takeSample(boolean, int, long)
	TakeSample(a bool, b int, c int64) *JavaLangObject

	// public org.apache.spark.rdd.RDD<T> union(org.apache.spark.rdd.RDD<T>)
	Union(a RddRDDInterface) *RddRDD

	// public org.apache.spark.rdd.RDD<T> intersection(org.apache.spark.rdd.RDD<T>)
	Intersection(a RddRDDInterface) *RddRDD

	// public org.apache.spark.rdd.RDD<T> intersection(org.apache.spark.rdd.RDD<T>, int)
	Intersection2(a RddRDDInterface, b int) *RddRDD

	// public org.apache.spark.rdd.RDD<java.lang.Object> glom()
	Glom() *RddRDD

	// public org.apache.spark.rdd.RDD<java.lang.String> pipe(java.lang.String)
	Pipe(a string) *RddRDD

	// public java.lang.Object collect()
	Collect() *JavaLangObject

	// public java.lang.Object toArray()
	ToArray() *JavaLangObject

	// public org.apache.spark.rdd.RDD<T> subtract(org.apache.spark.rdd.RDD<T>)
	Subtract(a RddRDDInterface) *RddRDD

	// public org.apache.spark.rdd.RDD<T> subtract(org.apache.spark.rdd.RDD<T>, int)
	Subtract2(a RddRDDInterface, b int) *RddRDD

	// public long count()
	Count() int64

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> countApprox(long, double)
	CountApprox(a int64, b float64) *PartialPartialResult

	// public long countApproxDistinct(int, int)
	CountApproxDistinct2(a int, b int) int64

	// public long countApproxDistinct(double)
	CountApproxDistinct(a float64) int64

	// public java.lang.Object take(int)
	Take(a int) *JavaLangObject

	// public T first()
	First() *JavaLangObject

	// public boolean isEmpty()
	IsEmpty() bool

	// public void saveAsTextFile(java.lang.String)
	SaveAsTextFile(a string) 

	// public void saveAsObjectFile(java.lang.String)
	SaveAsObjectFile(a string) 

	// public java.lang.Object[] collectPartitions()
	CollectPartitions() []*JavaLangObject

	// public void checkpoint()
	Checkpoint() 

	// public org.apache.spark.rdd.RDD<T> localCheckpoint()
	LocalCheckpoint() *RddRDD

	// public boolean isCheckpointed()
	IsCheckpointed() bool

	// public boolean isCheckpointedAndMaterialized()
	IsCheckpointedAndMaterialized() bool

	// public boolean isLocallyCheckpointed()
	IsLocallyCheckpointed() bool

	// public org.apache.spark.util.CallSite creationSite()
	CreationSite() *UtilCallSite

	// public java.lang.String getCreationSite()
	GetCreationSite() string

	// public org.apache.spark.SparkContext context()
	Context() *SparkContext

	// public void doCheckpoint()
	DoCheckpoint() 

	// public void markCheckpointed()
	MarkCheckpointed() 

	// public void clearDependencies()
	ClearDependencies() 

	// public java.lang.String toDebugString()
	ToDebugString() string

	// public org.apache.spark.api.java.JavaRDD<T> toJavaRDD()
	ToJavaRDD() *ApiJavaJavaRDD
}

type RddRDD struct {
	JavaLangObject
}

// public static org.apache.spark.rdd.DoubleRDDFunctions doubleRDDToDoubleRDDFunctions(org.apache.spark.rdd.RDD<java.lang.Object>)
func (jbobject *RddRDD) DoubleRDDToDoubleRDDFunctions(a RddRDDInterface) *RddDoubleRDDFunctions {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/rdd/RDD", "doubleRDDToDoubleRDDFunctions", "org/apache/spark/rdd/DoubleRDDFunctions", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &RddDoubleRDDFunctions{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String logName()
func (jbobject *RddRDD) LogName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "logName", "java/lang/String")
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

// public boolean isTraceEnabled()
func (jbobject *RddRDD) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.SparkConf conf()
func (jbobject *RddRDD) Conf() *SparkConf {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "conf", "org/apache/spark/SparkConf")
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract org.apache.spark.Partition[] getPartitions()
func (jbobject *RddRDD) GetPartitions() []*Partition {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPartitions", javabind.ObjectArrayType("org/apache/spark/Partition"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/Partition")
	dst := new([]*Partition)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.SparkContext sparkContext()
func (jbobject *RddRDD) SparkContext() *SparkContext {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sparkContext", "org/apache/spark/SparkContext")
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
func (jbobject *RddRDD) Id() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String name()
func (jbobject *RddRDD) Name() string {
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

// public org.apache.spark.rdd.RDD<T> setName(java.lang.String)
func (jbobject *RddRDD) SetName(a string) *RddRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setName", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> persist(org.apache.spark.storage.StorageLevel)
func (jbobject *RddRDD) Persist2(a StorageStorageLevelInterface) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "persist", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"))
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> persist()
func (jbobject *RddRDD) Persist() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "persist", "org/apache/spark/rdd/RDD")
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> cache()
func (jbobject *RddRDD) Cache() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cache", "org/apache/spark/rdd/RDD")
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> unpersist(boolean)
func (jbobject *RddRDD) Unpersist(a bool) *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/rdd/RDD", a)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.StorageLevel getStorageLevel()
func (jbobject *RddRDD) GetStorageLevel() *StorageStorageLevel {
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

// public final org.apache.spark.Partition[] partitions()
func (jbobject *RddRDD) Partitions() []*Partition {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitions", javabind.ObjectArrayType("org/apache/spark/Partition"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/Partition")
	dst := new([]*Partition)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final int getNumPartitions()
func (jbobject *RddRDD) GetNumPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNumPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.rdd.RDD<T> distinct()
func (jbobject *RddRDD) Distinct() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/rdd/RDD")
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> sample(boolean, double, long)
func (jbobject *RddRDD) Sample(a bool, b float64, c int64) *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/rdd/RDD", a, b, c)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T>[] randomSplit(double[], long)
func (jbobject *RddRDD) RandomSplit(a []float64, b int64) []*RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "randomSplit", javabind.ObjectArrayType("org/apache/spark/rdd/RDD<T>"), a, b)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/rdd/RDD<T>")
	dst := new([]*RddRDD)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.rdd.RDD<T> randomSampleWithRange(double, double, long)
func (jbobject *RddRDD) RandomSampleWithRange(a float64, b float64, c int64) *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "randomSampleWithRange", "org/apache/spark/rdd/RDD", a, b, c)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object takeSample(boolean, int, long)
func (jbobject *RddRDD) TakeSample(a bool, b int, c int64) *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "takeSample", "java/lang/Object", a, b, c)
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

// public org.apache.spark.rdd.RDD<T> union(org.apache.spark.rdd.RDD<T>)
func (jbobject *RddRDD) Union(a RddRDDInterface) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> intersection(org.apache.spark.rdd.RDD<T>)
func (jbobject *RddRDD) Intersection(a RddRDDInterface) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersection", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> intersection(org.apache.spark.rdd.RDD<T>, int)
func (jbobject *RddRDD) Intersection2(a RddRDDInterface, b int) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersection", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"), b)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<java.lang.Object> glom()
func (jbobject *RddRDD) Glom() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "glom", "org/apache/spark/rdd/RDD")
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<java.lang.String> pipe(java.lang.String)
func (jbobject *RddRDD) Pipe(a string) *RddRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "pipe", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object collect()
func (jbobject *RddRDD) Collect() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "collect", "java/lang/Object")
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

// public java.lang.Object toArray()
func (jbobject *RddRDD) ToArray() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toArray", "java/lang/Object")
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

// public org.apache.spark.rdd.RDD<T> subtract(org.apache.spark.rdd.RDD<T>)
func (jbobject *RddRDD) Subtract(a RddRDDInterface) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<T> subtract(org.apache.spark.rdd.RDD<T>, int)
func (jbobject *RddRDD) Subtract2(a RddRDDInterface, b int) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"), b)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public long count()
func (jbobject *RddRDD) Count() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "count", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> countApprox(long, double)
func (jbobject *RddRDD) CountApprox(a int64, b float64) *PartialPartialResult {
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

// public long countApproxDistinct(int, int)
func (jbobject *RddRDD) CountApproxDistinct2(a int, b int) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApproxDistinct", javabind.Long, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long countApproxDistinct(double)
func (jbobject *RddRDD) CountApproxDistinct(a float64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "countApproxDistinct", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.Object take(int)
func (jbobject *RddRDD) Take(a int) *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "take", "java/lang/Object", a)
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

// public T first()
func (jbobject *RddRDD) First() *JavaLangObject {
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
func (jbobject *RddRDD) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void saveAsTextFile(java.lang.String)
func (jbobject *RddRDD) SaveAsTextFile(a string)  {
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
func (jbobject *RddRDD) SaveAsObjectFile(a string)  {
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

// public java.lang.Object[] collectPartitions()
func (jbobject *RddRDD) CollectPartitions() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "collectPartitions", javabind.ObjectArrayType("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/lang/Object")
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void checkpoint()
func (jbobject *RddRDD) Checkpoint()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "checkpoint", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.apache.spark.rdd.RDD<T> localCheckpoint()
func (jbobject *RddRDD) LocalCheckpoint() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "localCheckpoint", "org/apache/spark/rdd/RDD")
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isCheckpointed()
func (jbobject *RddRDD) IsCheckpointed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCheckpointed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isCheckpointedAndMaterialized()
func (jbobject *RddRDD) IsCheckpointedAndMaterialized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCheckpointedAndMaterialized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isLocallyCheckpointed()
func (jbobject *RddRDD) IsLocallyCheckpointed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isLocallyCheckpointed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.util.CallSite creationSite()
func (jbobject *RddRDD) CreationSite() *UtilCallSite {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "creationSite", "org/apache/spark/util/CallSite")
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
	unique_x := &UtilCallSite{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getCreationSite()
func (jbobject *RddRDD) GetCreationSite() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreationSite", "java/lang/String")
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

// public org.apache.spark.SparkContext context()
func (jbobject *RddRDD) Context() *SparkContext {
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

// public void doCheckpoint()
func (jbobject *RddRDD) DoCheckpoint()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "doCheckpoint", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void markCheckpointed()
func (jbobject *RddRDD) MarkCheckpointed()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markCheckpointed", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void clearDependencies()
func (jbobject *RddRDD) ClearDependencies()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearDependencies", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String toDebugString()
func (jbobject *RddRDD) ToDebugString() string {
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

// public java.lang.String toString()
func (jbobject *RddRDD) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

// public org.apache.spark.api.java.JavaRDD<T> toJavaRDD()
func (jbobject *RddRDD) ToJavaRDD() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toJavaRDD", "org/apache/spark/api/java/JavaRDD")
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


