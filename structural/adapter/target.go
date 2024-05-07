package adapter

import (
	"errors"
	"fmt"
)

// MediaPlayer 媒体播放器
// 媒体播放器作为目标接口，也就是说我们的目标是让媒体播放器具备额外的功能
type MediaPlayer interface {
	Play(mediaType, data string) error
}

// AudioPlayer 音频播放器
// 音频播放器实现了媒体播放器接口，但当前支持持播放mp3类型
type AudioPlayer struct {
}

func (p *AudioPlayer) Play(mediaType, data string) error {
	if mediaType != "mp3" {
		return errors.New("unsupported media type")
	}
	fmt.Println("AudioPlayer play mp3:" + data)
	return nil
}
