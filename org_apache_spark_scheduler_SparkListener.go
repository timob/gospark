package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerInterface interface {

	// public abstract void onStageCompleted(org.apache.spark.scheduler.SparkListenerStageCompleted)
	OnStageCompleted(a SchedulerSparkListenerStageCompletedInterface) 

	// public abstract void onStageSubmitted(org.apache.spark.scheduler.SparkListenerStageSubmitted)
	OnStageSubmitted(a SchedulerSparkListenerStageSubmittedInterface) 

	// public abstract void onTaskStart(org.apache.spark.scheduler.SparkListenerTaskStart)
	OnTaskStart(a SchedulerSparkListenerTaskStartInterface) 

	// public abstract void onTaskGettingResult(org.apache.spark.scheduler.SparkListenerTaskGettingResult)
	OnTaskGettingResult(a SchedulerSparkListenerTaskGettingResultInterface) 

	// public abstract void onTaskEnd(org.apache.spark.scheduler.SparkListenerTaskEnd)
	OnTaskEnd(a SchedulerSparkListenerTaskEndInterface) 

	// public abstract void onJobStart(org.apache.spark.scheduler.SparkListenerJobStart)
	OnJobStart(a SchedulerSparkListenerJobStartInterface) 

	// public abstract void onJobEnd(org.apache.spark.scheduler.SparkListenerJobEnd)
	OnJobEnd(a SchedulerSparkListenerJobEndInterface) 

	// public abstract void onEnvironmentUpdate(org.apache.spark.scheduler.SparkListenerEnvironmentUpdate)
	OnEnvironmentUpdate(a SchedulerSparkListenerEnvironmentUpdateInterface) 

	// public abstract void onBlockManagerAdded(org.apache.spark.scheduler.SparkListenerBlockManagerAdded)
	OnBlockManagerAdded(a SchedulerSparkListenerBlockManagerAddedInterface) 

	// public abstract void onBlockManagerRemoved(org.apache.spark.scheduler.SparkListenerBlockManagerRemoved)
	OnBlockManagerRemoved(a SchedulerSparkListenerBlockManagerRemovedInterface) 

	// public abstract void onUnpersistRDD(org.apache.spark.scheduler.SparkListenerUnpersistRDD)
	OnUnpersistRDD(a SchedulerSparkListenerUnpersistRDDInterface) 

	// public abstract void onApplicationStart(org.apache.spark.scheduler.SparkListenerApplicationStart)
	OnApplicationStart(a SchedulerSparkListenerApplicationStartInterface) 

	// public abstract void onApplicationEnd(org.apache.spark.scheduler.SparkListenerApplicationEnd)
	OnApplicationEnd(a SchedulerSparkListenerApplicationEndInterface) 

	// public abstract void onExecutorMetricsUpdate(org.apache.spark.scheduler.SparkListenerExecutorMetricsUpdate)
	OnExecutorMetricsUpdate(a SchedulerSparkListenerExecutorMetricsUpdateInterface) 

	// public abstract void onExecutorAdded(org.apache.spark.scheduler.SparkListenerExecutorAdded)
	OnExecutorAdded(a SchedulerSparkListenerExecutorAddedInterface) 

	// public abstract void onExecutorRemoved(org.apache.spark.scheduler.SparkListenerExecutorRemoved)
	OnExecutorRemoved(a SchedulerSparkListenerExecutorRemovedInterface) 

	// public abstract void onBlockUpdated(org.apache.spark.scheduler.SparkListenerBlockUpdated)
	OnBlockUpdated(a SchedulerSparkListenerBlockUpdatedInterface) 
}

type SchedulerSparkListener struct {
	JavaLangObject
}

