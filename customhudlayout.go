package customhud

/*
#include "customhudlayout.h"
#cgo noescape IsCsScriptReady
#cgo noescape CreateCustomHud
#cgo noescape FindCustomHud
#cgo noescape RemoveCustomHud
#cgo noescape HideCustomHudFromOtherPlayers
#cgo noescape SetHudHasClass
#cgo noescape ResetHudHasClass
#cgo noescape SetHudDialogVariable
#cgo noescape SetHudHasClassForPlayer
#cgo noescape ResetHudHasClassForPlayer
#cgo noescape BHasClass
#cgo noescape ToggleClass
#cgo noescape SetHudDialogVariableForPlayer
#cgo noescape ResetHudDialogVariableForPlayer
#cgo noescape SetHudInputCapture
#cgo noescape IsHudInputCaptureEnabled
#cgo noescape ResetHud
#cgo noescape ResetHudForPlayer
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

var _CreateCustomHud = func(name string, layoutResource string) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	__layoutResource := plugify.ConstructString(layoutResource)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.CreateCustomHud((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__layoutResource))))
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
//  @brief Creates a custom_hud_layout entity with the given Panorama layout and returns its handle. The handle is used for all further calls to this hud and stays valid until the hud is removed or the map changes. cs_script cannot spawn entities, so creation goes through s2sdk.
//
//  @param name: Name (targetname) of the hud. Does not have to be unique.
//  @param layoutResource: Path to the Panorama layout (.xml), declared as a resource in the session manifest.
//
//  @return Hud handle, or -1 if the hud could not be created.
func CreateCustomHud(name string, layoutResource string) int32 {
	return _CreateCustomHud(name, layoutResource)
}

var _FindCustomHud = func(name string) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.FindCustomHud((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// FindCustomHud 
//  @brief Gets the handle of the first hud with this name (targetname), such as a hud placed on the map.
//
//  @param name: The hud's targetname.
//
//  @return Hud handle, or -1 if there is no hud with this name.
func FindCustomHud(name string) int32 {
	return _FindCustomHud(name)
}

var _RemoveCustomHud = func(hud int32) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__retVal = bool(C.RemoveCustomHud(__hud))
	return __retVal
}

// RemoveCustomHud 
//  @brief Removes a hud.
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//
//  @return False if the handle isn't a hud.
func RemoveCustomHud(hud int32) bool {
	return _RemoveCustomHud(hud)
}

var _HideCustomHudFromOtherPlayers = func(hud int32, playerSlot int32) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.HideCustomHudFromOtherPlayers(__hud, __playerSlot))
	return __retVal
}

// HideCustomHudFromOtherPlayers 
//  @brief Hides the hud entity from all players except the owner, at the transmit/PVS level. A stronger guarantee than a CSS class, which only hides the panel on clients the entity is still transmitted to. (s2sdk.HideTransmitEntityFromOtherPlayers)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: The owner player slot who will still see the entity.
//
//  @return False if the handle isn't a hud.
func HideCustomHudFromOtherPlayers(hud int32, playerSlot int32) bool {
	return _HideCustomHudFromOtherPlayers(hud, playerSlot)
}

var _SetHudHasClass = func(hud int32, panelId string, className string, hasClass bool) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	__hasClass := C.bool(hasClass)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudHasClass(__hud, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className)), __hasClass))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// SetHudHasClass 
//  @brief Set if a panel has a class. Applies to all players. (CustomHudLayout.SetHasClass)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//  @param hasClass: True to add the class, false to remove it.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func SetHudHasClass(hud int32, panelId string, className string, hasClass bool) bool {
	return _SetHudHasClass(hud, panelId, className, hasClass)
}

var _ResetHudHasClass = func(hud int32, panelId string, className string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.ResetHudHasClass(__hud, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// ResetHudHasClass 
//  @brief Revert a panel's class to the original value from the layout. Applies to all players. (CustomHudLayout.SetHasClass without hasClass)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func ResetHudHasClass(hud int32, panelId string, className string) bool {
	return _ResetHudHasClass(hud, panelId, className)
}

var _SetHudDialogVariable = func(hud int32, panelId string, variableName string, value string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__panelId := plugify.ConstructString(panelId)
	__variableName := plugify.ConstructString(variableName)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudDialogVariable(__hud, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__variableName)), (*C.String)(unsafe.Pointer(&__value))))
		},
		Finally: func() {
			// Perform cleanup.
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
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param panelId: The panel's id attribute in the layout.
//  @param variableName: Variable name.
//  @param value: Value to display.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func SetHudDialogVariable(hud int32, panelId string, variableName string, value string) bool {
	return _SetHudDialogVariable(hud, panelId, variableName, value)
}

var _SetHudHasClassForPlayer = func(hud int32, playerSlot int32, panelId string, className string, hasClass bool) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	__hasClass := C.bool(hasClass)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudHasClassForPlayer(__hud, __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className)), __hasClass))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// SetHudHasClassForPlayer 
//  @brief Set if a panel has a class for a single player. Will override the all player value. (CustomHudLayout.SetHasClassForPlayer)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot the override applies to.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//  @param hasClass: True to add the class, false to remove it.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func SetHudHasClassForPlayer(hud int32, playerSlot int32, panelId string, className string, hasClass bool) bool {
	return _SetHudHasClassForPlayer(hud, playerSlot, panelId, className, hasClass)
}

var _ResetHudHasClassForPlayer = func(hud int32, playerSlot int32, panelId string, className string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.ResetHudHasClassForPlayer(__hud, __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// ResetHudHasClassForPlayer 
//  @brief Remove a single player's value of a panel's class, so the all player value applies again. (CustomHudLayout.SetHasClassForPlayer without hasClass)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot the override applies to.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func ResetHudHasClassForPlayer(hud int32, playerSlot int32, panelId string, className string) bool {
	return _ResetHudHasClassForPlayer(hud, playerSlot, panelId, className)
}

var _BHasClass = func(hud int32, playerSlot int32, panelId string, className string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.BHasClass(__hud, __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// BHasClass 
//  @brief Get if a panel has a class for a player. The player value is used if set, otherwise the all player value. Classes from the layout file are not known to the server and are not reported. (Panel.BHasClass)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot, or -1 for the all player value.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//
//  @return True if the class is set. False if it isn't, or the handle isn't a hud.
func BHasClass(hud int32, playerSlot int32, panelId string, className string) bool {
	return _BHasClass(hud, playerSlot, panelId, className)
}

var _ToggleClass = func(hud int32, playerSlot int32, panelId string, className string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__className := plugify.ConstructString(className)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.ToggleClass(__hud, __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// ToggleClass 
//  @brief Toggle a class on a panel, based on the value reported by BHasClass. Pass -1 for `playerSlot` to toggle the all player value. (Panel.ToggleClass)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot, or -1 for the all player value.
//  @param panelId: The panel's id attribute in the layout.
//  @param className: CSS class.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func ToggleClass(hud int32, playerSlot int32, panelId string, className string) bool {
	return _ToggleClass(hud, playerSlot, panelId, className)
}

var _SetHudDialogVariableForPlayer = func(hud int32, playerSlot int32, panelId string, variableName string, value string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__variableName := plugify.ConstructString(variableName)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetHudDialogVariableForPlayer(__hud, __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__variableName)), (*C.String)(unsafe.Pointer(&__value))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__variableName)
			plugify.DestroyString(&__value)
		},
	}.Do()
	return __retVal
}

// SetHudDialogVariableForPlayer 
//  @brief Set the value of a dialog variable for a single player. Will override the all player value. (CustomHudLayout.SetDialogVariableStringForPlayer)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot the override applies to.
//  @param panelId: The panel's id attribute in the layout.
//  @param variableName: Variable name.
//  @param value: Value to display.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func SetHudDialogVariableForPlayer(hud int32, playerSlot int32, panelId string, variableName string, value string) bool {
	return _SetHudDialogVariableForPlayer(hud, playerSlot, panelId, variableName, value)
}

var _ResetHudDialogVariableForPlayer = func(hud int32, playerSlot int32, panelId string, variableName string) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__panelId := plugify.ConstructString(panelId)
	__variableName := plugify.ConstructString(variableName)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.ResetHudDialogVariableForPlayer(__hud, __playerSlot, (*C.String)(unsafe.Pointer(&__panelId)), (*C.String)(unsafe.Pointer(&__variableName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__panelId)
			plugify.DestroyString(&__variableName)
		},
	}.Do()
	return __retVal
}

// ResetHudDialogVariableForPlayer 
//  @brief Remove a single player's value of a dialog variable, so the all player value applies again. If no all player value has been set, the value will be an empty string. (CustomHudLayout.SetDialogVariableStringForPlayer without value)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot the override applies to.
//  @param panelId: The panel's id attribute in the layout.
//  @param variableName: Variable name.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func ResetHudDialogVariableForPlayer(hud int32, playerSlot int32, panelId string, variableName string) bool {
	return _ResetHudDialogVariableForPlayer(hud, playerSlot, panelId, variableName)
}

var _SetHudInputCapture = func(hud int32, playerSlot int32, enabled bool) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__enabled := C.bool(enabled)
	__retVal = bool(C.SetHudInputCapture(__hud, __playerSlot, __enabled))
	return __retVal
}

// SetHudInputCapture 
//  @brief Set to true to force a player into cursor mode and enable click detection on the panels of this hud. Set a callback with OnHudClicked_Register to listen for clicks. Multiple huds can have input captured at a time. Players will get movement control back once all huds have disabled input capture. (CustomHudLayout.SetInputCaptureEnabled)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot.
//  @param enabled: True to capture input, false to release control.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func SetHudInputCapture(hud int32, playerSlot int32, enabled bool) bool {
	return _SetHudInputCapture(hud, playerSlot, enabled)
}

var _IsHudInputCaptureEnabled = func(hud int32, playerSlot int32) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsHudInputCaptureEnabled(__hud, __playerSlot))
	return __retVal
}

// IsHudInputCaptureEnabled 
//  @brief Get if this hud is capturing input for a player. (CustomHudLayout.IsInputCaptureEnabled)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot.
//
//  @return False if input isn't captured, the handle isn't a hud, CS Script isn't ready or the call failed.
func IsHudInputCaptureEnabled(hud int32, playerSlot int32) bool {
	return _IsHudInputCaptureEnabled(hud, playerSlot)
}

var _ResetHud = func(hud int32) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__retVal = bool(C.ResetHud(__hud))
	return __retVal
}

// ResetHud 
//  @brief Reset to original state for all players. (CustomHudLayout.Reset)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func ResetHud(hud int32) bool {
	return _ResetHud(hud)
}

var _ResetHudForPlayer = func(hud int32, playerSlot int32) bool {
	var __retVal bool
	__hud := C.int32_t(hud)
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.ResetHudForPlayer(__hud, __playerSlot))
	return __retVal
}

// ResetHudForPlayer 
//  @brief Reset a single player's overrides to their original state. (CustomHudLayout.ResetForPlayer)
//
//  @param hud: Hud handle from CreateCustomHud or FindCustomHud.
//  @param playerSlot: Player slot whose overrides are reset.
//
//  @return False if the handle isn't a hud, CS Script isn't ready or the call failed.
func ResetHudForPlayer(hud int32, playerSlot int32) bool {
	return _ResetHudForPlayer(hud, playerSlot)
}

