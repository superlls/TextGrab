# TextGrab

一个基于 Go 语言开发的 macOS 菜单栏 OCR 工具，实现"截图 -> OCR 识别 -> 自动复制到剪切板"的完整闭环。

## 功能特性

- **菜单栏常驻**: 程序常驻系统菜单栏，随时可用
- **全局快捷键**: 使用 ⌘⇧W (Command+Shift+W) 快速触发截图识别
- **交互式截图**: 鼠标框选需要识别的屏幕区域
- **高精度 OCR**: 使用 macOS 原生 Vision 框架，专注简体中文识别
- **自动复制**: 识别结果自动复制到系统剪切板
- **系统通知**: 完成后发送 macOS 系统通知
- **可重复使用**: 识别完成后可立即进行下一次识别

## 使用方法

### 快速开始

1. 启动程序后，会在菜单栏显示 "TextGrab" 图标
2. 使用快捷键 **⌘⇧W** 或点击菜单栏图标选择"截图识别"
3. 用鼠标框选要识别的屏幕区域
4. 等待识别完成，文本会自动复制到剪切板
5. 可以继续使用快捷键进行下一次识别

### 菜单选项

- **截图识别 (⌘⇧W)**: 触发截图和 OCR 识别
- **关于**: 查看应用信息
- **退出**: 退出应用

## 技术架构

### 核心技术栈

- **Go 1.21+**: 主要开发语言
- **CGO**: 用于调用 Objective-C 和 macOS 框架
- **Vision Framework**: macOS 原生 OCR 引擎（仅简体中文，高精度）
- **Carbon Framework**: 全局快捷键监听
- **Systray**: 菜单栏应用框架

### 项目结构

```
TextGrab/
├── main.go                 # 主程序入口（菜单栏应用）
├── go.mod                  # Go 模块依赖
├── build.sh                # 编译脚本
├── README.md               # 项目说明
├── ocr/
│   └── ocr.go             # OCR 模块 (CGO + Vision)
├── screenshot/
│   └── screenshot.go      # 截图模块
├── clipboard/
│   └── clipboard.go       # 剪切板模块
└── hotkey/
    └── hotkey.go          # 全局快捷键模块 (CGO + Carbon)
```

## 编译说明

### 前置要求

- macOS 10.15+ (Catalina 或更高版本)
- Go 1.21 或更高版本
- Xcode Command Line Tools

### 编译步骤

1. **安装依赖**

```bash
go mod download
```

2. **编译项目**

```bash
chmod +x build.sh
./build.sh
```

或者手动编译:

```bash
CGO_ENABLED=1 go build -o textgrab .
```

### 链接的 macOS 框架

编译时会自动链接以下框架:

**OCR 模块** (`ocr/ocr.go`):
- `-framework Foundation`: 基础框架
- `-framework Vision`: OCR 识别引擎
- `-framework CoreGraphics`: 图像处理
- `-framework AppKit`: macOS UI 框架

**快捷键模块** (`hotkey/hotkey.go`):
- `-framework Carbon`: 全局快捷键监听
- `-framework Cocoa`: macOS 应用框架

## 运行说明

### 启动应用

```bash
./textgrab
```

应用会在菜单栏显示，不会打开窗口。

### 权限设置

首次运行需要授予以下权限:

1. **屏幕录制权限** (必需)
   - 系统偏好设置 → 安全性与隐私 → 隐私 → 屏幕录制
   - 将 textgrab 或终端添加到允许列表
   - 修改权限后需要重启应用

2. **辅助功能权限** (可选，用于全局快捷键)
   - 系统偏好设置 → 安全性与隐私 → 隐私 → 辅助功能
   - 将 textgrab 或终端添加到允许列表

### 使用技巧

- 按 **⌘⇧W** 后，用鼠标框选要识别的区域
- 按 **ESC** 可以取消截图
- 识别完成后会听到系统通知声音
- 可以立即使用 **⌘V** 粘贴识别的文本
- 程序会一直运行在后台，可以重复使用

## OCR 配置

OCR 模块使用 Vision 框架，配置如下:

- **识别语言**: 仅简体中文 (`zh-Hans`)
- **识别级别**: `VNRequestTextRecognitionLevelAccurate` (高精度模式)
- **语言校正**: 已启用

### 为什么只支持简体中文？

为了获得最佳的中文识别准确度，我们专注于简体中文识别。如果需要识别英文或其他语言，可以修改 `ocr/ocr.go:36`:

```objective-c
// 添加英文支持
request.recognitionLanguages = @[@"zh-Hans", @"en-US"];

// 或其他语言
request.recognitionLanguages = @[@"zh-Hans", @"ja", @"ko"];
```

支持的语言代码:
- `zh-Hans`: 简体中文
- `zh-Hant`: 繁体中文
- `en-US`: 英文
- `ja`: 日语
- `ko`: 韩语
- `fr`: 法语
- `de`: 德语
- `es`: 西班牙语
- `it`: 意大利语
- `pt`: 葡萄牙语

## 依赖库

- `github.com/getlantern/systray`: 菜单栏应用框架
- `github.com/atotto/clipboard`: 剪切板操作
- `github.com/gen2brain/beeep`: 系统通知

## 常见问题

### 1. 快捷键不工作

- 确保已授予辅助功能权限
- 检查是否有其他应用占用了 ⌘⇧W 快捷键
- 尝试重启应用

### 2. 截图失败: "operation not permitted"

需要授予屏幕录制权限:
- 系统偏好设置 → 安全性与隐私 → 隐私 → 屏幕录制
- 添加 textgrab 到允许列表
- 重启应用

### 3. 编译错误: "framework not found"

确保已安装 Xcode Command Line Tools:

```bash
xcode-select --install
```

### 4. OCR 识别不准确

- 确保截图清晰，文字大小适中
- 避免截取过小或模糊的文字
- 尽量选择纯文本区域，避免复杂背景

### 5. 如何退出应用？

点击菜单栏的 TextGrab 图标，选择"退出"。

### 6. 如何开机自启动？

1. 系统偏好设置 → 用户与群组 → 登录项
2. 点击 "+" 添加 textgrab 应用
3. 勾选"隐藏"选项

## 开发说明

### 修改快捷键

编辑 `hotkey/hotkey.go:40`，修改快捷键组合:

```c
// 当前: Command+Shift+W
RegisterEventHotKey(13, cmdKey + shiftKey, ...)

// 改为 Command+Shift+X
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
- 修改 `recognitionLanguages` 调整支持的语言
- 修改 `usesLanguageCorrection` 启用/禁用语言校正

## 许可证

MIT License

## 作者

Created with Claude Code