// public abstract void onStageCompleted(org.apache.spark.scheduler.SparkListenerStageCompleted)
func (jbobject *SchedulerSparkListener) OnStageCompleted(a SchedulerSparkListenerStageCompletedInterface)  {
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

// public abstract void onStageSubmitted(org.apache.spark.scheduler.SparkListenerStageSubmitted)
func (jbobject *SchedulerSparkListener) OnStageSubmitted(a SchedulerSparkListenerStageSubmittedInterface)  {
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

// public abstract void onTaskStart(org.apache.spark.scheduler.SparkListenerTaskStart)
func (jbobject *SchedulerSparkListener) OnTaskStart(a SchedulerSparkListenerTaskStartInterface)  {
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

// public abstract void onTaskGettingResult(org.apache.spark.scheduler.SparkListenerTaskGettingResult)
func (jbobject *SchedulerSparkListener) OnTaskGettingResult(a SchedulerSparkListenerTaskGettingResultInterface)  {
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

// public abstract void onTaskEnd(org.apache.spark.scheduler.SparkListenerTaskEnd)
func (jbobject *SchedulerSparkListener) OnTaskEnd(a SchedulerSparkListenerTaskEndInterface)  {
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

// public abstract void onJobStart(org.apache.spark.scheduler.SparkListenerJobStart)
func (jbobject *SchedulerSparkListener) OnJobStart(a SchedulerSparkListenerJobStartInterface)  {
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

// public abstract void onJobEnd(org.apache.spark.scheduler.SparkListenerJobEnd)
func (jbobject *SchedulerSparkListener) OnJobEnd(a SchedulerSparkListenerJobEndInterface)  {
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

// public abstract void onEnvironmentUpdate(org.apache.spark.scheduler.SparkListenerEnvironmentUpdate)
func (jbobject *SchedulerSparkListener) OnEnvironmentUpdate(a SchedulerSparkListenerEnvironmentUpdateInterface)  {
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

// public abstract void onBlockManagerAdded(org.apache.spark.scheduler.SparkListenerBlockManagerAdded)
func (jbobject *SchedulerSparkListener) OnBlockManagerAdded(a SchedulerSparkListenerBlockManagerAddedInterface)  {
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

// public abstract void onBlockManagerRemoved(org.apache.spark.scheduler.SparkListenerBlockManagerRemoved)
func (jbobject *SchedulerSparkListener) OnBlockManagerRemoved(a SchedulerSparkListenerBlockManagerRemovedInterface)  {
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

// public abstract void onUnpersistRDD(org.apache.spark.scheduler.SparkListenerUnpersistRDD)
func (jbobject *SchedulerSparkListener) OnUnpersistRDD(a SchedulerSparkListenerUnpersistRDDInterface)  {
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

// public abstract void onApplicationStart(org.apache.spark.scheduler.SparkListenerApplicationStart)
func (jbobject *SchedulerSparkListener) OnApplicationStart(a SchedulerSparkListenerApplicationStartInterface)  {
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

// public abstract void onApplicationEnd(org.apache.spark.scheduler.SparkListenerApplicationEnd)
func (jbobject *SchedulerSparkListener) OnApplicationEnd(a SchedulerSparkListenerApplicationEndInterface)  {
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

// public abstract void onExecutorMetricsUpdate(org.apache.spark.scheduler.SparkListenerExecutorMetricsUpdate)
func (jbobject *SchedulerSparkListener) OnExecutorMetricsUpdate(a SchedulerSparkListenerExecutorMetricsUpdateInterface)  {
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

// public abstract void onExecutorAdded(org.apache.spark.scheduler.SparkListenerExecutorAdded)
func (jbobject *SchedulerSparkListener) OnExecutorAdded(a SchedulerSparkListenerExecutorAddedInterface)  {
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

// public abstract void onExecutorRemoved(org.apache.spark.scheduler.SparkListenerExecutorRemoved)
func (jbobject *SchedulerSparkListener) OnExecutorRemoved(a SchedulerSparkListenerExecutorRemovedInterface)  {
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

// public abstract void onBlockUpdated(org.apache.spark.scheduler.SparkListenerBlockUpdated)
func (jbobject *SchedulerSparkListener) OnBlockUpdated(a SchedulerSparkListenerBlockUpdatedInterface)  {
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


