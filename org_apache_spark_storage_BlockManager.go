package gospark

import "github.com/timob/javabind"

type StorageBlockManagerInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.storage.BlockManagerMaster master()
	Master() *StorageBlockManagerMaster

	// public org.apache.spark.SparkConf conf()
	Conf() *SparkConf

	// public org.apache.spark.memory.MemoryManager memoryManager()
	MemoryManager() *MemoryMemoryManager

	// public org.apache.spark.storage.DiskBlockManager diskBlockManager()
	DiskBlockManager() *StorageDiskBlockManager

	// public org.apache.spark.storage.MemoryStore memoryStore()
	MemoryStore() *StorageMemoryStore

	// public org.apache.spark.storage.DiskStore diskStore()
	DiskStore() *StorageDiskStore

	// public org.apache.spark.storage.ExternalBlockStore externalBlockStore()
	ExternalBlockStore() *StorageExternalBlockStore

	// public boolean externalShuffleServiceEnabled()
	ExternalShuffleServiceEnabled() bool

	// public org.apache.spark.storage.BlockManagerId blockManagerId()
	BlockManagerId() *StorageBlockManagerId

	// public org.apache.spark.storage.BlockManagerId shuffleServerId()
	ShuffleServerId() *StorageBlockManagerId

	// public org.apache.spark.network.shuffle.ShuffleClient shuffleClient()
	ShuffleClient() *NetworkShuffleShuffleClient

	// public void initialize(java.lang.String)
	Initialize(a string) 

	// public void reregister()
	Reregister() 

	// public void waitForAsyncReregister()
	WaitForAsyncReregister() 

	// public org.apache.spark.network.buffer.ManagedBuffer getBlockData(org.apache.spark.storage.BlockId)
	GetBlockData(a StorageBlockIdInterface) *NetworkBufferManagedBuffer

	// public void putBlockData(org.apache.spark.storage.BlockId, org.apache.spark.network.buffer.ManagedBuffer, org.apache.spark.storage.StorageLevel)
	PutBlockData(a StorageBlockIdInterface, b NetworkBufferManagedBufferInterface, c StorageStorageLevelInterface) 

	// public int removeRdd(int)
	RemoveRdd(a int) int

	// public int removeBroadcast(long, boolean)
	RemoveBroadcast(a int64, b bool) int

	// public void removeBlock(org.apache.spark.storage.BlockId, boolean)
	RemoveBlock(a StorageBlockIdInterface, b bool) 

	// public org.apache.spark.storage.BlockInfo getBlockInfo(org.apache.spark.storage.BlockId)
	GetBlockInfo(a StorageBlockIdInterface) *StorageBlockInfo

	// public void stop()
	Stop() 
}

