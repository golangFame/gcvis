package trace

type Scvgtrace struct {
	ElapsedTime float64 // in seconds
	Inuse       int64
	Idle        int64
	Sys         int64
	Released    int64
	Consumed    int64
}

type Gctrace struct {
	ElapsedTime  float64 // in seconds
	NumGC        int64
	Nproc        int64
	t1           int64
	t2           int64
	t3           int64
	t4           int64
	Heap0        int64 // heap size before, in megabytes
	Heap1        int64 // heap size after, in megabytes
	Obj          int64
	NMalloc      int64
	NFree        int64
	NSpan        int64
	NGoRoutines  int64
	NBGSweep     int64
	NPauseSweep  int64
	NHandoff     int64
	NHandoffCnt  int64
	NSteal       int64
	NStealCnt    int64
	NProcYield   int64
	NOsYield     int64
	NSleep       int64
	STWSclock    float64
	MASclock     float64
	STWMclock    float64
	STWScpu      float64
	MASAssistcpu float64
	MASBGcpu     float64
	MASIdlecpu   float64
	STWMcpu      float64
}
