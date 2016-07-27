package gospark

import "github.com/timob/javabind"

type NetworkBlockTransferServiceInterface interface {
	NetworkShuffleShuffleClientInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public abstract void init(org.apache.spark.network.BlockDataManager)
	Init2(a NetworkBlockDataManagerInterface) 

	// public abstract void close()
	Close() 

	// public abstract int port()
	Port() int

	// public abstract java.lang.String hostName()
	HostName() string

	// public org.apache.spark.network.buffer.ManagedBuffer fetchBlockSync(java.lang.String, int, java.lang.String, java.lang.String)
	FetchBlockSync(a string, b int, c string, d string) *NetworkBufferManagedBuffer

	// public void uploadBlockSync(java.lang.String, int, java.lang.String, org.apache.spark.storage.BlockId, org.apache.spark.network.buffer.ManagedBuffer, org.apache.spark.storage.StorageLevel)
	UploadBlockSync(a string, b int, c string, d StorageBlockIdInterface, e NetworkBufferManagedBufferInterface, f StorageStorageLevelInterface) 
}

type NetworkBlockTransferService struct {
	NetworkShuffleShuffleClient
}

// public org.apache.spark.network.BlockTransferService()
func NewNetworkBlockTransferService() (*NetworkBlockTransferService) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/network/BlockTransferService")
	if err != nil {
		panic(err)
	}
	x := &NetworkBlockTransferService{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *NetworkBlockTransferService) LogName() string {
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
func (jbobject *NetworkBlockTransferService) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract void init(org.apache.spark.network.BlockDataManager)
func (jbobject *NetworkBlockTransferService) Init2(a NetworkBlockDataManagerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "init", javabind.Void, conv_a.Value().Cast("org/apache/spark/network/BlockDataManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void close()
func (jbobject *NetworkBlockTransferService) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract int port()
func (jbobject *NetworkBlockTransferService) Port() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "port", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.String hostName()
func (jbobject *NetworkBlockTransferService) HostName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hostName", "java/lang/String")
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

// public abstract void fetchBlocks(java.lang.String, int, java.lang.String, java.lang.String[], org.apache.spark.network.shuffle.BlockFetchingListener)
func (jbobject *NetworkBlockTransferService) FetchBlocks(a string, b int, c string, d []string, e NetworkShuffleBlockFetchingListenerInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
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
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fetchBlocks", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("org/apache/spark/network/shuffle/BlockFetchingListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()

}

// public org.apache.spark.network.buffer.ManagedBuffer fetchBlockSync(java.lang.String, int, java.lang.String, java.lang.String)
func (jbobject *NetworkBlockTransferService) FetchBlockSync(a string, b int, c string, d string) *NetworkBufferManagedBuffer {
	conv_a := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "fetchBlockSync", "org/apache/spark/network/buffer/ManagedBuffer", conv_a.Value().Cast("java/lang/String"), b, conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
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

// public void uploadBlockSync(java.lang.String, int, java.lang.String, org.apache.spark.storage.BlockId, org.apache.spark.network.buffer.ManagedBuffer, org.apache.spark.storage.StorageLevel)
func (jbobject *NetworkBlockTransferService) UploadBlockSync(a string, b int, c string, d StorageBlockIdInterface, e NetworkBufferManagedBufferInterface, f StorageStorageLevelInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
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
	_, err := jbobject.CallMethod(javabind.GetEnv(), "uploadBlockSync", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("org/apache/spark/storage/BlockId"), conv_e.Value().Cast("org/apache/spark/network/buffer/ManagedBuffer"), conv_f.Value().Cast("org/apache/spark/storage/StorageLevel"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()

}