type StorageBlockManager struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockManager(java.lang.String, org.apache.spark.rpc.RpcEnv, org.apache.spark.storage.BlockManagerMaster, org.apache.spark.serializer.Serializer, org.apache.spark.SparkConf, org.apache.spark.memory.MemoryManager, org.apache.spark.MapOutputTracker, org.apache.spark.shuffle.ShuffleManager, org.apache.spark.network.BlockTransferService, org.apache.spark.SecurityManager, int)
func NewStorageBlockManager(a string, b RpcRpcEnvInterface, c StorageBlockManagerMasterInterface, d SerializerSerializerInterface, e SparkConfInterface, f MemoryMemoryManagerInterface, g MapOutputTrackerInterface, h ShuffleShuffleManagerInterface, i NetworkBlockTransferServiceInterface, j SecurityManagerInterface, k int) (*StorageBlockManager) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaCallable()
	conv_i := javabind.NewGoToJavaCallable()
	conv_j := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}
	if err := conv_i.Convert(i); err != nil {
		panic(err)
	}
	if err := conv_j.Convert(j); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockManager", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/rpc/RpcEnv"), conv_c.Value().Cast("org/apache/spark/storage/BlockManagerMaster"), conv_d.Value().Cast("org/apache/spark/serializer/Serializer"), conv_e.Value().Cast("org/apache/spark/SparkConf"), conv_f.Value().Cast("org/apache/spark/memory/MemoryManager"), conv_g.Value().Cast("org/apache/spark/MapOutputTracker"), conv_h.Value().Cast("org/apache/spark/shuffle/ShuffleManager"), conv_i.Value().Cast("org/apache/spark/network/BlockTransferService"), conv_j.Value().Cast("org/apache/spark/SecurityManager"), k)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	conv_j.CleanUp()
	x := &StorageBlockManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *StorageBlockManager) LogName() string {
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
func (jbobject *StorageBlockManager) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.storage.BlockManagerMaster master()
func (jbobject *StorageBlockManager) Master() *StorageBlockManagerMaster {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "master", "org/apache/spark/storage/BlockManagerMaster")
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
	unique_x := &StorageBlockManagerMaster{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf conf()
func (jbobject *StorageBlockManager) Conf() *SparkConf {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "conf", "org/apache/spark/SparkConf")
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.memory.MemoryManager memoryManager()
func (jbobject *StorageBlockManager) MemoryManager() *MemoryMemoryManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryManager", "org/apache/spark/memory/MemoryManager")
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
	unique_x := &MemoryMemoryManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.DiskBlockManager diskBlockManager()
func (jbobject *StorageBlockManager) DiskBlockManager() *StorageDiskBlockManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskBlockManager", "org/apache/spark/storage/DiskBlockManager")
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
	unique_x := &StorageDiskBlockManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.MemoryStore memoryStore()
func (jbobject *StorageBlockManager) MemoryStore() *StorageMemoryStore {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryStore", "org/apache/spark/storage/MemoryStore")
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
	unique_x := &StorageMemoryStore{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.DiskStore diskStore()
func (jbobject *StorageBlockManager) DiskStore() *StorageDiskStore {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskStore", "org/apache/spark/storage/DiskStore")
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
	unique_x := &StorageDiskStore{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.ExternalBlockStore externalBlockStore()
func (jbobject *StorageBlockManager) ExternalBlockStore() *StorageExternalBlockStore {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "externalBlockStore", "org/apache/spark/storage/ExternalBlockStore")
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
	unique_x := &StorageExternalBlockStore{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean externalShuffleServiceEnabled()
func (jbobject *StorageBlockManager) ExternalShuffleServiceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "externalShuffleServiceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.storage.BlockManagerId blockManagerId()
func (jbobject *StorageBlockManager) BlockManagerId() *StorageBlockManagerId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "blockManagerId", "org/apache/spark/storage/BlockManagerId")
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
	unique_x := &StorageBlockManagerId{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.BlockManagerId shuffleServerId()
func (jbobject *StorageBlockManager) ShuffleServerId() *StorageBlockManagerId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleServerId", "org/apache/spark/storage/BlockManagerId")
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
	unique_x := &StorageBlockManagerId{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.network.shuffle.ShuffleClient shuffleClient()
func (jbobject *StorageBlockManager) ShuffleClient() *NetworkShuffleShuffleClient {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleClient", "org/apache/spark/network/shuffle/ShuffleClient")
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
	unique_x := &NetworkShuffleShuffleClient{}
	unique_x.Callable = dst
	return unique_x
}

// public void initialize(java.lang.String)
func (jbobject *StorageBlockManager) Initialize(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "initialize", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void reregister()
func (jbobject *StorageBlockManager) Reregister()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reregister", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void waitForAsyncReregister()
func (jbobject *StorageBlockManager) WaitForAsyncReregister()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "waitForAsyncReregister", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.apache.spark.network.buffer.ManagedBuffer getBlockData(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockManager) GetBlockData(a StorageBlockIdInterface) *NetworkBufferManagedBuffer {
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

// public void putBlockData(org.apache.spark.storage.BlockId, org.apache.spark.network.buffer.ManagedBuffer, org.apache.spark.storage.StorageLevel)
func (jbobject *StorageBlockManager) PutBlockData(a StorageBlockIdInterface, b NetworkBufferManagedBufferInterface, c StorageStorageLevelInterface)  {
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

// public int removeRdd(int)
func (jbobject *StorageBlockManager) RemoveRdd(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "removeRdd", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int removeBroadcast(long, boolean)
func (jbobject *StorageBlockManager) RemoveBroadcast(a int64, b bool) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "removeBroadcast", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void removeBlock(org.apache.spark.storage.BlockId, boolean)
func (jbobject *StorageBlockManager) RemoveBlock(a StorageBlockIdInterface, b bool)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeBlock", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/BlockId"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.apache.spark.storage.BlockInfo getBlockInfo(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockManager) GetBlockInfo(a StorageBlockIdInterface) *StorageBlockInfo {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockInfo", "org/apache/spark/storage/BlockInfo", conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
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
	unique_x := &StorageBlockInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public void stop()
func (jbobject *StorageBlockManager) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


