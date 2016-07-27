package main

/*
#cgo CFLAGS:-D_GNU_SOURCE -I/usr/lib/jvm/default-java/include
#cgo LDFLAGS: -ldl

#include <jni.h>

int Java_GoSparkInit_Start(JNIEnv *env, jobject jobj) {
	return start(env);
}

*/
import "C"
