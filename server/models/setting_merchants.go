package models

//商户设置
type SettingMerchants struct {
	Model
	Type                 *int    //类型 1pc 2手机端
	Logo                 *string //商户logo
	CustomerServiceImage *string //客服头像
	VisitorImage         *string //访客头像
	LeftAd               *string //左侧广告
	LeftAdUrl            *string //左侧广告连接
	RightAd              *string //右侧广告
	RightAdUrl           *string //右侧广告连接
	Color                *string //窗口颜色配置
}

func (m *SettingMerchants) GetTableName() string {
	return "setting_merchants"
}
