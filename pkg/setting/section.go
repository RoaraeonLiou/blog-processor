package setting

type BasicSettingS struct {
	BlogDir                string
	ImageDir               string
	TemplateFile           string
	OutputDir              string
	HttpBasePath           string
	DateLayout             string
	CommonHeaderFileName   string
	CommonHeaderFileExt    string
	CommonHeaderFileFormat string
}

type LogStrategySettingS struct {
	LogToFile bool
	LogFile   string
}

type DBSettingS struct {
	DBFile string
}

type ArchivesSettingS struct {
	Require bool
	Title   string
	Layout  string
	Url     string
	Summary string
	Type    string
}

type SearchSettingS struct {
	Require     bool
	Title       string
	Layout      string
	Summary     string
	Placeholder string
	Type        string
}

type GlobalHeaderSettingS struct {
	Author string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
