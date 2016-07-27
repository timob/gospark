package gospark

import "github.com/timob/javabind"

type UiJobsJobProgressListenerInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public void onUnpersistRDD(org.apache.spark.scheduler.SparkListenerUnpersistRDD)
	OnUnpersistRDD(a SchedulerSparkListenerUnpersistRDDInterface) 

	// public void onExecutorAdded(org.apache.spark.scheduler.SparkListenerExecutorAdded)
	OnExecutorAdded(a SchedulerSparkListenerExecutorAddedInterface) 

	// public void onExecutorRemoved(org.apache.spark.scheduler.SparkListenerExecutorRemoved)
	OnExecutorRemoved(a SchedulerSparkListenerExecutorRemovedInterface) 

	// public void onBlockUpdated(org.apache.spark.scheduler.SparkListenerBlockUpdated)
	OnBlockUpdated(a SchedulerSparkListenerBlockUpdatedInterface) 

	// public long startTime()
	StartTime() int64

	// public long endTime()
	EndTime() int64

	// public int numCompletedStages()
	NumCompletedStages() int

	// public int numFailedStages()
	NumFailedStages() int

	// public int numCompletedJobs()
	NumCompletedJobs() int

	// public int numFailedJobs()
	NumFailedJobs() int

	// public int retainedStages()
	RetainedStages() int

	// public int retainedJobs()
	RetainedJobs() int

	// public synchronized void onJobStart(org.apache.spark.scheduler.SparkListenerJobStart)
	OnJobStart(a SchedulerSparkListenerJobStartInterface) 

	// public synchronized void onJobEnd(org.apache.spark.scheduler.SparkListenerJobEnd)
	OnJobEnd(a SchedulerSparkListenerJobEndInterface) 

	// public synchronized void onStageCompleted(org.apache.spark.scheduler.SparkListenerStageCompleted)
	OnStageCompleted(a SchedulerSparkListenerStageCompletedInterface) 

	// public synchronized void onStageSubmitted(org.apache.spark.scheduler.SparkListenerStageSubmitted)
	OnStageSubmitted(a SchedulerSparkListenerStageSubmittedInterface) 

	// public synchronized void onTaskStart(org.apache.spark.scheduler.SparkListenerTaskStart)
	OnTaskStart(a SchedulerSparkListenerTaskStartInterface) 

	// public void onTaskGettingResult(org.apache.spark.scheduler.SparkListenerTaskGettingResult)
	OnTaskGettingResult(a SchedulerSparkListenerTaskGettingResultInterface) 

	// public synchronized void onTaskEnd(org.apache.spark.scheduler.SparkListenerTaskEnd)
	OnTaskEnd(a SchedulerSparkListenerTaskEndInterface) 

	// public void onExecutorMetricsUpdate(org.apache.spark.scheduler.SparkListenerExecutorMetricsUpdate)
	OnExecutorMetricsUpdate(a SchedulerSparkListenerExecutorMetricsUpdateInterface) 

	// public synchronized void onEnvironmentUpdate(org.apache.spark.scheduler.SparkListenerEnvironmentUpdate)
	OnEnvironmentUpdate(a SchedulerSparkListenerEnvironmentUpdateInterface) 

	// public synchronized void onBlockManagerAdded(org.apache.spark.scheduler.SparkListenerBlockManagerAdded)
	OnBlockManagerAdded(a SchedulerSparkListenerBlockManagerAddedInterface) 

	// public void onBlockManagerRemoved(org.apache.spark.scheduler.SparkListenerBlockManagerRemoved)
	OnBlockManagerRemoved(a SchedulerSparkListenerBlockManagerRemovedInterface) 

	// public void onApplicationStart(org.apache.spark.scheduler.SparkListenerApplicationStart)
	OnApplicationStart(a SchedulerSparkListenerApplicationStartInterface) 

	// public void onApplicationEnd(org.apache.spark.scheduler.SparkListenerApplicationEnd)
	OnApplicationEnd(a SchedulerSparkListenerApplicationEndInterface) 

	// public void waitUntilExecutorsUp(int, long)
	WaitUntilExecutorsUp(a int, b int64) 
}

type UiJobsJobProgressListener struct {
	JavaLangObject
}

