package planBuilder

import (
	logger "github.com/ShaoWenAcerLu/BCIR_main/utilities"
	validator "github.com/ShaoWenAcerLu/BCIR_validator"
)

func BuildPlan(policyName string) ([]string, error) {
	// validate the connection to source and target systems in policy
	var messages []string
	err := validator.CheckPolicySysConn(policyName)
	if err != nil {
		return messages, err
	} else {
		msg := "Success check connection to source and target system"
		logger.LogInfo(msg)
		messages = append(messages, msg)
	}

	return messages, nil
}
