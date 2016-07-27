package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerUnpersistRDDInterface interface {
	JavaLangObjectInterface

	// public int rddId()
	RddId() int

	// public org.apache.spark.scheduler.SparkListenerUnpersistRDD copy(int)
	Copy(a int) *SchedulerSparkListenerUnpersistRDD

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerUnpersistRDD struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerUnpersistRDD(int)
func NewSchedulerSparkListenerUnpersistRDD(a int) (*SchedulerSparkListenerUnpersistRDD) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerUnpersistRDD", a)
	if err != nil {
		panic(err)
	}
	x := &SchedulerSparkListenerUnpersistRDD{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int rddId()
func (jbobject *SchedulerSparkListenerUnpersistRDD) RddId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rddId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.scheduler.SparkListenerUnpersistRDD copy(int)
func (jbobject *SchedulerSparkListenerUnpersistRDD) Copy(a int) *SchedulerSparkListenerUnpersistRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerUnpersistRDD", a)
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
	unique_x := &SchedulerSparkListenerUnpersistRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerUnpersistRDD) ProductPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productPrefix", "java/lang/String")
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

// public int productArity()
func (jbobject *SchedulerSparkListenerUnpersistRDD) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerUnpersistRDD) ProductElement(a int) *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productElement", "java/lang/Object", a)
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

// public boolean canEqual(java.lang.Object)
func (jbobject *SchedulerSparkListenerUnpersistRDD) CanEqual(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canEqual", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *SchedulerSparkListenerUnpersistRDD) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerUnpersistRDD) ToString() string {
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

// public boolean equals(java.lang.Object)
func (jbobject *SchedulerSparkListenerUnpersistRDD) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}


