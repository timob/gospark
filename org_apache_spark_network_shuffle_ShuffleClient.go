package gospark

import "github.com/timob/javabind"

type NetworkShuffleShuffleClientInterface interface {
	JavaLangObjectInterface

	// public void init(java.lang.String)
	Init(a string) 

	// public abstract void fetchBlocks(java.lang.String, int, java.lang.String, java.lang.String[], org.apache.spark.network.shuffle.BlockFetchingListener)
	FetchBlocks(a string, b int, c string, d []string, e NetworkShuffleBlockFetchingListenerInterface) 
}

type NetworkShuffleShuffleClient struct {
	JavaLangObject
}

// public org.apache.spark.network.shuffle.ShuffleClient()
func NewNetworkShuffleShuffleClient() (*NetworkShuffleShuffleClient) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/network/shuffle/ShuffleClient")
	if err != nil {
		panic(err)
	}
	x := &NetworkShuffleShuffleClient{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void init(java.lang.String)
func (jbobject *NetworkShuffleShuffleClient) Init(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "init", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void fetchBlocks(java.lang.String, int, java.lang.String, java.lang.String[], org.apache.spark.network.shuffle.BlockFetchingListener)
func (jbobject *NetworkShuffleShuffleClient) FetchBlocks(a string, b int, c string, d []string, e NetworkShuffleBlockFetchingListenerInterface)  {
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


