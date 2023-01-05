//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	t "github.com/energye/energy/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

func init() {
	var resourceEventGet = func(fn interface{}, getVal func(idx int) uintptr, resp bool) (sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) {
		var (
			instance unsafe.Pointer
		)
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		senderPtr := getPtr(0)
		browser = &ICefBrowser{browseId: int32(getVal(1)), chromium: senderPtr}
		tempFrame := (*cefFrame)(getPtr(2))
		frame = &ICefFrame{
			Browser: browser,
			Name:    api.GoStr(tempFrame.Name),
			Url:     api.GoStr(tempFrame.Url),
			Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
		}
		cefRequest := (*rICefRequest)(getPtr(3))
		instance = GetInstancePtr(cefRequest.Instance)
		request = &ICefRequest{
			instance:             instance,
			Url:                  api.GoStr(cefRequest.Url),
			Method:               api.GoStr(cefRequest.Method),
			ReferrerUrl:          api.GoStr(cefRequest.ReferrerUrl),
			ReferrerPolicy:       TCefReferrerPolicy(cefRequest.ReferrerPolicy),
			Flags:                TCefUrlRequestFlags(cefRequest.Flags),
			FirstPartyForCookies: api.GoStr(cefRequest.FirstPartyForCookies),
			ResourceType:         TCefResourceType(cefRequest.ResourceType),
			TransitionType:       TCefTransitionType(cefRequest.TransitionType),
			Identifier:           *(*uint64)(GetParamPtr(cefRequest.Identifier, 0)),
		}
		if resp {
			cefResponse := (*iCefResponse)(getPtr(4))
			instance = GetInstancePtr(cefResponse.Instance)
			response = &ICefResponse{
				instance:   instance,
				Status:     int32(cefResponse.Status),
				StatusText: api.GoStr(cefResponse.StatusText),
				MimeType:   api.GoStr(cefResponse.MimeType),
				Charset:    api.GoStr(cefResponse.Charset),
				Error:      TCefErrorCode(cefResponse.Error),
				URL:        api.GoStr(cefResponse.URL),
			}
		}
		return lcl.AsObject(senderPtr), browser, frame, request, response
	}
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		//defer func() {
		//	if err := recover(); err != nil {
		//		logger.Error("CEF Events Error:", err)
		//	}
		//}()
		var (
			instance unsafe.Pointer
		)
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case ChromiumEventOnFindResult:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			cefRectPtr := (*tCefRectPtr)(getPtr(4))
			cefRect := &TCefRect{
				X:      int32(cefRectPtr.X),
				Y:      int32(cefRectPtr.Y),
				Width:  int32(cefRectPtr.Width),
				Height: int32(cefRectPtr.Height),
			}
			fn.(ChromiumEventOnFindResult)(lcl.AsObject(sender), browser, int32(getVal(2)), int32(getVal(3)), cefRect, int32(getVal(5)), api.GoBool(getVal(6)))
		case BrowseProcessMessageReceived:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			cefProcMsg := (*ipc.CefProcessMessagePtr)(getPtr(4))
			args := ipc.NewArgumentList()
			args.UnPackageBytePtr(cefProcMsg.Data, int32(cefProcMsg.DataLen))
			processMessage := &ipc.ICefProcessMessage{
				Name:         api.GoStr(cefProcMsg.Name),
				ArgumentList: args,
			}
			var result = (*bool)(getPtr(5))
			*result = fn.(BrowseProcessMessageReceived)(lcl.AsObject(sender), browser, frame, CefProcessId(getVal(3)), processMessage)
			args.Clear()
			cefProcMsg.Data = 0
			cefProcMsg.DataLen = 0
			cefProcMsg.Name = 0
			cefProcMsg = nil
			args = nil
		case ChromiumEventOnResourceLoadComplete:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(ChromiumEventOnResourceLoadComplete)(sender, browse, frame, request, response, *(*TCefUrlRequestStatus)(getPtr(5)), *(*int64)(getPtr(6)))
		case ChromiumEventOnResourceRedirect:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			var newStr = new(t.TString)
			var newStrPtr = (*uintptr)(getPtr(5))
			fn.(ChromiumEventOnResourceRedirect)(sender, browse, frame, request, response, newStr)
			*newStrPtr = newStr.ToPtr()
		case ChromiumEventOnResourceResponse:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(ChromiumEventOnResourceResponse)(sender, browse, frame, request, response, (*bool)(getPtr(5)))
		case ChromiumEventOnBeforeResourceLoad:
			sender, browse, frame, request, _ := resourceEventGet(fn, getVal, false)
			instance = getInstance(getVal(4))
			callback := &ICefCallback{instance: instance}
			fn.(ChromiumEventOnBeforeResourceLoad)(sender, browse, frame, request, callback, (*TCefReturnValue)(getPtr(5)))
		//menu begin
		case ChromiumEventOnBeforeContextMenu:
			sender := getPtr(0)
			instance = getInstance(getVal(1))
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			cefParams := (*iCefContextMenuParams)(getPtr(3))
			params := &ICefContextMenuParams{
				XCoord:            int32(cefParams.XCoord),
				YCoord:            int32(cefParams.YCoord),
				TypeFlags:         TCefContextMenuTypeFlags(cefParams.TypeFlags),
				LinkUrl:           api.GoStr(cefParams.LinkUrl),
				UnfilteredLinkUrl: api.GoStr(cefParams.UnfilteredLinkUrl),
				SourceUrl:         api.GoStr(cefParams.SourceUrl),
				TitleText:         api.GoStr(cefParams.TitleText),
				PageUrl:           api.GoStr(cefParams.PageUrl),
				FrameUrl:          api.GoStr(cefParams.FrameUrl),
				FrameCharset:      api.GoStr(cefParams.FrameCharset),
				MediaType:         TCefContextMenuMediaType(cefParams.MediaType),
				MediaStateFlags:   TCefContextMenuMediaStateFlags(cefParams.MediaStateFlags),
				SelectionText:     api.GoStr(cefParams.SelectionText),
				EditStateFlags:    TCefContextMenuEditStateFlags(cefParams.EditStateFlags),
			}
			instance = getInstance(getVal(4))
			KeyAccelerator.clear()
			model := &ICefMenuModel{instance: instance, CefMis: KeyAccelerator}
			fn.(ChromiumEventOnBeforeContextMenu)(lcl.AsObject(sender), browser, frame, params, model)
		case ChromiumEventOnContextMenuCommand:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			cefParams := (*iCefContextMenuParams)(getPtr(3))
			params := &ICefContextMenuParams{
				XCoord:            int32(cefParams.XCoord),
				YCoord:            int32(cefParams.YCoord),
				TypeFlags:         TCefContextMenuTypeFlags(cefParams.TypeFlags),
				LinkUrl:           api.GoStr(cefParams.LinkUrl),
				UnfilteredLinkUrl: api.GoStr(cefParams.UnfilteredLinkUrl),
				SourceUrl:         api.GoStr(cefParams.SourceUrl),
				TitleText:         api.GoStr(cefParams.TitleText),
				PageUrl:           api.GoStr(cefParams.PageUrl),
				FrameUrl:          api.GoStr(cefParams.FrameUrl),
				FrameCharset:      api.GoStr(cefParams.FrameCharset),
				MediaType:         TCefContextMenuMediaType(cefParams.MediaType),
				MediaStateFlags:   TCefContextMenuMediaStateFlags(cefParams.MediaStateFlags),
				SelectionText:     api.GoStr(cefParams.SelectionText),
				EditStateFlags:    TCefContextMenuEditStateFlags(cefParams.EditStateFlags),
			}
			commandId := MenuId(getVal(4))
			eventFlags := uint32(getVal(5))
			if !KeyAccelerator.commandIdEventCallback(browser, commandId, params, eventFlags, (*bool)(getPtr(5))) {
				fn.(ChromiumEventOnContextMenuCommand)(lcl.AsObject(sender), browser, frame, params, commandId, eventFlags, (*bool)(getPtr(6)))
			}
		case ChromiumEventOnContextMenuDismissed:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnContextMenuDismissed)(lcl.AsObject(sender), browser, frame)
		//menu end
		//---
		//cookie begin
		case ChromiumEventOnCookieSet:
			success := api.GoBool(getVal(1))
			ID := int32(getVal(2))
			fn.(ChromiumEventOnCookieSet)(lcl.AsObject(getVal(0)), success, ID)
		case ChromiumEventOnCookiesDeleted:
			numDeleted := int32(getVal(1))
			fn.(ChromiumEventOnCookiesDeleted)(lcl.AsObject(getVal(0)), numDeleted)
		case ChromiumEventOnCookiesFlushed:
			fn.(ChromiumEventOnCookiesFlushed)(lcl.AsObject(getVal(0)))
		case ChromiumEventOnCookiesVisited:
			cookie := *(*iCefCookiePtr)(getPtr(1))
			creation := *(*float64)(GetParamPtr(cookie.creation, 0))
			lastAccess := *(*float64)(GetParamPtr(cookie.lastAccess, 0))
			expires := *(*float64)(GetParamPtr(cookie.expires, 0))
			iCookie := &ICefCookie{
				Url:            api.GoStr(cookie.url),
				Name:           api.GoStr(cookie.name),
				Value:          api.GoStr(cookie.value),
				Domain:         api.GoStr(cookie.domain),
				Path:           api.GoStr(cookie.path),
				Secure:         *(*bool)(GetParamPtr(cookie.secure, 0)),
				Httponly:       *(*bool)(GetParamPtr(cookie.httponly, 0)),
				HasExpires:     *(*bool)(GetParamPtr(cookie.hasExpires, 0)),
				Creation:       DDateTimeToGoDateTime(creation),
				LastAccess:     DDateTimeToGoDateTime(lastAccess),
				Expires:        DDateTimeToGoDateTime(expires),
				Count:          int32(cookie.count),
				Total:          int32(cookie.total),
				ID:             int32(cookie.aID),
				SameSite:       TCefCookieSameSite(cookie.sameSite),
				Priority:       TCefCookiePriority(cookie.priority),
				SetImmediately: *(*bool)(GetParamPtr(cookie.aSetImmediately, 0)),
				DeleteCookie:   *(*bool)(GetParamPtr(cookie.aDeleteCookie, 0)),
				Result:         *(*bool)(GetParamPtr(cookie.aResult, 0)),
			}
			fn.(ChromiumEventOnCookiesVisited)(lcl.AsObject(getVal(0)), iCookie)
		case ChromiumEventOnCookieVisitorDestroyed:
			id := int32(getVal(1))
			fn.(ChromiumEventOnCookieVisitorDestroyed)(lcl.AsObject(getVal(0)), id)
		//cookie end
		//--- other
		case ChromiumEventOnScrollOffsetChanged:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnScrollOffsetChanged)(lcl.AsObject(sender), browser, float64(getVal(2)), float64(getVal(2)))
		case ChromiumEventOnRenderProcessTerminated:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnRenderProcessTerminated)(lcl.AsObject(sender), browser, TCefTerminationStatus(getVal(2)))
		case ChromiumEventOnRenderCompMsg:
			message := *(*types.TMessage)(getPtr(1))
			fn.(ChromiumEventOnRenderCompMsg)(lcl.AsObject(getVal(0)), message, api.GoBool(getVal(2)))
		case ChromiumEventOnCefBrowser:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnCefBrowser)(lcl.AsObject(sender), browser)
		case ChromiumEventOnTitleChange:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnTitleChange)(lcl.AsObject(sender), browser, api.GoStr(getVal(2)))
		case ChromiumEventOnKeyEvent:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			keyEvent := (*TCefKeyEvent)(getPtr(2))
			fn.(ChromiumEventOnKeyEvent)(lcl.AsObject(sender), browser, keyEvent, (*bool)(getPtr(3)))
		case ChromiumEventOnFullScreenModeChange:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnFullScreenModeChange)(lcl.AsObject(sender), browser, api.GoBool(getVal(2)))
		case ChromiumEventOnBeforeBrowser: //创建浏览器之前
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			var result = (*bool)(getPtr(3))
			*result = fn.(ChromiumEventOnBeforeBrowser)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnAddressChange: //创建浏览器之前
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnAddressChange)(lcl.AsObject(sender), browser, frame, api.GoStr(getVal(3)))
		case ChromiumEventOnAfterCreated: //创建浏览器之后
			sender := getPtr(0)
			//事件处理函数返回true将不继续执行
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnAfterCreated)(lcl.AsObject(sender), browser)
		case ChromiumEventOnBeforeClose: //关闭浏览器之前
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnBeforeClose)(lcl.AsObject(sender), browser)
		case ChromiumEventOnClose:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnClose)(lcl.AsObject(sender), browser, (*TCefCloseBrowsesAction)(getPtr(2)))
		case ChromiumEventOnResult: //通用Result bool事件
			fn.(ChromiumEventOnResult)(lcl.AsObject(getVal(0)), api.GoBool(getVal(1)))
		case ChromiumEventOnResultFloat: //通用Result float事件
			fn.(ChromiumEventOnResultFloat)(lcl.AsObject(getVal(0)), *(*float64)(getPtr(1)))
		case ChromiumEventOnLoadStart:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnLoadStart)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnLoadingStateChange:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnLoadingStateChange)(lcl.AsObject(sender), browser, api.GoBool(getVal(2)), api.GoBool(getVal(3)), api.GoBool(getVal(4)))
		case ChromiumEventOnLoadingProgressChange:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnLoadingProgressChange)(lcl.AsObject(sender), browser, *(*float64)(getPtr(2)))
		case ChromiumEventOnLoadError:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnLoadError)(lcl.AsObject(sender), browser, frame, CEF_NET_ERROR(getVal(3)), api.GoStr(getVal(4)), api.GoStr(getVal(5)))
		case ChromiumEventOnLoadEnd:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnLoadEnd)(lcl.AsObject(sender), browser, frame, int32(getVal(3)))
		case ChromiumEventOnBeforeDownload: //下载之前
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			item := (*downloadItem)(getPtr(2))
			downItem := &DownloadItem{
				Id:                 int32(item.Id),
				CurrentSpeed:       int64(item.CurrentSpeed),
				PercentComplete:    int32(item.PercentComplete),
				TotalBytes:         int64(item.TotalBytes),
				ReceivedBytes:      int64(item.ReceivedBytes),
				StartTime:          DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.StartTime, 0))),
				EndTime:            DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.EndTime, 0))),
				FullPath:           api.GoStr(item.FullPath),
				Url:                api.GoStr(item.Url),
				OriginalUrl:        api.GoStr(item.OriginalUrl),
				SuggestedFileName:  api.GoStr(item.SuggestedFileName),
				ContentDisposition: api.GoStr(item.ContentDisposition),
				MimeType:           api.GoStr(item.MimeType),
				IsValid:            *(*bool)(unsafe.Pointer(item.IsValid)),
				State:              int32(item.State),
			}
			suggestedName := api.GoStr(getVal(3))
			instance = getInstance(getVal(4))
			callback := &ICefBeforeDownloadCallback{
				instance: instance,
				browseId: browser.Identifier(),
				downId:   downItem.Id,
			}
			fn.(ChromiumEventOnBeforeDownload)(lcl.AsObject(sender), browser, downItem, suggestedName, callback)
		case ChromiumEventOnDownloadUpdated: //下载更新
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			item := *(*downloadItem)(getPtr(2))
			downItem := &DownloadItem{
				Id:                 int32(item.Id),
				CurrentSpeed:       int64(item.CurrentSpeed),
				PercentComplete:    int32(item.PercentComplete),
				TotalBytes:         int64(item.TotalBytes),
				ReceivedBytes:      int64(item.ReceivedBytes),
				StartTime:          DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.StartTime, 0))),
				EndTime:            DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.EndTime, 0))),
				FullPath:           api.GoStr(item.FullPath),
				Url:                api.GoStr(item.Url),
				OriginalUrl:        api.GoStr(item.OriginalUrl),
				SuggestedFileName:  api.GoStr(item.SuggestedFileName),
				ContentDisposition: api.GoStr(item.ContentDisposition),
				MimeType:           api.GoStr(item.MimeType),
				IsValid:            *(*bool)(unsafe.Pointer(item.IsValid)),
				State:              int32(item.State),
			}
			instance = getInstance(getVal(3))
			callback := &ICefDownloadItemCallback{
				instance: instance,
				browseId: browser.Identifier(),
				downId:   downItem.Id,
			}
			fn.(ChromiumEventOnDownloadUpdated)(lcl.AsObject(sender), browser, downItem, callback)
		//frame
		case ChromiumEventOnFrameAttached:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnFrameAttached)(lcl.AsObject(sender), browser, frame, api.GoBool(getVal(3)))
		case ChromiumEventOnFrameCreated:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
			}
			fn.(ChromiumEventOnFrameCreated)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnFrameDetached:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1))}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnFrameDetached)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnMainFrameChanged:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			var (
				oldFrame *ICefFrame = nil
				newFrame *ICefFrame = nil
			)
			tempOldFrame := (*cefFrame)(getPtr(2))
			if tempOldFrame != nil {
				oldFrame = &ICefFrame{
					Browser: browser,
					Name:    api.GoStr(tempOldFrame.Name),
					Url:     api.GoStr(tempOldFrame.Url),
					Id:      StrToInt64(api.GoStr(tempOldFrame.Identifier)),
				}
			}
			tempNewFrame := (*cefFrame)(getPtr(3))
			if tempNewFrame != nil {
				newFrame = &ICefFrame{
					Browser: browser,
					Name:    api.GoStr(tempNewFrame.Name),
					Url:     api.GoStr(tempNewFrame.Url),
					Id:      StrToInt64(api.GoStr(tempNewFrame.Identifier)),
				}
			}
			fn.(ChromiumEventOnMainFrameChanged)(lcl.AsObject(sender), browser, oldFrame, newFrame)
		//windowParent popup
		case ChromiumEventOnBeforePopup:
			sender := getPtr(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			beforePInfoPtr := (*beforePopupInfoPtr)(getPtr(3))
			beforePInfo := &BeforePopupInfo{
				TargetUrl:         api.GoStr(beforePInfoPtr.TargetUrl),
				TargetFrameName:   api.GoStr(beforePInfoPtr.TargetFrameName),
				TargetDisposition: TCefWindowOpenDisposition(beforePInfoPtr.TargetDisposition),
				UserGesture:       api.GoBool(beforePInfoPtr.UserGesture),
			}

			var (
				client             = &ICefClient{instance: getPtr(5)}
				noJavascriptAccess = (*bool)(getPtr(6))
				result             = (*bool)(getPtr(7))
			)
			//callback
			*result = fn.(ChromiumEventOnBeforePopup)(lcl.AsObject(sender), browser, frame, beforePInfo, client, noJavascriptAccess)
		//windowParent open url from tab
		case ChromiumEventOnOpenUrlFromTab:

		default:
			return false
		}
		return true
	})
}

func getInstance(value interface{}) unsafe.Pointer {
	var ptr uintptr
	switch value.(type) {
	case uintptr:
		ptr = value.(uintptr)
	case unsafe.Pointer:
		ptr = uintptr(value.(unsafe.Pointer))
	case lcl.IObject:
		ptr = lcl.CheckPtr(value)
	default:
		ptr = getUIntPtr(value)
	}
	return unsafe.Pointer(ptr)
}

func getUIntPtr(v interface{}) uintptr {
	switch v.(type) {
	case int:
		return uintptr(v.(int))
	case uint:
		return uintptr(v.(uint))
	case int8:
		return uintptr(v.(int8))
	case uint8:
		return uintptr(v.(uint8))
	case int16:
		return uintptr(v.(int16))
	case uint16:
		return uintptr(v.(uint16))
	case int32:
		return uintptr(v.(int32))
	case uint32:
		return uintptr(v.(uint32))
	case int64:
		return uintptr(v.(int64))
	case uint64:
		return uintptr(v.(uint64))
	}
	return 0
}
