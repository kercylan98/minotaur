package messages

type LocalPublishRequest struct {
	Topic   string
	Message any
}

func (s *Subscription) SubscriptionId() uint64 {
	return s.Id
}
func (s *Subscription) SubscriptionTopic() string {
	return s.Topic
}
