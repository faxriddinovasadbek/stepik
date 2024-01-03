var kw  sync.WaitGroup

	for i := 0; i < 10; i++{
		kw.Add(1)
		go func(){
			defer kw.Done()
			work()
		}()
	}

	kw.Wait()
