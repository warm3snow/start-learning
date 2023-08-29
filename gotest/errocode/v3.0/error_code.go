/**
 * @Author: xueyanghan
 * @File: error_code.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/6 18:22
 */

package errocode

import (
	"fmt"
)

const (
	DefaultErrCode = ErrCodePanic
)

const (
	ErrorTypeCloudApi = 0
	ErrorTypeEN       = 1
	ErrorTypeCN       = 2
)

// ErrCode - err code
type ErrCode uint32

// error code
const (
	ErrCodeOK ErrCode = 0

	// 系统缺省错误: 100~199
	ErrCodeInvalidParameter      ErrCode = iota + 100 // 参数错误（包括参数格式、类型等错误）
	ErrCodeInvalidParameterValue                      // 参数取值错误
	ErrCodeMissingParameter                           // 缺少参数错误，必传参数没填
	ErrCodeUnknownParameter                           // 未知参数错误，用户多传未定义的参数会导致错误
	ErrCodeInternalError
)
const (
	// 联盟及成员管理错误码: 200~299
	ErrCodeAddConsortium                     ErrCode = iota + 200 // 联盟信息上链失败
	ErrCodeUpdateConsortium                                       // 更新联盟信息上链失败
	ErrCodeGetConsortiumByName                                    // 根据联盟名称获取联盟详情失败
	ErrCodeConsortiumNotExists                                    // 指定联盟不存在
	ErrCodeBindConsMemberCrt                                      // 绑定联盟成员证书失败
	ErrCodeBindConsMemberEncryptPk                                // 绑定联盟成员加密公钥失败
	ErrCodeUpdateConsMemberStatus                                 // 更新联盟成员状态失败
	ErrCodeCreateInviteMemberEvent                                // 创建联盟成员邀请事件失败
	ErrCodeGetConsMemberByAppId                                   // 根据AppId从链上获取联盟成员失败
	ErrCodeMemberAlreadyBelongsToCons                             // 成员已归属该联盟
	ErrCodeMemberNotBelongsToAnyCons                              // 成员未归属任何联盟
	ErrCodeMemberNotBelongsToThisCons                             // 成员未归属该联盟
	ErrCodeGetConsortiumListForMember                             // 获取成员所在联盟列表失败
	ErrCodeGetConsortiumMembers                                   // 获取联盟拥有成员信息失败
	ErrCodeGetMemberConsNames                                     // 获取成员所属联盟名称列表失败
	ErrCodeGetJoinedMemberList                                    // 获取联盟已加入成员列表失败
	ErrCodeGetInvitedMemberList                                   // 获取联盟邀请中成员列表失败
	ErrCodeGetInvitingMemberList                                  // 获取邀请中的联盟成员列表失败
	ErrCodeMemberAlreadyInInviting                                // 联盟成员已经在邀请中
	ErrCodeExitConsortium                                         // 成员退出联盟失败
	ErrCodeDownloadMemberCert                                     // 下载联盟成员托管私钥和证书失败
	ErrCodeDownloadMemberEncryptPk                                // 下载联盟成员数据加密公钥失败
	ErrCodeInvalidMemberRole                                      // 成员不具备相关角色
	ErrCodeConsortiumExists                                       // 联盟已存在
	ErrCodeApplyCert                                              // 获取联盟成员链身份失败
	ErrCodeGetMemberCertKeyInfoByAppIdFromDB                      // 从DB获取成员链身份及数据加密公钥失败
	ErrCodeInsertMemberCertKeyInfoToDB                            // 成员链身份及数据加密公钥新增入库失败
	ErrCodeGetUserNicknameByAppId                                 // 根据AppId获取用户昵称失败
	ErrCodeGetMyTeeResourceEncryptPk                              // 获取自己可信计算加密公钥失败
	ErrCodeAddMemberSignCrt                                       // 联盟成员链身份证书上链失败
	ErrCodeAddMemberTeeResourceEncryptPkList                      // 联盟成员计算资源加密证书上链失败
	ErrCodeConsortiumNotOwner                                     // 非联盟所有者
	ErrCodeConsMemberExit                                         // 成员已退出
)
const (

	// 计算模型管理错误码: 300~399
	ErrCodeGetCalcModel                     ErrCode = iota + 300 // 从链上获取计算模型信息失败
	ErrCodeGetCOSPresignedURL                                    // 生成计算模型COS预签名URL失败
	ErrCodeInvalidAppId                                          // 无效的AppId
	ErrCodeNotVisible                                            // 不可见的计算模型
	ErrCodeInvalidCalcModelNo                                    // 无效的calcModelNo
	ErrCodeCalcModelExists                                       // 计算模型已存在
	ErrCodeGetCalcModelContent                                   // 获取计算模型内容失败
	ErrCodeReleaseStoreCalcModel                                 // 发布计算模型到模型仓库失败
	ErrCodeUpdateReleaseStatusCalcModel                          // 更改计算模型发布状态失败
	ErrCodeReleaseCalcModel                                      // 发布计算模型上链失败
	ErrCodeUpdateCalcModel                                       // 更新计算模型上链失败
	ErrCodeDeleteCalcModel                                       // 删除计算模型失败
	ErrCodeGetCalcModelList                                      // 从链上获取计算联盟列表信息失败
	ErrCodeCalcPkgChecksum                                       // 计算模型checksum失败
	ErrCodeUpdateCalcModelRelatedTaskStatus                      // 更新计算模型关联任务状态失败
	ErrCodeNoAnyCalcModel                                        // 该成员未拥有任何计算模型
	ErrCodeNoAppIdInCookie                                       // cookie里没有appid
	ErrCodeStatAppIdDir                                          // 该用户存储模型文件目录错误
	ErrCodeCreateAppIdDir                                        // 创建用户存储模型文件目录错误
	ErrCodeAppIdPathIsNotDir                                     // AppId路径不是目录
	ErrCodeFormHasNoModelFile                                    // 上传模型文件时的form-data没有名称为calc-model-file的文件
	ErrCodeSaveCalcModelFile                                     // 存储模型文件失败
	ErrCodeReadCalcModelFile                                     // 读取模型文件失败
	ErrCodeUploadCalcModelFile                                   // 上传模型文件失败
	ErrCodeDownloadCalcModelFile                                 // 下载模型文件失败
	ErrCodeUnmarshalCalcModelFileUploadReq                       // 解析上传模型文件请求失败
	ErrCodeUnmarshalCalcModelFileUploadRes                       // 解析上传模型文件响应失败
	ErrCodeAlreadyAppliedCalcModelAcl                            // 重复申请计算模型授权
	ErrCodeAddCalcModelAcl                                       // 申请计算模型授权请求上链失败
	ErrCodeCreateAuthorizeCalcModelEvent                         // 创建申请计算模型授权请求事件失败
	ErrCodeCalcModelNotRelease                                   // 计算模型未发布
)

