package randtest

type TestingTask interface {
	Run() error
	SaveData() error
}
