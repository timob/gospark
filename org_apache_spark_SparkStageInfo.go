package gospark

import "github.com/timob/javabind"

type SparkStageInfoInterface interface {

	// public abstract int stageId()
	StageId() int

	// public abstract int currentAttemptId()
	CurrentAttemptId() int

	// public abstract long submissionTime()
	SubmissionTime() int64

	// public abstract java.lang.String name()
	Name() string

	// public abstract int numTasks()
	NumTasks() int

	// public abstract int numActiveTasks()
	NumActiveTasks() int

	// public abstract int numCompletedTasks()
	NumCompletedTasks() int

	// public abstract int numFailedTasks()
	NumFailedTasks() int
}

type SparkStageInfo struct {
	*javabind.Callable
}

// public abstract int stageId()
func (jbobject *SparkStageInfo) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int currentAttemptId()
func (jbobject *SparkStageInfo) CurrentAttemptId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "currentAttemptId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract long submissionTime()
func (jbobject *SparkStageInfo) SubmissionTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "submissionTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public abstract java.lang.String name()
func (jbobject *SparkStageInfo) Name() string {
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

// public abstract int numTasks()
func (jbobject *SparkStageInfo) NumTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int numActiveTasks()
func (jbobject *SparkStageInfo) NumActiveTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numActiveTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int numCompletedTasks()
func (jbobject *SparkStageInfo) NumCompletedTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numCompletedTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int numFailedTasks()
func (jbobject *SparkStageInfo) NumFailedTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numFailedTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


