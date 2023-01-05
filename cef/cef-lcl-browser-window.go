//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	. "github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
	"time"
)

type IBaseWindow interface {
	lcl.IWinControl
	FormCreate()
	WindowParent() ITCefWindowParent
	Chromium() IChromium
	ChromiumCreate(config *tCefChromiumConfig, defaultUrl string)
	registerEvent()
	registerDefaultEvent()
}

//LCLBrowserWindow 基于chromium 和 lcl 的窗口组件
type LCLBrowserWindow struct {
	*lcl.TForm                          //
	chromium         IChromium          //
	browser          *ICefBrowser       //
	windowParent     ITCefWindowParent  //
	windowProperty   *WindowProperty    //
	windowId         int32              //
	windowType       consts.WINDOW_TYPE //0:browser 1:devTools 2:viewSource 默认:0
	isClosing        bool               //
	canClose         bool               //
	onResize         []TNotifyEvent     //
	onActivate       []TNotifyEvent     //
	onShow           []TNotifyEvent     //
	onClose          []TCloseEvent      //
	onCloseQuery     []TCloseQueryEvent //
	onActivateAfter  lcl.TNotifyEvent   //
	isFormCreate     bool               //是否创建完成 WindowForm
	isChromiumCreate bool               //是否创建完成 Chromium
	frames           TCEFFrame          //当前浏览器下的所有frame
	auxTools         *auxTools          //辅助工具
}

//创建一个带有 chromium 窗口
//
//该窗口默认不具备默认事件处理能力, 通过 EnableDefaultEvent 函数注册事件处理
func NewBrowserWindow(config *tCefChromiumConfig, windowProperty *WindowProperty) *LCLBrowserWindow {
	if windowProperty == nil {
		windowProperty = NewWindowProperty()
	}
	var window = NewWindow(windowProperty)
	window.ChromiumCreate(config, windowProperty.Url)
	window.putChromiumWindowInfo()
	//OnBeforeBrowser 是一个必须的默认事件，在浏览器创建时窗口序号会根据browserId生成
	window.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool { return false })
	return window
}

func (m *LCLBrowserWindow) Browser() *ICefBrowser {
	return m.browser
}

func (m *LCLBrowserWindow) Chromium() IChromium {
	return m.chromium
}

func (m *LCLBrowserWindow) Id() int32 {
	return m.windowId
}

func (m *LCLBrowserWindow) Show() {
	if m.TForm == nil {
		return
	}
	m.TForm.Show()
}

func (m *LCLBrowserWindow) Hide() {
	if m.TForm == nil {
		return
	}
	m.TForm.Hide()
}

func (m *LCLBrowserWindow) Visible() bool {
	if m.TForm == nil {
		return false
	}
	return m.TForm.Visible()
}

func (m *LCLBrowserWindow) SetVisible(value bool) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetVisible(value)
}

//以默认的方式展示在任务栏上
func (m *LCLBrowserWindow) SetDefaultInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StDefault)
}

//展示在任务栏上
func (m *LCLBrowserWindow) SetShowInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StAlways)
}

//不会展示在任务栏上
func (m *LCLBrowserWindow) SetNotInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StNever)
}

//返回chromium的父组件对象，该对象不是window组件对象,属于window的一个子组件
//
//在windows下它是 TCEFWindowParent, linux或macOSx下它是 TCEFLinkedWindowParent
//
//通过函数可调整该组件的属性
func (m *LCLBrowserWindow) WindowParent() ITCefWindowParent {
	return m.windowParent
}

//返回窗口关闭状态
func (m *LCLBrowserWindow) IsClosing() bool {
	return m.isClosing
}

// 设置窗口类型
func (m *LCLBrowserWindow) SetWindowType(windowType consts.WINDOW_TYPE) {
	m.windowType = windowType
}

// 返回窗口类型
func (m *LCLBrowserWindow) WindowType() consts.WINDOW_TYPE {
	return m.windowType
}

