package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateInstanceNameRequest struct {

	// 语言
	XLanguage *UpdateInstanceNameRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifiyInstanceNameRequest `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}

type UpdateInstanceNameRequestXLanguage struct {
	value string
}

type UpdateInstanceNameRequestXLanguageEnum struct {
	ZH_CN UpdateInstanceNameRequestXLanguage
	EN_US UpdateInstanceNameRequestXLanguage
}

func GetUpdateInstanceNameRequestXLanguageEnum() UpdateInstanceNameRequestXLanguageEnum {
	return UpdateInstanceNameRequestXLanguageEnum{
		ZH_CN: UpdateInstanceNameRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateInstanceNameRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateInstanceNameRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceNameRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
