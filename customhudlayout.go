package customhud

/*
#include "customhudlayout.h"
#cgo noescape IsCsScriptReady
#cgo noescape CreateCustomHud
#cgo noescape RemoveCustomHud
#cgo noescape HideCustomHudFromOtherPlayers
#cgo noescape SetHudHasClass
#cgo noescape SetHudDialogVariable
#cgo noescape SetHudHasClassForPlayer
#cgo noescape SetHudDialogVariableForPlayer
#cgo noescape SetHudInputCapture
#cgo noescape IsHudInputCaptureEnabled
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from customhud (group: customhudlayout)

var _IsCsScriptReady = func() bool {
	__retVal := bool(C.IsCsScriptReady())
	return __retVal
}

// IsCsScriptReady 
//  @brief Checks whether the CS Script system has come up and the plugin is connected to it.
//
//
//  @return True if Instance is available.
func IsCsScriptReady() bool {
	return _IsCsScriptReady()
}

var _CreateCustomHud = func(name string, layoutResource string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__layoutResource := plugify.ConstructString(layoutResource)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.CreateCustomHud((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__layoutResource))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__layoutResource)
		},
	}.Do()
	return __retVal
}

// CreateCustomHud 
//  @brief Creates a custom_hud_layout entity with the given Panorama layout. cs_script cannot spawn entities, so creation goes through s2sdk.
//
//  @param name: Name (targetname) used for all further calls to this hud.
//  @param layoutResource: Path to the Panorama layout (.xml), declared as a resource in the session manifest.
//
//  @return False if the engine doesn't know the custom_hud_layout class.
func CreateCustomHud(name string, layoutResource string) bool {
	return _CreateCustomHud(name, layoutResource)
}

var _RemoveCustomHud = func(name string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.RemoveCustomHud((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// RemoveCustomHud 
//  @brief Removes a previously created custom_hud_layout.
//
//  @param name: Name passed to CreateCustomHud.
//
//  @return False if no hud with this name was created by this plugin.
func RemoveCustomHud(name string) bool {
	return _RemoveCustomHud(name)
}

var _HideCustomHudFromOtherPlayers = func(name string, playerSlot int32) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.HideCustomHudFromOtherPlayers((*C.String)(unsafe.Pointer(&__name)), __playerSlot))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// HideCustomHudFromOtherPlayers 
//  @brief Hides the hud entity from all players except the owner, at the transmit/PVS level — a stronger guarantee than a CSS class, which only hides the panel on clients the entity is still transmitted to. (s2sdk.HideTransmitEntityFromOtherPlayers)
//
//  @param name: Name passed to CreateCustomHud.
//  @param playerSlot: The owner player slot who will still see the entity.
//
//  @return False if no hud with this name was created by this plugin.
func HideCustomHudFromOtherPlayers(name string, playerSlot int32) bool {
	return _HideCustomHudFromOtherPlayers(name, playerSlot)
}

var _SetHudHasClass = func(name string, panelId string, className string, hasClass bool) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	__hasClass := C.bool(hasClass)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudHasClass((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className)), __hasClass))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// SetHudHasClass 
//  @brief Set if a panel has a class. Applies to all players. Omit `hasClass` to revert to the original value. (CustomHudLayout.SetHasClass)
//
//  @param name: Name passed to CreateCustomHud.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//  @param hasClass: True to add the class, false to remove it.
//
//  @return False if the entity isn't found or CS Script isn't ready.
func SetHudHasClass(name string, panelId string, className string, hasClass bool) bool {
	return _SetHudHasClass(name, panelId, className, hasClass)
}

var _SetHudDialogVariable = func(name string, panelId string, variableName string, value string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__panelId := plugify.ConstructString(panelId)
	__variableName := plugify.ConstructString(variableName)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudDialogVariable((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__variableName)), (*C.String)(unsafe.Pointer(&__value))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__variableName)
			plugify.DestroyString(&__value)
		},
	}.Do()
	return __retVal
}

// SetHudDialogVariable 
//  @brief Set the value of a dialog variable. Applies to all players. In the layout it's read as text="{s:variableName}". (CustomHudLayout.SetDialogVariableString)
//
//  @param name: Name passed to CreateCustomHud.
//  @param panelId: The panel's id attribute in the layout.
//  @param variableName: Variable name.
//  @param value: Value to display.
//
//  @return False if the entity isn't found or CS Script isn't ready.
func SetHudDialogVariable(name string, panelId string, variableName string, value string) bool {
	return _SetHudDialogVariable(name, panelId, variableName, value)
}

var _SetHudHasClassForPlayer = func(name string, playerSlot int32, panelId string, className string, hasClass bool) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	__hasClass := C.bool(hasClass)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudHasClassForPlayer((*C.String)(unsafe.Pointer(&__name)), __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className)), __hasClass))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// SetHudHasClassForPlayer 
//  @brief Set if a panel has a class for a single player. Will override the all player value. Omit `hasClass` to defer to the all player value. (CustomHudLayout.SetHasClassForPlayer)
//
//  @param name: Name passed to CreateCustomHud.
//  @param playerSlot: Player slot the override applies to.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//  @param hasClass: True to add the class, false to remove it.
//
//  @return False if the entity isn't found or CS Script isn't ready.
func SetHudHasClassForPlayer(name string, playerSlot int32, panelId string, className string, hasClass bool) bool {
	return _SetHudHasClassForPlayer(name, playerSlot, panelId, className, hasClass)
}

var _SetHudDialogVariableForPlayer = func(name string, playerSlot int32, panelId string, variableName string, value string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__variableName := plugify.ConstructString(variableName)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudDialogVariableForPlayer((*C.String)(unsafe.Pointer(&__name)), __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__variableName)), (*C.String)(unsafe.Pointer(&__value))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__variableName)
			plugify.DestroyString(&__value)
		},
	}.Do()
	return __retVal
}

// SetHudDialogVariableForPlayer 
//  @brief Set the value of a dialog variable for a single player. Will override the all player value. Omit `value` to defer to the all player value. If no all player value has been set, the value will be an empty string. (CustomHudLayout.SetDialogVariableStringForPlayer)
//
//  @param name: Name passed to CreateCustomHud.
//  @param playerSlot: Player slot the override applies to.
//  @param panelId: The panel's id attribute in the layout.
//  @param variableName: Variable name.
//  @param value: Value to display.
//
//  @return False if the entity isn't found or CS Script isn't ready.
func SetHudDialogVariableForPlayer(name string, playerSlot int32, panelId string, variableName string, value string) bool {
	return _SetHudDialogVariableForPlayer(name, playerSlot, panelId, variableName, value)
}

var _SetHudInputCapture = func(name string, playerSlot int32, enabled bool) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__playerSlot := C.int32_t(playerSlot)
	__enabled := C.bool(enabled)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudInputCapture((*C.String)(unsafe.Pointer(&__name)), __playerSlot, __enabled))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// SetHudInputCapture 
//  @brief Set to true to force a player into cursor mode and enable click detection on the panels of this hud. Set a callback with OnHudClicked_Register to listen for clicks. Multiple huds can have input captured at a time. Players will get movement control back once all huds have disabled input capture. (CustomHudLayout.SetInputCaptureEnabled)
//
//  @param name: Name passed to CreateCustomHud.
//  @param playerSlot: Player slot.
//  @param enabled: True to capture input, false to release control.
//
//  @return False if the entity isn't found or CS Script isn't ready.
func SetHudInputCapture(name string, playerSlot int32, enabled bool) bool {
	return _SetHudInputCapture(name, playerSlot, enabled)
}

var _IsHudInputCaptureEnabled = func(name string, playerSlot int32) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.IsHudInputCaptureEnabled((*C.String)(unsafe.Pointer(&__name)), __playerSlot))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// IsHudInputCaptureEnabled 
//  @brief Get if this hud is capturing input for a player. (CustomHudLayout.IsInputCaptureEnabled)
//
//  @param name: Name passed to CreateCustomHud.
//  @param playerSlot: Player slot.
//
//  @return False if the entity isn't found, CS Script isn't ready, or input isn't captured.
func IsHudInputCaptureEnabled(name string, playerSlot int32) bool {
	return _IsHudInputCaptureEnabled(name, playerSlot)
}

