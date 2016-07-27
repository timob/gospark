package gospark

import "github.com/timob/javabind"

type ShuffleShuffleBlockResolverInterface interface {

	// public abstract org.apache.spark.network.buffer.ManagedBuffer getBlockData(org.apache.spark.storage.ShuffleBlockId)
	GetBlockData(a StorageShuffleBlockIdInterface) *NetworkBufferManagedBuffer

	// public abstract void stop()
	Stop() 
}

type ShuffleShuffleBlockResolver struct {
	JavaLangObject
}

// public abstract org.apache.spark.network.buffer.ManagedBuffer getBlockData(org.apache.spark.storage.ShuffleBlockId)
func (jbobject *ShuffleShuffleBlockResolver) GetBlockData(a StorageShuffleBlockIdInterface) *NetworkBufferManagedBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockData", "org/apache/spark/network/buffer/ManagedBuffer", conv_a.Value().Cast("org/apache/spark/storage/ShuffleBlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &NetworkBufferManagedBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void stop()
func (jbobject *ShuffleShuffleBlockResolver) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


