package gospark

import "github.com/timob/javabind"

type SparkJobInfoInterface interface {

	// public abstract int jobId()
	JobId() int

	// public abstract int[] stageIds()
	StageIds() []int

	// public abstract org.apache.spark.JobExecutionStatus status()
	Status() *JobExecutionStatus
}

type SparkJobInfo struct {
	*javabind.Callable
}

// public abstract int jobId()
func (jbobject *SparkJobInfo) JobId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jobId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int[] stageIds()
func (jbobject *SparkJobInfo) StageIds() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageIds", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public abstract org.apache.spark.JobExecutionStatus status()
func (jbobject *SparkJobInfo) Status() *JobExecutionStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "status", "org/apache/spark/JobExecutionStatus")
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}


