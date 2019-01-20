package concurrency_pattern

func MultiDone(channels ...<-chan interface{}) <-chan interface{} {
	cnt := len(channels)
	switch cnt {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	mDone := make(chan interface{})
	go func() {
		defer close(mDone)

		switch cnt {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-MultiDone(append(channels[3:], mDone)...):
			}
		}
	}()

	return mDone
}
