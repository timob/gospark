package gospark

import "github.com/timob/javabind"

type SchedulerTaskInfoInterface interface {
	JavaLangObjectInterface

	// public long taskId()
	TaskId() int64

	// public int index()
	Index() int

	// public int attemptNumber()
	AttemptNumber() int

	// public long launchTime()
	LaunchTime() int64

	// public java.lang.String executorId()
	ExecutorId() string

	// public java.lang.String host()
	Host() string

	// public boolean speculative()
	Speculative() bool

	// public long gettingResultTime()
	GettingResultTime() int64

	// public long finishTime()
	FinishTime() int64

	// public boolean failed()
	Failed() bool

	// public void markGettingResult(long)
	MarkGettingResult(a int64) 

	// public void markSuccessful(long)
	MarkSuccessful(a int64) 

	// public void markFailed(long)
	MarkFailed(a int64) 

	// public boolean gettingResult()
	GettingResult() bool

	// public boolean finished()
	Finished() bool

	// public boolean successful()
	Successful() bool

	// public boolean running()
	Running() bool

	// public java.lang.String status()
	Status() string

	// public int attempt()
	Attempt() int

	// public java.lang.String id()
	Id() string

	// public long duration()
	Duration() int64

	// public long timeRunning(long)
	TimeRunning(a int64) int64
}

type SchedulerTaskInfo struct {
	JavaLangObject
}

// public long taskId()
func (jbobject *SchedulerTaskInfo) TaskId() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskId", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public int index()
func (jbobject *SchedulerTaskInfo) Index() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "index", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int attemptNumber()
func (jbobject *SchedulerTaskInfo) AttemptNumber() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attemptNumber", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long launchTime()
func (jbobject *SchedulerTaskInfo) LaunchTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "launchTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.String executorId()
func (jbobject *SchedulerTaskInfo) ExecutorId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorId", "java/lang/String")
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

// public java.lang.String host()
func (jbobject *SchedulerTaskInfo) Host() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "host", "java/lang/String")
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

// public boolean speculative()
func (jbobject *SchedulerTaskInfo) Speculative() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "speculative", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long gettingResultTime()
func (jbobject *SchedulerTaskInfo) GettingResultTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "gettingResultTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long finishTime()
func (jbobject *SchedulerTaskInfo) FinishTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "finishTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean failed()
func (jbobject *SchedulerTaskInfo) Failed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "failed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void markGettingResult(long)
func (jbobject *SchedulerTaskInfo) MarkGettingResult(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markGettingResult", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void markSuccessful(long)
func (jbobject *SchedulerTaskInfo) MarkSuccessful(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markSuccessful", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void markFailed(long)
func (jbobject *SchedulerTaskInfo) MarkFailed(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markFailed", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean gettingResult()
func (jbobject *SchedulerTaskInfo) GettingResult() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "gettingResult", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean finished()
func (jbobject *SchedulerTaskInfo) Finished() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "finished", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean successful()
func (jbobject *SchedulerTaskInfo) Successful() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "successful", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean running()
func (jbobject *SchedulerTaskInfo) Running() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "running", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String status()
func (jbobject *SchedulerTaskInfo) Status() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "status", "java/lang/String")
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

// public int attempt()
func (jbobject *SchedulerTaskInfo) Attempt() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attempt", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String id()
func (jbobject *SchedulerTaskInfo) Id() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", "java/lang/String")
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

// public long duration()
func (jbobject *SchedulerTaskInfo) Duration() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "duration", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long timeRunning(long)
func (jbobject *SchedulerTaskInfo) TimeRunning(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "timeRunning", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


