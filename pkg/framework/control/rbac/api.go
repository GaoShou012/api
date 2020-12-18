package rbac

import (
	"errors"
	"framework/class/rbac"
	"framework/env"
)

func CreateApi(operator rbac.Operator, api rbac.Api) error {
	return env.ApiAdapter.Create(api)
}
func UpdateApi(operator rbac.Operator, apiId uint64, api rbac.Api) error {
	// 校验操作者权限
	ok, err := env.ApiAdapter.VerifyIdWithOperator(apiId, operator)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("权限不足")
	}

	// 更新API信息
	if err := env.ApiAdapter.Update(apiId, api); err != nil {
		return err
	}

	return nil
}

func SelectApiByPage(operator rbac.Operator, page uint64, pageSize uint64) ([]rbac.Api, error) {
	return env.ApiAdapter.SelectByPage(operator.GetTenantId(), page, pageSize)
}

