func main() {
	if something.OK() {
        something.Lock()
        defer something.Unlock()
        err := something.Do()
        if err == nil {
			stop := StartTimer()
			defer stop()
			log.Println("working...")
			doWork(something)
			<-something.Done() // wait for it
			log.Println("finished")
			return nil
        } else {
			return err
        }
	} else {
		return errors.New("something not ok")
	}
}
