package gospark

import "github.com/timob/javabind"

type SchedulerDAGSchedulerEventProcessLoopInterface interface {
	UtilEventLoopInterface

	// public void onReceive(org.apache.spark.scheduler.DAGSchedulerEvent)
	OnReceive3(a SchedulerDAGSchedulerEventInterface) 

	// public void onReceive(java.lang.Object)
	OnReceive2(a interface{}) 
}

type SchedulerDAGSchedulerEventProcessLoop struct {
	UtilEventLoop
}

// public org.apache.spark.scheduler.DAGSchedulerEventProcessLoop(org.apache.spark.scheduler.DAGScheduler)
func NewSchedulerDAGSchedulerEventProcessLoop(a SchedulerDAGSchedulerInterface) (*SchedulerDAGSchedulerEventProcessLoop) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/DAGSchedulerEventProcessLoop", conv_a.Value().Cast("org/apache/spark/scheduler/DAGScheduler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SchedulerDAGSchedulerEventProcessLoop{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void onReceive(org.apache.spark.scheduler.DAGSchedulerEvent)
func (jbobject *SchedulerDAGSchedulerEventProcessLoop) OnReceive3(a SchedulerDAGSchedulerEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onReceive", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/DAGSchedulerEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onStop()
func (jbobject *SchedulerDAGSchedulerEventProcessLoop) OnStop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onStop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void onReceive(java.lang.Object)
func (jbobject *SchedulerDAGSchedulerEventProcessLoop) OnReceive2(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onReceive", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


