package components

import (
	"log"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"bytes"
	//"fmt"
	"github.com/Amniversary/wedding-plugin-game/config"
	"bytes"
)

const (
	AccessCtrl       = "http://172.17.0.13:7777/rpc"
	SaiMaMessageType = 7
)

//TODO: 发送聚合短信
func SendJuHeSMS(phone string, tpId string, vCode string) (map[string]interface{}, bool) {
	key := "6962e47932431e9608350c1d5bfb523c"
	juheURL := "http://v.juhe.cn/sms/send"
	param := url.Values{}
	param.Set("mobile", phone)    //接收短信的手机号码
	param.Set("tpl_id", tpId)     //短信模板ID，请参考个人中心短信模板设置
	param.Set("tpl_value", vCode) //变量名和变量值对。如果你的变量名或者变量值中带有#&amp;=中的任意一个特殊符号，请先分别进行urlencode编码后再传递，&lt;a href=&quot;http://www.juhe.cn/news/index/id/50&quot; target=&quot;_blank&quot;&gt;详细说明&gt;&lt;/a&gt;
	param.Set("key", key)         //应用APPKEY(应用详细页查询)
	param.Set("dtype", "json")    //返回数据的格式,xml或json，默认json

	data, err := Get(juheURL, param)
	if err != nil {
		log.Printf("getJuhe Request err : %v", err)
		return nil, false
	}
	var netReturn map[string]interface{}
	json.Unmarshal(data, &netReturn)
	if netReturn["error_code"].(float64) == 0 {
		//log.Printf("接口返回result字段是:\r\n%v", netReturn)
		return netReturn, true
	}
	return nil, false
}

func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		log.Printf("parse url err: %v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Printf("get request err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Post(apiUrl string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiUrl, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
// todo: HLBUser sendBroadcast
func SendHLBUserBroadcast(idList []int64, content string, chanType string) bool {
	client := http.Client{}
	userContent := config.HLBUserContent{
		Type:    SaiMaMessageType,
		Content: content,
	}
	jsonContent, err := json.Marshal(userContent)
	if err != nil {
		log.Printf("json encode brodcast err: [%v]", err)
		return false
	}
	HLBUser := config.HLBUser{
		Type: chanType,
		IdList:idList,
		Content: string(jsonContent),
	}
	userJson, err := json.Marshal(HLBUser)
	if err != nil {
		log.Printf("json encode userJson err: [%v]", err)
		return false
	}
	newReq, err := http.NewRequest("POST", AccessCtrl, bytes.NewBuffer(userJson))
	if err != nil {
		log.Printf("http new request err: [%v]", err)
		return false
	}
	newReq.Header.Set("Content-Type", "application/json")
	newReq.Header.Set("ServerName", "AccessCtrl")
	newReq.Header.Set("MethodName", "Broadcast")
	resp, err := client.Do(newReq)
	if err != nil {
		log.Printf("http client do request err: [%v]", err)
		return false
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil realAll err: [%v]", err)
		return false
	}
	Response := &config.Response{}
	if err := json.Unmarshal(rspBody, Response); err != nil {
		log.Printf("HLBUser json decode err: [%v], [%v]", err, string(rspBody))
		return false
	}
	if Response.Code != 0 {
		return false
	}
	log.Printf("%v", string(rspBody))
	return true
}

//func SendGenCardQrcode(cardId int64) (bool, error) {
//	Url := "http://172.17.16.11:5607/api/response.do"
//	client := http.Client{}
//	data := &config.GetQrcode{CardId: cardId}
//	request := &config.RequestJson{
//		ActionName: "save_qrcode",
//		Data:       data,
//	}
//	reqBytes, err := json.Marshal(request)
//	if err != nil {
//		log.Printf("GenCardQrcode json encode err: %v", err)
//		return false, err
//	}
//	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(reqBytes))
//	if err != nil {
//		log.Printf("http new request err: %v", err)
//		return false, err
//	}
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Printf("http do request err : %v", err)
//		return false, err
//	}
//	defer resp.Body.Close()
//	rspBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("ioutil realAll err: %v", err)
//		return false, err
//	}
//	response := &config.Response{}
//	if err := json.Unmarshal(rspBody, response); err != nil {
//		log.Printf("json decode err: %v", err)
//		return false, err
//	}
//	if response.Code != config.RESPONSE_OK {
//		log.Printf("wedding card server is err: [%v], response-Code: [%d], errMsg:[%s]", request, response.Code, response.Msg)
//		return false, fmt.Errorf("wedding card service genCard error.")
//	}
//	return true, nil
//}
//
//func GetNewBusinessCard(data *config.NewBusinessCardReq) (*config.Response, bool) {
//	Url := "http://www.wedding.cn/api/response.do" //http://www.wedding.cn/api/response.do http://172.17.16.11:5607/api/response.do
//	client := http.Client{}
//	request := &config.RequestJson{
//		ActionName: "new_business_card",
//		Data:       data,
//	}
//	reqBytes, err := json.Marshal(request)
//	if err != nil {
//		log.Printf("newBusinessCard json encode err: [%v]", err)
//		return nil, false
//	}
//	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(reqBytes))
//	if err != nil {
//		log.Printf("http new request err: [%v]", err)
//		return nil, false
//	}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Printf("http do request err: [%v]", err)
//		return nil, false
//	}
//	defer resp.Body.Close()
//	rspBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("ioutil realAll err: [%v]", err)
//		return nil, false
//	}
//	response := &config.Response{}
//	if err := json.Unmarshal(rspBody, response); err != nil {
//		log.Printf("json decode err: [%v]", err)
//		return nil, false
//	}
//	if response.Code != config.RESPONSE_OK {
//		log.Printf("wedding card server is err: [%v], responseCode: [%v], errMsg: [%v]", request, response.Code, response.Msg)
//		return nil, false
//	}
//	return response, true
//}
