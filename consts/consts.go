//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package consts

import (
	"github.com/energye/golcl/energy/homedir"
	"github.com/energye/golcl/lcl/api"
	"os"
	"path/filepath"
	"unsafe"
)

var (
	ExePath       string                       //执行文件目录
	HomeDir, _    = homedir.Dir()              //系统用户目录
	Separator     = string(filepath.Separator) //
	SingleProcess = false                      //进程启动方式, true单进程 false多进程
)

const (
	Empty               = ""
	MAINARGS_NETIPCPORT = "net-ipc-port"
	download_dir        = "downloads"
	ipcAccIdx           = 8 //ipc args新的参数开始位置
	ENERGY_HOME_KEY     = "ENERGY_HOME"
	MemoryNetwork       = "unix"
)

func init() {
	ExePath, _ = os.Getwd()
}

type _int8 int8
type _int16 int16
type _int32 int32
type _int64 int64
type _int int
type _uint8 uint8
type _uint16 uint16
type _uint32 uint32
type _uint64 uint64
type _uint uint
type _uintptr uintptr
type _string string
type _boolean bool
type _float32 float32
type _float64 float64

func (m _int8) ToPtr() uintptr {
	return uintptr(m)
}

func (m _int16) ToPtr() uintptr {
	return uintptr(m)
}

func (m _int32) ToPtr() uintptr {
	return uintptr(m)
}

func (m _int64) ToPtr() uintptr {
	return uintptr(m)
}

func (m _int) ToPtr() uintptr {
	return uintptr(m)
}

func (m _uint8) ToPtr() uintptr {
	return uintptr(m)
}

func (m _uint16) ToPtr() uintptr {
	return uintptr(m)
}

func (m _uint32) ToPtr() uintptr {
	return uintptr(m)
}

func (m _uint64) ToPtr() uintptr {
	return uintptr(m)
}

func (m _uint) ToPtr() uintptr {
	return uintptr(m)
}

func (m _uintptr) ToPtr() uintptr {
	return uintptr(m)
}

func (m _string) ToPtr() uintptr {
	return api.PascalStr(string(m))
}

func (m _boolean) ToPtr() uintptr {
	return api.PascalBool(bool(m))
}

func (m _float32) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

func (m _float64) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

//0:net 1:unix
type IPC_TYPE = _int8

const (
	IPCT_NET = IPC_TYPE(iota)
	IPCT_UNIX
)

type ChannelType = _int8

const (
	Ct_Server = ChannelType(iota)
	Ct_Client
)

type TriggerMode = _int8

const (
	Tm_Async    = TriggerMode(iota) //异步
	Tm_Callback                     //异步，带回调函数返回结果
	Tm_Sync                         //同步，阻塞等待结果返回值
)

//功能和消息常量
const (
	WM_APP                   = 0x008000
	MINIBROWSER_SHOWDEVTOOLS = WM_APP + 0x00101 //开发工具展示
	MINIBROWSER_HIDEDEVTOOLS = WM_APP + 0x00102 //开发工具隐藏
)

//缩放、放大
type ZOOM = _int32

const (
	ZOOM_RESET = ZOOM(iota)
	ZOOM_INC
	ZOOM_DEC
)

//进程结束的状态
type TCefTerminationStatus = _int32

const (
	TS_ABNORMAL_TERMINATION = TCefTerminationStatus(iota)
	TS_PROCESS_WAS_KILLED
	TS_PROCESS_CRASHED
	TS_PROCESS_OOM
)

//前进 & 后退
type BF = _int32

const (
	BF_GOBACK = BF(iota)
	BF_GOFORWARD
)

//日志等级
type LOG = _uint32

const (
	LOGSEVERITY_DEFAULT LOG = 0
	LOGSEVERITY_VERBOSE LOG = 1
	LOGSEVERITY_DEBUG   LOG = LOGSEVERITY_VERBOSE
	LOGSEVERITY_INFO    LOG = 2
	LOGSEVERITY_WARNING LOG = 3
	LOGSEVERITY_ERROR   LOG = 4
	LOGSEVERITY_FATAL   LOG = 5
	LOGSEVERITY_DISABLE LOG = 99
)