// public org.apache.spark.ui.jobs.JobProgressListener(org.apache.spark.SparkConf)
func NewUiJobsJobProgressListener(a SparkConfInterface) (*UiJobsJobProgressListener) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/ui/jobs/JobProgressListener", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &UiJobsJobProgressListener{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *UiJobsJobProgressListener) LogName() string {
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
func (jbobject *UiJobsJobProgressListener) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void onUnpersistRDD(org.apache.spark.scheduler.SparkListenerUnpersistRDD)
func (jbobject *UiJobsJobProgressListener) OnUnpersistRDD(a SchedulerSparkListenerUnpersistRDDInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onUnpersistRDD", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerUnpersistRDD"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onExecutorAdded(org.apache.spark.scheduler.SparkListenerExecutorAdded)
func (jbobject *UiJobsJobProgressListener) OnExecutorAdded(a SchedulerSparkListenerExecutorAddedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onExecutorAdded", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerExecutorAdded"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onExecutorRemoved(org.apache.spark.scheduler.SparkListenerExecutorRemoved)
func (jbobject *UiJobsJobProgressListener) OnExecutorRemoved(a SchedulerSparkListenerExecutorRemovedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onExecutorRemoved", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerExecutorRemoved"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onBlockUpdated(org.apache.spark.scheduler.SparkListenerBlockUpdated)
func (jbobject *UiJobsJobProgressListener) OnBlockUpdated(a SchedulerSparkListenerBlockUpdatedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onBlockUpdated", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerBlockUpdated"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public long startTime()
func (jbobject *UiJobsJobProgressListener) StartTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "startTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long endTime()
func (jbobject *UiJobsJobProgressListener) EndTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "endTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public int numCompletedStages()
func (jbobject *UiJobsJobProgressListener) NumCompletedStages() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numCompletedStages", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int numFailedStages()
func (jbobject *UiJobsJobProgressListener) NumFailedStages() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numFailedStages", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int numCompletedJobs()
func (jbobject *UiJobsJobProgressListener) NumCompletedJobs() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numCompletedJobs", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int numFailedJobs()
func (jbobject *UiJobsJobProgressListener) NumFailedJobs() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numFailedJobs", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int retainedStages()
func (jbobject *UiJobsJobProgressListener) RetainedStages() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "retainedStages", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int retainedJobs()
func (jbobject *UiJobsJobProgressListener) RetainedJobs() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "retainedJobs", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public synchronized void onJobStart(org.apache.spark.scheduler.SparkListenerJobStart)
func (jbobject *UiJobsJobProgressListener) OnJobStart(a SchedulerSparkListenerJobStartInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onJobStart", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerJobStart"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onJobEnd(org.apache.spark.scheduler.SparkListenerJobEnd)
func (jbobject *UiJobsJobProgressListener) OnJobEnd(a SchedulerSparkListenerJobEndInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onJobEnd", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerJobEnd"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onStageCompleted(org.apache.spark.scheduler.SparkListenerStageCompleted)
func (jbobject *UiJobsJobProgressListener) OnStageCompleted(a SchedulerSparkListenerStageCompletedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStageCompleted", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerStageCompleted"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onStageSubmitted(org.apache.spark.scheduler.SparkListenerStageSubmitted)
func (jbobject *UiJobsJobProgressListener) OnStageSubmitted(a SchedulerSparkListenerStageSubmittedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStageSubmitted", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerStageSubmitted"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onTaskStart(org.apache.spark.scheduler.SparkListenerTaskStart)
func (jbobject *UiJobsJobProgressListener) OnTaskStart(a SchedulerSparkListenerTaskStartInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onTaskStart", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerTaskStart"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onTaskGettingResult(org.apache.spark.scheduler.SparkListenerTaskGettingResult)
func (jbobject *UiJobsJobProgressListener) OnTaskGettingResult(a SchedulerSparkListenerTaskGettingResultInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onTaskGettingResult", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerTaskGettingResult"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onTaskEnd(org.apache.spark.scheduler.SparkListenerTaskEnd)
func (jbobject *UiJobsJobProgressListener) OnTaskEnd(a SchedulerSparkListenerTaskEndInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onTaskEnd", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerTaskEnd"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onExecutorMetricsUpdate(org.apache.spark.scheduler.SparkListenerExecutorMetricsUpdate)
func (jbobject *UiJobsJobProgressListener) OnExecutorMetricsUpdate(a SchedulerSparkListenerExecutorMetricsUpdateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onExecutorMetricsUpdate", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerExecutorMetricsUpdate"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onEnvironmentUpdate(org.apache.spark.scheduler.SparkListenerEnvironmentUpdate)
func (jbobject *UiJobsJobProgressListener) OnEnvironmentUpdate(a SchedulerSparkListenerEnvironmentUpdateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onEnvironmentUpdate", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerEnvironmentUpdate"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void onBlockManagerAdded(org.apache.spark.scheduler.SparkListenerBlockManagerAdded)
func (jbobject *UiJobsJobProgressListener) OnBlockManagerAdded(a SchedulerSparkListenerBlockManagerAddedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onBlockManagerAdded", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerBlockManagerAdded"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onBlockManagerRemoved(org.apache.spark.scheduler.SparkListenerBlockManagerRemoved)
func (jbobject *UiJobsJobProgressListener) OnBlockManagerRemoved(a SchedulerSparkListenerBlockManagerRemovedInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onBlockManagerRemoved", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerBlockManagerRemoved"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onApplicationStart(org.apache.spark.scheduler.SparkListenerApplicationStart)
func (jbobject *UiJobsJobProgressListener) OnApplicationStart(a SchedulerSparkListenerApplicationStartInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onApplicationStart", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerApplicationStart"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onApplicationEnd(org.apache.spark.scheduler.SparkListenerApplicationEnd)
func (jbobject *UiJobsJobProgressListener) OnApplicationEnd(a SchedulerSparkListenerApplicationEndInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onApplicationEnd", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerApplicationEnd"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void waitUntilExecutorsUp(int, long)
func (jbobject *UiJobsJobProgressListener) WaitUntilExecutorsUp(a int, b int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "waitUntilExecutorsUp", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}


