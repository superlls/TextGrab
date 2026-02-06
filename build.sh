#!/bin/bash

# TextGrab 编译脚本
# 用于编译 macOS 平台的 OCR 截图工具

set -e

echo "开始编译 TextGrab..."

# 设置 CGO 环境变量
export CGO_ENABLED=1

# 编译项目
# 注意：macOS 框架会通过 ocr.go 中的 #cgo LDFLAGS 自动链接
go build -o textgrab .

echo "✓ 编译成功！"
echo ""
echo "可执行文件: ./textgrab"
echo ""
echo "运行方式:"
echo "  ./textgrab"
echo ""
echo "注意事项:"
echo "  1. 首次运行可能需要授予屏幕录制权限"
echo "  2. 系统偏好设置 -> 安全性与隐私 -> 隐私 -> 屏幕录制"
echo "  3. 将终端或 iTerm 添加到允许列表中"
