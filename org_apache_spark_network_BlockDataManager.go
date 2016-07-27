package gospark

import "github.com/timob/javabind"

type NetworkBlockDataManagerInterface interface {

	// public abstract org.apache.spark.network.buffer.ManagedBuffer getBlockData(org.apache.spark.storage.BlockId)
	GetBlockData(a StorageBlockIdInterface) *NetworkBufferManagedBuffer

	// public abstract void putBlockData(org.apache.spark.storage.BlockId, org.apache.spark.network.buffer.ManagedBuffer, org.apache.spark.storage.StorageLevel)
	PutBlockData(a StorageBlockIdInterface, b NetworkBufferManagedBufferInterface, c StorageStorageLevelInterface) 
}

type NetworkBlockDataManager struct {
	JavaLangObject
}

// public abstract org.apache.spark.network.buffer.ManagedBuffer getBlockData(org.apache.spark.storage.BlockId)
func (jbobject *NetworkBlockDataManager) GetBlockData(a StorageBlockIdInterface) *NetworkBufferManagedBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockData", "org/apache/spark/network/buffer/ManagedBuffer", conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
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

// public abstract void putBlockData(org.apache.spark.storage.BlockId, org.apache.spark.network.buffer.ManagedBuffer, org.apache.spark.storage.StorageLevel)
func (jbobject *NetworkBlockDataManager) PutBlockData(a StorageBlockIdInterface, b NetworkBufferManagedBufferInterface, c StorageStorageLevelInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "putBlockData", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/BlockId"), conv_b.Value().Cast("org/apache/spark/network/buffer/ManagedBuffer"), conv_c.Value().Cast("org/apache/spark/storage/StorageLevel"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()

}


