package gospark

import "github.com/timob/javabind"

type ApiJavaJavaRDDInterface interface {
	ApiJavaAbstractJavaRDDLikeInterface

	// public static <T> org.apache.spark.rdd.RDD<T> toRDD(org.apache.spark.api.java.JavaRDD<T>)
	ToRDD(a ApiJavaJavaRDDInterface) *RddRDD

	// public org.apache.spark.rdd.RDD<T> rdd()
	Rdd() *RddRDD

	// public org.apache.spark.api.java.JavaRDD<T> wrapRDD(org.apache.spark.rdd.RDD<T>)
	WrapRDD2(a RddRDDInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> cache()
	Cache() *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> persist(org.apache.spark.storage.StorageLevel)
	Persist(a StorageStorageLevelInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> unpersist()
	Unpersist() *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> unpersist(boolean)
	Unpersist2(a bool) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> distinct()
	Distinct() *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> distinct(int)
	Distinct2(a int) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> filter(org.apache.spark.api.java.function.Function<T, java.lang.Boolean>)
	Filter(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> coalesce(int)
	Coalesce(a int) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> coalesce(int, boolean)
	Coalesce2(a int, b bool) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> repartition(int)
	Repartition(a int) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> sample(boolean, double)
	Sample(a bool, b float64) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> sample(boolean, double, long)
	Sample2(a bool, b float64, c int64) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T>[] randomSplit(double[])
	RandomSplit(a []float64) []*ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T>[] randomSplit(double[], long)
	RandomSplit2(a []float64, b int64) []*ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>)
	Union(a ApiJavaJavaRDDInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> intersection(org.apache.spark.api.java.JavaRDD<T>)
	Intersection(a ApiJavaJavaRDDInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> subtract(org.apache.spark.api.java.JavaRDD<T>)
	Subtract(a ApiJavaJavaRDDInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> subtract(org.apache.spark.api.java.JavaRDD<T>, int)
	Subtract2(a ApiJavaJavaRDDInterface, b int) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> subtract(org.apache.spark.api.java.JavaRDD<T>, org.apache.spark.Partitioner)
	Subtract3(a ApiJavaJavaRDDInterface, b PartitionerInterface) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<T> setName(java.lang.String)
	SetName(a string) *ApiJavaJavaRDD

	// public <S> org.apache.spark.api.java.JavaRDD<T> sortBy(org.apache.spark.api.java.function.Function<T, S>, boolean, int)
	SortBy(a ApiJavaFunctionFunctionInterface, b bool, c int) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDDLike wrapRDD(org.apache.spark.rdd.RDD)
	WrapRDD(a RddRDDInterface) *ApiJavaJavaRDDLike
}

type ApiJavaJavaRDD struct {
	ApiJavaAbstractJavaRDDLike
}

// public static <T> org.apache.spark.rdd.RDD<T> toRDD(org.apache.spark.api.java.JavaRDD<T>)
func (jbobject *ApiJavaJavaRDD) ToRDD(a ApiJavaJavaRDDInterface) *RddRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/api/java/JavaRDD", "toRDD", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"))
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

// public org.apache.spark.rdd.RDD<T> rdd()
func (jbobject *ApiJavaJavaRDD) Rdd() *RddRDD {
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

// public org.apache.spark.api.java.JavaRDD<T> wrapRDD(org.apache.spark.rdd.RDD<T>)
func (jbobject *ApiJavaJavaRDD) WrapRDD2(a RddRDDInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "wrapRDD", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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

// public org.apache.spark.api.java.JavaRDD<T> cache()
func (jbobject *ApiJavaJavaRDD) Cache() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cache", "org/apache/spark/api/java/JavaRDD")
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

// public org.apache.spark.api.java.JavaRDD<T> persist(org.apache.spark.storage.StorageLevel)
func (jbobject *ApiJavaJavaRDD) Persist(a StorageStorageLevelInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "persist", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"))
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

// public org.apache.spark.api.java.JavaRDD<T> unpersist()
func (jbobject *ApiJavaJavaRDD) Unpersist() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/api/java/JavaRDD")
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

// public org.apache.spark.api.java.JavaRDD<T> unpersist(boolean)
func (jbobject *ApiJavaJavaRDD) Unpersist2(a bool) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unpersist", "org/apache/spark/api/java/JavaRDD", a)
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

// public org.apache.spark.api.java.JavaRDD<T> distinct()
func (jbobject *ApiJavaJavaRDD) Distinct() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/api/java/JavaRDD")
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

// public org.apache.spark.api.java.JavaRDD<T> distinct(int)
func (jbobject *ApiJavaJavaRDD) Distinct2(a int) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "distinct", "org/apache/spark/api/java/JavaRDD", a)
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

// public org.apache.spark.api.java.JavaRDD<T> filter(org.apache.spark.api.java.function.Function<T, java.lang.Boolean>)
func (jbobject *ApiJavaJavaRDD) Filter(a ApiJavaFunctionFunctionInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "filter", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"))
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

// public org.apache.spark.api.java.JavaRDD<T> coalesce(int)
func (jbobject *ApiJavaJavaRDD) Coalesce(a int) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "coalesce", "org/apache/spark/api/java/JavaRDD", a)
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

// public org.apache.spark.api.java.JavaRDD<T> coalesce(int, boolean)
func (jbobject *ApiJavaJavaRDD) Coalesce2(a int, b bool) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "coalesce", "org/apache/spark/api/java/JavaRDD", a, b)
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

// public org.apache.spark.api.java.JavaRDD<T> repartition(int)
func (jbobject *ApiJavaJavaRDD) Repartition(a int) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "repartition", "org/apache/spark/api/java/JavaRDD", a)
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

// public org.apache.spark.api.java.JavaRDD<T> sample(boolean, double)
func (jbobject *ApiJavaJavaRDD) Sample(a bool, b float64) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/api/java/JavaRDD", a, b)
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

// public org.apache.spark.api.java.JavaRDD<T> sample(boolean, double, long)
func (jbobject *ApiJavaJavaRDD) Sample2(a bool, b float64, c int64) *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sample", "org/apache/spark/api/java/JavaRDD", a, b, c)
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

// public org.apache.spark.api.java.JavaRDD<T>[] randomSplit(double[])
func (jbobject *ApiJavaJavaRDD) RandomSplit(a []float64) []*ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "randomSplit", javabind.ObjectArrayType("org/apache/spark/api/java/JavaRDD<T>"), a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/api/java/JavaRDD<T>")
	dst := new([]*ApiJavaJavaRDD)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.api.java.JavaRDD<T>[] randomSplit(double[], long)
func (jbobject *ApiJavaJavaRDD) RandomSplit2(a []float64, b int64) []*ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "randomSplit", javabind.ObjectArrayType("org/apache/spark/api/java/JavaRDD<T>"), a, b)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/api/java/JavaRDD<T>")
	dst := new([]*ApiJavaJavaRDD)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>)
func (jbobject *ApiJavaJavaRDD) Union(a ApiJavaJavaRDDInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
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

// public org.apache.spark.api.java.JavaRDD<T> intersection(org.apache.spark.api.java.JavaRDD<T>)
func (jbobject *ApiJavaJavaRDD) Intersection(a ApiJavaJavaRDDInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersection", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"))
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

// public org.apache.spark.api.java.JavaRDD<T> subtract(org.apache.spark.api.java.JavaRDD<T>)
func (jbobject *ApiJavaJavaRDD) Subtract(a ApiJavaJavaRDDInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"))
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

// public org.apache.spark.api.java.JavaRDD<T> subtract(org.apache.spark.api.java.JavaRDD<T>, int)
func (jbobject *ApiJavaJavaRDD) Subtract2(a ApiJavaJavaRDDInterface, b int) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"), b)
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

// public org.apache.spark.api.java.JavaRDD<T> subtract(org.apache.spark.api.java.JavaRDD<T>, org.apache.spark.Partitioner)
func (jbobject *ApiJavaJavaRDD) Subtract3(a ApiJavaJavaRDDInterface, b PartitionerInterface) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"), conv_b.Value().Cast("org/apache/spark/Partitioner"))
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

// public java.lang.String toString()
func (jbobject *ApiJavaJavaRDD) ToString() string {
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

// public org.apache.spark.api.java.JavaRDD<T> setName(java.lang.String)
func (jbobject *ApiJavaJavaRDD) SetName(a string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setName", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"))
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

// public <S> org.apache.spark.api.java.JavaRDD<T> sortBy(org.apache.spark.api.java.function.Function<T, S>, boolean, int)
func (jbobject *ApiJavaJavaRDD) SortBy(a ApiJavaFunctionFunctionInterface, b bool, c int) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sortBy", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/function/Function"), b, c)
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

// public org.apache.spark.api.java.JavaRDDLike wrapRDD(org.apache.spark.rdd.RDD)
func (jbobject *ApiJavaJavaRDD) WrapRDD(a RddRDDInterface) *ApiJavaJavaRDDLike {
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


