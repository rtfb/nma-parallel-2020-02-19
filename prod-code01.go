func (c *cache) periodicSync() {
        snapshot := c.getLatestSnapshot() // HLa
        c.Lock() // HLb
        defer c.Unlock() // HLb
        if snapshot != nil {
                c.snapshot = snapshot // HLc
                c.snapshotTime.Store(time.Now())
                err := c.Save()
                if err != nil {
                        c.log.Error("Failed to save cache", zap.Error(err))
                }
        }
		c.timer = time.AfterFunc(c.interval, c.periodicSync) // HLa
}