type LANGUAGE = _string

const (
	LANGUAGE_zh_CN  LANGUAGE = "zh-CN"
	LANGUAGE_zh_TW  LANGUAGE = "zh-TW"
	LANGUAGE_am     LANGUAGE = "am"
	LANGUAGE_ar     LANGUAGE = "ar"
	LANGUAGE_bg     LANGUAGE = "bg"
	LANGUAGE_bn     LANGUAGE = "bn"
	LANGUAGE_ca     LANGUAGE = "ca"
	LANGUAGE_cs     LANGUAGE = "cs"
	LANGUAGE_da     LANGUAGE = "da"
	LANGUAGE_de     LANGUAGE = "de"
	LANGUAGE_el     LANGUAGE = "el"
	LANGUAGE_en_GB  LANGUAGE = "en-GB"
	LANGUAGE_en_US  LANGUAGE = "en-US"
	LANGUAGE_es     LANGUAGE = "es"
	LANGUAGE_es_419 LANGUAGE = "es-419"
	LANGUAGE_et     LANGUAGE = "et"
	LANGUAGE_fa     LANGUAGE = "fa"
	LANGUAGE_fi     LANGUAGE = "fi"
	LANGUAGE_fil    LANGUAGE = "fil"
	LANGUAGE_fr     LANGUAGE = "fr"
	LANGUAGE_gu     LANGUAGE = "gu"
	LANGUAGE_he     LANGUAGE = "he"
	LANGUAGE_hi     LANGUAGE = "hi"
	LANGUAGE_hr     LANGUAGE = "hr"
	LANGUAGE_hu     LANGUAGE = "hu"
	LANGUAGE_id     LANGUAGE = "channelId"
	LANGUAGE_it     LANGUAGE = "it"
	LANGUAGE_ja     LANGUAGE = "ja"
	LANGUAGE_kn     LANGUAGE = "kn"
	LANGUAGE_ko     LANGUAGE = "ko"
	LANGUAGE_lt     LANGUAGE = "lt"
	LANGUAGE_lv     LANGUAGE = "lv"
	LANGUAGE_ml     LANGUAGE = "ml"
	LANGUAGE_mr     LANGUAGE = "mr"
	LANGUAGE_ms     LANGUAGE = "ms"
	LANGUAGE_nb     LANGUAGE = "nb"
	LANGUAGE_nl     LANGUAGE = "nl"
	LANGUAGE_pl     LANGUAGE = "pl"
	LANGUAGE_pt_BR  LANGUAGE = "pt-BR"
	LANGUAGE_pt_PT  LANGUAGE = "pt-PT"
	LANGUAGE_ro     LANGUAGE = "ro"
	LANGUAGE_ru     LANGUAGE = "ru"
	LANGUAGE_sk     LANGUAGE = "sk"
	LANGUAGE_sl     LANGUAGE = "sl"
	LANGUAGE_sr     LANGUAGE = "sr"
	LANGUAGE_sv     LANGUAGE = "sv"
	LANGUAGE_sw     LANGUAGE = "sw"
	LANGUAGE_ta     LANGUAGE = "ta"
	LANGUAGE_te     LANGUAGE = "te"
	LANGUAGE_th     LANGUAGE = "th"
	LANGUAGE_tr     LANGUAGE = "tr"
	LANGUAGE_uk     LANGUAGE = "uk"
	LANGUAGE_vi     LANGUAGE = "vi"
)

// Chromium关闭的操作类型
// 在 TChromium.Onclose 使用
// -------------------------
// cbaCancel : 停止关闭浏览器
// cbaClose  : 继续关闭浏览器
// cbaDelay  : 暂时停止关闭浏览器
//			 : 当应用程序需要在关闭浏览器之前执行一些自定义进程时使用。在关闭浏览器之前，通常需要在主线程中销毁TCEFWindowParent。
type CBS = _int32