// 创建window浏览器组件
//
// 不带有默认事件的chromium
func (m *LCLBrowserWindow) ChromiumCreate(config *tCefChromiumConfig, defaultUrl string) {
	if m.isChromiumCreate {
		return
	}
	m.isChromiumCreate = true
	m.windowId = BrowserWindow.GetNextWindowNum()
	if config == nil {
		config = NewChromiumConfig()
	}
	m.chromium = NewChromium(m, config)
	m.chromium.SetEnableMultiBrowserMode(true)
	if defaultUrl != "" {
		m.chromium.SetDefaultURL(defaultUrl)
	}
	//windowParent
	m.windowParent = NewCEFWindow(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	m.windowParent.SetChromium(m.chromium, 0)
	m.windowParent.SetOnEnter(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.chromium.Initialized()
		m.chromium.FrameIsFocused()
		m.chromium.SetFocus(true)
	})
	m.windowParent.SetOnExit(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.chromium.SendCaptureLostEvent()
	})
}

func (m *LCLBrowserWindow) putChromiumWindowInfo() {
	BrowserWindow.putWindowInfo(m.windowId, m)
}

//默认的chromium事件
func (m *LCLBrowserWindow) defaultChromiumEvent() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		AddGoForm(m.windowId, m.Instance())
		m.registerPopupEvent()
		m.registerDefaultEvent()
		m.registerDefaultChromiumCloseEvent()
	}
}

// 创建窗口
//
// 不带有默认事件的窗口
func (m *LCLBrowserWindow) FormCreate() {
	if m.isFormCreate {
		return
	}
	m.isFormCreate = true
	m.SetPosition(types.PoDesktopCenter)
	m.SetName(fmt.Sprintf("energy_window_name_%d", time.Now().UnixNano()/1e6))
}

//默认窗口活动/关闭处理事件
func (m *LCLBrowserWindow) defaultWindowEvent() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		m.SetOnResize(m.resize)
		m.SetOnActivate(m.activate)
	}
	m.SetOnShow(m.show)
}

//默认的窗口关闭事件
func (m *LCLBrowserWindow) defaultWindowCloseEvent() {
	m.SetOnClose(m.close)
	m.SetOnCloseQuery(m.closeQuery)
}

//启用默认关闭事件行为-该窗口将被关闭
func (m *LCLBrowserWindow) EnableDefaultClose() {
	m.defaultWindowCloseEvent()
	m.registerDefaultChromiumCloseEvent()
}

//启用所有默认事件行为
func (m *LCLBrowserWindow) EnableAllDefaultEvent() {
	m.defaultWindowCloseEvent()
	m.defaultChromiumEvent()
}

// 添加OnResize事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) AddOnResize(fn TNotifyEvent) {
	m.onResize = append(m.onResize, fn)
}

// 添加OnActivate事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) AddOnActivate(fn TNotifyEvent) {
	m.onActivate = append(m.onActivate, fn)
}

// 添加OnShow事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) AddOnShow(fn TNotifyEvent) {
	m.onShow = append(m.onShow, fn)
}

// 添加OnClose事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) AddOnClose(fn TCloseEvent) {
	m.onClose = append(m.onClose, fn)
}

// 添加OnCloseQuery事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) AddOnCloseQuery(fn TCloseQueryEvent) {
	m.onCloseQuery = append(m.onCloseQuery, fn)
}

//每次激活窗口之后执行一次
func (m *LCLBrowserWindow) SetOnActivateAfter(fn lcl.TNotifyEvent) {
	m.onActivateAfter = fn
}

func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		m.SetWindowState(types.WsMinimized)
	})
}

