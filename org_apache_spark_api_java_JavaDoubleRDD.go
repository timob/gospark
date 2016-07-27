package gospark

import "github.com/timob/javabind"

type ApiJavaJavaDoubleRDDInterface interface {
	ApiJavaAbstractJavaRDDLikeInterface

	// public static org.apache.spark.rdd.RDD<java.lang.Object> toRDD(org.apache.spark.api.java.JavaDoubleRDD)
	ToRDD(a ApiJavaJavaDoubleRDDInterface) *RddRDD

	// public static org.apache.spark.api.java.JavaDoubleRDD fromRDD(org.apache.spark.rdd.RDD<java.lang.Object>)
	FromRDD(a RddRDDInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.rdd.RDD<java.lang.Object> srdd()
	Srdd() *RddRDD

	// public org.apache.spark.rdd.RDD<java.lang.Double> rdd()
	Rdd() *RddRDD

	// public org.apache.spark.api.java.JavaDoubleRDD wrapRDD(org.apache.spark.rdd.RDD<java.lang.Double>)
	WrapRDD2(a RddRDDInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD cache()
	Cache() *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD persist(org.apache.spark.storage.StorageLevel)
	Persist(a StorageStorageLevelInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD unpersist()
	Unpersist() *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD unpersist(boolean)
	Unpersist2(a bool) *ApiJavaJavaDoubleRDD

	// public java.lang.Double first()
	First2() float64

	// public org.apache.spark.api.java.JavaDoubleRDD distinct()
	Distinct() *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD distinct(int)
	Distinct2(a int) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD filter(org.apache.spark.api.java.function.Function<java.lang.Double, java.lang.Boolean>)
	Filter(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD coalesce(int)
	Coalesce(a int) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD coalesce(int, boolean)
	Coalesce2(a int, b bool) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD repartition(int)
	Repartition(a int) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD subtract(org.apache.spark.api.java.JavaDoubleRDD)
	Subtract(a ApiJavaJavaDoubleRDDInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD subtract(org.apache.spark.api.java.JavaDoubleRDD, int)
	Subtract2(a ApiJavaJavaDoubleRDDInterface, b int) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD subtract(org.apache.spark.api.java.JavaDoubleRDD, org.apache.spark.Partitioner)
	Subtract3(a ApiJavaJavaDoubleRDDInterface, b PartitionerInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD sample(boolean, java.lang.Double)
	Sample(a bool, b float64) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD sample(boolean, java.lang.Double, long)
	Sample2(a bool, b float64, c int64) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD)
	Union(a ApiJavaJavaDoubleRDDInterface) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD intersection(org.apache.spark.api.java.JavaDoubleRDD)
	Intersection(a ApiJavaJavaDoubleRDDInterface) *ApiJavaJavaDoubleRDD

	// public java.lang.Double sum()
	Sum() float64

	// public java.lang.Double min()
	Min() float64

	// public java.lang.Double max()
	Max() float64

	// public org.apache.spark.util.StatCounter stats()
	Stats() *UtilStatCounter

	// public java.lang.Double mean()
	Mean() float64

	// public java.lang.Double variance()
	Variance() float64

	// public java.lang.Double stdev()
	Stdev() float64

	// public java.lang.Double sampleStdev()
	SampleStdev() float64

	// public java.lang.Double sampleVariance()
	SampleVariance() float64

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> meanApprox(long, java.lang.Double)
	MeanApprox2(a int64, b float64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> meanApprox(long)
	MeanApprox(a int64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> sumApprox(long, java.lang.Double)
	SumApprox2(a int64, b float64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> sumApprox(long)
	SumApprox(a int64) *PartialPartialResult

	// public long[] histogram(double[])
	Histogram(a []float64) []int64

	// public long[] histogram(java.lang.Double[], boolean)
	Histogram2(a []float64, b bool) []int64

	// public org.apache.spark.api.java.JavaDoubleRDD setName(java.lang.String)
	SetName(a string) *ApiJavaJavaDoubleRDD

	// public java.lang.Object first()
	First3() *JavaLangObject

	// public org.apache.spark.api.java.JavaRDDLike wrapRDD(org.apache.spark.rdd.RDD)
	WrapRDD(a RddRDDInterface) *ApiJavaJavaRDDLike
}

type ApiJavaJavaDoubleRDD struct {
	ApiJavaAbstractJavaRDDLike
}

// public org.apache.spark.api.java.JavaDoubleRDD(org.apache.spark.rdd.RDD<java.lang.Object>)
func NewApiJavaJavaDoubleRDD(a RddRDDInterface) (*ApiJavaJavaDoubleRDD) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ApiJavaJavaDoubleRDD{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.rdd.RDD<java.lang.Object> toRDD(org.apache.spark.api.java.JavaDoubleRDD)
func (jbobject *ApiJavaJavaDoubleRDD) ToRDD(a ApiJavaJavaDoubleRDDInterface) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/api/java/JavaDoubleRDD", "toRDD", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"))
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

// public static org.apache.spark.api.java.JavaDoubleRDD fromRDD(org.apache.spark.rdd.RDD<java.lang.Object>)
func (jbobject *ApiJavaJavaDoubleRDD) FromRDD(a RddRDDInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/api/java/JavaDoubleRDD", "fromRDD", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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

// public org.apache.spark.rdd.RDD<java.lang.Object> srdd()
func (jbobject *ApiJavaJavaDoubleRDD) Srdd() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "srdd", "org/apache/spark/rdd/RDD")
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

// public org.apache.spark.rdd.RDD<java.lang.Double> rdd()
func (jbobject *ApiJavaJavaDoubleRDD) Rdd() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rdd", "org/apache/spark/rdd/RDD")
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

// public org.apache.spark.api.java.JavaDoubleRDD wrapRDD(org.apache.spark.rdd.RDD<java.lang.Double>)
func (jbobject *ApiJavaJavaDoubleRDD) WrapRDD2(a RddRDDInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "wrapRDD", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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

// public org.apache.spark.api.java.JavaDoubleRDD cache()
func (jbobject *ApiJavaJavaDoubleRDD) Cache() *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cache", "org/apache/spark/api/java/JavaDoubleRDD")
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD persist(org.apache.spark.storage.StorageLevel)
func (jbobject *ApiJavaJavaDoubleRDD) Persist(a StorageStorageLevelInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "persist", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"))
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

// public org.apache.spark.api.java.JavaDoubleRDD unpersist()
func (jbobject *ApiJavaJavaDoubleRDD) Unpersist() *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/api/java/JavaDoubleRDD")
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD unpersist(boolean)
func (jbobject *ApiJavaJavaDoubleRDD) Unpersist2(a bool) *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/api/java/JavaDoubleRDD", a)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Double first()
func (jbobject *ApiJavaJavaDoubleRDD) First2() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "first", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.api.java.JavaDoubleRDD distinct()
func (jbobject *ApiJavaJavaDoubleRDD) Distinct() *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/api/java/JavaDoubleRDD")
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD distinct(int)
func (jbobject *ApiJavaJavaDoubleRDD) Distinct2(a int) *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/api/java/JavaDoubleRDD", a)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD filter(org.apache.spark.api.java.function.Function<java.lang.Double, java.lang.Boolean>)
func (jbobject *ApiJavaJavaDoubleRDD) Filter(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "filter", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public org.apache.spark.api.java.JavaDoubleRDD coalesce(int)
func (jbobject *ApiJavaJavaDoubleRDD) Coalesce(a int) *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "coalesce", "org/apache/spark/api/java/JavaDoubleRDD", a)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD coalesce(int, boolean)
func (jbobject *ApiJavaJavaDoubleRDD) Coalesce2(a int, b bool) *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "coalesce", "org/apache/spark/api/java/JavaDoubleRDD", a, b)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD repartition(int)
func (jbobject *ApiJavaJavaDoubleRDD) Repartition(a int) *ApiJavaJavaDoubleRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "repartition", "org/apache/spark/api/java/JavaDoubleRDD", a)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD subtract(org.apache.spark.api.java.JavaDoubleRDD)
func (jbobject *ApiJavaJavaDoubleRDD) Subtract(a ApiJavaJavaDoubleRDDInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"))
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

// public org.apache.spark.api.java.JavaDoubleRDD subtract(org.apache.spark.api.java.JavaDoubleRDD, int)
func (jbobject *ApiJavaJavaDoubleRDD) Subtract2(a ApiJavaJavaDoubleRDDInterface, b int) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"), b)
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

// public org.apache.spark.api.java.JavaDoubleRDD subtract(org.apache.spark.api.java.JavaDoubleRDD, org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaDoubleRDD) Subtract3(a ApiJavaJavaDoubleRDDInterface, b PartitionerInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"), conv_b.Value().Cast("org/apache/spark/Partitioner"))
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

// public org.apache.spark.api.java.JavaDoubleRDD sample(boolean, java.lang.Double)
func (jbobject *ApiJavaJavaDoubleRDD) Sample(a bool, b float64) *ApiJavaJavaDoubleRDD {
	conv_b := javabind.NewGoToJavaDouble()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/api/java/JavaDoubleRDD", a, conv_b.Value().Cast("java/lang/Double"))
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD sample(boolean, java.lang.Double, long)
func (jbobject *ApiJavaJavaDoubleRDD) Sample2(a bool, b float64, c int64) *ApiJavaJavaDoubleRDD {
	conv_b := javabind.NewGoToJavaDouble()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/api/java/JavaDoubleRDD", a, conv_b.Value().Cast("java/lang/Double"), c)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD)
func (jbobject *ApiJavaJavaDoubleRDD) Union(a ApiJavaJavaDoubleRDDInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
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

// public org.apache.spark.api.java.JavaDoubleRDD intersection(org.apache.spark.api.java.JavaDoubleRDD)
func (jbobject *ApiJavaJavaDoubleRDD) Intersection(a ApiJavaJavaDoubleRDDInterface) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersection", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"))
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

// public java.lang.Double sum()
func (jbobject *ApiJavaJavaDoubleRDD) Sum() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sum", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Double min()
func (jbobject *ApiJavaJavaDoubleRDD) Min() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "min", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Double max()
func (jbobject *ApiJavaJavaDoubleRDD) Max() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "max", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.util.StatCounter stats()
func (jbobject *ApiJavaJavaDoubleRDD) Stats() *UtilStatCounter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stats", "org/apache/spark/util/StatCounter")
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
	unique_x := &UtilStatCounter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Double mean()
func (jbobject *ApiJavaJavaDoubleRDD) Mean() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mean", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Double variance()
func (jbobject *ApiJavaJavaDoubleRDD) Variance() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "variance", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Double stdev()
func (jbobject *ApiJavaJavaDoubleRDD) Stdev() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stdev", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Double sampleStdev()
func (jbobject *ApiJavaJavaDoubleRDD) SampleStdev() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleStdev", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Double sampleVariance()
func (jbobject *ApiJavaJavaDoubleRDD) SampleVariance() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleVariance", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> meanApprox(long, java.lang.Double)
func (jbobject *ApiJavaJavaDoubleRDD) MeanApprox2(a int64, b float64) *PartialPartialResult {
	conv_b := javabind.NewGoToJavaDouble()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "meanApprox", "org/apache/spark/partial/PartialResult", a, conv_b.Value().Cast("java/lang/Double"))
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
	unique_x := &PartialPartialResult{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> meanApprox(long)
func (jbobject *ApiJavaJavaDoubleRDD) MeanApprox(a int64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "meanApprox", "org/apache/spark/partial/PartialResult", a)
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

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> sumApprox(long, java.lang.Double)
func (jbobject *ApiJavaJavaDoubleRDD) SumApprox2(a int64, b float64) *PartialPartialResult {
	conv_b := javabind.NewGoToJavaDouble()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sumApprox", "org/apache/spark/partial/PartialResult", a, conv_b.Value().Cast("java/lang/Double"))
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
	unique_x := &PartialPartialResult{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> sumApprox(long)
func (jbobject *ApiJavaJavaDoubleRDD) SumApprox(a int64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sumApprox", "org/apache/spark/partial/PartialResult", a)
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

// public long[] histogram(double[])
func (jbobject *ApiJavaJavaDoubleRDD) Histogram(a []float64) []int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "histogram", javabind.Long | javabind.Array, a)
	if err != nil {
		panic(err)
	}
	return jret.([]int64)
}

// public long[] histogram(java.lang.Double[], boolean)
func (jbobject *ApiJavaJavaDoubleRDD) Histogram2(a []float64, b bool) []int64 {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaDouble(), "java/lang/Double")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "histogram", javabind.Long | javabind.Array, conv_a.Value().Cast("java/lang/Double"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.([]int64)
}

// public org.apache.spark.api.java.JavaDoubleRDD setName(java.lang.String)
func (jbobject *ApiJavaJavaDoubleRDD) SetName(a string) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setName", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("java/lang/String"))
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

// public java.lang.Object first()
func (jbobject *ApiJavaJavaDoubleRDD) First3() *JavaLangObject {
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
func (jbobject *ApiJavaJavaDoubleRDD) WrapRDD(a RddRDDInterface) *ApiJavaJavaRDDLike {
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


