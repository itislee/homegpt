package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"time"
	"os"
)

const (
	Sensor = "sensor"
	BinarySensor = "binary_sensor"
	AlarmPanel = "alarm_control_panel"
	Light = "light"
	Switch = "switch"
	Script = "script"
)

var StatePath = "states/%s.%s"
var ServiceSetValuePath = "services/text/set_value"
var LastConversationTime string

type StatePost struct {
	State string	`json:"state"`
	Attributes struct {
		FriendlyName	string	`json:"friendly_name"`
		UnitMeasure	string	`json:"unit_of_measurement"`
	} `json:"attributes"`
}

type Conversation struct {
	EntifyId string `json:"entity_id"`
	State string `json:"state"`
	Attributes struct {
		EntityClass string `json:"entity_class"`
		ParentEntityId string `json:"parent_entity_id"`
		Content string `json:"content"`
		Answers []struct {
			Type string `json:"type"`
			Tts struct {
				Text string `json:"text"`
			}`json:"tts"`
		} `json:"answers"`
		History []string `json:"history"`
		Timestamp string `json:"timestamp`
		Icon string `json:"icon"`
		FriendlyName string `json:"friend_name"`
		SupportedFeatures int `json:"supported_features"`
	} `json:"attributes"`
	LastChange string `json:"last_change"`
	LastUpdated string `json:"last_updated"`
	Context struct {
		Id string `json:"id"`
		ParentId string `json:"parent_id`
		UserId string `json:"user_id`
	}`json:"context"`
}

type SetValueData struct{
	EntityId string `json:"entity_id"`
	Value string `json:"value"`
}

func (ha *HomeAssistant)Loop() {
	for {
		question := ha.NeedUseGPT()
		if question != "" {
			ha.ServicesTextSetValue("text.xiaomi_lx06_de09_play_text", fmt.Sprintf("正在询问%s，请稍等", ha.botClient.Name()))

			fmt.Printf("asking:%s\n", question)
			r := ha.botClient.Ask(question)
			fmt.Printf("result:%s", r)
			ha.ServicesTextSetValue("text.xiaomi_lx06_de09_play_text", r)
		}
		time.Sleep(1 * time.Second)
	}
}

func (ha *HomeAssistant)NeedUseGPT() string {
	con, t, err := ha.GetConversasion(Sensor, "xiaomi_lx06_de09_conversation")
	if err != nil {
		return ""
	}
	if len(con) < 1 {
		// 没有会话数据
		fmt.Printf("no msg\n")
		return ""
	}
	if LastConversationTime=="" || t==LastConversationTime {
		// 没有更新
		fmt.Printf("no new msg\n")
		return ""
	}
	fmt.Printf("new msg[%s]:%s\n", t, con)
	LastConversationTime = t

	// 优先从环境变量中获取，如果没有设置使用默认值
	promptWord := os.Getenv("PROMPT_WORD")
	if promptWord == "" {
		promptWord = "请问"
	}
	if strings.HasPrefix(con, promptWord) {
		return con
	}
	return ""
}

func (ha *HomeAssistant) GetConversasion(sensorType string, entityId string) (string, string, error) {
	path := fmt.Sprintf(StatePath, sensorType, entityId)
	rsp, err := ha.httpGet(path)
	if err != nil {
		return "", "", err
	}
	//fmt.Printf("rsp %s", string(rsp))
	con := Conversation{}
	err = json.Unmarshal(rsp, &con)
	if err != nil {
		return "", "", err
	}
	//fmt.Printf("conversation:%+v\n", con)
	return con.State, con.Attributes.Timestamp, nil
}

func (ha *HomeAssistant) ServicesTextSetValue(entityId, value string) error {
	d := SetValueData{
		EntityId: entityId,
		Value: value,
	}
	jsonObj, err := json.Marshal(d)
	if err != nil {
		return err
	}
	_, err = ha.httpPost(ServiceSetValuePath, jsonObj)
	if err != nil {
		return err
	}
	return nil
}

func (ha *HomeAssistant) SendState(sensorType string, entityId string, st StatePost) error {
	path := fmt.Sprintf(StatePath, sensorType, entityId)
	jsonObj, err := json.Marshal(st)
	if err != nil {
		return err
	}
	//the resp body is not needed at this moment, there is not useful information there
	_, err = ha.httpPost(path, jsonObj)
	if err != nil {
		return err
	}
	return nil
}