func (m *LCLBrowserWindow) Maximize() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		var bs = m.BorderStyle()
		var monitor = m.Monitor().WorkareaRect()
		if m.windowProperty == nil {
			m.windowProperty = &WindowProperty{}
		}
		if bs == types.BsNone {
			var ws = m.WindowState()
			var redWindowState types.TWindowState
			//默认状态0
			if m.windowProperty.WindowState == types.WsNormal && m.windowProperty.WindowState == ws {
				redWindowState = types.WsMaximized
			} else {
				if m.windowProperty.WindowState == types.WsNormal {
					redWindowState = types.WsMaximized
				} else if m.windowProperty.WindowState == types.WsMaximized {
					redWindowState = types.WsNormal
				}
			}
			m.windowProperty.WindowState = redWindowState
			if redWindowState == types.WsMaximized {
				m.windowProperty.X = m.Left()
				m.windowProperty.Y = m.Top()
				m.windowProperty.Width = m.Width()
				m.windowProperty.Height = m.Height()
				m.SetLeft(monitor.Left)
				m.SetTop(monitor.Top)
				m.SetWidth(monitor.Right - monitor.Left - 1)
				m.SetHeight(monitor.Bottom - monitor.Top - 1)
			} else if redWindowState == types.WsNormal {
				m.SetLeft(m.windowProperty.X)
				m.SetTop(m.windowProperty.Y)
				m.SetWidth(m.windowProperty.Width)
				m.SetHeight(m.windowProperty.Height)
			}
		} else {
			if m.WindowState() == types.WsMaximized {
				m.SetWindowState(types.WsNormal)
				if IsDarwin() {
					m.SetWindowState(types.WsMaximized)
					m.SetWindowState(types.WsNormal)
				}
			} else if m.WindowState() == types.WsNormal {
				m.SetWindowState(types.WsMaximized)
			}
			m.windowProperty.WindowState = m.WindowState()
		}
	})
}

// 关闭带有浏览器的窗口
func (m *LCLBrowserWindow) CloseBrowserWindow() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		if m == nil {
			logger.Error("关闭浏览器 WindowInfo 为空")
			return
		}
		if IsDarwin() {
			//main window close
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				m.Close()
			} else {
				//sub window close
				m.isClosing = true
				m.Hide()
				m.chromium.CloseBrowser(true)
			}
		} else {
			m.isClosing = true
			m.Hide()
			m.chromium.CloseBrowser(true)
		}
	})
}

//禁用口透明
func (m *LCLBrowserWindow) DisableTransparent() {
	if m.TForm == nil {
		return
	}
	m.SetAllowDropFiles(false)
	m.SetAlphaBlend(false)
	m.SetAlphaBlendValue(255)
}

//使窗口透明 value 0 ~ 255
func (m *LCLBrowserWindow) EnableTransparent(value uint8) {
	if m.TForm == nil {
		return
	}
	m.SetAllowDropFiles(true)
	m.SetAlphaBlend(true)
	m.SetAlphaBlendValue(value)
}

//禁用最小化按钮
func (m *LCLBrowserWindow) DisableMinimize() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Exclude(types.BiMinimize))
}

//禁用最大化按钮
func (m *LCLBrowserWindow) DisableMaximize() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Exclude(types.BiMaximize))
}

//禁用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *LCLBrowserWindow) DisableSystemMenu() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Exclude(types.BiSystemMenu))
}

//禁用帮助菜单
func (m *LCLBrowserWindow) DisableHelp() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Exclude(types.BiHelp))
}

//启用最小化按钮
func (m *LCLBrowserWindow) EnableMinimize() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Include(types.BiMinimize))
}

//启用最大化按钮
func (m *LCLBrowserWindow) EnableMaximize() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Include(types.BiMaximize))
}

//启用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *LCLBrowserWindow) EnableSystemMenu() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Include(types.BiSystemMenu))
}

//启用帮助菜单
func (m *LCLBrowserWindow) EnableHelp() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Include(types.BiHelp))
}

func (m *LCLBrowserWindow) show(sender lcl.IObject) {
	var ret bool
	if m.onShow != nil {
		for _, fn := range m.onShow {
			if fn(sender) {
				ret = true
			}
		}
	}
	if !ret {
		if m.windowParent != nil {
			QueueAsyncCall(func(id int) {
				m.windowParent.UpdateSize()
			})
		}
	}
}

func (m *LCLBrowserWindow) resize(sender lcl.IObject) {
	var ret bool
	if m.onResize != nil {
		for _, fn := range m.onResize {
			if fn(sender) {
				ret = true
			}
		}
	}
	if !ret {
		if m.isClosing {
			return
		}
		if m.chromium != nil {
			m.chromium.NotifyMoveOrResizeStarted()
		}
		if m.windowParent != nil {
			m.windowParent.UpdateSize()
		}
	}
}

