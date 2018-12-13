package code

const (
	//ActionSuccess 成功
	ActionSuccess = 20000
	// ActionArgsError 参数错误
	ActionArgsError = 20001
	// ActionDbError 数据库错误
	ActionDbError = 20002
	// ActionCreateUserError 创建用户错误
	ActionCreateUserError = 20003
	// ActionUserNotFound 用户不存在
	ActionUserNotFound = 20004
	// ActionPropNotExists 道具不存在
	ActionPropNotExists = 20005
	// ActionPropNotEnough 道具不足
	ActionPropNotEnough = 20006
	// ActionGoldNotEnough 金币不足
	ActionGoldNotEnough = 20007
	// ActionUnknownIllustrationID 未知的图集ID
	ActionUnknownIllustrationID = 20008
	// ActionNotGetIllustrationAward 不能获取图集奖励
	ActionNotGetIllustrationAward = 20009
	// ActionGetWeixinOpenIDError  获取微信ID错误
	ActionGetWeixinOpenIDError = 20010
	// ActionOpenIDIsNotEmpty  OpenID是不能为空的
	ActionOpenIDIsNotEmpty = 20011
	// ActionNotAccOfflineNum 没有离线能量可以领取
	ActionNotAccOfflineNum = 20012
	// ActionNotLegitimateConstellID 非合法的星座ID
	ActionNotLegitimateConstellID = 20013
	// ActionTimeNotEnough 时间不足
	ActionTimeNotEnough = 20014
	// ActionUnknownConstellationBuffType 未知的buff类型
	ActionUnknownConstellationBuffType = 20015
	// ActionNotLegitimateStarNum 未知的StarNum
	ActionNotLegitimateStarNum = 20016
	// ActionStarAlreadyStartUp 星星已经被点亮了
	ActionStarAlreadyStartUp = 20017
	// ActionPowerNotEnough 能量不够哇
	ActionPowerNotEnough = 20018
	// ActionNotRebirth 没有通关, 不能重生
	ActionNotRebirth = 20019
	// ActionNotChoiceConstellationBuff 不能选择buff 应为重生次数不为1
	ActionNotChoiceConstellationBuff = 20020
	// ActionUnkonownAchievementTypeOrOrder 未知的成就类型或者序列
	ActionUnkonownAchievementTypeOrOrder = 20021
	// ActionAchievementNotExists  成就不存在, 意思就是这个成就还没有满足
	ActionAchievementNotExists = 20022
	// ActionAlreadyHaveSpeedPlayer 该玩家已经结成了速配
	ActionAlreadyHaveSpeedPlayer = 20023
	// ActionUnknownMessageID 未知的消息ID
	ActionUnknownMessageID = 20024
	// ActionBeReplyMessageNotSpeedMsg 被回复的消息,不是速配消息
	ActionBeReplyMessageNotSpeedMsg = 20025
	// ActionUnknownConstellationID 未知的星座ID
	ActionUnknownConstellationID = 20026
	// ActionConstellationTagRepeated 星座标签重复
	ActionConstellationTagRepeated = 20027
	// ActionContentLegitimate 内容违法
	ActionContentLegitimate = 20028
	// ActionConstellationTagNotIsExists 星座标签不存在
	ActionConstellationTagNotIsExists = 20029
	// ActionUnknownWishingID 未知的许愿ID
	ActionUnknownWishingID = 20030
)
