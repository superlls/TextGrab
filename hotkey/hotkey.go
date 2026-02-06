package hotkey

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Carbon -framework Cocoa

#import <Carbon/Carbon.h>
#import <Cocoa/Cocoa.h>

// Forward declaration
extern void goHotkeyCallback();

// Event handler for hotkey
static OSStatus hotkeyHandler(EventHandlerCallRef nextHandler, EventRef theEvent, void *userData) {
    goHotkeyCallback();
    return noErr;
}

// Register hotkey: Command+Shift+W
static int registerHotkey() {
    EventHotKeyRef hotKeyRef;
    EventHotKeyID hotKeyID;
    EventTypeSpec eventType;

    eventType.eventClass = kEventClassKeyboard;
    eventType.eventKind = kEventHotKeyPressed;

    InstallApplicationEventHandler(&hotkeyHandler, 1, &eventType, NULL, NULL);

    hotKeyID.signature = 'htk1';
    hotKeyID.id = 1;

    // Register Command+Shift+W (keycode 13 = W)
    OSStatus status = RegisterEventHotKey(
        13,                                    // keycode for 'W'
        cmdKey + shiftKey,                     // Command + Shift
        hotKeyID,
        GetApplicationEventTarget(),
        0,
        &hotKeyRef
    );

    return (int)status;
}
*/
import "C"
import (
	"fmt"
	"sync"
)

var (
	callback     func()
	callbackLock sync.Mutex
)

//export goHotkeyCallback
func goHotkeyCallback() {
	callbackLock.Lock()
	cb := callback
	callbackLock.Unlock()

	if cb != nil {
		go cb()
	}
}

// Register registers a global hotkey (Command+Shift+W) with the given callback
func Register(cb func()) error {
	callbackLock.Lock()
	callback = cb
	callbackLock.Unlock()

	status := C.registerHotkey()
	if status != 0 {
		return fmt.Errorf("failed to register hotkey, status: %d", status)
	}

	return nil
}
