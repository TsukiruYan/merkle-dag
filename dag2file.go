package merkledag

import (
	"encoding/json"
	"strings"
)

// Hash 转换为文件
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) ([]byte, error) {
	// 从 KVStore 中获取对象
	value, err := store.Get(hash)
	if err != nil {
		// 处理错误
		return nil, err
	}

	// 将获取的对象反序列化为 Object 结构体
	var obj Object
	err = json.Unmarshal(value, &obj)
	if err != nil {
		// 处理错误
		return nil, err
	}

	// 如果路径为空，返回对象的数据
	if path == "" {
		return obj.Data, nil
	}

	// 将路径分割为多个部分
	parts := strings.Split(path, "/")

	// 获取哈希函数实例
	h := hp.Get()

	// 遍历对象中的链接
	for _, link := range obj.Links {
		
		if link.Name == parts[0] {
			
			h.Reset()
			h.Write(link.Hash)
			computedHash := h.Sum(nil)

	
			return Hash2File(store, computedHash, strings.Join(parts[1:], "/"), hp)
		}
	}

	return nil, nil
}
