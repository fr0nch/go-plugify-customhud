#include "shared.h"

PLUGIFY_EXPORT bool (*__customhud_IsCsScriptReady)() = NULL;


PLUGIFY_EXPORT bool (*__customhud_CreateCustomHud)(String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_RemoveCustomHud)(String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_HideCustomHudFromOtherPlayers)(String*, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudHasClass)(String*, String*, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudDialogVariable)(String*, String*, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudHasClassForPlayer)(String*, int32_t, String*, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudDialogVariableForPlayer)(String*, int32_t, String*, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__customhud_SetHudInputCapture)(String*, int32_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__customhud_IsHudInputCaptureEnabled)(String*, int32_t) = NULL;


