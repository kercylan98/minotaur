package reactor

type Reactor[P comparable] struct {
	chs []chan P
}

func (el *Reactor[P]) Send(producer P, event Event) {

}