const (
	// 数据源管理错误码: 400~499
	ErrCodeDescDataSourceDB                  ErrCode = iota + 400 // 获取数据源DB表结构失败
	ErrCodeCheckDataSourceDB                                      // 数据源数据库校验失败
	ErrCodeGetDataSource                                          // 从链上获取数据源信息失败
	ErrCodeDataSourceExists                                       // 数据源已存在
	ErrCodeAddDataSource                                          // 数据源上链失败
	ErrCodeUpdateDataSourceRelatedTaskStatus                      // 更新数据源关联任务状态失败
	ErrCodeGetDataSourceList                                      // 从链上获取数据源信息列表失败
	ErrCodeNoAnyDataSource                                        // 该成员未拥有任何数据源
	ErrCodeDataSourceNotRelease                                   // 数据源未发布
	ErrCodeUpdateDataSource                                       // 更新数据源上链失败
	ErrCodeDeleteDataSource                                       // 删除数据源失败
	ErrCodeReleaseDataSource                                      // 发布数据源上链失败
	ErrCodeApplyDataSourceAuth                                    // 申请数据源授权上链失败
)

const (
	// 计算任务管理错误码: 500~599
	ErrCodeAddCalcTask                          ErrCode = iota + 500 // 创建计算任务失败
	ErrCodeUpdateCalcTask                                            // 更新计算任务失败
	ErrCodeGetCalcTask                                               // 从链上获取计算任务详情失败
	ErrCodeGetCalcTaskList                                           // 从链上获取计算任务列表信息失败
	ErrCodeAuthorizeCalcTaskEvent                                    // 创建授权计算任务邀请事件失败
	ErrCodeRunSubmitCmd                                              // 计算任务提交失败
	ErrCodeAddCalcTaskInstance                                       // 创建计算任务实例失败
	ErrCodeCheckResultDB                                             // 计算结果接收DB探测失败
	ErrCodeGetCalcTaskInstanceListByNoAndStatus                      // 获取任务编号获取计算任务实例列表失败
	ErrCodeCalcTaskInstanceAlreadyRunning                            // 计算任务实例已经启动
	ErrCodeCalcTaskInstanceAlreadyCreated                            // 计算任务实例已经创建
	ErrCodeNoAnyCalcTask                                             // 该成员未创建任何计算任务
	ErrCodeCreateCalcTaskInfo                                        // 创建计算任务包体失败
	ErrCodeCreateTaskRequest                                         // 创建可信计算任务请求失败
	ErrCodeAddCalcTaskRequest                                        // 可信计算任务请求上链失败
	ErrCodeCreateAuthorizeCalcTaskEvent                              // 创建计算任务授权事件失败
	ErrCodeCalcTaskNotExists                                         // 计算任务不存在
	ErrCodeCalcTaskExists                                            // 计算任务已存在
	ErrCodeGetRunningCalcTaskInstanceCount                           // 获取计算资源运行中的任务实例数量失败
	ErrCodeStopCalcTask                                              // 终止计算任务失败
	ErrCodeCalcTaskUnauth                                            // 计算任务未授权完成
	ErrCodeCheckConsMemberValid                                      // 检查计算任务参与方有效性失败
	ErrCodeCalcTaskInvalidCreator                                    // 无效计算任务创建方
	ErrCodeCalcTaskInvalidCalcModelProvider                          // 无效计算计算模型提供方
	ErrCodeCalcTaskInvalidDataSourceProvider                         // 无效计算数据目录提供方
	ErrCodeCalcTaskInvalidResultUser                                 // 无效计算结果使用方
)

const (
	// 事件中心管理错误码: 600~699
	ErrCodeEventExists             = iota + 600 // 事件已存在
	ErrCodeEventNotExists                       // 事件不存在
	ErrCodeEventTypeNotMatch                    // 事件类型不匹配
	ErrCodeGetEventByEventId                    // 根据事件ID获取事件失败
	ErrCodeGetEventList                         // 获取事件列表失败
	ErrCodeEventStepNotExists                   // 事件步骤不存在
	ErrCodeEventStepExists                      // 事件步骤已存在
	ErrCodeEventStepNotMatch                    // 事件步骤不匹配
	ErrCodeGetEventStep                         // 获取事件步骤失败
	ErrCodeGetEventStepList                     // 获取事件步骤列表失败
	ErrCodeEventTaskNotExists                   // 事件任务不存在
	ErrCodeEventTaskExists                      // 事件任务已存在
	ErrCodeEventTaskTypeNotMatch                // 事件任务类型不匹配
	ErrCodeGetEventTask                         // 获取事件任务失败
	ErrCodeGetEventTaskList                     // 获取事件任务列表失败
	ErrcodeEventGetTaskCount                    // 获取事件任务数量失败
	ErrCodeEventTaskStatusNotMatch              // 事件任务状态不匹配
	ErrCodeHandleEventTask                      // 处理事件任务失败
)

