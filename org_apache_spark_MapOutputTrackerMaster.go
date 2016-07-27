package gospark

import "github.com/timob/javabind"

type MapOutputTrackerMasterInterface interface {
	MapOutputTrackerInterface

	// public org.apache.spark.util.TimeStampedHashMap<java.lang.Object, org.apache.spark.scheduler.MapStatus[]> mapStatuses()
	MapStatuses() *UtilTimeStampedHashMap

	// public void registerShuffle(int, int)
	RegisterShuffle(a int, b int) 

	// public void registerMapOutput(int, int, org.apache.spark.scheduler.MapStatus)
	RegisterMapOutput(a int, b int, c SchedulerMapStatusInterface) 

	// public void registerMapOutputs(int, org.apache.spark.scheduler.MapStatus[], boolean)
	RegisterMapOutputs(a int, b []*SchedulerMapStatus, c bool) 

	// public void unregisterMapOutput(int, int, org.apache.spark.storage.BlockManagerId)
	UnregisterMapOutput(a int, b int, c StorageBlockManagerIdInterface) 

	// public boolean containsShuffle(int)
	ContainsShuffle(a int) bool

	// public void incrementEpoch()
	IncrementEpoch() 

	// public byte[] getSerializedMapOutputStatuses(int)
	GetSerializedMapOutputStatuses(a int) []byte
}

type MapOutputTrackerMaster struct {
	MapOutputTracker
}

// public org.apache.spark.MapOutputTrackerMaster(org.apache.spark.SparkConf)
func NewMapOutputTrackerMaster(a SparkConfInterface) (*MapOutputTrackerMaster) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/MapOutputTrackerMaster", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &MapOutputTrackerMaster{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.util.TimeStampedHashMap<java.lang.Object, org.apache.spark.scheduler.MapStatus[]> mapStatuses()
func (jbobject *MapOutputTrackerMaster) MapStatuses() *UtilTimeStampedHashMap {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapStatuses", "org/apache/spark/util/TimeStampedHashMap")
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
	unique_x := &UtilTimeStampedHashMap{}
	unique_x.Callable = dst
	return unique_x
}

// public void registerShuffle(int, int)
func (jbobject *MapOutputTrackerMaster) RegisterShuffle(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "registerShuffle", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void registerMapOutput(int, int, org.apache.spark.scheduler.MapStatus)
func (jbobject *MapOutputTrackerMaster) RegisterMapOutput(a int, b int, c SchedulerMapStatusInterface)  {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "registerMapOutput", javabind.Void, a, b, conv_c.Value().Cast("org/apache/spark/scheduler/MapStatus"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void registerMapOutputs(int, org.apache.spark.scheduler.MapStatus[], boolean)
func (jbobject *MapOutputTrackerMaster) RegisterMapOutputs(a int, b []*SchedulerMapStatus, c bool)  {
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/apache/spark/scheduler/MapStatus")
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "registerMapOutputs", javabind.Void, a, conv_b.Value().Cast("org/apache/spark/scheduler/MapStatus"), c)
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void unregisterMapOutput(int, int, org.apache.spark.storage.BlockManagerId)
func (jbobject *MapOutputTrackerMaster) UnregisterMapOutput(a int, b int, c StorageBlockManagerIdInterface)  {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unregisterMapOutput", javabind.Void, a, b, conv_c.Value().Cast("org/apache/spark/storage/BlockManagerId"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void unregisterShuffle(int)
func (jbobject *MapOutputTrackerMaster) UnregisterShuffle(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unregisterShuffle", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean containsShuffle(int)
func (jbobject *MapOutputTrackerMaster) ContainsShuffle(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "containsShuffle", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void incrementEpoch()
func (jbobject *MapOutputTrackerMaster) IncrementEpoch()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incrementEpoch", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public byte[] getSerializedMapOutputStatuses(int)
func (jbobject *MapOutputTrackerMaster) GetSerializedMapOutputStatuses(a int) []byte {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSerializedMapOutputStatuses", javabind.Byte | javabind.Array, a)
	if err != nil {
		panic(err)
	}
	return jret.([]byte)
}

// public void stop()
func (jbobject *MapOutputTrackerMaster) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