const (
	CbaClose = CBS(iota)
	CbaDelay
	CbaCancel
)

//CEF 进程 ChannelId
type CefProcessId = _int32

const (
	PID_BROWSER = CefProcessId(iota)
	PID_RENDER
)

//支持的JS类型
type V8_JS_VALUE_TYPE = _int32

const (
	V8_VALUE_STRING = V8_JS_VALUE_TYPE(iota)
	V8_VALUE_INT
	V8_VALUE_DOUBLE
	V8_VALUE_BOOLEAN
	V8_VALUE_NULL
	V8_VALUE_UNDEFINED
	V8_VALUE_OBJECT
	V8_VALUE_ARRAY
	V8_VALUE_FUNCTION
	V8_VALUE_EXCEPTION
	V8_VALUE_ROOT_OBJECT
	V8_NO_OUT_VALUE
)

//支持的GO类型
type GO_VALUE_TYPE = _int32

const (
	GO_VALUE_STRING = GO_VALUE_TYPE(iota)
	GO_VALUE_INT
	GO_VALUE_INT8
	GO_VALUE_INT16
	GO_VALUE_INT32
	GO_VALUE_INT64
	GO_VALUE_UINT
	GO_VALUE_UINT8
	GO_VALUE_UINT16
	GO_VALUE_UINT32
	GO_VALUE_UINT64
	GO_VALUE_UINTPTR
	GO_VALUE_FLOAT32
	GO_VALUE_FLOAT64
	GO_VALUE_BOOL
	GO_VALUE_NIL
	GO_VALUE_STRUCT
	GO_VALUE_SLICE
	GO_VALUE_FUNC
	GO_VALUE_PTR
	GO_VALUE_EXCEPTION
	GO_VALUE_INVALID_TYPE //无效类型
	GO_VALUE_ARGUMENT     //argument
	GO_VALUE_DICTVALUE    //dictValue
)

//JS属性
type V8_PROPERTY_ATTRIBUTE = _int32

const (
	V8_PROPERTY_ATTRIBUTE_NONE       V8_PROPERTY_ATTRIBUTE = 0
	V8_PROPERTY_ATTRIBUTE_READONLY   V8_PROPERTY_ATTRIBUTE = 1 << 0
	V8_PROPERTY_ATTRIBUTE_DONTENUM   V8_PROPERTY_ATTRIBUTE = 1 << 1
	V8_PROPERTY_ATTRIBUTE_DONTDELETE V8_PROPERTY_ATTRIBUTE = 1 << 2
)

//JS交互绑定的事件类型
type BIND_EVENT = _int32

const (
	BE_SET = BIND_EVENT(iota)
	BE_GET
	BE_FUNC
)

//异常信息
type CEF_V8_EXCEPTION = _int32

