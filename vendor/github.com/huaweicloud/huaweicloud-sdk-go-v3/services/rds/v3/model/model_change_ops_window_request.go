package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ChangeOpsWindowRequest struct {

	// 语言
	XLanguage *ChangeOpsWindowRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OpsWindowRequest `json:"body,omitempty"`
}

func (o ChangeOpsWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOpsWindowRequest struct{}"
	}

	return strings.Join([]string{"ChangeOpsWindowRequest", string(data)}, " ")
}

type ChangeOpsWindowRequestXLanguage struct {
	value string
}

type ChangeOpsWindowRequestXLanguageEnum struct {
	ZH_CN ChangeOpsWindowRequestXLanguage
	EN_US ChangeOpsWindowRequestXLanguage
}

func GetChangeOpsWindowRequestXLanguageEnum() ChangeOpsWindowRequestXLanguageEnum {
	return ChangeOpsWindowRequestXLanguageEnum{
		ZH_CN: ChangeOpsWindowRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeOpsWindowRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeOpsWindowRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeOpsWindowRequestXLanguage) UnmarshalJSON(b []byte) error {
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
