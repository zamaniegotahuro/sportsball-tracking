
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            time.Sleep(time.Second)
            fmt.Println("Hello, World!")
        }
    }()
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            time.Sleep(time.Second)
            fmt.Println("Goodbye, World!")
        }
    }()
    wg.Wait()
}