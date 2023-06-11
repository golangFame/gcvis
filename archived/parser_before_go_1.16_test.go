package main

/*func TestParserWithMatchingInputGo15(t *testing.T) {
	line := "gc 88 @3.243s 9%: 0.040+16+1.0+5.9+0.34 ms clock, 0.16+16+0+18/5.7/11+1.3 ms cpu, 32->33->19 MB, 33 MB goal, 4 P"

	runParserWith(line)

	expectedGCTrace := &Gctrace{
		Heap1:       33,
		ElapsedTime: 3.243,
	}

	select {
	case gctrace := <-parser.GcChan:
		if !reflect.DeepEqual(gctrace, expectedGCTrace) {
			t.Errorf("Expected gctrace to equal %+v. Got %+v instead.", expectedGCTrace, gctrace)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}

func TestParserWithMatchingInputGo14(t *testing.T) {
	line := "gc76(1): 2+1+1390+1 us, 1 -> 3 MB, 16397 (1015746-999349) objects, 1436/1/0 sweeps, 0(0) handoff, 0(0) steal, 0/0/0 yields"

	runParserWith(line)

	expectedGCTrace := &Gctrace{
		Heap1: 3,
	}

	select {
	case gctrace := <-parser.GcChan:
		if !reflect.DeepEqual(gctrace, expectedGCTrace) {
			t.Errorf("Expected gctrace to equal %+v. Got %+v instead.", expectedGCTrace, gctrace)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}*/

/*func TestParserGoRoutinesInputGo14(t *testing.T) {
	line := "gc76(1): 2+1+1390+1 us, 1 -> 3 MB, 16397 (1015746-999349) objects, 12 goroutines, 1436/1/0 sweeps, 0(0) handoff, 0(0) steal, 0/0/0 yields"

	runParserWith(line)

	expectedGCTrace := &Gctrace{
		Heap1: 3,
	}

	select {
	case gctrace := <-parser.GcChan:
		if !reflect.DeepEqual(gctrace, expectedGCTrace) {
			t.Errorf("Expected gctrace to equal %+v. Got %+v instead.", expectedGCTrace, gctrace)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}*/
