package gospark

import "github.com/timob/javabind"

type UtilTaskCompletionListenerInterface interface {

	// public abstract void onTaskCompletion(org.apache.spark.TaskContext)
	OnTaskCompletion(a TaskContextInterface) 
}

type UtilTaskCompletionListener struct {
	JavaUtilEventListener
}

// public abstract void onTaskCompletion(org.apache.spark.TaskContext)
func (jbobject *UtilTaskCompletionListener) OnTaskCompletion(a TaskContextInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onTaskCompletion", javabind.Void, conv_a.Value().Cast("org/apache/spark/TaskContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


