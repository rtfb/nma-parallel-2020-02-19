func syncThingsC(ctx context.Context, cli upstream.Client, input []*upstream.Thing) ([]*thing, error) {
        var wg sync.WaitGroup
        const nWorkers = 4
        dataChan := make(chan *upstream.Thing, len(input))
        for _, i := range input {
                dataChan <- i
        }
        close(dataChan)
        things := make(chan []*thing, nWorkers)
        errc := make(chan error, nWorkers)
        wg.Add(nWorkers)
        for i := 0; i < nWorkers; i++ {
                go func() {
                        defer wg.Done()
                        var batch []*thing
                        for i := range dataChan {
                                batchlet, err := getOneBatch(ctx, cli, i)
                                if err != nil {
                                        errc <- err
                                        return
                                }
                                batch = append(batch, batchlet...)
                        }
                        things <- batch
                }()
        }
        wg.Wait()
        close(things)
        close(errc)
        for err := range errc {
                if err != nil {
                        return nil, err
                }
        }
        var newThings []*thing
        for chunk := range scheds {
                newThings = append(newThings, chunk...)
        }
        return newThings, nil
}
