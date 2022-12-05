package crons

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"os"
	"strings"
	"time"
)

var GetKnowledgeJson = &cGetKnowledgeJson{name: "getknowledgejson"}

type cGetKnowledgeJson struct {
	name string
}

func (c *cGetKnowledgeJson) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cGetKnowledgeJson) Execute(ctx context.Context) {
	c.GenerateKnowledgeJsonFile(ctx)
	g.Log().Infof(ctx, "cron GetKnowledgeJson Execute:%v", time.Now())
}

func (c *cGetKnowledgeJson) getTicketToken(ctx context.Context) string {
	r, err := g.Client().Post(
		ctx,
		"https://portal-pro.mycyclone.com/api/v1/cas/login",
		`{"username": "xin.zhang", "password": "#1995119zx"}`,
	)
	defer r.Close()
	if err != nil {
		glog.Error(ctx, err)
	}
	data := r.ReadAll()
	res := make(map[string]interface{})
	err = json.Unmarshal(data, &res)
	if err != nil {
		glog.Error(ctx, "解析json数据失败", data, err)
	}
	glog.Info(ctx, "token为：", res["token"].(string))
	return res["token"].(string)
}

func (c *cGetKnowledgeJson) getTicketData(ctx context.Context) ([]string, []map[string]interface{}) {
	token := c.getTicketToken(ctx)
	head := map[string]string{
		"Authorization": "Bearer " + token,
	}
	r, _ := g.Client().Header(head).Get(
		ctx,
		"http://ticket.mycyclone.com/api/questions_library?page=1&per_page=5000&is_open_state=0&keyword=",
	)
	defer r.Close()
	ticketJson := r.ReadAll()
	res := make(map[string]interface{})
	_ = json.Unmarshal(ticketJson, &res)
	tickeData := res["items"].([]interface{})
	//fmt.Println(res["items"].([]interface{}))
	var words string
	var key_word []string
	var allData []map[string]interface{}
	for _, j := range tickeData {
		oneData := make(map[string]interface{})

		function_path := j.(map[string]interface{})["function_path"].(string)
		if strings.Contains(j.(map[string]interface{})["function_path"].(string), "-") {
			word := strings.Split(function_path, "-")
			for _, i := range word {
				q := strings.Replace(i, "【", "", -1)
				w := strings.Replace(q, "】", "", -1)
				s := strings.TrimSpace(w)
				words = s
				if words != "" {
					key_word = append(key_word, words)
				}

			}
		} else {
			q := strings.Replace(function_path, "【", "", -1)
			s := strings.Replace(q, "】", "", -1)
			words = strings.TrimSpace(s)
			if words != "" {
				key_word = append(key_word, words)
			}

		}
		if j.(map[string]interface{})["open_state"].(bool) {
			oneData["Content"] = j.(map[string]interface{})["suggested_solution"].(string)
			oneData["Description"] = j.(map[string]interface{})["problem_desc"].(string)
			oneData["Title"] = j.(map[string]interface{})["summary"].(string)
			oneData["PostDate"] = j.(map[string]interface{})["created_at"].(string)
			oneData["Label"] = j.(map[string]interface{})["problem_lable_info"].([]interface{})[0].(map[string]interface{})["name"].(string)
			oneData["Keyword"] = words
			oneData["code"] = j.(map[string]interface{})["code"].(string)

		} else {
			oneData["Content"] = j.(map[string]interface{})["suggested_solution"].(string)
			oneData["Description"] = "具体方案请联系技术支持"
			oneData["Title"] = j.(map[string]interface{})["summary"].(string)
			oneData["PostDate"] = j.(map[string]interface{})["created_at"].(string)
			oneData["Label"] = j.(map[string]interface{})["problem_lable_info"].([]interface{})[0].(map[string]interface{})["name"].(string)
			oneData["Keyword"] = words
			oneData["code"] = j.(map[string]interface{})["code"].(string)
		}
		allData = append(allData, oneData)
	}
	fmt.Println(key_word)
	return key_word, allData
}

func (c *cGetKnowledgeJson) GenerateKnowledgeJsonFile(ctx context.Context) {
	keyWord, allData := c.getTicketData(ctx)
	// 去重
	keyWords := removeDuplicateElement(keyWord)
	glog.Info(ctx, keyWords)

	file, err := os.OpenFile("./keyword.josn", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(keyWords)
	if err != nil {
		glog.Error(ctx, err)
	}

	// TODO 文件的路径写在配置文件中
	file2, err := os.OpenFile("./Knowledge.josn", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		glog.Error(ctx, err)
	}
	encoder1 := json.NewEncoder(file2)
	// https://learnku.com/articles/58053
	// json.Marshal() 在进行序列化时，会进行 HTMLEscape 编码，会将 “<”, “>”, “&”, U+2028, 及 U+2029 转码成 “\u003c”,”\u003e”, “\u0026”, “\u2028”, 和 “\u2029”
	encoder1.SetEscapeHTML(false)
	err = encoder1.Encode(allData)
	if err != nil {
		glog.Error(ctx, err)
	}

}

// 字符串切片去重
func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok { //如果字典中找不到元素，ok=false，!ok为true，就往切片中append元素。
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
