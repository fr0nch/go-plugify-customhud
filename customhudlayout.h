#pragma once

#include "shared.h"

extern bool (*__customhud_IsCsScriptReady)();

static bool IsCsScriptReady() {
	return __customhud_IsCsScriptReady();
}

extern int32_t (*__customhud_CreateCustomHud)(String*, String*);

static int32_t CreateCustomHud(String* name, String* layoutResource) {
	return __customhud_CreateCustomHud(name, layoutResource);
}

extern int32_t (*__customhud_FindCustomHud)(String*);

static int32_t FindCustomHud(String* name) {
	return __customhud_FindCustomHud(name);
}

extern bool (*__customhud_RemoveCustomHud)(int32_t);

static bool RemoveCustomHud(int32_t hud) {
	return __customhud_RemoveCustomHud(hud);
}

extern bool (*__customhud_HideCustomHudFromOtherPlayers)(int32_t, int32_t);

static bool HideCustomHudFromOtherPlayers(int32_t hud, int32_t playerSlot) {
	return __customhud_HideCustomHudFromOtherPlayers(hud, playerSlot);
}

extern bool (*__customhud_SetHudHasClass)(int32_t, String*, String*, bool);

static bool SetHudHasClass(int32_t hud, String* panelId, String* className, bool hasClass) {
	return __customhud_SetHudHasClass(hud, panelId, className, hasClass);
}

extern bool (*__customhud_ResetHudHasClass)(int32_t, String*, String*);

static bool ResetHudHasClass(int32_t hud, String* panelId, String* className) {
	return __customhud_ResetHudHasClass(hud, panelId, className);
}

extern bool (*__customhud_SetHudDialogVariable)(int32_t, String*, String*, String*);

static bool SetHudDialogVariable(int32_t hud, String* panelId, String* variableName, String* value) {
	return __customhud_SetHudDialogVariable(hud, panelId, variableName, value);
}

extern bool (*__customhud_SetHudHasClassForPlayer)(int32_t, int32_t, String*, String*, bool);

static bool SetHudHasClassForPlayer(int32_t hud, int32_t playerSlot, String* panelId, String* className, bool hasClass) {
	return __customhud_SetHudHasClassForPlayer(hud, playerSlot, panelId, className, hasClass);
}

extern bool (*__customhud_ResetHudHasClassForPlayer)(int32_t, int32_t, String*, String*);

static bool ResetHudHasClassForPlayer(int32_t hud, int32_t playerSlot, String* panelId, String* className) {
	return __customhud_ResetHudHasClassForPlayer(hud, playerSlot, panelId, className);
}

extern bool (*__customhud_BHasClass)(int32_t, int32_t, String*, String*);

static bool BHasClass(int32_t hud, int32_t playerSlot, String* panelId, String* className) {
	return __customhud_BHasClass(hud, playerSlot, panelId, className);
}

extern bool (*__customhud_ToggleClass)(int32_t, int32_t, String*, String*);

static bool ToggleClass(int32_t hud, int32_t playerSlot, String* panelId, String* className) {
	return __customhud_ToggleClass(hud, playerSlot, panelId, className);
}

extern bool (*__customhud_SetHudDialogVariableForPlayer)(int32_t, int32_t, String*, String*, String*);

static bool SetHudDialogVariableForPlayer(int32_t hud, int32_t playerSlot, String* panelId, String* variableName, String* value) {
	return __customhud_SetHudDialogVariableForPlayer(hud, playerSlot, panelId, variableName, value);
}

extern bool (*__customhud_ResetHudDialogVariableForPlayer)(int32_t, int32_t, String*, String*);

static bool ResetHudDialogVariableForPlayer(int32_t hud, int32_t playerSlot, String* panelId, String* variableName) {
	return __customhud_ResetHudDialogVariableForPlayer(hud, playerSlot, panelId, variableName);
}

extern bool (*__customhud_SetHudInputCapture)(int32_t, int32_t, bool);

static bool SetHudInputCapture(int32_t hud, int32_t playerSlot, bool enabled) {
	return __customhud_SetHudInputCapture(hud, playerSlot, enabled);
}

extern bool (*__customhud_IsHudInputCaptureEnabled)(int32_t, int32_t);

static bool IsHudInputCaptureEnabled(int32_t hud, int32_t playerSlot) {
	return __customhud_IsHudInputCaptureEnabled(hud, playerSlot);
}

extern bool (*__customhud_ResetHud)(int32_t);

static bool ResetHud(int32_t hud) {
	return __customhud_ResetHud(hud);
}

extern bool (*__customhud_ResetHudForPlayer)(int32_t, int32_t);

static bool ResetHudForPlayer(int32_t hud, int32_t playerSlot) {
	return __customhud_ResetHudForPlayer(hud, playerSlot);
}