const (
	// 资源管理相关错误码：700+
	ErrCodeAddTeeResource                 ErrCode = iota + 700 // 创建可信计算资源失败
	ErrCodeGetTeeResourceInfoByName                            // 根据可信计算资源名称查详情失败
	ErrCodeGetTeeResourceInfoById                              // 根据可信计算资源ID查详情失败
	ErrCodeTeeResourceNotOwner                                 // 非计算资源所有者
	ErrCodeTeeResourceExists                                   // 可信计算资源已存在
	ErrCodeTeeResourceNotExists                                // 计算资源不存在
	ErrCodeGetTeeResourceList                                  // 获取计算资源列表失败
	ErrCodeUpdateTeeResource                                   // 更新计算资源信息失败
	ErrCodeAddChainConf                                        // 创建链配置失败
	ErrCodeGetChainConfInfoByName                              // 根据链配置名称查详情失败
	ErrCodeGetChainConfInfoById                                // 根据链配置ID查详情失败
	ErrCodeChainConfExists                                     // 链配置已存在
	ErrCodeChainConfNotExists                                  // 链配置不存在
	ErrCodeChainConfNotOwner                                   // 非链配置所有者
	ErrCodeGetChainConfList                                    // 获取链配置列表失败
	ErrCodeUpdateChainConf                                     // 更新链配置信息失败
	ErrCodeApplyEncryptKey                                     // 申请数据加密公钥失败
	ErrCodeGetChainDetailByChainNetworkId                      // 根据链网络ID获取链详情信息失败
	ErrCodeGetTKMSService                                      // 获取TKMS服务失败
	ErrCodeGetUsedMemSize                                      // 获取计算资源已使用安全内存大小失败
)

const (
	// 数据连接相关错误码：800+
	ErrCodeCheckDBConnect              ErrCode = iota + 800 // 数据库连接探测失败
	ErrCodeAddDataConnect                                   // 创建数据连接失败
	ErrCodeUpdateDataConnect                                // 更新数据连接失败
	ErrCodeUpdateDataConnectStatus                          // 更新数据连接状态失败
	ErrCodeUnknownDataConnectStatus                         // 未知数据连接状态
	ErrCodeDataConnectEnabled                               // 数据连接处于启用状态，无法修改
	ErrCodeGetDataConnectInfoByName                         // 根据名称获取数据连接信息失败
	ErrCodeGetDataConnectInfoById                           // 根据ID获取数据连接信息失败
	ErrCodeGetDataConnectList                               // 获取数据连接列表失败
	ErrCodeDataConnectExists                                // 数据连接已存在
	ErrCodeDataConnectNotExists                             // 数据连接不存在
	ErrCodeDataConnectNotOwner                              // 非数据连接所有者
	ErrCodeUpdateDataConnectSameStatus                      // 相同连接状态，无需修改
)

const (
	//账户相关错误码：900+
	ErrCodeInvalidUserName           ErrCode = iota + 900 //登录账号格式错误
	ErrCodeUserNameExist                                  //登录账号已存在
	ErrCodeInvalidPassword                                //登录密码错误
	ErrCodeInvalidNickName                                //用户昵称格式错误
	ErrCodeInvalidPhone                                   //用户手机号码格式错误
	ErrCodeInvalidEmail                                   //用户邮箱地址格式错误
	ErrCodeInvalidUserStatus                              //用户状态异常
	ErrCodeUserResourceExist                              //用户还存在未退出的联盟，无法删除
	ErrCodeUserPreset                                     //预设账号无法被修改
	ErrCodeGetUserAccount                                 //获取用户账号信息失败
	ErrCodeUserNumberLimit                                //用户数已达上限
	ErrCodeUserInfoNoChange                               //用户信息修改失败，修改项未发生变化
	ErrCodeInvalidUserNameOrPassword                      //用户名或密码不正确
	ErrCodeSessionExpired                                 //会话已失效，请重新登录
	ErrCodeSessionInvalidAppId                            //会话AppId不匹配
	ErrCodeNoSupportAction                                //平台不支持该接口
	ErrCodeUserNoBindGroup                                //用户未绑定用户组
	ErrCodeGroupNoBindPolicy                              //用户组未绑定策略
	ErrCodeActionAuthFailure                              //用户无接口访问权限
	ErrCodeSecretNoExist                                  //密钥不存在或已失效
	ErrCodeInvalidUserGroupStatus                         //用户组状态异常
	ErrCodeUserAlreadyLocked                              //用户已锁定
)

const (
	// 链交互相关: 1000+
	ErrCodeGetTBCCChainService      ErrCode = iota + 1000 // 获取TBCC管控台链服务失败
	ErrCodeGetMemberChainService                          // 获取联盟成员链访问服务失败
	ErrCodeCreateMemberChainService                       // 创建联盟成员链访问服务失败
)

const (
	// 其他: 2000+
	ErrCodeApplyChainAccessService ErrCode = iota + 2000
	ErrCodeAccountNotExists                // 该账号不存在
	ErrCodeAccountNicknameNotMatch         // 账号和昵称不匹配
	ErrCodePanic
)