const (
	CVE_ERROR_OK                             = CEF_V8_EXCEPTION(iota) //操作成功
	CVE_ERROR_NOT_FOUND_FIELD                                         //未找到字段 或字段未定义
	CVE_ERROR_NOT_FOUND_FUNC                                          //未找到函数 或函数未定义
	CVE_ERROR_TYPE_NOT_SUPPORTED                                      //不支持的变量类型 变量类型只支持[string int double bool null undefined]
	CVE_ERROR_TYPE_CANNOT_CHANGE                                      //字段为普通类型不能变更为 array、object、function
	CVE_ERROR_TYPE_INVALID                                            //类型无效
	CVE_ERROR_GET_STRING_FAIL                                         //获取string类型失败
	CVE_ERROR_GET_INT_FAIL                                            //获取int类型失败
	CVE_ERROR_GET_DOUBLE_FAIL                                         //获取double类型失败
	CVE_ERROR_GET_BOOL_FAIL                                           //获取bool类型失败
	CVE_ERROR_GET_NULL_FAIL                                           //获取null类型失败
	CVE_ERROR_GET_UNDEFINED_FAIL                                      //获取undefined类型失败
	CVE_ERROR_FUNC_INVALID_P_L_9                                      //该函数非法 类型不正确 或参数个数大于9个
	CVE_ERROR_FUNC_IN_PAM                                             //入参类型不正确 只能为string int double boolean
	CVE_ERROR_FUNC_OUT_PAM                                            //出参类型不正确 只能为EefError 或 可选的[string int double boolean]
	CVE_ERROR_FUNC_GET_IN_PAM_STRING_FAIL                             //入参获取string类型值失败
	CVE_ERROR_FUNC_GET_IN_PAM_INT_FAIL                                //入参获取int类型值失败
	CVE_ERROR_FUNC_GET_IN_PAM_DOUBLE_FAIL                             //入参获取double类型值失败
	CVE_ERROR_FUNC_GET_IN_PAM_BOOLEAN_FAIL                            //入参获取boolean类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_STRING_FAIL                            //出参获取string类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_INT_FAIL                               //出参获取int类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_DOUBLE_FAIL                            //出参获取double类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_BOOLEAN_FAIL                           //出参获取boolean类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_CEFERROR_FAIL                          //出参获取CefError值失败
	CVE_ERROR_IPC_GET_BIND_FIELD_VALUE_FAIL                           //IPC获取绑定值失败
	CVE_ERROR_UNKNOWN_ERROR                                           //未知错误
)

const (
	BIND_FUNC_IN_MAX_SUM  = 9 //函数最大入参数
	BIND_FUNC_OUT_MAX_SUM = 1 //函数最大出参数
)

//函数类型
type FN_TYPE = _int8

const (
	FN_TYPE_COMMON = FN_TYPE(iota) //普通函数，直接定义的
	FN_TYPE_OBJECT                 //对象函数，所属对象
)

//通用类型或对象类型
type IS_CO = _int8

const (
	IS_COMMON = IS_CO(iota)
	IS_OBJECT
)

//进程消息类型
type PROCESS_MESSAGE_TYPE = _int8

const (
	PMT_JS_CODE = PROCESS_MESSAGE_TYPE(iota) //执行JS代码消息
	PMT_TEXT                                 //文本传递消息
	PMT_BINARY                               //二进制消息
)

type TCefProcessType = _int8

const (
	PtBrowser = TCefProcessType(iota)
	PtRender
	PtZygote
	PtGPU
	PtUtility
	PtBroker
	PtCrashpad
	PtOther
)

type TDateTime = _float64

type TCefCookieSameSite = _int32

const (
	Ccss_CEF_COOKIE_SAME_SITE_UNSPECIFIED = TCefCookieSameSite(iota)
	Ccss_CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	Ccss_CEF_COOKIE_SAME_SITE_LAX_MODE
	Ccss_CEF_COOKIE_SAME_SITE_STRICT_MODE
)

type TCefCookiePriority = _int32

const (
	CEF_COOKIE_PRIORITY_LOW    TCefCookiePriority = -1
	CEF_COOKIE_PRIORITY_MEDIUM TCefCookiePriority = 0
	CEF_COOKIE_PRIORITY_HIGH   TCefCookiePriority = 1
)

type TCefProxyType = _int32

const (
	PtDirect = TCefProxyType(iota)
	PtAutodetect
	PtSystem
	PtFixedServers
	PtPACScript
)

type TCefProxyScheme = _int32

const (
	PsHTTP = TCefProxyScheme(iota)
	PsSOCKS4
	PsSOCKS5
)

type TCefContextMenuType = _int32

const (
	CMT_NONE = TCefContextMenuType(iota)
	CMT_CHECK
	CMT_RADIO
)

type TCefContextMenuMediaType = _int32

