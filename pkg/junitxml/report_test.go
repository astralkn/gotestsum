package junitxml

//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}
//
//func TestWrite(t *testing.T) {
//	out := new(bytes.Buffer)
//	exec := createExecution(t)
//
//	defer env.Patch(t, "GOVERSION", "go7.7.7")()
//	_, err := Write(out, exec, Config{})
//	assert.NilError(t, err)
//	golden.Assert(t, out.String(), "junitxml-report.golden")
//}
//
//func createExecution(t *testing.T) *testjson.Execution {
//	exec, err := testjson.ScanTestOutput(testjson.ScanConfig{
//		Stdout:  readTestData(t, "out"),
//		Stderr:  readTestData(t, "err"),
//		Handler: &noopHandler{},
//	})
//	assert.NilError(t, err)
//	return exec
//}
//
//func readTestData(t *testing.T, stream string) io.Reader {
//	wd, _ := os.Getwd()
//	raw, err := ioutil.ReadFile(wd + "/go-test-json." + stream)
//	assert.NilError(t, err)
//	return bytes.NewReader(raw)
//}
//
//type noopHandler struct{}
//
//func (s *noopHandler) Event(testjson.TestEvent, *testjson.Execution) error {
//	return nil
//}
//
//func (s *noopHandler) Err(string) error {
//	return nil
//}
//
//func TestGoVersion(t *testing.T) {
//	t.Run("unknown", func(t *testing.T) {
//		defer env.Patch(t, "PATH", "/bogus")()
//		assert.Equal(t, goVersion(), "unknown")
//	})
//
//	t.Run("current version", func(t *testing.T) {
//		expected := fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
//		assert.Equal(t, goVersion(), expected)
//	})
//}