// ErrMessage - err msg
var ErrMessage = map[ErrCode][]string{
	// 系统缺省错误: 100~199
	ErrCodeInvalidParameter:      {"InvalidParameter", "invalid parameter", "参数错误"},
	ErrCodeInvalidParameterValue: {"InvalidParameterValue", "invalid parameter value", "参数取值错误"},
	ErrCodeMissingParameter:      {"MissingParameter", "missing parameter", "必传参数没填"},
	ErrCodeUnknownParameter:      {"UnknownParameter", "unknown parameter", "未知参数错误"},
	ErrCodeInternalError:         {"InternalError", "service internal error", "服务内部错误"},

	// 联盟及成员管理错误码: 200~299
	ErrCodeAddConsortium:                     {"FailedOperation.AddConsortium", "add consortium failed", "联盟信息上链失败"},
	ErrCodeUpdateConsortium:                  {"FailedOperation.UpdateConsortium", "update consortium failed", "更新联盟信息上链失败"},
	ErrCodeGetConsortiumByName:               {"FailedOperation.GetConsortiumByName", "get consortium by name failed", "根据联盟名称获取联盟详情失败"},
	ErrCodeConsortiumNotExists:               {"FailedOperation.ConsortiumNotExists", "consortium not exists", "指定联盟不存在"},
	ErrCodeBindConsMemberCrt:                 {"FailedOperation.BindConsMemberCrt", "bind consortium member certs failed", "绑定联盟成员证书失败"},
	ErrCodeBindConsMemberEncryptPk:           {"FailedOperation.BindConsMemberEncryptPk", "bind consortium member encrypt pk failed", "绑定联盟成员加密公钥失败"},
	ErrCodeUpdateConsMemberStatus:            {"FailedOperation.UpdateConsMemberStatus", "update consortium member status failed", "更新联盟成员状态失败"},
	ErrCodeCreateInviteMemberEvent:           {"FailedOperation.CreateInviteMemberEvent", "create invite member event failed", "创建联盟成员邀请事件失败"},
	ErrCodeGetConsMemberByAppId:              {"FailedOperation.GetConsMemberByAppId", "get consortium member by appId from chain failed", "根据AppId从链上获取联盟成员失败"},
	ErrCodeMemberAlreadyBelongsToCons:        {"FailedOperation.MemberAlreadyBelongsToCons", "member already belongs to this consortium", "成员已归属该联盟"},
	ErrCodeMemberNotBelongsToAnyCons:         {"FailedOperation.MemberNotBelongsToAnyCons", "member not belongs to any consortium", "成员未归属任何联盟"},
	ErrCodeMemberNotBelongsToThisCons:        {"FailedOperation.MemberNotBelongsToThisCons", "member not belongs to this consortium", "成员未归属该联盟"},
	ErrCodeGetConsortiumListForMember:        {"FailedOperation.GetConsortiumsForMember", "get member consortium list failed", "获取成员所在联盟列表失败"},
	ErrCodeGetConsortiumMembers:              {"FailedOperation.GetConsortiumMembers", "get consortium members failed", "获取联盟拥有成员信息失败"},
	ErrCodeGetMemberConsNames:                {"FailedOperation.GetMemberConsNames", "get member consortium name list failed", "获取成员所属联盟名称列表失败"},
	ErrCodeGetJoinedMemberList:               {"FailedOperation.GetJoinedMemberList", "get consortium joined member list failed", "获取联盟已加入成员列表失败"},
	ErrCodeGetInvitedMemberList:              {"FailedOperation.GetInvitedMemberList", "get consortium invited member list failed", "获取联盟邀请中成员列表失败"},
	ErrCodeGetInvitingMemberList:             {"FailedOperation.GetInvitingMemberList", "get consortium inviting member list failed", "获取邀请中的联盟成员列表失败"},
	ErrCodeMemberAlreadyInInviting:           {"FailedOperation.MemberAlreadyInInviting", "member already in inviting", "联盟成员已经在邀请中"},
	ErrCodeExitConsortium:                    {"FailedOperation.ExitConsortium", "member exit failed", "成员退出联盟失败"},
	ErrCodeDownloadMemberCert:                {"FailedOperation.DownloadMemberCert", "download member crt failed", "下载联盟成员托管私钥和证书失败"},
	ErrCodeDownloadMemberEncryptPk:           {"FailedOperation.DownloadMemberEncryptPk", "download member encrypt pk failed", "下载联盟成员数据加密公钥失败"},
	ErrCodeInvalidMemberRole:                 {"FailedOperation.InvalidMemberRole", "invalid member role", "成员不具备相关角色"},
	ErrCodeConsortiumExists:                  {"FailedOperation.ConsortiumExists", "consortium exists", "联盟已存在"},
	ErrCodeApplyCert:                         {"FailedOperation.ApplyCert", "apply cert failed", "获取联盟成员链身份失败"},
	ErrCodeGetMemberCertKeyInfoByAppIdFromDB: {"FailedOperation.GetMemberCertKeyInfoByAppIdFromDB", "get member cert key info by appId from DB failed", "从DB获取成员链身份及数据加密公钥失败"},
	ErrCodeInsertMemberCertKeyInfoToDB:       {"FailedOperation.InsertMemberCertKeyInfoToDB", "insert member cert key info to DB failed", "成员链身份及数据加密公钥新增入库失败"},
	ErrCodeGetUserNicknameByAppId:            {"FailedOperation.GetUserNicknameByAppId", "get user nickname by appId failed", "根据AppId获取用户昵称失败"},
	ErrCodeGetMyTeeResourceEncryptPk:         {"FailedOperation.GetMyTeeResourceEncryptPk", "get my tee resource encrypt pk failed", "获取自己可信计算加密公钥失败"},
	ErrCodeAddMemberSignCrt:                  {"FailedOperation.AddMemberSignCrt", "add member sign crt failed", "联盟成员链身份证书上链失败"},
	ErrCodeAddMemberTeeResourceEncryptPkList: {"FailedOperation.AddMemberTeeResourceEncryptPkList", "add member tee resource encrypt pk list failed", "联盟成员计算资源加密证书上链失败"},
	ErrCodeConsortiumNotOwner:                {"FailedOperation.ConsortiumNotOwner", "not consortium owner", "非联盟所有者"},
	ErrCodeConsMemberExit:                    {"FailedOperation.ConsMemberExit", "member is already exited", "成员已退出"},

	// 计算模型管理错误码: 300~399
	ErrCodeGetCalcModel:                     {"FailedOperation.GetCalcModel", "get calc model from chain failed", "获取计算模型信息失败"},
	ErrCodeGetCOSPresignedURL:               {"FailedOperation.GetCOSPresignedURL", "get cos presigned url failed", "生成计算模型COS预签名URL失败"},
	ErrCodeInvalidAppId:                     {"FailedOperation.InvalidAppId", "invalid appId", "无效的AppId"},
	ErrCodeNotVisible:                       {"FailedOperation.NotVisible", "calc model is not visible for appid", "模型不可见"},
	ErrCodeInvalidCalcModelNo:               {"FailedOperation.Invalid calcModelNo", "invalid calcModelNo", "无效的模型编号"},
	ErrCodeCalcModelExists:                  {"FailedOperation.CalcModelExists", "calc model exists", "计算模型已存在"},
	ErrCodeGetCalcModelContent:              {"FailedOperation.GetCalcModelContent", "get calc model content failed", "获取计算模型内容失败"},
	ErrCodeReleaseStoreCalcModel:            {"FailedOperation.ReleaseStoreCalcModel", "release calc model to store failed", "发布计算模型到模型仓库失败"},
	ErrCodeUpdateReleaseStatusCalcModel:     {"FailedOperation.ErrCodeUpdateReleaseStatusCalcModel", "update release status of calc model failed", "更新计算模型发布状态失败"},
	ErrCodeReleaseCalcModel:                 {"FailedOperation.ReleaseCalcModel", "release calc model to chain failed", "发布计算模型上链失败"},
	ErrCodeUpdateCalcModel:                  {"FailedOperation.UpdateCalcModel", "update calc model failed", "更新计算模型失败"},
	ErrCodeDeleteCalcModel:                  {"FailedOperation.DeleteCalcModel", "delete calc model failed", "删除计算模型失败"},
	ErrCodeGetCalcModelList:                 {"FailedOperation.GetCalcModelList", "get calc model list from chain failed", "从链上获取计算联盟列表信息失败"},
	ErrCodeCalcPkgChecksum:                  {"FailedOperation.CalcPkgChecksum", "calc pkg checksum failed", "计算模型checksum失败"},
	ErrCodeUpdateCalcModelRelatedTaskStatus: {"FailedOperation.UpdateCalcModelRelatedTaskStatus", "update calc model related task status failed", "更新计算模型关联任务状态失败"},
	ErrCodeNoAnyCalcModel:                   {"FailedOperation.NoAnyCalcModel", "no any calc model", "该成员未拥有任何计算模型"},
	ErrCodeNoAppIdInCookie:                  {"FailedOperation.NoAppIdInCookie", "no appId in cookie", "cookie里没有appId"},
	ErrCodeStatAppIdDir:                     {"FailedOperation.StatAppIdDir", "stat appId dir failed", "该用户存储模型文件目录错误"},
	ErrCodeCreateAppIdDir:                   {"FailedOperation.CreateAppIdDir", "create appId dir failed", "创建用户存储模型文件目录错误"},
	ErrCodeAppIdPathIsNotDir:                {"FailedOperation.AppIdPathIsNotDir", "appId dir is not a dir", "AppId路径不是目录"},
	ErrCodeFormHasNoModelFile:               {"FailedOperation.FormHasNoModelFile", "form has no calc-model-file", "上传模型文件时的form-data没有名称为calc-model-file的form-data"},
	ErrCodeSaveCalcModelFile:                {"FailedOperation.SaveCalcModelFile", "save calc model file failed", "存储模型文件失败"},
	ErrCodeReadCalcModelFile:                {"FailedOperation.ReadCalcModelFile", "read calc model file failed", "读取模型文件失败"},
	ErrCodeUploadCalcModelFile:              {"FailedOperation.UploadCalcModelFile", "upload calc model file failed", "上传模型文件失败"},
	ErrCodeDownloadCalcModelFile:            {"FailedOperation.DownloadCalcModelFile", "download calc model file failed", "下载模型文件失败"},
	ErrCodeUnmarshalCalcModelFileUploadReq:  {"FailedOperation.UnmarshalCalcModelFileUploadReq", "unmarshal calc model file upload request failed", "解析上传模型文件请求失败"},
	ErrCodeUnmarshalCalcModelFileUploadRes:  {"FailedOperation.UnmarshalCalcModelFileUploadRes", "unmarshal calc model file upload response failed", "解析上传模型文件响应失败"},
	ErrCodeAlreadyAppliedCalcModelAcl:       {"FailedOperation.AlreadyAppliedCalcModelAcl", "already applied calc model file failed", "已经申请过计算模型权限"},
	ErrCodeAddCalcModelAcl:                  {"FailedOperation.ErrCodeAddCalcModelAcl", "add calc model acl to chain failed", "申请计算模型授权请求上链失败"},
	ErrCodeCreateAuthorizeCalcModelEvent:    {"FailedOperation.ErrCodeCreateAuthorizeCalcModelEvent", "create authorize calc model failed", "创建申请计算模型授权请求事件失败"},
	ErrCodeCalcModelNotRelease:              {"FailedOperation.CalcModelNotRelease", "calc model not release", "计算模型未发布"},

	// 数据源管理错误码: 400~499
	ErrCodeDescDataSourceDB:                  {"FailedOperation.DescDataSourceDB", "descirbe data source db struct failed", "获取数据源DB表结构失败"},
	ErrCodeCheckDataSourceDB:                 {"FailedOperation.CheckDataSourceDB", "check data source db failed", "数据源数据库校验失败"},
	ErrCodeGetDataSource:                     {"FailedOperation.GetDataSource", "get data source from chain failed", "从链上获取数据源信息失败"},
	ErrCodeDataSourceExists:                  {"FailedOperation.DataSourceExists", "data source exists", "数据源已存在"},
	ErrCodeAddDataSource:                     {"FailedOperation.AddDataSource", "add data source to chain failed", "数据源上链失败"},
	ErrCodeUpdateDataSourceRelatedTaskStatus: {"FailedOperation.UpdateDataSourceRelatedTaskStatus", "update data source related task status failed", "更新数据源关联任务状态失败"},
	ErrCodeGetDataSourceList:                 {"FailedOperation.GetDataSourceList", "get data source list from chain failed", "从链上获取数据源信息列表失败"},
	ErrCodeNoAnyDataSource:                   {"FailedOperation.NoAnyDataSource", "no any data source", "该成员未拥有任何数据源"},
	ErrCodeDataSourceNotRelease:              {"FailedOperation.DataSourceNotRelease", "data source not release", "数据源未发布"},
	ErrCodeUpdateDataSource:                  {"FailedOperation.UpdateDataSource", "update data source failed", "更新数据源失败"},
	ErrCodeDeleteDataSource:                  {"FailedOperation.DeleteDataSource", "delete data source failed", "删除数据源失败"},
	ErrCodeReleaseDataSource:                 {"FailedOperation.ReleaseDataSource", "release data source failed", "发布数据源失败"},
	ErrCodeApplyDataSourceAuth:               {"FailedOperation.ApplyDataSourceAuth", "apply data source auth failed", "申请数据源授权失败"},

	// 计算任务管理错误码: 500~599
	ErrCodeAddCalcTask:                          {"FailedOperation.AddCalcTask", "add calc task failed", "创建计算任务失败"},
	ErrCodeUpdateCalcTask:                       {"FailedOperation.UpdateCalcTask", "update calc task failed", "更新计算任务失败"},
	ErrCodeGetCalcTask:                          {"FailedOperation.GetCalcTask", "get calc task failed", "从链上获取计算任务详情失败"},
	ErrCodeGetCalcTaskList:                      {"FailedOperation.GetCalcTaskList", "get calc task list failed", "从链上获取计算任务列表信息失败"},
	ErrCodeAuthorizeCalcTaskEvent:               {"FailedOperation.AuthorizeCalcTaskEvent", "authorize calc task event failed", "创建授权计算任务邀请事件失败"},
	ErrCodeRunSubmitCmd:                         {"FailedOperation.RunSubmitCmd", "run submit cmd failed", "计算任务提交失败"},
	ErrCodeAddCalcTaskInstance:                  {"FailedOperation.AddCalcTaskInstance", "add calc task instance failed", "创建计算任务实例失败"},
	ErrCodeCheckResultDB:                        {"FailedOperation.CheckResultDB", "check result db failed", "计算结果接收DB探测失败"},
	ErrCodeGetCalcTaskInstanceListByNoAndStatus: {"FailedOperation.GetCalcTaskInstanceListByNoAndStatus", "get calc task instance list by no and status failed", "获取任务编号获取计算任务实例列表失败"},
	ErrCodeCalcTaskInstanceAlreadyRunning:       {"FailedOperation.CalcTaskInstanceAlreadyRunning", "calc task instance already running", "计算任务实例已经启动"},
	ErrCodeCalcTaskInstanceAlreadyCreated:       {"FailedOperation.CalcTaskInstanceAlreadyCreated", "calc task instance already created", "计算任务实例已经创建"},
	ErrCodeNoAnyCalcTask:                        {"FailedOperation.NoAnyCalcTask", "no any calc task", "该成员未创建任何计算任务"},
	ErrCodeCreateCalcTaskInfo:                   {"FailedOperation.CreateCalcTaskInfo", "create calc task info failed", "创建计算任务包体失败"},
	ErrCodeCreateTaskRequest:                    {"FailedOperation.CreateTaskRequest", "create calc task request failed", "创建可信计算任务请求失败"},
	ErrCodeAddCalcTaskRequest:                   {"FailedOperation.AddCalcTaskRequest", "add calc task request on chain failed", "可信计算任务请求上链失败"},
	ErrCodeCreateAuthorizeCalcTaskEvent:         {"FailedOperation.CreateAuthorizeCalcTaskEvent", "create authorize calc task event failed", "创建计算任务授权事件失败"},
	ErrCodeCalcTaskNotExists:                    {"FailedOperation.CalcTaskNotExists", "calc task not exists", "计算任务不存在"},
	ErrCodeCalcTaskExists:                       {"FailedOperation.CalcTaskExists", "calc task exists", "计算任务已存在"},
	ErrCodeGetRunningCalcTaskInstanceCount:      {"FailedOperation.GetRunningCalcTaskInstanceCount", "get running calc task instance count", "获取计算资源运行中的任务实例数量失败"},
	ErrCodeStopCalcTask:                         {"FailedOperation.StopCalcTask", "stop calc task failed", "终止计算任务失败"},
	ErrCodeCalcTaskUnauth:                       {"FailedOperation.CalcTaskUnauth", "calc task unauth", "计算任务未授权完成"},
	ErrCodeCheckConsMemberValid:                 {"FailedOperation.CheckConsMemberValid", "check calc task member valid", "检查计算任务参与方有效性失败"},
	ErrCodeCalcTaskInvalidCreator:               {"FailedOperation.CalcTaskInvalidCreator", "calc task invalid creator", "无效计算任务创建方"},
	ErrCodeCalcTaskInvalidCalcModelProvider:     {"FailedOperation.CalcTaskInvalidCalcModelProvider", "calc task invalid calc model provider", "无效计算计算模型提供方"},
	ErrCodeCalcTaskInvalidDataSourceProvider:    {"FailedOperation.CalcTaskInvalidDataSourceProvider", "calc task invalid data source provider", "无效计算数据目录提供方"},
	ErrCodeCalcTaskInvalidResultUser:            {"FailedOperation.ErrCodeCalcTaskInvalidResultUser", "calc task invalid result user", "无效计算结果使用方"},

	// 事件中心管理错误码: 600~699
	ErrCodeEventExists:             {"FailedOperation.EventExists", "event exists", "事件已存在"},
	ErrCodeEventNotExists:          {"FailedOperation.EventNotExists", "event not exists", "事件不存在"},
	ErrCodeEventTypeNotMatch:       {"FailedOperation.EventTypeNotMatch", "event type not match", "事件类型不匹配"},
	ErrCodeGetEventByEventId:       {"FailedOperation.GetEventByEventId", "get event by event id failed", "根据事件ID获取事件失败"},
	ErrCodeGetEventList:            {"FailedOperation.GetEventList", "get event list failed", "获取事件列表失败"},
	ErrCodeEventStepExists:         {"FailedOperation.EventStepExists", "event step exists", "事件步骤已存在"},
	ErrCodeEventStepNotExists:      {"FailedOperation.EventStepNotExists", "event step not exists", "事件步骤不存在"},
	ErrCodeEventStepNotMatch:       {"FailedOperation.EventStepNotMatch", "event step not match", "事件步骤不匹配"},
	ErrCodeGetEventStepList:        {"FailedOperation.GetEventStepList", "get event step list failed", "获取事件步骤列表失败"},
	ErrCodeGetEventStep:            {"FailedOperation.GetEventStep", "get event step failed", "获取事件步骤失败"},
	ErrCodeEventTaskExists:         {"FailedOperation.EventTaskExists", "event task exists", "事件任务已存在"},
	ErrCodeEventTaskNotExists:      {"FailedOperation.EventTaskNotExists", "event task not exists", "事件任务不存在"},
	ErrCodeEventTaskTypeNotMatch:   {"FailedOperation.EventTaskTypeNotMatch", "event task type not match", "事件任务类型不匹配"},
	ErrCodeGetEventTask:            {"FailedOperation.GetEventTask", "get event task failed", "获取事件任务失败"},
	ErrCodeGetEventTaskList:        {"FailedOperation.EventTaskList", "get event task list failed", "获取事件任务列表失败"},
	ErrcodeEventGetTaskCount:       {"FailedOperation.EventGetTaskCount", "get event task count failed", "获取事件任务数量失败"},
	ErrCodeEventTaskStatusNotMatch: {"FailedOperation.EventTaskStatusNotMatch", "event task status not match", "事件任务状态不匹配"},
	ErrCodeHandleEventTask:         {"FailedOperation.HandleEventTask", "handle event task failed", "处理事件任务失败"},

	// 资源管理相关错误码：700+
	ErrCodeAddTeeResource:                 {"FailedOperation.AddTeeResource", "add tee resource failed", "创建可信计算资源失败"},
	ErrCodeTeeResourceExists:              {"FailedOperation.TeeResourceExists", "tee resource already exists", "可信计算资源已存在"},
	ErrCodeTeeResourceNotExists:           {"FailedOperation.TeeResourceNotExists", "tee resource not exists", "可信计算资源不存在"},
	ErrCodeGetTeeResourceInfoByName:       {"FailedOperation.GetTeeResourceInfoByName", "get tee resource info by name failed", "根据可信计算资源名称查详情失败"},
	ErrCodeGetTeeResourceInfoById:         {"FailedOperation.GetTeeResourceInfoById", "get tee resource info by id failed", "根据可信计算资源ID查详情失败"},
	ErrCodeTeeResourceNotOwner:            {"FailedOperation.TeeResourceNotOwner", "not tee resource owner", "非计算资源所有者"},
	ErrCodeGetTeeResourceList:             {"FailedOperation.GetTeeResourceList", "get tee resource list failed", "获取计算资源列表失败"},
	ErrCodeUpdateTeeResource:              {"FailedOperation.UpdateTeeResource", "update tee resource failed", "更新计算资源信息失败"},
	ErrCodeAddChainConf:                   {"FailedOperation.AddChainConf", "add chain conf failed", "创建链配置失败"},
	ErrCodeGetChainConfInfoByName:         {"FailedOperation.GetChainConfInfoByName", "get chain conf info by name failed", "根据链配置名称查详情失败"},
	ErrCodeGetChainConfInfoById:           {"FailedOperation.GetChainConfInfoById", "get chain conf info by id failed", "根据链配置ID查详情失败"},
	ErrCodeChainConfExists:                {"FailedOperation.ChainConfExists", "chain conf already exists", "链配置已存在"},
	ErrCodeChainConfNotExists:             {"FailedOperation.ChainConfNotExists", "chain conf not exists", "链配置不存在"},
	ErrCodeChainConfNotOwner:              {"FailedOperation.ChainConfNotOwner", "not chain conf owner", "非链配置所有者"},
	ErrCodeGetChainConfList:               {"FailedOperation.GetChainConfList", "get chain conf list failed", "获取链配置列表失败"},
	ErrCodeUpdateChainConf:                {"FailedOperation.UpdateChainConf", "update chain conf info failed", "更新链配置信息失败"},
	ErrCodeApplyEncryptKey:                {"FailedOperation.ApplyEncryptKey", "apply tkms encrypt key failed", "申请数据加密公钥失败"},
	ErrCodeGetChainDetailByChainNetworkId: {"FailedOperation.GetChainDetailByChainNetworkId", "get chain detail by chain network id failed", "根据链网络ID获取链详情信息失败"},
	ErrCodeGetTKMSService:                 {"FailedOperation.GetTKMSService", "get TKMS service failed", "获取TKMS服务失败"},
	ErrCodeGetUsedMemSize:                 {"FailedOperation.GetUsedMemSize", "get use mem size failed", "获取计算资源已使用安全内存大小失败"},

	// 数据连接相关错误码：800+
	ErrCodeCheckDBConnect:              {"FailedOperation.CheckDBConnect", "check db connect failed", "数据库连接探测失败"},
	ErrCodeAddDataConnect:              {"FailedOperation.AddDataConnect", "add data connect failed", "创建数据连接失败"},
	ErrCodeUpdateDataConnect:           {"FailedOperation.UpdateDataConnect", "update data connect failed", "更新数据连接失败"},
	ErrCodeUpdateDataConnectStatus:     {"FailedOperation.UpdateDataConnectStatus", "update data connect status failed", "更新数据连接状态失败"},
	ErrCodeUnknownDataConnectStatus:    {"FailedOperation.UnknownDataConnectStatus", "unknown update data connect status", "未知数据连接状态"},
	ErrCodeDataConnectEnabled:          {"FailedOperation.DataConnectEnabled", "can't update ENABLE data connect failed", "数据连接处于启用状态，无法修改"},
	ErrCodeGetDataConnectInfoByName:    {"FailedOperation.GetDataConnectInfoByName", "get data connect info by name failed", "根据名称获取数据连接信息失败"},
	ErrCodeGetDataConnectInfoById:      {"FailedOperation.GetDataConnectInfoById", "get data connect info by id failed", "根据ID获取数据连接信息失败"},
	ErrCodeGetDataConnectList:          {"FailedOperation.GetDataConnectList", "get data connect list failed", "获取数据连接列表失败"},
	ErrCodeDataConnectExists:           {"FailedOperation.DataConnectExists", "data connect info exists", "数据连接已存在"},
	ErrCodeDataConnectNotExists:        {"FailedOperation.DataConnectNotExists", "data connect info not exists", "数据连接不存在"},
	ErrCodeDataConnectNotOwner:         {"FailedOperation.DataConnectNotOwner", "not data connect owner", "非数据连接所有者"},
	ErrCodeUpdateDataConnectSameStatus: {"FailedOperation.UpdateDataConnectSameStatus", "same data connect status", "相同连接状态，无需修改"},

	//账户相关错误码：900+
	ErrCodeInvalidUserName:   {"FailedOperation.InvalidUserName", "invalid user name", "登录账号格式错误"},
	ErrCodeUserNameExist:     {"FailedOperation.UserNameExist", "user name already existed", "登录账号已存在"},
	ErrCodeInvalidPassword:   {"FailedOperation.PasswordError", "invalid password", "登录密码错误"},
	ErrCodeInvalidNickName:   {"FailedOperation.InvalidNickName", "invalid nick name", "用户昵称格式错误"},
	ErrCodeInvalidPhone:      {"FailedOperation.InvalidPhone", "invalid phone", "用户手机号码格式错误"},
	ErrCodeInvalidEmail:      {"FailedOperation.InvalidEmail", "invalid email", "用户邮箱地址格式错误"},
	ErrCodeInvalidUserStatus: {"FailedOperation.InvalidUserStatus", "invalid user status", "用户状态异常"},
	ErrCodeUserResourceExist: {"FailedOperation.UserResourceExist", "user resource still existed, delete failed",
		"用户还存在未退出的联盟，无法删除"},
	ErrCodeUserPreset:      {"FailedOperation.UserPreset", "user is preset, modify failed", "预设账号无法被修改"},
	ErrCodeGetUserAccount:  {"FailedOperation.GetUserAccount", "get user account failed", "获取用户账号信息失败"},
	ErrCodeUserNumberLimit: {"FailedOperation.UserNumberLimit", "user number limited", "用户数已达上限"},
	ErrCodeUserInfoNoChange: {"FailedOperation.UserInfoNoChange", "user info not change, modify failed",
		"用户信息修改失败，修改项未发生变化"},
	ErrCodeActionAuthFailure: {"FailedOperation.InvalidAuth", "action auth failure", "用户无接口访问权限"},
	ErrCodeInvalidUserNameOrPassword: {"FailedOperation.InvalidUserNameOrPassword", "invalid username or password",
		"用户名或密码不正确"},
	ErrCodeSessionExpired:         {"FailedOperation.SessionExpired", "session expired", "会话已失效，请重新登录"},
	ErrCodeSessionInvalidAppId:    {"FailedOperation.ErrCodeSessionInvalidAppId", "session invalid appId", "会话AppId不匹配"},
	ErrCodeNoSupportAction:        {"FailedOperation.NoSupportAction", "action not supported", "平台不支持该接口"},
	ErrCodeUserNoBindGroup:        {"FailedOperation.UserNoBindGroup", "user not bind group", "用户未绑定用户组"},
	ErrCodeGroupNoBindPolicy:      {"FailedOperation.GroupNoBindPolicy", "group not bind policy", "用户组未绑定策略"},
	ErrCodeSecretNoExist:          {"FailedOperation.SecretNoExist", "secret not exist", "密钥不存在或已失效"},
	ErrCodeInvalidUserGroupStatus: {"FailedOperation.InvalidUserGroupStatus", "invalid user group status", "用户组状态异常"},
	ErrCodeUserAlreadyLocked:      {"FailedOperation.UserAlreadyLocked", "user has been locked", "用户已被锁定"},

	// 其他: 1000+
	ErrCodeGetTBCCChainService:      {"FailedOperation.GetTBCCChainService", "get TBCC chain service failed", "获取TBCC链服务失败"},
	ErrCodeGetMemberChainService:    {"FailedOperation.GetMemberChainService", "get member chain service failed", "获取联盟成员链访问服务失败"},
	ErrCodeCreateMemberChainService: {"FailedOperation.CreateMemberChainService", "create member chain service failed", "创建联盟成员链访问服务失败"},

	// 其他: 2000+
	ErrCodeApplyChainAccessService: {"FailedOperation.ApplyChainAccessService", "apply chain access service failed", "获取链访问服务失败"},
	ErrCodeAccountNotExists:        {"FailedOperation.AccountNotExists", "account not exists", "该账号不存在"},
	ErrCodeAccountNicknameNotMatch: {"FailedOperation.AccountNicknameNotMatch", "account appId not match nickname", "账号和昵称不匹配"},
	ErrCodePanic:                   {"InternalError.ServicePanic", "service panic", "服务异常，请重试"},
}

// String - return err string
func (e ErrCode) String() string {
	if s, ok := ErrMessage[e]; ok {
		if len(s) == 3 {
			return s[ErrorTypeEN]
		}
	}

	return fmt.Sprintf("unknown error code %d", uint32(e))
}