const (
	CM_MEDIATYPE_NONE = TCefContextMenuMediaType(iota)
	CM_MEDIATYPE_IMAGE
	CM_MEDIATYPE_VIDEO
	CM_MEDIATYPE_AUDIO
	CM_MEDIATYPE_CANVAS
	CM_MEDIATYPE_FILE
	CM_MEDIATYPE_PLUGIN
)

type MenuId = _int32

const (
	MENU_ID_BACK                       MenuId = 100
	MENU_ID_FORWARD                    MenuId = 101
	MENU_ID_RELOAD                     MenuId = 102
	MENU_ID_RELOAD_NOCACHE             MenuId = 103
	MENU_ID_STOPLOAD                   MenuId = 104
	MENU_ID_UNDO                       MenuId = 110
	MENU_ID_REDO                       MenuId = 111
	MENU_ID_CUT                        MenuId = 112
	MENU_ID_COPY                       MenuId = 113
	MENU_ID_PASTE                      MenuId = 114
	MENU_ID_DELETE                     MenuId = 115
	MENU_ID_SELECT_ALL                 MenuId = 116
	MENU_ID_FIND                       MenuId = 130
	MENU_ID_PRINT                      MenuId = 131
	MENU_ID_VIEW_SOURCE                MenuId = 132
	MENU_ID_SPELLCHECK_SUGGESTION_0    MenuId = 200
	MENU_ID_SPELLCHECK_SUGGESTION_1    MenuId = 201
	MENU_ID_SPELLCHECK_SUGGESTION_2    MenuId = 202
	MENU_ID_SPELLCHECK_SUGGESTION_3    MenuId = 203
	MENU_ID_SPELLCHECK_SUGGESTION_4    MenuId = 204
	MENU_ID_SPELLCHECK_SUGGESTION_LAST MenuId = 204
	MENU_ID_NO_SPELLING_SUGGESTIONS    MenuId = 205
	MENU_ID_ADD_TO_DICTIONARY          MenuId = 206
	MENU_ID_CUSTOM_FIRST               MenuId = 220
	MENU_ID_CUSTOM_LAST                MenuId = 250
	MENU_ID_USER_FIRST                 MenuId = 26500
	MENU_ID_USER_LAST                  MenuId = 28500
)

type TCefMenuColorType = _int32

const (
	CEF_MENU_COLOR_TEXT = TCefMenuColorType(iota)
	CEF_MENU_COLOR_TEXT_HOVERED
	CEF_MENU_COLOR_TEXT_ACCELERATOR
	CEF_MENU_COLOR_TEXT_ACCELERATOR_HOVERED
	CEF_MENU_COLOR_BACKGROUND
	CEF_MENU_COLOR_BACKGROUND_HOVERED
	CEF_MENU_COLOR_COUNT
)

type ARGB = _uint32

type TCefKeyEventType = _int32

