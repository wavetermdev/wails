package events

type ApplicationEventType uint
type WindowEventType uint

const (
	FilesDropped WindowEventType = iota
)

var Common = newCommonEvents()

type commonEvents struct {
	ApplicationStarted ApplicationEventType
}

func newCommonEvents() commonEvents {
	return commonEvents{
		ApplicationStarted: 1152,
	}
}

var Mac = newMacEvents()

type macEvents struct {
	ApplicationDidBecomeActive                              ApplicationEventType
	ApplicationDidChangeBackingProperties                   ApplicationEventType
	ApplicationDidChangeEffectiveAppearance                 ApplicationEventType
	ApplicationDidChangeIcon                                ApplicationEventType
	ApplicationDidChangeOcclusionState                      ApplicationEventType
	ApplicationDidChangeScreenParameters                    ApplicationEventType
	ApplicationDidChangeStatusBarFrame                      ApplicationEventType
	ApplicationDidChangeStatusBarOrientation                ApplicationEventType
	ApplicationDidFinishLaunching                           ApplicationEventType
	ApplicationDidHide                                      ApplicationEventType
	ApplicationDidResignActive                              ApplicationEventType
	ApplicationDidUnhide                                    ApplicationEventType
	ApplicationDidUpdate                                    ApplicationEventType
	ApplicationWillBecomeActive                             ApplicationEventType
	ApplicationWillFinishLaunching                          ApplicationEventType
	ApplicationWillHide                                     ApplicationEventType
	ApplicationWillResignActive                             ApplicationEventType
	ApplicationWillTerminate                                ApplicationEventType
	ApplicationWillUnhide                                   ApplicationEventType
	ApplicationWillUpdate                                   ApplicationEventType
	WindowDidBecomeKey                                      WindowEventType
	WindowDidBecomeMain                                     WindowEventType
	WindowDidBeginSheet                                     WindowEventType
	WindowDidChangeAlpha                                    WindowEventType
	WindowDidChangeBackingLocation                          WindowEventType
	WindowDidChangeBackingProperties                        WindowEventType
	WindowDidChangeCollectionBehavior                       WindowEventType
	WindowDidChangeEffectiveAppearance                      WindowEventType
	WindowDidChangeOcclusionState                           WindowEventType
	WindowDidChangeOrderingMode                             WindowEventType
	WindowDidChangeScreen                                   WindowEventType
	WindowDidChangeScreenParameters                         WindowEventType
	WindowDidChangeScreenProfile                            WindowEventType
	WindowDidChangeScreenSpace                              WindowEventType
	WindowDidChangeScreenSpaceProperties                    WindowEventType
	WindowDidChangeSharingType                              WindowEventType
	WindowDidChangeSpace                                    WindowEventType
	WindowDidChangeSpaceOrderingMode                        WindowEventType
	WindowDidChangeTitle                                    WindowEventType
	WindowDidChangeToolbar                                  WindowEventType
	WindowDidChangeVisibility                               WindowEventType
	WindowDidDeminiaturize                                  WindowEventType
	WindowDidEndSheet                                       WindowEventType
	WindowDidEnterFullScreen                                WindowEventType
	WindowDidEnterVersionBrowser                            WindowEventType
	WindowDidExitFullScreen                                 WindowEventType
	WindowDidExitVersionBrowser                             WindowEventType
	WindowDidExpose                                         WindowEventType
	WindowDidFocus                                          WindowEventType
	WindowDidMiniaturize                                    WindowEventType
	WindowDidMove                                           WindowEventType
	WindowDidOrderOffScreen                                 WindowEventType
	WindowDidOrderOnScreen                                  WindowEventType
	WindowDidResignKey                                      WindowEventType
	WindowDidResignMain                                     WindowEventType
	WindowDidResize                                         WindowEventType
	WindowDidUnfocus                                        WindowEventType
	WindowDidUpdate                                         WindowEventType
	WindowDidUpdateAlpha                                    WindowEventType
	WindowDidUpdateCollectionBehavior                       WindowEventType
	WindowDidUpdateCollectionProperties                     WindowEventType
	WindowDidUpdateShadow                                   WindowEventType
	WindowDidUpdateTitle                                    WindowEventType
	WindowDidUpdateToolbar                                  WindowEventType
	WindowDidUpdateVisibility                               WindowEventType
	WindowWillBecomeKey                                     WindowEventType
	WindowWillBecomeMain                                    WindowEventType
	WindowWillBeginSheet                                    WindowEventType
	WindowWillChangeOrderingMode                            WindowEventType
	WindowWillClose                                         WindowEventType
	WindowWillDeminiaturize                                 WindowEventType
	WindowWillEnterFullScreen                               WindowEventType
	WindowWillEnterVersionBrowser                           WindowEventType
	WindowWillExitFullScreen                                WindowEventType
	WindowWillExitVersionBrowser                            WindowEventType
	WindowWillFocus                                         WindowEventType
	WindowWillMiniaturize                                   WindowEventType
	WindowWillMove                                          WindowEventType
	WindowWillOrderOffScreen                                WindowEventType
	WindowWillOrderOnScreen                                 WindowEventType
	WindowWillResignMain                                    WindowEventType
	WindowWillResize                                        WindowEventType
	WindowWillUnfocus                                       WindowEventType
	WindowWillUpdate                                        WindowEventType
	WindowWillUpdateAlpha                                   WindowEventType
	WindowWillUpdateCollectionBehavior                      WindowEventType
	WindowWillUpdateCollectionProperties                    WindowEventType
	WindowWillUpdateShadow                                  WindowEventType
	WindowWillUpdateTitle                                   WindowEventType
	WindowWillUpdateToolbar                                 WindowEventType
	WindowWillUpdateVisibility                              WindowEventType
	WindowWillUseStandardFrame                              WindowEventType
	MenuWillOpen                                            ApplicationEventType
	MenuDidOpen                                             ApplicationEventType
	MenuDidClose                                            ApplicationEventType
	MenuWillSendAction                                      ApplicationEventType
	MenuDidSendAction                                       ApplicationEventType
	MenuWillHighlightItem                                   ApplicationEventType
	MenuDidHighlightItem                                    ApplicationEventType
	MenuWillDisplayItem                                     ApplicationEventType
	MenuDidDisplayItem                                      ApplicationEventType
	MenuWillAddItem                                         ApplicationEventType
	MenuDidAddItem                                          ApplicationEventType
	MenuWillRemoveItem                                      ApplicationEventType
	MenuDidRemoveItem                                       ApplicationEventType
	MenuWillBeginTracking                                   ApplicationEventType
	MenuDidBeginTracking                                    ApplicationEventType
	MenuWillEndTracking                                     ApplicationEventType
	MenuDidEndTracking                                      ApplicationEventType
	MenuWillUpdate                                          ApplicationEventType
	MenuDidUpdate                                           ApplicationEventType
	MenuWillPopUp                                           ApplicationEventType
	MenuDidPopUp                                            ApplicationEventType
	MenuWillSendActionToItem                                ApplicationEventType
	MenuDidSendActionToItem                                 ApplicationEventType
	WebViewDidStartProvisionalNavigation                    WindowEventType
	WebViewDidReceiveServerRedirectForProvisionalNavigation WindowEventType
	WebViewDidFinishNavigation                              WindowEventType
	WebViewDidCommitNavigation                              WindowEventType
	WindowFileDraggingEntered                               WindowEventType
	WindowFileDraggingPerformed                             WindowEventType
	WindowFileDraggingExited                                WindowEventType
}

