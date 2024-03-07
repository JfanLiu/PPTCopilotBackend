package gpt

import (
	"backend/conf"
	"fmt"
	"math/rand"
)

// ApiKeyPool 结构体定义了 API 密钥池，包括可用的 API 密钥列表和正在使用的密钥的状态。
type ApiKeyPool struct {
	ApiKeyList []string        // 存储 API 密钥列表
	InUse      map[string]bool // 记录 API 密钥是否正在使用的映射
}

// 实例化一个ApiKeyPool结构体
var pool = ApiKeyPool{
	ApiKeyList: conf.GetGptApiKeys(), // 从配置中获取 API 密钥列表
	InUse:      map[string]bool{},    // 初始化正在使用的 API 密钥映射为空
}

// 创建一个全局的互斥锁 mutex，用于在获取和释放 API 密钥时保证同步
var mutex = make(chan bool, 1)

// GetApiKey 函数用于从 ApiKeyPool 中获取一个可用的 API 密钥
func GetApiKey() (string, error) {
	mutex <- true              // 加锁
	defer func() { <-mutex }() // 在函数返回前解锁

	// 随机选择一个可用ApiKey
	var validKeys []string
	for _, key := range pool.ApiKeyList {
		if !pool.InUse[key] {
			validKeys = append(validKeys, key)
		}
	}
	if len(validKeys) > 0 {
		// 生成随机数
		randIndex := rand.Intn(len(validKeys))
		pool.InUse[validKeys[randIndex]] = true // 标记为正在使用
		return validKeys[randIndex], nil
	}

	// 如果没有可用的 apikey ，则返回错误
	return "", fmt.Errorf("no api key available")
}

// ReleaseApiKey 函数用于释放一个 API 密钥，将其标记为未使用状态
func ReleaseApiKey(key string) {
	mutex <- true
	defer func() { <-mutex }()

	// 将指定的 API 密钥标记为未使用
	pool.InUse[key] = false
}
