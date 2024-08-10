package effect

import "github.com/kercylan98/minotaur/engine/vivid"

func NewManager(configurator ...ManagerConfigurator) *Manager {
	m := &Manager{
		config:     newManagerConfiguration(),
		attributes: make(map[vivid.ActorRef]*Attributes),
	}

	for _, c := range configurator {
		c.Configure(m.config)
	}

	return m
}

// Manager 效果管理器
type Manager struct {
	config     *ManagerConfiguration
	attributes map[vivid.ActorRef]*Attributes      // Actor 属性
	buffs      map[vivid.ActorRef]map[BuffId]*buff // Actor Buff
}

func (m *Manager) ApplyBuff(actor vivid.ActorRef, id BuffId, configurator ...BuffDescriptorConfigurator) {
	descriptor := newBuffDescriptor()
	for _, c := range configurator {
		c.Configure(descriptor)
	}

	buffs, exist := m.buffs[actor]
	if !exist {
		buffs = make(map[BuffId]*buff)
		m.buffs[actor] = buffs
	}
	bf, exist := buffs[id]
	if !exist {
		bf = newBuff(id)
		buffs[id] = bf
	}

	for _, hook := range descriptor.applyHooks {
		m.attributes[actor] = hook(m.GetAttributes(actor))
	}
}

// GetAttributes 获取属性集合
func (m *Manager) GetAttributes(actor vivid.ActorRef) *Attributes {
	attrs, exist := m.attributes[actor]
	if !exist {
		attrs = newAttributes(m)
		m.attributes[actor] = attrs
	}
	return attrs
}

// GetAttribute 获取属性
func (m *Manager) GetAttribute(actor vivid.ActorRef, attributeType AttributeType) Attribute {
	attrs, exist := m.attributes[actor]
	if !exist {
		return m.config.defaultAttributes[attributeType]
	}
	return attrs.Get(attributeType)
}

// SetAttribute 设置属性
func (m *Manager) SetAttribute(actor vivid.ActorRef, attributeType AttributeType, attribute Attribute) {
	attrs, exist := m.attributes[actor]
	if !exist {
		attrs = newAttributes(m)
		m.attributes[actor] = attrs
	}
	attrs.Set(attributeType, attribute)
}
