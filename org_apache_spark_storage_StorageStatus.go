package gospark

import "github.com/timob/javabind"

type StorageStorageStatusInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.storage.BlockManagerId blockManagerId()
	BlockManagerId() *StorageBlockManagerId

	// public long maxMem()
	MaxMem() int64

	// public void addBlock(org.apache.spark.storage.BlockId, org.apache.spark.storage.BlockStatus)
	AddBlock(a StorageBlockIdInterface, b StorageBlockStatusInterface) 

	// public void updateBlock(org.apache.spark.storage.BlockId, org.apache.spark.storage.BlockStatus)
	UpdateBlock(a StorageBlockIdInterface, b StorageBlockStatusInterface) 

	// public boolean containsBlock(org.apache.spark.storage.BlockId)
	ContainsBlock(a StorageBlockIdInterface) bool

	// public int numBlocks()
	NumBlocks() int

	// public int numRddBlocks()
	NumRddBlocks() int

	// public int numRddBlocksById(int)
	NumRddBlocksById(a int) int

	// public long memRemaining()
	MemRemaining() int64

	// public long memUsed()
	MemUsed() int64

	// public long diskUsed()
	DiskUsed() int64

	// public long offHeapUsed()
	OffHeapUsed() int64

	// public long memUsedByRdd(int)
	MemUsedByRdd(a int) int64

	// public long diskUsedByRdd(int)
	DiskUsedByRdd(a int) int64

	// public long offHeapUsedByRdd(int)
	OffHeapUsedByRdd(a int) int64
}

type StorageStorageStatus struct {
	JavaLangObject
}

// public org.apache.spark.storage.StorageStatus(org.apache.spark.storage.BlockManagerId, long)
func NewStorageStorageStatus(a StorageBlockManagerIdInterface, b int64) (*StorageStorageStatus) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/StorageStatus", conv_a.Value().Cast("org/apache/spark/storage/BlockManagerId"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &StorageStorageStatus{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.storage.BlockManagerId blockManagerId()
func (jbobject *StorageStorageStatus) BlockManagerId() *StorageBlockManagerId {
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

// public long maxMem()
func (jbobject *StorageStorageStatus) MaxMem() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "maxMem", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void addBlock(org.apache.spark.storage.BlockId, org.apache.spark.storage.BlockStatus)
func (jbobject *StorageStorageStatus) AddBlock(a StorageBlockIdInterface, b StorageBlockStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addBlock", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/BlockId"), conv_b.Value().Cast("org/apache/spark/storage/BlockStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void updateBlock(org.apache.spark.storage.BlockId, org.apache.spark.storage.BlockStatus)
func (jbobject *StorageStorageStatus) UpdateBlock(a StorageBlockIdInterface, b StorageBlockStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "updateBlock", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/BlockId"), conv_b.Value().Cast("org/apache/spark/storage/BlockStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public boolean containsBlock(org.apache.spark.storage.BlockId)
func (jbobject *StorageStorageStatus) ContainsBlock(a StorageBlockIdInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "containsBlock", javabind.Boolean, conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int numBlocks()
func (jbobject *StorageStorageStatus) NumBlocks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numBlocks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int numRddBlocks()
func (jbobject *StorageStorageStatus) NumRddBlocks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numRddBlocks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int numRddBlocksById(int)
func (jbobject *StorageStorageStatus) NumRddBlocksById(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numRddBlocksById", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long memRemaining()
func (jbobject *StorageStorageStatus) MemRemaining() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memRemaining", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long memUsed()
func (jbobject *StorageStorageStatus) MemUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long diskUsed()
func (jbobject *StorageStorageStatus) DiskUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long offHeapUsed()
func (jbobject *StorageStorageStatus) OffHeapUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "offHeapUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long memUsedByRdd(int)
func (jbobject *StorageStorageStatus) MemUsedByRdd(a int) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memUsedByRdd", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long diskUsedByRdd(int)
func (jbobject *StorageStorageStatus) DiskUsedByRdd(a int) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskUsedByRdd", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long offHeapUsedByRdd(int)
func (jbobject *StorageStorageStatus) OffHeapUsedByRdd(a int) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "offHeapUsedByRdd", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


