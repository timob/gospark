package gospark

import "github.com/timob/javabind"

type SchedulerDAGSchedulerInterface interface {
	JavaLangObjectInterface

	// public static int RESUBMIT_TIMEOUT()
	RESUBMIT_TIMEOUT() int

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.SparkContext sc()
	Sc() *SparkContext

	// public org.apache.spark.scheduler.TaskScheduler taskScheduler()
	TaskScheduler() *SchedulerTaskScheduler

	// public org.apache.spark.scheduler.DAGSchedulerSource metricsSource()
	MetricsSource() *SchedulerDAGSchedulerSource

	// public int numTotalJobs()
	NumTotalJobs() int

	// public org.apache.spark.scheduler.OutputCommitCoordinator outputCommitCoordinator()
	OutputCommitCoordinator() *SchedulerOutputCommitCoordinator

	// public org.apache.spark.scheduler.DAGSchedulerEventProcessLoop eventProcessLoop()
	EventProcessLoop() *SchedulerDAGSchedulerEventProcessLoop

	// public void taskStarted(org.apache.spark.scheduler.Task<?>, org.apache.spark.scheduler.TaskInfo)
	TaskStarted(a SchedulerTaskInterface, b SchedulerTaskInfoInterface) 

	// public void taskGettingResult(org.apache.spark.scheduler.TaskInfo)
	TaskGettingResult(a SchedulerTaskInfoInterface) 

	// public void executorLost(java.lang.String)
	ExecutorLost(a string) 

	// public void executorAdded(java.lang.String, java.lang.String)
	ExecutorAdded(a string, b string) 

	// public void cancelJob(int)
	CancelJob(a int) 

	// public void cancelJobGroup(java.lang.String)
	CancelJobGroup(a string) 

	// public void cancelAllJobs()
	CancelAllJobs() 

	// public void doCancelAllJobs()
	DoCancelAllJobs() 

	// public void cancelStage(int)
	CancelStage(a int) 

	// public void resubmitFailedStages()
	ResubmitFailedStages() 

	// public void handleJobGroupCancelled(java.lang.String)
	HandleJobGroupCancelled(a string) 

	// public void handleBeginEvent(org.apache.spark.scheduler.Task<?>, org.apache.spark.scheduler.TaskInfo)
	HandleBeginEvent(a SchedulerTaskInterface, b SchedulerTaskInfoInterface) 

	// public void cleanUpAfterSchedulerStop()
	CleanUpAfterSchedulerStop() 

	// public void handleGetTaskResult(org.apache.spark.scheduler.TaskInfo)
	HandleGetTaskResult(a SchedulerTaskInfoInterface) 

	// public void handleMapStageSubmitted(int, org.apache.spark.ShuffleDependency<?, ?, ?>, org.apache.spark.util.CallSite, org.apache.spark.scheduler.JobListener, java.util.Properties)
	HandleMapStageSubmitted(a int, b ShuffleDependencyInterface, c UtilCallSiteInterface, d SchedulerJobListenerInterface, e JavaUtilPropertiesInterface) 

	// public void handleTaskCompletion(org.apache.spark.scheduler.CompletionEvent)
	HandleTaskCompletion(a SchedulerCompletionEventInterface) 

	// public void handleExecutorAdded(java.lang.String, java.lang.String)
	HandleExecutorAdded(a string, b string) 

	// public void handleStageCancellation(int)
	HandleStageCancellation(a int) 

	// public void handleJobCancellation(int, java.lang.String)
	HandleJobCancellation(a int, b string) 

	// public void markMapStageJobAsFinished(org.apache.spark.scheduler.ActiveJob, org.apache.spark.MapOutputStatistics)
	MarkMapStageJobAsFinished(a SchedulerActiveJobInterface, b MapOutputStatisticsInterface) 

	// public void stop()
	Stop() 
}