func (m *LCLBrowserWindow) activate(sender lcl.IObject) {
	var ret bool
	if m.onActivate != nil {
		for _, fn := range m.onActivate {
			if fn(sender) {
				ret = true
			}
		}
	}
	if !ret {
		if m.isClosing {
			return
		}
		if m.chromium != nil {
			if !m.chromium.Initialized() {
				m.chromium.CreateBrowser(m.windowParent)
			}
		}
	}
	if m.onActivateAfter != nil {
		m.onActivateAfter(sender)
	}
}
func (m *LCLBrowserWindow) registerPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !api.GoBool(BrowserWindow.Config.chromiumConfig.enableWindowPopup) {
			return true
		}
		BrowserWindow.popupWindow.SetWindowType(consts.WT_POPUP_SUB_BROWSER)
		BrowserWindow.popupWindow.ChromiumCreate(BrowserWindow.Config.chromiumConfig, beforePopupInfo.TargetUrl)
		BrowserWindow.popupWindow.chromium.EnableIndependentEvent()
		BrowserWindow.popupWindow.putChromiumWindowInfo()
		BrowserWindow.popupWindow.defaultChromiumEvent()
		var result = false
		defer func() {
			if result {
				QueueAsyncCall(func(id int) {
					winProperty := BrowserWindow.popupWindow.windowProperty
					if winProperty != nil {
						if winProperty.IsShowModel {
							BrowserWindow.popupWindow.ShowModal()
							return
						}
					}
					BrowserWindow.popupWindow.Show()
				})
			}
		}()
		if bwEvent.onBeforePopup != nil {
			result = !bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, BrowserWindow.popupWindow, noJavascriptAccess)
		}
		return result
	})
}

// 默认事件注册 部分事件允许被覆盖
func (m *LCLBrowserWindow) registerDefaultEvent() {
	var bwEvent = BrowserWindow.browserEvent
	//默认自定义快捷键
	defaultAcceleratorCustom()
	m.chromium.SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ipc.ICefProcessMessage) bool {
		if bwEvent.onProcessMessageReceived != nil {
			return bwEvent.onProcessMessageReceived(sender, browser, frame, sourceProcess, message)
		}
		return false
	})
	m.chromium.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
		}
		if bwEvent.onBeforeResourceLoad != nil {
			bwEvent.onBeforeResourceLoad(sender, browser, frame, request, callback, result)
		}
	})
	//事件可以被覆盖
	m.chromium.SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *DownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback)
		} else {
			callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
		}
	})
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		chromiumOnBeforeContextMenu(sender, browser, frame, params, model)
		if bwEvent.onBeforeContextMenu != nil {
			bwEvent.onBeforeContextMenu(sender, browser, frame, params, model)
		}
	})
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
		chromiumOnContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		if bwEvent.onContextMenuCommand != nil {
			bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		}
	})
	m.chromium.SetOnLoadingStateChange(func(sender lcl.IObject, browser *ICefBrowser, isLoading, canGoBack, canGoForward bool) {
		if bwEvent.onLoadingStateChange != nil {
			bwEvent.onLoadingStateChange(sender, browser, isLoading, canGoBack, canGoForward)
		}
	})
	m.chromium.SetOnFrameCreated(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) {
		QueueAsyncCall(func(id int) {
			BrowserWindow.putBrowserFrame(browser, frame)
		})
		if bwEvent.onFrameCreated != nil {
			bwEvent.onFrameCreated(sender, browser, frame)
		}
	})
	m.chromium.SetOnFrameDetached(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) {
		chromiumOnFrameDetached(browser, frame)
		if bwEvent.onFrameDetached != nil {
			bwEvent.onFrameDetached(sender, browser, frame)
		}
	})

	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
		if chromiumOnAfterCreate(browser) {
			return
		}
		if bwEvent.onAfterCreated != nil {
			bwEvent.onAfterCreated(sender, browser)
		}
	})
	//事件可以被覆盖
	m.chromium.SetOnKeyEvent(func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, result *bool) {
		if api.GoBool(BrowserWindow.Config.chromiumConfig.enableDevTools) {
			if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
				if winInfo.WindowType() == consts.WT_DEV_TOOLS || winInfo.WindowType() == consts.WT_VIEW_SOURCE {
					return
				}
			}
			if event.WindowsKeyCode == VkF12 && event.Kind == consts.KEYEVENT_RAW_KEYDOWN {
				browser.ShowDevTools()
				*result = true
			} else if event.WindowsKeyCode == VkF12 && event.Kind == consts.KEYEVENT_KEYUP {
				*result = true
			}
		}
		if KeyAccelerator.accelerator(browser, event, result) {
			return
		}
		if bwEvent.onKeyEvent != nil {
			bwEvent.onKeyEvent(sender, browser, event, result)
		}
	})
	m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool {
		chromiumOnBeforeBrowser(browser, frame)
		if bwEvent.onBeforeBrowser != nil {
			return bwEvent.onBeforeBrowser(sender, browser, frame)
		}
		return false
	})
	m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *ICefBrowser, title string) {
		updateBrowserDevTools(browser, title)
		updateBrowserViewSource(browser, title)
		if bwEvent.onTitleChange != nil {
			bwEvent.onTitleChange(sender, browser, title)
		}
	})
}

