package setting

type BasicSettingS struct {
	BlogDir      string
	ImageDir     string
	TemplateFile string
	OutputDir    string
}

type LogStrategySettingS struct {
	LogToFile bool
	LogFile   string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
