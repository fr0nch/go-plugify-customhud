#include "shared.h"

PLUGIFY_EXPORT bool (*__customhud_IsCsScriptReady)() = NULL;


PLUGIFY_EXPORT int32_t (*__customhud_CreateCustomHud)(String*, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__customhud_FindCustomHud)(String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_RemoveCustomHud)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__customhud_HideCustomHudFromOtherPlayers)(int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudHasClass)(int32_t, String*, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__customhud_ResetHudHasClass)(int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudDialogVariable)(int32_t, String*, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudHasClassForPlayer)(int32_t, int32_t, String*, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__customhud_ResetHudHasClassForPlayer)(int32_t, int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_BHasClass)(int32_t, int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_ToggleClass)(int32_t, int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudDialogVariableForPlayer)(int32_t, int32_t, String*, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_ResetHudDialogVariableForPlayer)(int32_t, int32_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudInputCapture)(int32_t, int32_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__customhud_IsHudInputCaptureEnabled)(int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__customhud_ResetHud)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__customhud_ResetHudForPlayer)(int32_t, int32_t) = NULL;


