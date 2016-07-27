package gospark

import "github.com/timob/javabind"

type SchedulerOutputCommitCoordinatorInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public boolean isEmpty()
	IsEmpty() bool

	// public boolean canCommit(int, int, int)
	CanCommit(a int, b int, c int) bool

	// public void stageStart(int, int)
	StageStart(a int, b int) 

	// public void stageEnd(int)
	StageEnd(a int) 

	// public synchronized void taskCompleted(int, int, int, org.apache.spark.TaskEndReason)
	TaskCompleted(a int, b int, c int, d TaskEndReasonInterface) 

	// public synchronized void stop()
	Stop() 

	// public synchronized boolean handleAskPermissionToCommit(int, int, int)
	HandleAskPermissionToCommit(a int, b int, c int) bool
}

type SchedulerOutputCommitCoordinator struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.OutputCommitCoordinator(org.apache.spark.SparkConf, boolean)
func NewSchedulerOutputCommitCoordinator(a SparkConfInterface, b bool) (*SchedulerOutputCommitCoordinator) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/OutputCommitCoordinator", conv_a.Value().Cast("org/apache/spark/SparkConf"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SchedulerOutputCommitCoordinator{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *SchedulerOutputCommitCoordinator) LogName() string {
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
func (jbobject *SchedulerOutputCommitCoordinator) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isEmpty()
func (jbobject *SchedulerOutputCommitCoordinator) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean canCommit(int, int, int)
func (jbobject *SchedulerOutputCommitCoordinator) CanCommit(a int, b int, c int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canCommit", javabind.Boolean, a, b, c)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void stageStart(int, int)
func (jbobject *SchedulerOutputCommitCoordinator) StageStart(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stageStart", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void stageEnd(int)
func (jbobject *SchedulerOutputCommitCoordinator) StageEnd(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stageEnd", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public synchronized void taskCompleted(int, int, int, org.apache.spark.TaskEndReason)
func (jbobject *SchedulerOutputCommitCoordinator) TaskCompleted(a int, b int, c int, d TaskEndReasonInterface)  {
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "taskCompleted", javabind.Void, a, b, c, conv_d.Value().Cast("org/apache/spark/TaskEndReason"))
	if err != nil {
		panic(err)
	}
	conv_d.CleanUp()

}

// public synchronized void stop()
func (jbobject *SchedulerOutputCommitCoordinator) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public synchronized boolean handleAskPermissionToCommit(int, int, int)
func (jbobject *SchedulerOutputCommitCoordinator) HandleAskPermissionToCommit(a int, b int, c int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "handleAskPermissionToCommit", javabind.Boolean, a, b, c)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


