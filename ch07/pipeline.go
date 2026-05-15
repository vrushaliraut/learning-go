package main

type StageContext struct {
	StageName    string
	PipeLineName string
	PipeLineId   string
	RepoPath     string
}
type StageResult struct {
	Success bool
	Error   string
	Output  string
}

func LinterStage(ctx StageContext) StageResult {
	// Linting logic here
	return StageResult{Success: true, Output: "Linting passed"}
}

func TestStage(ctx StageContext) StageResult {
	// Testing logic here
	return StageResult{Success: true, Output: "All tests passed"}
}

func BuildStage(ctx StageContext) StageResult {
	// Building logic here
	return StageResult{Success: true, Output: "Build successful"}
}

func runPipeLine(ctx StageContext, stages []func(ctx2 StageContext) StageResult) StageResult {
	for _, stage := range stages {
		if result := stage(ctx); !result.Success {
			return result
		}
	}
	return StageResult{Success: true, Output: "All tests passed"}
}

func main() {
	stages := []func(ctx StageContext) StageResult{LinterStage, TestStage, BuildStage}
	runPipeLine(StageContext{StageName: "CI Pipeline", PipeLineName: "MyPipeline", PipeLineId: "12345", RepoPath: "/path/to/repo"}, stages)
}
