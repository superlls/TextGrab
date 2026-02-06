# 📸 TextGrab

一个基于 Go 语言开发的 macOS 菜单栏 OCR 工具，实现"截图 → OCR 识别 → 自动复制到剪切板"的完整闭环。

## ✨ 功能特性

- 🎯 **菜单栏常驻** - 程序常驻系统菜单栏，随时可用
- ⌨️ **全局快捷键** - 使用 `⌘⇧W` 快速触发截图识别
- 🖱️ **交互式截图** - 鼠标框选需要识别的屏幕区域
- 🌏 **多语言识别** - 自动识别中文、日文、韩文、英文等多种语言
- 🎯 **高精度 OCR** - 使用 macOS 原生 Vision 框架
- 📋 **自动复制** - 识别结果自动复制到系统剪切板
- 🔄 **可重复使用** - 识别完成后可立即进行下一次识别

## 🚀 快速开始

### 启动应用

```bash
./textgrab
```

应用会在菜单栏显示 "TextGrab"，不会打开窗口。

### 使用方法

1. 按下快捷键 `⌘⇧W` (Command + Shift + W)
2. 用鼠标框选要识别的屏幕区域
3. 松开鼠标后自动识别
4. 识别完成后文本自动复制到剪切板
5. 按 `⌘V` 即可粘贴

### 菜单选项

- **截图识别 (⌘⇧W)** - 触发截图和 OCR 识别
- **关于** - 查看应用信息
- **退出** - 退出应用

## 🔧 编译说明

### 前置要求

- macOS 10.15+ (Catalina 或更高版本)
- Go 1.21 或更高版本
- Xcode Command Line Tools

### 编译步骤

```bash
# 安装依赖
go mod download

# 编译项目
chmod +x build.sh
./build.sh
```

或者手动编译:

```bash
CGO_ENABLED=1 go build -o textgrab .
```

## 🔐 权限设置

首次运行需要授予以下权限:

### 1️⃣ 屏幕录制权限 (必需)

- 打开 **系统偏好设置** → **安全性与隐私** → **隐私** → **屏幕录制**
- 勾选 **Terminal** 或 **iTerm**（取决于你使用的终端）
- 修改权限后需要重启应用

### 2️⃣ 辅助功能权限 (可选，用于全局快捷键)

- 打开 **系统偏好设置** → **安全性与隐私** → **隐私** → **辅助功能**
- 勾选 **Terminal** 或 **iTerm**
- 修改权限后需要重启应用

## 💡 使用技巧

- ✅ 按 `⌘⇧W` 后，用鼠标框选要识别的区域
- ❌ 按 `ESC` 可以取消截图
- 📋 识别完成后立即使用 `⌘V` 粘贴文本
- 🔄 程序会一直运行在后台，可以重复使用
- 🚪 点击菜单栏图标选择"退出"可以关闭程序

## 🌍 语言支持

TextGrab 使用 macOS Vision 框架的自动语言检测功能，可以识别系统支持的所有语言，包括但不限于:

- 🇨🇳 简体中文 / 繁体中文
- 🇯🇵 日语
- 🇰🇷 韩语
- 🇺🇸 英语
- 🇫🇷 法语
- 🇩🇪 德语
- 🇪🇸 西班牙语
- 🇮🇹 意大利语
- 🇵🇹 葡萄牙语

> **注意**: 实际支持的语言取决于你的 macOS 版本。较新的系统版本支持更多语言。

## 📁 项目结构

```
TextGrab/
├── main.go                 # 主程序入口（菜单栏应用）
├── go.mod                  # Go 模块依赖
├── build.sh                # 编译脚本
├── ocr/
│   └── ocr.go             # OCR 模块 (CGO + Vision)
├── screenshot/
│   └── screenshot.go      # 截图模块
├── clipboard/
│   └── clipboard.go       # 剪切板模块
└── hotkey/
    └── hotkey.go          # 全局快捷键模块 (CGO + Carbon)
```

## 🛠️ 技术栈

- **Go 1.21+** - 主要开发语言
- **CGO** - 用于调用 Objective-C 和 macOS 框架
- **Vision Framework** - macOS 原生 OCR 引擎
- **Carbon Framework** - 全局快捷键监听
- **Systray** - 菜单栏应用框架

### 依赖库

- `github.com/getlantern/systray` - 菜单栏应用框架
- `github.com/atotto/clipboard` - 剪切板操作
- `github.com/gen2brain/beeep` - 系统通知

## ❓ 常见问题

### 快捷键不工作

- 确保已授予辅助功能权限
- 检查是否有其他应用占用了 `⌘⇧W` 快捷键
- 尝试重启应用

### 截图失败: "operation not permitted"

- 检查是否授予了屏幕录制权限
- 修改权限后需要重启应用

### 编译错误: "framework not found"

确保已安装 Xcode Command Line Tools:

```bash
xcode-select --install
```

### OCR 识别不准确

- 确保截图清晰，文字大小适中
- 避免截取过小或模糊的文字
- 尽量选择纯文本区域，避免复杂背景

### 如何开机自启动？

1. 系统偏好设置 → 用户与群组 → 登录项
2. 点击 "+" 添加 textgrab 应用
3. 勾选"隐藏"选项

## 🔨 开发说明

### 修改快捷键

编辑 `hotkey/hotkey.go`，修改快捷键组合:

```c
// 当前: Command+Shift+W (keycode 13)
RegisterEventHotKey(13, cmdKey + shiftKey, ...)

// 改为 Command+Shift+X (keycode 7)
RegisterEventHotKey(7, cmdKey + shiftKey, ...)
```

常用按键代码:
- W: 13
- X: 7
- C: 8
- V: 9
- A: 0
- S: 1

### 自定义 OCR 参数

编辑 `ocr/ocr.go` 中的 `performOCR` 函数:

- 修改 `recognitionLevel` 调整识别精度
- 修改 `automaticallyDetectsLanguage` 启用/禁用自动语言检测
- 修改 `usesLanguageCorrection` 启用/禁用语言校正
- 修改置信度阈值 `confidence >= 0.3f` 调整识别敏感度
