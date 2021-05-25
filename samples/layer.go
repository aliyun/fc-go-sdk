package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aliyun/fc-go-sdk"
)

func main() {
	client, _ := fc.NewClient(
		os.Getenv("ENDPOINT"), "2016-08-15",
		os.Getenv("ACCESS_KEY_ID"),
		os.Getenv("ACCESS_KEY_SECRET"),
		fc.WithTransport(&http.Transport{MaxIdleConnsPerHost: 100}))

	// 层名称
	layerName := "test-layer"
	// 准备 Zip 格式的层文件
	layZipFile := "./hello_world.zip"
	// 指定兼容的运行时环境
	compatibleRuntime := []string{"python3", "nodejs12"}

	// 1. 发布层版本
	fmt.Println("Publish layer versions")
	data, err := ioutil.ReadFile(layZipFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	publishLayerVersionOutput, err := client.PublishLayerVersion(fc.NewPublishLayerVersionInput().
		WithLayerName(layerName).
		WithCode(fc.NewCode().WithZipFile(data)).
		WithCompatibleRuntime(compatibleRuntime).
		WithDescription("my layer"),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("PublishLayerVersion response: %+v \n\n", publishLayerVersionOutput)
	}

	// 2. 查询指定层版本信息
	fmt.Printf("Get the layer of version %d\n", publishLayerVersionOutput.Layer.Version)
	getLayerVersionOutput, err := client.GetLayerVersion(
		fc.NewGetLayerVersionInput(layerName, publishLayerVersionOutput.Layer.Version))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("GetLayerVersion response: %+v \n\n", getLayerVersionOutput.Layer)
	}

	// 3. 获取层列表
	fmt.Println("List layers")
	nextToken := ""
	layers := []*fc.Layer{}
	for {
		listLayersOutput, err := client.ListLayers(
			fc.NewListLayersInput().
				WithLimit(100).
				WithNextToken(nextToken))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if len(listLayersOutput.Layers) != 0 {
			layers = append(layers, listLayersOutput.Layers...)
		}
		if listLayersOutput.NextToken == nil {
			break
		}
		nextToken = *listLayersOutput.NextToken

	}
	fmt.Println("ListLayers response:")
	for _, layer := range layers {
		fmt.Printf("- layerName: %s, layerMaxVersion: %d\n", layer.LayerName, layer.Version)
	}

	// 4. 获取层版本列表
	fmt.Println("List layer versions")
	// 层的起始版本,默认从1开始
	startVersion := int32(1)
	fmt.Println("ListLayerVersions response:")
	layerVersions := []*fc.Layer{}
	for {
		listLayerVersionsOutput, err := client.ListLayerVersions(
			fc.NewListLayerVersionsInput(layerName, startVersion).
				WithLimit(100))
		if err != nil {
			if err, ok := err.(*fc.ServiceError); ok &&
				err.HTTPStatus == http.StatusNotFound {
				break
			}
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if len(listLayerVersionsOutput.Layers) > 0 {
			layerVersions = append(layerVersions, listLayerVersionsOutput.Layers...)
		}
		if listLayerVersionsOutput.NextVersion == nil ||
			*listLayerVersionsOutput.NextVersion == 0 {
			break
		}
		startVersion = *listLayerVersionsOutput.NextVersion
	}

	for _, layer := range layerVersions {
		fmt.Printf("- layerName: %s, layerVersion: %d\n", layer.LayerName, layer.Version)
	}

	// 5. 删除层版本
	fmt.Printf("Delete the layer of version %d \n", publishLayerVersionOutput.Layer.Version)
	deleteLayerVersionOutput, err := client.DeleteLayerVersion(
		fc.NewDeleteLayerVersionInput(layerName, publishLayerVersionOutput.Layer.Version))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("DeleteLayerVersion response: %+v \n\n", deleteLayerVersionOutput)
	}
}
