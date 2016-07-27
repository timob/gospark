package gospark

import "github.com/timob/javabind"

type ApiJavaJavaSparkStatusTrackerInterface interface {
	JavaLangObjectInterface

	// public int[] getJobIdsForGroup(java.lang.String)
	GetJobIdsForGroup(a string) []int

	// public int[] getActiveStageIds()
	GetActiveStageIds() []int

	// public int[] getActiveJobIds()
	GetActiveJobIds() []int

	// public org.apache.spark.SparkJobInfo getJobInfo(int)
	GetJobInfo(a int) *SparkJobInfo

	// public org.apache.spark.SparkStageInfo getStageInfo(int)
	GetStageInfo(a int) *SparkStageInfo
}

type ApiJavaJavaSparkStatusTracker struct {
	JavaLangObject
}

// public org.apache.spark.api.java.JavaSparkStatusTracker(org.apache.spark.SparkContext)
func NewApiJavaJavaSparkStatusTracker(a SparkContextInterface) (*ApiJavaJavaSparkStatusTracker) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkStatusTracker", conv_a.Value().Cast("org/apache/spark/SparkContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ApiJavaJavaSparkStatusTracker{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int[] getJobIdsForGroup(java.lang.String)
func (jbobject *ApiJavaJavaSparkStatusTracker) GetJobIdsForGroup(a string) []int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJobIdsForGroup", javabind.Int | javabind.Array, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.([]int)
}

// public int[] getActiveStageIds()
func (jbobject *ApiJavaJavaSparkStatusTracker) GetActiveStageIds() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActiveStageIds", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int[] getActiveJobIds()
func (jbobject *ApiJavaJavaSparkStatusTracker) GetActiveJobIds() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActiveJobIds", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.apache.spark.SparkJobInfo getJobInfo(int)
func (jbobject *ApiJavaJavaSparkStatusTracker) GetJobInfo(a int) *SparkJobInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJobInfo", "org/apache/spark/SparkJobInfo", a)
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
	unique_x := &SparkJobInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkStageInfo getStageInfo(int)
func (jbobject *ApiJavaJavaSparkStatusTracker) GetStageInfo(a int) *SparkStageInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStageInfo", "org/apache/spark/SparkStageInfo", a)
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
	unique_x := &SparkStageInfo{}
	unique_x.Callable = dst
	return unique_x
}


