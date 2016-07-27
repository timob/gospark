package gospark

import "github.com/timob/javabind"

type NetworkShuffleBlockFetchingListenerInterface interface {

	// public abstract void onBlockFetchSuccess(java.lang.String, org.apache.spark.network.buffer.ManagedBuffer)
	OnBlockFetchSuccess(a string, b NetworkBufferManagedBufferInterface) 
}

type NetworkShuffleBlockFetchingListener struct {
	JavaUtilEventListener
}

// public abstract void onBlockFetchSuccess(java.lang.String, org.apache.spark.network.buffer.ManagedBuffer)
func (jbobject *NetworkShuffleBlockFetchingListener) OnBlockFetchSuccess(a string, b NetworkBufferManagedBufferInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "onBlockFetchSuccess", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/network/buffer/ManagedBuffer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}


