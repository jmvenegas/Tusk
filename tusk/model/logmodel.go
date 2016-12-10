package model

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Level int

const (
	INFO Level = iota
	WARN
	FATAL
)

var LogLevelStrings = map[Level]string{
	0: "INFO",
	1: "WARN",
	2: "FATA",
}

type LogEntry struct {
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
	LogLevel Level     `json:"loglevel"`
}

func (le *LogEntry) ToString() string {
	return fmt.Sprintf("[%s]%s : %s",
		LogLevelStrings[le.LogLevel],
		le.Date.Format(time.RFC850),
		le.Message)
}

func (le *LogEntry) ToJson() []byte {
	jString, _ := json.Marshal(le)
	return jString
}

type LogModel struct {
	FileName     string
	FP           *os.File
	CurrentLevel Level
}

func NewLogModel(file string) *LogModel {
	lm := new(LogModel)
	lm.FileName = file
	lm.CurrentLevel = INFO
	return lm
}

func (l *LogModel) Init() {
	f, err := os.OpenFile(l.FileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Problem opening file for logging")
	}
	l.FP = f
}

// TODO - set a better way to close file, some fancy defer somewhere?
func (l *LogModel) Close() {
	l.FP.Close()
}

func (l *LogModel) SetLogLevel(level Level) {
	l.CurrentLevel = level
}

func (l *LogModel) Log(level Level, message string) {
	l.write(l.buildEntry(level, message))
}

func (l *LogModel) buildEntry(level Level, message string) *LogEntry {
	le := LogEntry{Message: message,
		Date:     time.Now(),
		LogLevel: level}
	return &le
}

func (l *LogModel) write(le *LogEntry) {
	if le.LogLevel <= l.CurrentLevel {
		l.FP.WriteString(le.ToString())
	}
}
