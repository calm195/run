package core

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Cutter
//
//	@Description: 实现 io.Writer 接口。
//	用于日志切割, strings.Join([]string{director,layout, formats..., level+".log"}, os.PathSeparator)
type Cutter struct {
	level        string        // 日志级别(debug, info, warn, error, panic, fatal)
	layout       string        // 时间格式 2006-01-02 15:04:05
	formats      []string      // 自定义参数([]string{Director,"2006-01-02", level+".log"}
	director     string        // 日志文件夹
	retentionDay int           // 日志保留天数
	file         *os.File      // 文件句柄
	mutex        *sync.RWMutex // 读写锁
}

type CutterOption func(*Cutter)

// CutterWithLayout
//
//	@Description: 时间格式函数包，返回函数func(c *Cutter)将参数设置到Cutter中
//	@param layout // 例如: "2006-01-02 15:04:05" 或 "2006-01-02"
//	@return CutterOption
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

// CutterWithFormats
//
//	@Description: 格式化参数函数包，返回函数func(c *Cutter)将参数设置到Cutter中
//	@param format
//	@return CutterOption
func CutterWithFormats(format ...string) CutterOption {
	return func(c *Cutter) {
		if len(format) > 0 {
			c.formats = format
		}
	}
}

// NewCutter
//
//	@Description: 根据给定的参数创建一个新的 Cutter 实例。
//	@param director  日志文件夹路径
//	@param level   日志级别(debug, info, warn, error, debug, panic, fatal)
//	@param retentionDay 日志保留天数，表示删除多少天前的目录，小于等于零的值默认忽略不再处理
//	@param options 可选参数，支持设置时间格式和自定义参数
//	@return *Cutter 返回一个新的 Cutter 实例
func NewCutter(director string, level string, retentionDay int, options ...CutterOption) *Cutter {
	rotate := &Cutter{
		level:        level,
		director:     director,
		retentionDay: retentionDay,
		mutex:        new(sync.RWMutex),
	}
	for i := 0; i < len(options); i++ {
		options[i](rotate)
	}
	return rotate
}

// Write
//
//	@Description: 自动创建目录并写入日志，并扫描是否需要删除过期日志。线程安全，确保在多线程环境下的安全写入操作。
//	@receiver c *Cutter
//	@param bytes []byte  日志内容
//	@return n int   写入的字节数
//	@return err error 如果发生错误则返回错误
func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()
	length := len(c.formats)
	values := make([]string, 0, 3+length)
	values = append(values, c.director)
	if c.layout != "" {
		values = append(values, time.Now().Format(c.layout))
	}
	for i := 0; i < length; i++ {
		values = append(values, c.formats[i])
	}
	values = append(values, c.level+".log")
	filename := filepath.Join(values...)
	director := filepath.Dir(filename)
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err
	}
	err = removeNDaysFolders(c.director, c.retentionDay)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}

// Sync
//
//	@Description: 同步文件内容到磁盘，线程安全，确保在多线程环境下的安全写入操作。
//	@receiver c *Cutter
//	@return error 如果发生错误则返回错误
func (c *Cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.file != nil {
		return c.file.Sync()
	}
	return nil
}

// removeNDaysFolders
//
//	@Description: 日志目录文件清理，小于等于零的值默认忽略不再处理
//	@param dir 日志目录
//	@param days 天数，表示删除多少天前的目录
//	@return error 如果发生错误则返回错误
func removeNDaysFolders(dir string, days int) error {
	if days <= 0 {
		return nil
	}
	cutoff := time.Now().AddDate(0, 0, -days)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.ModTime().Before(cutoff) && path != dir {
			err = os.RemoveAll(path)
		}
		return err
	})
}
