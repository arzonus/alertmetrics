package config

type Config struct {
	Database struct {
		Connection string `yaml:"connection"`
	} `yaml:"db"`

	Metrics []Metric `yaml:"metrics"`
	Period  uint     `yaml:"period"`

	Notifier Notifier `yaml:"notifiers"`
}

type Metric struct {
	Name       string `yaml:"name"`
	LowerBound uint   `yaml:"lowerBound"`
	UpperBound uint   `yaml:"upperBound"`
}

type Notifier struct {
	LogNotifier struct {
		Enable bool `yaml:"enable"`
	} `yaml:"log"`
}
