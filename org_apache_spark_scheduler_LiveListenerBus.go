package gospark

import "github.com/timob/javabind"

type SchedulerLiveListenerBusInterface interface {
	UtilAsynchronousListenerBusInterface

	// public void onPostEvent(org.apache.spark.scheduler.SparkListener, org.apache.spark.scheduler.SparkListenerEvent)
	OnPostEvent2(a SchedulerSparkListenerInterface, b SchedulerSparkListenerEventInterface) 

	// public void onDropEvent(org.apache.spark.scheduler.SparkListenerEvent)
	OnDropEvent3(a SchedulerSparkListenerEventInterface) 

	// public void onPostEvent(java.lang.Object, java.lang.Object)
	OnPostEvent(a interface{}, b interface{}) 

	// public void onDropEvent(java.lang.Object)
	OnDropEvent2(a interface{}) 
}

type SchedulerLiveListenerBus struct {
	UtilAsynchronousListenerBus
}

// public org.apache.spark.scheduler.LiveListenerBus()
func NewSchedulerLiveListenerBus() (*SchedulerLiveListenerBus) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/LiveListenerBus")
	if err != nil {
		panic(err)
	}
	x := &SchedulerLiveListenerBus{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void onPostEvent(org.apache.spark.scheduler.SparkListener, org.apache.spark.scheduler.SparkListenerEvent)
func (jbobject *SchedulerLiveListenerBus) OnPostEvent2(a SchedulerSparkListenerInterface, b SchedulerSparkListenerEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onPostEvent", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListener"), conv_b.Value().Cast("org/apache/spark/scheduler/SparkListenerEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void onDropEvent(org.apache.spark.scheduler.SparkListenerEvent)
func (jbobject *SchedulerLiveListenerBus) OnDropEvent3(a SchedulerSparkListenerEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onDropEvent", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListenerEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void onPostEvent(java.lang.Object, java.lang.Object)
func (jbobject *SchedulerLiveListenerBus) OnPostEvent(a interface{}, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onPostEvent", javabind.Void, conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void onDropEvent(java.lang.Object)
func (jbobject *SchedulerLiveListenerBus) OnDropEvent2(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onDropEvent", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


