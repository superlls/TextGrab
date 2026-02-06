package main

import (
	"log"
	"strings"

	"textgrab/clipboard"
	"textgrab/hotkey"
	"textgrab/ocr"
	"textgrab/screenshot"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// Set title and tooltip
	systray.SetTitle("TextGrab")
	systray.SetTooltip("TextGrab - 屏幕截图 OCR 工具\n快捷键: ⌘⇧W")

	// Add menu items
	mCapture := systray.AddMenuItem("截图识别 (⌘⇧W)", "截取屏幕区域并进行 OCR 识别")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("关于", "关于 TextGrab")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出 TextGrab")

	// Register global hotkey
	err := hotkey.Register(func() {
		performOCR()
	})
	if err != nil {
		log.Printf("注册快捷键失败: %v", err)
		beeep.Alert("TextGrab", "快捷键注册失败，请检查权限设置", "")
	}

	// Handle menu clicks
	go func() {
		for {
			select {
			case <-mCapture.ClickedCh:
				performOCR()
			case <-mAbout.ClickedCh:
				beeep.Notify("TextGrab", "屏幕截图 OCR 工具\n快捷键: ⌘⇧W (Command+Shift+W)\n\n支持简体中文识别", "")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleanup
}

func performOCR() {
	// Capture screenshot
	screenshotPath, err := screenshot.CaptureScreen()
	if err != nil {
		if err.Error() == "screenshot cancelled by user" {
			// User cancelled, do nothing
			return
		}
		log.Printf("截图失败: %v", err)
		beeep.Alert("TextGrab", "截图失败: "+err.Error(), "")
		return
	}

	// Perform OCR
	text, err := ocr.RecognizeText(screenshotPath)

	// Clean up screenshot file
	defer screenshot.Cleanup(screenshotPath)

	if err != nil {
		log.Printf("OCR 识别失败: %v", err)
		beeep.Alert("TextGrab", "OCR 识别失败: "+err.Error(), "")
		return
	}

	// Trim whitespace
	text = strings.TrimSpace(text)

	if text == "" {
		beeep.Alert("TextGrab", "未识别到任何文本", "")
		return
	}

	// Copy to clipboard
	if err := clipboard.WriteText(text); err != nil {
		log.Printf("复制到剪切板失败: %v", err)
		beeep.Alert("TextGrab", "复制到剪切板失败: "+err.Error(), "")
		return
	}

	// Send success notification
	beeep.Notify("TextGrab", "识别完成，文本已复制到剪切板", "")
}