//控制LCL创建的窗口事件
func (m *LCLBrowserWindow) registerControlLCLWindowEvent() {

}

func (m *LCLBrowserWindow) close(sender lcl.IObject, action *types.TCloseAction) {
	var ret bool
	if m.onClose != nil {
		for _, fn := range m.onClose {
			if fn(sender, action) {
				ret = true
			}
		}
	}
	if !ret {
		logger.Debug("window.onClose")
		*action = types.CaFree
	}
}

func (m *LCLBrowserWindow) closeQuery(sender lcl.IObject, close *bool) {
	var ret bool
	if m.onCloseQuery != nil {
		for _, fn := range m.onCloseQuery {
			if fn(sender, close) {
				ret = true
			}
		}
	}
	if !ret {
		logger.Debug("window.onCloseQuery windowType:", m.WindowType())
		if IsDarwin() {
			//main window close
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				*close = true
				desChildWind := m.windowParent.DestroyChildWindow()
				logger.Debug("window.onCloseQuery => windowParent.DestroyChildWindow:", desChildWind)
			} else {
				//sub window close
				*close = m.canClose
			}
			if !m.isClosing {
				m.isClosing = true
				m.Hide()
				m.chromium.CloseBrowser(true)
			}
		} else {
			*close = m.canClose
			if !m.isClosing {
				m.isClosing = true
				m.Hide()
				m.chromium.CloseBrowser(true)
			}
		}
	}
}

//默认的chromium关闭事件
func (m *LCLBrowserWindow) registerDefaultChromiumCloseEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowsesAction) {
		logger.Debug("chromium.onClose")
		if IsDarwin() { //MacOSX
			desChildWind := m.windowParent.DestroyChildWindow()
			logger.Debug("chromium.onClose => windowParent.DestroyChildWindow:", desChildWind)
			*aAction = consts.CbaClose
		} else if IsLinux() {
			*aAction = consts.CbaClose
		} else if IsWindows() {
			*aAction = consts.CbaDelay
		}
		if !IsDarwin() {
			QueueAsyncCall(func(id int) { //main thread run
				m.windowParent.Free()
				logger.Debug("chromium.onClose => windowParent.Free")
			})
		}
		if bwEvent.onClose != nil {
			bwEvent.onClose(sender, browser, aAction)
		}
	})
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("chromium.onBeforeClose")
		chromiumOnBeforeClose(browser)
		m.canClose = true
		var closeWindow = func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("chromium.OnBeforeClose Error:", err)
				}
			}()
			if m.auxTools != nil {
				if m.auxTools.viewSourceWindow != nil {
					m.auxTools.viewSourceWindow = nil
				}
				if m.auxTools.devToolsWindow != nil {
					m.auxTools.devToolsWindow.Close()
				}
			}
			BrowserWindow.removeWindowInfo(m.windowId)
			//主窗口关闭
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				if IsWindows() {
					rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
				} else {
					m.Close()
				}
			} else if IsDarwin() {
				m.Close()
			}
		}
		QueueAsyncCall(func(id int) { // main thread run
			closeWindow()
		})
		if bwEvent.onBeforeClose != nil {
			bwEvent.onBeforeClose(sender, browser)
		}
	})
}
