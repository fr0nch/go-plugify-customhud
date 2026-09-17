#pragma once

#include "shared.h"

extern bool (*__customhud_IsCsScriptReady)();

static bool IsCsScriptReady() {
	return __customhud_IsCsScriptReady();
}

extern bool (*__customhud_CreateCustomHud)(String*, String*);

static bool CreateCustomHud(String* name, String* layoutResource) {
	return __customhud_CreateCustomHud(name, layoutResource);
}

extern bool (*__customhud_RemoveCustomHud)(String*);

static bool RemoveCustomHud(String* name) {
	return __customhud_RemoveCustomHud(name);
}

extern bool (*__customhud_HideCustomHudFromOtherPlayers)(String*, int32_t);

static bool HideCustomHudFromOtherPlayers(String* name, int32_t playerSlot) {
	return __customhud_HideCustomHudFromOtherPlayers(name, playerSlot);
}

extern bool (*__customhud_SetHudHasClass)(String*, String*, String*, bool);

static bool SetHudHasClass(String* name, String* panelId, String* className, bool hasClass) {
	return __customhud_SetHudHasClass(name, panelId, className, hasClass);
}

extern bool (*__customhud_SetHudDialogVariable)(String*, String*, String*, String*);

static bool SetHudDialogVariable(String* name, String* panelId, String* variableName, String* value) {
	return __customhud_SetHudDialogVariable(name, panelId, variableName, value);
}

extern bool (*__customhud_SetHudHasClassForPlayer)(String*, int32_t, String*, String*, bool);

static bool SetHudHasClassForPlayer(String* name, int32_t playerSlot, String* panelId, String* className, bool hasClass) {
	return __customhud_SetHudHasClassForPlayer(name, playerSlot, panelId, className, hasClass);
}

extern bool (*__customhud_SetHudDialogVariableForPlayer)(String*, int32_t, String*, String*, String*);

static bool SetHudDialogVariableForPlayer(String* name, int32_t playerSlot, String* panelId, String* variableName, String* value) {
	return __customhud_SetHudDialogVariableForPlayer(name, playerSlot, panelId, variableName, value);
}

extern bool (*__customhud_SetHudInputCapture)(String*, int32_t, bool);

static bool SetHudInputCapture(String* name, int32_t playerSlot, bool enabled) {
	return __customhud_SetHudInputCapture(name, playerSlot, enabled);
}

extern bool (*__customhud_IsHudInputCaptureEnabled)(String*, int32_t);

static bool IsHudInputCaptureEnabled(String* name, int32_t playerSlot) {
	return __customhud_IsHudInputCaptureEnabled(name, playerSlot);
}

