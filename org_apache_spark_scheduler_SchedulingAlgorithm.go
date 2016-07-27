package gospark

import "github.com/timob/javabind"

type SchedulerSchedulingAlgorithmInterface interface {

	// public abstract boolean comparator(org.apache.spark.scheduler.Schedulable, org.apache.spark.scheduler.Schedulable)
	Comparator(a SchedulerSchedulableInterface, b SchedulerSchedulableInterface) bool
}

type SchedulerSchedulingAlgorithm struct {
	JavaLangObject
}

// public abstract boolean comparator(org.apache.spark.scheduler.Schedulable, org.apache.spark.scheduler.Schedulable)
func (jbobject *SchedulerSchedulingAlgorithm) Comparator(a SchedulerSchedulableInterface, b SchedulerSchedulableInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "comparator", javabind.Boolean, conv_a.Value().Cast("org/apache/spark/scheduler/Schedulable"), conv_b.Value().Cast("org/apache/spark/scheduler/Schedulable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}