const (
	KEYEVENT_RAW_KEYDOWN = TCefKeyEventType(iota)
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

type TCefEventFlags = _uint32

type TCefWindowHandleType = _int8

type TCefWindowHandle = _uintptr

const (
	Wht_WindowParent = TCefWindowHandleType(iota)
	Wht_LinkedWindowParent
)

type TCefReturnValue = _int32

const (
	RV_CANCEL = TCefReturnValue(iota)
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

type TCefReferrerPolicy = _int32

const (
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_FROM_SECURE_TO_INSECURE = TCefReferrerPolicy(iota) // same value as REFERRER_POLICY_DEFAULT
	REFERRER_POLICY_REDUCE_REFERRER_GRANULARITY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_ONLY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_NEVER_CLEAR_REFERRER
	REFERRER_POLICY_ORIGIN
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_CLEAR_ON_TRANSITION_FROM_SECURE_TO_INSECURE
	REFERRER_POLICY_NO_REFERRER // REFERRER_POLICY_LAST_VALUE = REFERRER_POLICY_NO_REFERRER
)

type TCefUrlRequestFlags = _int

type TCefErrorCode = _int32

type TCefResourceType = _int32

const (
	RT_MAIN_FRAME = TCefResourceType(iota)
	RT_SUB_FRAME
	RT_STYLESHEET
	RT_SCRIPT
	RT_IMAGE
	RT_FONT_RESOURCE
	RT_SUB_RESOURCE
	RT_OBJECT
	RT_MEDIA
	RT_WORKER
	RT_SHARED_WORKER
	RT_PREFETCH
	RT_FAVICON
	RT_XHR
	RT_PING
	RT_SERVICE_WORKER
	RT_CSP_REPORT
	RT_PLUGIN_RESOURCE
	RT_EMPTY_FILLER_TYPE_DO_NOT_USE  // This type doesn't exist in CEF and it's here just to fill this position.
	RT_NAVIGATION_PRELOAD_MAIN_FRAME // This type must have a value of 19
	RT_NAVIGATION_PRELOAD_SUB_FRAME
)

type TCefTransitionType = _int

type TCefUrlRequestStatus = _int32

const (
	UR_UNKNOWN = TCefUrlRequestStatus(iota)
	UR_SUCCESS
	UR_IO_PENDING
	UR_CANCELED
	UR_FAILED
)

type TCefState = _int32

const (
	STATE_DEFAULT = TCefState(iota)
	STATE_ENABLED
	STATE_DISABLE
)

type TCefTouchEeventType = _int32

const (
	CEF_TET_RELEASED = TCefTouchEeventType(iota)
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

type TCefPointerType = _int32

const (
	CEF_POINTER_TYPE_TOUCH = TCefPointerType(iota)
	CEF_POINTER_TYPE_MOUSE
	CEF_POINTER_TYPE_PEN
	CEF_POINTER_TYPE_ERASER
	CEF_POINTER_TYPE_UNKNOWN
)

type TCefMouseButtonType = _int32

const (
	MBT_LEFT = TCefMouseButtonType(iota)
	MBT_MIDDLE
	MBT_RIGHT
)

//进程消息错误码
type ProcessMessageError = _int32

const (
	PME_OK                        ProcessMessageError = iota + 1 //发送成功
	PMErr_NOT_FOUND_FRAME         ProcessMessageError = -1       //没找到Frame
	PMErr_TARGET_PROCESS          ProcessMessageError = -2       //目标进程标识错误
	PMErr_NAME_IS_NULL            ProcessMessageError = -3       //消息名称为空
	PMErr_NO_INVALID_FRAME        ProcessMessageError = -4       //无效的Frame
	PMErr_REQUIRED_PARAMS_IS_NULL ProcessMessageError = -5       //必要参数为空
	PMErr_NAME_CANNOT_USED        ProcessMessageError = -6       //不能使用的消息名称
)

type TCefWindowOpenDisposition = _int32

const (
	WOD_UNKNOWN = TCefWindowOpenDisposition(iota)
	WOD_CURRENT_TAB
	WOD_SINGLETON_TAB
	WOD_NEW_FOREGROUND_TAB
	WOD_NEW_BACKGROUND_TAB
	WOD_NEW_POPUP
	WOD_NEW_WINDOW
	WOD_SAVE_TO_DISK
	WOD_OFF_THE_RECORD
	WOD_IGNORE_ACTION
	WOD_SWITCH_TO_TAB
	WOD_NEW_PICTURE_IN_PICTURE
)

// Browser Window Type
type WINDOW_TYPE = _int8

const (
	WT_MAIN_BROWSER = WINDOW_TYPE(iota)
	WT_POPUP_SUB_BROWSER
	WT_DEV_TOOLS
	WT_VIEW_SOURCE
)

type TCefContextMenuTypeFlags = _uint32
type TCefContextMenuMediaStateFlags = _uint32
type TCefContextMenuEditStateFlags = _uint32

type TCefMenuAnchorPosition = _int32

const (
	CEF_MENU_ANCHOR_TOPLEFT = TCefMenuAnchorPosition(iota)
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

type TCefShowState = _int32
