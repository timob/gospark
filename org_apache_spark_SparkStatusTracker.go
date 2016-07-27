package gospark

import "github.com/timob/javabind"

type SparkStatusTrackerInterface interface {
	JavaLangObjectInterface

	// public int[] getJobIdsForGroup(java.lang.String)
	GetJobIdsForGroup(a string) []int

	// public int[] getActiveStageIds()
	GetActiveStageIds() []int

	// public int[] getActiveJobIds()
	GetActiveJobIds() []int
}

type SparkStatusTracker struct {
	JavaLangObject
}

// public org.apache.spark.SparkStatusTracker(org.apache.spark.SparkContext)
func NewSparkStatusTracker(a SparkContextInterface) (*SparkStatusTracker) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkStatusTracker", conv_a.Value().Cast("org/apache/spark/SparkContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SparkStatusTracker{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int[] getJobIdsForGroup(java.lang.String)
func (jbobject *SparkStatusTracker) GetJobIdsForGroup(a string) []int {
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
func (jbobject *SparkStatusTracker) GetActiveStageIds() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActiveStageIds", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int[] getActiveJobIds()
func (jbobject *SparkStatusTracker) GetActiveJobIds() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActiveJobIds", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}


