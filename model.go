package oilPaintingChaincode

// 油画创作的步骤1,2,3,4
const (
	Step1 = 0
	Step2 = 1
	Step3 = 2
	Step4 = 3
)

// confirmStatus
const (
	WU = 0
	// 待确认
	TCB = 1
	// 已确认
	CONFIRMED = 2
	// 已拒绝
	CONFUSED = 3
)

// 油画基本信息
type OilBaseInfo struct {
	// docType
	DocType string `json:"docType"`
	// 定制id
	CustomId string `json:"customId"`
	// 作品名称
	WorksName string `json:"worksName"`
	// 油画创作过程的步骤：1. 2. 3. 4
	ProductionProcess int `json:"productionProcess"`
	// 定制油画确认状态：0 1 2 3
	ConfirmStatus int `json:"confirmStatus"`
	// 发起时间
	CreateTime string `json:"createTime"`
	// 更新时间
	UpdateTime string `json:"updateTime"`
	// 确认时间
	ConfirmTime string `json:"confirmTime"`
	// 画家
	Painter string `json:"painter"`
	// 创作日期
	ManufactureDate string `json:"manufactureDate"`
	// 所属版权
	Copyright string `json:"copyright"`
	// 画布
	Canvas string `json:"canvas"`
	// 画框
	Frame string `json:"frame"`
	// 颜料
	Pigment string `json:"pigment"`
	// 画布图地址
	CanvasImage string `json:"canvasImage"`
	// 画布图hash
	CanvasImageHash string `json:"canvasImageHash"`
	// 画框图地址
	FrameImage string `json:"frameImage"`
	// 画框图hash
	FrameImageHash string `json:"frameImageHash"`
	// 颜料图地址
	PigmentImage string `json:"pigmentImage"`
	// 颜料图hash
	PigmentImageHash string `json:"pigmentImageHash"`
}

// 素描稿信息
type OilSketchInfo struct {
	//
	DocType string `json:"docType"`
	// 定制id
	CustomId string `json:"customId"`
	// 油画创作过程的步骤：1. 2. 3. 4
	ProductionProcess int `json:"productionProcess"`
	// 定制油画确认状态：0 1 2 3
	ConfirmStatus int `json:"confirmStatus"`
	// 发起时间
	CreateTime string `json:"createTime"`
	// 更新时间
	UpdateTime string `json:"updateTime"`
	// 确认时间
	ConfirmTime string `json:"confirmTime"`
	// 素描图片地址
	SketchImage string `json:"sketchImage"`
	// 素描图片hash
	SketchImageHash string `json:"sketchImageHash"`
}

// 上色信息
type OilColorInfo struct {
	//
	DocType string `json:"docType"`
	// 定制id
	CustomId string `json:"customId"`
	// 油画创作过程的步骤：1. 2. 3. 4
	ProductionProcess int `json:"productionProcess"`
	// 定制油画确认状态：0 1 2 3
	ConfirmStatus int `json:"confirmStatus"`
	// 发起时间
	CreateTime string `json:"createTime"`
	// 更新时间
	UpdateTime string `json:"updateTime"`
	// 确认时间
	ConfirmTime string `json:"confirmTime"`
	// 上色图片地址
	ColorImage string `json:"colorImage"`
	// 上色图片hash
	ColorImageHash string `json:"colorImageHash"`
}

// 细节绘制信息
type OilShadowInfo struct {
	//
	DocType string `json:"docType"`
	// 定制id
	CustomId string `json:"customId"`
	// 油画创作过程的步骤：1. 2. 3. 4
	ProductionProcess int `json:"productionProcess"`
	// 定制油画确认状态：0 1 2 3
	ConfirmStatus int `json:"confirmStatus"`
	// 发起时间
	CreateTime string `json:"createTime"`
	// 更新时间
	UpdateTime string `json:"updateTime"`
	// 确认时间
	ConfirmTime string `json:"confirmTime"`
	// 细节绘制图片地址
	ShadowImage string `json:"shadowImage"`
	// 细节绘制图片hash
	ShadowImageHash string `json:"shadowImageHash"`
}

// key 链上的key值
type Key struct {
	// DocType
	DocType string `json:"docType"`
	// 定制id
	CustomId string `json:"customId"`
	// 油画创作过程的步骤：0. 1. 2. 3
	ProductionProcess int `json:"productionProcess"`
}

type TimeValue struct {
	// 发起时间
	CreateTime string `json:"createTime"`
	// 更新时间
	UpdateTime string `json:"updateTime"`
	// 确认时间
	ConfirmTime string `json:"confirmTime"`
}

// request
type ChainCodeRequest struct {
	// DocType
	DocType string `json:"docType"`
	// 定制id
	CustomId string `json:"customId"`
	// 油画创作过程的步骤：0. 1. 2. 3
	ProductionProcess int `json:"productionProcess"`
	// confirm
	ConfirmStatus int `json:"confirmStatus"`
}
