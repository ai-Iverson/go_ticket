package scheduled

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"go_ticket/internal/service"
	"os"
	"strings"
)

type sScheduled struct{}

func init() {
	service.RegisterScheduled(New())
}

func New() *sScheduled {
	return &sScheduled{}
}

func (s *sScheduled) getTicketToken(ctx context.Context) string {
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

func (s *sScheduled) getTicketData(ctx context.Context) ([]string, []map[string]interface{}) {
	token := s.getTicketToken(ctx)
	head := map[string]string{
		"Authorization": "Bearer " + token,
	}
	r, _ := g.Client().Header(head).Get(
		ctx,
		"http://ticket.mycyclone.com/api/questions_library?page=1&per_page=10&is_open_state=0&keyword=",
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
				words = w
				if words != "" {
					key_word = append(key_word, words)
				}

			}
		} else {
			q := strings.Replace(function_path, "【", "", -1)
			words = strings.Replace(q, "】", "", -1)
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

func (s *sScheduled) GenerateKnowledgeJsonFile(ctx context.Context) {
	keyWord, allData := s.getTicketData(ctx)
	glog.Info(ctx, keyWord)
	file, err := os.OpenFile("./keyword.josn", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(keyWord)
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
	err = encoder1.Encode(allData)
	if err != nil {
		glog.Error(ctx, err)
	}

}
