package adapter

// AdvancedMediaPlayer 高级媒体播放器接口
// 高级媒体播放器接口作为被适配者，我们的目标是让目标接口具备被适配者的功能
type AdvancedMediaPlayer interface {
	Play(mediaType, data string) error
}

// Mp4Player Mp4播放器
type Mp4Player struct {
}

func (m *Mp4Player) Play(mediaType, data string) error {
	//TODO implement me
	panic("implement me")
}
