#pragma once

#include "shared.h"

extern void (*__panorama_api_OnCsScriptReady_Register)(void*);

static void OnCsScriptReady_Register(void* callback) {
	__panorama_api_OnCsScriptReady_Register(callback);
}

extern void (*__panorama_api_OnCsScriptReady_Unregister)(void*);

static void OnCsScriptReady_Unregister(void* callback) {
	__panorama_api_OnCsScriptReady_Unregister(callback);
}

extern void (*__panorama_api_OnHudClicked_Register)(void*);

static void OnHudClicked_Register(void* callback) {
	__panorama_api_OnHudClicked_Register(callback);
}

extern void (*__panorama_api_OnHudClicked_Unregister)(void*);

static void OnHudClicked_Unregister(void* callback) {
	__panorama_api_OnHudClicked_Unregister(callback);
}

