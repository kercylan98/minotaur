package lockstep

type StoppedEventHandle[ClientID comparable, Command any] func(lockstep *Lockstep[ClientID, Command])
