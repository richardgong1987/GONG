package system

import (
	"github.com/gin-gonic/gin"
	"github.com/richardgong1987/server/global"
	"github.com/richardgong1987/server/model/common/request"
	"github.com/richardgong1987/server/model/common/response"
	"github.com/richardgong1987/server/model/system"
	systemReq "github.com/richardgong1987/server/model/system/request"
	"github.com/richardgong1987/server/utils"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// DeleteSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysOperationRecord      true  "SysOperationRecord模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindJSON(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysOperationRecordByIds
// @Tags      SysOperationRecord
// @Summary   批量删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "批量删除SysOperationRecord"
// @Success   200   {object}  response.Response{msg=string}  "批量删除SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// FindSysOperationRecord
// @Tags      SysOperationRecord
// @Summary   用id查询SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysOperationRecord                                  true  "Id"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysOperationRecord"
// @Router    /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindQuery(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(sysOperationRecord, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysOperationRecord, err := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysOperationRecord": reSysOperationRecord}, "查询成功", c)
}

// GetSysOperationRecordList
// @Tags      SysOperationRecord
// @Summary   分页获取SysOperationRecord列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysOperationRecordSearch                        true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := operationRecordService.GetSysOperationRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
