package gospark

import "github.com/timob/javabind"

type UtilStatCounterInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.util.StatCounter merge(double)
	Merge(a float64) *UtilStatCounter

	// public org.apache.spark.util.StatCounter merge(org.apache.spark.util.StatCounter)
	Merge2(a UtilStatCounterInterface) *UtilStatCounter

	// public org.apache.spark.util.StatCounter copy()
	Copy() *UtilStatCounter

	// public long count()
	Count() int64

	// public double mean()
	Mean() float64

	// public double sum()
	Sum() float64

	// public double max()
	Max() float64

	// public double min()
	Min() float64

	// public double variance()
	Variance() float64

	// public double sampleVariance()
	SampleVariance() float64

	// public double stdev()
	Stdev() float64

	// public double sampleStdev()
	SampleStdev() float64
}

type UtilStatCounter struct {
	JavaLangObject
}

// public org.apache.spark.util.StatCounter()
func NewUtilStatCounter() (*UtilStatCounter) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/util/StatCounter")
	if err != nil {
		panic(err)
	}
	x := &UtilStatCounter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.util.StatCounter merge(double)
func (jbobject *UtilStatCounter) Merge(a float64) *UtilStatCounter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "merge", "org/apache/spark/util/StatCounter", a)
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

// public org.apache.spark.util.StatCounter merge(org.apache.spark.util.StatCounter)
func (jbobject *UtilStatCounter) Merge2(a UtilStatCounterInterface) *UtilStatCounter {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "merge", "org/apache/spark/util/StatCounter", conv_a.Value().Cast("org/apache/spark/util/StatCounter"))
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
	unique_x := &UtilStatCounter{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.util.StatCounter copy()
func (jbobject *UtilStatCounter) Copy() *UtilStatCounter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/util/StatCounter")
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

// public long count()
func (jbobject *UtilStatCounter) Count() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "count", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public double mean()
func (jbobject *UtilStatCounter) Mean() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mean", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double sum()
func (jbobject *UtilStatCounter) Sum() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sum", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double max()
func (jbobject *UtilStatCounter) Max() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "max", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double min()
func (jbobject *UtilStatCounter) Min() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "min", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double variance()
func (jbobject *UtilStatCounter) Variance() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "variance", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double sampleVariance()
func (jbobject *UtilStatCounter) SampleVariance() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleVariance", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double stdev()
func (jbobject *UtilStatCounter) Stdev() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stdev", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public double sampleStdev()
func (jbobject *UtilStatCounter) SampleStdev() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sampleStdev", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public java.lang.String toString()
func (jbobject *UtilStatCounter) ToString() string {
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


