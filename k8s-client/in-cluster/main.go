/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {

	// 	// os.Getenv("KUBERNETES_SERVICE_HOST")
	// 	KUBERNETES_SERVICE_HOST := "localhost"
	// 	os.Setenv("KUBERNETES_SERVICE_HOST", KUBERNETES_SERVICE_HOST)
	// 	// os.Getenv("KUBERNETES_SERVICE_PORT")
	// 	KUBERNETES_SERVICE_PORT := "49577"
	// 	os.Setenv("KUBERNETES_SERVICE_PORT", KUBERNETES_SERVICE_PORT)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		foundPod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "my-prometheus-alertmanager-0", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod my-prometheus-alertmanager-0 not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found my-prometheus-alertmanager-0 pod in default namespace\n")
			fmt.Printf("POD %s Status: %s\n", foundPod.Name, foundPod.Status.Conditions[0].Status)
		}

		fmt.Println("PODS STATUS")
		fmt.Printf("POD_NAME\t\t\t\tNAMESPACE\t\t\t\tSTATUS \n")
		for _, pod := range pods.Items {
			fmt.Printf("%s - %s - %s\n", pod.Name, pod.Namespace, pod.Status.Phase)
		}

		deploy, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d deployments in the cluster\n", len(deploy.Items))

		fmt.Println("DEPLOYMENTS STATUS")
		fmt.Printf("DEPLOYMENT_NAME\t\t\t\tNAMESPACE\t\t\t\tSTATUS \n")
		for _, dep := range deploy.Items {
			fmt.Printf("%s - %s - %s\n", dep.Name, dep.Namespace, &dep.Status.Conditions[0])
		}

		time.Sleep(30 * time.Second)
	}
}
