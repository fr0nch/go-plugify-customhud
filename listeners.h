#pragma once

#include "shared.h"

extern void (*__customhud_OnCsScriptReady_Register)(void*);

static void OnCsScriptReady_Register(void* callback) {
	__customhud_OnCsScriptReady_Register(callback);
}

extern void (*__customhud_OnCsScriptReady_Unregister)(void*);

static void OnCsScriptReady_Unregister(void* callback) {
	__customhud_OnCsScriptReady_Unregister(callback);
}

extern void (*__customhud_OnHudClicked_Register)(void*);

static void OnHudClicked_Register(void* callback) {
	__customhud_OnHudClicked_Register(callback);
}

extern void (*__customhud_OnHudClicked_Unregister)(void*);

static void OnHudClicked_Unregister(void* callback) {
	__customhud_OnHudClicked_Unregister(callback);
}

