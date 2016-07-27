package gospark

import "github.com/timob/javabind"

type RddDoubleRDDFunctionsInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public double sum()
	Sum() float64

	// public org.apache.spark.util.StatCounter stats()
	Stats() *UtilStatCounter

	// public double mean()
	Mean() float64

	// public double variance()
	Variance() float64

	// public double stdev()
	Stdev() float64

	// public double sampleStdev()
	SampleStdev() float64

	// public double sampleVariance()
	SampleVariance() float64

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> meanApprox(long, double)
	MeanApprox(a int64, b float64) *PartialPartialResult

	// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> sumApprox(long, double)
	SumApprox(a int64, b float64) *PartialPartialResult

	// public long[] histogram(double[], boolean)
	Histogram(a []float64, b bool) []int64
}

type RddDoubleRDDFunctions struct {
	JavaLangObject
}

// public org.apache.spark.rdd.DoubleRDDFunctions(org.apache.spark.rdd.RDD<java.lang.Object>)
func NewRddDoubleRDDFunctions(a RddRDDInterface) (*RddDoubleRDDFunctions) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/rdd/DoubleRDDFunctions", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &RddDoubleRDDFunctions{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *RddDoubleRDDFunctions) LogName() string {
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
func (jbobject *RddDoubleRDDFunctions) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public double sum()
func (jbobject *RddDoubleRDDFunctions) Sum() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sum", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public org.apache.spark.util.StatCounter stats()
func (jbobject *RddDoubleRDDFunctions) Stats() *UtilStatCounter {
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

// public double mean()
func (jbobject *RddDoubleRDDFunctions) Mean() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mean", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double variance()
func (jbobject *RddDoubleRDDFunctions) Variance() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "variance", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double stdev()
func (jbobject *RddDoubleRDDFunctions) Stdev() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stdev", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double sampleStdev()
func (jbobject *RddDoubleRDDFunctions) SampleStdev() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleStdev", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double sampleVariance()
func (jbobject *RddDoubleRDDFunctions) SampleVariance() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleVariance", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> meanApprox(long, double)
func (jbobject *RddDoubleRDDFunctions) MeanApprox(a int64, b float64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "meanApprox", "org/apache/spark/partial/PartialResult", a, b)
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

// public org.apache.spark.partial.PartialResult<org.apache.spark.partial.BoundedDouble> sumApprox(long, double)
func (jbobject *RddDoubleRDDFunctions) SumApprox(a int64, b float64) *PartialPartialResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sumApprox", "org/apache/spark/partial/PartialResult", a, b)
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

// public long[] histogram(double[], boolean)
func (jbobject *RddDoubleRDDFunctions) Histogram(a []float64, b bool) []int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "histogram", javabind.Long | javabind.Array, a, b)
	if err != nil {
		panic(err)
	}
	return jret.([]int64)
}


