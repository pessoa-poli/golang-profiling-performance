# Start a server at port 8080 to read the file cpu.pprof created on the profiling stage
read-cpu-pprof:
	go tool pprof -http=:8080 cpu.pprof

read-trace-profile: # Question mark shows you the hotkeys to navigate this visualization.
	go tool trace trace.out