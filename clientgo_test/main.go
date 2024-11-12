package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := "C:\\custom\\project\\go\\config"
	//将kubeconfig格式化微restclient.config类型
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	//通过config创建clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//通过clientset获取pods列表,明明空间是default
	//context.TODO() 控制上下文环境，比如请求超时时间
	//metav1.ListOptions{} 可填写对应的过滤配置，比如按标签
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pods.Items[0])

}
