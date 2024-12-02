package messagesv1

func (s *Subscription) SubscriptionId() uint64 {
	return s.Id
}
func (s *Subscription) SubscriptionTopic() string {
	return s.Topic
}