func newMacEvents() macEvents {
	return macEvents{
		ApplicationDidBecomeActive:               1024,
		ApplicationDidChangeBackingProperties:    1025,
		ApplicationDidChangeEffectiveAppearance:  1026,
		ApplicationDidChangeIcon:                 1027,
		ApplicationDidChangeOcclusionState:       1028,
		ApplicationDidChangeScreenParameters:     1029,
		ApplicationDidChangeStatusBarFrame:       1030,
		ApplicationDidChangeStatusBarOrientation: 1031,
		ApplicationDidFinishLaunching:            1032,
		ApplicationDidHide:                       1033,
		ApplicationDidResignActive:               1034,
		ApplicationDidUnhide:                     1035,
		ApplicationDidUpdate:                     1036,
		ApplicationWillBecomeActive:              1037,
		ApplicationWillFinishLaunching:           1038,
		ApplicationWillHide:                      1039,
		ApplicationWillResignActive:              1040,
		ApplicationWillTerminate:                 1041,
		ApplicationWillUnhide:                    1042,
		ApplicationWillUpdate:                    1043,
		WindowDidBecomeKey:                       1044,
		WindowDidBecomeMain:                      1045,
		WindowDidBeginSheet:                      1046,
		WindowDidChangeAlpha:                     1047,
		WindowDidChangeBackingLocation:           1048,
		WindowDidChangeBackingProperties:         1049,
		WindowDidChangeCollectionBehavior:        1050,
		WindowDidChangeEffectiveAppearance:       1051,
		WindowDidChangeOcclusionState:            1052,
		WindowDidChangeOrderingMode:              1053,
		WindowDidChangeScreen:                    1054,
		WindowDidChangeScreenParameters:          1055,
		WindowDidChangeScreenProfile:             1056,
		WindowDidChangeScreenSpace:               1057,
		WindowDidChangeScreenSpaceProperties:     1058,
		WindowDidChangeSharingType:               1059,
		WindowDidChangeSpace:                     1060,
		WindowDidChangeSpaceOrderingMode:         1061,
		WindowDidChangeTitle:                     1062,
		WindowDidChangeToolbar:                   1063,
		WindowDidChangeVisibility:                1064,
		WindowDidDeminiaturize:                   1065,
		WindowDidEndSheet:                        1066,
		WindowDidEnterFullScreen:                 1067,
		WindowDidEnterVersionBrowser:             1068,
		WindowDidExitFullScreen:                  1069,
		WindowDidExitVersionBrowser:              1070,
		WindowDidExpose:                          1071,
		WindowDidFocus:                           1072,
		WindowDidMiniaturize:                     1073,
		WindowDidMove:                            1074,
		WindowDidOrderOffScreen:                  1075,
		WindowDidOrderOnScreen:                   1076,
		WindowDidResignKey:                       1077,
		WindowDidResignMain:                      1078,
		WindowDidResize:                          1079,
		WindowDidUnfocus:                         1080,
		WindowDidUpdate:                          1081,
		WindowDidUpdateAlpha:                     1082,
		WindowDidUpdateCollectionBehavior:        1083,
		WindowDidUpdateCollectionProperties:      1084,
		WindowDidUpdateShadow:                    1085,
		WindowDidUpdateTitle:                     1086,
		WindowDidUpdateToolbar:                   1087,
		WindowDidUpdateVisibility:                1088,
		WindowWillBecomeKey:                      1089,
		WindowWillBecomeMain:                     1090,
		WindowWillBeginSheet:                     1091,
		WindowWillChangeOrderingMode:             1092,
		WindowWillClose:                          1093,
		WindowWillDeminiaturize:                  1094,
		WindowWillEnterFullScreen:                1095,
		WindowWillEnterVersionBrowser:            1096,
		WindowWillExitFullScreen:                 1097,
		WindowWillExitVersionBrowser:             1098,
		WindowWillFocus:                          1099,
		WindowWillMiniaturize:                    1100,
		WindowWillMove:                           1101,
		WindowWillOrderOffScreen:                 1102,
		WindowWillOrderOnScreen:                  1103,
		WindowWillResignMain:                     1104,
		WindowWillResize:                         1105,
		WindowWillUnfocus:                        1106,
		WindowWillUpdate:                         1107,
		WindowWillUpdateAlpha:                    1108,
		WindowWillUpdateCollectionBehavior:       1109,
		WindowWillUpdateCollectionProperties:     1110,
		WindowWillUpdateShadow:                   1111,
		WindowWillUpdateTitle:                    1112,
		WindowWillUpdateToolbar:                  1113,
		WindowWillUpdateVisibility:               1114,
		WindowWillUseStandardFrame:               1115,
		MenuWillOpen:                             1116,
		MenuDidOpen:                              1117,
		MenuDidClose:                             1118,
		MenuWillSendAction:                       1119,
		MenuDidSendAction:                        1120,
		MenuWillHighlightItem:                    1121,
		MenuDidHighlightItem:                     1122,
		MenuWillDisplayItem:                      1123,
		MenuDidDisplayItem:                       1124,
		MenuWillAddItem:                          1125,
		MenuDidAddItem:                           1126,
		MenuWillRemoveItem:                       1127,
		MenuDidRemoveItem:                        1128,
		MenuWillBeginTracking:                    1129,
		MenuDidBeginTracking:                     1130,
		MenuWillEndTracking:                      1131,
		MenuDidEndTracking:                       1132,
		MenuWillUpdate:                           1133,
		MenuDidUpdate:                            1134,
		MenuWillPopUp:                            1135,
		MenuDidPopUp:                             1136,
		MenuWillSendActionToItem:                 1137,
		MenuDidSendActionToItem:                  1138,
		WebViewDidStartProvisionalNavigation:     1139,
		WebViewDidReceiveServerRedirectForProvisionalNavigation: 1140,
		WebViewDidFinishNavigation:                              1141,
		WebViewDidCommitNavigation:                              1142,
		WindowFileDraggingEntered:                               1143,
		WindowFileDraggingPerformed:                             1144,
		WindowFileDraggingExited:                                1145,
	}
}

var Windows = newWindowsEvents()

type windowsEvents struct {
	SystemThemeChanged    ApplicationEventType
	APMPowerStatusChange  ApplicationEventType
	APMSuspend            ApplicationEventType
	APMResumeAutomatic    ApplicationEventType
	APMResumeSuspend      ApplicationEventType
	APMPowerSettingChange ApplicationEventType
}

func newWindowsEvents() windowsEvents {
	return windowsEvents{
		SystemThemeChanged:    1146,
		APMPowerStatusChange:  1147,
		APMSuspend:            1148,
		APMResumeAutomatic:    1149,
		APMResumeSuspend:      1150,
		APMPowerSettingChange: 1151,
	}
}
