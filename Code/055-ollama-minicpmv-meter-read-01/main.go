package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Ollama 请求结构体
// 参考 Ollama API 文档: https://github.com/ollama/ollama/blob/main/docs/api.md#generate-a-completion
type OllamaRequest struct {
	Model  string   `json:"model"`
	Prompt string   `json:"prompt"`
	Images []string `json:"images"`
	Stream bool     `json:"stream"` // 设置为 false 以一次性获取结果
}

// Ollama 响应结构体
type OllamaResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func main() {
	// 1. 配置参数
	imagePath := "images/5.jpg" // 替换为你本地的水表/气表图片路径
	modelName := "minicpm-v"    // 确保这里与你 pull 的模型名称一致

	// 2. 读取并转换图片为 Base64
	base64Image, err := imageToBase64(imagePath)
	if err != nil {
		fmt.Printf("❌ 图片处理失败: %v\n", err)
		return
	}

	// 3. 构建针对仪表识别的 Prompt
	// 关键点：明确要求输出 JSON 格式，方便后续代码解析
	promptText := `
	你是一个专业的工业仪表识别助手。
	请仔细观察这张图片中的水表或气表。
	1. 识别表盘上的数字读数。如果是机械指针表，请根据指针位置估算读数（读取黑色指针，忽略红色指针）。如果是液晶数字，请直接读取数字。
	2. 忽略小数点后的红色数字（如果有）。
	3. 请仅输出一个标准的 JSON 对象，不要包含 markdown 标记或其他废话。
	JSON 格式示例: {"reading": 123.4, "type": "mechanical"}
	`

	// 4. 调用 Ollama API
	result, err := callOllama(modelName, promptText, base64Image)
	if err != nil {
		fmt.Printf("❌ 调用 Ollama 失败: %v\n", err)
		return
	}

	// 5. 输出结果
	fmt.Println("✅ 识别成功！模型原始返回：")
	fmt.Println(result.Response)

	// (可选) 这里可以添加 JSON 解析逻辑，将结果存入数据库
}

// 图片转 Base64 函数
func imageToBase64(path string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("文件不存在: %s", path)
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Ollama 需要的是纯 Base64 字符串，不需要 "data:image/jpeg;base64," 前缀
	encoded := base64.StdEncoding.EncodeToString(fileBytes)
	return encoded, nil
}

// 调用 Ollama API 的核心函数
func callOllama(model, prompt, imageBase64 string) (*OllamaResponse, error) {
	// 构建请求体
	reqBody := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Images: []string{imageBase64},
		Stream: false, // 关闭流式，等待完整结果
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// 发送 POST 请求到本地 Ollama 服务
	url := "http://localhost:11434/api/generate"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 请求失败: %s, 内容: %s", resp.Status, string(body))
	}

	var result OllamaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
