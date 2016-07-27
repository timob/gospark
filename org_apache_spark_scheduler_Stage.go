package gospark

import "github.com/timob/javabind"

type SchedulerStageInterface interface {
	JavaLangObjectInterface

	// public static int MAX_CONSECUTIVE_FETCH_FAILURES()
	MAX_CONSECUTIVE_FETCH_FAILURES() int

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public int id()
	Id() int

	// public org.apache.spark.rdd.RDD<?> rdd()
	Rdd() *RddRDD

	// public int numTasks()
	NumTasks() int

	// public int firstJobId()
	FirstJobId() int

	// public org.apache.spark.util.CallSite callSite()
	CallSite() *UtilCallSite

	// public int numPartitions()
	NumPartitions() int

	// public java.lang.String name()
	Name() string

	// public java.lang.String details()
	Details() string

	// public void resetInternalAccumulators()
	ResetInternalAccumulators() 

	// public void clearFailures()
	ClearFailures() 

	// public boolean failedOnFetchAndShouldAbort(int)
	FailedOnFetchAndShouldAbort(a int) bool

	// public org.apache.spark.scheduler.StageInfo latestInfo()
	LatestInfo() *SchedulerStageInfo
}

type SchedulerStage struct {
	JavaLangObject
}

// public static int MAX_CONSECUTIVE_FETCH_FAILURES()
func (jbobject *SchedulerStage) MAX_CONSECUTIVE_FETCH_FAILURES() int {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/scheduler/Stage", "MAX_CONSECUTIVE_FETCH_FAILURES", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String logName()
func (jbobject *SchedulerStage) LogName() string {
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
func (jbobject *SchedulerStage) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int id()
func (jbobject *SchedulerStage) Id() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.rdd.RDD<?> rdd()
func (jbobject *SchedulerStage) Rdd() *RddRDD {
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

// public int numTasks()
func (jbobject *SchedulerStage) NumTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int firstJobId()
func (jbobject *SchedulerStage) FirstJobId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "firstJobId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.util.CallSite callSite()
func (jbobject *SchedulerStage) CallSite() *UtilCallSite {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "callSite", "org/apache/spark/util/CallSite")
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

// public int numPartitions()
func (jbobject *SchedulerStage) NumPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String name()
func (jbobject *SchedulerStage) Name() string {
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

// public java.lang.String details()
func (jbobject *SchedulerStage) Details() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "details", "java/lang/String")
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

// public void resetInternalAccumulators()
func (jbobject *SchedulerStage) ResetInternalAccumulators()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "resetInternalAccumulators", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void clearFailures()
func (jbobject *SchedulerStage) ClearFailures()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearFailures", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean failedOnFetchAndShouldAbort(int)
func (jbobject *SchedulerStage) FailedOnFetchAndShouldAbort(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "failedOnFetchAndShouldAbort", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.scheduler.StageInfo latestInfo()
func (jbobject *SchedulerStage) LatestInfo() *SchedulerStageInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "latestInfo", "org/apache/spark/scheduler/StageInfo")
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
	unique_x := &SchedulerStageInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public final int hashCode()
func (jbobject *SchedulerStage) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final boolean equals(java.lang.Object)
func (jbobject *SchedulerStage) Equals(a interface{}) bool {
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


