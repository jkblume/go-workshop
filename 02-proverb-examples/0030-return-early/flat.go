func main() {
	if !something.OK() {  // flipped
        return errors.New("something not ok")
	}
	something.Lock()
	defer something.Unlock()
	err := something.Do()
	if err != nil {       // flipped
		return err
	}
	stop := StartTimer()
	defer stop()
	log.Println("working...")
	doWork(something)
	<-something.Done() // wait for it
	log.Println("finished")
	return nil
}
