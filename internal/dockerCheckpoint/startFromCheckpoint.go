package dockerCheckpoint

import (
	"context"
	"fmt"
	"log"
	_ "sigs.k8s.io/controller-runtime"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	myTypes "kangmingfa/kubernetes-container-checkpoint/types"
)

func StartFromCheckpoint(checkpoint string) myTypes.Result {
	// 创建 Docker 客户端
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// 远程机器的 Docker 主机地址
	remoteHost := "tcp://192.168.195.43:2375"

	// 配置远程连接
	cli, err = client.NewClientWithOpts(client.WithHost(remoteHost))
	if err != nil {
		log.Fatal(err)
	}

	// 指定要加载的 checkpoint 的路径 5568
	checkpointPath := "checkpoint2"

	// 指定要启动的容器名称
	containerName := "cr"
	containerID := "4faf3222d027"

	// 检查容器是否存在
	_, err = cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return myTypes.Result{Msg: "container is not exist", Success: false}
	}

	// 创建启动容器的选项
	options := types.ContainerStartOptions{
		CheckpointID: checkpointPath,
	}

	// 启动容器
	err = cli.ContainerStart(context.Background(), containerName, options)
	if err != nil {
		return myTypes.Result{Msg: err.Error(), Success: false}
	}

	fmt.Printf("Container '%s' started from checkpoint '%s'\n", containerName, checkpointPath)
	return myTypes.Result{Msg: "start from checkpoint successful", Success: true}
}