type SchedulerDAGScheduler struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.DAGScheduler(org.apache.spark.SparkContext, org.apache.spark.scheduler.TaskScheduler, org.apache.spark.scheduler.LiveListenerBus, org.apache.spark.MapOutputTrackerMaster, org.apache.spark.storage.BlockManagerMaster, org.apache.spark.SparkEnv, org.apache.spark.util.Clock)
func NewSchedulerDAGScheduler3(a SparkContextInterface, b SchedulerTaskSchedulerInterface, c SchedulerLiveListenerBusInterface, d MapOutputTrackerMasterInterface, e StorageBlockManagerMasterInterface, f SparkEnvInterface, g UtilClockInterface) (*SchedulerDAGScheduler) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/DAGScheduler", conv_a.Value().Cast("org/apache/spark/SparkContext"), conv_b.Value().Cast("org/apache/spark/scheduler/TaskScheduler"), conv_c.Value().Cast("org/apache/spark/scheduler/LiveListenerBus"), conv_d.Value().Cast("org/apache/spark/MapOutputTrackerMaster"), conv_e.Value().Cast("org/apache/spark/storage/BlockManagerMaster"), conv_f.Value().Cast("org/apache/spark/SparkEnv"), conv_g.Value().Cast("org/apache/spark/util/Clock"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &SchedulerDAGScheduler{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.scheduler.DAGScheduler(org.apache.spark.SparkContext, org.apache.spark.scheduler.TaskScheduler)
func NewSchedulerDAGScheduler2(a SparkContextInterface, b SchedulerTaskSchedulerInterface) (*SchedulerDAGScheduler) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/DAGScheduler", conv_a.Value().Cast("org/apache/spark/SparkContext"), conv_b.Value().Cast("org/apache/spark/scheduler/TaskScheduler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &SchedulerDAGScheduler{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.scheduler.DAGScheduler(org.apache.spark.SparkContext)
func NewSchedulerDAGScheduler(a SparkContextInterface) (*SchedulerDAGScheduler) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/DAGScheduler", conv_a.Value().Cast("org/apache/spark/SparkContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SchedulerDAGScheduler{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static int RESUBMIT_TIMEOUT()
func (jbobject *SchedulerDAGScheduler) RESUBMIT_TIMEOUT() int {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/scheduler/DAGScheduler", "RESUBMIT_TIMEOUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String logName()
func (jbobject *SchedulerDAGScheduler) LogName() string {
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
func (jbobject *SchedulerDAGScheduler) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.SparkContext sc()
func (jbobject *SchedulerDAGScheduler) Sc() *SparkContext {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sc", "org/apache/spark/SparkContext")
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
	unique_x := &SparkContext{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.TaskScheduler taskScheduler()
func (jbobject *SchedulerDAGScheduler) TaskScheduler() *SchedulerTaskScheduler {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskScheduler", "org/apache/spark/scheduler/TaskScheduler")
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
	unique_x := &SchedulerTaskScheduler{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.DAGSchedulerSource metricsSource()
func (jbobject *SchedulerDAGScheduler) MetricsSource() *SchedulerDAGSchedulerSource {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "metricsSource", "org/apache/spark/scheduler/DAGSchedulerSource")
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
	unique_x := &SchedulerDAGSchedulerSource{}
	unique_x.Callable = dst
	return unique_x
}

// public int numTotalJobs()
func (jbobject *SchedulerDAGScheduler) NumTotalJobs() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numTotalJobs", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.scheduler.OutputCommitCoordinator outputCommitCoordinator()
func (jbobject *SchedulerDAGScheduler) OutputCommitCoordinator() *SchedulerOutputCommitCoordinator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "outputCommitCoordinator", "org/apache/spark/scheduler/OutputCommitCoordinator")
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
	unique_x := &SchedulerOutputCommitCoordinator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.DAGSchedulerEventProcessLoop eventProcessLoop()
func (jbobject *SchedulerDAGScheduler) EventProcessLoop() *SchedulerDAGSchedulerEventProcessLoop {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "eventProcessLoop", "org/apache/spark/scheduler/DAGSchedulerEventProcessLoop")
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
	unique_x := &SchedulerDAGSchedulerEventProcessLoop{}
	unique_x.Callable = dst
	return unique_x
}

// public void taskStarted(org.apache.spark.scheduler.Task<?>, org.apache.spark.scheduler.TaskInfo)
func (jbobject *SchedulerDAGScheduler) TaskStarted(a SchedulerTaskInterface, b SchedulerTaskInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "taskStarted", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/Task"), conv_b.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void taskGettingResult(org.apache.spark.scheduler.TaskInfo)
func (jbobject *SchedulerDAGScheduler) TaskGettingResult(a SchedulerTaskInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "taskGettingResult", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void executorLost(java.lang.String)
func (jbobject *SchedulerDAGScheduler) ExecutorLost(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "executorLost", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void executorAdded(java.lang.String, java.lang.String)
func (jbobject *SchedulerDAGScheduler) ExecutorAdded(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "executorAdded", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void cancelJob(int)
func (jbobject *SchedulerDAGScheduler) CancelJob(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelJob", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void cancelJobGroup(java.lang.String)
func (jbobject *SchedulerDAGScheduler) CancelJobGroup(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelJobGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void cancelAllJobs()
func (jbobject *SchedulerDAGScheduler) CancelAllJobs()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelAllJobs", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void doCancelAllJobs()
func (jbobject *SchedulerDAGScheduler) DoCancelAllJobs()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "doCancelAllJobs", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cancelStage(int)
func (jbobject *SchedulerDAGScheduler) CancelStage(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelStage", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void resubmitFailedStages()
func (jbobject *SchedulerDAGScheduler) ResubmitFailedStages()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "resubmitFailedStages", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void handleJobGroupCancelled(java.lang.String)
func (jbobject *SchedulerDAGScheduler) HandleJobGroupCancelled(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleJobGroupCancelled", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void handleBeginEvent(org.apache.spark.scheduler.Task<?>, org.apache.spark.scheduler.TaskInfo)
func (jbobject *SchedulerDAGScheduler) HandleBeginEvent(a SchedulerTaskInterface, b SchedulerTaskInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleBeginEvent", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/Task"), conv_b.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void cleanUpAfterSchedulerStop()
func (jbobject *SchedulerDAGScheduler) CleanUpAfterSchedulerStop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cleanUpAfterSchedulerStop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void handleGetTaskResult(org.apache.spark.scheduler.TaskInfo)
func (jbobject *SchedulerDAGScheduler) HandleGetTaskResult(a SchedulerTaskInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleGetTaskResult", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void handleMapStageSubmitted(int, org.apache.spark.ShuffleDependency<?, ?, ?>, org.apache.spark.util.CallSite, org.apache.spark.scheduler.JobListener, java.util.Properties)
func (jbobject *SchedulerDAGScheduler) HandleMapStageSubmitted(a int, b ShuffleDependencyInterface, c UtilCallSiteInterface, d SchedulerJobListenerInterface, e JavaUtilPropertiesInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleMapStageSubmitted", javabind.Void, a, conv_b.Value().Cast("org/apache/spark/ShuffleDependency"), conv_c.Value().Cast("org/apache/spark/util/CallSite"), conv_d.Value().Cast("org/apache/spark/scheduler/JobListener"), conv_e.Value().Cast("java/util/Properties"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()

}

// public void handleTaskCompletion(org.apache.spark.scheduler.CompletionEvent)
func (jbobject *SchedulerDAGScheduler) HandleTaskCompletion(a SchedulerCompletionEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleTaskCompletion", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/CompletionEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void handleExecutorAdded(java.lang.String, java.lang.String)
func (jbobject *SchedulerDAGScheduler) HandleExecutorAdded(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleExecutorAdded", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void handleStageCancellation(int)
func (jbobject *SchedulerDAGScheduler) HandleStageCancellation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleStageCancellation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void handleJobCancellation(int, java.lang.String)
func (jbobject *SchedulerDAGScheduler) HandleJobCancellation(a int, b string)  {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleJobCancellation", javabind.Void, a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void markMapStageJobAsFinished(org.apache.spark.scheduler.ActiveJob, org.apache.spark.MapOutputStatistics)
func (jbobject *SchedulerDAGScheduler) MarkMapStageJobAsFinished(a SchedulerActiveJobInterface, b MapOutputStatisticsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markMapStageJobAsFinished", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/ActiveJob"), conv_b.Value().Cast("org/apache/spark/MapOutputStatistics"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void stop()
func (jbobject *SchedulerDAGScheduler) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


